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

// GithubComKaytuIoKaytuEnginePkgInsightAPIInsightJobStatus github com kaytu io kaytu engine pkg insight api insight job status
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_insight_api.InsightJobStatus
type GithubComKaytuIoKaytuEnginePkgInsightAPIInsightJobStatus string

func NewGithubComKaytuIoKaytuEnginePkgInsightAPIInsightJobStatus(value GithubComKaytuIoKaytuEnginePkgInsightAPIInsightJobStatus) *GithubComKaytuIoKaytuEnginePkgInsightAPIInsightJobStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated GithubComKaytuIoKaytuEnginePkgInsightAPIInsightJobStatus.
func (m GithubComKaytuIoKaytuEnginePkgInsightAPIInsightJobStatus) Pointer() *GithubComKaytuIoKaytuEnginePkgInsightAPIInsightJobStatus {
	return &m
}

const (

	// GithubComKaytuIoKaytuEnginePkgInsightAPIInsightJobStatusINPROGRESS captures enum value "IN_PROGRESS"
	GithubComKaytuIoKaytuEnginePkgInsightAPIInsightJobStatusINPROGRESS GithubComKaytuIoKaytuEnginePkgInsightAPIInsightJobStatus = "IN_PROGRESS"

	// GithubComKaytuIoKaytuEnginePkgInsightAPIInsightJobStatusFAILED captures enum value "FAILED"
	GithubComKaytuIoKaytuEnginePkgInsightAPIInsightJobStatusFAILED GithubComKaytuIoKaytuEnginePkgInsightAPIInsightJobStatus = "FAILED"

	// GithubComKaytuIoKaytuEnginePkgInsightAPIInsightJobStatusSUCCEEDED captures enum value "SUCCEEDED"
	GithubComKaytuIoKaytuEnginePkgInsightAPIInsightJobStatusSUCCEEDED GithubComKaytuIoKaytuEnginePkgInsightAPIInsightJobStatus = "SUCCEEDED"
)

// for schema
var githubComKaytuIoKaytuEnginePkgInsightApiInsightJobStatusEnum []interface{}

func init() {
	var res []GithubComKaytuIoKaytuEnginePkgInsightAPIInsightJobStatus
	if err := json.Unmarshal([]byte(`["IN_PROGRESS","FAILED","SUCCEEDED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		githubComKaytuIoKaytuEnginePkgInsightApiInsightJobStatusEnum = append(githubComKaytuIoKaytuEnginePkgInsightApiInsightJobStatusEnum, v)
	}
}

func (m GithubComKaytuIoKaytuEnginePkgInsightAPIInsightJobStatus) validateGithubComKaytuIoKaytuEnginePkgInsightAPIInsightJobStatusEnum(path, location string, value GithubComKaytuIoKaytuEnginePkgInsightAPIInsightJobStatus) error {
	if err := validate.EnumCase(path, location, value, githubComKaytuIoKaytuEnginePkgInsightApiInsightJobStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this github com kaytu io kaytu engine pkg insight api insight job status
func (m GithubComKaytuIoKaytuEnginePkgInsightAPIInsightJobStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateGithubComKaytuIoKaytuEnginePkgInsightAPIInsightJobStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this github com kaytu io kaytu engine pkg insight api insight job status based on context it is used
func (m GithubComKaytuIoKaytuEnginePkgInsightAPIInsightJobStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}