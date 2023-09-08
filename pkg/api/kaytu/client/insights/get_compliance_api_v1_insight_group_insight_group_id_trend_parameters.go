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

// NewGetComplianceAPIV1InsightGroupInsightGroupIDTrendParams creates a new GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetComplianceAPIV1InsightGroupInsightGroupIDTrendParams() *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams {
	return &GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetComplianceAPIV1InsightGroupInsightGroupIDTrendParamsWithTimeout creates a new GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams object
// with the ability to set a timeout on a request.
func NewGetComplianceAPIV1InsightGroupInsightGroupIDTrendParamsWithTimeout(timeout time.Duration) *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams {
	return &GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams{
		timeout: timeout,
	}
}

// NewGetComplianceAPIV1InsightGroupInsightGroupIDTrendParamsWithContext creates a new GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams object
// with the ability to set a context for a request.
func NewGetComplianceAPIV1InsightGroupInsightGroupIDTrendParamsWithContext(ctx context.Context) *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams {
	return &GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams{
		Context: ctx,
	}
}

// NewGetComplianceAPIV1InsightGroupInsightGroupIDTrendParamsWithHTTPClient creates a new GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetComplianceAPIV1InsightGroupInsightGroupIDTrendParamsWithHTTPClient(client *http.Client) *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams {
	return &GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams{
		HTTPClient: client,
	}
}

/*
GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams contains all the parameters to send to the API endpoint

	for the get compliance API v1 insight group insight group ID trend operation.

	Typically these are written to a http.Request.
*/
type GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams struct {

	/* ConnectionGroup.

	   filter the result by connection group
	*/
	ConnectionGroup []string

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

	/* InsightGroupID.

	   Insight Group ID
	*/
	InsightGroupID string

	/* StartTime.

	   unix seconds for the start time of the trend
	*/
	StartTime *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get compliance API v1 insight group insight group ID trend params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams) WithDefaults() *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get compliance API v1 insight group insight group ID trend params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get compliance API v1 insight group insight group ID trend params
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams) WithTimeout(timeout time.Duration) *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get compliance API v1 insight group insight group ID trend params
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get compliance API v1 insight group insight group ID trend params
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams) WithContext(ctx context.Context) *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get compliance API v1 insight group insight group ID trend params
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get compliance API v1 insight group insight group ID trend params
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams) WithHTTPClient(client *http.Client) *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get compliance API v1 insight group insight group ID trend params
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionGroup adds the connectionGroup to the get compliance API v1 insight group insight group ID trend params
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams) WithConnectionGroup(connectionGroup []string) *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams {
	o.SetConnectionGroup(connectionGroup)
	return o
}

// SetConnectionGroup adds the connectionGroup to the get compliance API v1 insight group insight group ID trend params
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams) SetConnectionGroup(connectionGroup []string) {
	o.ConnectionGroup = connectionGroup
}

// WithConnectionID adds the connectionID to the get compliance API v1 insight group insight group ID trend params
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams) WithConnectionID(connectionID []string) *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams {
	o.SetConnectionID(connectionID)
	return o
}

// SetConnectionID adds the connectionId to the get compliance API v1 insight group insight group ID trend params
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams) SetConnectionID(connectionID []string) {
	o.ConnectionID = connectionID
}

// WithDatapointCount adds the datapointCount to the get compliance API v1 insight group insight group ID trend params
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams) WithDatapointCount(datapointCount *int64) *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams {
	o.SetDatapointCount(datapointCount)
	return o
}

// SetDatapointCount adds the datapointCount to the get compliance API v1 insight group insight group ID trend params
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams) SetDatapointCount(datapointCount *int64) {
	o.DatapointCount = datapointCount
}

// WithEndTime adds the endTime to the get compliance API v1 insight group insight group ID trend params
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams) WithEndTime(endTime *int64) *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams {
	o.SetEndTime(endTime)
	return o
}

// SetEndTime adds the endTime to the get compliance API v1 insight group insight group ID trend params
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams) SetEndTime(endTime *int64) {
	o.EndTime = endTime
}

// WithInsightGroupID adds the insightGroupID to the get compliance API v1 insight group insight group ID trend params
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams) WithInsightGroupID(insightGroupID string) *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams {
	o.SetInsightGroupID(insightGroupID)
	return o
}

// SetInsightGroupID adds the insightGroupId to the get compliance API v1 insight group insight group ID trend params
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams) SetInsightGroupID(insightGroupID string) {
	o.InsightGroupID = insightGroupID
}

// WithStartTime adds the startTime to the get compliance API v1 insight group insight group ID trend params
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams) WithStartTime(startTime *int64) *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams {
	o.SetStartTime(startTime)
	return o
}

// SetStartTime adds the startTime to the get compliance API v1 insight group insight group ID trend params
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams) SetStartTime(startTime *int64) {
	o.StartTime = startTime
}

// WriteToRequest writes these params to a swagger request
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param insightGroupId
	if err := r.SetPathParam("insightGroupId", o.InsightGroupID); err != nil {
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

// bindParamGetComplianceAPIV1InsightGroupInsightGroupIDTrend binds the parameter connectionGroup
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams) bindParamConnectionGroup(formats strfmt.Registry) []string {
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

// bindParamGetComplianceAPIV1InsightGroupInsightGroupIDTrend binds the parameter connectionId
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendParams) bindParamConnectionID(formats strfmt.Registry) []string {
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
