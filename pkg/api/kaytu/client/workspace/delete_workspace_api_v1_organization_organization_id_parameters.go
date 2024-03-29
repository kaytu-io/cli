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
	"github.com/go-openapi/swag"
)

// NewDeleteWorkspaceAPIV1OrganizationOrganizationIDParams creates a new DeleteWorkspaceAPIV1OrganizationOrganizationIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteWorkspaceAPIV1OrganizationOrganizationIDParams() *DeleteWorkspaceAPIV1OrganizationOrganizationIDParams {
	return &DeleteWorkspaceAPIV1OrganizationOrganizationIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteWorkspaceAPIV1OrganizationOrganizationIDParamsWithTimeout creates a new DeleteWorkspaceAPIV1OrganizationOrganizationIDParams object
// with the ability to set a timeout on a request.
func NewDeleteWorkspaceAPIV1OrganizationOrganizationIDParamsWithTimeout(timeout time.Duration) *DeleteWorkspaceAPIV1OrganizationOrganizationIDParams {
	return &DeleteWorkspaceAPIV1OrganizationOrganizationIDParams{
		timeout: timeout,
	}
}

// NewDeleteWorkspaceAPIV1OrganizationOrganizationIDParamsWithContext creates a new DeleteWorkspaceAPIV1OrganizationOrganizationIDParams object
// with the ability to set a context for a request.
func NewDeleteWorkspaceAPIV1OrganizationOrganizationIDParamsWithContext(ctx context.Context) *DeleteWorkspaceAPIV1OrganizationOrganizationIDParams {
	return &DeleteWorkspaceAPIV1OrganizationOrganizationIDParams{
		Context: ctx,
	}
}

// NewDeleteWorkspaceAPIV1OrganizationOrganizationIDParamsWithHTTPClient creates a new DeleteWorkspaceAPIV1OrganizationOrganizationIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteWorkspaceAPIV1OrganizationOrganizationIDParamsWithHTTPClient(client *http.Client) *DeleteWorkspaceAPIV1OrganizationOrganizationIDParams {
	return &DeleteWorkspaceAPIV1OrganizationOrganizationIDParams{
		HTTPClient: client,
	}
}

/*
DeleteWorkspaceAPIV1OrganizationOrganizationIDParams contains all the parameters to send to the API endpoint

	for the delete workspace API v1 organization organization ID operation.

	Typically these are written to a http.Request.
*/
type DeleteWorkspaceAPIV1OrganizationOrganizationIDParams struct {

	/* OrganizationID.

	   Organization ID
	*/
	OrganizationID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete workspace API v1 organization organization ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteWorkspaceAPIV1OrganizationOrganizationIDParams) WithDefaults() *DeleteWorkspaceAPIV1OrganizationOrganizationIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete workspace API v1 organization organization ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteWorkspaceAPIV1OrganizationOrganizationIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete workspace API v1 organization organization ID params
func (o *DeleteWorkspaceAPIV1OrganizationOrganizationIDParams) WithTimeout(timeout time.Duration) *DeleteWorkspaceAPIV1OrganizationOrganizationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete workspace API v1 organization organization ID params
func (o *DeleteWorkspaceAPIV1OrganizationOrganizationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete workspace API v1 organization organization ID params
func (o *DeleteWorkspaceAPIV1OrganizationOrganizationIDParams) WithContext(ctx context.Context) *DeleteWorkspaceAPIV1OrganizationOrganizationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete workspace API v1 organization organization ID params
func (o *DeleteWorkspaceAPIV1OrganizationOrganizationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete workspace API v1 organization organization ID params
func (o *DeleteWorkspaceAPIV1OrganizationOrganizationIDParams) WithHTTPClient(client *http.Client) *DeleteWorkspaceAPIV1OrganizationOrganizationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete workspace API v1 organization organization ID params
func (o *DeleteWorkspaceAPIV1OrganizationOrganizationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the delete workspace API v1 organization organization ID params
func (o *DeleteWorkspaceAPIV1OrganizationOrganizationIDParams) WithOrganizationID(organizationID int64) *DeleteWorkspaceAPIV1OrganizationOrganizationIDParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the delete workspace API v1 organization organization ID params
func (o *DeleteWorkspaceAPIV1OrganizationOrganizationIDParams) SetOrganizationID(organizationID int64) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteWorkspaceAPIV1OrganizationOrganizationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", swag.FormatInt64(o.OrganizationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
