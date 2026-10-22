package dto

import "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/analytics_reporting_service/model"

type CreateHandheldReaderActivateLogActivityParam struct {
	ActivateLogs []model.HandheldReaderActivateLog
	ClientID     string
}

type DeleteHandheldReaderActivateLogActivityParam struct {
	ActivateLogs []model.HandheldReaderActivateLog
	ClientID     string
}

type CreateHandheldReaderUsageLogActivityParam struct {
	UsageLogs []model.HandheldReaderUsageLog
	ClientID  string
}

type DeleteHandheldReaderUsageLogActivityParam struct {
	UsageLogs []model.HandheldReaderUsageLog
	ClientID  string
}
