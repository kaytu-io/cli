// Code generated by go-swagger; DO NOT EDIT.

package onboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

// NewGetOnboardAPIV1ConnectionsCountParams creates a new GetOnboardAPIV1ConnectionsCountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOnboardAPIV1ConnectionsCountParams() *GetOnboardAPIV1ConnectionsCountParams {
	return &GetOnboardAPIV1ConnectionsCountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOnboardAPIV1ConnectionsCountParamsWithTimeout creates a new GetOnboardAPIV1ConnectionsCountParams object
// with the ability to set a timeout on a request.
func NewGetOnboardAPIV1ConnectionsCountParamsWithTimeout(timeout time.Duration) *GetOnboardAPIV1ConnectionsCountParams {
	return &GetOnboardAPIV1ConnectionsCountParams{
		timeout: timeout,
	}
}

// NewGetOnboardAPIV1ConnectionsCountParamsWithContext creates a new GetOnboardAPIV1ConnectionsCountParams object
// with the ability to set a context for a request.
func NewGetOnboardAPIV1ConnectionsCountParamsWithContext(ctx context.Context) *GetOnboardAPIV1ConnectionsCountParams {
	return &GetOnboardAPIV1ConnectionsCountParams{
		Context: ctx,
	}
}

// NewGetOnboardAPIV1ConnectionsCountParamsWithHTTPClient creates a new GetOnboardAPIV1ConnectionsCountParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOnboardAPIV1ConnectionsCountParamsWithHTTPClient(client *http.Client) *GetOnboardAPIV1ConnectionsCountParams {
	return &GetOnboardAPIV1ConnectionsCountParams{
		HTTPClient: client,
	}
}

/*
GetOnboardAPIV1ConnectionsCountParams contains all the parameters to send to the API endpoint

	for the get onboard API v1 connections count operation.

	Typically these are written to a http.Request.
*/
type GetOnboardAPIV1ConnectionsCountParams struct {

	/* Type.

	   Request
	*/
	Type *models.GitlabComKeibiengineKeibiEnginePkgOnboardAPIConnectionCountRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get onboard API v1 connections count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOnboardAPIV1ConnectionsCountParams) WithDefaults() *GetOnboardAPIV1ConnectionsCountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get onboard API v1 connections count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOnboardAPIV1ConnectionsCountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get onboard API v1 connections count params
func (o *GetOnboardAPIV1ConnectionsCountParams) WithTimeout(timeout time.Duration) *GetOnboardAPIV1ConnectionsCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get onboard API v1 connections count params
func (o *GetOnboardAPIV1ConnectionsCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get onboard API v1 connections count params
func (o *GetOnboardAPIV1ConnectionsCountParams) WithContext(ctx context.Context) *GetOnboardAPIV1ConnectionsCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get onboard API v1 connections count params
func (o *GetOnboardAPIV1ConnectionsCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get onboard API v1 connections count params
func (o *GetOnboardAPIV1ConnectionsCountParams) WithHTTPClient(client *http.Client) *GetOnboardAPIV1ConnectionsCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get onboard API v1 connections count params
func (o *GetOnboardAPIV1ConnectionsCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithType adds the typeVar to the get onboard API v1 connections count params
func (o *GetOnboardAPIV1ConnectionsCountParams) WithType(typeVar *models.GitlabComKeibiengineKeibiEnginePkgOnboardAPIConnectionCountRequest) *GetOnboardAPIV1ConnectionsCountParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get onboard API v1 connections count params
func (o *GetOnboardAPIV1ConnectionsCountParams) SetType(typeVar *models.GitlabComKeibiengineKeibiEnginePkgOnboardAPIConnectionCountRequest) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetOnboardAPIV1ConnectionsCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Type != nil {
		if err := r.SetBodyParam(o.Type); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
