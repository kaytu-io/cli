// Code generated by go-swagger; DO NOT EDIT.

package alerting

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

// NewGetAlertingAPIV1TriggerListParams creates a new GetAlertingAPIV1TriggerListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAlertingAPIV1TriggerListParams() *GetAlertingAPIV1TriggerListParams {
	return &GetAlertingAPIV1TriggerListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAlertingAPIV1TriggerListParamsWithTimeout creates a new GetAlertingAPIV1TriggerListParams object
// with the ability to set a timeout on a request.
func NewGetAlertingAPIV1TriggerListParamsWithTimeout(timeout time.Duration) *GetAlertingAPIV1TriggerListParams {
	return &GetAlertingAPIV1TriggerListParams{
		timeout: timeout,
	}
}

// NewGetAlertingAPIV1TriggerListParamsWithContext creates a new GetAlertingAPIV1TriggerListParams object
// with the ability to set a context for a request.
func NewGetAlertingAPIV1TriggerListParamsWithContext(ctx context.Context) *GetAlertingAPIV1TriggerListParams {
	return &GetAlertingAPIV1TriggerListParams{
		Context: ctx,
	}
}

// NewGetAlertingAPIV1TriggerListParamsWithHTTPClient creates a new GetAlertingAPIV1TriggerListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAlertingAPIV1TriggerListParamsWithHTTPClient(client *http.Client) *GetAlertingAPIV1TriggerListParams {
	return &GetAlertingAPIV1TriggerListParams{
		HTTPClient: client,
	}
}

/*
GetAlertingAPIV1TriggerListParams contains all the parameters to send to the API endpoint

	for the get alerting API v1 trigger list operation.

	Typically these are written to a http.Request.
*/
type GetAlertingAPIV1TriggerListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get alerting API v1 trigger list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlertingAPIV1TriggerListParams) WithDefaults() *GetAlertingAPIV1TriggerListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get alerting API v1 trigger list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlertingAPIV1TriggerListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get alerting API v1 trigger list params
func (o *GetAlertingAPIV1TriggerListParams) WithTimeout(timeout time.Duration) *GetAlertingAPIV1TriggerListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get alerting API v1 trigger list params
func (o *GetAlertingAPIV1TriggerListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get alerting API v1 trigger list params
func (o *GetAlertingAPIV1TriggerListParams) WithContext(ctx context.Context) *GetAlertingAPIV1TriggerListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get alerting API v1 trigger list params
func (o *GetAlertingAPIV1TriggerListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get alerting API v1 trigger list params
func (o *GetAlertingAPIV1TriggerListParams) WithHTTPClient(client *http.Client) *GetAlertingAPIV1TriggerListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get alerting API v1 trigger list params
func (o *GetAlertingAPIV1TriggerListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAlertingAPIV1TriggerListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}