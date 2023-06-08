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

// GitlabComKeibiengineKeibiEnginePkgComplianceAPIComplianceReport gitlab com keibiengine keibi engine pkg compliance api compliance report
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_compliance_api.ComplianceReport
type GitlabComKeibiengineKeibiEnginePkgComplianceAPIComplianceReport struct {

	// failure message
	FailureMessage string `json:"failureMessage,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// report created at
	ReportCreatedAt int64 `json:"reportCreatedAt,omitempty"`

	// status
	Status GitlabComKeibiengineKeibiEnginePkgComplianceAPIComplianceReportJobStatus `json:"status,omitempty"`

	// updated at
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg compliance api compliance report
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIComplianceReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIComplianceReport) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

// ContextValidate validate this gitlab com keibiengine keibi engine pkg compliance api compliance report based on the context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIComplianceReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIComplianceReport) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Status.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIComplianceReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIComplianceReport) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgComplianceAPIComplianceReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
