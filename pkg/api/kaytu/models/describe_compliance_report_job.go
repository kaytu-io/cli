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

// DescribeComplianceReportJob describe compliance report job
//
// swagger:model describe.ComplianceReportJob
type DescribeComplianceReportJob struct {

	// Not the primary key but should be a unique identifier
	BenchmarkID string `json:"benchmarkID,omitempty"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// deleted at
	DeletedAt *GormDeletedAt `json:"deletedAt,omitempty"`

	// Should be NULLSTRING
	FailureMessage string `json:"failureMessage,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// report created at
	ReportCreatedAt int64 `json:"reportCreatedAt,omitempty"`

	// schedule job ID
	ScheduleJobID int64 `json:"scheduleJobID,omitempty"`

	// Not the primary key but should be a unique identifier
	SourceID string `json:"sourceID,omitempty"`

	// source type
	SourceType SourceType `json:"sourceType,omitempty"`

	// status
	Status GitlabComKeibiengineKeibiEnginePkgComplianceAPIComplianceReportJobStatus `json:"status,omitempty"`

	// updated at
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this describe compliance report job
func (m *DescribeComplianceReportJob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeletedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceType(formats); err != nil {
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

func (m *DescribeComplianceReportJob) validateDeletedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.DeletedAt) { // not required
		return nil
	}

	if m.DeletedAt != nil {
		if err := m.DeletedAt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deletedAt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deletedAt")
			}
			return err
		}
	}

	return nil
}

func (m *DescribeComplianceReportJob) validateSourceType(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceType) { // not required
		return nil
	}

	if err := m.SourceType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("sourceType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("sourceType")
		}
		return err
	}

	return nil
}

func (m *DescribeComplianceReportJob) validateStatus(formats strfmt.Registry) error {
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

// ContextValidate validate this describe compliance report job based on the context it is used
func (m *DescribeComplianceReportJob) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeletedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceType(ctx, formats); err != nil {
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

func (m *DescribeComplianceReportJob) contextValidateDeletedAt(ctx context.Context, formats strfmt.Registry) error {

	if m.DeletedAt != nil {

		if swag.IsZero(m.DeletedAt) { // not required
			return nil
		}

		if err := m.DeletedAt.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deletedAt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deletedAt")
			}
			return err
		}
	}

	return nil
}

func (m *DescribeComplianceReportJob) contextValidateSourceType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.SourceType) { // not required
		return nil
	}

	if err := m.SourceType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("sourceType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("sourceType")
		}
		return err
	}

	return nil
}

func (m *DescribeComplianceReportJob) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

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
func (m *DescribeComplianceReportJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeComplianceReportJob) UnmarshalBinary(b []byte) error {
	var res DescribeComplianceReportJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
