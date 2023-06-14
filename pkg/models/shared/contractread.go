// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"ContractifyProduction/pkg/types"
	"time"
)

type ContractRead struct {
	ContractTypes     []ContractTypeRead     `json:"contract_types,omitempty"`
	CreatedAt         *time.Time             `json:"created_at,omitempty"`
	CustomFieldValues []CustomFieldValueRead `json:"custom_field_values,omitempty"`
	Departments       []DepartmentRead       `json:"departments,omitempty"`
	Documents         []ContractDocumentRead `json:"documents,omitempty"`
	Dossier           *DossierRead           `json:"dossier,omitempty"`
	Duration          *string                `json:"duration,omitempty"`
	EndDate           *types.Date            `json:"end_date,omitempty"`
	ID                *int64                 `json:"id,omitempty"`
	IsOpenEnded       *bool                  `json:"is_open_ended,omitempty"`
	LegalEntities     []LegalEntityRead      `json:"legal_entities,omitempty"`
	Name              *string                `json:"name,omitempty"`
	Offices           []OfficeRead           `json:"offices,omitempty"`
	OwnerID           *int64                 `json:"owner_id,omitempty"`
	Permalink         *string                `json:"permalink,omitempty"`
	Phase             *ContractPhase         `json:"phase,omitempty"`
	Relations         []RelationRead         `json:"relations,omitempty"`
	Renewal           *ContractRenewal       `json:"renewal,omitempty"`
	StartDate         *types.Date            `json:"start_date,omitempty"`
	Termination       *ContractTermination   `json:"termination,omitempty"`
}
