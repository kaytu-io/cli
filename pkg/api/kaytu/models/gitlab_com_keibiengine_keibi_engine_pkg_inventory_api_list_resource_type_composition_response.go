// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GitlabComKeibiengineKeibiEnginePkgInventoryAPIListResourceTypeCompositionResponse gitlab com keibiengine keibi engine pkg inventory api list resource type composition response
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_inventory_api.ListResourceTypeCompositionResponse
type GitlabComKeibiengineKeibiEnginePkgInventoryAPIListResourceTypeCompositionResponse struct {

	// others
	Others int64 `json:"others,omitempty"`

	// top values
	TopValues map[string]int64 `json:"top_values,omitempty"`

	// total count
	TotalCount int64 `json:"total_count,omitempty"`

	// total value count
	TotalValueCount int64 `json:"total_value_count,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg inventory api list resource type composition response
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIListResourceTypeCompositionResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this gitlab com keibiengine keibi engine pkg inventory api list resource type composition response based on context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIListResourceTypeCompositionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIListResourceTypeCompositionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIListResourceTypeCompositionResponse) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgInventoryAPIListResourceTypeCompositionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}