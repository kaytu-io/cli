// Code generated by go-swagger; DO NOT EDIT.

package location

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

// NewGetInventoryAPIV1LocationsConnectorParams creates a new GetInventoryAPIV1LocationsConnectorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInventoryAPIV1LocationsConnectorParams() *GetInventoryAPIV1LocationsConnectorParams {
	return &GetInventoryAPIV1LocationsConnectorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInventoryAPIV1LocationsConnectorParamsWithTimeout creates a new GetInventoryAPIV1LocationsConnectorParams object
// with the ability to set a timeout on a request.
func NewGetInventoryAPIV1LocationsConnectorParamsWithTimeout(timeout time.Duration) *GetInventoryAPIV1LocationsConnectorParams {
	return &GetInventoryAPIV1LocationsConnectorParams{
		timeout: timeout,
	}
}

// NewGetInventoryAPIV1LocationsConnectorParamsWithContext creates a new GetInventoryAPIV1LocationsConnectorParams object
// with the ability to set a context for a request.
func NewGetInventoryAPIV1LocationsConnectorParamsWithContext(ctx context.Context) *GetInventoryAPIV1LocationsConnectorParams {
	return &GetInventoryAPIV1LocationsConnectorParams{
		Context: ctx,
	}
}

// NewGetInventoryAPIV1LocationsConnectorParamsWithHTTPClient creates a new GetInventoryAPIV1LocationsConnectorParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInventoryAPIV1LocationsConnectorParamsWithHTTPClient(client *http.Client) *GetInventoryAPIV1LocationsConnectorParams {
	return &GetInventoryAPIV1LocationsConnectorParams{
		HTTPClient: client,
	}
}

/*
GetInventoryAPIV1LocationsConnectorParams contains all the parameters to send to the API endpoint

	for the get inventory API v1 locations connector operation.

	Typically these are written to a http.Request.
*/
type GetInventoryAPIV1LocationsConnectorParams struct {

	/* Connector.

	   Connector
	*/
	Connector string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get inventory API v1 locations connector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInventoryAPIV1LocationsConnectorParams) WithDefaults() *GetInventoryAPIV1LocationsConnectorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get inventory API v1 locations connector params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInventoryAPIV1LocationsConnectorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get inventory API v1 locations connector params
func (o *GetInventoryAPIV1LocationsConnectorParams) WithTimeout(timeout time.Duration) *GetInventoryAPIV1LocationsConnectorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get inventory API v1 locations connector params
func (o *GetInventoryAPIV1LocationsConnectorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get inventory API v1 locations connector params
func (o *GetInventoryAPIV1LocationsConnectorParams) WithContext(ctx context.Context) *GetInventoryAPIV1LocationsConnectorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get inventory API v1 locations connector params
func (o *GetInventoryAPIV1LocationsConnectorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get inventory API v1 locations connector params
func (o *GetInventoryAPIV1LocationsConnectorParams) WithHTTPClient(client *http.Client) *GetInventoryAPIV1LocationsConnectorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get inventory API v1 locations connector params
func (o *GetInventoryAPIV1LocationsConnectorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnector adds the connector to the get inventory API v1 locations connector params
func (o *GetInventoryAPIV1LocationsConnectorParams) WithConnector(connector string) *GetInventoryAPIV1LocationsConnectorParams {
	o.SetConnector(connector)
	return o
}

// SetConnector adds the connector to the get inventory API v1 locations connector params
func (o *GetInventoryAPIV1LocationsConnectorParams) SetConnector(connector string) {
	o.Connector = connector
}

// WriteToRequest writes these params to a swagger request
func (o *GetInventoryAPIV1LocationsConnectorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param connector
	if err := r.SetPathParam("connector", o.Connector); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
