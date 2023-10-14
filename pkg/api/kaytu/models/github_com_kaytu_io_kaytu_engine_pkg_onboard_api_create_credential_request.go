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

// GithubComKaytuIoKaytuEnginePkgOnboardAPICreateCredentialRequest github com kaytu io kaytu engine pkg onboard api create credential request
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_onboard_api.CreateCredentialRequest
type GithubComKaytuIoKaytuEnginePkgOnboardAPICreateCredentialRequest struct {

	// config
	Config interface{} `json:"config,omitempty"`

	// source type
	// Example: Azure
	SourceType SourceType `json:"source_type,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg onboard api create credential request
func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICreateCredentialRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSourceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICreateCredentialRequest) validateSourceType(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceType) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this github com kaytu io kaytu engine pkg onboard api create credential request based on the context it is used
func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICreateCredentialRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSourceType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICreateCredentialRequest) contextValidateSourceType(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICreateCredentialRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICreateCredentialRequest) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgOnboardAPICreateCredentialRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
