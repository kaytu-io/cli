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

// NewPostOnboardAPIV1SourceSourceIDHealthcheckParams creates a new PostOnboardAPIV1SourceSourceIDHealthcheckParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostOnboardAPIV1SourceSourceIDHealthcheckParams() *PostOnboardAPIV1SourceSourceIDHealthcheckParams {
	return &PostOnboardAPIV1SourceSourceIDHealthcheckParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostOnboardAPIV1SourceSourceIDHealthcheckParamsWithTimeout creates a new PostOnboardAPIV1SourceSourceIDHealthcheckParams object
// with the ability to set a timeout on a request.
func NewPostOnboardAPIV1SourceSourceIDHealthcheckParamsWithTimeout(timeout time.Duration) *PostOnboardAPIV1SourceSourceIDHealthcheckParams {
	return &PostOnboardAPIV1SourceSourceIDHealthcheckParams{
		timeout: timeout,
	}
}

// NewPostOnboardAPIV1SourceSourceIDHealthcheckParamsWithContext creates a new PostOnboardAPIV1SourceSourceIDHealthcheckParams object
// with the ability to set a context for a request.
func NewPostOnboardAPIV1SourceSourceIDHealthcheckParamsWithContext(ctx context.Context) *PostOnboardAPIV1SourceSourceIDHealthcheckParams {
	return &PostOnboardAPIV1SourceSourceIDHealthcheckParams{
		Context: ctx,
	}
}

// NewPostOnboardAPIV1SourceSourceIDHealthcheckParamsWithHTTPClient creates a new PostOnboardAPIV1SourceSourceIDHealthcheckParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostOnboardAPIV1SourceSourceIDHealthcheckParamsWithHTTPClient(client *http.Client) *PostOnboardAPIV1SourceSourceIDHealthcheckParams {
	return &PostOnboardAPIV1SourceSourceIDHealthcheckParams{
		HTTPClient: client,
	}
}

/*
PostOnboardAPIV1SourceSourceIDHealthcheckParams contains all the parameters to send to the API endpoint

	for the post onboard API v1 source source ID healthcheck operation.

	Typically these are written to a http.Request.
*/
type PostOnboardAPIV1SourceSourceIDHealthcheckParams struct {

	/* SourceID.

	   Source ID
	*/
	SourceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post onboard API v1 source source ID healthcheck params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostOnboardAPIV1SourceSourceIDHealthcheckParams) WithDefaults() *PostOnboardAPIV1SourceSourceIDHealthcheckParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post onboard API v1 source source ID healthcheck params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostOnboardAPIV1SourceSourceIDHealthcheckParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post onboard API v1 source source ID healthcheck params
func (o *PostOnboardAPIV1SourceSourceIDHealthcheckParams) WithTimeout(timeout time.Duration) *PostOnboardAPIV1SourceSourceIDHealthcheckParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post onboard API v1 source source ID healthcheck params
func (o *PostOnboardAPIV1SourceSourceIDHealthcheckParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post onboard API v1 source source ID healthcheck params
func (o *PostOnboardAPIV1SourceSourceIDHealthcheckParams) WithContext(ctx context.Context) *PostOnboardAPIV1SourceSourceIDHealthcheckParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post onboard API v1 source source ID healthcheck params
func (o *PostOnboardAPIV1SourceSourceIDHealthcheckParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post onboard API v1 source source ID healthcheck params
func (o *PostOnboardAPIV1SourceSourceIDHealthcheckParams) WithHTTPClient(client *http.Client) *PostOnboardAPIV1SourceSourceIDHealthcheckParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post onboard API v1 source source ID healthcheck params
func (o *PostOnboardAPIV1SourceSourceIDHealthcheckParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSourceID adds the sourceID to the post onboard API v1 source source ID healthcheck params
func (o *PostOnboardAPIV1SourceSourceIDHealthcheckParams) WithSourceID(sourceID string) *PostOnboardAPIV1SourceSourceIDHealthcheckParams {
	o.SetSourceID(sourceID)
	return o
}

// SetSourceID adds the sourceId to the post onboard API v1 source source ID healthcheck params
func (o *PostOnboardAPIV1SourceSourceIDHealthcheckParams) SetSourceID(sourceID string) {
	o.SourceID = sourceID
}

// WriteToRequest writes these params to a swagger request
func (o *PostOnboardAPIV1SourceSourceIDHealthcheckParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
