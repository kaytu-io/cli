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

// SourceHealthStatus source health status
//
// swagger:model source.HealthStatus
type SourceHealthStatus string

func NewSourceHealthStatus(value SourceHealthStatus) *SourceHealthStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated SourceHealthStatus.
func (m SourceHealthStatus) Pointer() *SourceHealthStatus {
	return &m
}

const (

	// SourceHealthStatusEmpty captures enum value ""
	SourceHealthStatusEmpty SourceHealthStatus = ""

	// SourceHealthStatusHealthy captures enum value "healthy"
	SourceHealthStatusHealthy SourceHealthStatus = "healthy"

	// SourceHealthStatusUnhealthy captures enum value "unhealthy"
	SourceHealthStatusUnhealthy SourceHealthStatus = "unhealthy"
)

// for schema
var sourceHealthStatusEnum []interface{}

func init() {
	var res []SourceHealthStatus
	if err := json.Unmarshal([]byte(`["","healthy","unhealthy"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sourceHealthStatusEnum = append(sourceHealthStatusEnum, v)
	}
}

func (m SourceHealthStatus) validateSourceHealthStatusEnum(path, location string, value SourceHealthStatus) error {
	if err := validate.EnumCase(path, location, value, sourceHealthStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this source health status
func (m SourceHealthStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSourceHealthStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this source health status based on context it is used
func (m SourceHealthStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
