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

// GithubComKaytuIoKaytuEnginePkgComplianceAPIAssignedBenchmark github com kaytu io kaytu engine pkg compliance api assigned benchmark
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_compliance_api.AssignedBenchmark
type GithubComKaytuIoKaytuEnginePkgComplianceAPIAssignedBenchmark struct {

	// benchmark Id
	BenchmarkID *GithubComKaytuIoKaytuEnginePkgComplianceAPIBenchmark `json:"benchmarkId,omitempty"`

	// Status
	// Example: true
	Status bool `json:"status,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg compliance api assigned benchmark
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIAssignedBenchmark) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBenchmarkID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIAssignedBenchmark) validateBenchmarkID(formats strfmt.Registry) error {
	if swag.IsZero(m.BenchmarkID) { // not required
		return nil
	}

	if m.BenchmarkID != nil {
		if err := m.BenchmarkID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("benchmarkId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("benchmarkId")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this github com kaytu io kaytu engine pkg compliance api assigned benchmark based on the context it is used
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIAssignedBenchmark) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBenchmarkID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIAssignedBenchmark) contextValidateBenchmarkID(ctx context.Context, formats strfmt.Registry) error {

	if m.BenchmarkID != nil {

		if swag.IsZero(m.BenchmarkID) { // not required
			return nil
		}

		if err := m.BenchmarkID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("benchmarkId")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("benchmarkId")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIAssignedBenchmark) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIAssignedBenchmark) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgComplianceAPIAssignedBenchmark
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}