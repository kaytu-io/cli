// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GitlabComKeibiengineKeibiEnginePkgInventoryAPIRunQueryRequest gitlab com keibiengine keibi engine pkg inventory api run query request
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_inventory_api.RunQueryRequest
type GitlabComKeibiengineKeibiEnginePkgInventoryAPIRunQueryRequest struct {

	// page
	// Required: true
	Page *GitlabComKeibiengineKeibiEnginePkgInventoryAPIPage `json:"page"`

	// NOTE: we don't support multi-field sort for now, if sort is empty, results would be sorted by first column
	Sorts []*GitlabComKeibiengineKeibiEnginePkgInventoryAPISmartQuerySortItem `json:"sorts"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg inventory api run query request
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIRunQueryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSorts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIRunQueryRequest) validatePage(formats strfmt.Registry) error {

	if err := validate.Required("page", "body", m.Page); err != nil {
		return err
	}

	if m.Page != nil {
		if err := m.Page.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("page")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("page")
			}
			return err
		}
	}

	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIRunQueryRequest) validateSorts(formats strfmt.Registry) error {
	if swag.IsZero(m.Sorts) { // not required
		return nil
	}

	for i := 0; i < len(m.Sorts); i++ {
		if swag.IsZero(m.Sorts[i]) { // not required
			continue
		}

		if m.Sorts[i] != nil {
			if err := m.Sorts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sorts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sorts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this gitlab com keibiengine keibi engine pkg inventory api run query request based on the context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIRunQueryRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIRunQueryRequest) contextValidatePage(ctx context.Context, formats strfmt.Registry) error {

	if m.Page != nil {
		if err := m.Page.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("page")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("page")
			}
			return err
		}
	}

	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIRunQueryRequest) contextValidateSorts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Sorts); i++ {

		if m.Sorts[i] != nil {
			if err := m.Sorts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sorts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sorts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIRunQueryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIRunQueryRequest) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgInventoryAPIRunQueryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}