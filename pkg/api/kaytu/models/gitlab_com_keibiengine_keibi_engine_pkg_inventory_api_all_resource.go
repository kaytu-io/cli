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

// GitlabComKeibiengineKeibiEnginePkgInventoryAPIAllResource gitlab com keibiengine keibi engine pkg inventory api all resource
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_inventory_api.AllResource
type GitlabComKeibiengineKeibiEnginePkgInventoryAPIAllResource struct {

	// attributes
	Attributes map[string]string `json:"attributes,omitempty"`

	// The Region of the resource
	Location string `json:"location,omitempty"`

	// Resource Provider
	Provider GitlabComKeibiengineKeibiEnginePkgInventoryAPISourceType `json:"provider,omitempty"`

	// Provider Connection Id
	ProviderConnectionID string `json:"providerConnectionID,omitempty"`

	// Provider Connection Name
	ProviderConnectionName string `json:"providerConnectionName,omitempty"`

	// Resource Category
	ResourceCategory string `json:"resourceCategory,omitempty"`

	// Resource Id
	ResourceID string `json:"resourceID,omitempty"`

	// Resource Name
	ResourceName string `json:"resourceName,omitempty"`

	// Resource Type
	ResourceType string `json:"resourceType,omitempty"`

	// Resource Type Name
	ResourceTypeName string `json:"resourceTypeName,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg inventory api all resource
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIAllResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIAllResource) validateProvider(formats strfmt.Registry) error {
	if swag.IsZero(m.Provider) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this gitlab com keibiengine keibi engine pkg inventory api all resource based on the context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIAllResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProvider(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIAllResource) contextValidateProvider(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIAllResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIAllResource) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgInventoryAPIAllResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
