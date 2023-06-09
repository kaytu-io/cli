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

// GitlabComKeibiengineKeibiEnginePkgInventoryAPIConnectionResourceCountResponse gitlab com keibiengine keibi engine pkg inventory api connection resource count response
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_inventory_api.ConnectionResourceCountResponse
type GitlabComKeibiengineKeibiEnginePkgInventoryAPIConnectionResourceCountResponse struct {

	// Source Type
	Connector SourceType `json:"connector,omitempty"`

	// Provider Connection Id
	ConnectorConnectionID string `json:"connectorConnectionID,omitempty"`

	// Provider Connection Name
	ConnectorConnectionName string `json:"connectorConnectionName,omitempty"`

	// last inventory
	LastInventory string `json:"lastInventory,omitempty"`

	// lifecycle state
	LifecycleState string `json:"lifecycleState,omitempty"`

	// onboard date
	OnboardDate string `json:"onboardDate,omitempty"`

	// Number of resources
	ResourceCount int64 `json:"resourceCount,omitempty"`

	// Source Id
	SourceID string `json:"sourceID,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg inventory api connection resource count response
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIConnectionResourceCountResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIConnectionResourceCountResponse) validateConnector(formats strfmt.Registry) error {
	if swag.IsZero(m.Connector) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this gitlab com keibiengine keibi engine pkg inventory api connection resource count response based on the context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIConnectionResourceCountResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIConnectionResourceCountResponse) contextValidateConnector(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIConnectionResourceCountResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIConnectionResourceCountResponse) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgInventoryAPIConnectionResourceCountResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
