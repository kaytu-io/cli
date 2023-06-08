// Code generated by go-swagger; DO NOT EDIT.

package schedule

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

// NewGetScheduleAPIV1SourcesSourceIDJobsComplianceParams creates a new GetScheduleAPIV1SourcesSourceIDJobsComplianceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetScheduleAPIV1SourcesSourceIDJobsComplianceParams() *GetScheduleAPIV1SourcesSourceIDJobsComplianceParams {
	return &GetScheduleAPIV1SourcesSourceIDJobsComplianceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetScheduleAPIV1SourcesSourceIDJobsComplianceParamsWithTimeout creates a new GetScheduleAPIV1SourcesSourceIDJobsComplianceParams object
// with the ability to set a timeout on a request.
func NewGetScheduleAPIV1SourcesSourceIDJobsComplianceParamsWithTimeout(timeout time.Duration) *GetScheduleAPIV1SourcesSourceIDJobsComplianceParams {
	return &GetScheduleAPIV1SourcesSourceIDJobsComplianceParams{
		timeout: timeout,
	}
}

// NewGetScheduleAPIV1SourcesSourceIDJobsComplianceParamsWithContext creates a new GetScheduleAPIV1SourcesSourceIDJobsComplianceParams object
// with the ability to set a context for a request.
func NewGetScheduleAPIV1SourcesSourceIDJobsComplianceParamsWithContext(ctx context.Context) *GetScheduleAPIV1SourcesSourceIDJobsComplianceParams {
	return &GetScheduleAPIV1SourcesSourceIDJobsComplianceParams{
		Context: ctx,
	}
}

// NewGetScheduleAPIV1SourcesSourceIDJobsComplianceParamsWithHTTPClient creates a new GetScheduleAPIV1SourcesSourceIDJobsComplianceParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetScheduleAPIV1SourcesSourceIDJobsComplianceParamsWithHTTPClient(client *http.Client) *GetScheduleAPIV1SourcesSourceIDJobsComplianceParams {
	return &GetScheduleAPIV1SourcesSourceIDJobsComplianceParams{
		HTTPClient: client,
	}
}

/*
GetScheduleAPIV1SourcesSourceIDJobsComplianceParams contains all the parameters to send to the API endpoint

	for the get schedule API v1 sources source ID jobs compliance operation.

	Typically these are written to a http.Request.
*/
type GetScheduleAPIV1SourcesSourceIDJobsComplianceParams struct {

	/* From.

	   From Time (TimeRange)
	*/
	From *int64

	/* SourceID.

	   SourceID
	*/
	SourceID string

	/* To.

	   To Time (TimeRange)
	*/
	To *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get schedule API v1 sources source ID jobs compliance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScheduleAPIV1SourcesSourceIDJobsComplianceParams) WithDefaults() *GetScheduleAPIV1SourcesSourceIDJobsComplianceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get schedule API v1 sources source ID jobs compliance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScheduleAPIV1SourcesSourceIDJobsComplianceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get schedule API v1 sources source ID jobs compliance params
func (o *GetScheduleAPIV1SourcesSourceIDJobsComplianceParams) WithTimeout(timeout time.Duration) *GetScheduleAPIV1SourcesSourceIDJobsComplianceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get schedule API v1 sources source ID jobs compliance params
func (o *GetScheduleAPIV1SourcesSourceIDJobsComplianceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get schedule API v1 sources source ID jobs compliance params
func (o *GetScheduleAPIV1SourcesSourceIDJobsComplianceParams) WithContext(ctx context.Context) *GetScheduleAPIV1SourcesSourceIDJobsComplianceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get schedule API v1 sources source ID jobs compliance params
func (o *GetScheduleAPIV1SourcesSourceIDJobsComplianceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get schedule API v1 sources source ID jobs compliance params
func (o *GetScheduleAPIV1SourcesSourceIDJobsComplianceParams) WithHTTPClient(client *http.Client) *GetScheduleAPIV1SourcesSourceIDJobsComplianceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get schedule API v1 sources source ID jobs compliance params
func (o *GetScheduleAPIV1SourcesSourceIDJobsComplianceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFrom adds the from to the get schedule API v1 sources source ID jobs compliance params
func (o *GetScheduleAPIV1SourcesSourceIDJobsComplianceParams) WithFrom(from *int64) *GetScheduleAPIV1SourcesSourceIDJobsComplianceParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the get schedule API v1 sources source ID jobs compliance params
func (o *GetScheduleAPIV1SourcesSourceIDJobsComplianceParams) SetFrom(from *int64) {
	o.From = from
}

// WithSourceID adds the sourceID to the get schedule API v1 sources source ID jobs compliance params
func (o *GetScheduleAPIV1SourcesSourceIDJobsComplianceParams) WithSourceID(sourceID string) *GetScheduleAPIV1SourcesSourceIDJobsComplianceParams {
	o.SetSourceID(sourceID)
	return o
}

// SetSourceID adds the sourceId to the get schedule API v1 sources source ID jobs compliance params
func (o *GetScheduleAPIV1SourcesSourceIDJobsComplianceParams) SetSourceID(sourceID string) {
	o.SourceID = sourceID
}

// WithTo adds the to to the get schedule API v1 sources source ID jobs compliance params
func (o *GetScheduleAPIV1SourcesSourceIDJobsComplianceParams) WithTo(to *int64) *GetScheduleAPIV1SourcesSourceIDJobsComplianceParams {
	o.SetTo(to)
	return o
}

// SetTo adds the to to the get schedule API v1 sources source ID jobs compliance params
func (o *GetScheduleAPIV1SourcesSourceIDJobsComplianceParams) SetTo(to *int64) {
	o.To = to
}

// WriteToRequest writes these params to a swagger request
func (o *GetScheduleAPIV1SourcesSourceIDJobsComplianceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.From != nil {

		// query param from
		var qrFrom int64

		if o.From != nil {
			qrFrom = *o.From
		}
		qFrom := swag.FormatInt64(qrFrom)
		if qFrom != "" {

			if err := r.SetQueryParam("from", qFrom); err != nil {
				return err
			}
		}
	}

	// path param source_id
	if err := r.SetPathParam("source_id", o.SourceID); err != nil {
		return err
	}

	if o.To != nil {

		// query param to
		var qrTo int64

		if o.To != nil {
			qrTo = *o.To
		}
		qTo := swag.FormatInt64(qrTo)
		if qTo != "" {

			if err := r.SetQueryParam("to", qTo); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}