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
)

// NewGetOnboardAPIV1ConnectorParams creates a new GetOnboardAPIV1ConnectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOnboardAPIV1ConnectorParams() *GetOnboardAPIV1ConnectorParams {
	return &GetOnboardAPIV1ConnectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOnboardAPIV1ConnectorParamsWithTimeout creates a new GetOnboardAPIV1ConnectorParams object
// with the ability to set a timeout on a request.
func NewGetOnboardAPIV1ConnectorParamsWithTimeout(timeout time.Duration) *GetOnboardAPIV1ConnectorParams {
	return &GetOnboardAPIV1ConnectorParams{
		timeout: timeout,
	}
}

// NewGetOnboardAPIV1ConnectorParamsWithContext creates a new GetOnboardAPIV1ConnectorParams object
// with the ability to set a context for a request.
func NewGetOnboardAPIV1ConnectorParamsWithContext(ctx context.Context) *GetOnboardAPIV1ConnectorParams {
	return &GetOnboardAPIV1ConnectorParams{
		Context: ctx,
	}
}

// NewGetOnboardAPIV1ConnectorParamsWithHTTPClient creates a new GetOnboardAPIV1ConnectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOnboardAPIV1ConnectorParamsWithHTTPClient(client *http.Client) *GetOnboardAPIV1ConnectorParams {
	return &GetOnboardAPIV1ConnectorParams{
		HTTPClient: client,
	}
}

/*
GetOnboardAPIV1ConnectorParams contains all the parameters to send to the API endpoint

	for the get onboard API v1 connector operation.

	Typically these are written to a http.Request.
*/
type GetOnboardAPIV1ConnectorParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get onboard API v1 connector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOnboardAPIV1ConnectorParams) WithDefaults() *GetOnboardAPIV1ConnectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get onboard API v1 connector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOnboardAPIV1ConnectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get onboard API v1 connector params
func (o *GetOnboardAPIV1ConnectorParams) WithTimeout(timeout time.Duration) *GetOnboardAPIV1ConnectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get onboard API v1 connector params
func (o *GetOnboardAPIV1ConnectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get onboard API v1 connector params
func (o *GetOnboardAPIV1ConnectorParams) WithContext(ctx context.Context) *GetOnboardAPIV1ConnectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get onboard API v1 connector params
func (o *GetOnboardAPIV1ConnectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get onboard API v1 connector params
func (o *GetOnboardAPIV1ConnectorParams) WithHTTPClient(client *http.Client) *GetOnboardAPIV1ConnectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get onboard API v1 connector params
func (o *GetOnboardAPIV1ConnectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetOnboardAPIV1ConnectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
