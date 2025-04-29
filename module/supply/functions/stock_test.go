package functions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gitea.qwertysystem.net/BETS/ts-utils/module/supply/functions"
	supply_model "gitea.qwertysystem.net/BETS/ts-utils/module/supply/model"
)

func TestAddStockStatus(t *testing.T) {
	t.Run("successfully add and sort new status", func(t *testing.T) {
		currentStatus := []supply_model.StockStatus{
			supply_model.StockStatusExpired,
		}

		newStatus := functions.AddStockStatus(
			currentStatus,
			supply_model.StockStatusNotPairedYet,
			supply_model.StockStatusDefect,
			supply_model.StockStatusOnDisposalProcess,
		)

		assert.Equal(t, []supply_model.StockStatus{
			supply_model.StockStatusOnDisposalProcess,
			supply_model.StockStatusDefect,
			supply_model.StockStatusExpired,
			supply_model.StockStatusNotPairedYet,
		}, newStatus)
	})

	t.Run("skip adding status already in the old status", func(t *testing.T) {
		currentStatus := []supply_model.StockStatus{
			supply_model.StockStatusExpired,
		}

		newStatus := functions.AddStockStatus(
			currentStatus,
			supply_model.StockStatusNotPairedYet,
			supply_model.StockStatusDefect,
			supply_model.StockStatusOnDisposalProcess,
			supply_model.StockStatusExpired,
		)

		assert.Equal(t, []supply_model.StockStatus{
			supply_model.StockStatusOnDisposalProcess,
			supply_model.StockStatusDefect,
			supply_model.StockStatusExpired,
			supply_model.StockStatusNotPairedYet,
		}, newStatus)
	})

	t.Run("correctly add waiting for approval to the list (defect)", func(t *testing.T) {
		currentStatus := []supply_model.StockStatus{
			supply_model.StockStatusExpired,
		}

		newStatus := functions.AddStockStatus(
			currentStatus,
			supply_model.StockStatusNotPairedYet,
			supply_model.StockStatusDefect,
			supply_model.StockStatusOnDisposalProcess,
			supply_model.StockStatusExpired,
			supply_model.StockStatusWaitingForApproval,
		)

		assert.Equal(t, []supply_model.StockStatus{
			supply_model.StockStatusOnDisposalProcess,
			supply_model.StockStatusDefect,
			supply_model.StockStatusWaitingForApproval,
			supply_model.StockStatusExpired,
			supply_model.StockStatusNotPairedYet,
		}, newStatus)
	})

	t.Run("correctly add waiting for approval to the list (missing)", func(t *testing.T) {
		currentStatus := []supply_model.StockStatus{
			supply_model.StockStatusExpired,
		}

		newStatus := functions.AddStockStatus(
			currentStatus,
			supply_model.StockStatusNotPairedYet,
			supply_model.StockStatusMissing,
			supply_model.StockStatusOnDisposalProcess,
			supply_model.StockStatusExpired,
			supply_model.StockStatusWaitingForApproval,
		)

		assert.Equal(t, []supply_model.StockStatus{
			supply_model.StockStatusOnDisposalProcess,
			supply_model.StockStatusMissing,
			supply_model.StockStatusWaitingForApproval,
			supply_model.StockStatusExpired,
			supply_model.StockStatusNotPairedYet,
		}, newStatus)
	})

	t.Run("correctly add waiting for approval to the list (expired)", func(t *testing.T) {
		currentStatus := []supply_model.StockStatus{
			supply_model.StockStatusExpired,
		}

		newStatus := functions.AddStockStatus(
			currentStatus,
			supply_model.StockStatusNotPairedYet,
			supply_model.StockStatusOnDisposalProcess,
			supply_model.StockStatusWaitingForApproval,
		)

		assert.Equal(t, []supply_model.StockStatus{
			supply_model.StockStatusOnDisposalProcess,
			supply_model.StockStatusExpired,
			supply_model.StockStatusWaitingForApproval,
			supply_model.StockStatusNotPairedYet,
		}, newStatus)
	})
}

func TestRemoveStockStatus(t *testing.T) {
	t.Run("successfully remove status", func(t *testing.T) {
		currentStatus := []supply_model.StockStatus{
			supply_model.StockStatusOnDisposalProcess,
			supply_model.StockStatusDefect,
			supply_model.StockStatusExpired,
			supply_model.StockStatusNotPairedYet,
		}

		newStatus := functions.RemoveStockStatus(
			currentStatus,
			supply_model.StockStatusOnDisposalProcess,
			supply_model.StockStatusExpired,
		)

		assert.Equal(t, []supply_model.StockStatus{
			supply_model.StockStatusDefect,
			supply_model.StockStatusNotPairedYet,
		}, newStatus)
	})

	t.Run("skip and don't error when the desired status to be removed doesn't exist", func(t *testing.T) {
		currentStatus := []supply_model.StockStatus{
			supply_model.StockStatusOnDisposalProcess,
			supply_model.StockStatusDefect,
			supply_model.StockStatusExpired,
			supply_model.StockStatusNotPairedYet,
		}

		newStatus := functions.RemoveStockStatus(
			currentStatus,
			supply_model.StockStatusInStockOpname,
		)

		assert.Equal(t, currentStatus, newStatus)
	})

	t.Run("removing status until empty will automatically return \"Available\"", func(t *testing.T) {
		currentStatus := []supply_model.StockStatus{
			supply_model.StockStatusOnDisposalProcess,
			supply_model.StockStatusDefect,
			supply_model.StockStatusExpired,
			supply_model.StockStatusNotPairedYet,
		}

		newStatus := functions.RemoveStockStatus(
			currentStatus,
			supply_model.StockStatusOnDisposalProcess,
			supply_model.StockStatusDefect,
			supply_model.StockStatusExpired,
			supply_model.StockStatusNotPairedYet,
		)

		assert.Equal(t, []supply_model.StockStatus{supply_model.StockStatusAvailable}, newStatus)
	})
}

func TestSetStockStatus(t *testing.T) {
	t.Run("successfully set status", func(t *testing.T) {
		testableStatus := []supply_model.StockStatus{
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
			supply_model.StockStatusReturned,
		}

		for _, status := range testableStatus {
			currentStatus := []supply_model.StockStatus{supply_model.StockStatusAvailable}

			newStatus, err := functions.SetStockStatus(currentStatus, status)

			assert.NoError(t, err)
			assert.Equal(t, []supply_model.StockStatus{status}, newStatus)
		}
	})

	t.Run("can't set status more than 1", func(t *testing.T) {
		currentStatus := []supply_model.StockStatus{
			supply_model.StockStatusInStockOpname,
			supply_model.StockStatusUnplaced,
		}

		newStatus, err := functions.SetStockStatus(currentStatus, supply_model.StockStatusAvailable)

		assert.Error(t, err)
		assert.Equal(t, currentStatus, newStatus)
	})

	t.Run("can't set status directly", func(t *testing.T) {
		currentStatus := []supply_model.StockStatus{
			supply_model.StockStatusInStockOpname,
			supply_model.StockStatusUnplaced,
		}

		newStatus, err := functions.SetStockStatus(currentStatus, supply_model.StockStatusWaitingForApproval)

		assert.Error(t, err)
		assert.Equal(t, currentStatus, newStatus)
	})
}

func TestExampleSetAndAddOrRemoveStatus(t *testing.T) {
	t.Run("set status error, manually add status", func(t *testing.T) {
		currentStatus := []supply_model.StockStatus{
			supply_model.StockStatusInStockOpname,
			supply_model.StockStatusUnplaced,
		}

		newStatus, err := functions.SetStockStatus(currentStatus, supply_model.StockStatusExpired)
		if err != nil {
			newStatus = functions.AddStockStatus(currentStatus, supply_model.StockStatusExpired)
		}

		assert.Error(t, err)
		assert.Equal(t, []supply_model.StockStatus{
			supply_model.StockStatusInStockOpname,
			supply_model.StockStatusUnplaced,
			supply_model.StockStatusExpired,
		}, newStatus)
	})

	t.Run("set status error, manually remove status", func(t *testing.T) {
		currentStatus := []supply_model.StockStatus{
			supply_model.StockStatusOnTransfer,
			supply_model.StockStatusExpired,
		}

		newStatus, err := functions.SetStockStatus(currentStatus, supply_model.StockStatusAvailable)
		if err != nil {
			newStatus = functions.RemoveStockStatus(currentStatus, supply_model.StockStatusOnTransfer)
		}

		assert.Error(t, err)
		assert.Equal(t, []supply_model.StockStatus{
			supply_model.StockStatusExpired,
		}, newStatus)
	})
}
