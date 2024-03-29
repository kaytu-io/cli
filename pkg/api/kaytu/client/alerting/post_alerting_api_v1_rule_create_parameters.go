// Code generated by go-swagger; DO NOT EDIT.

package alerting

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

// NewPostAlertingAPIV1RuleCreateParams creates a new PostAlertingAPIV1RuleCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostAlertingAPIV1RuleCreateParams() *PostAlertingAPIV1RuleCreateParams {
	return &PostAlertingAPIV1RuleCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostAlertingAPIV1RuleCreateParamsWithTimeout creates a new PostAlertingAPIV1RuleCreateParams object
// with the ability to set a timeout on a request.
func NewPostAlertingAPIV1RuleCreateParamsWithTimeout(timeout time.Duration) *PostAlertingAPIV1RuleCreateParams {
	return &PostAlertingAPIV1RuleCreateParams{
		timeout: timeout,
	}
}

// NewPostAlertingAPIV1RuleCreateParamsWithContext creates a new PostAlertingAPIV1RuleCreateParams object
// with the ability to set a context for a request.
func NewPostAlertingAPIV1RuleCreateParamsWithContext(ctx context.Context) *PostAlertingAPIV1RuleCreateParams {
	return &PostAlertingAPIV1RuleCreateParams{
		Context: ctx,
	}
}

// NewPostAlertingAPIV1RuleCreateParamsWithHTTPClient creates a new PostAlertingAPIV1RuleCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostAlertingAPIV1RuleCreateParamsWithHTTPClient(client *http.Client) *PostAlertingAPIV1RuleCreateParams {
	return &PostAlertingAPIV1RuleCreateParams{
		HTTPClient: client,
	}
}

/*
PostAlertingAPIV1RuleCreateParams contains all the parameters to send to the API endpoint

	for the post alerting API v1 rule create operation.

	Typically these are written to a http.Request.
*/
type PostAlertingAPIV1RuleCreateParams struct {

	/* Request.

	   Request Body
	*/
	Request *models.GithubComKaytuIoKaytuEnginePkgAlertingAPICreateRuleRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post alerting API v1 rule create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAlertingAPIV1RuleCreateParams) WithDefaults() *PostAlertingAPIV1RuleCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post alerting API v1 rule create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostAlertingAPIV1RuleCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post alerting API v1 rule create params
func (o *PostAlertingAPIV1RuleCreateParams) WithTimeout(timeout time.Duration) *PostAlertingAPIV1RuleCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post alerting API v1 rule create params
func (o *PostAlertingAPIV1RuleCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post alerting API v1 rule create params
func (o *PostAlertingAPIV1RuleCreateParams) WithContext(ctx context.Context) *PostAlertingAPIV1RuleCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post alerting API v1 rule create params
func (o *PostAlertingAPIV1RuleCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post alerting API v1 rule create params
func (o *PostAlertingAPIV1RuleCreateParams) WithHTTPClient(client *http.Client) *PostAlertingAPIV1RuleCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post alerting API v1 rule create params
func (o *PostAlertingAPIV1RuleCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the post alerting API v1 rule create params
func (o *PostAlertingAPIV1RuleCreateParams) WithRequest(request *models.GithubComKaytuIoKaytuEnginePkgAlertingAPICreateRuleRequest) *PostAlertingAPIV1RuleCreateParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post alerting API v1 rule create params
func (o *PostAlertingAPIV1RuleCreateParams) SetRequest(request *models.GithubComKaytuIoKaytuEnginePkgAlertingAPICreateRuleRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostAlertingAPIV1RuleCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
