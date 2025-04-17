package functions

import (
	"gitea.qwertysystem.net/BETS/ts-utils/entity"
	supply_model "gitea.qwertysystem.net/BETS/ts-utils/module/supply/model"
)

// AddStockStatus adds a new stock status to an existing stock statuses,
// validate the transition according to the rules and returns the result
func AddStockStatus(s []supply_model.StockStatus, statuses ...supply_model.StockStatus) ([]supply_model.StockStatus, error) {
	ALLOWED_ADD_TRANSITIONS := entity.MakeMap[[]supply_model.StockStatus, []supply_model.StockStatus]()

	ALLOWED_ADD_TRANSITIONS.Set(
		[]supply_model.StockStatus{
			supply_model.StockStatusAvailable,
		},
		[]supply_model.StockStatus{
			supply_model.StockStatusOnCheckoutProcess,
			supply_model.StockStatusOnTransfer,
			supply_model.StockStatusInStockOpname,
			supply_model.StockStatusDefect,
			supply_model.StockStatusMissing,
			supply_model.StockStatusExpired,
			supply_model.StockStatusUnplaced,
			supply_model.StockStatusTAGReported,
			supply_model.StockStatusNotPairedYet,
		},
	)

	ALLOWED_ADD_TRANSITIONS.Set(
		[]supply_model.StockStatus{
			supply_model.StockStatusOnCheckoutProcess,
		},
		[]supply_model.StockStatus{
			supply_model.StockStatusAvailable,
			supply_model.StockStatusUnplaced,
		},
	)

	ALLOWED_ADD_TRANSITIONS.Set(
		[]supply_model.StockStatus{
			supply_model.StockStatusOnTransfer,
		},
		[]supply_model.StockStatus{
			supply_model.StockStatusAvailable,
		},
	)

	// TODO(refan): Programmatically adds status based on the transition rules
	s = append(s, s...)
	return s, nil
}

// SetStockStatus sets a new stock status,
// validate the transition according to the rules based from the existing stock statuses and returns the result
func SetStockStatus(s []supply_model.StockStatus, statuses ...supply_model.StockStatus) ([]supply_model.StockStatus, error) {
	// TODO(refan): Programmatically sets status ONLY if the transition rules allows it
	return statuses, nil
}

// RemoveStockStatus removes a stock status,
// validate the transition according to the rules based from the existing stock statuses and returns the result
func RemoveStockStatus(s []supply_model.StockStatus, statuses ...supply_model.StockStatus) ([]supply_model.StockStatus, error) {
	// TODO(refan): Programmatically removes status based on the transition rules
	return s, nil
}
