// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GithubComKaytuIoKaytuEnginePkgOnboardAPICreateConnectionResponse github com kaytu io kaytu engine pkg onboard api create connection response
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_onboard_api.CreateConnectionResponse
type GithubComKaytuIoKaytuEnginePkgOnboardAPICreateConnectionResponse struct {

	// id
	ID string `json:"id,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg onboard api create connection response
func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICreateConnectionResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this github com kaytu io kaytu engine pkg onboard api create connection response based on context it is used
func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICreateConnectionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICreateConnectionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgOnboardAPICreateConnectionResponse) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgOnboardAPICreateConnectionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
