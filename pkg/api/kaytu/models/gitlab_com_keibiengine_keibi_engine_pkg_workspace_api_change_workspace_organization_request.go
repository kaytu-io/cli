// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIChangeWorkspaceOrganizationRequest gitlab com keibiengine keibi engine pkg workspace api change workspace organization request
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_workspace_api.ChangeWorkspaceOrganizationRequest
type GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIChangeWorkspaceOrganizationRequest struct {

	// new org ID
	NewOrgID int64 `json:"newOrgID,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg workspace api change workspace organization request
func (m *GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIChangeWorkspaceOrganizationRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this gitlab com keibiengine keibi engine pkg workspace api change workspace organization request based on context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIChangeWorkspaceOrganizationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIChangeWorkspaceOrganizationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIChangeWorkspaceOrganizationRequest) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIChangeWorkspaceOrganizationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}