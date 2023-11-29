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

// NewPostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams creates a new PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams() *PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams {
	return &PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostWorkspaceAPIV1BootstrapWorkspaceNameFinishParamsWithTimeout creates a new PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams object
// with the ability to set a timeout on a request.
func NewPostWorkspaceAPIV1BootstrapWorkspaceNameFinishParamsWithTimeout(timeout time.Duration) *PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams {
	return &PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams{
		timeout: timeout,
	}
}

// NewPostWorkspaceAPIV1BootstrapWorkspaceNameFinishParamsWithContext creates a new PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams object
// with the ability to set a context for a request.
func NewPostWorkspaceAPIV1BootstrapWorkspaceNameFinishParamsWithContext(ctx context.Context) *PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams {
	return &PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams{
		Context: ctx,
	}
}

// NewPostWorkspaceAPIV1BootstrapWorkspaceNameFinishParamsWithHTTPClient creates a new PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostWorkspaceAPIV1BootstrapWorkspaceNameFinishParamsWithHTTPClient(client *http.Client) *PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams {
	return &PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams{
		HTTPClient: client,
	}
}

/*
PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams contains all the parameters to send to the API endpoint

	for the post workspace API v1 bootstrap workspace name finish operation.

	Typically these are written to a http.Request.
*/
type PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams struct {

	/* WorkspaceName.

	   Workspace Name
	*/
	WorkspaceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post workspace API v1 bootstrap workspace name finish params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams) WithDefaults() *PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post workspace API v1 bootstrap workspace name finish params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post workspace API v1 bootstrap workspace name finish params
func (o *PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams) WithTimeout(timeout time.Duration) *PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post workspace API v1 bootstrap workspace name finish params
func (o *PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post workspace API v1 bootstrap workspace name finish params
func (o *PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams) WithContext(ctx context.Context) *PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post workspace API v1 bootstrap workspace name finish params
func (o *PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post workspace API v1 bootstrap workspace name finish params
func (o *PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams) WithHTTPClient(client *http.Client) *PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post workspace API v1 bootstrap workspace name finish params
func (o *PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWorkspaceName adds the workspaceName to the post workspace API v1 bootstrap workspace name finish params
func (o *PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams) WithWorkspaceName(workspaceName string) *PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams {
	o.SetWorkspaceName(workspaceName)
	return o
}

// SetWorkspaceName adds the workspaceName to the post workspace API v1 bootstrap workspace name finish params
func (o *PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams) SetWorkspaceName(workspaceName string) {
	o.WorkspaceName = workspaceName
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkspaceAPIV1BootstrapWorkspaceNameFinishParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param workspace_name
	if err := r.SetPathParam("workspace_name", o.WorkspaceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}