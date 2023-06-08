// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIWorkspaceLimits gitlab com keibiengine keibi engine pkg workspace api workspace limits
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_workspace_api.WorkspaceLimits
type GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIWorkspaceLimits struct {

	// max connections
	MaxConnections int64 `json:"maxConnections,omitempty"`

	// max resources
	MaxResources int64 `json:"maxResources,omitempty"`

	// max users
	MaxUsers int64 `json:"maxUsers,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg workspace api workspace limits
func (m *GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIWorkspaceLimits) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this gitlab com keibiengine keibi engine pkg workspace api workspace limits based on context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIWorkspaceLimits) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIWorkspaceLimits) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIWorkspaceLimits) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIWorkspaceLimits
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
