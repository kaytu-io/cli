// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewPutAuthAPIV1UserRoleBindingParams creates a new PutAuthAPIV1UserRoleBindingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutAuthAPIV1UserRoleBindingParams() *PutAuthAPIV1UserRoleBindingParams {
	return &PutAuthAPIV1UserRoleBindingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutAuthAPIV1UserRoleBindingParamsWithTimeout creates a new PutAuthAPIV1UserRoleBindingParams object
// with the ability to set a timeout on a request.
func NewPutAuthAPIV1UserRoleBindingParamsWithTimeout(timeout time.Duration) *PutAuthAPIV1UserRoleBindingParams {
	return &PutAuthAPIV1UserRoleBindingParams{
		timeout: timeout,
	}
}

// NewPutAuthAPIV1UserRoleBindingParamsWithContext creates a new PutAuthAPIV1UserRoleBindingParams object
// with the ability to set a context for a request.
func NewPutAuthAPIV1UserRoleBindingParamsWithContext(ctx context.Context) *PutAuthAPIV1UserRoleBindingParams {
	return &PutAuthAPIV1UserRoleBindingParams{
		Context: ctx,
	}
}

// NewPutAuthAPIV1UserRoleBindingParamsWithHTTPClient creates a new PutAuthAPIV1UserRoleBindingParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutAuthAPIV1UserRoleBindingParamsWithHTTPClient(client *http.Client) *PutAuthAPIV1UserRoleBindingParams {
	return &PutAuthAPIV1UserRoleBindingParams{
		HTTPClient: client,
	}
}

/*
PutAuthAPIV1UserRoleBindingParams contains all the parameters to send to the API endpoint

	for the put auth API v1 user role binding operation.

	Typically these are written to a http.Request.
*/
type PutAuthAPIV1UserRoleBindingParams struct {

	/* Request.

	   Request Body
	*/
	Request *models.GitlabComKeibiengineKeibiEnginePkgAuthAPIPutRoleBindingRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put auth API v1 user role binding params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutAuthAPIV1UserRoleBindingParams) WithDefaults() *PutAuthAPIV1UserRoleBindingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put auth API v1 user role binding params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutAuthAPIV1UserRoleBindingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put auth API v1 user role binding params
func (o *PutAuthAPIV1UserRoleBindingParams) WithTimeout(timeout time.Duration) *PutAuthAPIV1UserRoleBindingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put auth API v1 user role binding params
func (o *PutAuthAPIV1UserRoleBindingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put auth API v1 user role binding params
func (o *PutAuthAPIV1UserRoleBindingParams) WithContext(ctx context.Context) *PutAuthAPIV1UserRoleBindingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put auth API v1 user role binding params
func (o *PutAuthAPIV1UserRoleBindingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put auth API v1 user role binding params
func (o *PutAuthAPIV1UserRoleBindingParams) WithHTTPClient(client *http.Client) *PutAuthAPIV1UserRoleBindingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put auth API v1 user role binding params
func (o *PutAuthAPIV1UserRoleBindingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the put auth API v1 user role binding params
func (o *PutAuthAPIV1UserRoleBindingParams) WithRequest(request *models.GitlabComKeibiengineKeibiEnginePkgAuthAPIPutRoleBindingRequest) *PutAuthAPIV1UserRoleBindingParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the put auth API v1 user role binding params
func (o *PutAuthAPIV1UserRoleBindingParams) SetRequest(request *models.GitlabComKeibiengineKeibiEnginePkgAuthAPIPutRoleBindingRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PutAuthAPIV1UserRoleBindingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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