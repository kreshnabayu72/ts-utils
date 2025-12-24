package dto

import "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/analytics_reporting_service/model"

type CreateIOTReaderUsageLogActivityParam struct {
	UsageLogs []model.IOTReaderUsageLog
	ClientID  string
	// Plan       string
}

type DeleteIOTReaderUsageLogActivityParam struct {
	UsageLogs []model.IOTReaderUsageLog
	ClientID  string
	// Plan         string
}
