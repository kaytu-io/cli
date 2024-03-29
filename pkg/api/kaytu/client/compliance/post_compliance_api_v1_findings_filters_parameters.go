// Code generated by go-swagger; DO NOT EDIT.

package compliance

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

	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

// NewPostComplianceAPIV1FindingsFiltersParams creates a new PostComplianceAPIV1FindingsFiltersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostComplianceAPIV1FindingsFiltersParams() *PostComplianceAPIV1FindingsFiltersParams {
	return &PostComplianceAPIV1FindingsFiltersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostComplianceAPIV1FindingsFiltersParamsWithTimeout creates a new PostComplianceAPIV1FindingsFiltersParams object
// with the ability to set a timeout on a request.
func NewPostComplianceAPIV1FindingsFiltersParamsWithTimeout(timeout time.Duration) *PostComplianceAPIV1FindingsFiltersParams {
	return &PostComplianceAPIV1FindingsFiltersParams{
		timeout: timeout,
	}
}

// NewPostComplianceAPIV1FindingsFiltersParamsWithContext creates a new PostComplianceAPIV1FindingsFiltersParams object
// with the ability to set a context for a request.
func NewPostComplianceAPIV1FindingsFiltersParamsWithContext(ctx context.Context) *PostComplianceAPIV1FindingsFiltersParams {
	return &PostComplianceAPIV1FindingsFiltersParams{
		Context: ctx,
	}
}

// NewPostComplianceAPIV1FindingsFiltersParamsWithHTTPClient creates a new PostComplianceAPIV1FindingsFiltersParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostComplianceAPIV1FindingsFiltersParamsWithHTTPClient(client *http.Client) *PostComplianceAPIV1FindingsFiltersParams {
	return &PostComplianceAPIV1FindingsFiltersParams{
		HTTPClient: client,
	}
}

/*
PostComplianceAPIV1FindingsFiltersParams contains all the parameters to send to the API endpoint

	for the post compliance API v1 findings filters operation.

	Typically these are written to a http.Request.
*/
type PostComplianceAPIV1FindingsFiltersParams struct {

	/* Request.

	   Request Body
	*/
	Request *models.GithubComKaytuIoKaytuEnginePkgComplianceAPIFindingFilters

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post compliance API v1 findings filters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostComplianceAPIV1FindingsFiltersParams) WithDefaults() *PostComplianceAPIV1FindingsFiltersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post compliance API v1 findings filters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostComplianceAPIV1FindingsFiltersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post compliance API v1 findings filters params
func (o *PostComplianceAPIV1FindingsFiltersParams) WithTimeout(timeout time.Duration) *PostComplianceAPIV1FindingsFiltersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post compliance API v1 findings filters params
func (o *PostComplianceAPIV1FindingsFiltersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post compliance API v1 findings filters params
func (o *PostComplianceAPIV1FindingsFiltersParams) WithContext(ctx context.Context) *PostComplianceAPIV1FindingsFiltersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post compliance API v1 findings filters params
func (o *PostComplianceAPIV1FindingsFiltersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post compliance API v1 findings filters params
func (o *PostComplianceAPIV1FindingsFiltersParams) WithHTTPClient(client *http.Client) *PostComplianceAPIV1FindingsFiltersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post compliance API v1 findings filters params
func (o *PostComplianceAPIV1FindingsFiltersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post compliance API v1 findings filters params
func (o *PostComplianceAPIV1FindingsFiltersParams) WithRequest(request *models.GithubComKaytuIoKaytuEnginePkgComplianceAPIFindingFilters) *PostComplianceAPIV1FindingsFiltersParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post compliance API v1 findings filters params
func (o *PostComplianceAPIV1FindingsFiltersParams) SetRequest(request *models.GithubComKaytuIoKaytuEnginePkgComplianceAPIFindingFilters) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostComplianceAPIV1FindingsFiltersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
