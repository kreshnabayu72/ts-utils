package dto

import "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/analytics_reporting_service/model"

// CreateTagLogsActivityParam appends one TAG Log event per row.
type CreateTagLogsActivityParam struct {
	ClientID string
	TagLogs  []model.TagLog
}

// DeleteTagLogsActivityParam removes every TAG Log event of the listed TAG ids.
type DeleteTagLogsActivityParam struct {
	ClientID string
	TagIDs   []string
}

// UpsertTagActivateLogActivityParam upserts one Activate/Delete session row by sessionId.
type UpsertTagActivateLogActivityParam struct {
	ClientID       string
	TagActivateLog model.TagActivateLog
}

// UpsertTagAuditLogActivityParam upserts one Audit session row by sessionId.
type UpsertTagAuditLogActivityParam struct {
	ClientID    string
	TagAuditLog model.TagAuditLog
}
