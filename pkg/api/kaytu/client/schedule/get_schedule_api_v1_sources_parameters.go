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

// NewGetScheduleAPIV1SourcesParams creates a new GetScheduleAPIV1SourcesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetScheduleAPIV1SourcesParams() *GetScheduleAPIV1SourcesParams {
	return &GetScheduleAPIV1SourcesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetScheduleAPIV1SourcesParamsWithTimeout creates a new GetScheduleAPIV1SourcesParams object
// with the ability to set a timeout on a request.
func NewGetScheduleAPIV1SourcesParamsWithTimeout(timeout time.Duration) *GetScheduleAPIV1SourcesParams {
	return &GetScheduleAPIV1SourcesParams{
		timeout: timeout,
	}
}

// NewGetScheduleAPIV1SourcesParamsWithContext creates a new GetScheduleAPIV1SourcesParams object
// with the ability to set a context for a request.
func NewGetScheduleAPIV1SourcesParamsWithContext(ctx context.Context) *GetScheduleAPIV1SourcesParams {
	return &GetScheduleAPIV1SourcesParams{
		Context: ctx,
	}
}

// NewGetScheduleAPIV1SourcesParamsWithHTTPClient creates a new GetScheduleAPIV1SourcesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetScheduleAPIV1SourcesParamsWithHTTPClient(client *http.Client) *GetScheduleAPIV1SourcesParams {
	return &GetScheduleAPIV1SourcesParams{
		HTTPClient: client,
	}
}

/*
GetScheduleAPIV1SourcesParams contains all the parameters to send to the API endpoint

	for the get schedule API v1 sources operation.

	Typically these are written to a http.Request.
*/
type GetScheduleAPIV1SourcesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get schedule API v1 sources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScheduleAPIV1SourcesParams) WithDefaults() *GetScheduleAPIV1SourcesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get schedule API v1 sources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScheduleAPIV1SourcesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get schedule API v1 sources params
func (o *GetScheduleAPIV1SourcesParams) WithTimeout(timeout time.Duration) *GetScheduleAPIV1SourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get schedule API v1 sources params
func (o *GetScheduleAPIV1SourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get schedule API v1 sources params
func (o *GetScheduleAPIV1SourcesParams) WithContext(ctx context.Context) *GetScheduleAPIV1SourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get schedule API v1 sources params
func (o *GetScheduleAPIV1SourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get schedule API v1 sources params
func (o *GetScheduleAPIV1SourcesParams) WithHTTPClient(client *http.Client) *GetScheduleAPIV1SourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get schedule API v1 sources params
func (o *GetScheduleAPIV1SourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetScheduleAPIV1SourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}