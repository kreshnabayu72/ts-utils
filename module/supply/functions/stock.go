package functions

import (
	"gitea.qwertysystem.net/BETS/ts-utils/entity"
	supply_model "gitea.qwertysystem.net/BETS/ts-utils/module/supply/model"
)

// AddStockStatus adds a new stock status to an existing stock status,
// validate the transition according to the rules and returns the result
func AddStockStatus(s []supply_model.StockStatus, statuses ...supply_model.StockStatus) ([]supply_model.StockStatus, error) {
	type addStockStatusTransition struct {
		Position string // Possible values: "right", "left"
		Status   supply_model.StockStatus
	}

	ALLOWED_ADD_TRANSITIONS := entity.MakeMap[[]supply_model.StockStatus, []addStockStatusTransition]()

	ALLOWED_ADD_TRANSITIONS.Set(
		[]supply_model.StockStatus{
			supply_model.StockStatusAvailable,
		},
		[]addStockStatusTransition{},
	)

	ALLOWED_ADD_TRANSITIONS.Set(
		[]supply_model.StockStatus{
			supply_model.StockStatusOnCheckoutProcess,
		},
		[]addStockStatusTransition{},
	)

	ALLOWED_ADD_TRANSITIONS.Set(
		[]supply_model.StockStatus{
			supply_model.StockStatusOnTransfer,
		},
		[]addStockStatusTransition{},
	)

	ALLOWED_ADD_TRANSITIONS.Set(
		[]supply_model.StockStatus{
			supply_model.StockStatusInStockOpname,
		},
		[]addStockStatusTransition{
			{
				Position: "right",
				Status:   supply_model.StockStatusTAGReported,
			},
		},
	)

	ALLOWED_ADD_TRANSITIONS.Set(
		[]supply_model.StockStatus{
			supply_model.StockStatusDefect,
		},
		[]addStockStatusTransition{
			{
				Position: "right",
				Status:   supply_model.StockStatusWaitingForApproval,
			},
			{
				Position: "right",
				Status:   supply_model.StockStatusTAGReported,
			},
			{
				Position: "left",
				Status:   supply_model.StockStatusInStockOpname,
			},
			{
				Position: "left",
				Status:   supply_model.StockStatusOnDisposalProcess,
			},
			{
				Position: "left",
				Status:   supply_model.StockStatusInSupplierReturnProcess,
			},
		},
	)

	ALLOWED_ADD_TRANSITIONS.Set(
		[]supply_model.StockStatus{
			supply_model.StockStatusMissing,
		},
		[]addStockStatusTransition{
			{
				Position: "left",
				Status:   supply_model.StockStatusInStockOpname,
			},
			{
				Position: "left",
				Status:   supply_model.StockStatusOnDisposalProcess,
			},
			{
				Position: "right",
				Status:   supply_model.StockStatusTAGReported,
			},
			{
				Position: "right",
				Status:   supply_model.StockStatusNotPairedYet,
			},
		},
	)

	ALLOWED_ADD_TRANSITIONS.Set(
		[]supply_model.StockStatus{
			supply_model.StockStatusExpired,
		},
		[]addStockStatusTransition{
			{
				Position: "right",
				Status:   supply_model.StockStatusWaitingForApproval,
			},
			{
				Position: "right",
				Status:   supply_model.StockStatusTAGReported,
			},
			{
				Position: "left",
				Status:   supply_model.StockStatusInStockOpname,
			},
			{
				Position: "left",
				Status:   supply_model.StockStatusOnDisposalProcess,
			},
			{
				Position: "left",
				Status:   supply_model.StockStatusInSupplierReturnProcess,
			},
		},
	)

	ALLOWED_ADD_TRANSITIONS.Set(
		[]supply_model.StockStatus{
			supply_model.StockStatusUnplaced,
		},
		[]addStockStatusTransition{
			{
				Position: "left",
				Status:   supply_model.StockStatusOnCheckoutProcess,
			},
			{
				Position: "left",
				Status:   supply_model.StockStatusInSupplierReturnProcess,
			},
			{
				Position: "right",
				Status:   supply_model.StockStatusExpired,
			},
		},
	)

	ALLOWED_ADD_TRANSITIONS.Set(
		[]supply_model.StockStatus{
			supply_model.StockStatusTAGReported,
		},
		[]addStockStatusTransition{
			{
				Position: "left",
				Status:   supply_model.StockStatusExpired,
			},
		},
	)

	ALLOWED_ADD_TRANSITIONS.Set(
		[]supply_model.StockStatus{
			supply_model.StockStatusNotPairedYet,
		},
		[]addStockStatusTransition{
			{
				Position: "left",
				Status:   supply_model.StockStatusExpired,
			},
			{
				Position: "left",
				Status:   supply_model.StockStatusInStockOpname,
			},
		},
	)

	ALLOWED_ADD_TRANSITIONS.Set(
		[]supply_model.StockStatus{
			supply_model.StockStatusOnCheckoutProcess,
			supply_model.StockStatusUnplaced,
		},
		[]addStockStatusTransition{},
	)

	ALLOWED_ADD_TRANSITIONS.Set(
		[]supply_model.StockStatus{
			supply_model.StockStatusMissing,
			supply_model.StockStatusNotPairedYet,
		},
		[]addStockStatusTransition{
			{
				Position: "left",
				Status:   supply_model.StockStatusInStockOpname,
			},
			{
				Position: "left",
				Status:   supply_model.StockStatusOnDisposalProcess,
			},
		},
	)

	ALLOWED_ADD_TRANSITIONS.Set(
		[]supply_model.StockStatus{
			supply_model.StockStatusMissing,
			supply_model.StockStatusTAGReported,
		},
		[]addStockStatusTransition{
			{
				Position: "left",
				Status:   supply_model.StockStatusInStockOpname,
			},
			{
				Position: "left",
				Status:   supply_model.StockStatusOnDisposalProcess,
			},
		},
	)

	ALLOWED_ADD_TRANSITIONS.Set(
		[]supply_model.StockStatus{
			supply_model.StockStatusDefect,
			supply_model.StockStatusWaitingForApproval,
		},
		[]addStockStatusTransition{
			{
				Position: "left",
				Status:   supply_model.StockStatusInStockOpname,
			},
		},
	)

	ALLOWED_ADD_TRANSITIONS.Set(
		[]supply_model.StockStatus{
			supply_model.StockStatusExpired,
			supply_model.StockStatusWaitingForApproval,
		},
		[]addStockStatusTransition{
			{
				Position: "left",
				Status:   supply_model.StockStatusInStockOpname,
			},
		},
	)

	ALLOWED_ADD_TRANSITIONS.Set(
		[]supply_model.StockStatus{
			supply_model.StockStatusExpired,
			supply_model.StockStatusNotPairedYet,
		},
		[]addStockStatusTransition{
			{
				Position: "left",
				Status:   supply_model.StockStatusInStockOpname,
			},
			{
				Position: "left",
				Status:   supply_model.StockStatusOnDisposalProcess,
			},
		},
	)

	ALLOWED_ADD_TRANSITIONS.Set(
		[]supply_model.StockStatus{
			supply_model.StockStatusDefect,
			supply_model.StockStatusNotPairedYet,
		},
		[]addStockStatusTransition{
			{
				Position: "left",
				Status:   supply_model.StockStatusInStockOpname,
			},
			{
				Position: "left",
				Status:   supply_model.StockStatusOnDisposalProcess,
			},
		},
	)

	ALLOWED_ADD_TRANSITIONS.Set(
		[]supply_model.StockStatus{
			supply_model.StockStatusUnplaced,
			supply_model.StockStatusDefect,
		},
		[]addStockStatusTransition{
			{
				Position: "left",
				Status:   supply_model.StockStatusOnDisposalProcess,
			},
			{
				Position: "left",
				Status:   supply_model.StockStatusInSupplierReturnProcess,
			},
			{
				Position: "right",
				Status:   supply_model.StockStatusWaitingForApproval,
			},
		},
	)

	ALLOWED_ADD_TRANSITIONS.Set(
		[]supply_model.StockStatus{
			supply_model.StockStatusUnplaced,
			supply_model.StockStatusExpired,
		},
		[]addStockStatusTransition{
			{
				Position: "left",
				Status:   supply_model.StockStatusOnDisposalProcess,
			},
			{
				Position: "left",
				Status:   supply_model.StockStatusInSupplierReturnProcess,
			},
			{
				Position: "right",
				Status:   supply_model.StockStatusWaitingForApproval,
			},
		},
	)

	ALLOWED_ADD_TRANSITIONS.Set(
		[]supply_model.StockStatus{
			supply_model.StockStatusUnplaced,
			supply_model.StockStatusDefect,
			supply_model.StockStatusWaitingForApproval,
		},
		[]addStockStatusTransition{},
	)

	ALLOWED_ADD_TRANSITIONS.Set(
		[]supply_model.StockStatus{
			supply_model.StockStatusUnplaced,
			supply_model.StockStatusExpired,
			supply_model.StockStatusWaitingForApproval,
		},
		[]addStockStatusTransition{},
	)

	ALLOWED_ADD_TRANSITIONS.Set(
		[]supply_model.StockStatus{
			supply_model.StockStatusInSupplierReturnProcess,
			supply_model.StockStatusExpired,
		},
		[]addStockStatusTransition{},
	)

	ALLOWED_ADD_TRANSITIONS.Set(
		[]supply_model.StockStatus{
			supply_model.StockStatusInSupplierReturnProcess,
			supply_model.StockStatusUnplaced,
		},
		[]addStockStatusTransition{},
	)

	ALLOWED_ADD_TRANSITIONS.Set(
		[]supply_model.StockStatus{
			supply_model.StockStatusInSupplierReturnProcess,
			supply_model.StockStatusUnplaced,
			supply_model.StockStatusExpired,
		},
		[]addStockStatusTransition{},
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
