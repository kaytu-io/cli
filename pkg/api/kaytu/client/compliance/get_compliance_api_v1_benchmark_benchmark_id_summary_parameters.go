// Code generated by go-swagger; DO NOT EDIT.

package compliance

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

// NewGetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams creates a new GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams() *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams {
	return &GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetComplianceAPIV1BenchmarkBenchmarkIDSummaryParamsWithTimeout creates a new GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams object
// with the ability to set a timeout on a request.
func NewGetComplianceAPIV1BenchmarkBenchmarkIDSummaryParamsWithTimeout(timeout time.Duration) *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams {
	return &GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams{
		timeout: timeout,
	}
}

// NewGetComplianceAPIV1BenchmarkBenchmarkIDSummaryParamsWithContext creates a new GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams object
// with the ability to set a context for a request.
func NewGetComplianceAPIV1BenchmarkBenchmarkIDSummaryParamsWithContext(ctx context.Context) *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams {
	return &GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams{
		Context: ctx,
	}
}

// NewGetComplianceAPIV1BenchmarkBenchmarkIDSummaryParamsWithHTTPClient creates a new GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetComplianceAPIV1BenchmarkBenchmarkIDSummaryParamsWithHTTPClient(client *http.Client) *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams {
	return &GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams{
		HTTPClient: client,
	}
}

/*
GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams contains all the parameters to send to the API endpoint

	for the get compliance API v1 benchmark benchmark ID summary operation.

	Typically these are written to a http.Request.
*/
type GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams struct {

	/* BenchmarkID.

	   BenchmarkID
	*/
	BenchmarkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get compliance API v1 benchmark benchmark ID summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams) WithDefaults() *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get compliance API v1 benchmark benchmark ID summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get compliance API v1 benchmark benchmark ID summary params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams) WithTimeout(timeout time.Duration) *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get compliance API v1 benchmark benchmark ID summary params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get compliance API v1 benchmark benchmark ID summary params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams) WithContext(ctx context.Context) *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get compliance API v1 benchmark benchmark ID summary params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get compliance API v1 benchmark benchmark ID summary params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams) WithHTTPClient(client *http.Client) *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get compliance API v1 benchmark benchmark ID summary params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBenchmarkID adds the benchmarkID to the get compliance API v1 benchmark benchmark ID summary params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams) WithBenchmarkID(benchmarkID string) *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams {
	o.SetBenchmarkID(benchmarkID)
	return o
}

// SetBenchmarkID adds the benchmarkId to the get compliance API v1 benchmark benchmark ID summary params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams) SetBenchmarkID(benchmarkID string) {
	o.BenchmarkID = benchmarkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param benchmark_id
	if err := r.SetPathParam("benchmark_id", o.BenchmarkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}