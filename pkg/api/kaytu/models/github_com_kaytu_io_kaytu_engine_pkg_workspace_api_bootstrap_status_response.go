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

// GithubComKaytuIoKaytuEnginePkgWorkspaceAPIBootstrapStatusResponse github com kaytu io kaytu engine pkg workspace api bootstrap status response
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_workspace_api.BootstrapStatusResponse
type GithubComKaytuIoKaytuEnginePkgWorkspaceAPIBootstrapStatusResponse struct {

	// status
	Status GithubComKaytuIoKaytuEnginePkgWorkspaceAPIBootstrapStatus `json:"status,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg workspace api bootstrap status response
func (m *GithubComKaytuIoKaytuEnginePkgWorkspaceAPIBootstrapStatusResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgWorkspaceAPIBootstrapStatusResponse) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

// ContextValidate validate this github com kaytu io kaytu engine pkg workspace api bootstrap status response based on the context it is used
func (m *GithubComKaytuIoKaytuEnginePkgWorkspaceAPIBootstrapStatusResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgWorkspaceAPIBootstrapStatusResponse) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgWorkspaceAPIBootstrapStatusResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgWorkspaceAPIBootstrapStatusResponse) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgWorkspaceAPIBootstrapStatusResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}