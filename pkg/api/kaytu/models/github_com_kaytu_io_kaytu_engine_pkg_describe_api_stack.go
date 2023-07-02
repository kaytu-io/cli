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

// GithubComKaytuIoKaytuEnginePkgDescribeAPIStack github com kaytu io kaytu engine pkg describe api stack
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_describe_api.Stack
type GithubComKaytuIoKaytuEnginePkgDescribeAPIStack struct {

	// Accounts included in the stack
	// Example: ["[0123456789]"]
	AccountIds []string `json:"accountIds"`

	// Stack creation date
	// Example: 2023-06-01T17:00:00.000000Z
	CreatedAt string `json:"createdAt,omitempty"`

	// Stack evaluations history, including insight evaluations and compliance evaluations
	Evaluations []*GithubComKaytuIoKaytuEnginePkgDescribeAPIStackEvaluation `json:"evaluations"`

	// Stack resource types
	// Example: ["[Microsoft.Compute/virtualMachines]"]
	ResourceTypes []string `json:"resourceTypes"`

	// Stack resources list
	// Example: ["[/subscriptions/123/resourceGroups/rg-1/providers/Microsoft.Compute/virtualMachines/vm-1]"]
	Resources []string `json:"resources"`

	// Stack unique identifier
	// Example: stack-twr32a5d-5as5-4ffe-b1cc-e32w1ast87s0
	// Required: true
	StackID *string `json:"stackId"`

	// Stack tags
	Tags map[string][]string `json:"tags,omitempty"`

	// Stack last update date
	// Example: 2023-06-01T17:00:00.000000Z
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg describe api stack
func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIStack) Validate(formats strfmt.Registry) error {
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

func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIStack) validateEvaluations(formats strfmt.Registry) error {
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

func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIStack) validateStackID(formats strfmt.Registry) error {

	if err := validate.Required("stackId", "body", m.StackID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this github com kaytu io kaytu engine pkg describe api stack based on the context it is used
func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIStack) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEvaluations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIStack) contextValidateEvaluations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Evaluations); i++ {

		if m.Evaluations[i] != nil {

			if swag.IsZero(m.Evaluations[i]) { // not required
				return nil
			}

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
func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIStack) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgDescribeAPIStack) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgDescribeAPIStack
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}