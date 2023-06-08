// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GitlabComKeibiengineKeibiEnginePkgInventoryAPITopServiceCostResponse gitlab com keibiengine keibi engine pkg inventory api top service cost response
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_inventory_api.TopServiceCostResponse
type GitlabComKeibiengineKeibiEnginePkgInventoryAPITopServiceCostResponse struct {

	// Service Cost
	Cost float64 `json:"cost,omitempty"`

	// Service Name
	ServiceName string `json:"serviceName,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg inventory api top service cost response
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPITopServiceCostResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this gitlab com keibiengine keibi engine pkg inventory api top service cost response based on context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPITopServiceCostResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPITopServiceCostResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPITopServiceCostResponse) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgInventoryAPITopServiceCostResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}