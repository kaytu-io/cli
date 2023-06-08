// Code generated by go-swagger; DO NOT EDIT.

package smart_query

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

// NewGetInventoryAPIV1QueryCountParams creates a new GetInventoryAPIV1QueryCountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInventoryAPIV1QueryCountParams() *GetInventoryAPIV1QueryCountParams {
	return &GetInventoryAPIV1QueryCountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInventoryAPIV1QueryCountParamsWithTimeout creates a new GetInventoryAPIV1QueryCountParams object
// with the ability to set a timeout on a request.
func NewGetInventoryAPIV1QueryCountParamsWithTimeout(timeout time.Duration) *GetInventoryAPIV1QueryCountParams {
	return &GetInventoryAPIV1QueryCountParams{
		timeout: timeout,
	}
}

// NewGetInventoryAPIV1QueryCountParamsWithContext creates a new GetInventoryAPIV1QueryCountParams object
// with the ability to set a context for a request.
func NewGetInventoryAPIV1QueryCountParamsWithContext(ctx context.Context) *GetInventoryAPIV1QueryCountParams {
	return &GetInventoryAPIV1QueryCountParams{
		Context: ctx,
	}
}

// NewGetInventoryAPIV1QueryCountParamsWithHTTPClient creates a new GetInventoryAPIV1QueryCountParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInventoryAPIV1QueryCountParamsWithHTTPClient(client *http.Client) *GetInventoryAPIV1QueryCountParams {
	return &GetInventoryAPIV1QueryCountParams{
		HTTPClient: client,
	}
}

/*
GetInventoryAPIV1QueryCountParams contains all the parameters to send to the API endpoint

	for the get inventory API v1 query count operation.

	Typically these are written to a http.Request.
*/
type GetInventoryAPIV1QueryCountParams struct {

	/* Request.

	   Request Body
	*/
	Request *models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIListQueryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get inventory API v1 query count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInventoryAPIV1QueryCountParams) WithDefaults() *GetInventoryAPIV1QueryCountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get inventory API v1 query count params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInventoryAPIV1QueryCountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get inventory API v1 query count params
func (o *GetInventoryAPIV1QueryCountParams) WithTimeout(timeout time.Duration) *GetInventoryAPIV1QueryCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get inventory API v1 query count params
func (o *GetInventoryAPIV1QueryCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get inventory API v1 query count params
func (o *GetInventoryAPIV1QueryCountParams) WithContext(ctx context.Context) *GetInventoryAPIV1QueryCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get inventory API v1 query count params
func (o *GetInventoryAPIV1QueryCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get inventory API v1 query count params
func (o *GetInventoryAPIV1QueryCountParams) WithHTTPClient(client *http.Client) *GetInventoryAPIV1QueryCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get inventory API v1 query count params
func (o *GetInventoryAPIV1QueryCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the get inventory API v1 query count params
func (o *GetInventoryAPIV1QueryCountParams) WithRequest(request *models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIListQueryRequest) *GetInventoryAPIV1QueryCountParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the get inventory API v1 query count params
func (o *GetInventoryAPIV1QueryCountParams) SetRequest(request *models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIListQueryRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *GetInventoryAPIV1QueryCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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