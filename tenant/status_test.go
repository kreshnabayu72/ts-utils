package tenant

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

// A degraded service must heal itself. Staying degraded until a human restarts it is not graceful
// degradation — it is an outage waiting for somebody to notice.
func TestReconcile_RecoversWhenEnterpriseManagementReturns(t *testing.T) {
	var up atomic.Bool

	ems := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		if !up.Load() {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		_, _ = w.Write([]byte(`{"data":{"clientIds":["acme","globex"]}}`))
	}))
	defer ems.Close()

	cfg := config{
		EMSBaseURL:    ems.URL,
		RetryEvery:    time.Millisecond,
		MaxRetryEvery: 5 * time.Millisecond,
		GiveUpAfter:   20 * time.Millisecond,
	}

	// Boot while it is down: the service starts, degraded, on the schemas it already had.
	p := &fakeMigrator{}
	res, err := catchUp(context.Background(), p, cfg)
	if err != nil {
		t.Fatalf("Provision: %v", err)
	}
	if !res.Degraded {
		t.Fatal("should have started degraded")
	}

	status := newStatus(res)
	if degraded, _, _ := status.snapshot(); !degraded {
		t.Fatal("status does not report the degradation, so nothing can surface it")
	}

	// Enterprise-management comes back a moment later.
	go func() {
		time.Sleep(10 * time.Millisecond)
		up.Store(true)
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	status.reconcile(ctx, p, cfg)

	if degraded, _, _ := status.snapshot(); degraded {
		t.Error("enterprise-management came back but the service is still degraded")
	}
	if len(p.got) != 2 {
		t.Errorf("recovered but provisioned %v — the tenants missed while it was down are still missing", p.got)
	}
}

// A service that started healthy has nothing to reconcile, and must not sit polling forever.
func TestReconcile_NoopWhenHealthy(t *testing.T) {
	status := newStatus(result{Clients: 3})

	done := make(chan struct{})
	go func() {
		status.reconcile(context.Background(), &fakeMigrator{}, config{EMSBaseURL: "http://127.0.0.1:1"})
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("Reconcile kept polling despite a healthy start")
	}
}
