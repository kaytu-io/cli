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

// NewPostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams creates a new PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams() *PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams {
	return &PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParamsWithTimeout creates a new PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams object
// with the ability to set a timeout on a request.
func NewPostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParamsWithTimeout(timeout time.Duration) *PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams {
	return &PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams{
		timeout: timeout,
	}
}

// NewPostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParamsWithContext creates a new PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams object
// with the ability to set a context for a request.
func NewPostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParamsWithContext(ctx context.Context) *PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams {
	return &PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams{
		Context: ctx,
	}
}

// NewPostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParamsWithHTTPClient creates a new PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParamsWithHTTPClient(client *http.Client) *PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams {
	return &PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams{
		HTTPClient: client,
	}
}

/*
PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams contains all the parameters to send to the API endpoint

	for the post workspace API v1 workspace workspace ID owner operation.

	Typically these are written to a http.Request.
*/
type PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams struct {

	/* Request.

	   Change ownership request
	*/
	Request *models.GithubComKaytuIoKaytuEnginePkgWorkspaceAPIChangeWorkspaceOwnershipRequest

	/* WorkspaceID.

	   WorkspaceID
	*/
	WorkspaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post workspace API v1 workspace workspace ID owner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams) WithDefaults() *PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post workspace API v1 workspace workspace ID owner params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post workspace API v1 workspace workspace ID owner params
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams) WithTimeout(timeout time.Duration) *PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post workspace API v1 workspace workspace ID owner params
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post workspace API v1 workspace workspace ID owner params
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams) WithContext(ctx context.Context) *PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post workspace API v1 workspace workspace ID owner params
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post workspace API v1 workspace workspace ID owner params
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams) WithHTTPClient(client *http.Client) *PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post workspace API v1 workspace workspace ID owner params
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post workspace API v1 workspace workspace ID owner params
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams) WithRequest(request *models.GithubComKaytuIoKaytuEnginePkgWorkspaceAPIChangeWorkspaceOwnershipRequest) *PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post workspace API v1 workspace workspace ID owner params
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams) SetRequest(request *models.GithubComKaytuIoKaytuEnginePkgWorkspaceAPIChangeWorkspaceOwnershipRequest) {
	o.Request = request
}

// WithWorkspaceID adds the workspaceID to the post workspace API v1 workspace workspace ID owner params
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams) WithWorkspaceID(workspaceID string) *PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the post workspace API v1 workspace workspace ID owner params
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams) SetWorkspaceID(workspaceID string) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDOwnerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	// path param workspace_id
	if err := r.SetPathParam("workspace_id", o.WorkspaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
