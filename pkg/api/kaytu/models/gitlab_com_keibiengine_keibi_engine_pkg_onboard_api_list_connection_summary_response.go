// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GitlabComKeibiengineKeibiEnginePkgOnboardAPIListConnectionSummaryResponse gitlab com keibiengine keibi engine pkg onboard api list connection summary response
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_onboard_api.ListConnectionSummaryResponse
type GitlabComKeibiengineKeibiEnginePkgOnboardAPIListConnectionSummaryResponse struct {

	// total resource count
	TotalResourceCount int64 `json:"TotalResourceCount,omitempty"`

	// connection count
	ConnectionCount int64 `json:"connectionCount,omitempty"`

	// connections
	Connections []*GitlabComKeibiengineKeibiEnginePkgOnboardAPIConnection `json:"connections"`

	// total cost
	TotalCost float64 `json:"totalCost,omitempty"`

	// total disabled count
	TotalDisabledCount int64 `json:"totalDisabledCount,omitempty"`

	// total unhealthy count
	TotalUnhealthyCount int64 `json:"totalUnhealthyCount,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg onboard api list connection summary response
func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPIListConnectionSummaryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnections(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPIListConnectionSummaryResponse) validateConnections(formats strfmt.Registry) error {
	if swag.IsZero(m.Connections) { // not required
		return nil
	}

	for i := 0; i < len(m.Connections); i++ {
		if swag.IsZero(m.Connections[i]) { // not required
			continue
		}

		if m.Connections[i] != nil {
			if err := m.Connections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("connections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this gitlab com keibiengine keibi engine pkg onboard api list connection summary response based on the context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPIListConnectionSummaryResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnections(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPIListConnectionSummaryResponse) contextValidateConnections(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Connections); i++ {

		if m.Connections[i] != nil {
			if err := m.Connections[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("connections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPIListConnectionSummaryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgOnboardAPIListConnectionSummaryResponse) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgOnboardAPIListConnectionSummaryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
