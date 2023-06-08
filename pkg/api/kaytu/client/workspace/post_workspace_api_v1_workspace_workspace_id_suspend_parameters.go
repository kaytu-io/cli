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
)

// NewPostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams creates a new PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams() *PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams {
	return &PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParamsWithTimeout creates a new PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams object
// with the ability to set a timeout on a request.
func NewPostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParamsWithTimeout(timeout time.Duration) *PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams {
	return &PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams{
		timeout: timeout,
	}
}

// NewPostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParamsWithContext creates a new PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams object
// with the ability to set a context for a request.
func NewPostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParamsWithContext(ctx context.Context) *PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams {
	return &PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams{
		Context: ctx,
	}
}

// NewPostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParamsWithHTTPClient creates a new PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParamsWithHTTPClient(client *http.Client) *PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams {
	return &PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams{
		HTTPClient: client,
	}
}

/*
PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams contains all the parameters to send to the API endpoint

	for the post workspace API v1 workspace workspace ID suspend operation.

	Typically these are written to a http.Request.
*/
type PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams struct {

	/* WorkspaceID.

	   Workspace ID
	*/
	WorkspaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post workspace API v1 workspace workspace ID suspend params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams) WithDefaults() *PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post workspace API v1 workspace workspace ID suspend params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post workspace API v1 workspace workspace ID suspend params
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams) WithTimeout(timeout time.Duration) *PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post workspace API v1 workspace workspace ID suspend params
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post workspace API v1 workspace workspace ID suspend params
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams) WithContext(ctx context.Context) *PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post workspace API v1 workspace workspace ID suspend params
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post workspace API v1 workspace workspace ID suspend params
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams) WithHTTPClient(client *http.Client) *PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post workspace API v1 workspace workspace ID suspend params
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWorkspaceID adds the workspaceID to the post workspace API v1 workspace workspace ID suspend params
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams) WithWorkspaceID(workspaceID string) *PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the post workspace API v1 workspace workspace ID suspend params
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams) SetWorkspaceID(workspaceID string) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDSuspendParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param workspace_id
	if err := r.SetPathParam("workspace_id", o.WorkspaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
