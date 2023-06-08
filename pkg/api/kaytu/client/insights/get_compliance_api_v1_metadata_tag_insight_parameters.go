// Code generated by go-swagger; DO NOT EDIT.

package insights

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

// NewGetComplianceAPIV1MetadataTagInsightParams creates a new GetComplianceAPIV1MetadataTagInsightParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetComplianceAPIV1MetadataTagInsightParams() *GetComplianceAPIV1MetadataTagInsightParams {
	return &GetComplianceAPIV1MetadataTagInsightParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetComplianceAPIV1MetadataTagInsightParamsWithTimeout creates a new GetComplianceAPIV1MetadataTagInsightParams object
// with the ability to set a timeout on a request.
func NewGetComplianceAPIV1MetadataTagInsightParamsWithTimeout(timeout time.Duration) *GetComplianceAPIV1MetadataTagInsightParams {
	return &GetComplianceAPIV1MetadataTagInsightParams{
		timeout: timeout,
	}
}

// NewGetComplianceAPIV1MetadataTagInsightParamsWithContext creates a new GetComplianceAPIV1MetadataTagInsightParams object
// with the ability to set a context for a request.
func NewGetComplianceAPIV1MetadataTagInsightParamsWithContext(ctx context.Context) *GetComplianceAPIV1MetadataTagInsightParams {
	return &GetComplianceAPIV1MetadataTagInsightParams{
		Context: ctx,
	}
}

// NewGetComplianceAPIV1MetadataTagInsightParamsWithHTTPClient creates a new GetComplianceAPIV1MetadataTagInsightParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetComplianceAPIV1MetadataTagInsightParamsWithHTTPClient(client *http.Client) *GetComplianceAPIV1MetadataTagInsightParams {
	return &GetComplianceAPIV1MetadataTagInsightParams{
		HTTPClient: client,
	}
}

/*
GetComplianceAPIV1MetadataTagInsightParams contains all the parameters to send to the API endpoint

	for the get compliance API v1 metadata tag insight operation.

	Typically these are written to a http.Request.
*/
type GetComplianceAPIV1MetadataTagInsightParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get compliance API v1 metadata tag insight params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComplianceAPIV1MetadataTagInsightParams) WithDefaults() *GetComplianceAPIV1MetadataTagInsightParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get compliance API v1 metadata tag insight params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComplianceAPIV1MetadataTagInsightParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get compliance API v1 metadata tag insight params
func (o *GetComplianceAPIV1MetadataTagInsightParams) WithTimeout(timeout time.Duration) *GetComplianceAPIV1MetadataTagInsightParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get compliance API v1 metadata tag insight params
func (o *GetComplianceAPIV1MetadataTagInsightParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get compliance API v1 metadata tag insight params
func (o *GetComplianceAPIV1MetadataTagInsightParams) WithContext(ctx context.Context) *GetComplianceAPIV1MetadataTagInsightParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get compliance API v1 metadata tag insight params
func (o *GetComplianceAPIV1MetadataTagInsightParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get compliance API v1 metadata tag insight params
func (o *GetComplianceAPIV1MetadataTagInsightParams) WithHTTPClient(client *http.Client) *GetComplianceAPIV1MetadataTagInsightParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get compliance API v1 metadata tag insight params
func (o *GetComplianceAPIV1MetadataTagInsightParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetComplianceAPIV1MetadataTagInsightParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}