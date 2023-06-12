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

// GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIWorkspaceResponse gitlab com keibiengine keibi engine pkg workspace api workspace response
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_workspace_api.WorkspaceResponse
type GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIWorkspaceResponse struct {

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// organization
	Organization *GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIOrganizationResponse `json:"organization,omitempty"`

	// owner Id
	OwnerID string `json:"ownerId,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// tier
	Tier string `json:"tier,omitempty"`

	// uri
	URI string `json:"uri,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this gitlab com keibiengine keibi engine pkg workspace api workspace response
func (m *GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIWorkspaceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIWorkspaceResponse) validateOrganization(formats strfmt.Registry) error {
	if swag.IsZero(m.Organization) { // not required
		return nil
	}

	if m.Organization != nil {
		if err := m.Organization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("organization")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this gitlab com keibiengine keibi engine pkg workspace api workspace response based on the context it is used
func (m *GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIWorkspaceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrganization(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIWorkspaceResponse) contextValidateOrganization(ctx context.Context, formats strfmt.Registry) error {

	if m.Organization != nil {

		if swag.IsZero(m.Organization) { // not required
			return nil
		}

		if err := m.Organization.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("organization")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIWorkspaceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIWorkspaceResponse) UnmarshalBinary(b []byte) error {
	var res GitlabComKeibiengineKeibiEnginePkgWorkspaceAPIWorkspaceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
