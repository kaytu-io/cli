// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GithubComKaytuIoKaytuEnginePkgComplianceAPIInsightConnection github com kaytu io kaytu engine pkg compliance api insight connection
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_compliance_api.InsightConnection
type GithubComKaytuIoKaytuEnginePkgComplianceAPIInsightConnection struct {

	// connection id
	ConnectionID string `json:"connection_id,omitempty"`

	// original id
	OriginalID string `json:"original_id,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg compliance api insight connection
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIInsightConnection) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this github com kaytu io kaytu engine pkg compliance api insight connection based on context it is used
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIInsightConnection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIInsightConnection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIInsightConnection) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgComplianceAPIInsightConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}