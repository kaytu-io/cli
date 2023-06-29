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
	"github.com/go-openapi/swag"
)

// NewGetOnboardAPIV1SourceSourceIDParams creates a new GetOnboardAPIV1SourceSourceIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOnboardAPIV1SourceSourceIDParams() *GetOnboardAPIV1SourceSourceIDParams {
	return &GetOnboardAPIV1SourceSourceIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOnboardAPIV1SourceSourceIDParamsWithTimeout creates a new GetOnboardAPIV1SourceSourceIDParams object
// with the ability to set a timeout on a request.
func NewGetOnboardAPIV1SourceSourceIDParamsWithTimeout(timeout time.Duration) *GetOnboardAPIV1SourceSourceIDParams {
	return &GetOnboardAPIV1SourceSourceIDParams{
		timeout: timeout,
	}
}

// NewGetOnboardAPIV1SourceSourceIDParamsWithContext creates a new GetOnboardAPIV1SourceSourceIDParams object
// with the ability to set a context for a request.
func NewGetOnboardAPIV1SourceSourceIDParamsWithContext(ctx context.Context) *GetOnboardAPIV1SourceSourceIDParams {
	return &GetOnboardAPIV1SourceSourceIDParams{
		Context: ctx,
	}
}

// NewGetOnboardAPIV1SourceSourceIDParamsWithHTTPClient creates a new GetOnboardAPIV1SourceSourceIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOnboardAPIV1SourceSourceIDParamsWithHTTPClient(client *http.Client) *GetOnboardAPIV1SourceSourceIDParams {
	return &GetOnboardAPIV1SourceSourceIDParams{
		HTTPClient: client,
	}
}

/*
GetOnboardAPIV1SourceSourceIDParams contains all the parameters to send to the API endpoint

	for the get onboard API v1 source source ID operation.

	Typically these are written to a http.Request.
*/
type GetOnboardAPIV1SourceSourceIDParams struct {

	/* SourceID.

	   Source ID
	*/
	SourceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get onboard API v1 source source ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOnboardAPIV1SourceSourceIDParams) WithDefaults() *GetOnboardAPIV1SourceSourceIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get onboard API v1 source source ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOnboardAPIV1SourceSourceIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get onboard API v1 source source ID params
func (o *GetOnboardAPIV1SourceSourceIDParams) WithTimeout(timeout time.Duration) *GetOnboardAPIV1SourceSourceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get onboard API v1 source source ID params
func (o *GetOnboardAPIV1SourceSourceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get onboard API v1 source source ID params
func (o *GetOnboardAPIV1SourceSourceIDParams) WithContext(ctx context.Context) *GetOnboardAPIV1SourceSourceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get onboard API v1 source source ID params
func (o *GetOnboardAPIV1SourceSourceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get onboard API v1 source source ID params
func (o *GetOnboardAPIV1SourceSourceIDParams) WithHTTPClient(client *http.Client) *GetOnboardAPIV1SourceSourceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get onboard API v1 source source ID params
func (o *GetOnboardAPIV1SourceSourceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSourceID adds the sourceID to the get onboard API v1 source source ID params
func (o *GetOnboardAPIV1SourceSourceIDParams) WithSourceID(sourceID int64) *GetOnboardAPIV1SourceSourceIDParams {
	o.SetSourceID(sourceID)
	return o
}

// SetSourceID adds the sourceId to the get onboard API v1 source source ID params
func (o *GetOnboardAPIV1SourceSourceIDParams) SetSourceID(sourceID int64) {
	o.SourceID = sourceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOnboardAPIV1SourceSourceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param sourceId
	if err := r.SetPathParam("sourceId", swag.FormatInt64(o.SourceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
