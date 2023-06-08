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

// GitlabComKeibiengineKeibiEnginePkgOnboardAPIConnectionCountRequest gitlab com keibiengine keibi engine pkg onboard api connection count request
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_onboard_api.ConnectionCountRequest
type GitlabComKeibiengineKeibiEnginePkgOnboardAPIConnectionCountRequest struct {

	// connectors
	Connectors []string `json:"connectors"`

	// health
	Health SourceHealthStatus `json:"health,omitempty"`

	// state
	State GitlabComKeibiengineKeibiEnginePkgOnboardAPIConnectionState `json:"state,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg onboard api connection count request
func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPIConnectionCountRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHealth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPIConnectionCountRequest) validateHealth(formats strfmt.Registry) error {
	if swag.IsZero(m.Health) { // not required
		return nil
	}

	if err := m.Health.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("health")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("health")
		}
		return err
	}

	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPIConnectionCountRequest) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("state")
		}
		return err
	}

	return nil
}

// ContextValidate validate this gitlab com keibiengine keibi engine pkg onboard api connection count request based on the context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPIConnectionCountRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHealth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPIConnectionCountRequest) contextValidateHealth(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Health.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("health")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("health")
		}
		return err
	}

	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPIConnectionCountRequest) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := m.State.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPIConnectionCountRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPIConnectionCountRequest) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgOnboardAPIConnectionCountRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}