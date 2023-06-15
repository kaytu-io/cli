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

// DescribeDescribeSourceJob describe describe source job
//
// swagger:model describe.DescribeSourceJob
type DescribeDescribeSourceJob struct {

	// account ID
	AccountID string `json:"accountID,omitempty"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// deleted at
	DeletedAt *GormDeletedAt `json:"deletedAt,omitempty"`

	// describe resource jobs
	DescribeResourceJobs []*DescribeDescribeResourceJob `json:"describeResourceJobs"`

	// described at
	DescribedAt string `json:"describedAt,omitempty"`

	// full discovery
	FullDiscovery bool `json:"fullDiscovery,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// Not the primary key but should be a unique identifier
	SourceID string `json:"sourceID,omitempty"`

	// source type
	SourceType SourceType `json:"sourceType,omitempty"`

	// status
	Status GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeSourceJobStatus `json:"status,omitempty"`

	// trigger type
	TriggerType GitlabComKeibiengineKeibiEnginePkgDescribeEnumsDescribeTriggerType `json:"triggerType,omitempty"`

	// updated at
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this describe describe source job
func (m *DescribeDescribeSourceJob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeletedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescribeResourceJobs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTriggerType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DescribeDescribeSourceJob) validateDeletedAt(formats strfmt.Registry) error {
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

func (m *DescribeDescribeSourceJob) validateDescribeResourceJobs(formats strfmt.Registry) error {
	if swag.IsZero(m.DescribeResourceJobs) { // not required
		return nil
	}

	for i := 0; i < len(m.DescribeResourceJobs); i++ {
		if swag.IsZero(m.DescribeResourceJobs[i]) { // not required
			continue
		}

		if m.DescribeResourceJobs[i] != nil {
			if err := m.DescribeResourceJobs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("describeResourceJobs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("describeResourceJobs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DescribeDescribeSourceJob) validateSourceType(formats strfmt.Registry) error {
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

func (m *DescribeDescribeSourceJob) validateStatus(formats strfmt.Registry) error {
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

func (m *DescribeDescribeSourceJob) validateTriggerType(formats strfmt.Registry) error {
	if swag.IsZero(m.TriggerType) { // not required
		return nil
	}

	if err := m.TriggerType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("triggerType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("triggerType")
		}
		return err
	}

	return nil
}

// ContextValidate validate this describe describe source job based on the context it is used
func (m *DescribeDescribeSourceJob) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeletedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDescribeResourceJobs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSourceType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTriggerType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DescribeDescribeSourceJob) contextValidateDeletedAt(ctx context.Context, formats strfmt.Registry) error {

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

func (m *DescribeDescribeSourceJob) contextValidateDescribeResourceJobs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DescribeResourceJobs); i++ {

		if m.DescribeResourceJobs[i] != nil {

			if swag.IsZero(m.DescribeResourceJobs[i]) { // not required
				return nil
			}

			if err := m.DescribeResourceJobs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("describeResourceJobs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("describeResourceJobs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DescribeDescribeSourceJob) contextValidateSourceType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *DescribeDescribeSourceJob) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *DescribeDescribeSourceJob) contextValidateTriggerType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.TriggerType) { // not required
		return nil
	}

	if err := m.TriggerType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("triggerType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("triggerType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DescribeDescribeSourceJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeDescribeSourceJob) UnmarshalBinary(b []byte) error {
	var res DescribeDescribeSourceJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
