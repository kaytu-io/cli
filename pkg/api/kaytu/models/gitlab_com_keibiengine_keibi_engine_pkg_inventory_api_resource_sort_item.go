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

// GitlabComKeibiengineKeibiEnginePkgInventoryAPIResourceSortItem gitlab com keibiengine keibi engine pkg inventory api resource sort item
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_inventory_api.ResourceSortItem
type GitlabComKeibiengineKeibiEnginePkgInventoryAPIResourceSortItem struct {

	// direction
	// Enum: [asc desc]
	Direction struct {
		GitlabComKeibiengineKeibiEnginePkgInventoryAPIDirectionType
	} `json:"direction,omitempty"`

	// field
	// Enum: [resourceID resourceName provider resourceType resourceGroup location connectionID]
	Field struct {
		GitlabComKeibiengineKeibiEnginePkgInventoryAPISortFieldType
	} `json:"field,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg inventory api resource sort item
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIResourceSortItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateField(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var gitlabComKeibiengineKeibiEnginePkgInventoryApiResourceSortItemTypeDirectionPropEnum []interface{}

func init() {
	var res []GitlabComKeibiengineKeibiEnginePkgInventoryAPIDirectionType
	if err := json.Unmarshal([]byte(`["asc","desc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gitlabComKeibiengineKeibiEnginePkgInventoryApiResourceSortItemTypeDirectionPropEnum = append(gitlabComKeibiengineKeibiEnginePkgInventoryApiResourceSortItemTypeDirectionPropEnum, v)
	}
}

// prop value enum
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIResourceSortItem) validateDirectionEnum(path, location string, value *struct {
	GitlabComKeibiengineKeibiEnginePkgInventoryAPIDirectionType
}) error {
	if err := validate.EnumCase(path, location, value, gitlabComKeibiengineKeibiEnginePkgInventoryApiResourceSortItemTypeDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIResourceSortItem) validateDirection(formats strfmt.Registry) error {
	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	return nil
}

var gitlabComKeibiengineKeibiEnginePkgInventoryApiResourceSortItemTypeFieldPropEnum []interface{}

func init() {
	var res []GitlabComKeibiengineKeibiEnginePkgInventoryAPISortFieldType
	if err := json.Unmarshal([]byte(`["resourceID","resourceName","provider","resourceType","resourceGroup","location","connectionID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gitlabComKeibiengineKeibiEnginePkgInventoryApiResourceSortItemTypeFieldPropEnum = append(gitlabComKeibiengineKeibiEnginePkgInventoryApiResourceSortItemTypeFieldPropEnum, v)
	}
}

// prop value enum
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIResourceSortItem) validateFieldEnum(path, location string, value *struct {
	GitlabComKeibiengineKeibiEnginePkgInventoryAPISortFieldType
}) error {
	if err := validate.EnumCase(path, location, value, gitlabComKeibiengineKeibiEnginePkgInventoryApiResourceSortItemTypeFieldPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIResourceSortItem) validateField(formats strfmt.Registry) error {
	if swag.IsZero(m.Field) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this gitlab com keibiengine keibi engine pkg inventory api resource sort item based on the context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIResourceSortItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDirection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateField(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIResourceSortItem) contextValidateDirection(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIResourceSortItem) contextValidateField(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIResourceSortItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIResourceSortItem) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgInventoryAPIResourceSortItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}