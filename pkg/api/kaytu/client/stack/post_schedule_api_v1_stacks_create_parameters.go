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

	/* Config.

	   Config json structure
	*/
	Config string

	/* Tag.

	   Tags Map[string][]string
	*/
	Tag *string

	/* TerraformFile.

	   ُTerraform StateFile full path
	*/
	TerraformFile runtime.NamedReadCloser

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

// WithConfig adds the config to the post schedule API v1 stacks create params
func (o *PostScheduleAPIV1StacksCreateParams) WithConfig(config string) *PostScheduleAPIV1StacksCreateParams {
	o.SetConfig(config)
	return o
}

// SetConfig adds the config to the post schedule API v1 stacks create params
func (o *PostScheduleAPIV1StacksCreateParams) SetConfig(config string) {
	o.Config = config
}

// WithTag adds the tag to the post schedule API v1 stacks create params
func (o *PostScheduleAPIV1StacksCreateParams) WithTag(tag *string) *PostScheduleAPIV1StacksCreateParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the post schedule API v1 stacks create params
func (o *PostScheduleAPIV1StacksCreateParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTerraformFile adds the terraformFile to the post schedule API v1 stacks create params
func (o *PostScheduleAPIV1StacksCreateParams) WithTerraformFile(terraformFile runtime.NamedReadCloser) *PostScheduleAPIV1StacksCreateParams {
	o.SetTerraformFile(terraformFile)
	return o
}

// SetTerraformFile adds the terraformFile to the post schedule API v1 stacks create params
func (o *PostScheduleAPIV1StacksCreateParams) SetTerraformFile(terraformFile runtime.NamedReadCloser) {
	o.TerraformFile = terraformFile
}

// WriteToRequest writes these params to a swagger request
func (o *PostScheduleAPIV1StacksCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form param config
	frConfig := o.Config
	fConfig := frConfig
	if fConfig != "" {
		if err := r.SetFormParam("config", fConfig); err != nil {
			return err
		}
	}

	if o.Tag != nil {

		// form param tag
		var frTag string
		if o.Tag != nil {
			frTag = *o.Tag
		}
		fTag := frTag
		if fTag != "" {
			if err := r.SetFormParam("tag", fTag); err != nil {
				return err
			}
		}
	}
	// form file param terraformFile
	if err := r.SetFileParam("terraformFile", o.TerraformFile); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
