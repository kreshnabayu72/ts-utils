package dto

import "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/analytics_reporting_service/model"

type CreateHandheldReaderUsageLogActivityParam struct {
	UsageLogs []model.HandheldReaderUsageLog
	ClientID  string
	// Plan       string
}

type DeleteHandheldReaderUsageLogActivityParam struct {
	UsageLogs []model.HandheldReaderUsageLog
	ClientID  string
	// Plan         string
}
