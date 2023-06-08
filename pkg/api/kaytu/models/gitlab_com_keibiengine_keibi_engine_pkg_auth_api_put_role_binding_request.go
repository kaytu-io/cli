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

// GitlabComKeibiengineKeibiEnginePkgAuthAPIPutRoleBindingRequest gitlab com keibiengine keibi engine pkg auth api put role binding request
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_auth_api.PutRoleBindingRequest
type GitlabComKeibiengineKeibiEnginePkgAuthAPIPutRoleBindingRequest struct {

	// Name of the role
	// Example: admin
	// Required: true
	// Enum: [admin editor viewer]
	RoleName struct {
		GitlabComKeibiengineKeibiEnginePkgAuthAPIRole
	} `json:"roleName"`

	// Unique identifier for the User
	// Example: sampleID
	// Required: true
	UserID *string `json:"userId"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg auth api put role binding request
func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIPutRoleBindingRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRoleName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var gitlabComKeibiengineKeibiEnginePkgAuthApiPutRoleBindingRequestTypeRoleNamePropEnum []interface{}

func init() {
	var res []GitlabComKeibiengineKeibiEnginePkgAuthAPIRole
	if err := json.Unmarshal([]byte(`["admin","editor","viewer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gitlabComKeibiengineKeibiEnginePkgAuthApiPutRoleBindingRequestTypeRoleNamePropEnum = append(gitlabComKeibiengineKeibiEnginePkgAuthApiPutRoleBindingRequestTypeRoleNamePropEnum, v)
	}
}

// prop value enum
func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIPutRoleBindingRequest) validateRoleNameEnum(path, location string, value *struct {
	GitlabComKeibiengineKeibiEnginePkgAuthAPIRole
}) error {
	if err := validate.EnumCase(path, location, value, gitlabComKeibiengineKeibiEnginePkgAuthApiPutRoleBindingRequestTypeRoleNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIPutRoleBindingRequest) validateRoleName(formats strfmt.Registry) error {

	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIPutRoleBindingRequest) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this gitlab com keibiengine keibi engine pkg auth api put role binding request based on the context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIPutRoleBindingRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRoleName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIPutRoleBindingRequest) contextValidateRoleName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIPutRoleBindingRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgAuthAPIPutRoleBindingRequest) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgAuthAPIPutRoleBindingRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}