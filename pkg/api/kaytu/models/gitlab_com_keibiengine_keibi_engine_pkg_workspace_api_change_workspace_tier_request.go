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

// GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIChangeWorkspaceTierRequest gitlab com keibiengine keibi engine pkg workspace api change workspace tier request
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_workspace_api.ChangeWorkspaceTierRequest
type GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIChangeWorkspaceTierRequest struct {

	// new name
	NewName GitlabComKeibiengineKeibiEnginePkgWorkspaceAPITier `json:"newName,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg workspace api change workspace tier request
func (m *GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIChangeWorkspaceTierRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNewName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIChangeWorkspaceTierRequest) validateNewName(formats strfmt.Registry) error {
	if swag.IsZero(m.NewName) { // not required
		return nil
	}

	if err := m.NewName.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("newName")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("newName")
		}
		return err
	}

	return nil
}

// ContextValidate validate this gitlab com keibiengine keibi engine pkg workspace api change workspace tier request based on the context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIChangeWorkspaceTierRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNewName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIChangeWorkspaceTierRequest) contextValidateNewName(ctx context.Context, formats strfmt.Registry) error {

	if err := m.NewName.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("newName")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("newName")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIChangeWorkspaceTierRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIChangeWorkspaceTierRequest) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIChangeWorkspaceTierRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}