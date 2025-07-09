package new_model

import "time"

type Client struct {
	ID                        int       `json:"_id"                         column:"id"`
	Name                      string    `json:"name"                        column:"name"`
	Status                    string    `json:"status"                      column:"status"` // TODO(refan): change this to an enum.
	ClientId                  string    `json:"clientId"                    column:"client_id"`
	Email                     *string   `json:"email"                       column:"email"`
	IsEmailConfirmed          bool      `json:"isEmailConfirmed"            column:"is_email_confirmed"`
	Country                   string    `json:"country"                     column:"country"`
	State                     string    `json:"state"                       column:"state"`
	City                      string    `json:"city"                        column:"city"`
	Zip                       string    `json:"zip"                         column:"zip"`
	PhoneNumber               string    `json:"phoneNumber"                 column:"phone_number"`
	Fax                       *string   `json:"fax"                         column:"fax"`
	Address                   string    `json:"address"                     column:"address"`
	ContactPersonName         string    `json:"contactPersonName"           column:"contact_person_name"`
	ContactPersonEmail        string    `json:"contactPersonEmail"          column:"contact_person_email"`
	PendingEmailChange        *string   `json:"pendingEmailChange"          column:"pending_email_change"`
	Image                     *string   `json:"image"                       column:"image"`
	LoginSessionMinutes       int       `json:"loginSessionMinutes"         column:"login_session_minutes"`
	PlanFixedAsset            string    `json:"planFixedAsset"              column:"plan_fixed_asset"`            // TODO(refan): change this to an enum.
	PlanSupplyAsset           string    `json:"planSupplyAsset"             column:"plan_supply_asset"`           // TODO(refan): change this to an enum.
	PlanInventoryAsset        string    `json:"planInventoryAsset"          column:"plan_inventory_asset"`        // TODO(refan): change this to an enum.
	PlanProcurementAsset      string    `json:"planProcurementAsset"        column:"plan_procurement_asset"`      // TODO(refan): change this to an enum.
	ModifiedById              int       `json:"modifiedById"                column:"modified_by_id"`              // REFERENCES user.
	DepreciationMethod        string    `json:"depreciationMethod"          column:"depreciation_method"`         // TODO(refan): change this to an enum.
	PendingDepreciationMethod *string   `json:"pendingDepreciationMethod"   column:"pending_depreciation_method"` // TODO(refan): change this to an enum.
	AddonRepairTicketing      bool      `json:"addonRepairTicketing"        column:"addon_repair_ticketing"`
	AddonMaintenance          bool      `json:"addonMaintenance"            column:"addon_maintenance"`
	AddonTracking             bool      `json:"addonTracking"               column:"addon_tracking"`
	AddonAssetControl         bool      `json:"addonAssetControl"           column:"addon_asset_control"`
	AddonAudit                bool      `json:"addonAudit"                  column:"addon_audit"`
	DateFormatLocale          string    `json:"dateFormatLocale"            column:"date_format_locale"`
	DateFormatYear            string    `json:"dateFormatYear"              column:"date_format_year"`
	DateFormatMonth           string    `json:"dateFormatMonth"             column:"date_format_month"`
	DateFormatDay             string    `json:"dateFormatDay"               column:"date_format_day"`
	DateFormatWeekday         string    `json:"dateFormatWeekday"           column:"date_format_weekday"`
	TimeFormat                bool      `json:"timeFormat"                  column:"time_format"`
	Geolocation               bool      `json:"geolocation"                 column:"geolocation"`
	TimeZone                  string    `json:"timeZone"                    column:"time_zone"`
	Pin                       string    `json:"pin"                         column:"pin"`
	CurrencySymbol            string    `json:"currencySymbol"              column:"currency_symbol"`
	CurrencyCurrency          string    `json:"currencyCurrency"            column:"currency_currency"`
	CurrencyLocale            string    `json:"currencyLocale"              column:"currency_locale"`
	CurrencyLabel             string    `json:"currencyLabel"               column:"currency_label"`
	ExpiredAt                 time.Time `json:"expiredAt"                   column:"expired_at"`
	CreatedAt                 time.Time `json:"createdAt"                   column:"created_at"`
	UpdatedAt                 time.Time `json:"updatedAt"                   column:"updated_at"`
}
