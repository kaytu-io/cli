// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIWorkspaceLimitsUsage gitlab com keibiengine keibi engine pkg workspace api workspace limits usage
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_workspace_api.WorkspaceLimitsUsage
type GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIWorkspaceLimitsUsage struct {

	// current connections
	CurrentConnections int64 `json:"currentConnections,omitempty"`

	// current resources
	CurrentResources int64 `json:"currentResources,omitempty"`

	// current users
	CurrentUsers int64 `json:"currentUsers,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// max connections
	MaxConnections int64 `json:"maxConnections,omitempty"`

	// max resources
	MaxResources int64 `json:"maxResources,omitempty"`

	// max users
	MaxUsers int64 `json:"maxUsers,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg workspace api workspace limits usage
func (m *GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIWorkspaceLimitsUsage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this gitlab com keibiengine keibi engine pkg workspace api workspace limits usage based on context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIWorkspaceLimitsUsage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIWorkspaceLimitsUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIWorkspaceLimitsUsage) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIWorkspaceLimitsUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
