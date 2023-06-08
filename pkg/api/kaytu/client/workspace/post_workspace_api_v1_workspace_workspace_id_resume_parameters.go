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

// NewPostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams creates a new PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams() *PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams {
	return &PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParamsWithTimeout creates a new PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams object
// with the ability to set a timeout on a request.
func NewPostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParamsWithTimeout(timeout time.Duration) *PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams {
	return &PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams{
		timeout: timeout,
	}
}

// NewPostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParamsWithContext creates a new PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams object
// with the ability to set a context for a request.
func NewPostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParamsWithContext(ctx context.Context) *PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams {
	return &PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams{
		Context: ctx,
	}
}

// NewPostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParamsWithHTTPClient creates a new PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParamsWithHTTPClient(client *http.Client) *PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams {
	return &PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams{
		HTTPClient: client,
	}
}

/*
PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams contains all the parameters to send to the API endpoint

	for the post workspace API v1 workspace workspace ID resume operation.

	Typically these are written to a http.Request.
*/
type PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams struct {

	/* WorkspaceID.

	   Workspace ID
	*/
	WorkspaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post workspace API v1 workspace workspace ID resume params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams) WithDefaults() *PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post workspace API v1 workspace workspace ID resume params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post workspace API v1 workspace workspace ID resume params
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams) WithTimeout(timeout time.Duration) *PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post workspace API v1 workspace workspace ID resume params
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post workspace API v1 workspace workspace ID resume params
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams) WithContext(ctx context.Context) *PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post workspace API v1 workspace workspace ID resume params
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post workspace API v1 workspace workspace ID resume params
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams) WithHTTPClient(client *http.Client) *PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post workspace API v1 workspace workspace ID resume params
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWorkspaceID adds the workspaceID to the post workspace API v1 workspace workspace ID resume params
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams) WithWorkspaceID(workspaceID string) *PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the post workspace API v1 workspace workspace ID resume params
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams) SetWorkspaceID(workspaceID string) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkspaceAPIV1WorkspaceWorkspaceIDResumeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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