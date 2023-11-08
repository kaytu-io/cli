// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GithubComKaytuIoKaytuEnginePkgAlertingAPIJiraAndStackResponse github com kaytu io kaytu engine pkg alerting api jira and stack response
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_alerting_api.JiraAndStackResponse
type GithubComKaytuIoKaytuEnginePkgAlertingAPIJiraAndStackResponse struct {

	// action id
	ActionID int64 `json:"action_id,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg alerting api jira and stack response
func (m *GithubComKaytuIoKaytuEnginePkgAlertingAPIJiraAndStackResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this github com kaytu io kaytu engine pkg alerting api jira and stack response based on context it is used
func (m *GithubComKaytuIoKaytuEnginePkgAlertingAPIJiraAndStackResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgAlertingAPIJiraAndStackResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgAlertingAPIJiraAndStackResponse) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgAlertingAPIJiraAndStackResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
