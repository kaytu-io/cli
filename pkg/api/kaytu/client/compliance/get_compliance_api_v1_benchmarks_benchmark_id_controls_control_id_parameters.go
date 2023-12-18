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

// NewGetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams creates a new GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams() *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams {
	return &GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParamsWithTimeout creates a new GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams object
// with the ability to set a timeout on a request.
func NewGetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParamsWithTimeout(timeout time.Duration) *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams {
	return &GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams{
		timeout: timeout,
	}
}

// NewGetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParamsWithContext creates a new GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams object
// with the ability to set a context for a request.
func NewGetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParamsWithContext(ctx context.Context) *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams {
	return &GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams{
		Context: ctx,
	}
}

// NewGetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParamsWithHTTPClient creates a new GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParamsWithHTTPClient(client *http.Client) *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams {
	return &GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams{
		HTTPClient: client,
	}
}

/*
GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams contains all the parameters to send to the API endpoint

	for the get compliance API v1 benchmarks benchmark ID controls control ID operation.

	Typically these are written to a http.Request.
*/
type GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams struct {

	/* BenchmarkID.

	   Benchmark ID
	*/
	BenchmarkID string

	/* ConnectionGroup.

	   Connection groups to filter by
	*/
	ConnectionGroup []string

	/* ConnectionID.

	   Connection IDs to filter by
	*/
	ConnectionID []string

	/* ControlID.

	   Control ID
	*/
	ControlID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get compliance API v1 benchmarks benchmark ID controls control ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams) WithDefaults() *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get compliance API v1 benchmarks benchmark ID controls control ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get compliance API v1 benchmarks benchmark ID controls control ID params
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams) WithTimeout(timeout time.Duration) *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get compliance API v1 benchmarks benchmark ID controls control ID params
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get compliance API v1 benchmarks benchmark ID controls control ID params
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams) WithContext(ctx context.Context) *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get compliance API v1 benchmarks benchmark ID controls control ID params
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get compliance API v1 benchmarks benchmark ID controls control ID params
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams) WithHTTPClient(client *http.Client) *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get compliance API v1 benchmarks benchmark ID controls control ID params
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBenchmarkID adds the benchmarkID to the get compliance API v1 benchmarks benchmark ID controls control ID params
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams) WithBenchmarkID(benchmarkID string) *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams {
	o.SetBenchmarkID(benchmarkID)
	return o
}

// SetBenchmarkID adds the benchmarkId to the get compliance API v1 benchmarks benchmark ID controls control ID params
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams) SetBenchmarkID(benchmarkID string) {
	o.BenchmarkID = benchmarkID
}

// WithConnectionGroup adds the connectionGroup to the get compliance API v1 benchmarks benchmark ID controls control ID params
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams) WithConnectionGroup(connectionGroup []string) *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams {
	o.SetConnectionGroup(connectionGroup)
	return o
}

// SetConnectionGroup adds the connectionGroup to the get compliance API v1 benchmarks benchmark ID controls control ID params
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams) SetConnectionGroup(connectionGroup []string) {
	o.ConnectionGroup = connectionGroup
}

// WithConnectionID adds the connectionID to the get compliance API v1 benchmarks benchmark ID controls control ID params
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams) WithConnectionID(connectionID []string) *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams {
	o.SetConnectionID(connectionID)
	return o
}

// SetConnectionID adds the connectionId to the get compliance API v1 benchmarks benchmark ID controls control ID params
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams) SetConnectionID(connectionID []string) {
	o.ConnectionID = connectionID
}

// WithControlID adds the controlID to the get compliance API v1 benchmarks benchmark ID controls control ID params
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams) WithControlID(controlID string) *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams {
	o.SetControlID(controlID)
	return o
}

// SetControlID adds the controlId to the get compliance API v1 benchmarks benchmark ID controls control ID params
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams) SetControlID(controlID string) {
	o.ControlID = controlID
}

// WriteToRequest writes these params to a swagger request
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param benchmark_id
	if err := r.SetPathParam("benchmark_id", o.BenchmarkID); err != nil {
		return err
	}

	if o.ConnectionGroup != nil {

		// binding items for connectionGroup
		joinedConnectionGroup := o.bindParamConnectionGroup(reg)

		// query array param connectionGroup
		if err := r.SetQueryParam("connectionGroup", joinedConnectionGroup...); err != nil {
			return err
		}
	}

	if o.ConnectionID != nil {

		// binding items for connectionId
		joinedConnectionID := o.bindParamConnectionID(reg)

		// query array param connectionId
		if err := r.SetQueryParam("connectionId", joinedConnectionID...); err != nil {
			return err
		}
	}

	// path param controlId
	if err := r.SetPathParam("controlId", o.ControlID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetComplianceAPIV1BenchmarksBenchmarkIDControlsControlID binds the parameter connectionGroup
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams) bindParamConnectionGroup(formats strfmt.Registry) []string {
	connectionGroupIR := o.ConnectionGroup

	var connectionGroupIC []string
	for _, connectionGroupIIR := range connectionGroupIR { // explode []string

		connectionGroupIIV := connectionGroupIIR // string as string
		connectionGroupIC = append(connectionGroupIC, connectionGroupIIV)
	}

	// items.CollectionFormat: "csv"
	connectionGroupIS := swag.JoinByFormat(connectionGroupIC, "csv")

	return connectionGroupIS
}

// bindParamGetComplianceAPIV1BenchmarksBenchmarkIDControlsControlID binds the parameter connectionId
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDControlsControlIDParams) bindParamConnectionID(formats strfmt.Registry) []string {
	connectionIDIR := o.ConnectionID

	var connectionIDIC []string
	for _, connectionIDIIR := range connectionIDIR { // explode []string

		connectionIDIIV := connectionIDIIR // string as string
		connectionIDIC = append(connectionIDIC, connectionIDIIV)
	}

	// items.CollectionFormat: "csv"
	connectionIDIS := swag.JoinByFormat(connectionIDIC, "csv")

	return connectionIDIS
}
