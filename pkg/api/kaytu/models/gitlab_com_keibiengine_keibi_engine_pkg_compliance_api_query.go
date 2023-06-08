// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GitlabComKeibiengineKeibiEnginePkgComplianceAPIQuery gitlab com keibiengine keibi engine pkg compliance api query
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_compliance_api.Query
type GitlabComKeibiengineKeibiEnginePkgComplianceAPIQuery struct {

	// connector
	Connector string `json:"connector,omitempty"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// engine
	Engine string `json:"engine,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// list of tables
	ListOfTables []string `json:"listOfTables"`

	// query to execute
	QueryToExecute string `json:"queryToExecute,omitempty"`

	// updated at
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg compliance api query
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIQuery) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this gitlab com keibiengine keibi engine pkg compliance api query based on context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIQuery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIQuery) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgComplianceAPIQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}