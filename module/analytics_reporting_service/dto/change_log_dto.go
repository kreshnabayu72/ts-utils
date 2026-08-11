// Package dto defines the analytics-reporting-service activity parameters.
package dto

import "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/analytics_reporting_service/model"

type CreateChangeLogActivityParam struct {
	ChangeLogs []model.Changelog
	ClientID   string
}

type DeleteChangeLogActivityParam struct {
	ChangeLogs []model.Changelog
	ClientID   string
}
