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

// GithubComKaytuIoKaytuEnginePkgMetadataModelsConfigMetadata github com kaytu io kaytu engine pkg metadata models config metadata
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_metadata_models.ConfigMetadata
type GithubComKaytuIoKaytuEnginePkgMetadataModelsConfigMetadata struct {

	// key
	Key GithubComKaytuIoKaytuEnginePkgMetadataModelsMetadataKey `json:"key,omitempty"`

	// type
	Type GithubComKaytuIoKaytuEnginePkgMetadataModelsConfigMetadataType `json:"type,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg metadata models config metadata
func (m *GithubComKaytuIoKaytuEnginePkgMetadataModelsConfigMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgMetadataModelsConfigMetadata) validateKey(formats strfmt.Registry) error {
	if swag.IsZero(m.Key) { // not required
		return nil
	}

	if err := m.Key.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("key")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("key")
		}
		return err
	}

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgMetadataModelsConfigMetadata) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// ContextValidate validate this github com kaytu io kaytu engine pkg metadata models config metadata based on the context it is used
func (m *GithubComKaytuIoKaytuEnginePkgMetadataModelsConfigMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgMetadataModelsConfigMetadata) contextValidateKey(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Key) { // not required
		return nil
	}

	if err := m.Key.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("key")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("key")
		}
		return err
	}

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgMetadataModelsConfigMetadata) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgMetadataModelsConfigMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgMetadataModelsConfigMetadata) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgMetadataModelsConfigMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
