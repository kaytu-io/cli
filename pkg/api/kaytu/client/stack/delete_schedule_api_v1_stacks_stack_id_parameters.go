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

// NewDeleteScheduleAPIV1StacksStackIDParams creates a new DeleteScheduleAPIV1StacksStackIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteScheduleAPIV1StacksStackIDParams() *DeleteScheduleAPIV1StacksStackIDParams {
	return &DeleteScheduleAPIV1StacksStackIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteScheduleAPIV1StacksStackIDParamsWithTimeout creates a new DeleteScheduleAPIV1StacksStackIDParams object
// with the ability to set a timeout on a request.
func NewDeleteScheduleAPIV1StacksStackIDParamsWithTimeout(timeout time.Duration) *DeleteScheduleAPIV1StacksStackIDParams {
	return &DeleteScheduleAPIV1StacksStackIDParams{
		timeout: timeout,
	}
}

// NewDeleteScheduleAPIV1StacksStackIDParamsWithContext creates a new DeleteScheduleAPIV1StacksStackIDParams object
// with the ability to set a context for a request.
func NewDeleteScheduleAPIV1StacksStackIDParamsWithContext(ctx context.Context) *DeleteScheduleAPIV1StacksStackIDParams {
	return &DeleteScheduleAPIV1StacksStackIDParams{
		Context: ctx,
	}
}

// NewDeleteScheduleAPIV1StacksStackIDParamsWithHTTPClient creates a new DeleteScheduleAPIV1StacksStackIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteScheduleAPIV1StacksStackIDParamsWithHTTPClient(client *http.Client) *DeleteScheduleAPIV1StacksStackIDParams {
	return &DeleteScheduleAPIV1StacksStackIDParams{
		HTTPClient: client,
	}
}

/*
DeleteScheduleAPIV1StacksStackIDParams contains all the parameters to send to the API endpoint

	for the delete schedule API v1 stacks stack ID operation.

	Typically these are written to a http.Request.
*/
type DeleteScheduleAPIV1StacksStackIDParams struct {

	/* StackID.

	   StackID
	*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete schedule API v1 stacks stack ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteScheduleAPIV1StacksStackIDParams) WithDefaults() *DeleteScheduleAPIV1StacksStackIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete schedule API v1 stacks stack ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteScheduleAPIV1StacksStackIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete schedule API v1 stacks stack ID params
func (o *DeleteScheduleAPIV1StacksStackIDParams) WithTimeout(timeout time.Duration) *DeleteScheduleAPIV1StacksStackIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete schedule API v1 stacks stack ID params
func (o *DeleteScheduleAPIV1StacksStackIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete schedule API v1 stacks stack ID params
func (o *DeleteScheduleAPIV1StacksStackIDParams) WithContext(ctx context.Context) *DeleteScheduleAPIV1StacksStackIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete schedule API v1 stacks stack ID params
func (o *DeleteScheduleAPIV1StacksStackIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete schedule API v1 stacks stack ID params
func (o *DeleteScheduleAPIV1StacksStackIDParams) WithHTTPClient(client *http.Client) *DeleteScheduleAPIV1StacksStackIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete schedule API v1 stacks stack ID params
func (o *DeleteScheduleAPIV1StacksStackIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStackID adds the stackID to the delete schedule API v1 stacks stack ID params
func (o *DeleteScheduleAPIV1StacksStackIDParams) WithStackID(stackID string) *DeleteScheduleAPIV1StacksStackIDParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the delete schedule API v1 stacks stack ID params
func (o *DeleteScheduleAPIV1StacksStackIDParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteScheduleAPIV1StacksStackIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param stackId
	if err := r.SetPathParam("stackId", o.StackID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
