// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GitlabComKeibiengineKeibiEnginePkgInventoryAPIAzureResource gitlab com keibiengine keibi engine pkg inventory api azure resource
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_inventory_api.AzureResource
type GitlabComKeibiengineKeibiEnginePkgInventoryAPIAzureResource struct {

	// attributes
	Attributes map[string]string `json:"attributes,omitempty"`

	// The Region of the resource
	Location string `json:"location,omitempty"`

	// Provider Connection Id
	ProviderConnectionID string `json:"providerConnectionID,omitempty"`

	// Provider Connection Name
	ProviderConnectionName string `json:"providerConnectionName,omitempty"`

	// Resource Category
	ResourceCategory string `json:"resourceCategory,omitempty"`

	// Resource Group
	ResourceGroup string `json:"resourceGroup,omitempty"`

	// Resource Id
	ResourceID string `json:"resourceID,omitempty"`

	// Resource Name
	ResourceName string `json:"resourceName,omitempty"`

	// Resource Type
	ResourceType string `json:"resourceType,omitempty"`

	// Resource Type Name
	ResourceTypeName string `json:"resourceTypeName,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg inventory api azure resource
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIAzureResource) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this gitlab com keibiengine keibi engine pkg inventory api azure resource based on context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIAzureResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIAzureResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIAzureResource) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgInventoryAPIAzureResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
