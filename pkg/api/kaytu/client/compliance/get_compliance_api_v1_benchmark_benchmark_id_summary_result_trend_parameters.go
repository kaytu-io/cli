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
	"github.com/go-openapi/swag"
)

// NewGetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams creates a new GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams() *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams {
	return &GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParamsWithTimeout creates a new GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams object
// with the ability to set a timeout on a request.
func NewGetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParamsWithTimeout(timeout time.Duration) *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams {
	return &GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams{
		timeout: timeout,
	}
}

// NewGetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParamsWithContext creates a new GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams object
// with the ability to set a context for a request.
func NewGetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParamsWithContext(ctx context.Context) *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams {
	return &GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams{
		Context: ctx,
	}
}

// NewGetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParamsWithHTTPClient creates a new GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParamsWithHTTPClient(client *http.Client) *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams {
	return &GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams{
		HTTPClient: client,
	}
}

/*
GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams contains all the parameters to send to the API endpoint

	for the get compliance API v1 benchmark benchmark ID summary result trend operation.

	Typically these are written to a http.Request.
*/
type GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams struct {

	/* BenchmarkID.

	   Benchmark ID
	*/
	BenchmarkID string

	/* End.

	   End time
	*/
	End int64

	/* Start.

	   Start time
	*/
	Start int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get compliance API v1 benchmark benchmark ID summary result trend params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams) WithDefaults() *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get compliance API v1 benchmark benchmark ID summary result trend params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get compliance API v1 benchmark benchmark ID summary result trend params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams) WithTimeout(timeout time.Duration) *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get compliance API v1 benchmark benchmark ID summary result trend params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get compliance API v1 benchmark benchmark ID summary result trend params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams) WithContext(ctx context.Context) *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get compliance API v1 benchmark benchmark ID summary result trend params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get compliance API v1 benchmark benchmark ID summary result trend params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams) WithHTTPClient(client *http.Client) *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get compliance API v1 benchmark benchmark ID summary result trend params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBenchmarkID adds the benchmarkID to the get compliance API v1 benchmark benchmark ID summary result trend params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams) WithBenchmarkID(benchmarkID string) *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams {
	o.SetBenchmarkID(benchmarkID)
	return o
}

// SetBenchmarkID adds the benchmarkId to the get compliance API v1 benchmark benchmark ID summary result trend params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams) SetBenchmarkID(benchmarkID string) {
	o.BenchmarkID = benchmarkID
}

// WithEnd adds the end to the get compliance API v1 benchmark benchmark ID summary result trend params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams) WithEnd(end int64) *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the get compliance API v1 benchmark benchmark ID summary result trend params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams) SetEnd(end int64) {
	o.End = end
}

// WithStart adds the start to the get compliance API v1 benchmark benchmark ID summary result trend params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams) WithStart(start int64) *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get compliance API v1 benchmark benchmark ID summary result trend params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams) SetStart(start int64) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param benchmark_id
	if err := r.SetPathParam("benchmark_id", o.BenchmarkID); err != nil {
		return err
	}

	// query param end
	qrEnd := o.End
	qEnd := swag.FormatInt64(qrEnd)
	if qEnd != "" {

		if err := r.SetQueryParam("end", qEnd); err != nil {
			return err
		}
	}

	// query param start
	qrStart := o.Start
	qStart := swag.FormatInt64(qrStart)
	if qStart != "" {

		if err := r.SetQueryParam("start", qStart); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
