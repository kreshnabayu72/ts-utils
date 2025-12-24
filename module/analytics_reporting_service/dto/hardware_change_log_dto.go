package dto

import "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/analytics_reporting_service/model"

type CreateHardwareChangeLogActivityParam struct {
	ChangeLogs []model.HardwareChangeLog
	ClientID   string
	// Plan       string
}

type DeleteHardwareChangeLogActivityParam struct {
	ChangeLogs []model.HardwareChangeLog
	ClientID   string
	// Plan         string
}
