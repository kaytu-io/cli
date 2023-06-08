// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GitlabComKeibiengineKeibiEnginePkgInventoryAPISmartQuerySortItem gitlab com keibiengine keibi engine pkg inventory api smart query sort item
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_inventory_api.SmartQuerySortItem
type GitlabComKeibiengineKeibiEnginePkgInventoryAPISmartQuerySortItem struct {

	// direction
	// Enum: [asc desc]
	Direction struct {
		GitlabComKeibiengineKeibiEnginePkgInventoryAPIDirectionType
	} `json:"direction,omitempty"`

	// fill this with column name
	Field string `json:"field,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg inventory api smart query sort item
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPISmartQuerySortItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var gitlabComKeibiengineKeibiEnginePkgInventoryApiSmartQuerySortItemTypeDirectionPropEnum []interface{}

func init() {
	var res []GitlabComKeibiengineKeibiEnginePkgInventoryAPIDirectionType
	if err := json.Unmarshal([]byte(`["asc","desc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gitlabComKeibiengineKeibiEnginePkgInventoryApiSmartQuerySortItemTypeDirectionPropEnum = append(gitlabComKeibiengineKeibiEnginePkgInventoryApiSmartQuerySortItemTypeDirectionPropEnum, v)
	}
}

// prop value enum
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPISmartQuerySortItem) validateDirectionEnum(path, location string, value *struct {
	GitlabComKeibiengineKeibiEnginePkgInventoryAPIDirectionType
}) error {
	if err := validate.EnumCase(path, location, value, gitlabComKeibiengineKeibiEnginePkgInventoryApiSmartQuerySortItemTypeDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPISmartQuerySortItem) validateDirection(formats strfmt.Registry) error {
	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this gitlab com keibiengine keibi engine pkg inventory api smart query sort item based on the context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPISmartQuerySortItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDirection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPISmartQuerySortItem) contextValidateDirection(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPISmartQuerySortItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPISmartQuerySortItem) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgInventoryAPISmartQuerySortItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}