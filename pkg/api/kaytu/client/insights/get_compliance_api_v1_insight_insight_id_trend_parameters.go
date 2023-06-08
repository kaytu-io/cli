// Code generated by go-swagger; DO NOT EDIT.

package insights

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

// NewGetComplianceAPIV1InsightInsightIDTrendParams creates a new GetComplianceAPIV1InsightInsightIDTrendParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetComplianceAPIV1InsightInsightIDTrendParams() *GetComplianceAPIV1InsightInsightIDTrendParams {
	return &GetComplianceAPIV1InsightInsightIDTrendParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetComplianceAPIV1InsightInsightIDTrendParamsWithTimeout creates a new GetComplianceAPIV1InsightInsightIDTrendParams object
// with the ability to set a timeout on a request.
func NewGetComplianceAPIV1InsightInsightIDTrendParamsWithTimeout(timeout time.Duration) *GetComplianceAPIV1InsightInsightIDTrendParams {
	return &GetComplianceAPIV1InsightInsightIDTrendParams{
		timeout: timeout,
	}
}

// NewGetComplianceAPIV1InsightInsightIDTrendParamsWithContext creates a new GetComplianceAPIV1InsightInsightIDTrendParams object
// with the ability to set a context for a request.
func NewGetComplianceAPIV1InsightInsightIDTrendParamsWithContext(ctx context.Context) *GetComplianceAPIV1InsightInsightIDTrendParams {
	return &GetComplianceAPIV1InsightInsightIDTrendParams{
		Context: ctx,
	}
}

// NewGetComplianceAPIV1InsightInsightIDTrendParamsWithHTTPClient creates a new GetComplianceAPIV1InsightInsightIDTrendParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetComplianceAPIV1InsightInsightIDTrendParamsWithHTTPClient(client *http.Client) *GetComplianceAPIV1InsightInsightIDTrendParams {
	return &GetComplianceAPIV1InsightInsightIDTrendParams{
		HTTPClient: client,
	}
}

/*
GetComplianceAPIV1InsightInsightIDTrendParams contains all the parameters to send to the API endpoint

	for the get compliance API v1 insight insight ID trend operation.

	Typically these are written to a http.Request.
*/
type GetComplianceAPIV1InsightInsightIDTrendParams struct {

	/* ConnectionID.

	   filter the result by source id
	*/
	ConnectionID []string

	/* DatapointCount.

	   number of datapoints to return
	*/
	DatapointCount *int64

	/* EndTime.

	   unix seconds for the end time of the trend
	*/
	EndTime *int64

	/* InsightID.

	   InsightID
	*/
	InsightID string

	/* StartTime.

	   unix seconds for the start time of the trend
	*/
	StartTime *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get compliance API v1 insight insight ID trend params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComplianceAPIV1InsightInsightIDTrendParams) WithDefaults() *GetComplianceAPIV1InsightInsightIDTrendParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get compliance API v1 insight insight ID trend params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComplianceAPIV1InsightInsightIDTrendParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get compliance API v1 insight insight ID trend params
func (o *GetComplianceAPIV1InsightInsightIDTrendParams) WithTimeout(timeout time.Duration) *GetComplianceAPIV1InsightInsightIDTrendParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get compliance API v1 insight insight ID trend params
func (o *GetComplianceAPIV1InsightInsightIDTrendParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get compliance API v1 insight insight ID trend params
func (o *GetComplianceAPIV1InsightInsightIDTrendParams) WithContext(ctx context.Context) *GetComplianceAPIV1InsightInsightIDTrendParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get compliance API v1 insight insight ID trend params
func (o *GetComplianceAPIV1InsightInsightIDTrendParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get compliance API v1 insight insight ID trend params
func (o *GetComplianceAPIV1InsightInsightIDTrendParams) WithHTTPClient(client *http.Client) *GetComplianceAPIV1InsightInsightIDTrendParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get compliance API v1 insight insight ID trend params
func (o *GetComplianceAPIV1InsightInsightIDTrendParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionID adds the connectionID to the get compliance API v1 insight insight ID trend params
func (o *GetComplianceAPIV1InsightInsightIDTrendParams) WithConnectionID(connectionID []string) *GetComplianceAPIV1InsightInsightIDTrendParams {
	o.SetConnectionID(connectionID)
	return o
}

// SetConnectionID adds the connectionId to the get compliance API v1 insight insight ID trend params
func (o *GetComplianceAPIV1InsightInsightIDTrendParams) SetConnectionID(connectionID []string) {
	o.ConnectionID = connectionID
}

// WithDatapointCount adds the datapointCount to the get compliance API v1 insight insight ID trend params
func (o *GetComplianceAPIV1InsightInsightIDTrendParams) WithDatapointCount(datapointCount *int64) *GetComplianceAPIV1InsightInsightIDTrendParams {
	o.SetDatapointCount(datapointCount)
	return o
}

// SetDatapointCount adds the datapointCount to the get compliance API v1 insight insight ID trend params
func (o *GetComplianceAPIV1InsightInsightIDTrendParams) SetDatapointCount(datapointCount *int64) {
	o.DatapointCount = datapointCount
}

// WithEndTime adds the endTime to the get compliance API v1 insight insight ID trend params
func (o *GetComplianceAPIV1InsightInsightIDTrendParams) WithEndTime(endTime *int64) *GetComplianceAPIV1InsightInsightIDTrendParams {
	o.SetEndTime(endTime)
	return o
}

// SetEndTime adds the endTime to the get compliance API v1 insight insight ID trend params
func (o *GetComplianceAPIV1InsightInsightIDTrendParams) SetEndTime(endTime *int64) {
	o.EndTime = endTime
}

// WithInsightID adds the insightID to the get compliance API v1 insight insight ID trend params
func (o *GetComplianceAPIV1InsightInsightIDTrendParams) WithInsightID(insightID string) *GetComplianceAPIV1InsightInsightIDTrendParams {
	o.SetInsightID(insightID)
	return o
}

// SetInsightID adds the insightId to the get compliance API v1 insight insight ID trend params
func (o *GetComplianceAPIV1InsightInsightIDTrendParams) SetInsightID(insightID string) {
	o.InsightID = insightID
}

// WithStartTime adds the startTime to the get compliance API v1 insight insight ID trend params
func (o *GetComplianceAPIV1InsightInsightIDTrendParams) WithStartTime(startTime *int64) *GetComplianceAPIV1InsightInsightIDTrendParams {
	o.SetStartTime(startTime)
	return o
}

// SetStartTime adds the startTime to the get compliance API v1 insight insight ID trend params
func (o *GetComplianceAPIV1InsightInsightIDTrendParams) SetStartTime(startTime *int64) {
	o.StartTime = startTime
}

// WriteToRequest writes these params to a swagger request
func (o *GetComplianceAPIV1InsightInsightIDTrendParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.DatapointCount != nil {

		// query param datapointCount
		var qrDatapointCount int64

		if o.DatapointCount != nil {
			qrDatapointCount = *o.DatapointCount
		}
		qDatapointCount := swag.FormatInt64(qrDatapointCount)
		if qDatapointCount != "" {

			if err := r.SetQueryParam("datapointCount", qDatapointCount); err != nil {
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

	// path param insightId
	if err := r.SetPathParam("insightId", o.InsightID); err != nil {
		return err
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

// bindParamGetComplianceAPIV1InsightInsightIDTrend binds the parameter connectionId
func (o *GetComplianceAPIV1InsightInsightIDTrendParams) bindParamConnectionID(formats strfmt.Registry) []string {
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
