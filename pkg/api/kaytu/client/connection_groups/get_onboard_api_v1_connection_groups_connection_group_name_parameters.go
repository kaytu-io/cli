// Code generated by go-swagger; DO NOT EDIT.

package connection_groups

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

// NewGetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams creates a new GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams() *GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams {
	return &GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOnboardAPIV1ConnectionGroupsConnectionGroupNameParamsWithTimeout creates a new GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams object
// with the ability to set a timeout on a request.
func NewGetOnboardAPIV1ConnectionGroupsConnectionGroupNameParamsWithTimeout(timeout time.Duration) *GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams {
	return &GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams{
		timeout: timeout,
	}
}

// NewGetOnboardAPIV1ConnectionGroupsConnectionGroupNameParamsWithContext creates a new GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams object
// with the ability to set a context for a request.
func NewGetOnboardAPIV1ConnectionGroupsConnectionGroupNameParamsWithContext(ctx context.Context) *GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams {
	return &GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams{
		Context: ctx,
	}
}

// NewGetOnboardAPIV1ConnectionGroupsConnectionGroupNameParamsWithHTTPClient creates a new GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOnboardAPIV1ConnectionGroupsConnectionGroupNameParamsWithHTTPClient(client *http.Client) *GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams {
	return &GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams{
		HTTPClient: client,
	}
}

/*
GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams contains all the parameters to send to the API endpoint

	for the get onboard API v1 connection groups connection group name operation.

	Typically these are written to a http.Request.
*/
type GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams struct {

	/* ConnectionGroupName.

	   ConnectionGroupName
	*/
	ConnectionGroupName string

	/* PopulateConnections.

	   Populate connections
	*/
	PopulateConnections *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get onboard API v1 connection groups connection group name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams) WithDefaults() *GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get onboard API v1 connection groups connection group name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams) SetDefaults() {
	var (
		populateConnectionsDefault = bool(false)
	)

	val := GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams{
		PopulateConnections: &populateConnectionsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get onboard API v1 connection groups connection group name params
func (o *GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams) WithTimeout(timeout time.Duration) *GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get onboard API v1 connection groups connection group name params
func (o *GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get onboard API v1 connection groups connection group name params
func (o *GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams) WithContext(ctx context.Context) *GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get onboard API v1 connection groups connection group name params
func (o *GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get onboard API v1 connection groups connection group name params
func (o *GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams) WithHTTPClient(client *http.Client) *GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get onboard API v1 connection groups connection group name params
func (o *GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionGroupName adds the connectionGroupName to the get onboard API v1 connection groups connection group name params
func (o *GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams) WithConnectionGroupName(connectionGroupName string) *GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams {
	o.SetConnectionGroupName(connectionGroupName)
	return o
}

// SetConnectionGroupName adds the connectionGroupName to the get onboard API v1 connection groups connection group name params
func (o *GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams) SetConnectionGroupName(connectionGroupName string) {
	o.ConnectionGroupName = connectionGroupName
}

// WithPopulateConnections adds the populateConnections to the get onboard API v1 connection groups connection group name params
func (o *GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams) WithPopulateConnections(populateConnections *bool) *GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams {
	o.SetPopulateConnections(populateConnections)
	return o
}

// SetPopulateConnections adds the populateConnections to the get onboard API v1 connection groups connection group name params
func (o *GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams) SetPopulateConnections(populateConnections *bool) {
	o.PopulateConnections = populateConnections
}

// WriteToRequest writes these params to a swagger request
func (o *GetOnboardAPIV1ConnectionGroupsConnectionGroupNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param connectionGroupName
	if err := r.SetPathParam("connectionGroupName", o.ConnectionGroupName); err != nil {
		return err
	}

	if o.PopulateConnections != nil {

		// query param populateConnections
		var qrPopulateConnections bool

		if o.PopulateConnections != nil {
			qrPopulateConnections = *o.PopulateConnections
		}
		qPopulateConnections := swag.FormatBool(qrPopulateConnections)
		if qPopulateConnections != "" {

			if err := r.SetQueryParam("populateConnections", qPopulateConnections); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
