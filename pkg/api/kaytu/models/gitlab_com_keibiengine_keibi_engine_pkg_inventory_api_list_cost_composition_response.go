// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GitlabComKeibiengineKeibiEnginePkgInventoryAPIListCostCompositionResponse gitlab com keibiengine keibi engine pkg inventory api list cost composition response
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_inventory_api.ListCostCompositionResponse
type GitlabComKeibiengineKeibiEnginePkgInventoryAPIListCostCompositionResponse struct {

	// others
	Others float64 `json:"others,omitempty"`

	// top values
	TopValues map[string]float64 `json:"top_values,omitempty"`

	// total cost value
	TotalCostValue float64 `json:"total_cost_value,omitempty"`

	// total count
	TotalCount int64 `json:"total_count,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg inventory api list cost composition response
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIListCostCompositionResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this gitlab com keibiengine keibi engine pkg inventory api list cost composition response based on context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIListCostCompositionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIListCostCompositionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIListCostCompositionResponse) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgInventoryAPIListCostCompositionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
