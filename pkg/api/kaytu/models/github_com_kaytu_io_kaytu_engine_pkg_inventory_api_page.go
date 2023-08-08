// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GithubComKaytuIoKaytuEnginePkgInventoryAPIPage github com kaytu io kaytu engine pkg inventory api page
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_inventory_api.Page
type GithubComKaytuIoKaytuEnginePkgInventoryAPIPage struct {

	// no
	No int64 `json:"no,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg inventory api page
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIPage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this github com kaytu io kaytu engine pkg inventory api page based on context it is used
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIPage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIPage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIPage) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgInventoryAPIPage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
