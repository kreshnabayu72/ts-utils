package new_model

import (
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/common/model"
)

type SampHistory struct {
	ID int `json:"id" column:"samp_histories.id"`

	SampID                           string                 `json:"sampId" column:"samp_histories.samp_id"`
	Groups                           model.JSONBArrayString `json:"groups" column:"samp_histories.groups"`
	Name                             string                 `json:"name" column:"samp_histories.name"`
	Version                          float64                `json:"version" column:"samp_histories.version"`
	Status                           string                 `json:"status" column:"samp_histories.status"`
	StrategicDirection               string                 `json:"strategicDirection" column:"samp_histories.strategic_direction"`
	OrganizationalGoals              string                 `json:"organizationalGoals" column:"samp_histories.organizational_goals"`
	TopManagementCommitmentStatement string                 `json:"topManagementCommitmentStatement" column:"samp_histories.top_management_commitment_statement"`
	AMSystemPolicyReferenceTitle     string                 `json:"amSystemPolicyReferenceTitle" column:"samp_histories.am_system_policy_reference_title"`
	AMSystemPolicyReferenceFilePath  string                 `json:"amSystemPolicyReferenceFilePath" column:"samp_histories.am_system_policy_reference_file_path"`
	AMSystemPolicyReferenceFileName  string                 `json:"amSystemPolicyReferenceFileName" column:"samp_histories.am_system_policy_reference_file_name"`
	AMSystemStructureDescription     string                 `json:"amSystemStructureDescription" column:"samp_histories.am_system_structure_description"`
	AMSystemStructureFilePath        string                 `json:"amSystemStructureFilePath" column:"samp_histories.am_system_structure_file_path"`
	AMSystemStructureFileName        string                 `json:"amSystemStructureFileName" column:"samp_histories.am_system_structure_file_name"`
	ScopeOfAMSystem                  string                 `json:"scopeOfAMSystem" column:"samp_histories.scope_of_am_system"`
	LegalAndRegulatoryRequirements   string                 `json:"legalAndRegulatoryRequirements" column:"samp_histories.legal_and_regulatory_requirements"`

	ModifiedByID int `json:"modifiedById" column:"samp_histories.modified_by_id"`

	CreatedAt time.Time `json:"createdAt" column:"samp_histories.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"samp_histories.updated_at"`
}
