package new_model

import (
	"time"
)

type SampCurrentVersion struct {
	ID int `json:"id" column:"samp_current_versions.id"`

	IsSubmitted bool `json:"isSubmitted" column:"samp_current_versions.is_submitted"`

	// Need pointers, after init (before submit), the values are null
	SampID                           *string `json:"sampId" column:"samp_current_versions.samp_id"`
	Name                             *string `json:"name" column:"samp_current_versions.name"`
	Version                          *string `json:"version" column:"samp_current_versions.version"`
	Status                           *string `json:"status" column:"samp_current_versions.status"`
	StrategicDirection               *string `json:"strategicDirection" column:"samp_current_versions.strategic_direction"`
	OrganizationalGoals              *string `json:"organizationalGoals" column:"samp_current_versions.organizational_goals"`
	TopManagementCommitmentStatement *string `json:"topManagementCommitmentStatement" column:"samp_current_versions.top_management_commitment_statement"`
	AMSystemPolicyReferenceTitle     *string `json:"amSystemPolicyReferenceTitle" column:"samp_current_versions.am_system_policy_reference_title"`
	AMSystemPolicyReferenceFilePath  *string `json:"amSystemPolicyReferenceFilePath" column:"samp_current_versions.am_system_policy_reference_file_path"`
	AMSystemPolicyReferenceFileName  *string `json:"amSystemPolicyReferenceFileName" column:"samp_current_versions.am_system_policy_reference_file_name"`
	AMSystemStructureDescription     *string `json:"amSystemStructureDescription" column:"samp_current_versions.am_system_structure_description"`
	AMSystemStructureFilePath        *string `json:"amSystemStructureFilePath" column:"samp_current_versions.am_system_structure_file_path"`
	AMSystemStructureFileName        *string `json:"amSystemStructureFileName" column:"samp_current_versions.am_system_structure_file_name"`
	ScopeOfAMSystem                  *string `json:"scopeOfAMSystem" column:"samp_current_versions.scope_of_am_system"`
	LegalAndRegulatoryRequirements   *string `json:"legalAndRegulatoryRequirements" column:"samp_current_versions.legal_and_regulatory_requirements"`

	ModifiedByID int `json:"modifiedById" column:"samp_current_versions.modified_by_id"`

	CreatedAt time.Time `json:"createdAt" column:"samp_current_versions.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"samp_current_versions.updated_at"`
}
