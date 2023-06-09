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

// GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetFiltersResponse gitlab com keibiengine keibi engine pkg inventory api get filters response
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_inventory_api.GetFiltersResponse
type GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetFiltersResponse struct {

	// search filters
	Filters GitlabComKeibiengineKeibiEnginePkgInventoryAPIResourceFiltersResponse `json:"filters,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg inventory api get filters response
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetFiltersResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetFiltersResponse) validateFilters(formats strfmt.Registry) error {
	if swag.IsZero(m.Filters) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this gitlab com keibiengine keibi engine pkg inventory api get filters response based on the context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetFiltersResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetFiltersResponse) contextValidateFilters(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetFiltersResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetFiltersResponse) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetFiltersResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
