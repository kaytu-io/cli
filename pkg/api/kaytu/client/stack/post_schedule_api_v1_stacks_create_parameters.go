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
	"github.com/go-openapi/swag"
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

	/* Resources.

	   Additional Resources
	*/
	Resources []string

	/* Tags.

	   Tags Map[string][]string
	*/
	Tags *string

	/* TerrafromFile.

	   File to upload
	*/
	TerrafromFile runtime.NamedReadCloser

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

// WithResources adds the resources to the post schedule API v1 stacks create params
func (o *PostScheduleAPIV1StacksCreateParams) WithResources(resources []string) *PostScheduleAPIV1StacksCreateParams {
	o.SetResources(resources)
	return o
}

// SetResources adds the resources to the post schedule API v1 stacks create params
func (o *PostScheduleAPIV1StacksCreateParams) SetResources(resources []string) {
	o.Resources = resources
}

// WithTags adds the tags to the post schedule API v1 stacks create params
func (o *PostScheduleAPIV1StacksCreateParams) WithTags(tags *string) *PostScheduleAPIV1StacksCreateParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the post schedule API v1 stacks create params
func (o *PostScheduleAPIV1StacksCreateParams) SetTags(tags *string) {
	o.Tags = tags
}

// WithTerrafromFile adds the terrafromFile to the post schedule API v1 stacks create params
func (o *PostScheduleAPIV1StacksCreateParams) WithTerrafromFile(terrafromFile runtime.NamedReadCloser) *PostScheduleAPIV1StacksCreateParams {
	o.SetTerrafromFile(terrafromFile)
	return o
}

// SetTerrafromFile adds the terrafromFile to the post schedule API v1 stacks create params
func (o *PostScheduleAPIV1StacksCreateParams) SetTerrafromFile(terrafromFile runtime.NamedReadCloser) {
	o.TerrafromFile = terrafromFile
}

// WriteToRequest writes these params to a swagger request
func (o *PostScheduleAPIV1StacksCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Resources != nil {

		// binding items for resources
		joinedResources := o.bindParamResources(reg)

		// form array param resources
		if err := r.SetFormParam("resources", joinedResources...); err != nil {
			return err
		}
	}

	if o.Tags != nil {

		// form param tags
		var frTags string
		if o.Tags != nil {
			frTags = *o.Tags
		}
		fTags := frTags
		if fTags != "" {
			if err := r.SetFormParam("tags", fTags); err != nil {
				return err
			}
		}
	}

	if o.TerrafromFile != nil {

		if o.TerrafromFile != nil {
			// form file param terrafromFile
			if err := r.SetFileParam("terrafromFile", o.TerrafromFile); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamPostScheduleAPIV1StacksCreate binds the parameter resources
func (o *PostScheduleAPIV1StacksCreateParams) bindParamResources(formats strfmt.Registry) []string {
	resourcesIR := o.Resources

	var resourcesIC []string
	for _, resourcesIIR := range resourcesIR { // explode []string

		resourcesIIV := resourcesIIR // string as string
		resourcesIC = append(resourcesIC, resourcesIIV)
	}

	// items.CollectionFormat: ""
	resourcesIS := swag.JoinByFormat(resourcesIC, "")

	return resourcesIS
}