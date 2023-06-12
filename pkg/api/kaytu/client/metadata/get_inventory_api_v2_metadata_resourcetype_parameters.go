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

// NewGetInventoryAPIV2MetadataResourcetypeParams creates a new GetInventoryAPIV2MetadataResourcetypeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInventoryAPIV2MetadataResourcetypeParams() *GetInventoryAPIV2MetadataResourcetypeParams {
	return &GetInventoryAPIV2MetadataResourcetypeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInventoryAPIV2MetadataResourcetypeParamsWithTimeout creates a new GetInventoryAPIV2MetadataResourcetypeParams object
// with the ability to set a timeout on a request.
func NewGetInventoryAPIV2MetadataResourcetypeParamsWithTimeout(timeout time.Duration) *GetInventoryAPIV2MetadataResourcetypeParams {
	return &GetInventoryAPIV2MetadataResourcetypeParams{
		timeout: timeout,
	}
}

// NewGetInventoryAPIV2MetadataResourcetypeParamsWithContext creates a new GetInventoryAPIV2MetadataResourcetypeParams object
// with the ability to set a context for a request.
func NewGetInventoryAPIV2MetadataResourcetypeParamsWithContext(ctx context.Context) *GetInventoryAPIV2MetadataResourcetypeParams {
	return &GetInventoryAPIV2MetadataResourcetypeParams{
		Context: ctx,
	}
}

// NewGetInventoryAPIV2MetadataResourcetypeParamsWithHTTPClient creates a new GetInventoryAPIV2MetadataResourcetypeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInventoryAPIV2MetadataResourcetypeParamsWithHTTPClient(client *http.Client) *GetInventoryAPIV2MetadataResourcetypeParams {
	return &GetInventoryAPIV2MetadataResourcetypeParams{
		HTTPClient: client,
	}
}

/*
GetInventoryAPIV2MetadataResourcetypeParams contains all the parameters to send to the API endpoint

	for the get inventory API v2 metadata resourcetype operation.

	Typically these are written to a http.Request.
*/
type GetInventoryAPIV2MetadataResourcetypeParams struct {

	/* Connector.

	   Filter by Connector
	*/
	Connector []string

	/* PageNumber.

	   page number - default is 1
	*/
	PageNumber *int64

	/* PageSize.

	   page size - default is 20
	*/
	PageSize *int64

	/* Service.

	   Filter by service name
	*/
	Service []string

	/* Tag.

	   Key-Value tags in key=value format to filter by
	*/
	Tag []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get inventory API v2 metadata resourcetype params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInventoryAPIV2MetadataResourcetypeParams) WithDefaults() *GetInventoryAPIV2MetadataResourcetypeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get inventory API v2 metadata resourcetype params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInventoryAPIV2MetadataResourcetypeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get inventory API v2 metadata resourcetype params
func (o *GetInventoryAPIV2MetadataResourcetypeParams) WithTimeout(timeout time.Duration) *GetInventoryAPIV2MetadataResourcetypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get inventory API v2 metadata resourcetype params
func (o *GetInventoryAPIV2MetadataResourcetypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get inventory API v2 metadata resourcetype params
func (o *GetInventoryAPIV2MetadataResourcetypeParams) WithContext(ctx context.Context) *GetInventoryAPIV2MetadataResourcetypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get inventory API v2 metadata resourcetype params
func (o *GetInventoryAPIV2MetadataResourcetypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get inventory API v2 metadata resourcetype params
func (o *GetInventoryAPIV2MetadataResourcetypeParams) WithHTTPClient(client *http.Client) *GetInventoryAPIV2MetadataResourcetypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get inventory API v2 metadata resourcetype params
func (o *GetInventoryAPIV2MetadataResourcetypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnector adds the connector to the get inventory API v2 metadata resourcetype params
func (o *GetInventoryAPIV2MetadataResourcetypeParams) WithConnector(connector []string) *GetInventoryAPIV2MetadataResourcetypeParams {
	o.SetConnector(connector)
	return o
}

// SetConnector adds the connector to the get inventory API v2 metadata resourcetype params
func (o *GetInventoryAPIV2MetadataResourcetypeParams) SetConnector(connector []string) {
	o.Connector = connector
}

// WithPageNumber adds the pageNumber to the get inventory API v2 metadata resourcetype params
func (o *GetInventoryAPIV2MetadataResourcetypeParams) WithPageNumber(pageNumber *int64) *GetInventoryAPIV2MetadataResourcetypeParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get inventory API v2 metadata resourcetype params
func (o *GetInventoryAPIV2MetadataResourcetypeParams) SetPageNumber(pageNumber *int64) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get inventory API v2 metadata resourcetype params
func (o *GetInventoryAPIV2MetadataResourcetypeParams) WithPageSize(pageSize *int64) *GetInventoryAPIV2MetadataResourcetypeParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get inventory API v2 metadata resourcetype params
func (o *GetInventoryAPIV2MetadataResourcetypeParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithService adds the service to the get inventory API v2 metadata resourcetype params
func (o *GetInventoryAPIV2MetadataResourcetypeParams) WithService(service []string) *GetInventoryAPIV2MetadataResourcetypeParams {
	o.SetService(service)
	return o
}

// SetService adds the service to the get inventory API v2 metadata resourcetype params
func (o *GetInventoryAPIV2MetadataResourcetypeParams) SetService(service []string) {
	o.Service = service
}

// WithTag adds the tag to the get inventory API v2 metadata resourcetype params
func (o *GetInventoryAPIV2MetadataResourcetypeParams) WithTag(tag []string) *GetInventoryAPIV2MetadataResourcetypeParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the get inventory API v2 metadata resourcetype params
func (o *GetInventoryAPIV2MetadataResourcetypeParams) SetTag(tag []string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *GetInventoryAPIV2MetadataResourcetypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Service != nil {

		// binding items for service
		joinedService := o.bindParamService(reg)

		// query array param service
		if err := r.SetQueryParam("service", joinedService...); err != nil {
			return err
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

// bindParamGetInventoryAPIV2MetadataResourcetype binds the parameter connector
func (o *GetInventoryAPIV2MetadataResourcetypeParams) bindParamConnector(formats strfmt.Registry) []string {
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

// bindParamGetInventoryAPIV2MetadataResourcetype binds the parameter service
func (o *GetInventoryAPIV2MetadataResourcetypeParams) bindParamService(formats strfmt.Registry) []string {
	serviceIR := o.Service

	var serviceIC []string
	for _, serviceIIR := range serviceIR { // explode []string

		serviceIIV := serviceIIR // string as string
		serviceIC = append(serviceIC, serviceIIV)
	}

	// items.CollectionFormat: "csv"
	serviceIS := swag.JoinByFormat(serviceIC, "csv")

	return serviceIS
}

// bindParamGetInventoryAPIV2MetadataResourcetype binds the parameter tag
func (o *GetInventoryAPIV2MetadataResourcetypeParams) bindParamTag(formats strfmt.Registry) []string {
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
