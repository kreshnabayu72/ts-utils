package dto

import "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/analytics_reporting_service/model"

type CreateReaderUserTagLogActivityParam struct {
	Logs     []model.ReaderUserTagLog
	ClientID string
}

type DeleteReaderUserTagLogActivityParam struct {
	Logs     []model.ReaderUserTagLog
	ClientID string
}
