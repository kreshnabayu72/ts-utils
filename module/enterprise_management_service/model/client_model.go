package model

import "time"

type Client struct {
	ID                              int        `json:"_id"                             column:"clients.id"`
	Name                            string     `json:"name"                            column:"clients.name"`
	ClientID                        string     `json:"clientId"                        column:"clients.client_id"`
	PartnerID                       *int       `json:"partnerId"                       column:"clients.partner_id"`
	Email                           *string    `json:"email"                           column:"clients.email"`
	Country                         string     `json:"country"                         column:"clients.country"`
	State                           string     `json:"state"                           column:"clients.state"`
	City                            string     `json:"city"                            column:"clients.city"`
	Zip                             string     `json:"zip"                             column:"clients.zip"`
	PhoneNumber                     string     `json:"phoneNumber"                     column:"clients.phone_number"`
	Fax                             *string    `json:"fax"                             column:"clients.fax"`
	Address                         string     `json:"address"                         column:"clients.address"`
	GroupID                         int        `json:"groupId"                         column:"clients.group_id"`
	Image                           *string    `json:"image"                           column:"clients.image"`
	ContactPersonName               string     `json:"contactPersonName"               column:"clients.contact_person_name"`
	ContactPersonEmail              string     `json:"contactPersonEmail"              column:"clients.contact_person_email"`
	ContactPersonPendingEmailChange *string    `json:"contactPersonPendingEmailChange" column:"clients.contact_person_pending_email_change"`
	LoginSessionMinutes             int        `json:"loginSessionMinutes"             column:"clients.login_session_minutes"`
	PinNumber                       string     `json:"pinNumber"                       column:"clients.pin_number"`
	IsDeleted                       bool       `json:"isDeleted"                       column:"clients.is_deleted"`
	DeletedAt                       *time.Time `json:"deletedAt"                       column:"clients.deleted_at"`
	CreatedAt                       time.Time  `json:"createdAt"                       column:"clients.created_at"`
	UpdatedAt                       time.Time  `json:"updatedAt"                       column:"clients.updated_at"`
	IsEmailConfirmed                bool       `json:"isEmailConfirmed"                column:"clients.is_email_confirmed"`
	DateFormat                      *string    `json:"dateFormat"                      column:"clients.date_format"`
	TimeFormat                      *string    `json:"timeFormat"                      column:"clients.time_format"`
	TimeZone                        *string    `json:"timeZone"                        column:"clients.time_zone"`
	Geolocation                     *bool      `json:"geolocation"                     column:"clients.geolocation"`
	PlanFixedAsset                  *string    `json:"planFixedAsset"                  column:"clients.plan_fixed_asset"`
	PlanSupplyAsset                 *string    `json:"planSupplyAsset"                 column:"clients.plan_supply_asset"`
	PlanInventoryAsset              *string    `json:"planInventoryAsset"              column:"clients.plan_inventory_asset"`
	PlanProcurementAsset            *string    `json:"planProcurementAsset"            column:"clients.plan_procurement_asset"`
	CurrencyCurrency                *string    `json:"currencyCurrency"                column:"clients.currency_currency"`
	CurrencySymbol                  *string    `json:"currencySymbol"                  column:"clients.currency_symbol"`
	CurrencyLocale                  *string    `json:"currencyLocale"                  column:"clients.currency_locale"`
	CurrencyLabel                   *string    `json:"currencyLabel"                   column:"clients.currency_label"`
	FileStorageTotalKB              *float64   `json:"fileStorageTotalKB"              column:"clients.file_storage_total_kb"`
	FileStorageAvailableKB          *float64   `json:"fileStorageAvailableKB"          column:"clients.file_storage_available_kb"`
	FileStorageUsedKB               *float64   `json:"fileStorageUsedKB"               column:"clients.file_storage_used_kb"`
	FileStorageAttachmentKB         *float64   `json:"fileStorageAttachmentKB"         column:"clients.file_storage_attachment_kb"`
	FileStorageDamagedKB            *float64   `json:"fileStorageDamagedKB"            column:"clients.file_storage_damage_kb"`
	FileStorageDisposalKB           *float64   `json:"fileStorageDisposalKB"           column:"clients.file_storage_disposal_kb"`
	FileStorageRepairTicketingKB    *float64   `json:"fileStorageRepairTicketingKB"    column:"clients.file_storage_repair_ticketing_kb"`
	FileStorageRoutineKB            *float64   `json:"fileStorageRoutineKB"            column:"clients.file_storage_routine_kb"`
}
