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

// GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmark gitlab com keibiengine keibi engine pkg compliance api benchmark
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_compliance_api.Benchmark
type GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmark struct {

	// auto assign
	AutoAssign bool `json:"autoAssign,omitempty"`

	// baseline
	Baseline bool `json:"baseline,omitempty"`

	// category
	Category string `json:"category,omitempty"`

	// children
	Children []string `json:"children"`

	// connectors
	Connectors []SourceType `json:"connectors"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// document URI
	DocumentURI string `json:"documentURI,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// logo URI
	LogoURI string `json:"logoURI,omitempty"`

	// managed
	Managed bool `json:"managed,omitempty"`

	// policies
	Policies []string `json:"policies"`

	// tags
	Tags map[string][]string `json:"tags,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// updated at
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg compliance api benchmark
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmark) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmark) validateConnectors(formats strfmt.Registry) error {
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

// ContextValidate validate this gitlab com keibiengine keibi engine pkg compliance api benchmark based on the context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmark) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnectors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmark) contextValidateConnectors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Connectors); i++ {

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
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmark) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmark) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmark
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
