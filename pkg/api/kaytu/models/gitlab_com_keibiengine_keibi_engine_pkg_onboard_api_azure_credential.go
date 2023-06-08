// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GitlabComKeibiengineKeibiEnginePkgOnboardAPIAzureCredential gitlab com keibiengine keibi engine pkg onboard api azure credential
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_onboard_api.AzureCredential
type GitlabComKeibiengineKeibiEnginePkgOnboardAPIAzureCredential struct {

	// client ID
	ClientID string `json:"clientID,omitempty"`

	// client secret
	ClientSecret string `json:"clientSecret,omitempty"`

	// tenant ID
	TenantID string `json:"tenantID,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg onboard api azure credential
func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPIAzureCredential) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this gitlab com keibiengine keibi engine pkg onboard api azure credential based on context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPIAzureCredential) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPIAzureCredential) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPIAzureCredential) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgOnboardAPIAzureCredential
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}