package new_model

import (
	"time"
)

type ServiceCenter struct {
	ID             int        `json:"_id"            column:"service_centers.id"`
	IsActive       bool       `json:"isActive"       column:"service_centers.is_active"`
	IsDefault      bool       `json:"isDefault"      column:"service_centers.is_default"`
	Name           string     `json:"name"           column:"service_centers.name"`
	City           string     `json:"city"           column:"service_centers.city"`
	Country        string     `json:"country"        column:"service_centers.country"`
	CountryIso     string     `json:"countryIso"        column:"service_centers.country_iso"`
	State          string     `json:"state"          column:"service_centers.state"`
	StateIso       string     `json:"stateIso"        column:"service_centers.state_iso"`
	ZipCode        string     `json:"zipCode"        column:"service_centers.zip_code"`
	Address        string     `json:"address"        column:"service_centers.address"`
	PhoneNumber    string     `json:"phoneNumber"    column:"service_centers.phone_number"`
	PicPhoneNumber string     `json:"picPhoneNumber" column:"service_centers.pic_phone_number"`
	PicName        string     `json:"picName"        column:"service_centers.pic_name"`
	ModifiedById   *int       `json:"modifiedById"   column:"service_centers.modified_by_id"`
	IsDeleted      bool       `json:"isDeleted"      column:"service_centers.is_deleted"`
	DeletedAt      *time.Time `json:"deletedAt"      column:"service_centers.deleted_at"`
	CreatedAt      time.Time  `json:"createdAt"      column:"service_centers.created_at"`
	UpdatedAt      time.Time  `json:"updatedAt"      column:"service_centers.updated_at"`
}
