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

// NewGetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams creates a new GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams() *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams {
	return &GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParamsWithTimeout creates a new GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams object
// with the ability to set a timeout on a request.
func NewGetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParamsWithTimeout(timeout time.Duration) *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams {
	return &GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams{
		timeout: timeout,
	}
}

// NewGetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParamsWithContext creates a new GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams object
// with the ability to set a context for a request.
func NewGetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParamsWithContext(ctx context.Context) *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams {
	return &GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams{
		Context: ctx,
	}
}

// NewGetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParamsWithHTTPClient creates a new GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParamsWithHTTPClient(client *http.Client) *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams {
	return &GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams{
		HTTPClient: client,
	}
}

/*
GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams contains all the parameters to send to the API endpoint

	for the get compliance API v1 benchmarks benchmark ID policies policy ID operation.

	Typically these are written to a http.Request.
*/
type GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams struct {

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

	/* PolicyID.

	   Policy ID
	*/
	PolicyID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get compliance API v1 benchmarks benchmark ID policies policy ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams) WithDefaults() *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get compliance API v1 benchmarks benchmark ID policies policy ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get compliance API v1 benchmarks benchmark ID policies policy ID params
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams) WithTimeout(timeout time.Duration) *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get compliance API v1 benchmarks benchmark ID policies policy ID params
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get compliance API v1 benchmarks benchmark ID policies policy ID params
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams) WithContext(ctx context.Context) *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get compliance API v1 benchmarks benchmark ID policies policy ID params
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get compliance API v1 benchmarks benchmark ID policies policy ID params
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams) WithHTTPClient(client *http.Client) *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get compliance API v1 benchmarks benchmark ID policies policy ID params
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBenchmarkID adds the benchmarkID to the get compliance API v1 benchmarks benchmark ID policies policy ID params
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams) WithBenchmarkID(benchmarkID string) *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams {
	o.SetBenchmarkID(benchmarkID)
	return o
}

// SetBenchmarkID adds the benchmarkId to the get compliance API v1 benchmarks benchmark ID policies policy ID params
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams) SetBenchmarkID(benchmarkID string) {
	o.BenchmarkID = benchmarkID
}

// WithConnectionGroup adds the connectionGroup to the get compliance API v1 benchmarks benchmark ID policies policy ID params
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams) WithConnectionGroup(connectionGroup []string) *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams {
	o.SetConnectionGroup(connectionGroup)
	return o
}

// SetConnectionGroup adds the connectionGroup to the get compliance API v1 benchmarks benchmark ID policies policy ID params
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams) SetConnectionGroup(connectionGroup []string) {
	o.ConnectionGroup = connectionGroup
}

// WithConnectionID adds the connectionID to the get compliance API v1 benchmarks benchmark ID policies policy ID params
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams) WithConnectionID(connectionID []string) *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams {
	o.SetConnectionID(connectionID)
	return o
}

// SetConnectionID adds the connectionId to the get compliance API v1 benchmarks benchmark ID policies policy ID params
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams) SetConnectionID(connectionID []string) {
	o.ConnectionID = connectionID
}

// WithPolicyID adds the policyID to the get compliance API v1 benchmarks benchmark ID policies policy ID params
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams) WithPolicyID(policyID string) *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams {
	o.SetPolicyID(policyID)
	return o
}

// SetPolicyID adds the policyId to the get compliance API v1 benchmarks benchmark ID policies policy ID params
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams) SetPolicyID(policyID string) {
	o.PolicyID = policyID
}

// WriteToRequest writes these params to a swagger request
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param policyId
	if err := r.SetPathParam("policyId", o.PolicyID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyID binds the parameter connectionGroup
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams) bindParamConnectionGroup(formats strfmt.Registry) []string {
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

// bindParamGetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyID binds the parameter connectionId
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesPolicyIDParams) bindParamConnectionID(formats strfmt.Registry) []string {
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
