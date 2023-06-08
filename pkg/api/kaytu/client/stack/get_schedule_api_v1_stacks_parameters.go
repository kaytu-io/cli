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

// NewGetScheduleAPIV1StacksParams creates a new GetScheduleAPIV1StacksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetScheduleAPIV1StacksParams() *GetScheduleAPIV1StacksParams {
	return &GetScheduleAPIV1StacksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetScheduleAPIV1StacksParamsWithTimeout creates a new GetScheduleAPIV1StacksParams object
// with the ability to set a timeout on a request.
func NewGetScheduleAPIV1StacksParamsWithTimeout(timeout time.Duration) *GetScheduleAPIV1StacksParams {
	return &GetScheduleAPIV1StacksParams{
		timeout: timeout,
	}
}

// NewGetScheduleAPIV1StacksParamsWithContext creates a new GetScheduleAPIV1StacksParams object
// with the ability to set a context for a request.
func NewGetScheduleAPIV1StacksParamsWithContext(ctx context.Context) *GetScheduleAPIV1StacksParams {
	return &GetScheduleAPIV1StacksParams{
		Context: ctx,
	}
}

// NewGetScheduleAPIV1StacksParamsWithHTTPClient creates a new GetScheduleAPIV1StacksParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetScheduleAPIV1StacksParamsWithHTTPClient(client *http.Client) *GetScheduleAPIV1StacksParams {
	return &GetScheduleAPIV1StacksParams{
		HTTPClient: client,
	}
}

/*
GetScheduleAPIV1StacksParams contains all the parameters to send to the API endpoint

	for the get schedule API v1 stacks operation.

	Typically these are written to a http.Request.
*/
type GetScheduleAPIV1StacksParams struct {

	/* AccounIds.

	   Account IDs to filter by
	*/
	AccounIds []string

	/* Tag.

	   Key-Value tags in key=value format to filter by
	*/
	Tag []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get schedule API v1 stacks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScheduleAPIV1StacksParams) WithDefaults() *GetScheduleAPIV1StacksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get schedule API v1 stacks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScheduleAPIV1StacksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get schedule API v1 stacks params
func (o *GetScheduleAPIV1StacksParams) WithTimeout(timeout time.Duration) *GetScheduleAPIV1StacksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get schedule API v1 stacks params
func (o *GetScheduleAPIV1StacksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get schedule API v1 stacks params
func (o *GetScheduleAPIV1StacksParams) WithContext(ctx context.Context) *GetScheduleAPIV1StacksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get schedule API v1 stacks params
func (o *GetScheduleAPIV1StacksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get schedule API v1 stacks params
func (o *GetScheduleAPIV1StacksParams) WithHTTPClient(client *http.Client) *GetScheduleAPIV1StacksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get schedule API v1 stacks params
func (o *GetScheduleAPIV1StacksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccounIds adds the accounIds to the get schedule API v1 stacks params
func (o *GetScheduleAPIV1StacksParams) WithAccounIds(accounIds []string) *GetScheduleAPIV1StacksParams {
	o.SetAccounIds(accounIds)
	return o
}

// SetAccounIds adds the accounIds to the get schedule API v1 stacks params
func (o *GetScheduleAPIV1StacksParams) SetAccounIds(accounIds []string) {
	o.AccounIds = accounIds
}

// WithTag adds the tag to the get schedule API v1 stacks params
func (o *GetScheduleAPIV1StacksParams) WithTag(tag []string) *GetScheduleAPIV1StacksParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the get schedule API v1 stacks params
func (o *GetScheduleAPIV1StacksParams) SetTag(tag []string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *GetScheduleAPIV1StacksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccounIds != nil {

		// binding items for accounIds
		joinedAccounIds := o.bindParamAccounIds(reg)

		// query array param accounIds
		if err := r.SetQueryParam("accounIds", joinedAccounIds...); err != nil {
			return err
		}
	}

	if o.Tag != nil {

		// binding items for tag
		joinedTag := o.bindParamTag(reg)

		// query array param tag
		if err := r.SetQueryParam("tag", joinedTag...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetScheduleAPIV1Stacks binds the parameter accounIds
func (o *GetScheduleAPIV1StacksParams) bindParamAccounIds(formats strfmt.Registry) []string {
	accounIdsIR := o.AccounIds

	var accounIdsIC []string
	for _, accounIdsIIR := range accounIdsIR { // explode []string

		accounIdsIIV := accounIdsIIR // string as string
		accounIdsIC = append(accounIdsIC, accounIdsIIV)
	}

	// items.CollectionFormat: ""
	accounIdsIS := swag.JoinByFormat(accounIdsIC, "")

	return accounIdsIS
}

// bindParamGetScheduleAPIV1Stacks binds the parameter tag
func (o *GetScheduleAPIV1StacksParams) bindParamTag(formats strfmt.Registry) []string {
	tagIR := o.Tag

	var tagIC []string
	for _, tagIIR := range tagIR { // explode []string

		tagIIV := tagIIR // string as string
		tagIC = append(tagIC, tagIIV)
	}

	// items.CollectionFormat: ""
	tagIS := swag.JoinByFormat(tagIC, "")

	return tagIS
}
