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

// GitlabComKeibiengineKeibiEnginePkgOnboardAPICatalogConnector gitlab com keibiengine keibi engine pkg onboard api catalog connector
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_onboard_api.CatalogConnector
type GitlabComKeibiengineKeibiEnginePkgOnboardAPICatalogConnector struct {

	// ID
	ID int64 `json:"ID,omitempty"`

	// allow new connections
	AllowNewConnections bool `json:"allowNewConnections,omitempty"`

	// category
	Category string `json:"category,omitempty"`

	// connection count
	ConnectionCount int64 `json:"connectionCount,omitempty"`

	// connection federator
	ConnectionFederator string `json:"connectionFederator,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// logo
	Logo string `json:"logo,omitempty"`

	// max connections limit
	MaxConnectionsLimit int64 `json:"maxConnectionsLimit,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// source type
	SourceType SourceType `json:"sourceType,omitempty"`

	// state
	State GitlabComKeibiengineKeibiEnginePkgOnboardAPIConnectorState `json:"state,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg onboard api catalog connector
func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPICatalogConnector) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSourceType(formats); err != nil {
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

func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPICatalogConnector) validateSourceType(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceType) { // not required
		return nil
	}

	if err := m.SourceType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("sourceType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("sourceType")
		}
		return err
	}

	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPICatalogConnector) validateState(formats strfmt.Registry) error {
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

// ContextValidate validate this gitlab com keibiengine keibi engine pkg onboard api catalog connector based on the context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPICatalogConnector) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSourceType(ctx, formats); err != nil {
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

func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPICatalogConnector) contextValidateSourceType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.SourceType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("sourceType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("sourceType")
		}
		return err
	}

	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPICatalogConnector) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPICatalogConnector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPICatalogConnector) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgOnboardAPICatalogConnector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}