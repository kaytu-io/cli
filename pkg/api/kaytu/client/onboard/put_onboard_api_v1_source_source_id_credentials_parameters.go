// Code generated by go-swagger; DO NOT EDIT.

package onboard

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

// NewPutOnboardAPIV1SourceSourceIDCredentialsParams creates a new PutOnboardAPIV1SourceSourceIDCredentialsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutOnboardAPIV1SourceSourceIDCredentialsParams() *PutOnboardAPIV1SourceSourceIDCredentialsParams {
	return &PutOnboardAPIV1SourceSourceIDCredentialsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutOnboardAPIV1SourceSourceIDCredentialsParamsWithTimeout creates a new PutOnboardAPIV1SourceSourceIDCredentialsParams object
// with the ability to set a timeout on a request.
func NewPutOnboardAPIV1SourceSourceIDCredentialsParamsWithTimeout(timeout time.Duration) *PutOnboardAPIV1SourceSourceIDCredentialsParams {
	return &PutOnboardAPIV1SourceSourceIDCredentialsParams{
		timeout: timeout,
	}
}

// NewPutOnboardAPIV1SourceSourceIDCredentialsParamsWithContext creates a new PutOnboardAPIV1SourceSourceIDCredentialsParams object
// with the ability to set a context for a request.
func NewPutOnboardAPIV1SourceSourceIDCredentialsParamsWithContext(ctx context.Context) *PutOnboardAPIV1SourceSourceIDCredentialsParams {
	return &PutOnboardAPIV1SourceSourceIDCredentialsParams{
		Context: ctx,
	}
}

// NewPutOnboardAPIV1SourceSourceIDCredentialsParamsWithHTTPClient creates a new PutOnboardAPIV1SourceSourceIDCredentialsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutOnboardAPIV1SourceSourceIDCredentialsParamsWithHTTPClient(client *http.Client) *PutOnboardAPIV1SourceSourceIDCredentialsParams {
	return &PutOnboardAPIV1SourceSourceIDCredentialsParams{
		HTTPClient: client,
	}
}

/*
PutOnboardAPIV1SourceSourceIDCredentialsParams contains all the parameters to send to the API endpoint

	for the put onboard API v1 source source ID credentials operation.

	Typically these are written to a http.Request.
*/
type PutOnboardAPIV1SourceSourceIDCredentialsParams struct {

	/* SourceID.

	   Source ID
	*/
	SourceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put onboard API v1 source source ID credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutOnboardAPIV1SourceSourceIDCredentialsParams) WithDefaults() *PutOnboardAPIV1SourceSourceIDCredentialsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put onboard API v1 source source ID credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutOnboardAPIV1SourceSourceIDCredentialsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put onboard API v1 source source ID credentials params
func (o *PutOnboardAPIV1SourceSourceIDCredentialsParams) WithTimeout(timeout time.Duration) *PutOnboardAPIV1SourceSourceIDCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put onboard API v1 source source ID credentials params
func (o *PutOnboardAPIV1SourceSourceIDCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put onboard API v1 source source ID credentials params
func (o *PutOnboardAPIV1SourceSourceIDCredentialsParams) WithContext(ctx context.Context) *PutOnboardAPIV1SourceSourceIDCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put onboard API v1 source source ID credentials params
func (o *PutOnboardAPIV1SourceSourceIDCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put onboard API v1 source source ID credentials params
func (o *PutOnboardAPIV1SourceSourceIDCredentialsParams) WithHTTPClient(client *http.Client) *PutOnboardAPIV1SourceSourceIDCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put onboard API v1 source source ID credentials params
func (o *PutOnboardAPIV1SourceSourceIDCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSourceID adds the sourceID to the put onboard API v1 source source ID credentials params
func (o *PutOnboardAPIV1SourceSourceIDCredentialsParams) WithSourceID(sourceID string) *PutOnboardAPIV1SourceSourceIDCredentialsParams {
	o.SetSourceID(sourceID)
	return o
}

// SetSourceID adds the sourceId to the put onboard API v1 source source ID credentials params
func (o *PutOnboardAPIV1SourceSourceIDCredentialsParams) SetSourceID(sourceID string) {
	o.SourceID = sourceID
}

// WriteToRequest writes these params to a swagger request
func (o *PutOnboardAPIV1SourceSourceIDCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param sourceId
	if err := r.SetPathParam("sourceId", o.SourceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}