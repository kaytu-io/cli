// Code generated by go-swagger; DO NOT EDIT.

package stack

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

// NewGetScheduleAPIV1StacksResourceResourceIDParams creates a new GetScheduleAPIV1StacksResourceResourceIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetScheduleAPIV1StacksResourceResourceIDParams() *GetScheduleAPIV1StacksResourceResourceIDParams {
	return &GetScheduleAPIV1StacksResourceResourceIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetScheduleAPIV1StacksResourceResourceIDParamsWithTimeout creates a new GetScheduleAPIV1StacksResourceResourceIDParams object
// with the ability to set a timeout on a request.
func NewGetScheduleAPIV1StacksResourceResourceIDParamsWithTimeout(timeout time.Duration) *GetScheduleAPIV1StacksResourceResourceIDParams {
	return &GetScheduleAPIV1StacksResourceResourceIDParams{
		timeout: timeout,
	}
}

// NewGetScheduleAPIV1StacksResourceResourceIDParamsWithContext creates a new GetScheduleAPIV1StacksResourceResourceIDParams object
// with the ability to set a context for a request.
func NewGetScheduleAPIV1StacksResourceResourceIDParamsWithContext(ctx context.Context) *GetScheduleAPIV1StacksResourceResourceIDParams {
	return &GetScheduleAPIV1StacksResourceResourceIDParams{
		Context: ctx,
	}
}

// NewGetScheduleAPIV1StacksResourceResourceIDParamsWithHTTPClient creates a new GetScheduleAPIV1StacksResourceResourceIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetScheduleAPIV1StacksResourceResourceIDParamsWithHTTPClient(client *http.Client) *GetScheduleAPIV1StacksResourceResourceIDParams {
	return &GetScheduleAPIV1StacksResourceResourceIDParams{
		HTTPClient: client,
	}
}

/*
GetScheduleAPIV1StacksResourceResourceIDParams contains all the parameters to send to the API endpoint

	for the get schedule API v1 stacks resource resource ID operation.

	Typically these are written to a http.Request.
*/
type GetScheduleAPIV1StacksResourceResourceIDParams struct {

	/* ResourceID.

	   Resource ID
	*/
	ResourceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get schedule API v1 stacks resource resource ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScheduleAPIV1StacksResourceResourceIDParams) WithDefaults() *GetScheduleAPIV1StacksResourceResourceIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get schedule API v1 stacks resource resource ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScheduleAPIV1StacksResourceResourceIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get schedule API v1 stacks resource resource ID params
func (o *GetScheduleAPIV1StacksResourceResourceIDParams) WithTimeout(timeout time.Duration) *GetScheduleAPIV1StacksResourceResourceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get schedule API v1 stacks resource resource ID params
func (o *GetScheduleAPIV1StacksResourceResourceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get schedule API v1 stacks resource resource ID params
func (o *GetScheduleAPIV1StacksResourceResourceIDParams) WithContext(ctx context.Context) *GetScheduleAPIV1StacksResourceResourceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get schedule API v1 stacks resource resource ID params
func (o *GetScheduleAPIV1StacksResourceResourceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get schedule API v1 stacks resource resource ID params
func (o *GetScheduleAPIV1StacksResourceResourceIDParams) WithHTTPClient(client *http.Client) *GetScheduleAPIV1StacksResourceResourceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get schedule API v1 stacks resource resource ID params
func (o *GetScheduleAPIV1StacksResourceResourceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceID adds the resourceID to the get schedule API v1 stacks resource resource ID params
func (o *GetScheduleAPIV1StacksResourceResourceIDParams) WithResourceID(resourceID string) *GetScheduleAPIV1StacksResourceResourceIDParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the get schedule API v1 stacks resource resource ID params
func (o *GetScheduleAPIV1StacksResourceResourceIDParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetScheduleAPIV1StacksResourceResourceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param resourceId
	if err := r.SetPathParam("resourceId", o.ResourceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
