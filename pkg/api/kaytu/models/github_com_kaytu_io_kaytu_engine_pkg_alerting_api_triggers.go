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

// GithubComKaytuIoKaytuEnginePkgAlertingAPITriggers github com kaytu io kaytu engine pkg alerting api triggers
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_alerting_api.Triggers
type GithubComKaytuIoKaytuEnginePkgAlertingAPITriggers struct {

	// action
	Action *GithubComKaytuIoKaytuEnginePkgAlertingAPIAction `json:"action,omitempty"`

	// response status
	ResponseStatus int64 `json:"response_status,omitempty"`

	// rule
	Rule *GithubComKaytuIoKaytuEnginePkgAlertingAPIRule `json:"rule,omitempty"`

	// triggered at
	TriggeredAt string `json:"triggered_at,omitempty"`

	// value
	Value int64 `json:"value,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg alerting api triggers
func (m *GithubComKaytuIoKaytuEnginePkgAlertingAPITriggers) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgAlertingAPITriggers) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if m.Action != nil {
		if err := m.Action.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("action")
			}
			return err
		}
	}

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgAlertingAPITriggers) validateRule(formats strfmt.Registry) error {
	if swag.IsZero(m.Rule) { // not required
		return nil
	}

	if m.Rule != nil {
		if err := m.Rule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rule")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this github com kaytu io kaytu engine pkg alerting api triggers based on the context it is used
func (m *GithubComKaytuIoKaytuEnginePkgAlertingAPITriggers) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgAlertingAPITriggers) contextValidateAction(ctx context.Context, formats strfmt.Registry) error {

	if m.Action != nil {

		if swag.IsZero(m.Action) { // not required
			return nil
		}

		if err := m.Action.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("action")
			}
			return err
		}
	}

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgAlertingAPITriggers) contextValidateRule(ctx context.Context, formats strfmt.Registry) error {

	if m.Rule != nil {

		if swag.IsZero(m.Rule) { // not required
			return nil
		}

		if err := m.Rule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgAlertingAPITriggers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgAlertingAPITriggers) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgAlertingAPITriggers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
