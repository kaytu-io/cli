// Code generated by go-swagger; DO NOT EDIT.

package onboard

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

// NewGetOnboardAPIV1ConnectorsConnectorNameParams creates a new GetOnboardAPIV1ConnectorsConnectorNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOnboardAPIV1ConnectorsConnectorNameParams() *GetOnboardAPIV1ConnectorsConnectorNameParams {
	return &GetOnboardAPIV1ConnectorsConnectorNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOnboardAPIV1ConnectorsConnectorNameParamsWithTimeout creates a new GetOnboardAPIV1ConnectorsConnectorNameParams object
// with the ability to set a timeout on a request.
func NewGetOnboardAPIV1ConnectorsConnectorNameParamsWithTimeout(timeout time.Duration) *GetOnboardAPIV1ConnectorsConnectorNameParams {
	return &GetOnboardAPIV1ConnectorsConnectorNameParams{
		timeout: timeout,
	}
}

// NewGetOnboardAPIV1ConnectorsConnectorNameParamsWithContext creates a new GetOnboardAPIV1ConnectorsConnectorNameParams object
// with the ability to set a context for a request.
func NewGetOnboardAPIV1ConnectorsConnectorNameParamsWithContext(ctx context.Context) *GetOnboardAPIV1ConnectorsConnectorNameParams {
	return &GetOnboardAPIV1ConnectorsConnectorNameParams{
		Context: ctx,
	}
}

// NewGetOnboardAPIV1ConnectorsConnectorNameParamsWithHTTPClient creates a new GetOnboardAPIV1ConnectorsConnectorNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOnboardAPIV1ConnectorsConnectorNameParamsWithHTTPClient(client *http.Client) *GetOnboardAPIV1ConnectorsConnectorNameParams {
	return &GetOnboardAPIV1ConnectorsConnectorNameParams{
		HTTPClient: client,
	}
}

/*
GetOnboardAPIV1ConnectorsConnectorNameParams contains all the parameters to send to the API endpoint

	for the get onboard API v1 connectors connector name operation.

	Typically these are written to a http.Request.
*/
type GetOnboardAPIV1ConnectorsConnectorNameParams struct {

	/* ConnectorName.

	   Connector name
	*/
	ConnectorName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get onboard API v1 connectors connector name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOnboardAPIV1ConnectorsConnectorNameParams) WithDefaults() *GetOnboardAPIV1ConnectorsConnectorNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get onboard API v1 connectors connector name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOnboardAPIV1ConnectorsConnectorNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get onboard API v1 connectors connector name params
func (o *GetOnboardAPIV1ConnectorsConnectorNameParams) WithTimeout(timeout time.Duration) *GetOnboardAPIV1ConnectorsConnectorNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get onboard API v1 connectors connector name params
func (o *GetOnboardAPIV1ConnectorsConnectorNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get onboard API v1 connectors connector name params
func (o *GetOnboardAPIV1ConnectorsConnectorNameParams) WithContext(ctx context.Context) *GetOnboardAPIV1ConnectorsConnectorNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get onboard API v1 connectors connector name params
func (o *GetOnboardAPIV1ConnectorsConnectorNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get onboard API v1 connectors connector name params
func (o *GetOnboardAPIV1ConnectorsConnectorNameParams) WithHTTPClient(client *http.Client) *GetOnboardAPIV1ConnectorsConnectorNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get onboard API v1 connectors connector name params
func (o *GetOnboardAPIV1ConnectorsConnectorNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectorName adds the connectorName to the get onboard API v1 connectors connector name params
func (o *GetOnboardAPIV1ConnectorsConnectorNameParams) WithConnectorName(connectorName string) *GetOnboardAPIV1ConnectorsConnectorNameParams {
	o.SetConnectorName(connectorName)
	return o
}

// SetConnectorName adds the connectorName to the get onboard API v1 connectors connector name params
func (o *GetOnboardAPIV1ConnectorsConnectorNameParams) SetConnectorName(connectorName string) {
	o.ConnectorName = connectorName
}

// WriteToRequest writes these params to a swagger request
func (o *GetOnboardAPIV1ConnectorsConnectorNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param connectorName
	if err := r.SetPathParam("connectorName", o.ConnectorName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}