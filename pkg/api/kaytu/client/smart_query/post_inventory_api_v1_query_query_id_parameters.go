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

// NewPostInventoryAPIV1QueryQueryIDParams creates a new PostInventoryAPIV1QueryQueryIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostInventoryAPIV1QueryQueryIDParams() *PostInventoryAPIV1QueryQueryIDParams {
	return &PostInventoryAPIV1QueryQueryIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostInventoryAPIV1QueryQueryIDParamsWithTimeout creates a new PostInventoryAPIV1QueryQueryIDParams object
// with the ability to set a timeout on a request.
func NewPostInventoryAPIV1QueryQueryIDParamsWithTimeout(timeout time.Duration) *PostInventoryAPIV1QueryQueryIDParams {
	return &PostInventoryAPIV1QueryQueryIDParams{
		timeout: timeout,
	}
}

// NewPostInventoryAPIV1QueryQueryIDParamsWithContext creates a new PostInventoryAPIV1QueryQueryIDParams object
// with the ability to set a context for a request.
func NewPostInventoryAPIV1QueryQueryIDParamsWithContext(ctx context.Context) *PostInventoryAPIV1QueryQueryIDParams {
	return &PostInventoryAPIV1QueryQueryIDParams{
		Context: ctx,
	}
}

// NewPostInventoryAPIV1QueryQueryIDParamsWithHTTPClient creates a new PostInventoryAPIV1QueryQueryIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostInventoryAPIV1QueryQueryIDParamsWithHTTPClient(client *http.Client) *PostInventoryAPIV1QueryQueryIDParams {
	return &PostInventoryAPIV1QueryQueryIDParams{
		HTTPClient: client,
	}
}

/*
PostInventoryAPIV1QueryQueryIDParams contains all the parameters to send to the API endpoint

	for the post inventory API v1 query query ID operation.

	Typically these are written to a http.Request.
*/
type PostInventoryAPIV1QueryQueryIDParams struct {

	/* Accept.

	   Accept header
	*/
	Accept string

	/* QueryID.

	   QueryID
	*/
	QueryID string

	/* Request.

	   Request Body
	*/
	Request *models.GithubComKaytuIoKaytuEnginePkgInventoryAPIRunQueryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post inventory API v1 query query ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostInventoryAPIV1QueryQueryIDParams) WithDefaults() *PostInventoryAPIV1QueryQueryIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post inventory API v1 query query ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostInventoryAPIV1QueryQueryIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post inventory API v1 query query ID params
func (o *PostInventoryAPIV1QueryQueryIDParams) WithTimeout(timeout time.Duration) *PostInventoryAPIV1QueryQueryIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post inventory API v1 query query ID params
func (o *PostInventoryAPIV1QueryQueryIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post inventory API v1 query query ID params
func (o *PostInventoryAPIV1QueryQueryIDParams) WithContext(ctx context.Context) *PostInventoryAPIV1QueryQueryIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post inventory API v1 query query ID params
func (o *PostInventoryAPIV1QueryQueryIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post inventory API v1 query query ID params
func (o *PostInventoryAPIV1QueryQueryIDParams) WithHTTPClient(client *http.Client) *PostInventoryAPIV1QueryQueryIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post inventory API v1 query query ID params
func (o *PostInventoryAPIV1QueryQueryIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccept adds the accept to the post inventory API v1 query query ID params
func (o *PostInventoryAPIV1QueryQueryIDParams) WithAccept(accept string) *PostInventoryAPIV1QueryQueryIDParams {
	o.SetAccept(accept)
	return o
}

// SetAccept adds the accept to the post inventory API v1 query query ID params
func (o *PostInventoryAPIV1QueryQueryIDParams) SetAccept(accept string) {
	o.Accept = accept
}

// WithQueryID adds the queryID to the post inventory API v1 query query ID params
func (o *PostInventoryAPIV1QueryQueryIDParams) WithQueryID(queryID string) *PostInventoryAPIV1QueryQueryIDParams {
	o.SetQueryID(queryID)
	return o
}

// SetQueryID adds the queryId to the post inventory API v1 query query ID params
func (o *PostInventoryAPIV1QueryQueryIDParams) SetQueryID(queryID string) {
	o.QueryID = queryID
}

// WithRequest adds the request to the post inventory API v1 query query ID params
func (o *PostInventoryAPIV1QueryQueryIDParams) WithRequest(request *models.GithubComKaytuIoKaytuEnginePkgInventoryAPIRunQueryRequest) *PostInventoryAPIV1QueryQueryIDParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post inventory API v1 query query ID params
func (o *PostInventoryAPIV1QueryQueryIDParams) SetRequest(request *models.GithubComKaytuIoKaytuEnginePkgInventoryAPIRunQueryRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostInventoryAPIV1QueryQueryIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param accept
	if err := r.SetHeaderParam("accept", o.Accept); err != nil {
		return err
	}

	// path param queryId
	if err := r.SetPathParam("queryId", o.QueryID); err != nil {
		return err
	}
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
