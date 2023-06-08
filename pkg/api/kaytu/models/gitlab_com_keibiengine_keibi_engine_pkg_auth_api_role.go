// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// GitlabComKeibiengineKeibiEnginePkgAuthAPIRole gitlab com keibiengine keibi engine pkg auth api role
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_auth_api.Role
type GitlabComKeibiengineKeibiEnginePkgAuthAPIRole string

func NewGitlabComKeibiengineKeibiEnginePkgAuthAPIRole(value GitlabComKeibiengineKeibiEnginePkgAuthAPIRole) *GitlabComKeibiengineKeibiEnginePkgAuthAPIRole {
	return &value
}

// Pointer returns a pointer to a freshly-allocated GitlabComKeibiengineKeibiEnginePkgAuthAPIRole.
func (m GitlabComKeibiengineKeibiEnginePkgAuthAPIRole) Pointer() *GitlabComKeibiengineKeibiEnginePkgAuthAPIRole {
	return &m
}

const (

	// GitlabComKeibiengineKeibiEnginePkgAuthAPIRoleAdmin captures enum value "admin"
	GitlabComKeibiengineKeibiEnginePkgAuthAPIRoleAdmin GitlabComKeibiengineKeibiEnginePkgAuthAPIRole = "admin"

	// GitlabComKeibiengineKeibiEnginePkgAuthAPIRoleEditor captures enum value "editor"
	GitlabComKeibiengineKeibiEnginePkgAuthAPIRoleEditor GitlabComKeibiengineKeibiEnginePkgAuthAPIRole = "editor"

	// GitlabComKeibiengineKeibiEnginePkgAuthAPIRoleViewer captures enum value "viewer"
	GitlabComKeibiengineKeibiEnginePkgAuthAPIRoleViewer GitlabComKeibiengineKeibiEnginePkgAuthAPIRole = "viewer"
)

// for schema
var gitlabComKeibiengineKeibiEnginePkgAuthApiRoleEnum []interface{}

func init() {
	var res []GitlabComKeibiengineKeibiEnginePkgAuthAPIRole
	if err := json.Unmarshal([]byte(`["admin","editor","viewer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gitlabComKeibiengineKeibiEnginePkgAuthApiRoleEnum = append(gitlabComKeibiengineKeibiEnginePkgAuthApiRoleEnum, v)
	}
}

func (m GitlabComKeibiengineKeibiEnginePkgAuthAPIRole) validateGitlabComKeibiengineKeibiEnginePkgAuthAPIRoleEnum(path, location string, value GitlabComKeibiengineKeibiEnginePkgAuthAPIRole) error {
	if err := validate.EnumCase(path, location, value, gitlabComKeibiengineKeibiEnginePkgAuthApiRoleEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this gitlab com keibiengine keibi engine pkg auth api role
func (m GitlabComKeibiengineKeibiEnginePkgAuthAPIRole) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateGitlabComKeibiengineKeibiEnginePkgAuthAPIRoleEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this gitlab com keibiengine keibi engine pkg auth api role based on context it is used
func (m GitlabComKeibiengineKeibiEnginePkgAuthAPIRole) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
