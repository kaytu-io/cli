// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GitlabComKeibiengineKeibiEnginePkgAuthAPIGetRoleBindingsResponse gitlab com keibiengine keibi engine pkg auth api get role bindings response
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_auth_api.GetRoleBindingsResponse
type GitlabComKeibiengineKeibiEnginePkgAuthAPIGetRoleBindingsResponse struct {

	// Global Access
	// Example: admin
	// Enum: [admin editor viewer]
	GlobalRoles GitlabComKeibiengineKeibiEnginePkgAuthAPIRole `json:"globalRoles,omitempty"`

	// List of user roles in each workspace
	RoleBindings []*GitlabComKeibiengineKeibiEnginePkgAuthAPIUserRoleBinding `json:"roleBindings"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg auth api get role bindings response
func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIGetRoleBindingsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGlobalRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleBindings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var gitlabComKeibiengineKeibiEnginePkgAuthApiGetRoleBindingsResponseTypeGlobalRolesPropEnum []interface{}

func init() {
	var res []GitlabComKeibiengineKeibiEnginePkgAuthAPIRole
	if err := json.Unmarshal([]byte(`["admin","editor","viewer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gitlabComKeibiengineKeibiEnginePkgAuthApiGetRoleBindingsResponseTypeGlobalRolesPropEnum = append(gitlabComKeibiengineKeibiEnginePkgAuthApiGetRoleBindingsResponseTypeGlobalRolesPropEnum, v)
	}
}

// prop value enum
func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIGetRoleBindingsResponse) validateGlobalRolesEnum(path, location string, value *struct {
	GitlabComKeibiengineKeibiEnginePkgAuthAPIRole
}) error {
	if err := validate.EnumCase(path, location, value, gitlabComKeibiengineKeibiEnginePkgAuthApiGetRoleBindingsResponseTypeGlobalRolesPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIGetRoleBindingsResponse) validateGlobalRoles(formats strfmt.Registry) error {
	if swag.IsZero(m.GlobalRoles) { // not required
		return nil
	}

	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIGetRoleBindingsResponse) validateRoleBindings(formats strfmt.Registry) error {
	if swag.IsZero(m.RoleBindings) { // not required
		return nil
	}

	for i := 0; i < len(m.RoleBindings); i++ {
		if swag.IsZero(m.RoleBindings[i]) { // not required
			continue
		}

		if m.RoleBindings[i] != nil {
			if err := m.RoleBindings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roleBindings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("roleBindings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this gitlab com keibiengine keibi engine pkg auth api get role bindings response based on the context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIGetRoleBindingsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGlobalRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoleBindings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIGetRoleBindingsResponse) contextValidateGlobalRoles(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIGetRoleBindingsResponse) contextValidateRoleBindings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RoleBindings); i++ {

		if m.RoleBindings[i] != nil {

			if swag.IsZero(m.RoleBindings[i]) { // not required
				return nil
			}

			if err := m.RoleBindings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roleBindings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("roleBindings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIGetRoleBindingsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIGetRoleBindingsResponse) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgAuthAPIGetRoleBindingsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
