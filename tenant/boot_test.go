package tenant

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

type fakeMigrator struct {
	got           []string
	fellBackToAll bool
	// failFor names a schema whose migration blows up, as a tenant whose own data conflicts with a
	// new constraint would.
	failFor string
	// verifyErr makes the dry-run fail, as broken migration SQL would.
	verifyErr error
	verified  bool
}

func (f *fakeMigrator) Verify(_ context.Context) error {
	f.verified = true
	return f.verifyErr
}

func (f *fakeMigrator) MigrateTenant(_ context.Context, schemas []string) error {
	// Mirrors the real runner: one transaction across every schema. If any of them fails, NONE of
	// them moved — not the ones before it in the list, and not the ones after.
	for _, s := range schemas {
		if s == f.failFor {
			return errors.New("duplicate key violates the new unique constraint")
		}
	}
	f.got = append(f.got, schemas...)
	return nil
}

func (f *fakeMigrator) MigrateExisting(_ context.Context) error {
	f.fellBackToAll = true
	return nil
}

func emsServing(clientIDs ...string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		body := `{"data":{"clientIds":[`
		for i, id := range clientIDs {
			if i > 0 {
				body += ","
			}
			body += `"` + id + `"`
		}
		body += `]}}`
		_, _ = w.Write([]byte(body))
	}))
}

func fastConfig(url string) config {
	return config{
		EMSBaseURL:    url,
		RetryEvery:    time.Millisecond,
		MaxRetryEvery: 2 * time.Millisecond,
		GiveUpAfter:   100 * time.Millisecond,
	}
}

// Every tenant the platform knows about is brought up to date — including ones this service has
// never seen, which is the whole reason the list comes from elsewhere.
func TestCatchUp_MigratesEveryClient(t *testing.T) {
	ems := emsServing("acme", "globex")
	defer ems.Close()

	m := &fakeMigrator{}
	res, err := catchUp(context.Background(), m, fastConfig(ems.URL))
	if err != nil {
		t.Fatalf("catchUp: %v", err)
	}
	if res.Degraded {
		t.Error("reported degraded despite reaching enterprise-management")
	}
	if len(m.got) != 2 || m.got[0] != "acme" || m.got[1] != "globex" {
		t.Errorf("migrated %v, want [acme globex]", m.got)
	}
}

// A fresh platform has no clients. Not an error, and nothing to do.
func TestCatchUp_NoClientsIsFine(t *testing.T) {
	ems := emsServing()
	defer ems.Close()

	m := &fakeMigrator{}
	if _, err := catchUp(context.Background(), m, fastConfig(ems.URL)); err != nil {
		t.Fatalf("a platform with no clients must boot cleanly: %v", err)
	}
	if len(m.got) != 0 {
		t.Errorf("migrated %v on an empty platform", m.got)
	}
}

// Services come up in whatever order they come up. If enterprise-management is not ready yet, this
// one waits and picks the tenants up when it appears.
func TestCatchUp_WaitsForEnterpriseManagement(t *testing.T) {
	var ready atomic.Bool

	ems := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		if !ready.Load() {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		_, _ = w.Write([]byte(`{"data":{"clientIds":["acme"]}}`))
	}))
	defer ems.Close()

	go func() {
		time.Sleep(20 * time.Millisecond)
		ready.Store(true)
	}()

	m := &fakeMigrator{}
	if _, err := catchUp(context.Background(), m, fastConfig(ems.URL)); err != nil {
		t.Fatalf("catchUp should have waited and succeeded: %v", err)
	}
	if len(m.got) != 1 || m.got[0] != "acme" {
		t.Errorf("migrated %v, want [acme] once enterprise-management came up", m.got)
	}
}

// Waiting must still answer to a shutdown, rather than hanging on an enterprise-management that may
// never arrive.
func TestCatchUp_ShutdownWhileWaitingExitsCleanly(t *testing.T) {
	cfg := fastConfig("http://127.0.0.1:1") // nothing is listening
	cfg.GiveUpAfter = time.Hour             // the cancellation must stop it, not the deadline

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(20 * time.Millisecond)
		cancel()
	}()

	done := make(chan error, 1)
	go func() {
		_, err := catchUp(ctx, &fakeMigrator{}, cfg)
		done <- err
	}()

	select {
	case err := <-done:
		if err == nil {
			t.Fatal("a cancelled wait must report that it was stopped, not that it succeeded")
		}
	case <-time.After(2 * time.Second):
		t.Fatal("catchUp ignored the shutdown signal and kept waiting")
	}
}

// A service that has been running for months already holds schemas for the tenants it serves.
// Refusing to restart because enterprise-management is down would take those tenants offline for a
// reason that has nothing to do with them. So it starts — on what it already has — and says clearly
// that it is degraded.
func TestCatchUp_DegradesRatherThanRefusingToStart(t *testing.T) {
	cfg := fastConfig("http://127.0.0.1:1") // enterprise-management is down

	m := &fakeMigrator{}
	res, err := catchUp(context.Background(), m, cfg)
	if err != nil {
		t.Fatalf("an unreachable enterprise-management must not stop the service from starting: %v", err)
	}
	if !res.Degraded {
		t.Error("started without the tenant list but did not report itself degraded")
	}
	if res.Reason == "" {
		t.Error("degraded with no reason — an operator has nothing to act on")
	}
	if !m.fellBackToAll {
		t.Error("did not fall back to the schemas it already had, so its existing tenants are stranded")
	}
}

// One tenant's data conflicting with a migration takes the WHOLE run down. Every other tenant stays
// exactly where it was.
//
// That is what the atomic run buys. Migrating the healthy ones and leaving the odd one behind puts
// the platform on two schema versions at once — and this binary can only have been written for one
// of them, so whoever is on the other is broken regardless. Keeping them "serving" buys nothing and
// hides the split.
//
// It is affordable precisely because rolling back costs nothing here: no down migration runs and no
// data is destroyed. The DDL simply never commits.
func TestCatchUp_OneBadTenantRollsBackEveryTenant(t *testing.T) {
	ems := emsServing("acme", "badcorp", "globex")
	defer ems.Close()

	m := &fakeMigrator{failFor: "badcorp"}
	_, err := catchUp(context.Background(), m, fastConfig(ems.URL))

	if err == nil {
		t.Fatal("a tenant that could not migrate must stop the boot, not be quietly left behind")
	}
	if len(m.got) != 0 {
		t.Errorf("migrated %v — the run is atomic, so NO tenant may be left on the new version", m.got)
	}
}

// Broken migration SQL is shared, so it is broken for everyone. The dry-run on a scratch schema
// catches it before the real run is even attempted — no tenant is touched, and nothing goes dirty.
func TestCatchUp_BadMigrationSQLStopsBeforeTouchingAnyTenant(t *testing.T) {
	ems := emsServing("acme", "globex")
	defer ems.Close()

	m := &fakeMigrator{verifyErr: errors.New(`column "typo" does not exist`)}

	_, err := catchUp(context.Background(), m, fastConfig(ems.URL))
	if err == nil {
		t.Fatal("migrations that cannot apply to an empty schema must stop the boot")
	}
	if !m.verified {
		t.Error("the migrations were never verified against a scratch schema")
	}
	if len(m.got) != 0 {
		t.Errorf("touched %v — no tenant may be migrated once the SQL is known to be broken", m.got)
	}
}

// The dry-run runs against an EMPTY schema, so it cannot catch a failure that depends on a tenant's
// data. Those surface at the real run — and take it down with them, which is what keeps every tenant
// on the same version.
func TestCatchUp_VerifyDoesNotHideDataSpecificFailures(t *testing.T) {
	ems := emsServing("acme", "badcorp", "globex")
	defer ems.Close()

	// Verify passes: the SQL is fine. badcorp's data is not.
	m := &fakeMigrator{failFor: "badcorp"}

	_, err := catchUp(context.Background(), m, fastConfig(ems.URL))
	if err == nil {
		t.Fatal("a data-specific failure must still stop the run")
	}
	if !m.verified {
		t.Error("the dry-run did not happen")
	}
}
