// Code generated by go-swagger; DO NOT EDIT.

package insight

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

// NewGetInventoryAPIV2InsightsParams creates a new GetInventoryAPIV2InsightsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInventoryAPIV2InsightsParams() *GetInventoryAPIV2InsightsParams {
	return &GetInventoryAPIV2InsightsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInventoryAPIV2InsightsParamsWithTimeout creates a new GetInventoryAPIV2InsightsParams object
// with the ability to set a timeout on a request.
func NewGetInventoryAPIV2InsightsParamsWithTimeout(timeout time.Duration) *GetInventoryAPIV2InsightsParams {
	return &GetInventoryAPIV2InsightsParams{
		timeout: timeout,
	}
}

// NewGetInventoryAPIV2InsightsParamsWithContext creates a new GetInventoryAPIV2InsightsParams object
// with the ability to set a context for a request.
func NewGetInventoryAPIV2InsightsParamsWithContext(ctx context.Context) *GetInventoryAPIV2InsightsParams {
	return &GetInventoryAPIV2InsightsParams{
		Context: ctx,
	}
}

// NewGetInventoryAPIV2InsightsParamsWithHTTPClient creates a new GetInventoryAPIV2InsightsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInventoryAPIV2InsightsParamsWithHTTPClient(client *http.Client) *GetInventoryAPIV2InsightsParams {
	return &GetInventoryAPIV2InsightsParams{
		HTTPClient: client,
	}
}

/*
GetInventoryAPIV2InsightsParams contains all the parameters to send to the API endpoint

	for the get inventory API v2 insights operation.

	Typically these are written to a http.Request.
*/
type GetInventoryAPIV2InsightsParams struct {

	/* ConnectionID.

	   filter the result by source id
	*/
	ConnectionID []string

	/* Connector.

	   filter insights by connector
	*/
	Connector []string

	/* InsightID.

	   filter the result by insight id
	*/
	InsightID []string

	/* Time.

	   unix seconds for the time to get the insight result for
	*/
	Time *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get inventory API v2 insights params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInventoryAPIV2InsightsParams) WithDefaults() *GetInventoryAPIV2InsightsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get inventory API v2 insights params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInventoryAPIV2InsightsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get inventory API v2 insights params
func (o *GetInventoryAPIV2InsightsParams) WithTimeout(timeout time.Duration) *GetInventoryAPIV2InsightsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get inventory API v2 insights params
func (o *GetInventoryAPIV2InsightsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get inventory API v2 insights params
func (o *GetInventoryAPIV2InsightsParams) WithContext(ctx context.Context) *GetInventoryAPIV2InsightsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get inventory API v2 insights params
func (o *GetInventoryAPIV2InsightsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get inventory API v2 insights params
func (o *GetInventoryAPIV2InsightsParams) WithHTTPClient(client *http.Client) *GetInventoryAPIV2InsightsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get inventory API v2 insights params
func (o *GetInventoryAPIV2InsightsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionID adds the connectionID to the get inventory API v2 insights params
func (o *GetInventoryAPIV2InsightsParams) WithConnectionID(connectionID []string) *GetInventoryAPIV2InsightsParams {
	o.SetConnectionID(connectionID)
	return o
}

// SetConnectionID adds the connectionId to the get inventory API v2 insights params
func (o *GetInventoryAPIV2InsightsParams) SetConnectionID(connectionID []string) {
	o.ConnectionID = connectionID
}

// WithConnector adds the connector to the get inventory API v2 insights params
func (o *GetInventoryAPIV2InsightsParams) WithConnector(connector []string) *GetInventoryAPIV2InsightsParams {
	o.SetConnector(connector)
	return o
}

// SetConnector adds the connector to the get inventory API v2 insights params
func (o *GetInventoryAPIV2InsightsParams) SetConnector(connector []string) {
	o.Connector = connector
}

// WithInsightID adds the insightID to the get inventory API v2 insights params
func (o *GetInventoryAPIV2InsightsParams) WithInsightID(insightID []string) *GetInventoryAPIV2InsightsParams {
	o.SetInsightID(insightID)
	return o
}

// SetInsightID adds the insightId to the get inventory API v2 insights params
func (o *GetInventoryAPIV2InsightsParams) SetInsightID(insightID []string) {
	o.InsightID = insightID
}

// WithTime adds the time to the get inventory API v2 insights params
func (o *GetInventoryAPIV2InsightsParams) WithTime(time *int64) *GetInventoryAPIV2InsightsParams {
	o.SetTime(time)
	return o
}

// SetTime adds the time to the get inventory API v2 insights params
func (o *GetInventoryAPIV2InsightsParams) SetTime(time *int64) {
	o.Time = time
}

// WriteToRequest writes these params to a swagger request
func (o *GetInventoryAPIV2InsightsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.InsightID != nil {

		// binding items for insightId
		joinedInsightID := o.bindParamInsightID(reg)

		// query array param insightId
		if err := r.SetQueryParam("insightId", joinedInsightID...); err != nil {
			return err
		}
	}

	if o.Time != nil {

		// query param time
		var qrTime int64

		if o.Time != nil {
			qrTime = *o.Time
		}
		qTime := swag.FormatInt64(qrTime)
		if qTime != "" {

			if err := r.SetQueryParam("time", qTime); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetInventoryAPIV2Insights binds the parameter connectionId
func (o *GetInventoryAPIV2InsightsParams) bindParamConnectionID(formats strfmt.Registry) []string {
	connectionIDIR := o.ConnectionID

	var connectionIDIC []string
	for _, connectionIDIIR := range connectionIDIR { // explode []string

		connectionIDIIV := connectionIDIIR // string as string
		connectionIDIC = append(connectionIDIC, connectionIDIIV)
	}

	// items.CollectionFormat: ""
	connectionIDIS := swag.JoinByFormat(connectionIDIC, "")

	return connectionIDIS
}

// bindParamGetInventoryAPIV2Insights binds the parameter connector
func (o *GetInventoryAPIV2InsightsParams) bindParamConnector(formats strfmt.Registry) []string {
	connectorIR := o.Connector

	var connectorIC []string
	for _, connectorIIR := range connectorIR { // explode []string

		connectorIIV := connectorIIR // string as string
		connectorIC = append(connectorIC, connectorIIV)
	}

	// items.CollectionFormat: ""
	connectorIS := swag.JoinByFormat(connectorIC, "")

	return connectorIS
}

// bindParamGetInventoryAPIV2Insights binds the parameter insightId
func (o *GetInventoryAPIV2InsightsParams) bindParamInsightID(formats strfmt.Registry) []string {
	insightIDIR := o.InsightID

	var insightIDIC []string
	for _, insightIDIIR := range insightIDIR { // explode []string

		insightIDIIV := insightIDIIR // string as string
		insightIDIC = append(insightIDIC, insightIDIIV)
	}

	// items.CollectionFormat: ""
	insightIDIS := swag.JoinByFormat(insightIDIC, "")

	return insightIDIS
}