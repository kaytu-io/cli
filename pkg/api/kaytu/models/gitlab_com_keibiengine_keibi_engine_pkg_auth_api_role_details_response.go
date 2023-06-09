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

// GitlabComKeibiengineKeibiEnginePkgAuthAPIRoleDetailsResponse gitlab com keibiengine keibi engine pkg auth api role details response
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_auth_api.RoleDetailsResponse
type GitlabComKeibiengineKeibiEnginePkgAuthAPIRoleDetailsResponse struct {

	// Role Description and accesses
	// Example: The Administrator role is a super user role with all of the capabilities that can be assigned to a role, and its enables access to all data \u0026 configuration on a Kaytu Workspace. You cannot edit or delete the Administrator role.
	Description string `json:"description,omitempty"`

	// Name of the role
	// Example: admin
	// Enum: [admin editor viewer]
	Role GitlabComKeibiengineKeibiEnginePkgAuthAPIRole `json:"role,omitempty"`

	// Number of users having this role
	// Example: 1
	UserCount int64 `json:"userCount,omitempty"`

	// List of users having this role
	Users []*GitlabComKeibiengineKeibiEnginePkgAuthAPIGetUsersResponse `json:"users"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg auth api role details response
func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIRoleDetailsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var gitlabComKeibiengineKeibiEnginePkgAuthApiRoleDetailsResponseTypeRolePropEnum []interface{}

func init() {
	var res []GitlabComKeibiengineKeibiEnginePkgAuthAPIRole
	if err := json.Unmarshal([]byte(`["admin","editor","viewer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gitlabComKeibiengineKeibiEnginePkgAuthApiRoleDetailsResponseTypeRolePropEnum = append(gitlabComKeibiengineKeibiEnginePkgAuthApiRoleDetailsResponseTypeRolePropEnum, v)
	}
}

// prop value enum
func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIRoleDetailsResponse) validateRoleEnum(path, location string, value *struct {
	GitlabComKeibiengineKeibiEnginePkgAuthAPIRole
}) error {
	if err := validate.EnumCase(path, location, value, gitlabComKeibiengineKeibiEnginePkgAuthApiRoleDetailsResponseTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIRoleDetailsResponse) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIRoleDetailsResponse) validateUsers(formats strfmt.Registry) error {
	if swag.IsZero(m.Users) { // not required
		return nil
	}

	for i := 0; i < len(m.Users); i++ {
		if swag.IsZero(m.Users[i]) { // not required
			continue
		}

		if m.Users[i] != nil {
			if err := m.Users[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("users" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this gitlab com keibiengine keibi engine pkg auth api role details response based on the context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIRoleDetailsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRole(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIRoleDetailsResponse) contextValidateRole(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIRoleDetailsResponse) contextValidateUsers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Users); i++ {

		if m.Users[i] != nil {
			if err := m.Users[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("users" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIRoleDetailsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIRoleDetailsResponse) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgAuthAPIRoleDetailsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
