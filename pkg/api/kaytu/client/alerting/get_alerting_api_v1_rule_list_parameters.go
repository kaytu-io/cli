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

// NewGetAlertingAPIV1RuleListParams creates a new GetAlertingAPIV1RuleListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAlertingAPIV1RuleListParams() *GetAlertingAPIV1RuleListParams {
	return &GetAlertingAPIV1RuleListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAlertingAPIV1RuleListParamsWithTimeout creates a new GetAlertingAPIV1RuleListParams object
// with the ability to set a timeout on a request.
func NewGetAlertingAPIV1RuleListParamsWithTimeout(timeout time.Duration) *GetAlertingAPIV1RuleListParams {
	return &GetAlertingAPIV1RuleListParams{
		timeout: timeout,
	}
}

// NewGetAlertingAPIV1RuleListParamsWithContext creates a new GetAlertingAPIV1RuleListParams object
// with the ability to set a context for a request.
func NewGetAlertingAPIV1RuleListParamsWithContext(ctx context.Context) *GetAlertingAPIV1RuleListParams {
	return &GetAlertingAPIV1RuleListParams{
		Context: ctx,
	}
}

// NewGetAlertingAPIV1RuleListParamsWithHTTPClient creates a new GetAlertingAPIV1RuleListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAlertingAPIV1RuleListParamsWithHTTPClient(client *http.Client) *GetAlertingAPIV1RuleListParams {
	return &GetAlertingAPIV1RuleListParams{
		HTTPClient: client,
	}
}

/*
GetAlertingAPIV1RuleListParams contains all the parameters to send to the API endpoint

	for the get alerting API v1 rule list operation.

	Typically these are written to a http.Request.
*/
type GetAlertingAPIV1RuleListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get alerting API v1 rule list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlertingAPIV1RuleListParams) WithDefaults() *GetAlertingAPIV1RuleListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get alerting API v1 rule list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlertingAPIV1RuleListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get alerting API v1 rule list params
func (o *GetAlertingAPIV1RuleListParams) WithTimeout(timeout time.Duration) *GetAlertingAPIV1RuleListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get alerting API v1 rule list params
func (o *GetAlertingAPIV1RuleListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get alerting API v1 rule list params
func (o *GetAlertingAPIV1RuleListParams) WithContext(ctx context.Context) *GetAlertingAPIV1RuleListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get alerting API v1 rule list params
func (o *GetAlertingAPIV1RuleListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get alerting API v1 rule list params
func (o *GetAlertingAPIV1RuleListParams) WithHTTPClient(client *http.Client) *GetAlertingAPIV1RuleListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get alerting API v1 rule list params
func (o *GetAlertingAPIV1RuleListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAlertingAPIV1RuleListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}