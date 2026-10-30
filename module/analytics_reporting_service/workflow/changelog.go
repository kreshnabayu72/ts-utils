// Package workflow holds Temporal helpers for calling analytics-reporting-service.
package workflow

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/constant"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/analytics_reporting_service/dto"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/analytics_reporting_service/model"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/temporal_helper"
)

// SpawnChangelogWorkflow starts the analytics RecordChangeLogs child workflow
// fire-and-forget (ABANDON) on taskQueue and never returns an error, so a
// changelog cannot fail or roll back the caller's committed write.
//
// The handoff is durable: the start runs on a disconnected context, so caller
// cancellation does not drop it, and a failed start is retried every minute
// until the child starts. The child's own retry policy keeps redelivering the
// rows across an analytics-reporting outage.
func SpawnChangelogWorkflow(ctx workflow.Context, clientID, taskQueue string, logs []model.Changelog) {
	if len(logs) == 0 {
		return
	}
	ctx, _ = workflow.NewDisconnectedContext(ctx)
	opts := temporal_helper.FireAndForgetChildWorkflowOptions(ctx, "", taskQueue)
	opts.RetryPolicy = &temporal.RetryPolicy{
		InitialInterval: time.Second,
		MaximumInterval: time.Minute,
	}
	childCtx := workflow.WithChildOptions(ctx, opts)
	param := dto.CreateChangeLogActivityParam{ClientID: clientID, ChangeLogs: logs}
	for {
		future := workflow.ExecuteChildWorkflow(childCtx, constant.AnalyticsReportingRecordChangeLogsWorkflow, param)
		var childExecution workflow.Execution
		err := future.GetChildWorkflowExecution().Get(ctx, &childExecution)
		if err == nil {
			return
		}
		workflow.GetLogger(ctx).Error("failed to spawn changelog child workflow; retrying", "error", err)
		_ = workflow.Sleep(ctx, time.Minute)
	}
}
