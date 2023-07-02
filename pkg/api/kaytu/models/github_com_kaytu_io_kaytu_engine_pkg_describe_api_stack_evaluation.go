// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GithubComKaytuIoKaytuEnginePkgDescribeAPIStackEvaluation github com kaytu io kaytu engine pkg describe api stack evaluation
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_describe_api.StackEvaluation
type GithubComKaytuIoKaytuEnginePkgDescribeAPIStackEvaluation struct {

	// Evaluation creation date
	// Example: 2020-01-01T00:00:00Z
	CreatedAt string `json:"createdAt,omitempty"`

	// Benchmark ID or Insight ID
	// Example: azure_cis_v140
	EvaluatorID string `json:"evaluatorId,omitempty"`

	// Evaluation Job ID to find the job results
	// Example: 1
	JobID int64 `json:"jobId,omitempty"`

	// BENCHMARK or INSIGHT
	// Example: BENCHMARK
	Type string `json:"type,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg describe api stack evaluation
func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIStackEvaluation) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this github com kaytu io kaytu engine pkg describe api stack evaluation based on context it is used
func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIStackEvaluation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIStackEvaluation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIStackEvaluation) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgDescribeAPIStackEvaluation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}