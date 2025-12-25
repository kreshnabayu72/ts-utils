package dto

import "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/analytics_reporting_service/model"

type CreateReaderTaggingLogActivityParam struct {
	Logs     []model.ReaderTaggingLog
	ClientID string
}

type DeleteReaderTaggingLogActivityParam struct {
	Logs     []model.ReaderTaggingLog
	ClientID string
}
