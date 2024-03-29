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
)

// NewPostAiAPIV1GptRunParams creates a new PostAiAPIV1GptRunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAiAPIV1GptRunParams() *PostAiAPIV1GptRunParams {
	return &PostAiAPIV1GptRunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAiAPIV1GptRunParamsWithTimeout creates a new PostAiAPIV1GptRunParams object
// with the ability to set a timeout on a request.
func NewPostAiAPIV1GptRunParamsWithTimeout(timeout time.Duration) *PostAiAPIV1GptRunParams {
	return &PostAiAPIV1GptRunParams{
		timeout: timeout,
	}
}

// NewPostAiAPIV1GptRunParamsWithContext creates a new PostAiAPIV1GptRunParams object
// with the ability to set a context for a request.
func NewPostAiAPIV1GptRunParamsWithContext(ctx context.Context) *PostAiAPIV1GptRunParams {
	return &PostAiAPIV1GptRunParams{
		Context: ctx,
	}
}

// NewPostAiAPIV1GptRunParamsWithHTTPClient creates a new PostAiAPIV1GptRunParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAiAPIV1GptRunParamsWithHTTPClient(client *http.Client) *PostAiAPIV1GptRunParams {
	return &PostAiAPIV1GptRunParams{
		HTTPClient: client,
	}
}

/*
PostAiAPIV1GptRunParams contains all the parameters to send to the API endpoint

	for the post ai API v1 gpt run operation.

	Typically these are written to a http.Request.
*/
type PostAiAPIV1GptRunParams struct {

	/* Query.

	   Description of query for KaytuGPT
	*/
	Query string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post ai API v1 gpt run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAiAPIV1GptRunParams) WithDefaults() *PostAiAPIV1GptRunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post ai API v1 gpt run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAiAPIV1GptRunParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post ai API v1 gpt run params
func (o *PostAiAPIV1GptRunParams) WithTimeout(timeout time.Duration) *PostAiAPIV1GptRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post ai API v1 gpt run params
func (o *PostAiAPIV1GptRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post ai API v1 gpt run params
func (o *PostAiAPIV1GptRunParams) WithContext(ctx context.Context) *PostAiAPIV1GptRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post ai API v1 gpt run params
func (o *PostAiAPIV1GptRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post ai API v1 gpt run params
func (o *PostAiAPIV1GptRunParams) WithHTTPClient(client *http.Client) *PostAiAPIV1GptRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post ai API v1 gpt run params
func (o *PostAiAPIV1GptRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuery adds the query to the post ai API v1 gpt run params
func (o *PostAiAPIV1GptRunParams) WithQuery(query string) *PostAiAPIV1GptRunParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the post ai API v1 gpt run params
func (o *PostAiAPIV1GptRunParams) SetQuery(query string) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *PostAiAPIV1GptRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Query); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
