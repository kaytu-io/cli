// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatus github com kaytu io kaytu engine pkg workspace api workspace status
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_workspace_api.WorkspaceStatus
type GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatus string

func NewGithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatus(value GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatus) *GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatus.
func (m GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatus) Pointer() *GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatus {
	return &m
}

const (

	// GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatusPROVISIONED captures enum value "PROVISIONED"
	GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatusPROVISIONED GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatus = "PROVISIONED"

	// GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatusBOOTSTRAPPING captures enum value "BOOTSTRAPPING"
	GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatusBOOTSTRAPPING GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatus = "BOOTSTRAPPING"

	// GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatusPROVISIONING captures enum value "PROVISIONING"
	GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatusPROVISIONING GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatus = "PROVISIONING"

	// GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatusPROVISIONINGFAILED captures enum value "PROVISIONING_FAILED"
	GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatusPROVISIONINGFAILED GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatus = "PROVISIONING_FAILED"

	// GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatusDELETING captures enum value "DELETING"
	GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatusDELETING GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatus = "DELETING"

	// GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatusDELETED captures enum value "DELETED"
	GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatusDELETED GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatus = "DELETED"

	// GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatusSUSPENDING captures enum value "SUSPENDING"
	GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatusSUSPENDING GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatus = "SUSPENDING"

	// GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatusSUSPENDED captures enum value "SUSPENDED"
	GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatusSUSPENDED GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatus = "SUSPENDED"
)

// for schema
var githubComKaytuIoKaytuEnginePkgWorkspaceApiWorkspaceStatusEnum []interface{}

func init() {
	var res []GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatus
	if err := json.Unmarshal([]byte(`["PROVISIONED","BOOTSTRAPPING","PROVISIONING","PROVISIONING_FAILED","DELETING","DELETED","SUSPENDING","SUSPENDED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		githubComKaytuIoKaytuEnginePkgWorkspaceApiWorkspaceStatusEnum = append(githubComKaytuIoKaytuEnginePkgWorkspaceApiWorkspaceStatusEnum, v)
	}
}

func (m GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatus) validateGithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatusEnum(path, location string, value GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatus) error {
	if err := validate.EnumCase(path, location, value, githubComKaytuIoKaytuEnginePkgWorkspaceApiWorkspaceStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this github com kaytu io kaytu engine pkg workspace api workspace status
func (m GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateGithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this github com kaytu io kaytu engine pkg workspace api workspace status based on context it is used
func (m GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
