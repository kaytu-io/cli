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

// DescribeDescribeResourceJob describe describe resource job
//
// swagger:model describe.DescribeResourceJob
type DescribeDescribeResourceJob struct {

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// deleted at
	DeletedAt *GormDeletedAt `json:"deletedAt,omitempty"`

	// described resource count
	DescribedResourceCount int64 `json:"describedResourceCount,omitempty"`

	// Should be NULLSTRING
	ErrorCode string `json:"errorCode,omitempty"`

	// Should be NULLSTRING
	FailureMessage string `json:"failureMessage,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// parent job ID
	ParentJobID int64 `json:"parentJobID,omitempty"`

	// resource type
	ResourceType string `json:"resourceType,omitempty"`

	// retry count
	RetryCount int64 `json:"retryCount,omitempty"`

	// status
	Status GithubComKaytuIoKaytuEnginePkgDescribeAPIDescribeResourceJobStatus `json:"status,omitempty"`

	// updated at
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this describe describe resource job
func (m *DescribeDescribeResourceJob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeletedAt(formats); err != nil {
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

func (m *DescribeDescribeResourceJob) validateDeletedAt(formats strfmt.Registry) error {
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

func (m *DescribeDescribeResourceJob) validateStatus(formats strfmt.Registry) error {
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

// ContextValidate validate this describe describe resource job based on the context it is used
func (m *DescribeDescribeResourceJob) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeletedAt(ctx, formats); err != nil {
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

func (m *DescribeDescribeResourceJob) contextValidateDeletedAt(ctx context.Context, formats strfmt.Registry) error {

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

func (m *DescribeDescribeResourceJob) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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
func (m *DescribeDescribeResourceJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeDescribeResourceJob) UnmarshalBinary(b []byte) error {
	var res DescribeDescribeResourceJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
