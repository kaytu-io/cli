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

// GitlabComKeibiengineKeibiEnginePkgComplianceEsFinding gitlab com keibiengine keibi engine pkg compliance es finding
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_compliance_es.Finding
type GitlabComKeibiengineKeibiEnginePkgComplianceEsFinding struct {

	// ID
	ID string `json:"ID,omitempty"`

	// benchmark ID
	BenchmarkID string `json:"benchmarkID,omitempty"`

	// compliance job ID
	ComplianceJobID int64 `json:"complianceJobID,omitempty"`

	// connection ID
	ConnectionID string `json:"connectionID,omitempty"`

	// connector
	Connector SourceType `json:"connector,omitempty"`

	// described at
	DescribedAt int64 `json:"describedAt,omitempty"`

	// evaluated at
	EvaluatedAt int64 `json:"evaluatedAt,omitempty"`

	// evaluator
	Evaluator string `json:"evaluator,omitempty"`

	// policy ID
	PolicyID string `json:"policyID,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// resource ID
	ResourceID string `json:"resourceID,omitempty"`

	// resource location
	ResourceLocation string `json:"resourceLocation,omitempty"`

	// resource name
	ResourceName string `json:"resourceName,omitempty"`

	// resource type
	ResourceType string `json:"resourceType,omitempty"`

	// result
	Result TypesComplianceResult `json:"result,omitempty"`

	// schedule job ID
	ScheduleJobID int64 `json:"scheduleJobID,omitempty"`

	// severity
	Severity string `json:"severity,omitempty"`

	// state active
	StateActive bool `json:"stateActive,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg compliance es finding
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceEsFinding) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgComplianceEsFinding) validateConnector(formats strfmt.Registry) error {
	if swag.IsZero(m.Connector) { // not required
		return nil
	}

	if err := m.Connector.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("connector")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("connector")
		}
		return err
	}

	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgComplianceEsFinding) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(m.Result) { // not required
		return nil
	}

	if err := m.Result.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("result")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("result")
		}
		return err
	}

	return nil
}

// ContextValidate validate this gitlab com keibiengine keibi engine pkg compliance es finding based on the context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceEsFinding) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgComplianceEsFinding) contextValidateConnector(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Connector) { // not required
		return nil
	}

	if err := m.Connector.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("connector")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("connector")
		}
		return err
	}

	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgComplianceEsFinding) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	if err := m.Result.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("result")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("result")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceEsFinding) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgComplianceEsFinding) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgComplianceEsFinding
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
