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
)

// NewGetScheduleAPIV1ComplianceReportLastCompletedParams creates a new GetScheduleAPIV1ComplianceReportLastCompletedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetScheduleAPIV1ComplianceReportLastCompletedParams() *GetScheduleAPIV1ComplianceReportLastCompletedParams {
	return &GetScheduleAPIV1ComplianceReportLastCompletedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetScheduleAPIV1ComplianceReportLastCompletedParamsWithTimeout creates a new GetScheduleAPIV1ComplianceReportLastCompletedParams object
// with the ability to set a timeout on a request.
func NewGetScheduleAPIV1ComplianceReportLastCompletedParamsWithTimeout(timeout time.Duration) *GetScheduleAPIV1ComplianceReportLastCompletedParams {
	return &GetScheduleAPIV1ComplianceReportLastCompletedParams{
		timeout: timeout,
	}
}

// NewGetScheduleAPIV1ComplianceReportLastCompletedParamsWithContext creates a new GetScheduleAPIV1ComplianceReportLastCompletedParams object
// with the ability to set a context for a request.
func NewGetScheduleAPIV1ComplianceReportLastCompletedParamsWithContext(ctx context.Context) *GetScheduleAPIV1ComplianceReportLastCompletedParams {
	return &GetScheduleAPIV1ComplianceReportLastCompletedParams{
		Context: ctx,
	}
}

// NewGetScheduleAPIV1ComplianceReportLastCompletedParamsWithHTTPClient creates a new GetScheduleAPIV1ComplianceReportLastCompletedParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetScheduleAPIV1ComplianceReportLastCompletedParamsWithHTTPClient(client *http.Client) *GetScheduleAPIV1ComplianceReportLastCompletedParams {
	return &GetScheduleAPIV1ComplianceReportLastCompletedParams{
		HTTPClient: client,
	}
}

/*
GetScheduleAPIV1ComplianceReportLastCompletedParams contains all the parameters to send to the API endpoint

	for the get schedule API v1 compliance report last completed operation.

	Typically these are written to a http.Request.
*/
type GetScheduleAPIV1ComplianceReportLastCompletedParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get schedule API v1 compliance report last completed params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScheduleAPIV1ComplianceReportLastCompletedParams) WithDefaults() *GetScheduleAPIV1ComplianceReportLastCompletedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get schedule API v1 compliance report last completed params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScheduleAPIV1ComplianceReportLastCompletedParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get schedule API v1 compliance report last completed params
func (o *GetScheduleAPIV1ComplianceReportLastCompletedParams) WithTimeout(timeout time.Duration) *GetScheduleAPIV1ComplianceReportLastCompletedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get schedule API v1 compliance report last completed params
func (o *GetScheduleAPIV1ComplianceReportLastCompletedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get schedule API v1 compliance report last completed params
func (o *GetScheduleAPIV1ComplianceReportLastCompletedParams) WithContext(ctx context.Context) *GetScheduleAPIV1ComplianceReportLastCompletedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get schedule API v1 compliance report last completed params
func (o *GetScheduleAPIV1ComplianceReportLastCompletedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get schedule API v1 compliance report last completed params
func (o *GetScheduleAPIV1ComplianceReportLastCompletedParams) WithHTTPClient(client *http.Client) *GetScheduleAPIV1ComplianceReportLastCompletedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get schedule API v1 compliance report last completed params
func (o *GetScheduleAPIV1ComplianceReportLastCompletedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetScheduleAPIV1ComplianceReportLastCompletedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}