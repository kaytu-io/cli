// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GithubComKaytuIoKaytuEnginePkgAlertingAPICreateActionReq github com kaytu io kaytu engine pkg alerting api create action req
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_alerting_api.CreateActionReq
type GithubComKaytuIoKaytuEnginePkgAlertingAPICreateActionReq struct {

	// body
	Body string `json:"body,omitempty"`

	// headers
	Headers map[string]string `json:"headers,omitempty"`

	// method
	Method string `json:"method,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg alerting api create action req
func (m *GithubComKaytuIoKaytuEnginePkgAlertingAPICreateActionReq) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this github com kaytu io kaytu engine pkg alerting api create action req based on context it is used
func (m *GithubComKaytuIoKaytuEnginePkgAlertingAPICreateActionReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgAlertingAPICreateActionReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgAlertingAPICreateActionReq) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgAlertingAPICreateActionReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
