package functions

import (
	"errors"
	"fmt"
	"slices"

	"gitea.qwertysystem.net/BETS/ts-utils/v2/functions"
	supply_model "gitea.qwertysystem.net/BETS/ts-utils/v2/module/supply/model"
)

// StockStatusPriority is a struct that contains the priority of a stock status for ordering purposes.
//
// The status will be sorted from the lowest to the highest priority.
type StockStatusPriority = map[int]supply_model.StockStatus

// StockStatusPriorityList will be used as a list of stock status priorities.
//
// 8, 10 and 12 are reserved for manually inserting "Waiting for Approval" status.
var StockStatusPriorityList = StockStatusPriority{
	0:  supply_model.StockStatusAvailable,
	1:  supply_model.StockStatusInStockOpname,
	2:  supply_model.StockStatusOnCheckoutProcess,
	3:  supply_model.StockStatusOnTransfer,
	4:  supply_model.StockStatusInSupplierReturnProcess,
	5:  supply_model.StockStatusOnDisposalProcess,
	6:  supply_model.StockStatusUnplaced,
	7:  supply_model.StockStatusMissing,
	8:  supply_model.StockStatusWaitingForApproval,
	9:  supply_model.StockStatusDefect,
	10: supply_model.StockStatusWaitingForApproval,
	11: supply_model.StockStatusExpired,
	12: supply_model.StockStatusWaitingForApproval,
	13: supply_model.StockStatusTAGReported,
	14: supply_model.StockStatusNotPairedYet,
	15: supply_model.StockStatusDisposed,
	16: supply_model.StockStatusCheckedOut,
	17: supply_model.StockStatusExchanged,
	18: supply_model.StockStatusReturned,
}

// AddStockStatus adds a new stock status to an existing stock status,
//
// Existing stock status will not be further added to prevent duplication.
// If the stock status initially contains "Available", it will be overwritten by the next-in-line new status
func AddStockStatus(s []supply_model.StockStatus, statuses ...supply_model.StockStatus) []supply_model.StockStatus {
	addedStatus := []supply_model.StockStatus{}

	if !functions.Contains(addedStatus, supply_model.StockStatusAvailable) {
		addedStatus = append(addedStatus, s...)
	}

	orderedStatusKey := []int{}
	orderedStatusValue := []supply_model.StockStatus{}

	for _, status := range statuses {
		// Skip appending the status if it already exists
		if functions.Contains(addedStatus, status) {
			continue
		}

		addedStatus = append(addedStatus, status)
	}

	for _, status := range addedStatus {
		for k, v := range StockStatusPriorityList {
			// Manually insert the "Waiting for Approval" status according to the currently available status.
			if status == supply_model.StockStatusWaitingForApproval {
				if functions.Contains(addedStatus, supply_model.StockStatusMissing) {
					orderedStatusKey = append(orderedStatusKey, 8)
					break
				} else if functions.Contains(addedStatus, supply_model.StockStatusDefect) {
					orderedStatusKey = append(orderedStatusKey, 10)
					break
				} else {
					orderedStatusKey = append(orderedStatusKey, 12)
					break
				}
			}

			if v == status {
				orderedStatusKey = append(orderedStatusKey, k)
				break
			}
		}
	}

	slices.Sort(orderedStatusKey)

	for _, k := range orderedStatusKey {
		orderedStatusValue = append(orderedStatusValue, StockStatusPriorityList[k])
	}

	return orderedStatusValue
}

// SetStockStatus attempts to set a stock status from an existing stock status.
//
// Will return an error if the rule doesn't allow the status to be set.
// You can use this error for an if case to determine whether to add / remove stock status manually.
func SetStockStatus(s []supply_model.StockStatus, status supply_model.StockStatus) ([]supply_model.StockStatus, error) {
	currentStatus := make([]supply_model.StockStatus, len(s))
	copy(currentStatus, s)

	commonMessage := "you might want to add / remove stock status manually."

	switch status {
	case
		supply_model.StockStatusAvailable,
		supply_model.StockStatusOnCheckoutProcess,
		supply_model.StockStatusOnTransfer,
		supply_model.StockStatusDefect,
		supply_model.StockStatusMissing,
		supply_model.StockStatusExpired,
		supply_model.StockStatusUnplaced,
		supply_model.StockStatusNotPairedYet,
		supply_model.StockStatusInStockOpname,
		supply_model.StockStatusDisposed,
		supply_model.StockStatusCheckedOut,
		supply_model.StockStatusExchanged,
		supply_model.StockStatusReturned:

		if len(currentStatus) <= 1 {
			return []supply_model.StockStatus{status}, nil
		}

		return currentStatus, errors.New(fmt.Sprintf("can't set \"%s\" status when there are currently more than one status, %s", status, commonMessage))
	}

	return currentStatus, errors.New(fmt.Sprintf("can't set \"%s\" directly, %s", status, commonMessage))
}

// RemoveStockStatus removes an existing stock status and preserves the sorting order.
//
// Skip deleting stock status if it doesn't exist in the list. If the resulting list is empty,
// this function will automatically add "Available" to the result.
func RemoveStockStatus(s []supply_model.StockStatus, statuses ...supply_model.StockStatus) []supply_model.StockStatus {
	removedStatus := make([]supply_model.StockStatus, len(s))
	copy(removedStatus, s)

	for _, status := range statuses {
		i := slices.Index(removedStatus, status)
		// Skip if the desired status doesn't exist in the list.
		if i == -1 {
			continue
		}

		removedStatus = append(removedStatus[:i], removedStatus[i+1:]...)
	}

	if len(removedStatus) == 0 {
		return []supply_model.StockStatus{
			supply_model.StockStatusAvailable,
		}
	}

	return removedStatus
}
