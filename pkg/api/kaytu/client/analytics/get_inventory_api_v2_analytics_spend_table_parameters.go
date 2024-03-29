// Code generated by go-swagger; DO NOT EDIT.

package analytics

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

// NewGetInventoryAPIV2AnalyticsSpendTableParams creates a new GetInventoryAPIV2AnalyticsSpendTableParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInventoryAPIV2AnalyticsSpendTableParams() *GetInventoryAPIV2AnalyticsSpendTableParams {
	return &GetInventoryAPIV2AnalyticsSpendTableParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInventoryAPIV2AnalyticsSpendTableParamsWithTimeout creates a new GetInventoryAPIV2AnalyticsSpendTableParams object
// with the ability to set a timeout on a request.
func NewGetInventoryAPIV2AnalyticsSpendTableParamsWithTimeout(timeout time.Duration) *GetInventoryAPIV2AnalyticsSpendTableParams {
	return &GetInventoryAPIV2AnalyticsSpendTableParams{
		timeout: timeout,
	}
}

// NewGetInventoryAPIV2AnalyticsSpendTableParamsWithContext creates a new GetInventoryAPIV2AnalyticsSpendTableParams object
// with the ability to set a context for a request.
func NewGetInventoryAPIV2AnalyticsSpendTableParamsWithContext(ctx context.Context) *GetInventoryAPIV2AnalyticsSpendTableParams {
	return &GetInventoryAPIV2AnalyticsSpendTableParams{
		Context: ctx,
	}
}

// NewGetInventoryAPIV2AnalyticsSpendTableParamsWithHTTPClient creates a new GetInventoryAPIV2AnalyticsSpendTableParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInventoryAPIV2AnalyticsSpendTableParamsWithHTTPClient(client *http.Client) *GetInventoryAPIV2AnalyticsSpendTableParams {
	return &GetInventoryAPIV2AnalyticsSpendTableParams{
		HTTPClient: client,
	}
}

/*
GetInventoryAPIV2AnalyticsSpendTableParams contains all the parameters to send to the API endpoint

	for the get inventory API v2 analytics spend table operation.

	Typically these are written to a http.Request.
*/
type GetInventoryAPIV2AnalyticsSpendTableParams struct {

	/* ConnectionGroup.

	   Connection group to filter by - mutually exclusive with connectionId
	*/
	ConnectionGroup []string

	/* ConnectionID.

	   Connection IDs to filter by - mutually exclusive with connectionGroup
	*/
	ConnectionID []string

	/* Connector.

	   Connector
	*/
	Connector []string

	/* Dimension.

	   Dimension of the table, default is metric
	*/
	Dimension *string

	/* EndTime.

	   timestamp for end in epoch seconds
	*/
	EndTime *int64

	/* Granularity.

	   Granularity of the table, default is daily
	*/
	Granularity *string

	/* MetricIds.

	   Metrics IDs
	*/
	MetricIds []string

	/* StartTime.

	   timestamp for start in epoch seconds
	*/
	StartTime *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get inventory API v2 analytics spend table params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) WithDefaults() *GetInventoryAPIV2AnalyticsSpendTableParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get inventory API v2 analytics spend table params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get inventory API v2 analytics spend table params
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) WithTimeout(timeout time.Duration) *GetInventoryAPIV2AnalyticsSpendTableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get inventory API v2 analytics spend table params
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get inventory API v2 analytics spend table params
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) WithContext(ctx context.Context) *GetInventoryAPIV2AnalyticsSpendTableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get inventory API v2 analytics spend table params
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get inventory API v2 analytics spend table params
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) WithHTTPClient(client *http.Client) *GetInventoryAPIV2AnalyticsSpendTableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get inventory API v2 analytics spend table params
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionGroup adds the connectionGroup to the get inventory API v2 analytics spend table params
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) WithConnectionGroup(connectionGroup []string) *GetInventoryAPIV2AnalyticsSpendTableParams {
	o.SetConnectionGroup(connectionGroup)
	return o
}

// SetConnectionGroup adds the connectionGroup to the get inventory API v2 analytics spend table params
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) SetConnectionGroup(connectionGroup []string) {
	o.ConnectionGroup = connectionGroup
}

// WithConnectionID adds the connectionID to the get inventory API v2 analytics spend table params
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) WithConnectionID(connectionID []string) *GetInventoryAPIV2AnalyticsSpendTableParams {
	o.SetConnectionID(connectionID)
	return o
}

// SetConnectionID adds the connectionId to the get inventory API v2 analytics spend table params
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) SetConnectionID(connectionID []string) {
	o.ConnectionID = connectionID
}

// WithConnector adds the connector to the get inventory API v2 analytics spend table params
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) WithConnector(connector []string) *GetInventoryAPIV2AnalyticsSpendTableParams {
	o.SetConnector(connector)
	return o
}

// SetConnector adds the connector to the get inventory API v2 analytics spend table params
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) SetConnector(connector []string) {
	o.Connector = connector
}

// WithDimension adds the dimension to the get inventory API v2 analytics spend table params
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) WithDimension(dimension *string) *GetInventoryAPIV2AnalyticsSpendTableParams {
	o.SetDimension(dimension)
	return o
}

// SetDimension adds the dimension to the get inventory API v2 analytics spend table params
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) SetDimension(dimension *string) {
	o.Dimension = dimension
}

// WithEndTime adds the endTime to the get inventory API v2 analytics spend table params
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) WithEndTime(endTime *int64) *GetInventoryAPIV2AnalyticsSpendTableParams {
	o.SetEndTime(endTime)
	return o
}

// SetEndTime adds the endTime to the get inventory API v2 analytics spend table params
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) SetEndTime(endTime *int64) {
	o.EndTime = endTime
}

// WithGranularity adds the granularity to the get inventory API v2 analytics spend table params
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) WithGranularity(granularity *string) *GetInventoryAPIV2AnalyticsSpendTableParams {
	o.SetGranularity(granularity)
	return o
}

// SetGranularity adds the granularity to the get inventory API v2 analytics spend table params
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) SetGranularity(granularity *string) {
	o.Granularity = granularity
}

// WithMetricIds adds the metricIds to the get inventory API v2 analytics spend table params
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) WithMetricIds(metricIds []string) *GetInventoryAPIV2AnalyticsSpendTableParams {
	o.SetMetricIds(metricIds)
	return o
}

// SetMetricIds adds the metricIds to the get inventory API v2 analytics spend table params
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) SetMetricIds(metricIds []string) {
	o.MetricIds = metricIds
}

// WithStartTime adds the startTime to the get inventory API v2 analytics spend table params
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) WithStartTime(startTime *int64) *GetInventoryAPIV2AnalyticsSpendTableParams {
	o.SetStartTime(startTime)
	return o
}

// SetStartTime adds the startTime to the get inventory API v2 analytics spend table params
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) SetStartTime(startTime *int64) {
	o.StartTime = startTime
}

// WriteToRequest writes these params to a swagger request
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Dimension != nil {

		// query param dimension
		var qrDimension string

		if o.Dimension != nil {
			qrDimension = *o.Dimension
		}
		qDimension := qrDimension
		if qDimension != "" {

			if err := r.SetQueryParam("dimension", qDimension); err != nil {
				return err
			}
		}
	}

	if o.EndTime != nil {

		// query param endTime
		var qrEndTime int64

		if o.EndTime != nil {
			qrEndTime = *o.EndTime
		}
		qEndTime := swag.FormatInt64(qrEndTime)
		if qEndTime != "" {

			if err := r.SetQueryParam("endTime", qEndTime); err != nil {
				return err
			}
		}
	}

	if o.Granularity != nil {

		// query param granularity
		var qrGranularity string

		if o.Granularity != nil {
			qrGranularity = *o.Granularity
		}
		qGranularity := qrGranularity
		if qGranularity != "" {

			if err := r.SetQueryParam("granularity", qGranularity); err != nil {
				return err
			}
		}
	}

	if o.MetricIds != nil {

		// binding items for metricIds
		joinedMetricIds := o.bindParamMetricIds(reg)

		// query array param metricIds
		if err := r.SetQueryParam("metricIds", joinedMetricIds...); err != nil {
			return err
		}
	}

	if o.StartTime != nil {

		// query param startTime
		var qrStartTime int64

		if o.StartTime != nil {
			qrStartTime = *o.StartTime
		}
		qStartTime := swag.FormatInt64(qrStartTime)
		if qStartTime != "" {

			if err := r.SetQueryParam("startTime", qStartTime); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetInventoryAPIV2AnalyticsSpendTable binds the parameter connectionGroup
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) bindParamConnectionGroup(formats strfmt.Registry) []string {
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

// bindParamGetInventoryAPIV2AnalyticsSpendTable binds the parameter connectionId
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) bindParamConnectionID(formats strfmt.Registry) []string {
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

// bindParamGetInventoryAPIV2AnalyticsSpendTable binds the parameter connector
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) bindParamConnector(formats strfmt.Registry) []string {
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

// bindParamGetInventoryAPIV2AnalyticsSpendTable binds the parameter metricIds
func (o *GetInventoryAPIV2AnalyticsSpendTableParams) bindParamMetricIds(formats strfmt.Registry) []string {
	metricIdsIR := o.MetricIds

	var metricIdsIC []string
	for _, metricIdsIIR := range metricIdsIR { // explode []string

		metricIdsIIV := metricIdsIIR // string as string
		metricIdsIC = append(metricIdsIC, metricIdsIIV)
	}

	// items.CollectionFormat: "csv"
	metricIdsIS := swag.JoinByFormat(metricIdsIC, "csv")

	return metricIdsIS
}
