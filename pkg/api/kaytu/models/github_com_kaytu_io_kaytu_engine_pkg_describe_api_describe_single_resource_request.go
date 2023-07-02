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

// GithubComKaytuIoKaytuEnginePkgDescribeAPIDescribeSingleResourceRequest github com kaytu io kaytu engine pkg describe api describe single resource request
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_describe_api.DescribeSingleResourceRequest
type GithubComKaytuIoKaytuEnginePkgDescribeAPIDescribeSingleResourceRequest struct {

	// access key
	AccessKey string `json:"accessKey,omitempty"`

	// account ID
	AccountID string `json:"accountID,omitempty"`

	// additional fields
	AdditionalFields map[string]string `json:"additionalFields,omitempty"`

	// provider
	Provider SourceType `json:"provider,omitempty"`

	// resource type
	ResourceType string `json:"resourceType,omitempty"`

	// secret key
	SecretKey string `json:"secretKey,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg describe api describe single resource request
func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIDescribeSingleResourceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIDescribeSingleResourceRequest) validateProvider(formats strfmt.Registry) error {
	if swag.IsZero(m.Provider) { // not required
		return nil
	}

	if err := m.Provider.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("provider")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("provider")
		}
		return err
	}

	return nil
}

// ContextValidate validate this github com kaytu io kaytu engine pkg describe api describe single resource request based on the context it is used
func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIDescribeSingleResourceRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProvider(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIDescribeSingleResourceRequest) contextValidateProvider(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Provider) { // not required
		return nil
	}

	if err := m.Provider.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("provider")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("provider")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIDescribeSingleResourceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIDescribeSingleResourceRequest) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgDescribeAPIDescribeSingleResourceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}