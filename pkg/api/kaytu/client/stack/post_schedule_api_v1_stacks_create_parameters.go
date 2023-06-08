// Code generated by go-swagger; DO NOT EDIT.

package stack

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

// NewPostScheduleAPIV1StacksCreateParams creates a new PostScheduleAPIV1StacksCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostScheduleAPIV1StacksCreateParams() *PostScheduleAPIV1StacksCreateParams {
	return &PostScheduleAPIV1StacksCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostScheduleAPIV1StacksCreateParamsWithTimeout creates a new PostScheduleAPIV1StacksCreateParams object
// with the ability to set a timeout on a request.
func NewPostScheduleAPIV1StacksCreateParamsWithTimeout(timeout time.Duration) *PostScheduleAPIV1StacksCreateParams {
	return &PostScheduleAPIV1StacksCreateParams{
		timeout: timeout,
	}
}

// NewPostScheduleAPIV1StacksCreateParamsWithContext creates a new PostScheduleAPIV1StacksCreateParams object
// with the ability to set a context for a request.
func NewPostScheduleAPIV1StacksCreateParamsWithContext(ctx context.Context) *PostScheduleAPIV1StacksCreateParams {
	return &PostScheduleAPIV1StacksCreateParams{
		Context: ctx,
	}
}

// NewPostScheduleAPIV1StacksCreateParamsWithHTTPClient creates a new PostScheduleAPIV1StacksCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostScheduleAPIV1StacksCreateParamsWithHTTPClient(client *http.Client) *PostScheduleAPIV1StacksCreateParams {
	return &PostScheduleAPIV1StacksCreateParams{
		HTTPClient: client,
	}
}

/*
PostScheduleAPIV1StacksCreateParams contains all the parameters to send to the API endpoint

	for the post schedule API v1 stacks create operation.

	Typically these are written to a http.Request.
*/
type PostScheduleAPIV1StacksCreateParams struct {

	/* Request.

	   Request Body
	*/
	Request *models.GitlabComKeibiengineKeibiEnginePkgDescribeAPICreateStackRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post schedule API v1 stacks create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostScheduleAPIV1StacksCreateParams) WithDefaults() *PostScheduleAPIV1StacksCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post schedule API v1 stacks create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostScheduleAPIV1StacksCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post schedule API v1 stacks create params
func (o *PostScheduleAPIV1StacksCreateParams) WithTimeout(timeout time.Duration) *PostScheduleAPIV1StacksCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post schedule API v1 stacks create params
func (o *PostScheduleAPIV1StacksCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post schedule API v1 stacks create params
func (o *PostScheduleAPIV1StacksCreateParams) WithContext(ctx context.Context) *PostScheduleAPIV1StacksCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post schedule API v1 stacks create params
func (o *PostScheduleAPIV1StacksCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post schedule API v1 stacks create params
func (o *PostScheduleAPIV1StacksCreateParams) WithHTTPClient(client *http.Client) *PostScheduleAPIV1StacksCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post schedule API v1 stacks create params
func (o *PostScheduleAPIV1StacksCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post schedule API v1 stacks create params
func (o *PostScheduleAPIV1StacksCreateParams) WithRequest(request *models.GitlabComKeibiengineKeibiEnginePkgDescribeAPICreateStackRequest) *PostScheduleAPIV1StacksCreateParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post schedule API v1 stacks create params
func (o *PostScheduleAPIV1StacksCreateParams) SetRequest(request *models.GitlabComKeibiengineKeibiEnginePkgDescribeAPICreateStackRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostScheduleAPIV1StacksCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
