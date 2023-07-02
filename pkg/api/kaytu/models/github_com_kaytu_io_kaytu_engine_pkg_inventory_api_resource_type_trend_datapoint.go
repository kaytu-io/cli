// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceTypeTrendDatapoint github com kaytu io kaytu engine pkg inventory api resource type trend datapoint
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_inventory_api.ResourceTypeTrendDatapoint
type GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceTypeTrendDatapoint struct {

	// count
	Count int64 `json:"count,omitempty"`

	// date
	Date string `json:"date,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg inventory api resource type trend datapoint
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceTypeTrendDatapoint) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this github com kaytu io kaytu engine pkg inventory api resource type trend datapoint based on context it is used
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceTypeTrendDatapoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceTypeTrendDatapoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceTypeTrendDatapoint) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceTypeTrendDatapoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}