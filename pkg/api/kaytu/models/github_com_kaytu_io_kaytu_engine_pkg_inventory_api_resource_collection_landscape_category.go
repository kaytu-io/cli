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

// GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceCollectionLandscapeCategory github com kaytu io kaytu engine pkg inventory api resource collection landscape category
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_inventory_api.ResourceCollectionLandscapeCategory
type GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceCollectionLandscapeCategory struct {

	// description
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// subcategories
	Subcategories []*GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceCollectionLandscapeSubcategory `json:"subcategories"`
}

// Validate validates this github com kaytu io kaytu engine pkg inventory api resource collection landscape category
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceCollectionLandscapeCategory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubcategories(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceCollectionLandscapeCategory) validateSubcategories(formats strfmt.Registry) error {
	if swag.IsZero(m.Subcategories) { // not required
		return nil
	}

	for i := 0; i < len(m.Subcategories); i++ {
		if swag.IsZero(m.Subcategories[i]) { // not required
			continue
		}

		if m.Subcategories[i] != nil {
			if err := m.Subcategories[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subcategories" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subcategories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this github com kaytu io kaytu engine pkg inventory api resource collection landscape category based on the context it is used
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceCollectionLandscapeCategory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSubcategories(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceCollectionLandscapeCategory) contextValidateSubcategories(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Subcategories); i++ {

		if m.Subcategories[i] != nil {

			if swag.IsZero(m.Subcategories[i]) { // not required
				return nil
			}

			if err := m.Subcategories[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subcategories" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subcategories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceCollectionLandscapeCategory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceCollectionLandscapeCategory) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceCollectionLandscapeCategory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
