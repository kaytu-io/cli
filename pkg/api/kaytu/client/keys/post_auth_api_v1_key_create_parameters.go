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

// NewPostAuthAPIV1KeyCreateParams creates a new PostAuthAPIV1KeyCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAuthAPIV1KeyCreateParams() *PostAuthAPIV1KeyCreateParams {
	return &PostAuthAPIV1KeyCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAuthAPIV1KeyCreateParamsWithTimeout creates a new PostAuthAPIV1KeyCreateParams object
// with the ability to set a timeout on a request.
func NewPostAuthAPIV1KeyCreateParamsWithTimeout(timeout time.Duration) *PostAuthAPIV1KeyCreateParams {
	return &PostAuthAPIV1KeyCreateParams{
		timeout: timeout,
	}
}

// NewPostAuthAPIV1KeyCreateParamsWithContext creates a new PostAuthAPIV1KeyCreateParams object
// with the ability to set a context for a request.
func NewPostAuthAPIV1KeyCreateParamsWithContext(ctx context.Context) *PostAuthAPIV1KeyCreateParams {
	return &PostAuthAPIV1KeyCreateParams{
		Context: ctx,
	}
}

// NewPostAuthAPIV1KeyCreateParamsWithHTTPClient creates a new PostAuthAPIV1KeyCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAuthAPIV1KeyCreateParamsWithHTTPClient(client *http.Client) *PostAuthAPIV1KeyCreateParams {
	return &PostAuthAPIV1KeyCreateParams{
		HTTPClient: client,
	}
}

/*
PostAuthAPIV1KeyCreateParams contains all the parameters to send to the API endpoint

	for the post auth API v1 key create operation.

	Typically these are written to a http.Request.
*/
type PostAuthAPIV1KeyCreateParams struct {

	/* Request.

	   Request Body
	*/
	Request *models.GitlabComKeibiengineKeibiEnginePkgAuthAPICreateAPIKeyRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post auth API v1 key create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAuthAPIV1KeyCreateParams) WithDefaults() *PostAuthAPIV1KeyCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post auth API v1 key create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAuthAPIV1KeyCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post auth API v1 key create params
func (o *PostAuthAPIV1KeyCreateParams) WithTimeout(timeout time.Duration) *PostAuthAPIV1KeyCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post auth API v1 key create params
func (o *PostAuthAPIV1KeyCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post auth API v1 key create params
func (o *PostAuthAPIV1KeyCreateParams) WithContext(ctx context.Context) *PostAuthAPIV1KeyCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post auth API v1 key create params
func (o *PostAuthAPIV1KeyCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post auth API v1 key create params
func (o *PostAuthAPIV1KeyCreateParams) WithHTTPClient(client *http.Client) *PostAuthAPIV1KeyCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post auth API v1 key create params
func (o *PostAuthAPIV1KeyCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post auth API v1 key create params
func (o *PostAuthAPIV1KeyCreateParams) WithRequest(request *models.GitlabComKeibiengineKeibiEnginePkgAuthAPICreateAPIKeyRequest) *PostAuthAPIV1KeyCreateParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post auth API v1 key create params
func (o *PostAuthAPIV1KeyCreateParams) SetRequest(request *models.GitlabComKeibiengineKeibiEnginePkgAuthAPICreateAPIKeyRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostAuthAPIV1KeyCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
