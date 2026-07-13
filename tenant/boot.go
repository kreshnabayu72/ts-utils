// Package tenant answers one question for every service that stores tenant data: which tenants
// exist, and therefore which schemas this service owes.
//
// A service cannot answer that alone. A schema it has never created leaves no trace in its own
// database, so looking inward only ever finds the tenants it already knows about — which is
// exactly the set that does not need anything done. The tenant list lives in
// enterprise-management, and every other service reads it.
//
// That single call collapses three problems into one:
//
//   - Fresh install — no clients yet, so there is nothing to create.
//   - New client — the create-client saga provisions the services that are running; any that were
//     down catch up on their next boot, from this list.
//   - New service — it boots owning nothing, reads the list, and creates every tenant schema from
//     scratch. Clients that predate it are not a special case and need no backfill tool.
//
// Each schema carries its own migration bookkeeping, so a brand-new tenant runs the whole history
// and an existing one runs only the delta. Nothing has to work out which is which.
package tenant

import (
	"context"
	"encoding/json"
	"fmt"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/constant"
	"log/slog"
	"net/http"
	"time"
)

// Migrator is the service's own tenant-provisioning activity — the SAME object it registers
// with Temporal for the client-creation saga to call.
//
// That is deliberate. Provisioning a tenant is one operation with two callers: the saga, when a
// client is created (n=1), and boot, when this service has to catch up on every tenant that
// already exists (n=all). Giving each its own code path would be two implementations of the same
// thing, identical right up until somebody changes one of them.
type Migrator interface {
	// Verify round-trips the migrations on a throwaway schema, so broken SQL is caught before any
	// tenant is touched.
	Verify(ctx context.Context) error
	// MigrateTenant provisions the named tenants, creating any schema that does not exist.
	MigrateTenant(ctx context.Context, clientIDs []string) error
	// MigrateExisting brings up to date only the schemas already here. It is the fallback when the
	// tenant list cannot be reached — those schemas ARE a stale copy of that list, and a stale copy
	// beats nothing.
	MigrateExisting(ctx context.Context) error
}

type config struct {
	// EMSBaseURL is enterprise-management's internal address.
	EMSBaseURL string
	// ServiceToken authenticates this service to EMS's /internal endpoints. Boot is itself an
	// internal caller, so it must present the secret like any other.
	ServiceToken string
	// RetryEvery is the gap between attempts while waiting, and MaxRetryEvery caps it as the wait
	// backs off. Boot order matters — enterprise-management and its database have to be healthy
	// before anything else can learn which tenants exist — and in a cold start everything comes up
	// at once, so waiting is the normal case, not the failure case.
	RetryEvery    time.Duration
	MaxRetryEvery time.Duration
	// GiveUpAfter is how long to keep waiting before starting anyway, in a degraded state. It is
	// not a failure deadline: it is the point at which refusing to serve does more harm than
	// serving with a stale tenant list.
	GiveUpAfter time.Duration
}

// result reports what the service is actually able to do, so the caller can say so out loud
// rather than presenting itself as healthy.
type result struct {
	// Degraded is true when the tenant list could not be reached and this service fell back to the
	// schemas it already had.
	Degraded bool
	// Reason explains the degradation in a form worth putting in front of an operator.
	Reason string
	// Clients is how many tenants are migrated. There is no "how many are behind": the migration
	// run is all-or-nothing, so either every tenant is on this version or none of them moved.
	Clients int
}

func (c config) withDefaults() config {
	if c.RetryEvery == 0 {
		c.RetryEvery = time.Second
	}
	if c.MaxRetryEvery == 0 {
		c.MaxRetryEvery = 30 * time.Second
	}
	if c.GiveUpAfter == 0 {
		c.GiveUpAfter = 2 * time.Minute
	}
	return c
}

// catchUp brings every tenant's schema up to date — the ones this service already has, and the
// ones it has never seen — from the list enterprise-management holds.
//
// It waits for enterprise-management rather than failing fast — services come up in whatever order
// they come up, and enterprise-management being slow is ordinary, not exceptional.
//
// But it does not wait forever, and this is the part worth being careful about. A service that has
// been running for months already holds schemas for every tenant it serves. Refusing to restart
// because enterprise-management happens to be down would take those tenants offline too — for a
// reason that has nothing to do with them. The schemas already here ARE a stale copy of the tenant
// list, and a stale copy beats nothing.
//
// So after GiveUpAfter it starts anyway, migrating what it already has, and reports itself
// DEGRADED. Degraded is not healthy, and the caller must not present it as such:
//
//   - A client created while enterprise-management was unreachable has no schema here. Requests
//     for that tenant will fail on a missing table until this service is restarted with
//     enterprise-management up.
//   - Anything that needs enterprise-management at request time — licensing, client lookup —
//     fails for as long as it stays down.
//
// The one thing this must never do is start quietly. A service that cannot see the tenant list and
// says nothing is worse than one that refuses to start, because the failure surfaces later, to a
// user, as a missing table.
func catchUp(ctx context.Context, m Migrator, cfg config) (result, error) {
	cfg = cfg.withDefaults()

	// Prove the migrations apply at all before letting them near a tenant.
	//
	// The SQL is shared, so broken SQL is broken everywhere: run it against real schemas and every
	// tenant is attempted, every tenant fails, and every tenant is left marked dirty — which is
	// deliberately manual to clear, one schema at a time. A scratch schema catches that for free,
	// and no tenant is touched.
	//
	// Failing the boot here is right, and it is the one case where it is: the migrations cannot
	// apply to ANY schema, so there is no healthy tenant to keep serving and nothing to degrade
	// into. This is a broken deploy, and it should stop like one.
	if err := m.Verify(ctx); err != nil {
		return result{}, fmt.Errorf("tenant: refusing to migrate any tenant: %w", err)
	}

	clientIDs, err := waitForClients(ctx, cfg)
	if err != nil {
		if ctx.Err() != nil {
			// Told to shut down while waiting. Not a failure — just stop.
			return result{}, err
		}

		// Enterprise-management never answered. Fall back to the schemas already here, so the
		// tenants this service has been serving all along keep working.
		reason := fmt.Sprintf(
			"tenant list unavailable: enterprise-management (%s) unreachable for %s. "+
				"Serving existing tenants only — clients created since the last successful start have no schema here, "+
				"and anything requiring enterprise-management (licensing, client lookup) will fail: %v",
			cfg.EMSBaseURL, cfg.GiveUpAfter, err)

		slog.ErrorContext(ctx, "tenant: DEGRADED — starting without the tenant list", "reason", reason)

		if err := m.MigrateExisting(ctx); err != nil {
			// The fallback failed too. There is nothing left to serve with.
			return result{}, fmt.Errorf("tenant: degraded fallback failed, cannot serve any tenant: %w", err)
		}

		return result{Degraded: true, Reason: reason}, nil
	}

	if len(clientIDs) == 0 {
		// A fresh platform. Not an error and not a warning — there is simply nothing to do until
		// somebody creates a client.
		slog.InfoContext(ctx, "tenant: no clients yet, nothing to provision")
		return result{}, nil
	}

	if err := m.MigrateTenant(ctx, clientIDs); err != nil {
		// The migration run is atomic across every tenant, so this means NO tenant moved — they are
		// all still on the version they were already on, and none is half-migrated.
		//
		// That is why this stops the boot instead of degrading. There is no healthy subset to keep
		// serving: the code in this binary expects the new schema and no tenant has it. Starting
		// would mean failing every request from every tenant while claiming to be up.
		return result{}, fmt.Errorf("tenant: migration failed, no tenant was migrated: %w", err)
	}

	slog.InfoContext(ctx, "tenant: schemas up to date", "clients", len(clientIDs))
	return result{Clients: len(clientIDs)}, nil
}

// waitForClients polls enterprise-management until it answers, the deadline passes, or the service
// is told to shut down.
func waitForClients(ctx context.Context, cfg config) ([]string, error) {
	deadline := time.Now().Add(cfg.GiveUpAfter)
	wait := cfg.RetryEvery
	var lastErr error

	for attempt := 1; ; attempt++ {
		ids, err := fetchClients(ctx, cfg.EMSBaseURL, cfg.ServiceToken)
		if err == nil {
			if attempt > 1 {
				slog.InfoContext(ctx, "tenant: enterprise-management reachable", "attempts", attempt)
			}
			return ids, nil
		}
		lastErr = err

		if time.Now().After(deadline) {
			return nil, lastErr
		}

		slog.WarnContext(ctx, "tenant: waiting for enterprise-management",
			"url", cfg.EMSBaseURL, "attempt", attempt, "retryIn", wait, "err", err)

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(wait):
		}

		// Back off, but keep checking often enough that an enterprise-management which appears is
		// picked up promptly rather than sitting unused.
		if wait < cfg.MaxRetryEvery {
			wait *= 2
			if wait > cfg.MaxRetryEvery {
				wait = cfg.MaxRetryEvery
			}
		}
	}
}

func fetchClients(ctx context.Context, baseURL, serviceToken string) ([]string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL+"/internal/clients", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(constant.HEADER_SERVICE_TOKEN, serviceToken)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %d", resp.StatusCode)
	}

	var body struct {
		Data struct {
			ClientIDs []string `json:"clientIds"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, fmt.Errorf("decode: %w", err)
	}
	return body.Data.ClientIDs, nil
}

// Start is the whole tenant-provisioning boot step, in one call.
//
// A service's app wiring should say WHAT happens, not how. What happens is: this service's tenant
// schemas are brought up to date before it serves anything, and it keeps trying if it cannot.
// Everything below that — verifying the migrations on a scratch schema, waiting for
// enterprise-management, degrading rather than dying, healing in the background — is this
// package's business, and repeating it in three app.go files would be three chances to get it
// subtly different.
//
// Returns the live Status, which the caller serves on /health. It is the one thing the app does
// need to hold: degradation has to be visible, and only the app knows where to show it.
func Start(ctx context.Context, m Migrator, emsBaseURL, serviceToken string) (*Status, error) {
	cfg := config{EMSBaseURL: emsBaseURL, ServiceToken: serviceToken}

	res, err := catchUp(ctx, m, cfg)
	if err != nil {
		return nil, err
	}

	status := newStatus(res)

	// Recovering from a degraded start happens on its own. Reconcile returns immediately when the
	// start was healthy, so this costs nothing in the normal case.
	//
	// Deliberately NOT the boot ctx: that one is cancelled once wiring is done, and this has to
	// outlive it. It ends when the process does.
	go status.reconcile(context.WithoutCancel(ctx), m, cfg)

	return status, nil
}
