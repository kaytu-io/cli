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

// GitlabComKeibiengineKeibiEnginePkgComplianceAPIInsightResult gitlab com keibiengine keibi engine pkg compliance api insight result
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_compliance_api.InsightResult
type GitlabComKeibiengineKeibiEnginePkgComplianceAPIInsightResult struct {

	// Connection ID
	ConnectionID string `json:"connectionID,omitempty"`

	// Connections
	Connections []*GitlabComKeibiengineKeibiEnginePkgComplianceAPIInsightConnection `json:"connections"`

	// Insight Details
	Details GitlabComKeibiengineKeibiEnginePkgComplianceAPIInsightDetail `json:"details,omitempty"`

	// Time of Execution
	ExecutedAt string `json:"executedAt,omitempty"`

	// Insight ID
	InsightID int64 `json:"insightID,omitempty"`

	// Job ID
	JobID int64 `json:"jobID,omitempty"`

	// Locations
	Locations []string `json:"locations"`

	// Result
	Result int64 `json:"result,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg compliance api insight result
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIInsightResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIInsightResult) validateConnections(formats strfmt.Registry) error {
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

func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIInsightResult) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.Details) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this gitlab com keibiengine keibi engine pkg compliance api insight result based on the context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIInsightResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnections(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIInsightResult) contextValidateConnections(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIInsightResult) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIInsightResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceAPIInsightResult) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgComplianceAPIInsightResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
