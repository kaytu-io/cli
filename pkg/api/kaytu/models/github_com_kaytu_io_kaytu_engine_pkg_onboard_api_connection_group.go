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

// GithubComKaytuIoKaytuEnginePkgOnboardAPIConnectionGroup github com kaytu io kaytu engine pkg onboard api connection group
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_onboard_api.ConnectionGroup
type GithubComKaytuIoKaytuEnginePkgOnboardAPIConnectionGroup struct {

	// connection ids
	// Example: ["[\"1e8ac3bf-c268-4a87-9374-ce04cc40a596\"]"]
	ConnectionIds []string `json:"connectionIds"`

	// connections
	Connections []*GithubComKaytuIoKaytuEnginePkgOnboardAPIConnection `json:"connections"`

	// name
	// Example: UltraSightApplication
	Name string `json:"name,omitempty"`

	// query
	// Example: SELECT kaytu_id FROM kaytu_connections WHERE tags-\u003e'application' IS NOT NULL AND tags-\u003e'application' @\u003e '\"UltraSight\"'
	Query string `json:"query,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg onboard api connection group
func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPIConnectionGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnections(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPIConnectionGroup) validateConnections(formats strfmt.Registry) error {
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

// ContextValidate validate this github com kaytu io kaytu engine pkg onboard api connection group based on the context it is used
func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPIConnectionGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnections(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPIConnectionGroup) contextValidateConnections(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Connections); i++ {

		if m.Connections[i] != nil {

			if swag.IsZero(m.Connections[i]) { // not required
				return nil
			}

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
func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPIConnectionGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPIConnectionGroup) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgOnboardAPIConnectionGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
