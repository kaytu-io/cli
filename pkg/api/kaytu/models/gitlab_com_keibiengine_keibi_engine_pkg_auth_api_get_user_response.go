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

// GitlabComKeibiengineKeibiEnginePkgAuthAPIGetUserResponse gitlab com keibiengine keibi engine pkg auth api get user response
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_auth_api.GetUserResponse
type GitlabComKeibiengineKeibiEnginePkgAuthAPIGetUserResponse struct {

	// Is the user blocked or not
	// Example: false
	Blocked bool `json:"blocked,omitempty"`

	// Creation timestamp in UTC
	// Example: 2023-03-31T09:36:09.855Z
	CreatedAt string `json:"createdAt,omitempty"`

	// Email address of the user
	// Example: sample@gmail.com
	Email string `json:"email,omitempty"`

	// Is email verified or not
	// Example: true
	EmailVerified bool `json:"emailVerified,omitempty"`

	// Last activity timestamp in UTC
	// Example: 2023-04-21T08:53:09.928Z
	LastActivity string `json:"lastActivity,omitempty"`

	// Name of the role
	// Example: admin
	// Enum: [admin editor viewer]
	RoleName GitlabComKeibiengineKeibiEnginePkgAuthAPIRole `json:"roleName,omitempty"`

	// Invite status
	// Example: pending
	// Enum: [accepted pending]
	Status GitlabComKeibiengineKeibiEnginePkgAuthAPIInviteStatus `json:"status,omitempty"`

	// Unique identifier for the user
	// Example: sampleID
	UserID string `json:"userId,omitempty"`

	// Username
	// Example: sampleName
	UserName string `json:"userName,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg auth api get user response
func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIGetUserResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRoleName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var gitlabComKeibiengineKeibiEnginePkgAuthApiGetUserResponseTypeRoleNamePropEnum []interface{}

func init() {
	var res []GitlabComKeibiengineKeibiEnginePkgAuthAPIRole
	if err := json.Unmarshal([]byte(`["admin","editor","viewer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gitlabComKeibiengineKeibiEnginePkgAuthApiGetUserResponseTypeRoleNamePropEnum = append(gitlabComKeibiengineKeibiEnginePkgAuthApiGetUserResponseTypeRoleNamePropEnum, v)
	}
}

// prop value enum
func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIGetUserResponse) validateRoleNameEnum(path, location string, value *struct {
	GitlabComKeibiengineKeibiEnginePkgAuthAPIRole
}) error {
	if err := validate.EnumCase(path, location, value, gitlabComKeibiengineKeibiEnginePkgAuthApiGetUserResponseTypeRoleNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIGetUserResponse) validateRoleName(formats strfmt.Registry) error {
	if swag.IsZero(m.RoleName) { // not required
		return nil
	}

	return nil
}

var gitlabComKeibiengineKeibiEnginePkgAuthApiGetUserResponseTypeStatusPropEnum []interface{}

func init() {
	var res []GitlabComKeibiengineKeibiEnginePkgAuthAPIInviteStatus
	if err := json.Unmarshal([]byte(`["accepted","pending"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gitlabComKeibiengineKeibiEnginePkgAuthApiGetUserResponseTypeStatusPropEnum = append(gitlabComKeibiengineKeibiEnginePkgAuthApiGetUserResponseTypeStatusPropEnum, v)
	}
}

// prop value enum
func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIGetUserResponse) validateStatusEnum(path, location string, value *struct {
	GitlabComKeibiengineKeibiEnginePkgAuthAPIInviteStatus
}) error {
	if err := validate.EnumCase(path, location, value, gitlabComKeibiengineKeibiEnginePkgAuthApiGetUserResponseTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIGetUserResponse) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this gitlab com keibiengine keibi engine pkg auth api get user response based on the context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIGetUserResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRoleName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIGetUserResponse) contextValidateRoleName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIGetUserResponse) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIGetUserResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIGetUserResponse) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgAuthAPIGetUserResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
