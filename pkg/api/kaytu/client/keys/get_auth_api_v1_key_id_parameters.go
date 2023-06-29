// Code generated by go-swagger; DO NOT EDIT.

package keys

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

// NewGetAuthAPIV1KeyIDParams creates a new GetAuthAPIV1KeyIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAuthAPIV1KeyIDParams() *GetAuthAPIV1KeyIDParams {
	return &GetAuthAPIV1KeyIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthAPIV1KeyIDParamsWithTimeout creates a new GetAuthAPIV1KeyIDParams object
// with the ability to set a timeout on a request.
func NewGetAuthAPIV1KeyIDParamsWithTimeout(timeout time.Duration) *GetAuthAPIV1KeyIDParams {
	return &GetAuthAPIV1KeyIDParams{
		timeout: timeout,
	}
}

// NewGetAuthAPIV1KeyIDParamsWithContext creates a new GetAuthAPIV1KeyIDParams object
// with the ability to set a context for a request.
func NewGetAuthAPIV1KeyIDParamsWithContext(ctx context.Context) *GetAuthAPIV1KeyIDParams {
	return &GetAuthAPIV1KeyIDParams{
		Context: ctx,
	}
}

// NewGetAuthAPIV1KeyIDParamsWithHTTPClient creates a new GetAuthAPIV1KeyIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAuthAPIV1KeyIDParamsWithHTTPClient(client *http.Client) *GetAuthAPIV1KeyIDParams {
	return &GetAuthAPIV1KeyIDParams{
		HTTPClient: client,
	}
}

/*
GetAuthAPIV1KeyIDParams contains all the parameters to send to the API endpoint

	for the get auth API v1 key ID operation.

	Typically these are written to a http.Request.
*/
type GetAuthAPIV1KeyIDParams struct {

	/* ID.

	   Key ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get auth API v1 key ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAuthAPIV1KeyIDParams) WithDefaults() *GetAuthAPIV1KeyIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get auth API v1 key ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAuthAPIV1KeyIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get auth API v1 key ID params
func (o *GetAuthAPIV1KeyIDParams) WithTimeout(timeout time.Duration) *GetAuthAPIV1KeyIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get auth API v1 key ID params
func (o *GetAuthAPIV1KeyIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get auth API v1 key ID params
func (o *GetAuthAPIV1KeyIDParams) WithContext(ctx context.Context) *GetAuthAPIV1KeyIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get auth API v1 key ID params
func (o *GetAuthAPIV1KeyIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get auth API v1 key ID params
func (o *GetAuthAPIV1KeyIDParams) WithHTTPClient(client *http.Client) *GetAuthAPIV1KeyIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get auth API v1 key ID params
func (o *GetAuthAPIV1KeyIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get auth API v1 key ID params
func (o *GetAuthAPIV1KeyIDParams) WithID(id string) *GetAuthAPIV1KeyIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get auth API v1 key ID params
func (o *GetAuthAPIV1KeyIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthAPIV1KeyIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
