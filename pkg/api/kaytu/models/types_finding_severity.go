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

// TypesFindingSeverity types finding severity
//
// swagger:model types.FindingSeverity
type TypesFindingSeverity string

func NewTypesFindingSeverity(value TypesFindingSeverity) *TypesFindingSeverity {
	return &value
}

// Pointer returns a pointer to a freshly-allocated TypesFindingSeverity.
func (m TypesFindingSeverity) Pointer() *TypesFindingSeverity {
	return &m
}

const (

	// TypesFindingSeverityNone captures enum value "none"
	TypesFindingSeverityNone TypesFindingSeverity = "none"

	// TypesFindingSeverityPassed captures enum value "passed"
	TypesFindingSeverityPassed TypesFindingSeverity = "passed"

	// TypesFindingSeverityLow captures enum value "low"
	TypesFindingSeverityLow TypesFindingSeverity = "low"

	// TypesFindingSeverityMedium captures enum value "medium"
	TypesFindingSeverityMedium TypesFindingSeverity = "medium"

	// TypesFindingSeverityHigh captures enum value "high"
	TypesFindingSeverityHigh TypesFindingSeverity = "high"

	// TypesFindingSeverityCritical captures enum value "critical"
	TypesFindingSeverityCritical TypesFindingSeverity = "critical"
)

// for schema
var typesFindingSeverityEnum []interface{}

func init() {
	var res []TypesFindingSeverity
	if err := json.Unmarshal([]byte(`["none","passed","low","medium","high","critical"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		typesFindingSeverityEnum = append(typesFindingSeverityEnum, v)
	}
}

func (m TypesFindingSeverity) validateTypesFindingSeverityEnum(path, location string, value TypesFindingSeverity) error {
	if err := validate.EnumCase(path, location, value, typesFindingSeverityEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this types finding severity
func (m TypesFindingSeverity) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTypesFindingSeverityEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this types finding severity based on context it is used
func (m TypesFindingSeverity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
