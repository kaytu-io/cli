// Code generated by go-swagger; DO NOT EDIT.

package metadata

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

// NewGetInventoryAPIV2MetadataServicesParams creates a new GetInventoryAPIV2MetadataServicesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInventoryAPIV2MetadataServicesParams() *GetInventoryAPIV2MetadataServicesParams {
	return &GetInventoryAPIV2MetadataServicesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInventoryAPIV2MetadataServicesParamsWithTimeout creates a new GetInventoryAPIV2MetadataServicesParams object
// with the ability to set a timeout on a request.
func NewGetInventoryAPIV2MetadataServicesParamsWithTimeout(timeout time.Duration) *GetInventoryAPIV2MetadataServicesParams {
	return &GetInventoryAPIV2MetadataServicesParams{
		timeout: timeout,
	}
}

// NewGetInventoryAPIV2MetadataServicesParamsWithContext creates a new GetInventoryAPIV2MetadataServicesParams object
// with the ability to set a context for a request.
func NewGetInventoryAPIV2MetadataServicesParamsWithContext(ctx context.Context) *GetInventoryAPIV2MetadataServicesParams {
	return &GetInventoryAPIV2MetadataServicesParams{
		Context: ctx,
	}
}

// NewGetInventoryAPIV2MetadataServicesParamsWithHTTPClient creates a new GetInventoryAPIV2MetadataServicesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInventoryAPIV2MetadataServicesParamsWithHTTPClient(client *http.Client) *GetInventoryAPIV2MetadataServicesParams {
	return &GetInventoryAPIV2MetadataServicesParams{
		HTTPClient: client,
	}
}

/*
GetInventoryAPIV2MetadataServicesParams contains all the parameters to send to the API endpoint

	for the get inventory API v2 metadata services operation.

	Typically these are written to a http.Request.
*/
type GetInventoryAPIV2MetadataServicesParams struct {

	/* Connector.

	   Connector
	*/
	Connector []string

	/* CostSupport.

	   Filter by cost support
	*/
	CostSupport *bool

	/* PageNumber.

	   page number - default is 1
	*/
	PageNumber *int64

	/* PageSize.

	   page size - default is 20
	*/
	PageSize *int64

	/* Tag.

	   Key-Value tags in key=value format to filter by
	*/
	Tag []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get inventory API v2 metadata services params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInventoryAPIV2MetadataServicesParams) WithDefaults() *GetInventoryAPIV2MetadataServicesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get inventory API v2 metadata services params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInventoryAPIV2MetadataServicesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get inventory API v2 metadata services params
func (o *GetInventoryAPIV2MetadataServicesParams) WithTimeout(timeout time.Duration) *GetInventoryAPIV2MetadataServicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get inventory API v2 metadata services params
func (o *GetInventoryAPIV2MetadataServicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get inventory API v2 metadata services params
func (o *GetInventoryAPIV2MetadataServicesParams) WithContext(ctx context.Context) *GetInventoryAPIV2MetadataServicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get inventory API v2 metadata services params
func (o *GetInventoryAPIV2MetadataServicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get inventory API v2 metadata services params
func (o *GetInventoryAPIV2MetadataServicesParams) WithHTTPClient(client *http.Client) *GetInventoryAPIV2MetadataServicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get inventory API v2 metadata services params
func (o *GetInventoryAPIV2MetadataServicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnector adds the connector to the get inventory API v2 metadata services params
func (o *GetInventoryAPIV2MetadataServicesParams) WithConnector(connector []string) *GetInventoryAPIV2MetadataServicesParams {
	o.SetConnector(connector)
	return o
}

// SetConnector adds the connector to the get inventory API v2 metadata services params
func (o *GetInventoryAPIV2MetadataServicesParams) SetConnector(connector []string) {
	o.Connector = connector
}

// WithCostSupport adds the costSupport to the get inventory API v2 metadata services params
func (o *GetInventoryAPIV2MetadataServicesParams) WithCostSupport(costSupport *bool) *GetInventoryAPIV2MetadataServicesParams {
	o.SetCostSupport(costSupport)
	return o
}

// SetCostSupport adds the costSupport to the get inventory API v2 metadata services params
func (o *GetInventoryAPIV2MetadataServicesParams) SetCostSupport(costSupport *bool) {
	o.CostSupport = costSupport
}

// WithPageNumber adds the pageNumber to the get inventory API v2 metadata services params
func (o *GetInventoryAPIV2MetadataServicesParams) WithPageNumber(pageNumber *int64) *GetInventoryAPIV2MetadataServicesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get inventory API v2 metadata services params
func (o *GetInventoryAPIV2MetadataServicesParams) SetPageNumber(pageNumber *int64) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get inventory API v2 metadata services params
func (o *GetInventoryAPIV2MetadataServicesParams) WithPageSize(pageSize *int64) *GetInventoryAPIV2MetadataServicesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get inventory API v2 metadata services params
func (o *GetInventoryAPIV2MetadataServicesParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithTag adds the tag to the get inventory API v2 metadata services params
func (o *GetInventoryAPIV2MetadataServicesParams) WithTag(tag []string) *GetInventoryAPIV2MetadataServicesParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the get inventory API v2 metadata services params
func (o *GetInventoryAPIV2MetadataServicesParams) SetTag(tag []string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *GetInventoryAPIV2MetadataServicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Connector != nil {

		// binding items for connector
		joinedConnector := o.bindParamConnector(reg)

		// query array param connector
		if err := r.SetQueryParam("connector", joinedConnector...); err != nil {
			return err
		}
	}

	if o.CostSupport != nil {

		// query param costSupport
		var qrCostSupport bool

		if o.CostSupport != nil {
			qrCostSupport = *o.CostSupport
		}
		qCostSupport := swag.FormatBool(qrCostSupport)
		if qCostSupport != "" {

			if err := r.SetQueryParam("costSupport", qCostSupport); err != nil {
				return err
			}
		}
	}

	if o.PageNumber != nil {

		// query param pageNumber
		var qrPageNumber int64

		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := swag.FormatInt64(qrPageNumber)
		if qPageNumber != "" {

			if err := r.SetQueryParam("pageNumber", qPageNumber); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.Tag != nil {

		// binding items for tag
		joinedTag := o.bindParamTag(reg)

		// query array param tag
		if err := r.SetQueryParam("tag", joinedTag...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetInventoryAPIV2MetadataServices binds the parameter connector
func (o *GetInventoryAPIV2MetadataServicesParams) bindParamConnector(formats strfmt.Registry) []string {
	connectorIR := o.Connector

	var connectorIC []string
	for _, connectorIIR := range connectorIR { // explode []string

		connectorIIV := connectorIIR // string as string
		connectorIC = append(connectorIC, connectorIIV)
	}

	// items.CollectionFormat: "csv"
	connectorIS := swag.JoinByFormat(connectorIC, "csv")

	return connectorIS
}

// bindParamGetInventoryAPIV2MetadataServices binds the parameter tag
func (o *GetInventoryAPIV2MetadataServicesParams) bindParamTag(formats strfmt.Registry) []string {
	tagIR := o.Tag

	var tagIC []string
	for _, tagIIR := range tagIR { // explode []string

		tagIIV := tagIIR // string as string
		tagIC = append(tagIC, tagIIV)
	}

	// items.CollectionFormat: "csv"
	tagIS := swag.JoinByFormat(tagIC, "csv")

	return tagIS
}
