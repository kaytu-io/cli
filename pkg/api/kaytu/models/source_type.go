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

// SourceType source type
//
// swagger:model source.Type
type SourceType string

func NewSourceType(value SourceType) *SourceType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated SourceType.
func (m SourceType) Pointer() *SourceType {
	return &m
}

const (

	// SourceTypeEmpty captures enum value ""
	SourceTypeEmpty SourceType = ""

	// SourceTypeAWS captures enum value "AWS"
	SourceTypeAWS SourceType = "AWS"

	// SourceTypeAzure captures enum value "Azure"
	SourceTypeAzure SourceType = "Azure"
)

// for schema
var sourceTypeEnum []interface{}

func init() {
	var res []SourceType
	if err := json.Unmarshal([]byte(`["","AWS","Azure"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sourceTypeEnum = append(sourceTypeEnum, v)
	}
}

func (m SourceType) validateSourceTypeEnum(path, location string, value SourceType) error {
	if err := validate.EnumCase(path, location, value, sourceTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this source type
func (m SourceType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSourceTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this source type based on context it is used
func (m SourceType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
