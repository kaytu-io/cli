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

	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

// NewPostOnboardAPIV1ConnectionsConnectionIDStateParams creates a new PostOnboardAPIV1ConnectionsConnectionIDStateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostOnboardAPIV1ConnectionsConnectionIDStateParams() *PostOnboardAPIV1ConnectionsConnectionIDStateParams {
	return &PostOnboardAPIV1ConnectionsConnectionIDStateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostOnboardAPIV1ConnectionsConnectionIDStateParamsWithTimeout creates a new PostOnboardAPIV1ConnectionsConnectionIDStateParams object
// with the ability to set a timeout on a request.
func NewPostOnboardAPIV1ConnectionsConnectionIDStateParamsWithTimeout(timeout time.Duration) *PostOnboardAPIV1ConnectionsConnectionIDStateParams {
	return &PostOnboardAPIV1ConnectionsConnectionIDStateParams{
		timeout: timeout,
	}
}

// NewPostOnboardAPIV1ConnectionsConnectionIDStateParamsWithContext creates a new PostOnboardAPIV1ConnectionsConnectionIDStateParams object
// with the ability to set a context for a request.
func NewPostOnboardAPIV1ConnectionsConnectionIDStateParamsWithContext(ctx context.Context) *PostOnboardAPIV1ConnectionsConnectionIDStateParams {
	return &PostOnboardAPIV1ConnectionsConnectionIDStateParams{
		Context: ctx,
	}
}

// NewPostOnboardAPIV1ConnectionsConnectionIDStateParamsWithHTTPClient creates a new PostOnboardAPIV1ConnectionsConnectionIDStateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostOnboardAPIV1ConnectionsConnectionIDStateParamsWithHTTPClient(client *http.Client) *PostOnboardAPIV1ConnectionsConnectionIDStateParams {
	return &PostOnboardAPIV1ConnectionsConnectionIDStateParams{
		HTTPClient: client,
	}
}

/*
PostOnboardAPIV1ConnectionsConnectionIDStateParams contains all the parameters to send to the API endpoint

	for the post onboard API v1 connections connection ID state operation.

	Typically these are written to a http.Request.
*/
type PostOnboardAPIV1ConnectionsConnectionIDStateParams struct {

	/* ConnectionID.

	   Connection ID
	*/
	ConnectionID string

	/* Request.

	   Request
	*/
	Request *models.GithubComKaytuIoKaytuEnginePkgOnboardAPIChangeConnectionLifecycleStateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post onboard API v1 connections connection ID state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostOnboardAPIV1ConnectionsConnectionIDStateParams) WithDefaults() *PostOnboardAPIV1ConnectionsConnectionIDStateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post onboard API v1 connections connection ID state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostOnboardAPIV1ConnectionsConnectionIDStateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post onboard API v1 connections connection ID state params
func (o *PostOnboardAPIV1ConnectionsConnectionIDStateParams) WithTimeout(timeout time.Duration) *PostOnboardAPIV1ConnectionsConnectionIDStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post onboard API v1 connections connection ID state params
func (o *PostOnboardAPIV1ConnectionsConnectionIDStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post onboard API v1 connections connection ID state params
func (o *PostOnboardAPIV1ConnectionsConnectionIDStateParams) WithContext(ctx context.Context) *PostOnboardAPIV1ConnectionsConnectionIDStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post onboard API v1 connections connection ID state params
func (o *PostOnboardAPIV1ConnectionsConnectionIDStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post onboard API v1 connections connection ID state params
func (o *PostOnboardAPIV1ConnectionsConnectionIDStateParams) WithHTTPClient(client *http.Client) *PostOnboardAPIV1ConnectionsConnectionIDStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post onboard API v1 connections connection ID state params
func (o *PostOnboardAPIV1ConnectionsConnectionIDStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionID adds the connectionID to the post onboard API v1 connections connection ID state params
func (o *PostOnboardAPIV1ConnectionsConnectionIDStateParams) WithConnectionID(connectionID string) *PostOnboardAPIV1ConnectionsConnectionIDStateParams {
	o.SetConnectionID(connectionID)
	return o
}

// SetConnectionID adds the connectionId to the post onboard API v1 connections connection ID state params
func (o *PostOnboardAPIV1ConnectionsConnectionIDStateParams) SetConnectionID(connectionID string) {
	o.ConnectionID = connectionID
}

// WithRequest adds the request to the post onboard API v1 connections connection ID state params
func (o *PostOnboardAPIV1ConnectionsConnectionIDStateParams) WithRequest(request *models.GithubComKaytuIoKaytuEnginePkgOnboardAPIChangeConnectionLifecycleStateRequest) *PostOnboardAPIV1ConnectionsConnectionIDStateParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post onboard API v1 connections connection ID state params
func (o *PostOnboardAPIV1ConnectionsConnectionIDStateParams) SetRequest(request *models.GithubComKaytuIoKaytuEnginePkgOnboardAPIChangeConnectionLifecycleStateRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostOnboardAPIV1ConnectionsConnectionIDStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param connectionId
	if err := r.SetPathParam("connectionId", o.ConnectionID); err != nil {
		return err
	}
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
