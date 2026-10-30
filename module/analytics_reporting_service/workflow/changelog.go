// Package workflow holds Temporal helpers for calling analytics-reporting-service.
package workflow

import (
	"go.temporal.io/sdk/workflow"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/constant"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/analytics_reporting_service/dto"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/analytics_reporting_service/model"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/temporal_helper"
)

// SpawnChangelogWorkflow starts the analytics RecordChangeLogs child workflow
// fire-and-forget on taskQueue. It waits only until the child has started;
// a start failure is logged, never returned, so changelogs cannot fail the caller.
func SpawnChangelogWorkflow(ctx workflow.Context, clientID, taskQueue string, logs []model.Changelog) {
	if len(logs) == 0 {
		return
	}
	opts := temporal_helper.FireAndForgetChildWorkflowOptions(ctx, "", taskQueue)
	future := workflow.ExecuteChildWorkflow(
		workflow.WithChildOptions(ctx, opts),
		constant.AnalyticsReportingRecordChangeLogsWorkflow,
		dto.CreateChangeLogActivityParam{ClientID: clientID, ChangeLogs: logs},
	)
	var childExecution workflow.Execution
	if err := future.GetChildWorkflowExecution().Get(ctx, &childExecution); err != nil {
		workflow.GetLogger(ctx).Error("failed to spawn changelog child workflow", "error", err)
	}
}
