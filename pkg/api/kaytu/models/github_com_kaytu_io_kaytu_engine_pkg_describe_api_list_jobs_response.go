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

// GithubComKaytuIoKaytuEnginePkgDescribeAPIListJobsResponse github com kaytu io kaytu engine pkg describe api list jobs response
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_describe_api.ListJobsResponse
type GithubComKaytuIoKaytuEnginePkgDescribeAPIListJobsResponse struct {

	// jobs
	Jobs []*GithubComKaytuIoKaytuEnginePkgDescribeAPIJob `json:"jobs"`

	// summaries
	Summaries []*GithubComKaytuIoKaytuEnginePkgDescribeAPIJobSummary `json:"summaries"`
}

// Validate validates this github com kaytu io kaytu engine pkg describe api list jobs response
func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIListJobsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJobs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSummaries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIListJobsResponse) validateJobs(formats strfmt.Registry) error {
	if swag.IsZero(m.Jobs) { // not required
		return nil
	}

	for i := 0; i < len(m.Jobs); i++ {
		if swag.IsZero(m.Jobs[i]) { // not required
			continue
		}

		if m.Jobs[i] != nil {
			if err := m.Jobs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("jobs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("jobs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIListJobsResponse) validateSummaries(formats strfmt.Registry) error {
	if swag.IsZero(m.Summaries) { // not required
		return nil
	}

	for i := 0; i < len(m.Summaries); i++ {
		if swag.IsZero(m.Summaries[i]) { // not required
			continue
		}

		if m.Summaries[i] != nil {
			if err := m.Summaries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("summaries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("summaries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this github com kaytu io kaytu engine pkg describe api list jobs response based on the context it is used
func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIListJobsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateJobs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSummaries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIListJobsResponse) contextValidateJobs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Jobs); i++ {

		if m.Jobs[i] != nil {

			if swag.IsZero(m.Jobs[i]) { // not required
				return nil
			}

			if err := m.Jobs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("jobs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("jobs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIListJobsResponse) contextValidateSummaries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Summaries); i++ {

		if m.Summaries[i] != nil {

			if swag.IsZero(m.Summaries[i]) { // not required
				return nil
			}

			if err := m.Summaries[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("summaries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("summaries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIListJobsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIListJobsResponse) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgDescribeAPIListJobsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}