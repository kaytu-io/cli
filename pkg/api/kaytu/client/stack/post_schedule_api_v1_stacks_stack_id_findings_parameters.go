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

// NewPostScheduleAPIV1StacksStackIDFindingsParams creates a new PostScheduleAPIV1StacksStackIDFindingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostScheduleAPIV1StacksStackIDFindingsParams() *PostScheduleAPIV1StacksStackIDFindingsParams {
	return &PostScheduleAPIV1StacksStackIDFindingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostScheduleAPIV1StacksStackIDFindingsParamsWithTimeout creates a new PostScheduleAPIV1StacksStackIDFindingsParams object
// with the ability to set a timeout on a request.
func NewPostScheduleAPIV1StacksStackIDFindingsParamsWithTimeout(timeout time.Duration) *PostScheduleAPIV1StacksStackIDFindingsParams {
	return &PostScheduleAPIV1StacksStackIDFindingsParams{
		timeout: timeout,
	}
}

// NewPostScheduleAPIV1StacksStackIDFindingsParamsWithContext creates a new PostScheduleAPIV1StacksStackIDFindingsParams object
// with the ability to set a context for a request.
func NewPostScheduleAPIV1StacksStackIDFindingsParamsWithContext(ctx context.Context) *PostScheduleAPIV1StacksStackIDFindingsParams {
	return &PostScheduleAPIV1StacksStackIDFindingsParams{
		Context: ctx,
	}
}

// NewPostScheduleAPIV1StacksStackIDFindingsParamsWithHTTPClient creates a new PostScheduleAPIV1StacksStackIDFindingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostScheduleAPIV1StacksStackIDFindingsParamsWithHTTPClient(client *http.Client) *PostScheduleAPIV1StacksStackIDFindingsParams {
	return &PostScheduleAPIV1StacksStackIDFindingsParams{
		HTTPClient: client,
	}
}

/*
PostScheduleAPIV1StacksStackIDFindingsParams contains all the parameters to send to the API endpoint

	for the post schedule API v1 stacks stack ID findings operation.

	Typically these are written to a http.Request.
*/
type PostScheduleAPIV1StacksStackIDFindingsParams struct {

	/* Request.

	   Request Body
	*/
	Request *models.GithubComKaytuIoKaytuEnginePkgDescribeAPIGetStackFindings

	/* StackID.

	   StackId
	*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post schedule API v1 stacks stack ID findings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostScheduleAPIV1StacksStackIDFindingsParams) WithDefaults() *PostScheduleAPIV1StacksStackIDFindingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post schedule API v1 stacks stack ID findings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostScheduleAPIV1StacksStackIDFindingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post schedule API v1 stacks stack ID findings params
func (o *PostScheduleAPIV1StacksStackIDFindingsParams) WithTimeout(timeout time.Duration) *PostScheduleAPIV1StacksStackIDFindingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post schedule API v1 stacks stack ID findings params
func (o *PostScheduleAPIV1StacksStackIDFindingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post schedule API v1 stacks stack ID findings params
func (o *PostScheduleAPIV1StacksStackIDFindingsParams) WithContext(ctx context.Context) *PostScheduleAPIV1StacksStackIDFindingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post schedule API v1 stacks stack ID findings params
func (o *PostScheduleAPIV1StacksStackIDFindingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post schedule API v1 stacks stack ID findings params
func (o *PostScheduleAPIV1StacksStackIDFindingsParams) WithHTTPClient(client *http.Client) *PostScheduleAPIV1StacksStackIDFindingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post schedule API v1 stacks stack ID findings params
func (o *PostScheduleAPIV1StacksStackIDFindingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post schedule API v1 stacks stack ID findings params
func (o *PostScheduleAPIV1StacksStackIDFindingsParams) WithRequest(request *models.GithubComKaytuIoKaytuEnginePkgDescribeAPIGetStackFindings) *PostScheduleAPIV1StacksStackIDFindingsParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post schedule API v1 stacks stack ID findings params
func (o *PostScheduleAPIV1StacksStackIDFindingsParams) SetRequest(request *models.GithubComKaytuIoKaytuEnginePkgDescribeAPIGetStackFindings) {
	o.Request = request
}

// WithStackID adds the stackID to the post schedule API v1 stacks stack ID findings params
func (o *PostScheduleAPIV1StacksStackIDFindingsParams) WithStackID(stackID string) *PostScheduleAPIV1StacksStackIDFindingsParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the post schedule API v1 stacks stack ID findings params
func (o *PostScheduleAPIV1StacksStackIDFindingsParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *PostScheduleAPIV1StacksStackIDFindingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	// path param stackId
	if err := r.SetPathParam("stackId", o.StackID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
