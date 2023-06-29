// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GithubComKaytuIoKaytuEnginePkgAuthAPIMembership github com kaytu io kaytu engine pkg auth api membership
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_auth_api.Membership
type GithubComKaytuIoKaytuEnginePkgAuthAPIMembership struct {

	// Assignment timestamp in UTC
	// Example: 2023-03-31T09:36:09.855Z
	AssignedAt string `json:"assignedAt,omitempty"`

	// Last activity timestamp in UTC
	// Example: 2023-04-21T08:53:09.928Z
	LastActivity string `json:"lastActivity,omitempty"`

	// Name of the role
	// Example: admin
	// Enum: [admin editor viewer]
	RoleName GithubComKaytuIoKaytuEnginePkgAuthAPIRole `json:"roleName,omitempty"`

	// Unique identifier for the workspace
	// Example: ws123456789
	WorkspaceID string `json:"workspaceID,omitempty"`

	// Name of the Workspace
	// Example: demo
	WorkspaceName string `json:"workspaceName,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg auth api membership
func (m *GithubComKaytuIoKaytuEnginePkgAuthAPIMembership) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRoleName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var githubComKaytuIoKaytuEnginePkgAuthApiMembershipTypeRoleNamePropEnum []interface{}

func init() {
	var res []GithubComKaytuIoKaytuEnginePkgAuthAPIRole
	if err := json.Unmarshal([]byte(`["admin","editor","viewer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		githubComKaytuIoKaytuEnginePkgAuthApiMembershipTypeRoleNamePropEnum = append(githubComKaytuIoKaytuEnginePkgAuthApiMembershipTypeRoleNamePropEnum, v)
	}
}

// prop value enum
func (m *GithubComKaytuIoKaytuEnginePkgAuthAPIMembership) validateRoleNameEnum(path, location string, value *struct {
	GithubComKaytuIoKaytuEnginePkgAuthAPIRole
}) error {
	if err := validate.EnumCase(path, location, value, githubComKaytuIoKaytuEnginePkgAuthApiMembershipTypeRoleNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgAuthAPIMembership) validateRoleName(formats strfmt.Registry) error {
	if swag.IsZero(m.RoleName) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this github com kaytu io kaytu engine pkg auth api membership based on the context it is used
func (m *GithubComKaytuIoKaytuEnginePkgAuthAPIMembership) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRoleName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgAuthAPIMembership) contextValidateRoleName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgAuthAPIMembership) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgAuthAPIMembership) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgAuthAPIMembership
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
