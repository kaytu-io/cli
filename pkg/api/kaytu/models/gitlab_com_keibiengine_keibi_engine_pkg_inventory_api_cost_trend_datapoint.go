// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GitlabComKeibiengineKeibiEnginePkgInventoryAPICostTrendDatapoint gitlab com keibiengine keibi engine pkg inventory api cost trend datapoint
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_inventory_api.CostTrendDatapoint
type GitlabComKeibiengineKeibiEnginePkgInventoryAPICostTrendDatapoint struct {

	// count
	Count float64 `json:"count,omitempty"`

	// date
	Date string `json:"date,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg inventory api cost trend datapoint
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPICostTrendDatapoint) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this gitlab com keibiengine keibi engine pkg inventory api cost trend datapoint based on context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPICostTrendDatapoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPICostTrendDatapoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPICostTrendDatapoint) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgInventoryAPICostTrendDatapoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}