// Package dto defines the enterprise-management-service activity parameters for
// the admin-console Hardware Stock module.
package dto

import hardwareTaggingDto "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/hardware_tagging_service/dto"

// HardwareStockReceiptWorkflowParam is the HardwareStockReceiptWorkflow input.
type HardwareStockReceiptWorkflowParam struct {
	ClientID      string
	EntityType    string
	EntityID      *int64
	EntityName    string
	UserID        int64
	UserFirstName string
	UserLastName  string
	UserFullName  string

	SkuID            int64
	SkuCode          string
	DeviceName       string
	HardwareType     string
	RangeType        string
	CategoryID       string
	CategorySnapshot string

	InputSource    string
	TransactionID  string
	IsTotalControl bool

	SerialNumbers []string
	Units         []hardwareTaggingDto.ReceiptUnit
}

// HardwareStockApprovalDecisionWorkflowParam is the
// HardwareStockApprovalDecisionWorkflow input.
type HardwareStockApprovalDecisionWorkflowParam struct {
	ClientID      string
	ApprovalID    int64
	UserID        int64
	UserFirstName string
	UserLastName  string
	UserFullName  string

	IsApproved bool
	Note       string
}
