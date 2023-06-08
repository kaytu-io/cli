// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GitlabComKeibiengineKeibiEnginePkgComplianceAPIDatapoint gitlab com keibiengine keibi engine pkg compliance api datapoint
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_compliance_api.Datapoint
type GitlabComKeibiengineKeibiEnginePkgComplianceAPIDatapoint struct {

	// time
	Time int64 `json:"time,omitempty"`

	// value
	Value int64 `json:"value,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg compliance api datapoint
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIDatapoint) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this gitlab com keibiengine keibi engine pkg compliance api datapoint based on context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIDatapoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIDatapoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIDatapoint) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgComplianceAPIDatapoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
