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

// GithubComKaytuIoKaytuEnginePkgComplianceAPIBenchmark github com kaytu io kaytu engine pkg compliance api benchmark
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_compliance_api.Benchmark
type GithubComKaytuIoKaytuEnginePkgComplianceAPIBenchmark struct {

	// Whether the benchmark is auto assigned or not
	// Example: true
	AutoAssign bool `json:"autoAssign,omitempty"`

	// Whether the benchmark is baseline or not
	// Example: true
	Baseline bool `json:"baseline,omitempty"`

	// Benchmark category
	Category string `json:"category,omitempty"`

	// Benchmark children
	// Example: ["[azure_cis_v140_1"," azure_cis_v140_2]"]
	Children []string `json:"children"`

	// Benchmark connectors
	// Example: ["[azure]"]
	Connectors []SourceType `json:"connectors"`

	// Benchmark controls
	// Example: ["[azure_cis_v140_1_1"," azure_cis_v140_1_2]"]
	Controls []string `json:"controls"`

	// Benchmark creation date
	// Example: 2020-01-01T00:00:00Z
	CreatedAt string `json:"createdAt,omitempty"`

	// Benchmark description
	// Example: The CIS Microsoft Azure Foundations Security Benchmark provides prescriptive guidance for establishing a secure baseline configuration for Microsoft Azure.
	Description string `json:"description,omitempty"`

	// Benchmark document URI
	// Example: benchmarks/azure_cis_v140.md
	DocumentURI string `json:"documentURI,omitempty"`

	// Whether the benchmark is enabled or not
	// Example: true
	Enabled bool `json:"enabled,omitempty"`

	// Benchmark ID
	// Example: azure_cis_v140
	ID string `json:"id,omitempty"`

	// Benchmark logo URI
	LogoURI string `json:"logoURI,omitempty"`

	// Whether the benchmark is managed or not
	// Example: true
	Managed bool `json:"managed,omitempty"`

	// Benchmark tags
	Tags map[string][]string `json:"tags,omitempty"`

	// Benchmark title
	// Example: Azure CIS v1.4.0
	Title string `json:"title,omitempty"`

	// Benchmark last update date
	// Example: 2020-01-01T00:00:00Z
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg compliance api benchmark
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIBenchmark) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIBenchmark) validateConnectors(formats strfmt.Registry) error {
	if swag.IsZero(m.Connectors) { // not required
		return nil
	}

	for i := 0; i < len(m.Connectors); i++ {

		if err := m.Connectors[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectors" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectors" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this github com kaytu io kaytu engine pkg compliance api benchmark based on the context it is used
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIBenchmark) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnectors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIBenchmark) contextValidateConnectors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Connectors); i++ {

		if swag.IsZero(m.Connectors[i]) { // not required
			return nil
		}

		if err := m.Connectors[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectors" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectors" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIBenchmark) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIBenchmark) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgComplianceAPIBenchmark
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
