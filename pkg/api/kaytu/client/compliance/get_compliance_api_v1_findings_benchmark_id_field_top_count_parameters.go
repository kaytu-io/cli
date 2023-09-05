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

// NewGetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams creates a new GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams() *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams {
	return &GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParamsWithTimeout creates a new GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams object
// with the ability to set a timeout on a request.
func NewGetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParamsWithTimeout(timeout time.Duration) *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams {
	return &GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams{
		timeout: timeout,
	}
}

// NewGetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParamsWithContext creates a new GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams object
// with the ability to set a context for a request.
func NewGetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParamsWithContext(ctx context.Context) *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams {
	return &GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams{
		Context: ctx,
	}
}

// NewGetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParamsWithHTTPClient creates a new GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParamsWithHTTPClient(client *http.Client) *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams {
	return &GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams{
		HTTPClient: client,
	}
}

/*
GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams contains all the parameters to send to the API endpoint

	for the get compliance API v1 findings benchmark ID field top count operation.

	Typically these are written to a http.Request.
*/
type GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams struct {

	/* BenchmarkID.

	   BenchmarkID
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

	/* Connector.

	   Connector type to filter by
	*/
	Connector []string

	/* Count.

	   Count
	*/
	Count int64

	/* Field.

	   Field
	*/
	Field string

	/* Severities.

	   Severities to filter by
	*/
	Severities []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get compliance API v1 findings benchmark ID field top count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams) WithDefaults() *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get compliance API v1 findings benchmark ID field top count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get compliance API v1 findings benchmark ID field top count params
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams) WithTimeout(timeout time.Duration) *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get compliance API v1 findings benchmark ID field top count params
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get compliance API v1 findings benchmark ID field top count params
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams) WithContext(ctx context.Context) *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get compliance API v1 findings benchmark ID field top count params
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get compliance API v1 findings benchmark ID field top count params
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams) WithHTTPClient(client *http.Client) *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get compliance API v1 findings benchmark ID field top count params
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBenchmarkID adds the benchmarkID to the get compliance API v1 findings benchmark ID field top count params
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams) WithBenchmarkID(benchmarkID string) *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams {
	o.SetBenchmarkID(benchmarkID)
	return o
}

// SetBenchmarkID adds the benchmarkId to the get compliance API v1 findings benchmark ID field top count params
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams) SetBenchmarkID(benchmarkID string) {
	o.BenchmarkID = benchmarkID
}

// WithConnectionGroup adds the connectionGroup to the get compliance API v1 findings benchmark ID field top count params
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams) WithConnectionGroup(connectionGroup []string) *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams {
	o.SetConnectionGroup(connectionGroup)
	return o
}

// SetConnectionGroup adds the connectionGroup to the get compliance API v1 findings benchmark ID field top count params
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams) SetConnectionGroup(connectionGroup []string) {
	o.ConnectionGroup = connectionGroup
}

// WithConnectionID adds the connectionID to the get compliance API v1 findings benchmark ID field top count params
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams) WithConnectionID(connectionID []string) *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams {
	o.SetConnectionID(connectionID)
	return o
}

// SetConnectionID adds the connectionId to the get compliance API v1 findings benchmark ID field top count params
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams) SetConnectionID(connectionID []string) {
	o.ConnectionID = connectionID
}

// WithConnector adds the connector to the get compliance API v1 findings benchmark ID field top count params
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams) WithConnector(connector []string) *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams {
	o.SetConnector(connector)
	return o
}

// SetConnector adds the connector to the get compliance API v1 findings benchmark ID field top count params
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams) SetConnector(connector []string) {
	o.Connector = connector
}

// WithCount adds the count to the get compliance API v1 findings benchmark ID field top count params
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams) WithCount(count int64) *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get compliance API v1 findings benchmark ID field top count params
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams) SetCount(count int64) {
	o.Count = count
}

// WithField adds the field to the get compliance API v1 findings benchmark ID field top count params
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams) WithField(field string) *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams {
	o.SetField(field)
	return o
}

// SetField adds the field to the get compliance API v1 findings benchmark ID field top count params
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams) SetField(field string) {
	o.Field = field
}

// WithSeverities adds the severities to the get compliance API v1 findings benchmark ID field top count params
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams) WithSeverities(severities []string) *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams {
	o.SetSeverities(severities)
	return o
}

// SetSeverities adds the severities to the get compliance API v1 findings benchmark ID field top count params
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams) SetSeverities(severities []string) {
	o.Severities = severities
}

// WriteToRequest writes these params to a swagger request
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param benchmarkId
	if err := r.SetPathParam("benchmarkId", o.BenchmarkID); err != nil {
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

	if o.Connector != nil {

		// binding items for connector
		joinedConnector := o.bindParamConnector(reg)

		// query array param connector
		if err := r.SetQueryParam("connector", joinedConnector...); err != nil {
			return err
		}
	}

	// path param count
	if err := r.SetPathParam("count", swag.FormatInt64(o.Count)); err != nil {
		return err
	}

	// path param field
	if err := r.SetPathParam("field", o.Field); err != nil {
		return err
	}

	if o.Severities != nil {

		// binding items for severities
		joinedSeverities := o.bindParamSeverities(reg)

		// query array param severities
		if err := r.SetQueryParam("severities", joinedSeverities...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetComplianceAPIV1FindingsBenchmarkIDFieldTopCount binds the parameter connectionGroup
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams) bindParamConnectionGroup(formats strfmt.Registry) []string {
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

// bindParamGetComplianceAPIV1FindingsBenchmarkIDFieldTopCount binds the parameter connectionId
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams) bindParamConnectionID(formats strfmt.Registry) []string {
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

// bindParamGetComplianceAPIV1FindingsBenchmarkIDFieldTopCount binds the parameter connector
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams) bindParamConnector(formats strfmt.Registry) []string {
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

// bindParamGetComplianceAPIV1FindingsBenchmarkIDFieldTopCount binds the parameter severities
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams) bindParamSeverities(formats strfmt.Registry) []string {
	severitiesIR := o.Severities

	var severitiesIC []string
	for _, severitiesIIR := range severitiesIR { // explode []string

		severitiesIIV := severitiesIIR // string as string
		severitiesIC = append(severitiesIC, severitiesIIV)
	}

	// items.CollectionFormat: "csv"
	severitiesIS := swag.JoinByFormat(severitiesIC, "csv")

	return severitiesIS
}
