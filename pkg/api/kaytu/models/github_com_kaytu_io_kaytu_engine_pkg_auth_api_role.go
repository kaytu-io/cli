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

// GithubComKaytuIoKaytuEnginePkgAuthAPIRole github com kaytu io kaytu engine pkg auth api role
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_auth_api.Role
type GithubComKaytuIoKaytuEnginePkgAuthAPIRole string

func NewGithubComKaytuIoKaytuEnginePkgAuthAPIRole(value GithubComKaytuIoKaytuEnginePkgAuthAPIRole) *GithubComKaytuIoKaytuEnginePkgAuthAPIRole {
	return &value
}

// Pointer returns a pointer to a freshly-allocated GithubComKaytuIoKaytuEnginePkgAuthAPIRole.
func (m GithubComKaytuIoKaytuEnginePkgAuthAPIRole) Pointer() *GithubComKaytuIoKaytuEnginePkgAuthAPIRole {
	return &m
}

const (

	// GithubComKaytuIoKaytuEnginePkgAuthAPIRoleAdmin captures enum value "admin"
	GithubComKaytuIoKaytuEnginePkgAuthAPIRoleAdmin GithubComKaytuIoKaytuEnginePkgAuthAPIRole = "admin"

	// GithubComKaytuIoKaytuEnginePkgAuthAPIRoleEditor captures enum value "editor"
	GithubComKaytuIoKaytuEnginePkgAuthAPIRoleEditor GithubComKaytuIoKaytuEnginePkgAuthAPIRole = "editor"

	// GithubComKaytuIoKaytuEnginePkgAuthAPIRoleViewer captures enum value "viewer"
	GithubComKaytuIoKaytuEnginePkgAuthAPIRoleViewer GithubComKaytuIoKaytuEnginePkgAuthAPIRole = "viewer"
)

// for schema
var githubComKaytuIoKaytuEnginePkgAuthApiRoleEnum []interface{}

func init() {
	var res []GithubComKaytuIoKaytuEnginePkgAuthAPIRole
	if err := json.Unmarshal([]byte(`["admin","editor","viewer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		githubComKaytuIoKaytuEnginePkgAuthApiRoleEnum = append(githubComKaytuIoKaytuEnginePkgAuthApiRoleEnum, v)
	}
}

func (m GithubComKaytuIoKaytuEnginePkgAuthAPIRole) validateGithubComKaytuIoKaytuEnginePkgAuthAPIRoleEnum(path, location string, value GithubComKaytuIoKaytuEnginePkgAuthAPIRole) error {
	if err := validate.EnumCase(path, location, value, githubComKaytuIoKaytuEnginePkgAuthApiRoleEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this github com kaytu io kaytu engine pkg auth api role
func (m GithubComKaytuIoKaytuEnginePkgAuthAPIRole) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateGithubComKaytuIoKaytuEnginePkgAuthAPIRoleEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this github com kaytu io kaytu engine pkg auth api role based on context it is used
func (m GithubComKaytuIoKaytuEnginePkgAuthAPIRole) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
