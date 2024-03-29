// Code generated by go-swagger; DO NOT EDIT.

package scheduler

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

// NewGetScheduleAPIV1DiscoveryResourcetypesListParams creates a new GetScheduleAPIV1DiscoveryResourcetypesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetScheduleAPIV1DiscoveryResourcetypesListParams() *GetScheduleAPIV1DiscoveryResourcetypesListParams {
	return &GetScheduleAPIV1DiscoveryResourcetypesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetScheduleAPIV1DiscoveryResourcetypesListParamsWithTimeout creates a new GetScheduleAPIV1DiscoveryResourcetypesListParams object
// with the ability to set a timeout on a request.
func NewGetScheduleAPIV1DiscoveryResourcetypesListParamsWithTimeout(timeout time.Duration) *GetScheduleAPIV1DiscoveryResourcetypesListParams {
	return &GetScheduleAPIV1DiscoveryResourcetypesListParams{
		timeout: timeout,
	}
}

// NewGetScheduleAPIV1DiscoveryResourcetypesListParamsWithContext creates a new GetScheduleAPIV1DiscoveryResourcetypesListParams object
// with the ability to set a context for a request.
func NewGetScheduleAPIV1DiscoveryResourcetypesListParamsWithContext(ctx context.Context) *GetScheduleAPIV1DiscoveryResourcetypesListParams {
	return &GetScheduleAPIV1DiscoveryResourcetypesListParams{
		Context: ctx,
	}
}

// NewGetScheduleAPIV1DiscoveryResourcetypesListParamsWithHTTPClient creates a new GetScheduleAPIV1DiscoveryResourcetypesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetScheduleAPIV1DiscoveryResourcetypesListParamsWithHTTPClient(client *http.Client) *GetScheduleAPIV1DiscoveryResourcetypesListParams {
	return &GetScheduleAPIV1DiscoveryResourcetypesListParams{
		HTTPClient: client,
	}
}

/*
GetScheduleAPIV1DiscoveryResourcetypesListParams contains all the parameters to send to the API endpoint

	for the get schedule API v1 discovery resourcetypes list operation.

	Typically these are written to a http.Request.
*/
type GetScheduleAPIV1DiscoveryResourcetypesListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get schedule API v1 discovery resourcetypes list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScheduleAPIV1DiscoveryResourcetypesListParams) WithDefaults() *GetScheduleAPIV1DiscoveryResourcetypesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get schedule API v1 discovery resourcetypes list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScheduleAPIV1DiscoveryResourcetypesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get schedule API v1 discovery resourcetypes list params
func (o *GetScheduleAPIV1DiscoveryResourcetypesListParams) WithTimeout(timeout time.Duration) *GetScheduleAPIV1DiscoveryResourcetypesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get schedule API v1 discovery resourcetypes list params
func (o *GetScheduleAPIV1DiscoveryResourcetypesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get schedule API v1 discovery resourcetypes list params
func (o *GetScheduleAPIV1DiscoveryResourcetypesListParams) WithContext(ctx context.Context) *GetScheduleAPIV1DiscoveryResourcetypesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get schedule API v1 discovery resourcetypes list params
func (o *GetScheduleAPIV1DiscoveryResourcetypesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get schedule API v1 discovery resourcetypes list params
func (o *GetScheduleAPIV1DiscoveryResourcetypesListParams) WithHTTPClient(client *http.Client) *GetScheduleAPIV1DiscoveryResourcetypesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get schedule API v1 discovery resourcetypes list params
func (o *GetScheduleAPIV1DiscoveryResourcetypesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetScheduleAPIV1DiscoveryResourcetypesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
