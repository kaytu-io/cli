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

// NewGetAuthAPIV1UsersParams creates a new GetAuthAPIV1UsersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAuthAPIV1UsersParams() *GetAuthAPIV1UsersParams {
	return &GetAuthAPIV1UsersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthAPIV1UsersParamsWithTimeout creates a new GetAuthAPIV1UsersParams object
// with the ability to set a timeout on a request.
func NewGetAuthAPIV1UsersParamsWithTimeout(timeout time.Duration) *GetAuthAPIV1UsersParams {
	return &GetAuthAPIV1UsersParams{
		timeout: timeout,
	}
}

// NewGetAuthAPIV1UsersParamsWithContext creates a new GetAuthAPIV1UsersParams object
// with the ability to set a context for a request.
func NewGetAuthAPIV1UsersParamsWithContext(ctx context.Context) *GetAuthAPIV1UsersParams {
	return &GetAuthAPIV1UsersParams{
		Context: ctx,
	}
}

// NewGetAuthAPIV1UsersParamsWithHTTPClient creates a new GetAuthAPIV1UsersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAuthAPIV1UsersParamsWithHTTPClient(client *http.Client) *GetAuthAPIV1UsersParams {
	return &GetAuthAPIV1UsersParams{
		HTTPClient: client,
	}
}

/*
GetAuthAPIV1UsersParams contains all the parameters to send to the API endpoint

	for the get auth API v1 users operation.

	Typically these are written to a http.Request.
*/
type GetAuthAPIV1UsersParams struct {

	/* Request.

	   Request Body
	*/
	Request *models.GitlabComKeibiengineKeibiEnginePkgAuthAPIGetUsersRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get auth API v1 users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAuthAPIV1UsersParams) WithDefaults() *GetAuthAPIV1UsersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get auth API v1 users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAuthAPIV1UsersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get auth API v1 users params
func (o *GetAuthAPIV1UsersParams) WithTimeout(timeout time.Duration) *GetAuthAPIV1UsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get auth API v1 users params
func (o *GetAuthAPIV1UsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get auth API v1 users params
func (o *GetAuthAPIV1UsersParams) WithContext(ctx context.Context) *GetAuthAPIV1UsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get auth API v1 users params
func (o *GetAuthAPIV1UsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get auth API v1 users params
func (o *GetAuthAPIV1UsersParams) WithHTTPClient(client *http.Client) *GetAuthAPIV1UsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get auth API v1 users params
func (o *GetAuthAPIV1UsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the get auth API v1 users params
func (o *GetAuthAPIV1UsersParams) WithRequest(request *models.GitlabComKeibiengineKeibiEnginePkgAuthAPIGetUsersRequest) *GetAuthAPIV1UsersParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the get auth API v1 users params
func (o *GetAuthAPIV1UsersParams) SetRequest(request *models.GitlabComKeibiengineKeibiEnginePkgAuthAPIGetUsersRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthAPIV1UsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
