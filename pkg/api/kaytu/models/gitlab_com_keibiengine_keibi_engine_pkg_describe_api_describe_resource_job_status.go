// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatus gitlab com keibiengine keibi engine pkg describe api describe resource job status
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_describe_api.DescribeResourceJobStatus
type GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatus string

func NewGitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatus(value GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatus) *GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatus.
func (m GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatus) Pointer() *GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatus {
	return &m
}

const (

	// GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatusCREATED captures enum value "CREATED"
	GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatusCREATED GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatus = "CREATED"

	// GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatusQUEUED captures enum value "QUEUED"
	GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatusQUEUED GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatus = "QUEUED"

	// GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatusINPROGRESS captures enum value "IN_PROGRESS"
	GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatusINPROGRESS GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatus = "IN_PROGRESS"

	// GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatusTIMEOUT captures enum value "TIMEOUT"
	GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatusTIMEOUT GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatus = "TIMEOUT"

	// GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatusFAILED captures enum value "FAILED"
	GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatusFAILED GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatus = "FAILED"

	// GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatusSUCCEEDED captures enum value "SUCCEEDED"
	GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatusSUCCEEDED GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatus = "SUCCEEDED"
)

// for schema
var gitlabComKeibiengineKeibiEnginePkgDescribeApiDescribeResourceJobStatusEnum []interface{}

func init() {
	var res []GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatus
	if err := json.Unmarshal([]byte(`["CREATED","QUEUED","IN_PROGRESS","TIMEOUT","FAILED","SUCCEEDED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gitlabComKeibiengineKeibiEnginePkgDescribeApiDescribeResourceJobStatusEnum = append(gitlabComKeibiengineKeibiEnginePkgDescribeApiDescribeResourceJobStatusEnum, v)
	}
}

func (m GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatus) validateGitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatusEnum(path, location string, value GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatus) error {
	if err := validate.EnumCase(path, location, value, gitlabComKeibiengineKeibiEnginePkgDescribeApiDescribeResourceJobStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this gitlab com keibiengine keibi engine pkg describe api describe resource job status
func (m GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateGitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this gitlab com keibiengine keibi engine pkg describe api describe resource job status based on context it is used
func (m GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeResourceJobStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}