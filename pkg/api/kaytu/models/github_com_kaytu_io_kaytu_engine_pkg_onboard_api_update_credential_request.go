// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GithubComKaytuIoKaytuEnginePkgOnboardAPIUpdateCredentialRequest github com kaytu io kaytu engine pkg onboard api update credential request
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_onboard_api.UpdateCredentialRequest
type GithubComKaytuIoKaytuEnginePkgOnboardAPIUpdateCredentialRequest struct {

	// config
	Config interface{} `json:"config,omitempty"`

	// connector
	// Example: Azure
	Connector SourceType `json:"connector,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg onboard api update credential request
func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPIUpdateCredentialRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPIUpdateCredentialRequest) validateConnector(formats strfmt.Registry) error {
	if swag.IsZero(m.Connector) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this github com kaytu io kaytu engine pkg onboard api update credential request based on the context it is used
func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPIUpdateCredentialRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPIUpdateCredentialRequest) contextValidateConnector(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPIUpdateCredentialRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPIUpdateCredentialRequest) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgOnboardAPIUpdateCredentialRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
