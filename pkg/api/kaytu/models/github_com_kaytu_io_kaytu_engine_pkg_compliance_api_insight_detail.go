// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GithubComKaytuIoKaytuEnginePkgComplianceAPIInsightDetail github com kaytu io kaytu engine pkg compliance api insight detail
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_compliance_api.InsightDetail
type GithubComKaytuIoKaytuEnginePkgComplianceAPIInsightDetail struct {

	// headers
	Headers []string `json:"headers"`

	// rows
	Rows [][]interface{} `json:"rows"`
}

// Validate validates this github com kaytu io kaytu engine pkg compliance api insight detail
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIInsightDetail) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this github com kaytu io kaytu engine pkg compliance api insight detail based on context it is used
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIInsightDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIInsightDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIInsightDetail) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgComplianceAPIInsightDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
