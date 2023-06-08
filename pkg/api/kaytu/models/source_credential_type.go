// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// SourceCredentialType source credential type
//
// swagger:model source.CredentialType
type SourceCredentialType string

func NewSourceCredentialType(value SourceCredentialType) *SourceCredentialType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated SourceCredentialType.
func (m SourceCredentialType) Pointer() *SourceCredentialType {
	return &m
}

const (

	// SourceCredentialTypeAutoDashGenerated captures enum value "auto-generated"
	SourceCredentialTypeAutoDashGenerated SourceCredentialType = "auto-generated"

	// SourceCredentialTypeManual captures enum value "manual"
	SourceCredentialTypeManual SourceCredentialType = "manual"
)

// for schema
var sourceCredentialTypeEnum []interface{}

func init() {
	var res []SourceCredentialType
	if err := json.Unmarshal([]byte(`["auto-generated","manual"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sourceCredentialTypeEnum = append(sourceCredentialTypeEnum, v)
	}
}

func (m SourceCredentialType) validateSourceCredentialTypeEnum(path, location string, value SourceCredentialType) error {
	if err := validate.EnumCase(path, location, value, sourceCredentialTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this source credential type
func (m SourceCredentialType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSourceCredentialTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this source credential type based on context it is used
func (m SourceCredentialType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
