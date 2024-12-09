package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Struct for the auditAttrs
type TransactionSettingAuditAttrs struct {
	MaxGracePeriod int `json:"maxGracePeriod" bson:"maxGracePeriod"`
}

// Struct for the borrowingAttrs
type TransactionSettingBorrowingAttrs struct {
	MaxBorrowingTime     int `json:"maxBorrowingTime" bson:"maxBorrowingTime"`
	TransactionTimeout   int `json:"transactionTimeout" bson:"transactionTimeout"`
	DefaultBorrowingTime int `json:"defaultBorrowingTime" bson:"defaultBorrowingTime"`
	ConfirmationTimeout  int `json:"confirmationTimeout" bson:"confirmationTimeout"`
}

// Struct for the transferAttrs
type TransactionSettingTransferAttrs struct {
	TransactionTimeout int `json:"transactionTimeout" bson:"transactionTimeout"`
}

// Struct for the disposalAttrs
type TransactionSettingDisposalAttrs struct {
	TransactionTimeout int `json:"transactionTimeout" bson:"transactionTimeout"`
}

type TransactionSetting struct {
	ID             primitive.ObjectID               `json:"_id" bson:"_id"`
	AuditAttrs     TransactionSettingAuditAttrs     `json:"auditAttrs" bson:"auditAttrs"`
	BorrowingAttrs TransactionSettingBorrowingAttrs `json:"borrowingAttrs" bson:"borrowingAttrs"`
	TransferAttrs  TransactionSettingTransferAttrs  `json:"transferAttrs" bson:"transferAttrs"`
	DisposalAttrs  TransactionSettingDisposalAttrs  `json:"disposalAttrs" bson:"disposalAttrs"`
}
