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

// NewGetComplianceAPIV1BenchmarkBenchmarkIDTreeParams creates a new GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetComplianceAPIV1BenchmarkBenchmarkIDTreeParams() *GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams {
	return &GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetComplianceAPIV1BenchmarkBenchmarkIDTreeParamsWithTimeout creates a new GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams object
// with the ability to set a timeout on a request.
func NewGetComplianceAPIV1BenchmarkBenchmarkIDTreeParamsWithTimeout(timeout time.Duration) *GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams {
	return &GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams{
		timeout: timeout,
	}
}

// NewGetComplianceAPIV1BenchmarkBenchmarkIDTreeParamsWithContext creates a new GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams object
// with the ability to set a context for a request.
func NewGetComplianceAPIV1BenchmarkBenchmarkIDTreeParamsWithContext(ctx context.Context) *GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams {
	return &GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams{
		Context: ctx,
	}
}

// NewGetComplianceAPIV1BenchmarkBenchmarkIDTreeParamsWithHTTPClient creates a new GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetComplianceAPIV1BenchmarkBenchmarkIDTreeParamsWithHTTPClient(client *http.Client) *GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams {
	return &GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams{
		HTTPClient: client,
	}
}

/*
GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams contains all the parameters to send to the API endpoint

	for the get compliance API v1 benchmark benchmark ID tree operation.

	Typically these are written to a http.Request.
*/
type GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams struct {

	/* BenchmarkID.

	   Benchmark ID
	*/
	BenchmarkID string

	/* Status.

	   Status
	*/
	Status []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get compliance API v1 benchmark benchmark ID tree params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams) WithDefaults() *GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get compliance API v1 benchmark benchmark ID tree params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get compliance API v1 benchmark benchmark ID tree params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams) WithTimeout(timeout time.Duration) *GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get compliance API v1 benchmark benchmark ID tree params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get compliance API v1 benchmark benchmark ID tree params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams) WithContext(ctx context.Context) *GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get compliance API v1 benchmark benchmark ID tree params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get compliance API v1 benchmark benchmark ID tree params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams) WithHTTPClient(client *http.Client) *GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get compliance API v1 benchmark benchmark ID tree params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBenchmarkID adds the benchmarkID to the get compliance API v1 benchmark benchmark ID tree params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams) WithBenchmarkID(benchmarkID string) *GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams {
	o.SetBenchmarkID(benchmarkID)
	return o
}

// SetBenchmarkID adds the benchmarkId to the get compliance API v1 benchmark benchmark ID tree params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams) SetBenchmarkID(benchmarkID string) {
	o.BenchmarkID = benchmarkID
}

// WithStatus adds the status to the get compliance API v1 benchmark benchmark ID tree params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams) WithStatus(status []string) *GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get compliance API v1 benchmark benchmark ID tree params
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams) SetStatus(status []string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param benchmark_id
	if err := r.SetPathParam("benchmark_id", o.BenchmarkID); err != nil {
		return err
	}

	if o.Status != nil {

		// binding items for status
		joinedStatus := o.bindParamStatus(reg)

		// query array param status
		if err := r.SetQueryParam("status", joinedStatus...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetComplianceAPIV1BenchmarkBenchmarkIDTree binds the parameter status
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDTreeParams) bindParamStatus(formats strfmt.Registry) []string {
	statusIR := o.Status

	var statusIC []string
	for _, statusIIR := range statusIR { // explode []string

		statusIIV := statusIIR // string as string
		statusIC = append(statusIC, statusIIV)
	}

	// items.CollectionFormat: "csv"
	statusIS := swag.JoinByFormat(statusIC, "csv")

	return statusIS
}
