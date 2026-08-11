package dto

import "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/analytics_reporting_service/model"

type CreateSessionLogActivityParam struct {
	SessionLogs []model.SessionLog
	ClientID    string
}

type DeleteSessionLogActivityParam struct {
	SessionLogs []model.SessionLog
	ClientID    string
}
