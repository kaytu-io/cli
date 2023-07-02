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

// GithubComKaytuIoKaytuEnginePkgDescribeAPIGetStackFindings github com kaytu io kaytu engine pkg describe api get stack findings
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_describe_api.GetStackFindings
type GithubComKaytuIoKaytuEnginePkgDescribeAPIGetStackFindings struct {

	// Benchmark IDs to filter
	// Example: ["azure_cis_v140"]
	BenchmarkIds []string `json:"benchmarkIds"`

	// Pages count to retrieve
	// Required: true
	Page GithubComKaytuIoKaytuEnginePkgComplianceAPIPage `json:"page"`

	// Sorts to apply
	Sorts []*GithubComKaytuIoKaytuEnginePkgComplianceAPIFindingSortItem `json:"sorts"`
}

// Validate validates this github com kaytu io kaytu engine pkg describe api get stack findings
func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIGetStackFindings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSorts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIGetStackFindings) validatePage(formats strfmt.Registry) error {

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIGetStackFindings) validateSorts(formats strfmt.Registry) error {
	if swag.IsZero(m.Sorts) { // not required
		return nil
	}

	for i := 0; i < len(m.Sorts); i++ {
		if swag.IsZero(m.Sorts[i]) { // not required
			continue
		}

		if m.Sorts[i] != nil {
			if err := m.Sorts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sorts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sorts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this github com kaytu io kaytu engine pkg describe api get stack findings based on the context it is used
func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIGetStackFindings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSorts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIGetStackFindings) contextValidatePage(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIGetStackFindings) contextValidateSorts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Sorts); i++ {

		if m.Sorts[i] != nil {

			if swag.IsZero(m.Sorts[i]) { // not required
				return nil
			}

			if err := m.Sorts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sorts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sorts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIGetStackFindings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIGetStackFindings) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgDescribeAPIGetStackFindings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}