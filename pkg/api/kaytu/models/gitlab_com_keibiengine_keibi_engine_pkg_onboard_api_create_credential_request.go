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

// GitlabComKeibiengineKeibiEnginePkgOnboardAPICreateCredentialRequest gitlab com keibiengine keibi engine pkg onboard api create credential request
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_onboard_api.CreateCredentialRequest
type GitlabComKeibiengineKeibiEnginePkgOnboardAPICreateCredentialRequest struct {

	// config
	Config interface{} `json:"config,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// source type
	SourceType SourceType `json:"source_type,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg onboard api create credential request
func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPICreateCredentialRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSourceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPICreateCredentialRequest) validateSourceType(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceType) { // not required
		return nil
	}

	if err := m.SourceType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("source_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("source_type")
		}
		return err
	}

	return nil
}

// ContextValidate validate this gitlab com keibiengine keibi engine pkg onboard api create credential request based on the context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPICreateCredentialRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSourceType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPICreateCredentialRequest) contextValidateSourceType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.SourceType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("source_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("source_type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPICreateCredentialRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPICreateCredentialRequest) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgOnboardAPICreateCredentialRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
