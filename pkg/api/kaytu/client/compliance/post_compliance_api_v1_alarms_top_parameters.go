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

// NewPostComplianceAPIV1AlarmsTopParams creates a new PostComplianceAPIV1AlarmsTopParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostComplianceAPIV1AlarmsTopParams() *PostComplianceAPIV1AlarmsTopParams {
	return &PostComplianceAPIV1AlarmsTopParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostComplianceAPIV1AlarmsTopParamsWithTimeout creates a new PostComplianceAPIV1AlarmsTopParams object
// with the ability to set a timeout on a request.
func NewPostComplianceAPIV1AlarmsTopParamsWithTimeout(timeout time.Duration) *PostComplianceAPIV1AlarmsTopParams {
	return &PostComplianceAPIV1AlarmsTopParams{
		timeout: timeout,
	}
}

// NewPostComplianceAPIV1AlarmsTopParamsWithContext creates a new PostComplianceAPIV1AlarmsTopParams object
// with the ability to set a context for a request.
func NewPostComplianceAPIV1AlarmsTopParamsWithContext(ctx context.Context) *PostComplianceAPIV1AlarmsTopParams {
	return &PostComplianceAPIV1AlarmsTopParams{
		Context: ctx,
	}
}

// NewPostComplianceAPIV1AlarmsTopParamsWithHTTPClient creates a new PostComplianceAPIV1AlarmsTopParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostComplianceAPIV1AlarmsTopParamsWithHTTPClient(client *http.Client) *PostComplianceAPIV1AlarmsTopParams {
	return &PostComplianceAPIV1AlarmsTopParams{
		HTTPClient: client,
	}
}

/*
PostComplianceAPIV1AlarmsTopParams contains all the parameters to send to the API endpoint

	for the post compliance API v1 alarms top operation.

	Typically these are written to a http.Request.
*/
type PostComplianceAPIV1AlarmsTopParams struct {

	/* Request.

	   Request Body
	*/
	Request *models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIGetTopFieldRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post compliance API v1 alarms top params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostComplianceAPIV1AlarmsTopParams) WithDefaults() *PostComplianceAPIV1AlarmsTopParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post compliance API v1 alarms top params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostComplianceAPIV1AlarmsTopParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post compliance API v1 alarms top params
func (o *PostComplianceAPIV1AlarmsTopParams) WithTimeout(timeout time.Duration) *PostComplianceAPIV1AlarmsTopParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post compliance API v1 alarms top params
func (o *PostComplianceAPIV1AlarmsTopParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post compliance API v1 alarms top params
func (o *PostComplianceAPIV1AlarmsTopParams) WithContext(ctx context.Context) *PostComplianceAPIV1AlarmsTopParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post compliance API v1 alarms top params
func (o *PostComplianceAPIV1AlarmsTopParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post compliance API v1 alarms top params
func (o *PostComplianceAPIV1AlarmsTopParams) WithHTTPClient(client *http.Client) *PostComplianceAPIV1AlarmsTopParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post compliance API v1 alarms top params
func (o *PostComplianceAPIV1AlarmsTopParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post compliance API v1 alarms top params
func (o *PostComplianceAPIV1AlarmsTopParams) WithRequest(request *models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIGetTopFieldRequest) *PostComplianceAPIV1AlarmsTopParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post compliance API v1 alarms top params
func (o *PostComplianceAPIV1AlarmsTopParams) SetRequest(request *models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIGetTopFieldRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostComplianceAPIV1AlarmsTopParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
