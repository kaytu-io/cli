// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GithubComKaytuIoKaytuEnginePkgComplianceAPIFindingFilterWithMetadata github com kaytu io kaytu engine pkg compliance api finding filter with metadata
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_compliance_api.FindingFilterWithMetadata
type GithubComKaytuIoKaytuEnginePkgComplianceAPIFindingFilterWithMetadata struct {

	// Display Name
	// Example: displayName
	DisplayName string `json:"displayName,omitempty"`

	// Key
	// Example: key
	Key string `json:"key,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg compliance api finding filter with metadata
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIFindingFilterWithMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this github com kaytu io kaytu engine pkg compliance api finding filter with metadata based on context it is used
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIFindingFilterWithMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIFindingFilterWithMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIFindingFilterWithMetadata) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgComplianceAPIFindingFilterWithMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
