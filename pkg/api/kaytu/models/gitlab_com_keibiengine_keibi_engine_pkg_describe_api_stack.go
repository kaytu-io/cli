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

// GitlabComKeibiengineKeibiEnginePkgDescribeAPIStack gitlab com keibiengine keibi engine pkg describe api stack
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_describe_api.Stack
type GitlabComKeibiengineKeibiEnginePkgDescribeAPIStack struct {

	// account ids
	AccountIds []string `json:"accountIds"`

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// evaluations
	Evaluations []*GitlabComKeibiengineKeibiEnginePkgDescribeAPIStackEvaluation `json:"evaluations"`

	// resources
	Resources []string `json:"resources"`

	// stack Id
	// Required: true
	StackID *string `json:"stackId"`

	// tags
	Tags map[string][]string `json:"tags,omitempty"`

	// updated at
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg describe api stack
func (m *GitlabComKeibiengineKeibiEnginePkgDescribeAPIStack) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvaluations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStackID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgDescribeAPIStack) validateEvaluations(formats strfmt.Registry) error {
	if swag.IsZero(m.Evaluations) { // not required
		return nil
	}

	for i := 0; i < len(m.Evaluations); i++ {
		if swag.IsZero(m.Evaluations[i]) { // not required
			continue
		}

		if m.Evaluations[i] != nil {
			if err := m.Evaluations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("evaluations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("evaluations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgDescribeAPIStack) validateStackID(formats strfmt.Registry) error {

	if err := validate.Required("stackId", "body", m.StackID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this gitlab com keibiengine keibi engine pkg describe api stack based on the context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgDescribeAPIStack) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEvaluations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgDescribeAPIStack) contextValidateEvaluations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Evaluations); i++ {

		if m.Evaluations[i] != nil {
			if err := m.Evaluations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("evaluations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("evaluations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgDescribeAPIStack) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgDescribeAPIStack) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgDescribeAPIStack
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}