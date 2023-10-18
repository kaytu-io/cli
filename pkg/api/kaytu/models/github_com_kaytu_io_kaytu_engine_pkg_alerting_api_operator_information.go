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

// GithubComKaytuIoKaytuEnginePkgAlertingAPIOperatorInformation github com kaytu io kaytu engine pkg alerting api operator information
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_alerting_api.OperatorInformation
type GithubComKaytuIoKaytuEnginePkgAlertingAPIOperatorInformation struct {

	// operator type
	OperatorType GithubComKaytuIoKaytuEnginePkgAlertingAPIOperatorType `json:"operator_type,omitempty"`

	// value
	Value int64 `json:"value,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg alerting api operator information
func (m *GithubComKaytuIoKaytuEnginePkgAlertingAPIOperatorInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperatorType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgAlertingAPIOperatorInformation) validateOperatorType(formats strfmt.Registry) error {
	if swag.IsZero(m.OperatorType) { // not required
		return nil
	}

	if err := m.OperatorType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("operator_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("operator_type")
		}
		return err
	}

	return nil
}

// ContextValidate validate this github com kaytu io kaytu engine pkg alerting api operator information based on the context it is used
func (m *GithubComKaytuIoKaytuEnginePkgAlertingAPIOperatorInformation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOperatorType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgAlertingAPIOperatorInformation) contextValidateOperatorType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.OperatorType) { // not required
		return nil
	}

	if err := m.OperatorType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("operator_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("operator_type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgAlertingAPIOperatorInformation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgAlertingAPIOperatorInformation) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgAlertingAPIOperatorInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}