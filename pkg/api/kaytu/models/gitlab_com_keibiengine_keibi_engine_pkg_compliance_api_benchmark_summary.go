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

// GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmarkSummary gitlab com keibiengine keibi engine pkg compliance api benchmark summary
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_compliance_api.BenchmarkSummary
type GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmarkSummary struct {

	// checks
	Checks *TypesSeverityResult `json:"checks,omitempty"`

	// compliancy trend
	CompliancyTrend []*GitlabComKeibiengineKeibiEnginePkgComplianceAPIDatapoint `json:"compliancyTrend"`

	// connectors
	Connectors []SourceType `json:"connectors"`

	// coverage
	Coverage float64 `json:"coverage,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// failed resources
	FailedResources int64 `json:"failedResources,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// passed resources
	PassedResources int64 `json:"passedResources,omitempty"`

	// result
	Result *TypesComplianceResultSummary `json:"result,omitempty"`

	// tags
	Tags map[string][]string `json:"tags,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg compliance api benchmark summary
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmarkSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChecks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompliancyTrend(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmarkSummary) validateChecks(formats strfmt.Registry) error {
	if swag.IsZero(m.Checks) { // not required
		return nil
	}

	if m.Checks != nil {
		if err := m.Checks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("checks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("checks")
			}
			return err
		}
	}

	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmarkSummary) validateCompliancyTrend(formats strfmt.Registry) error {
	if swag.IsZero(m.CompliancyTrend) { // not required
		return nil
	}

	for i := 0; i < len(m.CompliancyTrend); i++ {
		if swag.IsZero(m.CompliancyTrend[i]) { // not required
			continue
		}

		if m.CompliancyTrend[i] != nil {
			if err := m.CompliancyTrend[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("compliancyTrend" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("compliancyTrend" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmarkSummary) validateConnectors(formats strfmt.Registry) error {
	if swag.IsZero(m.Connectors) { // not required
		return nil
	}

	for i := 0; i < len(m.Connectors); i++ {

		if err := m.Connectors[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectors" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectors" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmarkSummary) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(m.Result) { // not required
		return nil
	}

	if m.Result != nil {
		if err := m.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this gitlab com keibiengine keibi engine pkg compliance api benchmark summary based on the context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmarkSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChecks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCompliancyTrend(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmarkSummary) contextValidateChecks(ctx context.Context, formats strfmt.Registry) error {

	if m.Checks != nil {

		if swag.IsZero(m.Checks) { // not required
			return nil
		}

		if err := m.Checks.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("checks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("checks")
			}
			return err
		}
	}

	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmarkSummary) contextValidateCompliancyTrend(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CompliancyTrend); i++ {

		if m.CompliancyTrend[i] != nil {

			if swag.IsZero(m.CompliancyTrend[i]) { // not required
				return nil
			}

			if err := m.CompliancyTrend[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("compliancyTrend" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("compliancyTrend" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmarkSummary) contextValidateConnectors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Connectors); i++ {

		if swag.IsZero(m.Connectors[i]) { // not required
			return nil
		}

		if err := m.Connectors[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectors" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectors" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmarkSummary) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if m.Result != nil {

		if swag.IsZero(m.Result) { // not required
			return nil
		}

		if err := m.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmarkSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmarkSummary) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmarkSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
