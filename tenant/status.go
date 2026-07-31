package tenant

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"log/slog"
	"sync"
	"time"
)

// Status is a service's live answer to "can I actually do my job right now".
//
// It exists so degradation is VISIBLE. A service that cannot see the tenant list and reports
// itself healthy is the worst of the available outcomes: the failure surfaces later, to a user, as
// a missing table — long after the cause is gone from the logs.
//
// Safe for concurrent use: the health handler reads it while the reconciler writes it.
type Status struct {
	mu       sync.RWMutex
	degraded bool
	reason   string
	since    time.Time
}

func newStatus(r result) *Status {
	s := &Status{}
	s.set(r)
	return s
}

func (s *Status) set(r result) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if r.Degraded && !s.degraded {
		s.since = time.Now()
	}
	s.degraded = r.Degraded
	s.reason = r.Reason
	if !r.Degraded {
		s.since = time.Time{}
	}
}

// Health is the /health handler. Mount it: fiberApp.Get("/health", tenantStatus.Health).
//
// It lives here, not in each service's wiring, because all three would otherwise carry the same
// ten lines — and the day one of them decided degraded should be a 200 "for the load balancer",
// the platform would quietly stop reporting its own degradation.
//
// 503 when degraded, and that is the point: a service that cannot see the tenant list, or has
// tenants it could not migrate, is NOT healthy. Saying so is what stops the failure from surfacing
// later, to a user, as a missing table.
func (s *Status) Health(c fiber.Ctx) error {
	degraded, reason, since := s.snapshot()
	if !degraded {
		return c.JSON(fiber.Map{"status": "ok"})
	}
	return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
		"status":      "degraded",
		"reason":      reason,
		"degradedFor": time.Since(since).Truncate(time.Second).String(),
	})
}

func (s *Status) snapshot() (degraded bool, reason string, since time.Time) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.degraded, s.reason, s.since
}

// Reconcile keeps trying, in the background, to recover from a degraded start.
//
// Without this, a service that started while enterprise-management was down stays degraded until
// somebody notices and restarts it. That is not graceful degradation, it is a deferred outage: the
// service recovers only when a human intervenes, and the tenants created in the meantime stay
// broken until they do.
//
// With it, the service converges on its own. Enterprise-management comes back — in a minute or an
// hour — and the next tick picks up the tenant list, creates the schemas that were missed, and
// clears the degraded flag. Nobody is paged, and no user notices anything beyond the window in
// which their own tenant did not yet exist here.
//
// It is a no-op when the service started healthy: there is nothing to recover from, and polling
// enterprise-management forever for no reason is just noise.
//
// Runs until ctx is cancelled, so shutdown stops it.
func (s *Status) reconcile(ctx context.Context, p Migrator, cfg config) {
	if degraded, _, _ := s.snapshot(); !degraded {
		return
	}
	cfg = cfg.withDefaults()

	ticker := time.NewTicker(cfg.MaxRetryEvery)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}

		clientIDs, err := fetchClients(ctx, cfg.EMSBaseURL, cfg.ServiceToken)
		if err != nil {
			// Still down. Stay degraded and stay quiet-ish: this is the expected state while it is
			// down, and logging an error every tick would bury the one that mattered.
			slog.DebugContext(ctx, "tenant: still degraded, enterprise-management unreachable", "err", err)
			continue
		}

		if err := p.MigrateTenant(ctx, clientIDs); err != nil {
			// The list came back but migrating it failed. That is a real fault, and it stays
			// degraded — serving a half-migrated schema would be worse than not serving the tenant.
			slog.ErrorContext(ctx, "tenant: recovered the tenant list but provisioning failed", "err", err)
			continue
		}

		_, _, since := s.snapshot()
		slog.InfoContext(ctx, "tenant: RECOVERED — tenant list reachable, schemas up to date",
			"clients", len(clientIDs), "degradedFor", time.Since(since).Truncate(time.Second))

		s.set(result{Clients: len(clientIDs)})
		return
	}
}
