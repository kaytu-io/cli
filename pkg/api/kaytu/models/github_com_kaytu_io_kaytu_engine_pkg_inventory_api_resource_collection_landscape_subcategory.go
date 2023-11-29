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
)

// GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceCollectionLandscapeSubcategory github com kaytu io kaytu engine pkg inventory api resource collection landscape subcategory
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_inventory_api.ResourceCollectionLandscapeSubcategory
type GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceCollectionLandscapeSubcategory struct {

	// description
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// items
	Items []*GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceCollectionLandscapeItem `json:"items"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg inventory api resource collection landscape subcategory
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceCollectionLandscapeSubcategory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceCollectionLandscapeSubcategory) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this github com kaytu io kaytu engine pkg inventory api resource collection landscape subcategory based on the context it is used
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceCollectionLandscapeSubcategory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceCollectionLandscapeSubcategory) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Items); i++ {

		if m.Items[i] != nil {

			if swag.IsZero(m.Items[i]) { // not required
				return nil
			}

			if err := m.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceCollectionLandscapeSubcategory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceCollectionLandscapeSubcategory) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceCollectionLandscapeSubcategory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}