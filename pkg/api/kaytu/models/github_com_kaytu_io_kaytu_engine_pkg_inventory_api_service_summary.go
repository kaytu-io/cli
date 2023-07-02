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

// GithubComKaytuIoKaytuEnginePkgInventoryAPIServiceSummary github com kaytu io kaytu engine pkg inventory api service summary
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_inventory_api.ServiceSummary
type GithubComKaytuIoKaytuEnginePkgInventoryAPIServiceSummary struct {

	// Cloud provider
	// Example: Azure
	Connector SourceType `json:"connector,omitempty"`

	// Number of Resources
	// Example: 100
	ResourceCount int64 `json:"resourceCount,omitempty"`

	// Service Label
	// Example: Compute
	ServiceLabel string `json:"serviceLabel,omitempty"`

	// Service Name
	// Example: compute
	ServiceName string `json:"serviceName,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg inventory api service summary
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIServiceSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIServiceSummary) validateConnector(formats strfmt.Registry) error {
	if swag.IsZero(m.Connector) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this github com kaytu io kaytu engine pkg inventory api service summary based on the context it is used
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIServiceSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIServiceSummary) contextValidateConnector(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIServiceSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgInventoryAPIServiceSummary) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgInventoryAPIServiceSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}