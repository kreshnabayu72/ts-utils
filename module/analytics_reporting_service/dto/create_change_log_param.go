package dto

import "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/analytics_reporting_service/model"

type CreateChangeLogActivityParam struct {
	ChangeLogs []model.ChangeLog
	ClientID   string
	// Plan       string
}

type DeleteChangeLogActivityParam struct {
	ChangeLogs []model.ChangeLog
	ClientID   string
	// Plan         string
}
