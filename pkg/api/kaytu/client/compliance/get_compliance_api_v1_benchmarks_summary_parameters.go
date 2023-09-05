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

// NewGetComplianceAPIV1BenchmarksSummaryParams creates a new GetComplianceAPIV1BenchmarksSummaryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetComplianceAPIV1BenchmarksSummaryParams() *GetComplianceAPIV1BenchmarksSummaryParams {
	return &GetComplianceAPIV1BenchmarksSummaryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetComplianceAPIV1BenchmarksSummaryParamsWithTimeout creates a new GetComplianceAPIV1BenchmarksSummaryParams object
// with the ability to set a timeout on a request.
func NewGetComplianceAPIV1BenchmarksSummaryParamsWithTimeout(timeout time.Duration) *GetComplianceAPIV1BenchmarksSummaryParams {
	return &GetComplianceAPIV1BenchmarksSummaryParams{
		timeout: timeout,
	}
}

// NewGetComplianceAPIV1BenchmarksSummaryParamsWithContext creates a new GetComplianceAPIV1BenchmarksSummaryParams object
// with the ability to set a context for a request.
func NewGetComplianceAPIV1BenchmarksSummaryParamsWithContext(ctx context.Context) *GetComplianceAPIV1BenchmarksSummaryParams {
	return &GetComplianceAPIV1BenchmarksSummaryParams{
		Context: ctx,
	}
}

// NewGetComplianceAPIV1BenchmarksSummaryParamsWithHTTPClient creates a new GetComplianceAPIV1BenchmarksSummaryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetComplianceAPIV1BenchmarksSummaryParamsWithHTTPClient(client *http.Client) *GetComplianceAPIV1BenchmarksSummaryParams {
	return &GetComplianceAPIV1BenchmarksSummaryParams{
		HTTPClient: client,
	}
}

/*
GetComplianceAPIV1BenchmarksSummaryParams contains all the parameters to send to the API endpoint

	for the get compliance API v1 benchmarks summary operation.

	Typically these are written to a http.Request.
*/
type GetComplianceAPIV1BenchmarksSummaryParams struct {

	/* ConnectionGroup.

	   Connection groups to filter by
	*/
	ConnectionGroup []string

	/* ConnectionID.

	   Connection IDs to filter by
	*/
	ConnectionID []string

	/* Connector.

	   Connector type to filter by
	*/
	Connector []string

	/* TimeAt.

	   timestamp for values in epoch seconds
	*/
	TimeAt *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get compliance API v1 benchmarks summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComplianceAPIV1BenchmarksSummaryParams) WithDefaults() *GetComplianceAPIV1BenchmarksSummaryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get compliance API v1 benchmarks summary params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComplianceAPIV1BenchmarksSummaryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get compliance API v1 benchmarks summary params
func (o *GetComplianceAPIV1BenchmarksSummaryParams) WithTimeout(timeout time.Duration) *GetComplianceAPIV1BenchmarksSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get compliance API v1 benchmarks summary params
func (o *GetComplianceAPIV1BenchmarksSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get compliance API v1 benchmarks summary params
func (o *GetComplianceAPIV1BenchmarksSummaryParams) WithContext(ctx context.Context) *GetComplianceAPIV1BenchmarksSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get compliance API v1 benchmarks summary params
func (o *GetComplianceAPIV1BenchmarksSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get compliance API v1 benchmarks summary params
func (o *GetComplianceAPIV1BenchmarksSummaryParams) WithHTTPClient(client *http.Client) *GetComplianceAPIV1BenchmarksSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get compliance API v1 benchmarks summary params
func (o *GetComplianceAPIV1BenchmarksSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionGroup adds the connectionGroup to the get compliance API v1 benchmarks summary params
func (o *GetComplianceAPIV1BenchmarksSummaryParams) WithConnectionGroup(connectionGroup []string) *GetComplianceAPIV1BenchmarksSummaryParams {
	o.SetConnectionGroup(connectionGroup)
	return o
}

// SetConnectionGroup adds the connectionGroup to the get compliance API v1 benchmarks summary params
func (o *GetComplianceAPIV1BenchmarksSummaryParams) SetConnectionGroup(connectionGroup []string) {
	o.ConnectionGroup = connectionGroup
}

// WithConnectionID adds the connectionID to the get compliance API v1 benchmarks summary params
func (o *GetComplianceAPIV1BenchmarksSummaryParams) WithConnectionID(connectionID []string) *GetComplianceAPIV1BenchmarksSummaryParams {
	o.SetConnectionID(connectionID)
	return o
}

// SetConnectionID adds the connectionId to the get compliance API v1 benchmarks summary params
func (o *GetComplianceAPIV1BenchmarksSummaryParams) SetConnectionID(connectionID []string) {
	o.ConnectionID = connectionID
}

// WithConnector adds the connector to the get compliance API v1 benchmarks summary params
func (o *GetComplianceAPIV1BenchmarksSummaryParams) WithConnector(connector []string) *GetComplianceAPIV1BenchmarksSummaryParams {
	o.SetConnector(connector)
	return o
}

// SetConnector adds the connector to the get compliance API v1 benchmarks summary params
func (o *GetComplianceAPIV1BenchmarksSummaryParams) SetConnector(connector []string) {
	o.Connector = connector
}

// WithTimeAt adds the timeAt to the get compliance API v1 benchmarks summary params
func (o *GetComplianceAPIV1BenchmarksSummaryParams) WithTimeAt(timeAt *int64) *GetComplianceAPIV1BenchmarksSummaryParams {
	o.SetTimeAt(timeAt)
	return o
}

// SetTimeAt adds the timeAt to the get compliance API v1 benchmarks summary params
func (o *GetComplianceAPIV1BenchmarksSummaryParams) SetTimeAt(timeAt *int64) {
	o.TimeAt = timeAt
}

// WriteToRequest writes these params to a swagger request
func (o *GetComplianceAPIV1BenchmarksSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.Connector != nil {

		// binding items for connector
		joinedConnector := o.bindParamConnector(reg)

		// query array param connector
		if err := r.SetQueryParam("connector", joinedConnector...); err != nil {
			return err
		}
	}

	if o.TimeAt != nil {

		// query param timeAt
		var qrTimeAt int64

		if o.TimeAt != nil {
			qrTimeAt = *o.TimeAt
		}
		qTimeAt := swag.FormatInt64(qrTimeAt)
		if qTimeAt != "" {

			if err := r.SetQueryParam("timeAt", qTimeAt); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetComplianceAPIV1BenchmarksSummary binds the parameter connectionGroup
func (o *GetComplianceAPIV1BenchmarksSummaryParams) bindParamConnectionGroup(formats strfmt.Registry) []string {
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

// bindParamGetComplianceAPIV1BenchmarksSummary binds the parameter connectionId
func (o *GetComplianceAPIV1BenchmarksSummaryParams) bindParamConnectionID(formats strfmt.Registry) []string {
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

// bindParamGetComplianceAPIV1BenchmarksSummary binds the parameter connector
func (o *GetComplianceAPIV1BenchmarksSummaryParams) bindParamConnector(formats strfmt.Registry) []string {
	connectorIR := o.Connector

	var connectorIC []string
	for _, connectorIIR := range connectorIR { // explode []string

		connectorIIV := connectorIIR // string as string
		connectorIC = append(connectorIC, connectorIIV)
	}

	// items.CollectionFormat: "csv"
	connectorIS := swag.JoinByFormat(connectorIC, "csv")

	return connectorIS
}
