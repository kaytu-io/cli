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

// GithubComKaytuIoKaytuEnginePkgComplianceAPIGetFindingsRequest github com kaytu io kaytu engine pkg compliance api get findings request
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_compliance_api.GetFindingsRequest
type GithubComKaytuIoKaytuEnginePkgComplianceAPIGetFindingsRequest struct {

	// after sort key
	AfterSortKey []interface{} `json:"afterSortKey"`

	// filters
	Filters *GithubComKaytuIoKaytuEnginePkgComplianceAPIFindingFilters `json:"filters,omitempty"`

	// limit
	// Example: 100
	Limit int64 `json:"limit,omitempty"`

	// sort
	Sort map[string]string `json:"sort,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg compliance api get findings request
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIGetFindingsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIGetFindingsRequest) validateFilters(formats strfmt.Registry) error {
	if swag.IsZero(m.Filters) { // not required
		return nil
	}

	if m.Filters != nil {
		if err := m.Filters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filters")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this github com kaytu io kaytu engine pkg compliance api get findings request based on the context it is used
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIGetFindingsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIGetFindingsRequest) contextValidateFilters(ctx context.Context, formats strfmt.Registry) error {

	if m.Filters != nil {

		if swag.IsZero(m.Filters) { // not required
			return nil
		}

		if err := m.Filters.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filters")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIGetFindingsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIGetFindingsRequest) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgComplianceAPIGetFindingsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
