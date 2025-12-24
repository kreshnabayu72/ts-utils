package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IOTReaderUsageLog struct {
	ID               primitive.ObjectID `json:"_id"              bson:"_id"`
	Port             string             `json:"port"             bson:"port"`
	AntennaGroup     string             `json:"antennaGroup"     bson:"antennaGroup"`
	ReaderID         string             `json:"readerId"         bson:"readerId"`
	TransactionID    string             `json:"transactionId"    bson:"transactionId"`
	Name             string             `json:"name"             bson:"name"`
	NameWithSequence string             `json:"nameWithSequence" bson:"nameWithSequence"`
	AuditType        string             `json:"auditType"        bson:"auditType"`
	Manager          string             `json:"manager"          bson:"manager"`
	Activity         string             `json:"activity"         bson:"activity"`
	PortGroup        string             `json:"portGroup"        bson:"portGroup"`
	Stock            string             `json:"stock"            bson:"stock"`
	SKU              string             `json:"sku"              bson:"sku"`
	PIC              string             `json:"pic"              bson:"pic"`
	CreatedAt        time.Time          `json:"createdAt"        bson:"createdAt"`
	UpdatedAt        time.Time          `json:"updatedAt"        bson:"updatedAt"`
}
