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

// GithubComKaytuIoKaytuEnginePkgOnboardAPICreateAwsConnectionRequest github com kaytu io kaytu engine pkg onboard api create aws connection request
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_onboard_api.CreateAwsConnectionRequest
type GithubComKaytuIoKaytuEnginePkgOnboardAPICreateAwsConnectionRequest struct {

	// aws config
	AwsConfig *GithubComKaytuIoKaytuEnginePkgOnboardAPIV2AWSCredentialV2Config `json:"awsConfig,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg onboard api create aws connection request
func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICreateAwsConnectionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICreateAwsConnectionRequest) validateAwsConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsConfig) { // not required
		return nil
	}

	if m.AwsConfig != nil {
		if err := m.AwsConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this github com kaytu io kaytu engine pkg onboard api create aws connection request based on the context it is used
func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICreateAwsConnectionRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAwsConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICreateAwsConnectionRequest) contextValidateAwsConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsConfig != nil {

		if swag.IsZero(m.AwsConfig) { // not required
			return nil
		}

		if err := m.AwsConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICreateAwsConnectionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICreateAwsConnectionRequest) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgOnboardAPICreateAwsConnectionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
