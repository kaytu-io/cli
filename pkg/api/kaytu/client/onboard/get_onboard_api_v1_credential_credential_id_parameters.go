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

// NewGetOnboardAPIV1CredentialCredentialIDParams creates a new GetOnboardAPIV1CredentialCredentialIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOnboardAPIV1CredentialCredentialIDParams() *GetOnboardAPIV1CredentialCredentialIDParams {
	return &GetOnboardAPIV1CredentialCredentialIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOnboardAPIV1CredentialCredentialIDParamsWithTimeout creates a new GetOnboardAPIV1CredentialCredentialIDParams object
// with the ability to set a timeout on a request.
func NewGetOnboardAPIV1CredentialCredentialIDParamsWithTimeout(timeout time.Duration) *GetOnboardAPIV1CredentialCredentialIDParams {
	return &GetOnboardAPIV1CredentialCredentialIDParams{
		timeout: timeout,
	}
}

// NewGetOnboardAPIV1CredentialCredentialIDParamsWithContext creates a new GetOnboardAPIV1CredentialCredentialIDParams object
// with the ability to set a context for a request.
func NewGetOnboardAPIV1CredentialCredentialIDParamsWithContext(ctx context.Context) *GetOnboardAPIV1CredentialCredentialIDParams {
	return &GetOnboardAPIV1CredentialCredentialIDParams{
		Context: ctx,
	}
}

// NewGetOnboardAPIV1CredentialCredentialIDParamsWithHTTPClient creates a new GetOnboardAPIV1CredentialCredentialIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOnboardAPIV1CredentialCredentialIDParamsWithHTTPClient(client *http.Client) *GetOnboardAPIV1CredentialCredentialIDParams {
	return &GetOnboardAPIV1CredentialCredentialIDParams{
		HTTPClient: client,
	}
}

/*
GetOnboardAPIV1CredentialCredentialIDParams contains all the parameters to send to the API endpoint

	for the get onboard API v1 credential credential ID operation.

	Typically these are written to a http.Request.
*/
type GetOnboardAPIV1CredentialCredentialIDParams struct {

	/* CredentialID.

	   CredentialID
	*/
	CredentialID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get onboard API v1 credential credential ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOnboardAPIV1CredentialCredentialIDParams) WithDefaults() *GetOnboardAPIV1CredentialCredentialIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get onboard API v1 credential credential ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOnboardAPIV1CredentialCredentialIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get onboard API v1 credential credential ID params
func (o *GetOnboardAPIV1CredentialCredentialIDParams) WithTimeout(timeout time.Duration) *GetOnboardAPIV1CredentialCredentialIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get onboard API v1 credential credential ID params
func (o *GetOnboardAPIV1CredentialCredentialIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get onboard API v1 credential credential ID params
func (o *GetOnboardAPIV1CredentialCredentialIDParams) WithContext(ctx context.Context) *GetOnboardAPIV1CredentialCredentialIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get onboard API v1 credential credential ID params
func (o *GetOnboardAPIV1CredentialCredentialIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get onboard API v1 credential credential ID params
func (o *GetOnboardAPIV1CredentialCredentialIDParams) WithHTTPClient(client *http.Client) *GetOnboardAPIV1CredentialCredentialIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get onboard API v1 credential credential ID params
func (o *GetOnboardAPIV1CredentialCredentialIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredentialID adds the credentialID to the get onboard API v1 credential credential ID params
func (o *GetOnboardAPIV1CredentialCredentialIDParams) WithCredentialID(credentialID string) *GetOnboardAPIV1CredentialCredentialIDParams {
	o.SetCredentialID(credentialID)
	return o
}

// SetCredentialID adds the credentialId to the get onboard API v1 credential credential ID params
func (o *GetOnboardAPIV1CredentialCredentialIDParams) SetCredentialID(credentialID string) {
	o.CredentialID = credentialID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOnboardAPIV1CredentialCredentialIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param credentialId
	if err := r.SetPathParam("credentialId", o.CredentialID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
