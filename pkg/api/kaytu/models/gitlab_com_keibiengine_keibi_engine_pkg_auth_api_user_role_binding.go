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

// GitlabComKeibiengineKeibiEnginePkgAuthAPIUserRoleBinding gitlab com keibiengine keibi engine pkg auth api user role binding
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_auth_api.UserRoleBinding
type GitlabComKeibiengineKeibiEnginePkgAuthAPIUserRoleBinding struct {

	// Name of the binding Role
	// Example: admin
	// Enum: [admin editor viewer]
	RoleName struct {
		GitlabComKeibiengineKeibiEnginePkgAuthAPIRole
	} `json:"roleName,omitempty"`

	// Unique identifier for the Workspace
	// Example: sampleID
	WorkspaceID string `json:"workspaceID,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg auth api user role binding
func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIUserRoleBinding) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRoleName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var gitlabComKeibiengineKeibiEnginePkgAuthApiUserRoleBindingTypeRoleNamePropEnum []interface{}

func init() {
	var res []GitlabComKeibiengineKeibiEnginePkgAuthAPIRole
	if err := json.Unmarshal([]byte(`["admin","editor","viewer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gitlabComKeibiengineKeibiEnginePkgAuthApiUserRoleBindingTypeRoleNamePropEnum = append(gitlabComKeibiengineKeibiEnginePkgAuthApiUserRoleBindingTypeRoleNamePropEnum, v)
	}
}

// prop value enum
func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIUserRoleBinding) validateRoleNameEnum(path, location string, value *struct {
	GitlabComKeibiengineKeibiEnginePkgAuthAPIRole
}) error {
	if err := validate.EnumCase(path, location, value, gitlabComKeibiengineKeibiEnginePkgAuthApiUserRoleBindingTypeRoleNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIUserRoleBinding) validateRoleName(formats strfmt.Registry) error {
	if swag.IsZero(m.RoleName) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this gitlab com keibiengine keibi engine pkg auth api user role binding based on the context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIUserRoleBinding) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRoleName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIUserRoleBinding) contextValidateRoleName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIUserRoleBinding) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIUserRoleBinding) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgAuthAPIUserRoleBinding
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
