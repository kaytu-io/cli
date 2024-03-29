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

// NewPostOnboardAPIV1CredentialCredentialIDAutoonboardParams creates a new PostOnboardAPIV1CredentialCredentialIDAutoonboardParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostOnboardAPIV1CredentialCredentialIDAutoonboardParams() *PostOnboardAPIV1CredentialCredentialIDAutoonboardParams {
	return &PostOnboardAPIV1CredentialCredentialIDAutoonboardParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostOnboardAPIV1CredentialCredentialIDAutoonboardParamsWithTimeout creates a new PostOnboardAPIV1CredentialCredentialIDAutoonboardParams object
// with the ability to set a timeout on a request.
func NewPostOnboardAPIV1CredentialCredentialIDAutoonboardParamsWithTimeout(timeout time.Duration) *PostOnboardAPIV1CredentialCredentialIDAutoonboardParams {
	return &PostOnboardAPIV1CredentialCredentialIDAutoonboardParams{
		timeout: timeout,
	}
}

// NewPostOnboardAPIV1CredentialCredentialIDAutoonboardParamsWithContext creates a new PostOnboardAPIV1CredentialCredentialIDAutoonboardParams object
// with the ability to set a context for a request.
func NewPostOnboardAPIV1CredentialCredentialIDAutoonboardParamsWithContext(ctx context.Context) *PostOnboardAPIV1CredentialCredentialIDAutoonboardParams {
	return &PostOnboardAPIV1CredentialCredentialIDAutoonboardParams{
		Context: ctx,
	}
}

// NewPostOnboardAPIV1CredentialCredentialIDAutoonboardParamsWithHTTPClient creates a new PostOnboardAPIV1CredentialCredentialIDAutoonboardParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostOnboardAPIV1CredentialCredentialIDAutoonboardParamsWithHTTPClient(client *http.Client) *PostOnboardAPIV1CredentialCredentialIDAutoonboardParams {
	return &PostOnboardAPIV1CredentialCredentialIDAutoonboardParams{
		HTTPClient: client,
	}
}

/*
PostOnboardAPIV1CredentialCredentialIDAutoonboardParams contains all the parameters to send to the API endpoint

	for the post onboard API v1 credential credential ID autoonboard operation.

	Typically these are written to a http.Request.
*/
type PostOnboardAPIV1CredentialCredentialIDAutoonboardParams struct {

	/* CredentialID.

	   CredentialID
	*/
	CredentialID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post onboard API v1 credential credential ID autoonboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostOnboardAPIV1CredentialCredentialIDAutoonboardParams) WithDefaults() *PostOnboardAPIV1CredentialCredentialIDAutoonboardParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post onboard API v1 credential credential ID autoonboard params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostOnboardAPIV1CredentialCredentialIDAutoonboardParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post onboard API v1 credential credential ID autoonboard params
func (o *PostOnboardAPIV1CredentialCredentialIDAutoonboardParams) WithTimeout(timeout time.Duration) *PostOnboardAPIV1CredentialCredentialIDAutoonboardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post onboard API v1 credential credential ID autoonboard params
func (o *PostOnboardAPIV1CredentialCredentialIDAutoonboardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post onboard API v1 credential credential ID autoonboard params
func (o *PostOnboardAPIV1CredentialCredentialIDAutoonboardParams) WithContext(ctx context.Context) *PostOnboardAPIV1CredentialCredentialIDAutoonboardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post onboard API v1 credential credential ID autoonboard params
func (o *PostOnboardAPIV1CredentialCredentialIDAutoonboardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post onboard API v1 credential credential ID autoonboard params
func (o *PostOnboardAPIV1CredentialCredentialIDAutoonboardParams) WithHTTPClient(client *http.Client) *PostOnboardAPIV1CredentialCredentialIDAutoonboardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post onboard API v1 credential credential ID autoonboard params
func (o *PostOnboardAPIV1CredentialCredentialIDAutoonboardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredentialID adds the credentialID to the post onboard API v1 credential credential ID autoonboard params
func (o *PostOnboardAPIV1CredentialCredentialIDAutoonboardParams) WithCredentialID(credentialID string) *PostOnboardAPIV1CredentialCredentialIDAutoonboardParams {
	o.SetCredentialID(credentialID)
	return o
}

// SetCredentialID adds the credentialId to the post onboard API v1 credential credential ID autoonboard params
func (o *PostOnboardAPIV1CredentialCredentialIDAutoonboardParams) SetCredentialID(credentialID string) {
	o.CredentialID = credentialID
}

// WriteToRequest writes these params to a swagger request
func (o *PostOnboardAPIV1CredentialCredentialIDAutoonboardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
