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

// NewGetOnboardAPIV1CatalogMetricsParams creates a new GetOnboardAPIV1CatalogMetricsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOnboardAPIV1CatalogMetricsParams() *GetOnboardAPIV1CatalogMetricsParams {
	return &GetOnboardAPIV1CatalogMetricsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOnboardAPIV1CatalogMetricsParamsWithTimeout creates a new GetOnboardAPIV1CatalogMetricsParams object
// with the ability to set a timeout on a request.
func NewGetOnboardAPIV1CatalogMetricsParamsWithTimeout(timeout time.Duration) *GetOnboardAPIV1CatalogMetricsParams {
	return &GetOnboardAPIV1CatalogMetricsParams{
		timeout: timeout,
	}
}

// NewGetOnboardAPIV1CatalogMetricsParamsWithContext creates a new GetOnboardAPIV1CatalogMetricsParams object
// with the ability to set a context for a request.
func NewGetOnboardAPIV1CatalogMetricsParamsWithContext(ctx context.Context) *GetOnboardAPIV1CatalogMetricsParams {
	return &GetOnboardAPIV1CatalogMetricsParams{
		Context: ctx,
	}
}

// NewGetOnboardAPIV1CatalogMetricsParamsWithHTTPClient creates a new GetOnboardAPIV1CatalogMetricsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOnboardAPIV1CatalogMetricsParamsWithHTTPClient(client *http.Client) *GetOnboardAPIV1CatalogMetricsParams {
	return &GetOnboardAPIV1CatalogMetricsParams{
		HTTPClient: client,
	}
}

/*
GetOnboardAPIV1CatalogMetricsParams contains all the parameters to send to the API endpoint

	for the get onboard API v1 catalog metrics operation.

	Typically these are written to a http.Request.
*/
type GetOnboardAPIV1CatalogMetricsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get onboard API v1 catalog metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOnboardAPIV1CatalogMetricsParams) WithDefaults() *GetOnboardAPIV1CatalogMetricsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get onboard API v1 catalog metrics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOnboardAPIV1CatalogMetricsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get onboard API v1 catalog metrics params
func (o *GetOnboardAPIV1CatalogMetricsParams) WithTimeout(timeout time.Duration) *GetOnboardAPIV1CatalogMetricsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get onboard API v1 catalog metrics params
func (o *GetOnboardAPIV1CatalogMetricsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get onboard API v1 catalog metrics params
func (o *GetOnboardAPIV1CatalogMetricsParams) WithContext(ctx context.Context) *GetOnboardAPIV1CatalogMetricsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get onboard API v1 catalog metrics params
func (o *GetOnboardAPIV1CatalogMetricsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get onboard API v1 catalog metrics params
func (o *GetOnboardAPIV1CatalogMetricsParams) WithHTTPClient(client *http.Client) *GetOnboardAPIV1CatalogMetricsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get onboard API v1 catalog metrics params
func (o *GetOnboardAPIV1CatalogMetricsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetOnboardAPIV1CatalogMetricsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
