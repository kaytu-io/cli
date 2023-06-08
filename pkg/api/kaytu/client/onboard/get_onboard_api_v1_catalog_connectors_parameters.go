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

// NewGetOnboardAPIV1CatalogConnectorsParams creates a new GetOnboardAPIV1CatalogConnectorsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOnboardAPIV1CatalogConnectorsParams() *GetOnboardAPIV1CatalogConnectorsParams {
	return &GetOnboardAPIV1CatalogConnectorsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOnboardAPIV1CatalogConnectorsParamsWithTimeout creates a new GetOnboardAPIV1CatalogConnectorsParams object
// with the ability to set a timeout on a request.
func NewGetOnboardAPIV1CatalogConnectorsParamsWithTimeout(timeout time.Duration) *GetOnboardAPIV1CatalogConnectorsParams {
	return &GetOnboardAPIV1CatalogConnectorsParams{
		timeout: timeout,
	}
}

// NewGetOnboardAPIV1CatalogConnectorsParamsWithContext creates a new GetOnboardAPIV1CatalogConnectorsParams object
// with the ability to set a context for a request.
func NewGetOnboardAPIV1CatalogConnectorsParamsWithContext(ctx context.Context) *GetOnboardAPIV1CatalogConnectorsParams {
	return &GetOnboardAPIV1CatalogConnectorsParams{
		Context: ctx,
	}
}

// NewGetOnboardAPIV1CatalogConnectorsParamsWithHTTPClient creates a new GetOnboardAPIV1CatalogConnectorsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOnboardAPIV1CatalogConnectorsParamsWithHTTPClient(client *http.Client) *GetOnboardAPIV1CatalogConnectorsParams {
	return &GetOnboardAPIV1CatalogConnectorsParams{
		HTTPClient: client,
	}
}

/*
GetOnboardAPIV1CatalogConnectorsParams contains all the parameters to send to the API endpoint

	for the get onboard API v1 catalog connectors operation.

	Typically these are written to a http.Request.
*/
type GetOnboardAPIV1CatalogConnectorsParams struct {

	/* Category.

	   Category filter
	*/
	Category *string

	/* ID.

	   ID filter
	*/
	ID *string

	/* MinConnection.

	   Minimum connection filter
	*/
	MinConnection *string

	/* State.

	   State filter
	*/
	State *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get onboard API v1 catalog connectors params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOnboardAPIV1CatalogConnectorsParams) WithDefaults() *GetOnboardAPIV1CatalogConnectorsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get onboard API v1 catalog connectors params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOnboardAPIV1CatalogConnectorsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get onboard API v1 catalog connectors params
func (o *GetOnboardAPIV1CatalogConnectorsParams) WithTimeout(timeout time.Duration) *GetOnboardAPIV1CatalogConnectorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get onboard API v1 catalog connectors params
func (o *GetOnboardAPIV1CatalogConnectorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get onboard API v1 catalog connectors params
func (o *GetOnboardAPIV1CatalogConnectorsParams) WithContext(ctx context.Context) *GetOnboardAPIV1CatalogConnectorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get onboard API v1 catalog connectors params
func (o *GetOnboardAPIV1CatalogConnectorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get onboard API v1 catalog connectors params
func (o *GetOnboardAPIV1CatalogConnectorsParams) WithHTTPClient(client *http.Client) *GetOnboardAPIV1CatalogConnectorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get onboard API v1 catalog connectors params
func (o *GetOnboardAPIV1CatalogConnectorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategory adds the category to the get onboard API v1 catalog connectors params
func (o *GetOnboardAPIV1CatalogConnectorsParams) WithCategory(category *string) *GetOnboardAPIV1CatalogConnectorsParams {
	o.SetCategory(category)
	return o
}

// SetCategory adds the category to the get onboard API v1 catalog connectors params
func (o *GetOnboardAPIV1CatalogConnectorsParams) SetCategory(category *string) {
	o.Category = category
}

// WithID adds the id to the get onboard API v1 catalog connectors params
func (o *GetOnboardAPIV1CatalogConnectorsParams) WithID(id *string) *GetOnboardAPIV1CatalogConnectorsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get onboard API v1 catalog connectors params
func (o *GetOnboardAPIV1CatalogConnectorsParams) SetID(id *string) {
	o.ID = id
}

// WithMinConnection adds the minConnection to the get onboard API v1 catalog connectors params
func (o *GetOnboardAPIV1CatalogConnectorsParams) WithMinConnection(minConnection *string) *GetOnboardAPIV1CatalogConnectorsParams {
	o.SetMinConnection(minConnection)
	return o
}

// SetMinConnection adds the minConnection to the get onboard API v1 catalog connectors params
func (o *GetOnboardAPIV1CatalogConnectorsParams) SetMinConnection(minConnection *string) {
	o.MinConnection = minConnection
}

// WithState adds the state to the get onboard API v1 catalog connectors params
func (o *GetOnboardAPIV1CatalogConnectorsParams) WithState(state *string) *GetOnboardAPIV1CatalogConnectorsParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the get onboard API v1 catalog connectors params
func (o *GetOnboardAPIV1CatalogConnectorsParams) SetState(state *string) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *GetOnboardAPIV1CatalogConnectorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Category != nil {

		// query param category
		var qrCategory string

		if o.Category != nil {
			qrCategory = *o.Category
		}
		qCategory := qrCategory
		if qCategory != "" {

			if err := r.SetQueryParam("category", qCategory); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.MinConnection != nil {

		// query param minConnection
		var qrMinConnection string

		if o.MinConnection != nil {
			qrMinConnection = *o.MinConnection
		}
		qMinConnection := qrMinConnection
		if qMinConnection != "" {

			if err := r.SetQueryParam("minConnection", qMinConnection); err != nil {
				return err
			}
		}
	}

	if o.State != nil {

		// query param state
		var qrState string

		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {

			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
