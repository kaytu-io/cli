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

// GithubComKaytuIoKaytuEnginePkgComplianceAPIFindingFilters github com kaytu io kaytu engine pkg compliance api finding filters
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_compliance_api.FindingFilters
type GithubComKaytuIoKaytuEnginePkgComplianceAPIFindingFilters struct {

	// active only
	ActiveOnly bool `json:"activeOnly,omitempty"`

	// Benchmark ID
	// Example: ["azure_cis_v140"]
	BenchmarkID []string `json:"benchmarkID"`

	// Connection ID
	// Example: ["8e0f8e7a-1b1c-4e6f-b7e4-9c6af9d2b1c8"]
	ConnectionID []string `json:"connectionID"`

	// Clout Provider
	// Example: ["Azure"]
	Connector []SourceType `json:"connector"`

	// Policy ID
	// Example: ["azure_cis_v140_7_5"]
	PolicyID []string `json:"policyID"`

	// Resource Collection ID
	// Example: ["example-rc"]
	ResourceCollection []string `json:"resourceCollection"`

	// Resource unique identifier
	// Example: ["/subscriptions/123/resourceGroups/rg-1/providers/Microsoft.Compute/virtualMachines/vm-1"]
	ResourceID []string `json:"resourceID"`

	// Resource type
	// Example: ["/subscriptions/123/resourceGroups/rg-1/providers/Microsoft.Compute/virtualMachines"]
	ResourceTypeID []string `json:"resourceTypeID"`

	// Severity
	// Example: ["low"]
	Severity []string `json:"severity"`

	// Compliance result status
	// Example: ["alarm"]
	Status []TypesComplianceResult `json:"status"`
}

// Validate validates this github com kaytu io kaytu engine pkg compliance api finding filters
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIFindingFilters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIFindingFilters) validateConnector(formats strfmt.Registry) error {
	if swag.IsZero(m.Connector) { // not required
		return nil
	}

	for i := 0; i < len(m.Connector); i++ {

		if err := m.Connector[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connector" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connector" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIFindingFilters) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	for i := 0; i < len(m.Status); i++ {

		if err := m.Status[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this github com kaytu io kaytu engine pkg compliance api finding filters based on the context it is used
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIFindingFilters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIFindingFilters) contextValidateConnector(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Connector); i++ {

		if swag.IsZero(m.Connector[i]) { // not required
			return nil
		}

		if err := m.Connector[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connector" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connector" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIFindingFilters) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Status); i++ {

		if swag.IsZero(m.Status[i]) { // not required
			return nil
		}

		if err := m.Status[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIFindingFilters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIFindingFilters) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgComplianceAPIFindingFilters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
