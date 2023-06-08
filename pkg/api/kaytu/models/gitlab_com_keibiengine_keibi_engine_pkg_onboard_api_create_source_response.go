// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GitlabComKeibiengineKeibiEnginePkgOnboardAPICreateSourceResponse gitlab com keibiengine keibi engine pkg onboard api create source response
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_onboard_api.CreateSourceResponse
type GitlabComKeibiengineKeibiEnginePkgOnboardAPICreateSourceResponse struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg onboard api create source response
func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPICreateSourceResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this gitlab com keibiengine keibi engine pkg onboard api create source response based on context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPICreateSourceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPICreateSourceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPICreateSourceResponse) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgOnboardAPICreateSourceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
