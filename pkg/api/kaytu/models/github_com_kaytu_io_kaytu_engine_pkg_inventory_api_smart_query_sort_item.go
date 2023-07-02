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

// GithubComKaytuIoKaytuEnginePkgInventoryAPISmartQuerySortItem github com kaytu io kaytu engine pkg inventory api smart query sort item
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_inventory_api.SmartQuerySortItem
type GithubComKaytuIoKaytuEnginePkgInventoryAPISmartQuerySortItem struct {

	// direction
	// Enum: [asc desc]
	Direction GithubComKaytuIoKaytuEnginePkgInventoryAPIDirectionType `json:"direction,omitempty"`

	// fill this with column name
	Field string `json:"field,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg inventory api smart query sort item
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPISmartQuerySortItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var githubComKaytuIoKaytuEnginePkgInventoryApiSmartQuerySortItemTypeDirectionPropEnum []interface{}

func init() {
	var res []GithubComKaytuIoKaytuEnginePkgInventoryAPIDirectionType
	if err := json.Unmarshal([]byte(`["asc","desc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		githubComKaytuIoKaytuEnginePkgInventoryApiSmartQuerySortItemTypeDirectionPropEnum = append(githubComKaytuIoKaytuEnginePkgInventoryApiSmartQuerySortItemTypeDirectionPropEnum, v)
	}
}

// prop value enum
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPISmartQuerySortItem) validateDirectionEnum(path, location string, value *struct {
	GithubComKaytuIoKaytuEnginePkgInventoryAPIDirectionType
}) error {
	if err := validate.EnumCase(path, location, value, githubComKaytuIoKaytuEnginePkgInventoryApiSmartQuerySortItemTypeDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPISmartQuerySortItem) validateDirection(formats strfmt.Registry) error {
	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this github com kaytu io kaytu engine pkg inventory api smart query sort item based on the context it is used
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPISmartQuerySortItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDirection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPISmartQuerySortItem) contextValidateDirection(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPISmartQuerySortItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPISmartQuerySortItem) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgInventoryAPISmartQuerySortItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}