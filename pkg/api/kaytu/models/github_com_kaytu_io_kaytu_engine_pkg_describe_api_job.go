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

// GithubComKaytuIoKaytuEnginePkgDescribeAPIJob github com kaytu io kaytu engine pkg describe api job
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_describe_api.Job
type GithubComKaytuIoKaytuEnginePkgDescribeAPIJob struct {

	// connection ID
	ConnectionID string `json:"connectionID,omitempty"`

	// connection provider ID
	ConnectionProviderID string `json:"connectionProviderID,omitempty"`

	// connection provider name
	ConnectionProviderName string `json:"connectionProviderName,omitempty"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// status
	Status GithubComKaytuIoKaytuEnginePkgDescribeAPIJobStatus `json:"status,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type GithubComKaytuIoKaytuEnginePkgDescribeAPIJobType `json:"type,omitempty"`

	// updated at
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg describe api job
func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIJob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIJob) validateStatus(formats strfmt.Registry) error {
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

func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIJob) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// ContextValidate validate this github com kaytu io kaytu engine pkg describe api job based on the context it is used
func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIJob) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIJob) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIJob) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIJob) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgDescribeAPIJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
