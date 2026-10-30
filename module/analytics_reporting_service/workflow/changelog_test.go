package workflow

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/workflow"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/constant"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/analytics_reporting_service/dto"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/analytics_reporting_service/model"
)

type testEnv struct {
	env           *testsuite.TestWorkflowEnvironment
	payload       []dto.CreateChangeLogActivityParam
	auditFailures int
}

func newEnv() *testEnv {
	var s testsuite.WorkflowTestSuite
	e := &testEnv{env: s.NewTestWorkflowEnvironment()}
	e.env.RegisterWorkflowWithOptions(
		func(_ workflow.Context, p dto.CreateChangeLogActivityParam) error {
			e.payload = append(e.payload, p)
			if e.auditFailures > 0 {
				e.auditFailures--
				return errors.New("audit sink unavailable")
			}
			return nil
		},
		workflow.RegisterOptions{Name: constant.AnalyticsReportingRecordChangeLogsWorkflow},
	)
	return e
}

func spawnWF(ctx workflow.Context, logs []model.Changelog) error {
	SpawnChangelogWorkflow(ctx, "acme", "iam-service", logs)
	return nil
}

func TestSpawnChangelogWorkflow(t *testing.T) {
	rows := []model.Changelog{{Object: "Position"}}

	t.Run("no rows spawns nothing", func(t *testing.T) {
		e := newEnv()
		e.env.ExecuteWorkflow(spawnWF, []model.Changelog{})
		if !e.env.IsWorkflowCompleted() || e.env.GetWorkflowError() != nil || len(e.payload) != 0 {
			t.Fatalf("err=%v payload=%v", e.env.GetWorkflowError(), e.payload)
		}
	})
	t.Run("rows reach the child", func(t *testing.T) {
		e := newEnv()
		e.env.ExecuteWorkflow(spawnWF, rows)
		if e.env.GetWorkflowError() != nil || len(e.payload) != 1 || e.payload[0].ClientID != "acme" {
			t.Fatalf("err=%v payload=%v", e.env.GetWorkflowError(), e.payload)
		}
	})
	t.Run("child retries beyond two attempts", func(t *testing.T) {
		e := newEnv()
		e.auditFailures = 3
		e.env.ExecuteWorkflow(func(ctx workflow.Context) error {
			SpawnChangelogWorkflow(ctx, "acme", "iam-service", rows)
			return workflow.Sleep(ctx, time.Minute)
		})
		if e.env.GetWorkflowError() != nil || e.auditFailures != 0 || len(e.payload) != 4 {
			t.Fatalf("err=%v failures=%d attempts=%d", e.env.GetWorkflowError(), e.auditFailures, len(e.payload))
		}
	})
	t.Run("failed start retries after parent cancellation", func(t *testing.T) {
		e := newEnv()
		attempts := 0
		e.env.OnWorkflow(constant.AnalyticsReportingRecordChangeLogsWorkflow, mock.Anything, mock.Anything).
			Return(func(workflow.Context, dto.CreateChangeLogActivityParam) error {
				attempts++
				if attempts == 1 {
					return testsuite.ErrMockStartChildWorkflowFailed
				}
				return nil
			})
		e.env.RegisterDelayedCallback(func() { e.env.CancelWorkflow() }, time.Second)
		e.env.ExecuteWorkflow(spawnWF, rows)
		if e.env.GetWorkflowError() != nil {
			t.Fatal(e.env.GetWorkflowError())
		}
		if attempts < 2 {
			t.Fatal("failed changelog handoff must be retried")
		}
		e.env.AssertExpectations(t)
	})
}
