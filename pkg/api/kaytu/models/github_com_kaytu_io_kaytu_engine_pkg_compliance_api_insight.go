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
)

// GithubComKaytuIoKaytuEnginePkgComplianceAPIInsight github com kaytu io kaytu engine pkg compliance api insight
//
// swagger:model github_com_kaytu-io_kaytu-engine_pkg_compliance_api.Insight
type GithubComKaytuIoKaytuEnginePkgComplianceAPIInsight struct {

	// Cloud Provider
	// Example: Azure
	Connector SourceType `json:"connector,omitempty"`

	// Description
	// Example: List clusters that have role-based access control (RBAC) disabled
	Description string `json:"description,omitempty"`

	// enabled
	// Example: true
	Enabled bool `json:"enabled,omitempty"`

	// Old Total Result Date
	FirstOldResultDate string `json:"firstOldResultDate,omitempty"`

	// Insight ID
	// Example: 23
	ID int64 `json:"id,omitempty"`

	// internal
	// Example: false
	Internal bool `json:"internal,omitempty"`

	// Links
	Links []string `json:"links"`

	// Logo URL
	LogoURL string `json:"logoURL,omitempty"`

	// Long Title
	// Example: List clusters that have role-based access control (RBAC) disabled
	LongTitle string `json:"longTitle,omitempty"`

	// Number of Old Total Result Value
	// Example: 0
	OldTotalResultValue int64 `json:"oldTotalResultValue,omitempty"`

	// Query
	Query GithubComKaytuIoKaytuEnginePkgComplianceAPIQuery `json:"query,omitempty"`

	// Insight Results
	Result []*GithubComKaytuIoKaytuEnginePkgComplianceAPIInsightResult `json:"result"`

	// Short Title
	// Example: Clusters with no RBAC
	ShortTitle string `json:"shortTitle,omitempty"`

	// Tags
	Tags map[string][]string `json:"tags,omitempty"`

	// Number of Total Result Value
	// Example: 10
	TotalResultValue int64 `json:"totalResultValue,omitempty"`
}

// Validate validates this github com kaytu io kaytu engine pkg compliance api insight
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIInsight) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIInsight) validateConnector(formats strfmt.Registry) error {
	if swag.IsZero(m.Connector) { // not required
		return nil
	}

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIInsight) validateQuery(formats strfmt.Registry) error {
	if swag.IsZero(m.Query) { // not required
		return nil
	}

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIInsight) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(m.Result) { // not required
		return nil
	}

	for i := 0; i < len(m.Result); i++ {
		if swag.IsZero(m.Result[i]) { // not required
			continue
		}

		if m.Result[i] != nil {
			if err := m.Result[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("result" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("result" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this github com kaytu io kaytu engine pkg compliance api insight based on the context it is used
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIInsight) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQuery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIInsight) contextValidateConnector(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIInsight) contextValidateQuery(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIInsight) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Result); i++ {

		if m.Result[i] != nil {

			if swag.IsZero(m.Result[i]) { // not required
				return nil
			}

			if err := m.Result[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("result" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("result" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIInsight) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GithubComKaytuIoKaytuEnginePkgComplianceAPIInsight) UnmarshalBinary(b []byte) error {
	var res GithubComKaytuIoKaytuEnginePkgComplianceAPIInsight
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
