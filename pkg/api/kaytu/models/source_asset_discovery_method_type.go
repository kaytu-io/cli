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

// SourceAssetDiscoveryMethodType source asset discovery method type
//
// swagger:model source.AssetDiscoveryMethodType
type SourceAssetDiscoveryMethodType string

func NewSourceAssetDiscoveryMethodType(value SourceAssetDiscoveryMethodType) *SourceAssetDiscoveryMethodType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated SourceAssetDiscoveryMethodType.
func (m SourceAssetDiscoveryMethodType) Pointer() *SourceAssetDiscoveryMethodType {
	return &m
}

const (

	// SourceAssetDiscoveryMethodTypeScheduled captures enum value "scheduled"
	SourceAssetDiscoveryMethodTypeScheduled SourceAssetDiscoveryMethodType = "scheduled"
)

// for schema
var sourceAssetDiscoveryMethodTypeEnum []interface{}

func init() {
	var res []SourceAssetDiscoveryMethodType
	if err := json.Unmarshal([]byte(`["scheduled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sourceAssetDiscoveryMethodTypeEnum = append(sourceAssetDiscoveryMethodTypeEnum, v)
	}
}

func (m SourceAssetDiscoveryMethodType) validateSourceAssetDiscoveryMethodTypeEnum(path, location string, value SourceAssetDiscoveryMethodType) error {
	if err := validate.EnumCase(path, location, value, sourceAssetDiscoveryMethodTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this source asset discovery method type
func (m SourceAssetDiscoveryMethodType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSourceAssetDiscoveryMethodTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this source asset discovery method type based on context it is used
func (m SourceAssetDiscoveryMethodType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}