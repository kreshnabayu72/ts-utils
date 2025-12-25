package dto

import "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/analytics_reporting_service/model"

type CreateTagEventLogAuditActivityParam struct {
	Logs     []model.TagEventLogAudit
	ClientID string
}

type DeleteTagEventLogAuditActivityParam struct {
	Logs     []model.TagEventLogAudit
	ClientID string
}
