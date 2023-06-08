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

// GitlabComKeibiengineKeibiEnginePkgComplianceAPIDirectionType gitlab com keibiengine keibi engine pkg compliance api direction type
//
// swagger:model gitlab_com_keibiengine_keibi-engine_pkg_compliance_api.DirectionType
type GitlabComKeibiengineKeibiEnginePkgComplianceAPIDirectionType string

func NewGitlabComKeibiengineKeibiEnginePkgComplianceAPIDirectionType(value GitlabComKeibiengineKeibiEnginePkgComplianceAPIDirectionType) *GitlabComKeibiengineKeibiEnginePkgComplianceAPIDirectionType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated GitlabComKeibiengineKeibiEnginePkgComplianceAPIDirectionType.
func (m GitlabComKeibiengineKeibiEnginePkgComplianceAPIDirectionType) Pointer() *GitlabComKeibiengineKeibiEnginePkgComplianceAPIDirectionType {
	return &m
}

const (

	// GitlabComKeibiengineKeibiEnginePkgComplianceAPIDirectionTypeAsc captures enum value "asc"
	GitlabComKeibiengineKeibiEnginePkgComplianceAPIDirectionTypeAsc GitlabComKeibiengineKeibiEnginePkgComplianceAPIDirectionType = "asc"

	// GitlabComKeibiengineKeibiEnginePkgComplianceAPIDirectionTypeDesc captures enum value "desc"
	GitlabComKeibiengineKeibiEnginePkgComplianceAPIDirectionTypeDesc GitlabComKeibiengineKeibiEnginePkgComplianceAPIDirectionType = "desc"
)

// for schema
var gitlabComKeibiengineKeibiEnginePkgComplianceApiDirectionTypeEnum []interface{}

func init() {
	var res []GitlabComKeibiengineKeibiEnginePkgComplianceAPIDirectionType
	if err := json.Unmarshal([]byte(`["asc","desc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gitlabComKeibiengineKeibiEnginePkgComplianceApiDirectionTypeEnum = append(gitlabComKeibiengineKeibiEnginePkgComplianceApiDirectionTypeEnum, v)
	}
}

func (m GitlabComKeibiengineKeibiEnginePkgComplianceAPIDirectionType) validateGitlabComKeibiengineKeibiEnginePkgComplianceAPIDirectionTypeEnum(path, location string, value GitlabComKeibiengineKeibiEnginePkgComplianceAPIDirectionType) error {
	if err := validate.EnumCase(path, location, value, gitlabComKeibiengineKeibiEnginePkgComplianceApiDirectionTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this gitlab com keibiengine keibi engine pkg compliance api direction type
func (m GitlabComKeibiengineKeibiEnginePkgComplianceAPIDirectionType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateGitlabComKeibiengineKeibiEnginePkgComplianceAPIDirectionTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this gitlab com keibiengine keibi engine pkg compliance api direction type based on context it is used
func (m GitlabComKeibiengineKeibiEnginePkgComplianceAPIDirectionType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
