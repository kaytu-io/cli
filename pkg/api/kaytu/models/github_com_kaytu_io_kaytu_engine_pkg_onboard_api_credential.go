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
	"github.com/go-openapi/validate"
)

// GithubComKaytuIoKaytuEnginePkgOnboardAPICredential github com kaytu io kaytu engine pkg onboard api credential
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_onboard_api.Credential
type GithubComKaytuIoKaytuEnginePkgOnboardAPICredential struct {

	// archived connections
	// Example: 0
	// Maximum: 1000
	// Minimum: 0
	ArchivedConnections *int64 `json:"archived_connections,omitempty"`

	// auto onboard enabled
	// Example: false
	AutoOnboardEnabled bool `json:"autoOnboardEnabled,omitempty"`

	// config
	Config interface{} `json:"config,omitempty"`

	// connections
	Connections []*GithubComKaytuIoKaytuEnginePkgOnboardAPIConnection `json:"connections"`

	// connector type
	// Example: AWS
	ConnectorType SourceType `json:"connectorType,omitempty"`

	// credential type
	// Example: manual-aws-org
	CredentialType GithubComKaytuIoKaytuEnginePkgOnboardAPICredentialType `json:"credentialType,omitempty"`

	// disabled connections
	// Example: 0
	// Maximum: 1000
	// Minimum: 0
	DisabledConnections *int64 `json:"disabled_connections,omitempty"`

	// discovered connections
	// Example: 50
	// Maximum: 100
	// Minimum: 0
	DiscoveredConnections *int64 `json:"discovered_connections,omitempty"`

	// enabled
	// Example: true
	Enabled bool `json:"enabled,omitempty"`

	// health reason
	HealthReason string `json:"healthReason,omitempty"`

	// health status
	// Example: healthy
	HealthStatus SourceHealthStatus `json:"healthStatus,omitempty"`

	// id
	// Example: 1028642a-b22e-26ha-c5h2-22nl254678m5
	ID string `json:"id,omitempty"`

	// last health check time
	// Example: 2023-06-03T12:21:33.406928Z
	// Format: date-time
	LastHealthCheckTime strfmt.DateTime `json:"lastHealthCheckTime,omitempty"`

	// metadata
	Metadata interface{} `json:"metadata,omitempty"`

	// name
	// Example: a-1mahsl7lzk
	Name string `json:"name,omitempty"`

	// onboard date
	// Example: 2023-06-03T12:21:33.406928Z
	// Format: date-time
	OnboardDate strfmt.DateTime `json:"onboardDate,omitempty"`

	// onboard connections
	// Example: 250
	// Maximum: 1000
	// Minimum: 0
	OnboardConnections *int64 `json:"onboard_connections,omitempty"`

	// spend discovery
	SpendDiscovery bool `json:"spendDiscovery,omitempty"`

	// total connections
	// Example: 300
	// Maximum: 1000
	// Minimum: 0
	TotalConnections *int64 `json:"total_connections,omitempty"`

	// unhealthy connections
	// Example: 50
	// Maximum: 100
	// Minimum: 0
	UnhealthyConnections *int64 `json:"unhealthy_connections,omitempty"`

	// version
	Version int64 `json:"version,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg onboard api credential
func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICredential) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArchivedConnections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectorType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisabledConnections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiscoveredConnections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastHealthCheckTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnboardDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnboardConnections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalConnections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnhealthyConnections(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICredential) validateArchivedConnections(formats strfmt.Registry) error {
	if swag.IsZero(m.ArchivedConnections) { // not required
		return nil
	}

	if err := validate.MinimumInt("archived_connections", "body", *m.ArchivedConnections, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("archived_connections", "body", *m.ArchivedConnections, 1000, false); err != nil {
		return err
	}

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICredential) validateConnections(formats strfmt.Registry) error {
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

func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICredential) validateConnectorType(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectorType) { // not required
		return nil
	}

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICredential) validateCredentialType(formats strfmt.Registry) error {
	if swag.IsZero(m.CredentialType) { // not required
		return nil
	}

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICredential) validateDisabledConnections(formats strfmt.Registry) error {
	if swag.IsZero(m.DisabledConnections) { // not required
		return nil
	}

	if err := validate.MinimumInt("disabled_connections", "body", *m.DisabledConnections, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("disabled_connections", "body", *m.DisabledConnections, 1000, false); err != nil {
		return err
	}

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICredential) validateDiscoveredConnections(formats strfmt.Registry) error {
	if swag.IsZero(m.DiscoveredConnections) { // not required
		return nil
	}

	if err := validate.MinimumInt("discovered_connections", "body", *m.DiscoveredConnections, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("discovered_connections", "body", *m.DiscoveredConnections, 100, false); err != nil {
		return err
	}

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICredential) validateHealthStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.HealthStatus) { // not required
		return nil
	}

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICredential) validateLastHealthCheckTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastHealthCheckTime) { // not required
		return nil
	}

	if err := validate.FormatOf("lastHealthCheckTime", "body", "date-time", m.LastHealthCheckTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICredential) validateOnboardDate(formats strfmt.Registry) error {
	if swag.IsZero(m.OnboardDate) { // not required
		return nil
	}

	if err := validate.FormatOf("onboardDate", "body", "date-time", m.OnboardDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICredential) validateOnboardConnections(formats strfmt.Registry) error {
	if swag.IsZero(m.OnboardConnections) { // not required
		return nil
	}

	if err := validate.MinimumInt("onboard_connections", "body", *m.OnboardConnections, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("onboard_connections", "body", *m.OnboardConnections, 1000, false); err != nil {
		return err
	}

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICredential) validateTotalConnections(formats strfmt.Registry) error {
	if swag.IsZero(m.TotalConnections) { // not required
		return nil
	}

	if err := validate.MinimumInt("total_connections", "body", *m.TotalConnections, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("total_connections", "body", *m.TotalConnections, 1000, false); err != nil {
		return err
	}

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICredential) validateUnhealthyConnections(formats strfmt.Registry) error {
	if swag.IsZero(m.UnhealthyConnections) { // not required
		return nil
	}

	if err := validate.MinimumInt("unhealthy_connections", "body", *m.UnhealthyConnections, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("unhealthy_connections", "body", *m.UnhealthyConnections, 100, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this github com kaytu io kaytu engine pkg onboard api credential based on the context it is used
func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICredential) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnections(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectorType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCredentialType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHealthStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICredential) contextValidateConnections(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICredential) contextValidateConnectorType(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICredential) contextValidateCredentialType(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICredential) contextValidateHealthStatus(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICredential) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICredential) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgOnboardAPICredential
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
