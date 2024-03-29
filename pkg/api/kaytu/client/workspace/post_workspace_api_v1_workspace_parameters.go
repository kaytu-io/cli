// Code generated by go-swagger; DO NOT EDIT.

package workspace

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

// NewPostWorkspaceAPIV1WorkspaceParams creates a new PostWorkspaceAPIV1WorkspaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostWorkspaceAPIV1WorkspaceParams() *PostWorkspaceAPIV1WorkspaceParams {
	return &PostWorkspaceAPIV1WorkspaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostWorkspaceAPIV1WorkspaceParamsWithTimeout creates a new PostWorkspaceAPIV1WorkspaceParams object
// with the ability to set a timeout on a request.
func NewPostWorkspaceAPIV1WorkspaceParamsWithTimeout(timeout time.Duration) *PostWorkspaceAPIV1WorkspaceParams {
	return &PostWorkspaceAPIV1WorkspaceParams{
		timeout: timeout,
	}
}

// NewPostWorkspaceAPIV1WorkspaceParamsWithContext creates a new PostWorkspaceAPIV1WorkspaceParams object
// with the ability to set a context for a request.
func NewPostWorkspaceAPIV1WorkspaceParamsWithContext(ctx context.Context) *PostWorkspaceAPIV1WorkspaceParams {
	return &PostWorkspaceAPIV1WorkspaceParams{
		Context: ctx,
	}
}

// NewPostWorkspaceAPIV1WorkspaceParamsWithHTTPClient creates a new PostWorkspaceAPIV1WorkspaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostWorkspaceAPIV1WorkspaceParamsWithHTTPClient(client *http.Client) *PostWorkspaceAPIV1WorkspaceParams {
	return &PostWorkspaceAPIV1WorkspaceParams{
		HTTPClient: client,
	}
}

/*
PostWorkspaceAPIV1WorkspaceParams contains all the parameters to send to the API endpoint

	for the post workspace API v1 workspace operation.

	Typically these are written to a http.Request.
*/
type PostWorkspaceAPIV1WorkspaceParams struct {

	/* Request.

	   Create workspace request
	*/
	Request *models.GithubComKaytuIoKaytuEnginePkgWorkspaceAPICreateWorkspaceRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post workspace API v1 workspace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkspaceAPIV1WorkspaceParams) WithDefaults() *PostWorkspaceAPIV1WorkspaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post workspace API v1 workspace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkspaceAPIV1WorkspaceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post workspace API v1 workspace params
func (o *PostWorkspaceAPIV1WorkspaceParams) WithTimeout(timeout time.Duration) *PostWorkspaceAPIV1WorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post workspace API v1 workspace params
func (o *PostWorkspaceAPIV1WorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post workspace API v1 workspace params
func (o *PostWorkspaceAPIV1WorkspaceParams) WithContext(ctx context.Context) *PostWorkspaceAPIV1WorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post workspace API v1 workspace params
func (o *PostWorkspaceAPIV1WorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post workspace API v1 workspace params
func (o *PostWorkspaceAPIV1WorkspaceParams) WithHTTPClient(client *http.Client) *PostWorkspaceAPIV1WorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post workspace API v1 workspace params
func (o *PostWorkspaceAPIV1WorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post workspace API v1 workspace params
func (o *PostWorkspaceAPIV1WorkspaceParams) WithRequest(request *models.GithubComKaytuIoKaytuEnginePkgWorkspaceAPICreateWorkspaceRequest) *PostWorkspaceAPIV1WorkspaceParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post workspace API v1 workspace params
func (o *PostWorkspaceAPIV1WorkspaceParams) SetRequest(request *models.GithubComKaytuIoKaytuEnginePkgWorkspaceAPICreateWorkspaceRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkspaceAPIV1WorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
