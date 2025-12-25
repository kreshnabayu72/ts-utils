package dto

import "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/analytics_reporting_service/model"

type CreateTagEventLogActivityParam struct {
	Logs     []model.TagEventLog
	ClientID string
}

type DeleteTagEventLogActivityParam struct {
	Logs     []model.TagEventLog
	ClientID string
}
