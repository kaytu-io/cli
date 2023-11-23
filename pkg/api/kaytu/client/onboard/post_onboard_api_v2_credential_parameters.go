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

	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

// NewPostOnboardAPIV2CredentialParams creates a new PostOnboardAPIV2CredentialParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostOnboardAPIV2CredentialParams() *PostOnboardAPIV2CredentialParams {
	return &PostOnboardAPIV2CredentialParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostOnboardAPIV2CredentialParamsWithTimeout creates a new PostOnboardAPIV2CredentialParams object
// with the ability to set a timeout on a request.
func NewPostOnboardAPIV2CredentialParamsWithTimeout(timeout time.Duration) *PostOnboardAPIV2CredentialParams {
	return &PostOnboardAPIV2CredentialParams{
		timeout: timeout,
	}
}

// NewPostOnboardAPIV2CredentialParamsWithContext creates a new PostOnboardAPIV2CredentialParams object
// with the ability to set a context for a request.
func NewPostOnboardAPIV2CredentialParamsWithContext(ctx context.Context) *PostOnboardAPIV2CredentialParams {
	return &PostOnboardAPIV2CredentialParams{
		Context: ctx,
	}
}

// NewPostOnboardAPIV2CredentialParamsWithHTTPClient creates a new PostOnboardAPIV2CredentialParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostOnboardAPIV2CredentialParamsWithHTTPClient(client *http.Client) *PostOnboardAPIV2CredentialParams {
	return &PostOnboardAPIV2CredentialParams{
		HTTPClient: client,
	}
}

/*
PostOnboardAPIV2CredentialParams contains all the parameters to send to the API endpoint

	for the post onboard API v2 credential operation.

	Typically these are written to a http.Request.
*/
type PostOnboardAPIV2CredentialParams struct {

	/* Config.

	   config
	*/
	Config *models.GithubComKaytuIoKaytuEnginePkgOnboardAPIV2CreateCredentialV2Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post onboard API v2 credential params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostOnboardAPIV2CredentialParams) WithDefaults() *PostOnboardAPIV2CredentialParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post onboard API v2 credential params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostOnboardAPIV2CredentialParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post onboard API v2 credential params
func (o *PostOnboardAPIV2CredentialParams) WithTimeout(timeout time.Duration) *PostOnboardAPIV2CredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post onboard API v2 credential params
func (o *PostOnboardAPIV2CredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post onboard API v2 credential params
func (o *PostOnboardAPIV2CredentialParams) WithContext(ctx context.Context) *PostOnboardAPIV2CredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post onboard API v2 credential params
func (o *PostOnboardAPIV2CredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post onboard API v2 credential params
func (o *PostOnboardAPIV2CredentialParams) WithHTTPClient(client *http.Client) *PostOnboardAPIV2CredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post onboard API v2 credential params
func (o *PostOnboardAPIV2CredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfig adds the config to the post onboard API v2 credential params
func (o *PostOnboardAPIV2CredentialParams) WithConfig(config *models.GithubComKaytuIoKaytuEnginePkgOnboardAPIV2CreateCredentialV2Request) *PostOnboardAPIV2CredentialParams {
	o.SetConfig(config)
	return o
}

// SetConfig adds the config to the post onboard API v2 credential params
func (o *PostOnboardAPIV2CredentialParams) SetConfig(config *models.GithubComKaytuIoKaytuEnginePkgOnboardAPIV2CreateCredentialV2Request) {
	o.Config = config
}

// WriteToRequest writes these params to a swagger request
func (o *PostOnboardAPIV2CredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Config != nil {
		if err := r.SetBodyParam(o.Config); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
