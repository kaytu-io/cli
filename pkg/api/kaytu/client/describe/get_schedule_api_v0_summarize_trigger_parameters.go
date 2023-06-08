// Code generated by go-swagger; DO NOT EDIT.

package describe

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

// NewGetScheduleAPIV0SummarizeTriggerParams creates a new GetScheduleAPIV0SummarizeTriggerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetScheduleAPIV0SummarizeTriggerParams() *GetScheduleAPIV0SummarizeTriggerParams {
	return &GetScheduleAPIV0SummarizeTriggerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetScheduleAPIV0SummarizeTriggerParamsWithTimeout creates a new GetScheduleAPIV0SummarizeTriggerParams object
// with the ability to set a timeout on a request.
func NewGetScheduleAPIV0SummarizeTriggerParamsWithTimeout(timeout time.Duration) *GetScheduleAPIV0SummarizeTriggerParams {
	return &GetScheduleAPIV0SummarizeTriggerParams{
		timeout: timeout,
	}
}

// NewGetScheduleAPIV0SummarizeTriggerParamsWithContext creates a new GetScheduleAPIV0SummarizeTriggerParams object
// with the ability to set a context for a request.
func NewGetScheduleAPIV0SummarizeTriggerParamsWithContext(ctx context.Context) *GetScheduleAPIV0SummarizeTriggerParams {
	return &GetScheduleAPIV0SummarizeTriggerParams{
		Context: ctx,
	}
}

// NewGetScheduleAPIV0SummarizeTriggerParamsWithHTTPClient creates a new GetScheduleAPIV0SummarizeTriggerParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetScheduleAPIV0SummarizeTriggerParamsWithHTTPClient(client *http.Client) *GetScheduleAPIV0SummarizeTriggerParams {
	return &GetScheduleAPIV0SummarizeTriggerParams{
		HTTPClient: client,
	}
}

/*
GetScheduleAPIV0SummarizeTriggerParams contains all the parameters to send to the API endpoint

	for the get schedule API v0 summarize trigger operation.

	Typically these are written to a http.Request.
*/
type GetScheduleAPIV0SummarizeTriggerParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get schedule API v0 summarize trigger params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScheduleAPIV0SummarizeTriggerParams) WithDefaults() *GetScheduleAPIV0SummarizeTriggerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get schedule API v0 summarize trigger params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScheduleAPIV0SummarizeTriggerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get schedule API v0 summarize trigger params
func (o *GetScheduleAPIV0SummarizeTriggerParams) WithTimeout(timeout time.Duration) *GetScheduleAPIV0SummarizeTriggerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get schedule API v0 summarize trigger params
func (o *GetScheduleAPIV0SummarizeTriggerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get schedule API v0 summarize trigger params
func (o *GetScheduleAPIV0SummarizeTriggerParams) WithContext(ctx context.Context) *GetScheduleAPIV0SummarizeTriggerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get schedule API v0 summarize trigger params
func (o *GetScheduleAPIV0SummarizeTriggerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get schedule API v0 summarize trigger params
func (o *GetScheduleAPIV0SummarizeTriggerParams) WithHTTPClient(client *http.Client) *GetScheduleAPIV0SummarizeTriggerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get schedule API v0 summarize trigger params
func (o *GetScheduleAPIV0SummarizeTriggerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetScheduleAPIV0SummarizeTriggerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
