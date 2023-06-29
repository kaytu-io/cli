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

	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

// NewPostAuthAPIV1KeyRoleParams creates a new PostAuthAPIV1KeyRoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAuthAPIV1KeyRoleParams() *PostAuthAPIV1KeyRoleParams {
	return &PostAuthAPIV1KeyRoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAuthAPIV1KeyRoleParamsWithTimeout creates a new PostAuthAPIV1KeyRoleParams object
// with the ability to set a timeout on a request.
func NewPostAuthAPIV1KeyRoleParamsWithTimeout(timeout time.Duration) *PostAuthAPIV1KeyRoleParams {
	return &PostAuthAPIV1KeyRoleParams{
		timeout: timeout,
	}
}

// NewPostAuthAPIV1KeyRoleParamsWithContext creates a new PostAuthAPIV1KeyRoleParams object
// with the ability to set a context for a request.
func NewPostAuthAPIV1KeyRoleParamsWithContext(ctx context.Context) *PostAuthAPIV1KeyRoleParams {
	return &PostAuthAPIV1KeyRoleParams{
		Context: ctx,
	}
}

// NewPostAuthAPIV1KeyRoleParamsWithHTTPClient creates a new PostAuthAPIV1KeyRoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAuthAPIV1KeyRoleParamsWithHTTPClient(client *http.Client) *PostAuthAPIV1KeyRoleParams {
	return &PostAuthAPIV1KeyRoleParams{
		HTTPClient: client,
	}
}

/*
PostAuthAPIV1KeyRoleParams contains all the parameters to send to the API endpoint

	for the post auth API v1 key role operation.

	Typically these are written to a http.Request.
*/
type PostAuthAPIV1KeyRoleParams struct {

	/* Request.

	   Request Body
	*/
	Request *models.GithubComKaytuIoKaytuEnginePkgAuthAPIUpdateKeyRoleRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post auth API v1 key role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAuthAPIV1KeyRoleParams) WithDefaults() *PostAuthAPIV1KeyRoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post auth API v1 key role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAuthAPIV1KeyRoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post auth API v1 key role params
func (o *PostAuthAPIV1KeyRoleParams) WithTimeout(timeout time.Duration) *PostAuthAPIV1KeyRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post auth API v1 key role params
func (o *PostAuthAPIV1KeyRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post auth API v1 key role params
func (o *PostAuthAPIV1KeyRoleParams) WithContext(ctx context.Context) *PostAuthAPIV1KeyRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post auth API v1 key role params
func (o *PostAuthAPIV1KeyRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post auth API v1 key role params
func (o *PostAuthAPIV1KeyRoleParams) WithHTTPClient(client *http.Client) *PostAuthAPIV1KeyRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post auth API v1 key role params
func (o *PostAuthAPIV1KeyRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post auth API v1 key role params
func (o *PostAuthAPIV1KeyRoleParams) WithRequest(request *models.GithubComKaytuIoKaytuEnginePkgAuthAPIUpdateKeyRoleRequest) *PostAuthAPIV1KeyRoleParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post auth API v1 key role params
func (o *PostAuthAPIV1KeyRoleParams) SetRequest(request *models.GithubComKaytuIoKaytuEnginePkgAuthAPIUpdateKeyRoleRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostAuthAPIV1KeyRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
