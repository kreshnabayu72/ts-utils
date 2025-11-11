package new_model

import (
	"time"
)

type AmpHistory struct {
	ID int `json:"id" column:"amp_histories.id"`

	AmpID     string `json:"ampId" column:"amp_histories.amp_id"`
	SampID    string `json:"sampId" column:"amp_histories.samp_id"`
	GroupName string `json:"groupName" column:"amp_histories.group_name"`
	Name      string `json:"name" column:"amp_histories.name"`
	Version   string `json:"version" column:"amp_histories.version"`

	CreatedAt time.Time `json:"createdAt" column:"amp_histories.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"amp_histories.updated_at"`
}
