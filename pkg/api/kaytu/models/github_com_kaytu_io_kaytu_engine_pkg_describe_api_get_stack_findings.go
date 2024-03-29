// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GithubComKaytuIoKaytuEnginePkgDescribeAPIGetStackFindings github com kaytu io kaytu engine pkg describe api get stack findings
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_describe_api.GetStackFindings
type GithubComKaytuIoKaytuEnginePkgDescribeAPIGetStackFindings struct {

	// Benchmark IDs to filter
	// Example: ["azure_cis_v140"]
	BenchmarkIds []string `json:"benchmarkIds"`
}

// Validate validates this github com kaytu io kaytu engine pkg describe api get stack findings
func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIGetStackFindings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this github com kaytu io kaytu engine pkg describe api get stack findings based on context it is used
func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIGetStackFindings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIGetStackFindings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIGetStackFindings) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgDescribeAPIGetStackFindings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
