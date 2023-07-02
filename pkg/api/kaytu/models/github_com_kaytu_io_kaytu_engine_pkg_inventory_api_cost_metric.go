// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GithubComKaytuIoKaytuEnginePkgInventoryAPICostMetric github com kaytu io kaytu engine pkg inventory api cost metric
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_inventory_api.CostMetric
type GithubComKaytuIoKaytuEnginePkgInventoryAPICostMetric struct {

	// connector
	// Example: azure
	Connector SourceType `json:"connector,omitempty"`

	// cost dimension name
	CostDimensionName string `json:"cost_dimension_name,omitempty"`

	// daily cost at end time
	DailyCostAtEndTime float64 `json:"daily_cost_at_end_time,omitempty"`

	// daily cost at start time
	DailyCostAtStartTime float64 `json:"daily_cost_at_start_time,omitempty"`

	// total cost
	TotalCost float64 `json:"total_cost,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg inventory api cost metric
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPICostMetric) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPICostMetric) validateConnector(formats strfmt.Registry) error {
	if swag.IsZero(m.Connector) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this github com kaytu io kaytu engine pkg inventory api cost metric based on the context it is used
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPICostMetric) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPICostMetric) contextValidateConnector(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPICostMetric) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPICostMetric) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgInventoryAPICostMetric
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}