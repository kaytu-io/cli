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

// GithubComKaytuIoKaytuEnginePkgInventoryAPICostTrendDatapoint github com kaytu io kaytu engine pkg inventory api cost trend datapoint
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_inventory_api.CostTrendDatapoint
type GithubComKaytuIoKaytuEnginePkgInventoryAPICostTrendDatapoint struct {

	// count
	// Minimum: 0
	Count *float64 `json:"count,omitempty"`

	// date
	// Format: date-time
	Date strfmt.DateTime `json:"date,omitempty"`

	// total connection count
	TotalConnectionCount int64 `json:"totalConnectionCount,omitempty"`

	// total successful described connection count
	TotalSuccessfulDescribedConnectionCount int64 `json:"totalSuccessfulDescribedConnectionCount,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg inventory api cost trend datapoint
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPICostTrendDatapoint) Validate(formats strfmt.Registry) error {
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

func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPICostTrendDatapoint) validateCount(formats strfmt.Registry) error {
	if swag.IsZero(m.Count) { // not required
		return nil
	}

	if err := validate.Minimum("count", "body", *m.Count, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPICostTrendDatapoint) validateDate(formats strfmt.Registry) error {
	if swag.IsZero(m.Date) { // not required
		return nil
	}

	if err := validate.FormatOf("date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this github com kaytu io kaytu engine pkg inventory api cost trend datapoint based on context it is used
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPICostTrendDatapoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPICostTrendDatapoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPICostTrendDatapoint) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgInventoryAPICostTrendDatapoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
