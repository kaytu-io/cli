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

// NewGetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams creates a new GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams() *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams {
	return &GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParamsWithTimeout creates a new GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams object
// with the ability to set a timeout on a request.
func NewGetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParamsWithTimeout(timeout time.Duration) *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams {
	return &GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams{
		timeout: timeout,
	}
}

// NewGetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParamsWithContext creates a new GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams object
// with the ability to set a context for a request.
func NewGetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParamsWithContext(ctx context.Context) *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams {
	return &GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams{
		Context: ctx,
	}
}

// NewGetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParamsWithHTTPClient creates a new GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParamsWithHTTPClient(client *http.Client) *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams {
	return &GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams{
		HTTPClient: client,
	}
}

/*
GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams contains all the parameters to send to the API endpoint

	for the get workspace API v1 workspaces limits byid workspace ID operation.

	Typically these are written to a http.Request.
*/
type GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams struct {

	/* WorkspaceID.

	   Workspace Name
	*/
	WorkspaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get workspace API v1 workspaces limits byid workspace ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams) WithDefaults() *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get workspace API v1 workspaces limits byid workspace ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get workspace API v1 workspaces limits byid workspace ID params
func (o *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams) WithTimeout(timeout time.Duration) *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workspace API v1 workspaces limits byid workspace ID params
func (o *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workspace API v1 workspaces limits byid workspace ID params
func (o *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams) WithContext(ctx context.Context) *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workspace API v1 workspaces limits byid workspace ID params
func (o *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workspace API v1 workspaces limits byid workspace ID params
func (o *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams) WithHTTPClient(client *http.Client) *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workspace API v1 workspaces limits byid workspace ID params
func (o *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWorkspaceID adds the workspaceID to the get workspace API v1 workspaces limits byid workspace ID params
func (o *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams) WithWorkspaceID(workspaceID string) *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get workspace API v1 workspaces limits byid workspace ID params
func (o *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams) SetWorkspaceID(workspaceID string) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
