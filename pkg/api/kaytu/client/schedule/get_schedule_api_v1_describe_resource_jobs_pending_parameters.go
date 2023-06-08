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

// NewGetScheduleAPIV1DescribeResourceJobsPendingParams creates a new GetScheduleAPIV1DescribeResourceJobsPendingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetScheduleAPIV1DescribeResourceJobsPendingParams() *GetScheduleAPIV1DescribeResourceJobsPendingParams {
	return &GetScheduleAPIV1DescribeResourceJobsPendingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetScheduleAPIV1DescribeResourceJobsPendingParamsWithTimeout creates a new GetScheduleAPIV1DescribeResourceJobsPendingParams object
// with the ability to set a timeout on a request.
func NewGetScheduleAPIV1DescribeResourceJobsPendingParamsWithTimeout(timeout time.Duration) *GetScheduleAPIV1DescribeResourceJobsPendingParams {
	return &GetScheduleAPIV1DescribeResourceJobsPendingParams{
		timeout: timeout,
	}
}

// NewGetScheduleAPIV1DescribeResourceJobsPendingParamsWithContext creates a new GetScheduleAPIV1DescribeResourceJobsPendingParams object
// with the ability to set a context for a request.
func NewGetScheduleAPIV1DescribeResourceJobsPendingParamsWithContext(ctx context.Context) *GetScheduleAPIV1DescribeResourceJobsPendingParams {
	return &GetScheduleAPIV1DescribeResourceJobsPendingParams{
		Context: ctx,
	}
}

// NewGetScheduleAPIV1DescribeResourceJobsPendingParamsWithHTTPClient creates a new GetScheduleAPIV1DescribeResourceJobsPendingParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetScheduleAPIV1DescribeResourceJobsPendingParamsWithHTTPClient(client *http.Client) *GetScheduleAPIV1DescribeResourceJobsPendingParams {
	return &GetScheduleAPIV1DescribeResourceJobsPendingParams{
		HTTPClient: client,
	}
}

/*
GetScheduleAPIV1DescribeResourceJobsPendingParams contains all the parameters to send to the API endpoint

	for the get schedule API v1 describe resource jobs pending operation.

	Typically these are written to a http.Request.
*/
type GetScheduleAPIV1DescribeResourceJobsPendingParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get schedule API v1 describe resource jobs pending params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScheduleAPIV1DescribeResourceJobsPendingParams) WithDefaults() *GetScheduleAPIV1DescribeResourceJobsPendingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get schedule API v1 describe resource jobs pending params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScheduleAPIV1DescribeResourceJobsPendingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get schedule API v1 describe resource jobs pending params
func (o *GetScheduleAPIV1DescribeResourceJobsPendingParams) WithTimeout(timeout time.Duration) *GetScheduleAPIV1DescribeResourceJobsPendingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get schedule API v1 describe resource jobs pending params
func (o *GetScheduleAPIV1DescribeResourceJobsPendingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get schedule API v1 describe resource jobs pending params
func (o *GetScheduleAPIV1DescribeResourceJobsPendingParams) WithContext(ctx context.Context) *GetScheduleAPIV1DescribeResourceJobsPendingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get schedule API v1 describe resource jobs pending params
func (o *GetScheduleAPIV1DescribeResourceJobsPendingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get schedule API v1 describe resource jobs pending params
func (o *GetScheduleAPIV1DescribeResourceJobsPendingParams) WithHTTPClient(client *http.Client) *GetScheduleAPIV1DescribeResourceJobsPendingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get schedule API v1 describe resource jobs pending params
func (o *GetScheduleAPIV1DescribeResourceJobsPendingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetScheduleAPIV1DescribeResourceJobsPendingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
