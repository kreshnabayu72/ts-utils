package model

type DefaultCurrency struct {
	Currency string `json:"currency" bson:"currency"`
	Symbol   string `json:"symbol"   bson:"symbol"`
	Locale   string `json:"locale"   bson:"locale"`
	Label    string `json:"label"    bson:"label"`
}

type ReceptionAttrs struct {
	TransactionTimeout int `json:"transactionTimeout" bson:"transactionTimeout"`
}

type CheckoutAttrs struct {
	ExpiryReminder           int `json:"expiryReminder"           bson:"expiryReminder"`
	ConfirmationEmailTimeout int `json:"confirmationEmailTimeout" bson:"confirmationEmailTimeout"`
	TransactionTimeout       int `json:"transactionTimeout"       bson:"transactionTimeout"`
}

type TransferAttrs struct {
	TransactionTimeout int `json:"transactionTimeout" bson:"transactionTimeout"`
}

type DisposalAttrs struct {
	TransactionTimeout int `json:"transactionTimeout" bson:"transactionTimeout"`
}

type StockOpnameAttrs struct {
	MaxGracePeriod int `json:"maxGracePeriod" bson:"maxGracePeriod"`
}

type ConsoleSettings struct {
	// General Settings
	DefaultCurrency DefaultCurrency `json:"defaultCurrency" bson:"defaultCurrency"`
	Timezone        string          `json:"timezone"        bson:"timezone"`
	DateFormat      string          `json:"dateFormat"      bson:"dateFormat"`
	TimeFormat      string          `json:"timeFormat"      bson:"timeFormat"`
	Geolocation     bool            `json:"geolocation"     bson:"geolocation"`

	// Transaction Settings
	ReceptionAttrs   ReceptionAttrs   `json:"receptionAttrs"   bson:"receptionAttrs"`
	CheckoutAttrs    CheckoutAttrs    `json:"checkoutAttrs"    bson:"checkoutAttrs"`
	TransferAttrs    TransferAttrs    `json:"transferAttrs"    bson:"transferAttrs"`
	DisposalAttrs    DisposalAttrs    `json:"disposalAttrs"    bson:"disposalAttrs"`
	StockOpnameAttrs StockOpnameAttrs `json:"stockOpnameAttrs" bson:"stockOpnameAttrs"`
}
