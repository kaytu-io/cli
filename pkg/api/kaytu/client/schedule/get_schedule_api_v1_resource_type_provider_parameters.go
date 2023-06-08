// Code generated by go-swagger; DO NOT EDIT.

package schedule

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

// NewGetScheduleAPIV1ResourceTypeProviderParams creates a new GetScheduleAPIV1ResourceTypeProviderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetScheduleAPIV1ResourceTypeProviderParams() *GetScheduleAPIV1ResourceTypeProviderParams {
	return &GetScheduleAPIV1ResourceTypeProviderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetScheduleAPIV1ResourceTypeProviderParamsWithTimeout creates a new GetScheduleAPIV1ResourceTypeProviderParams object
// with the ability to set a timeout on a request.
func NewGetScheduleAPIV1ResourceTypeProviderParamsWithTimeout(timeout time.Duration) *GetScheduleAPIV1ResourceTypeProviderParams {
	return &GetScheduleAPIV1ResourceTypeProviderParams{
		timeout: timeout,
	}
}

// NewGetScheduleAPIV1ResourceTypeProviderParamsWithContext creates a new GetScheduleAPIV1ResourceTypeProviderParams object
// with the ability to set a context for a request.
func NewGetScheduleAPIV1ResourceTypeProviderParamsWithContext(ctx context.Context) *GetScheduleAPIV1ResourceTypeProviderParams {
	return &GetScheduleAPIV1ResourceTypeProviderParams{
		Context: ctx,
	}
}

// NewGetScheduleAPIV1ResourceTypeProviderParamsWithHTTPClient creates a new GetScheduleAPIV1ResourceTypeProviderParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetScheduleAPIV1ResourceTypeProviderParamsWithHTTPClient(client *http.Client) *GetScheduleAPIV1ResourceTypeProviderParams {
	return &GetScheduleAPIV1ResourceTypeProviderParams{
		HTTPClient: client,
	}
}

/*
GetScheduleAPIV1ResourceTypeProviderParams contains all the parameters to send to the API endpoint

	for the get schedule API v1 resource type provider operation.

	Typically these are written to a http.Request.
*/
type GetScheduleAPIV1ResourceTypeProviderParams struct {

	/* Provider.

	   Provider
	*/
	Provider string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get schedule API v1 resource type provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScheduleAPIV1ResourceTypeProviderParams) WithDefaults() *GetScheduleAPIV1ResourceTypeProviderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get schedule API v1 resource type provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScheduleAPIV1ResourceTypeProviderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get schedule API v1 resource type provider params
func (o *GetScheduleAPIV1ResourceTypeProviderParams) WithTimeout(timeout time.Duration) *GetScheduleAPIV1ResourceTypeProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get schedule API v1 resource type provider params
func (o *GetScheduleAPIV1ResourceTypeProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get schedule API v1 resource type provider params
func (o *GetScheduleAPIV1ResourceTypeProviderParams) WithContext(ctx context.Context) *GetScheduleAPIV1ResourceTypeProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get schedule API v1 resource type provider params
func (o *GetScheduleAPIV1ResourceTypeProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get schedule API v1 resource type provider params
func (o *GetScheduleAPIV1ResourceTypeProviderParams) WithHTTPClient(client *http.Client) *GetScheduleAPIV1ResourceTypeProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get schedule API v1 resource type provider params
func (o *GetScheduleAPIV1ResourceTypeProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProvider adds the provider to the get schedule API v1 resource type provider params
func (o *GetScheduleAPIV1ResourceTypeProviderParams) WithProvider(provider string) *GetScheduleAPIV1ResourceTypeProviderParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the get schedule API v1 resource type provider params
func (o *GetScheduleAPIV1ResourceTypeProviderParams) SetProvider(provider string) {
	o.Provider = provider
}

// WriteToRequest writes these params to a swagger request
func (o *GetScheduleAPIV1ResourceTypeProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param provider
	if err := r.SetPathParam("provider", o.Provider); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}