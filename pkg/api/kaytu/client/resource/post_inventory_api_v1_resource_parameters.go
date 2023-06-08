// Code generated by go-swagger; DO NOT EDIT.

package resource

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

// NewPostInventoryAPIV1ResourceParams creates a new PostInventoryAPIV1ResourceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostInventoryAPIV1ResourceParams() *PostInventoryAPIV1ResourceParams {
	return &PostInventoryAPIV1ResourceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostInventoryAPIV1ResourceParamsWithTimeout creates a new PostInventoryAPIV1ResourceParams object
// with the ability to set a timeout on a request.
func NewPostInventoryAPIV1ResourceParamsWithTimeout(timeout time.Duration) *PostInventoryAPIV1ResourceParams {
	return &PostInventoryAPIV1ResourceParams{
		timeout: timeout,
	}
}

// NewPostInventoryAPIV1ResourceParamsWithContext creates a new PostInventoryAPIV1ResourceParams object
// with the ability to set a context for a request.
func NewPostInventoryAPIV1ResourceParamsWithContext(ctx context.Context) *PostInventoryAPIV1ResourceParams {
	return &PostInventoryAPIV1ResourceParams{
		Context: ctx,
	}
}

// NewPostInventoryAPIV1ResourceParamsWithHTTPClient creates a new PostInventoryAPIV1ResourceParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostInventoryAPIV1ResourceParamsWithHTTPClient(client *http.Client) *PostInventoryAPIV1ResourceParams {
	return &PostInventoryAPIV1ResourceParams{
		HTTPClient: client,
	}
}

/*
PostInventoryAPIV1ResourceParams contains all the parameters to send to the API endpoint

	for the post inventory API v1 resource operation.

	Typically these are written to a http.Request.
*/
type PostInventoryAPIV1ResourceParams struct {

	/* Request.

	   Request Body
	*/
	Request *models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetResourceRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post inventory API v1 resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostInventoryAPIV1ResourceParams) WithDefaults() *PostInventoryAPIV1ResourceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post inventory API v1 resource params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostInventoryAPIV1ResourceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post inventory API v1 resource params
func (o *PostInventoryAPIV1ResourceParams) WithTimeout(timeout time.Duration) *PostInventoryAPIV1ResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post inventory API v1 resource params
func (o *PostInventoryAPIV1ResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post inventory API v1 resource params
func (o *PostInventoryAPIV1ResourceParams) WithContext(ctx context.Context) *PostInventoryAPIV1ResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post inventory API v1 resource params
func (o *PostInventoryAPIV1ResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post inventory API v1 resource params
func (o *PostInventoryAPIV1ResourceParams) WithHTTPClient(client *http.Client) *PostInventoryAPIV1ResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post inventory API v1 resource params
func (o *PostInventoryAPIV1ResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post inventory API v1 resource params
func (o *PostInventoryAPIV1ResourceParams) WithRequest(request *models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetResourceRequest) *PostInventoryAPIV1ResourceParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post inventory API v1 resource params
func (o *PostInventoryAPIV1ResourceParams) SetRequest(request *models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetResourceRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostInventoryAPIV1ResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
