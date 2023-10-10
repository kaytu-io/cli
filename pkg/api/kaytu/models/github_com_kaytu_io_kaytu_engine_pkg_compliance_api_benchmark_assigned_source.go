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

// GithubComKaytuIoKaytuEnginePkgComplianceAPIBenchmarkAssignedSource github com kaytu io kaytu engine pkg compliance api benchmark assigned source
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_compliance_api.BenchmarkAssignedSource
type GithubComKaytuIoKaytuEnginePkgComplianceAPIBenchmarkAssignedSource struct {

	// Connection ID
	// Example: 8e0f8e7a-1b1c-4e6f-b7e4-9c6af9d2b1c8
	ConnectionID string `json:"connectionID,omitempty"`

	// Clout Provider
	// Example: Azure
	Connector struct {
		SourceType
	} `json:"connector,omitempty"`

	// Provider Connection ID
	// Example: 1283192749
	ProviderConnectionID string `json:"providerConnectionID,omitempty"`

	// Provider Connection Name
	ProviderConnectionName string `json:"providerConnectionName,omitempty"`

	// Status
	// Example: true
	Status bool `json:"status,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg compliance api benchmark assigned source
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIBenchmarkAssignedSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIBenchmarkAssignedSource) validateConnector(formats strfmt.Registry) error {
	if swag.IsZero(m.Connector) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this github com kaytu io kaytu engine pkg compliance api benchmark assigned source based on the context it is used
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIBenchmarkAssignedSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIBenchmarkAssignedSource) contextValidateConnector(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIBenchmarkAssignedSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIBenchmarkAssignedSource) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgComplianceAPIBenchmarkAssignedSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
