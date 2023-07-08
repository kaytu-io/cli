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

// GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceSortItem github com kaytu io kaytu engine pkg inventory api resource sort item
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_inventory_api.ResourceSortItem
type GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceSortItem struct {

	// direction
	// Enum: [asc desc]
	Direction GithubComKaytuIoKaytuEnginePkgInventoryAPIDirectionType `json:"direction,omitempty"`

	// field
	// Enum: [resourceID connector resourceType resourceGroup location connectionID]
	Field GithubComKaytuIoKaytuEnginePkgInventoryAPISortFieldType `json:"field,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg inventory api resource sort item
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceSortItem) Validate(formats strfmt.Registry) error {
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

var githubComKaytuIoKaytuEnginePkgInventoryApiResourceSortItemTypeDirectionPropEnum []interface{}

func init() {
	var res []GithubComKaytuIoKaytuEnginePkgInventoryAPIDirectionType
	if err := json.Unmarshal([]byte(`["asc","desc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		githubComKaytuIoKaytuEnginePkgInventoryApiResourceSortItemTypeDirectionPropEnum = append(githubComKaytuIoKaytuEnginePkgInventoryApiResourceSortItemTypeDirectionPropEnum, v)
	}
}

// prop value enum
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceSortItem) validateDirectionEnum(path, location string, value *struct {
	GithubComKaytuIoKaytuEnginePkgInventoryAPIDirectionType
}) error {
	if err := validate.EnumCase(path, location, value, githubComKaytuIoKaytuEnginePkgInventoryApiResourceSortItemTypeDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceSortItem) validateDirection(formats strfmt.Registry) error {
	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	return nil
}

var githubComKaytuIoKaytuEnginePkgInventoryApiResourceSortItemTypeFieldPropEnum []interface{}

func init() {
	var res []GithubComKaytuIoKaytuEnginePkgInventoryAPISortFieldType
	if err := json.Unmarshal([]byte(`["resourceID","connector","resourceType","resourceGroup","location","connectionID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		githubComKaytuIoKaytuEnginePkgInventoryApiResourceSortItemTypeFieldPropEnum = append(githubComKaytuIoKaytuEnginePkgInventoryApiResourceSortItemTypeFieldPropEnum, v)
	}
}

// prop value enum
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceSortItem) validateFieldEnum(path, location string, value *struct {
	GithubComKaytuIoKaytuEnginePkgInventoryAPISortFieldType
}) error {
	if err := validate.EnumCase(path, location, value, githubComKaytuIoKaytuEnginePkgInventoryApiResourceSortItemTypeFieldPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceSortItem) validateField(formats strfmt.Registry) error {
	if swag.IsZero(m.Field) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this github com kaytu io kaytu engine pkg inventory api resource sort item based on the context it is used
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceSortItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceSortItem) contextValidateDirection(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceSortItem) contextValidateField(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceSortItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceSortItem) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceSortItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
