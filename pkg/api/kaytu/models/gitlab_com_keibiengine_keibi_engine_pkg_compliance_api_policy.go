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

// GitlabComKeibiengineKeibiEnginePkgComplianceAPIPolicy gitlab com keibiengine keibi engine pkg compliance api policy
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_compliance_api.Policy
type GitlabComKeibiengineKeibiEnginePkgComplianceAPIPolicy struct {

	// connector
	Connector SourceType `json:"connector,omitempty"`

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

	// managed
	Managed bool `json:"managed,omitempty"`

	// manual verification
	ManualVerification bool `json:"manualVerification,omitempty"`

	// query ID
	QueryID string `json:"queryID,omitempty"`

	// severity
	Severity string `json:"severity,omitempty"`

	// tags
	Tags map[string][]string `json:"tags,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// updated at
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg compliance api policy
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIPolicy) validateConnector(formats strfmt.Registry) error {
	if swag.IsZero(m.Connector) { // not required
		return nil
	}

	if err := m.Connector.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("connector")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("connector")
		}
		return err
	}

	return nil
}

// ContextValidate validate this gitlab com keibiengine keibi engine pkg compliance api policy based on the context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIPolicy) contextValidateConnector(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Connector.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("connector")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("connector")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIPolicy) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgComplianceAPIPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}