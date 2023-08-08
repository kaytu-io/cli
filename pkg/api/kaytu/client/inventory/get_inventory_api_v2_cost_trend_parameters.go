// Code generated by go-swagger; DO NOT EDIT.

package inventory

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

// NewGetInventoryAPIV2CostTrendParams creates a new GetInventoryAPIV2CostTrendParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInventoryAPIV2CostTrendParams() *GetInventoryAPIV2CostTrendParams {
	return &GetInventoryAPIV2CostTrendParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInventoryAPIV2CostTrendParamsWithTimeout creates a new GetInventoryAPIV2CostTrendParams object
// with the ability to set a timeout on a request.
func NewGetInventoryAPIV2CostTrendParamsWithTimeout(timeout time.Duration) *GetInventoryAPIV2CostTrendParams {
	return &GetInventoryAPIV2CostTrendParams{
		timeout: timeout,
	}
}

// NewGetInventoryAPIV2CostTrendParamsWithContext creates a new GetInventoryAPIV2CostTrendParams object
// with the ability to set a context for a request.
func NewGetInventoryAPIV2CostTrendParamsWithContext(ctx context.Context) *GetInventoryAPIV2CostTrendParams {
	return &GetInventoryAPIV2CostTrendParams{
		Context: ctx,
	}
}

// NewGetInventoryAPIV2CostTrendParamsWithHTTPClient creates a new GetInventoryAPIV2CostTrendParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInventoryAPIV2CostTrendParamsWithHTTPClient(client *http.Client) *GetInventoryAPIV2CostTrendParams {
	return &GetInventoryAPIV2CostTrendParams{
		HTTPClient: client,
	}
}

/*
GetInventoryAPIV2CostTrendParams contains all the parameters to send to the API endpoint

	for the get inventory API v2 cost trend operation.

	Typically these are written to a http.Request.
*/
type GetInventoryAPIV2CostTrendParams struct {

	/* ConnectionGroup.

	   Connection group to filter by - mutually exclusive with connectionId
	*/
	ConnectionGroup *string

	/* ConnectionID.

	   Connection IDs to filter by - mutually exclusive with connectionGroup
	*/
	ConnectionID []string

	/* Connector.

	   Connector type to filter by
	*/
	Connector []string

	/* DatapointCount.

	   maximum number of datapoints to return, default is 30
	*/
	DatapointCount *string

	/* EndTime.

	   timestamp for end in epoch seconds
	*/
	EndTime *string

	/* StartTime.

	   timestamp for start in epoch seconds
	*/
	StartTime *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get inventory API v2 cost trend params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInventoryAPIV2CostTrendParams) WithDefaults() *GetInventoryAPIV2CostTrendParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get inventory API v2 cost trend params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInventoryAPIV2CostTrendParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get inventory API v2 cost trend params
func (o *GetInventoryAPIV2CostTrendParams) WithTimeout(timeout time.Duration) *GetInventoryAPIV2CostTrendParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get inventory API v2 cost trend params
func (o *GetInventoryAPIV2CostTrendParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get inventory API v2 cost trend params
func (o *GetInventoryAPIV2CostTrendParams) WithContext(ctx context.Context) *GetInventoryAPIV2CostTrendParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get inventory API v2 cost trend params
func (o *GetInventoryAPIV2CostTrendParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get inventory API v2 cost trend params
func (o *GetInventoryAPIV2CostTrendParams) WithHTTPClient(client *http.Client) *GetInventoryAPIV2CostTrendParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get inventory API v2 cost trend params
func (o *GetInventoryAPIV2CostTrendParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionGroup adds the connectionGroup to the get inventory API v2 cost trend params
func (o *GetInventoryAPIV2CostTrendParams) WithConnectionGroup(connectionGroup *string) *GetInventoryAPIV2CostTrendParams {
	o.SetConnectionGroup(connectionGroup)
	return o
}

// SetConnectionGroup adds the connectionGroup to the get inventory API v2 cost trend params
func (o *GetInventoryAPIV2CostTrendParams) SetConnectionGroup(connectionGroup *string) {
	o.ConnectionGroup = connectionGroup
}

// WithConnectionID adds the connectionID to the get inventory API v2 cost trend params
func (o *GetInventoryAPIV2CostTrendParams) WithConnectionID(connectionID []string) *GetInventoryAPIV2CostTrendParams {
	o.SetConnectionID(connectionID)
	return o
}

// SetConnectionID adds the connectionId to the get inventory API v2 cost trend params
func (o *GetInventoryAPIV2CostTrendParams) SetConnectionID(connectionID []string) {
	o.ConnectionID = connectionID
}

// WithConnector adds the connector to the get inventory API v2 cost trend params
func (o *GetInventoryAPIV2CostTrendParams) WithConnector(connector []string) *GetInventoryAPIV2CostTrendParams {
	o.SetConnector(connector)
	return o
}

// SetConnector adds the connector to the get inventory API v2 cost trend params
func (o *GetInventoryAPIV2CostTrendParams) SetConnector(connector []string) {
	o.Connector = connector
}

// WithDatapointCount adds the datapointCount to the get inventory API v2 cost trend params
func (o *GetInventoryAPIV2CostTrendParams) WithDatapointCount(datapointCount *string) *GetInventoryAPIV2CostTrendParams {
	o.SetDatapointCount(datapointCount)
	return o
}

// SetDatapointCount adds the datapointCount to the get inventory API v2 cost trend params
func (o *GetInventoryAPIV2CostTrendParams) SetDatapointCount(datapointCount *string) {
	o.DatapointCount = datapointCount
}

// WithEndTime adds the endTime to the get inventory API v2 cost trend params
func (o *GetInventoryAPIV2CostTrendParams) WithEndTime(endTime *string) *GetInventoryAPIV2CostTrendParams {
	o.SetEndTime(endTime)
	return o
}

// SetEndTime adds the endTime to the get inventory API v2 cost trend params
func (o *GetInventoryAPIV2CostTrendParams) SetEndTime(endTime *string) {
	o.EndTime = endTime
}

// WithStartTime adds the startTime to the get inventory API v2 cost trend params
func (o *GetInventoryAPIV2CostTrendParams) WithStartTime(startTime *string) *GetInventoryAPIV2CostTrendParams {
	o.SetStartTime(startTime)
	return o
}

// SetStartTime adds the startTime to the get inventory API v2 cost trend params
func (o *GetInventoryAPIV2CostTrendParams) SetStartTime(startTime *string) {
	o.StartTime = startTime
}

// WriteToRequest writes these params to a swagger request
func (o *GetInventoryAPIV2CostTrendParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ConnectionGroup != nil {

		// query param connectionGroup
		var qrConnectionGroup string

		if o.ConnectionGroup != nil {
			qrConnectionGroup = *o.ConnectionGroup
		}
		qConnectionGroup := qrConnectionGroup
		if qConnectionGroup != "" {

			if err := r.SetQueryParam("connectionGroup", qConnectionGroup); err != nil {
				return err
			}
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

	if o.DatapointCount != nil {

		// query param datapointCount
		var qrDatapointCount string

		if o.DatapointCount != nil {
			qrDatapointCount = *o.DatapointCount
		}
		qDatapointCount := qrDatapointCount
		if qDatapointCount != "" {

			if err := r.SetQueryParam("datapointCount", qDatapointCount); err != nil {
				return err
			}
		}
	}

	if o.EndTime != nil {

		// query param endTime
		var qrEndTime string

		if o.EndTime != nil {
			qrEndTime = *o.EndTime
		}
		qEndTime := qrEndTime
		if qEndTime != "" {

			if err := r.SetQueryParam("endTime", qEndTime); err != nil {
				return err
			}
		}
	}

	if o.StartTime != nil {

		// query param startTime
		var qrStartTime string

		if o.StartTime != nil {
			qrStartTime = *o.StartTime
		}
		qStartTime := qrStartTime
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

// bindParamGetInventoryAPIV2CostTrend binds the parameter connectionId
func (o *GetInventoryAPIV2CostTrendParams) bindParamConnectionID(formats strfmt.Registry) []string {
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

// bindParamGetInventoryAPIV2CostTrend binds the parameter connector
func (o *GetInventoryAPIV2CostTrendParams) bindParamConnector(formats strfmt.Registry) []string {
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
