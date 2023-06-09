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

// GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetFiltersRequest gitlab com keibiengine keibi engine pkg inventory api get filters request
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_inventory_api.GetFiltersRequest
type GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetFiltersRequest struct {

	// search filters
	// Required: true
	Filters GitlabComKeibiengineKeibiEnginePkgInventoryAPIResourceFilters `json:"filters"`

	// search query
	Query string `json:"query,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg inventory api get filters request
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetFiltersRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetFiltersRequest) validateFilters(formats strfmt.Registry) error {

	return nil
}

// ContextValidate validate this gitlab com keibiengine keibi engine pkg inventory api get filters request based on the context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetFiltersRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetFiltersRequest) contextValidateFilters(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetFiltersRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetFiltersRequest) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetFiltersRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
