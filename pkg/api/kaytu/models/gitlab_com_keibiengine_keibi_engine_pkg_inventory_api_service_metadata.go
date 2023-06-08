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

// GitlabComKeibiengineKeibiEnginePkgInventoryAPIServiceMetadata gitlab com keibiengine keibi engine pkg inventory api service metadata
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_inventory_api.ServiceMetadata
type GitlabComKeibiengineKeibiEnginePkgInventoryAPIServiceMetadata struct {

	// Service Connector
	Connector struct {
		SourceType
	} `json:"connector,omitempty"`

	// List of Cost map service names
	CostMapServiceNames []string `json:"cost_map_service_names"`

	// Cost is supported [yes/no]
	CostSupport bool `json:"cost_support,omitempty"`

	// logo uri
	LogoURI string `json:"logo_uri,omitempty"`

	// List of resource types
	ResourceTypes []string `json:"resource_types"`

	// resource types count
	ResourceTypesCount int64 `json:"resource_types_count,omitempty"`

	// Service Lable
	ServiceLabel string `json:"service_label,omitempty"`

	// Service Name
	ServiceName string `json:"service_name,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg inventory api service metadata
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIServiceMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIServiceMetadata) validateConnector(formats strfmt.Registry) error {
	if swag.IsZero(m.Connector) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this gitlab com keibiengine keibi engine pkg inventory api service metadata based on the context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIServiceMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIServiceMetadata) contextValidateConnector(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIServiceMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgInventoryAPIServiceMetadata) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgInventoryAPIServiceMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
