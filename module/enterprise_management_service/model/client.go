// Package model defines the enterprise-management-service database row types.
package model

import (
	"time"
)

// Client represents the 'clients' table in the enterprise management database.
type Client struct {
	ID                              int        `column:"id"                                  json:"id"`
	Status                          string     `column:"status"                              json:"status"`
	Prefix                          string     `column:"prefix"                              json:"prefix"`
	Suffix                          string     `column:"suffix"                              json:"suffix"`
	Name                            string     `column:"name"                                json:"name"`
	ClientID                        string     `column:"client_id"                           json:"clientId"`
	PartnerID                       *int64     `column:"partner_id"                          json:"partnerId"`
	Email                           *string    `column:"email"                               json:"email"`
	Country                         string     `column:"country"                             json:"country"`
	State                           string     `column:"state"                               json:"state"`
	City                            string     `column:"city"                                json:"city"`
	Zip                             string     `column:"zip"                                 json:"zip"`
	PhoneNumber                     string     `column:"phone_number"                        json:"phoneNumber"`
	Fax                             *string    `column:"fax"                                 json:"fax"`
	Address                         string     `column:"address"                             json:"address"`
	GroupID                         int64      `column:"group_id"                            json:"groupId"`
	Image                           *string    `column:"image"                               json:"image"`
	ContactPersonName               string     `column:"contact_person_name"                 json:"contactPersonName"`
	ContactPersonEmail              string     `column:"contact_person_email"                json:"contactPersonEmail"`
	ContactPersonPendingEmailChange *string    `column:"contact_person_pending_email_change" json:"contactPersonPendingEmailChange"`
	LoginSessionMinutes             int        `column:"login_session_minutes"               json:"loginSessionMinutes"`
	PinNumber                       string     `column:"pin_number"                          json:"pinNumber"`
	IsDeleted                       bool       `column:"is_deleted"                          json:"isDeleted"`
	DeletedAt                       *time.Time `column:"deleted_at"                          json:"deletedAt"`
	CreatedAt                       time.Time  `column:"created_at"                          json:"createdAt"`
	UpdatedAt                       time.Time  `column:"updated_at"                          json:"updatedAt"`
	IsEmailConfirmed                bool       `column:"is_email_confirmed"                  json:"isEmailConfirmed"`
	DateFormat                      *string    `column:"date_format"                         json:"dateFormat"`
	TimeFormat                      *string    `column:"time_format"                         json:"timeFormat"`
	TimeZone                        *string    `column:"time_zone"                           json:"timeZone"`
	Geolocation                     *bool      `column:"geolocation"                         json:"geolocation"`
	PlanFixedAsset                  *string    `column:"plan_fixed_asset"                    json:"planFixedAsset"`
	PlanSupplyAsset                 *string    `column:"plan_supply_asset"                   json:"planSupplyAsset"`
	PlanInventoryAsset              *string    `column:"plan_inventory_asset"                json:"planInventoryAsset"`
	PlanProcurementAsset            *string    `column:"plan_procurement_asset"              json:"planProcurementAsset"`
	CurrencyCurrency                *string    `column:"currency_currency"                   json:"currencyCurrency"`
	CurrencySymbol                  *string    `column:"currency_symbol"                     json:"currencySymbol"`
	CurrencyLocale                  *string    `column:"currency_locale"                     json:"currencyLocale"`
	CurrencyLabel                   *string    `column:"currency_label"                      json:"currencyLabel"`
	FileStorageTotalKb              *float64   `column:"file_storage_total_kb"               json:"fileStorageTotalKb"`
	FileStorageAvailableKb          *float64   `column:"file_storage_available_kb"           json:"fileStorageAvailableKb"`
	FileStorageUsedKb               *float64   `column:"file_storage_used_kb"                json:"fileStorageUsedKb"`
	FileStorageAttachmentKb         *float64   `column:"file_storage_attachment_kb"          json:"fileStorageAttachmentKb"`
	FileStorageDamageKb             *float64   `column:"file_storage_damage_kb"              json:"fileStorageDamageKb"`
	FileStorageDisposalKb           *float64   `column:"file_storage_disposal_kb"            json:"fileStorageDisposalKb"`
	FileStorageRepairTicketingKb    *float64   `column:"file_storage_repair_ticketing_kb"    json:"fileStorageRepairTicketingKb"`
	FileStorageRoutineKb            *float64   `column:"file_storage_routine_kb"             json:"fileStorageRoutineKb"`
	AddonRepairTicketing            *bool      `column:"addon_repair_ticketing"              json:"addonRepairTicketing"`
	AddonMaintenance                *bool      `column:"addon_maintenance"                   json:"addonMaintenance"`
	AddonTracking                   *bool      `column:"addon_tracking"                      json:"addonTracking"`
	AddonAssetControl               *bool      `column:"addon_asset_control"                 json:"addonAssetControl"`
	AddonMap                        *bool      `column:"addon_map"                           json:"addonMap"`
	AddonAudit                      *bool      `column:"addon_audit"                         json:"addonAudit"`
	FileStorageDeleteCountdown      *int64     `column:"file_storage_delete_countdown"       json:"fileStorageDeleteCountdown"`
}

