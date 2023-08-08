// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GithubComKaytuIoKaytuEnginePkgWorkspaceAPIChangeWorkspaceOrganizationRequest github com kaytu io kaytu engine pkg workspace api change workspace organization request
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_workspace_api.ChangeWorkspaceOrganizationRequest
type GithubComKaytuIoKaytuEnginePkgWorkspaceAPIChangeWorkspaceOrganizationRequest struct {

	// new org ID
	NewOrgID int64 `json:"newOrgID,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg workspace api change workspace organization request
func (m *GithubComKaytuIoKaytuEnginePkgWorkspaceAPIChangeWorkspaceOrganizationRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this github com kaytu io kaytu engine pkg workspace api change workspace organization request based on context it is used
func (m *GithubComKaytuIoKaytuEnginePkgWorkspaceAPIChangeWorkspaceOrganizationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgWorkspaceAPIChangeWorkspaceOrganizationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgWorkspaceAPIChangeWorkspaceOrganizationRequest) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgWorkspaceAPIChangeWorkspaceOrganizationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
