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

// GithubComKaytuIoKaytuEnginePkgAlertingAPIScope github com kaytu io kaytu engine pkg alerting api scope
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_alerting_api.Scope
type GithubComKaytuIoKaytuEnginePkgAlertingAPIScope struct {

	// connection group
	ConnectionGroup string `json:"connection_group,omitempty"`

	// connection id
	ConnectionID string `json:"connection_id,omitempty"`

	// connector
	Connector SourceType `json:"connector,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg alerting api scope
func (m *GithubComKaytuIoKaytuEnginePkgAlertingAPIScope) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgAlertingAPIScope) validateConnector(formats strfmt.Registry) error {
	if swag.IsZero(m.Connector) { // not required
		return nil
	}

	if err := m.Connector.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("connector")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("connector")
		}
		return err
	}

	return nil
}

// ContextValidate validate this github com kaytu io kaytu engine pkg alerting api scope based on the context it is used
func (m *GithubComKaytuIoKaytuEnginePkgAlertingAPIScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgAlertingAPIScope) contextValidateConnector(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Connector) { // not required
		return nil
	}

	if err := m.Connector.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("connector")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("connector")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgAlertingAPIScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgAlertingAPIScope) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgAlertingAPIScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
