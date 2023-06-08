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

// NewGetScheduleAPIV1DescribeSourceJobsPendingParams creates a new GetScheduleAPIV1DescribeSourceJobsPendingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetScheduleAPIV1DescribeSourceJobsPendingParams() *GetScheduleAPIV1DescribeSourceJobsPendingParams {
	return &GetScheduleAPIV1DescribeSourceJobsPendingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetScheduleAPIV1DescribeSourceJobsPendingParamsWithTimeout creates a new GetScheduleAPIV1DescribeSourceJobsPendingParams object
// with the ability to set a timeout on a request.
func NewGetScheduleAPIV1DescribeSourceJobsPendingParamsWithTimeout(timeout time.Duration) *GetScheduleAPIV1DescribeSourceJobsPendingParams {
	return &GetScheduleAPIV1DescribeSourceJobsPendingParams{
		timeout: timeout,
	}
}

// NewGetScheduleAPIV1DescribeSourceJobsPendingParamsWithContext creates a new GetScheduleAPIV1DescribeSourceJobsPendingParams object
// with the ability to set a context for a request.
func NewGetScheduleAPIV1DescribeSourceJobsPendingParamsWithContext(ctx context.Context) *GetScheduleAPIV1DescribeSourceJobsPendingParams {
	return &GetScheduleAPIV1DescribeSourceJobsPendingParams{
		Context: ctx,
	}
}

// NewGetScheduleAPIV1DescribeSourceJobsPendingParamsWithHTTPClient creates a new GetScheduleAPIV1DescribeSourceJobsPendingParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetScheduleAPIV1DescribeSourceJobsPendingParamsWithHTTPClient(client *http.Client) *GetScheduleAPIV1DescribeSourceJobsPendingParams {
	return &GetScheduleAPIV1DescribeSourceJobsPendingParams{
		HTTPClient: client,
	}
}

/*
GetScheduleAPIV1DescribeSourceJobsPendingParams contains all the parameters to send to the API endpoint

	for the get schedule API v1 describe source jobs pending operation.

	Typically these are written to a http.Request.
*/
type GetScheduleAPIV1DescribeSourceJobsPendingParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get schedule API v1 describe source jobs pending params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScheduleAPIV1DescribeSourceJobsPendingParams) WithDefaults() *GetScheduleAPIV1DescribeSourceJobsPendingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get schedule API v1 describe source jobs pending params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScheduleAPIV1DescribeSourceJobsPendingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get schedule API v1 describe source jobs pending params
func (o *GetScheduleAPIV1DescribeSourceJobsPendingParams) WithTimeout(timeout time.Duration) *GetScheduleAPIV1DescribeSourceJobsPendingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get schedule API v1 describe source jobs pending params
func (o *GetScheduleAPIV1DescribeSourceJobsPendingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get schedule API v1 describe source jobs pending params
func (o *GetScheduleAPIV1DescribeSourceJobsPendingParams) WithContext(ctx context.Context) *GetScheduleAPIV1DescribeSourceJobsPendingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get schedule API v1 describe source jobs pending params
func (o *GetScheduleAPIV1DescribeSourceJobsPendingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get schedule API v1 describe source jobs pending params
func (o *GetScheduleAPIV1DescribeSourceJobsPendingParams) WithHTTPClient(client *http.Client) *GetScheduleAPIV1DescribeSourceJobsPendingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get schedule API v1 describe source jobs pending params
func (o *GetScheduleAPIV1DescribeSourceJobsPendingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetScheduleAPIV1DescribeSourceJobsPendingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
