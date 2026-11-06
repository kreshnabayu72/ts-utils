// Package dto defines the analytics-reporting-service activity parameters for the
// admin-console Hardware Stock log surface.
package dto

import "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/analytics_reporting_service/model"

// CreateHardwareLogsActivityParam appends one Hardware Log event per row.
type CreateHardwareLogsActivityParam struct {
	ClientID     string
	HardwareLogs []model.HardwareLog
}

// DeleteHardwareLogsActivityParam removes every Hardware Log event of the listed units.
type DeleteHardwareLogsActivityParam struct {
	ClientID string
	UnitIDs  []string
}

// CreateHardwareStockEventLogActivityParam appends one Add Stock or Delete Stock session.
type CreateHardwareStockEventLogActivityParam struct {
	ClientID      string
	StockEventLog model.HardwareStockEventLog
}
