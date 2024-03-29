// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceTypeTrendDatapoint github com kaytu io kaytu engine pkg inventory api resource type trend datapoint
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_inventory_api.ResourceTypeTrendDatapoint
type GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceTypeTrendDatapoint struct {

	// count
	// Example: 100
	// Minimum: 0
	Count *int64 `json:"count,omitempty"`

	// date
	// Format: date-time
	Date strfmt.DateTime `json:"date,omitempty"`

	// total connection count
	TotalConnectionCount int64 `json:"totalConnectionCount,omitempty"`

	// total successful described connection count
	TotalSuccessfulDescribedConnectionCount int64 `json:"totalSuccessfulDescribedConnectionCount,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg inventory api resource type trend datapoint
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceTypeTrendDatapoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceTypeTrendDatapoint) validateCount(formats strfmt.Registry) error {
	if swag.IsZero(m.Count) { // not required
		return nil
	}

	if err := validate.MinimumInt("count", "body", *m.Count, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceTypeTrendDatapoint) validateDate(formats strfmt.Registry) error {
	if swag.IsZero(m.Date) { // not required
		return nil
	}

	if err := validate.FormatOf("date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

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
