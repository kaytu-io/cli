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

// NewPutAlertingAPIV1RuleUpdateRuleIDParams creates a new PutAlertingAPIV1RuleUpdateRuleIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutAlertingAPIV1RuleUpdateRuleIDParams() *PutAlertingAPIV1RuleUpdateRuleIDParams {
	return &PutAlertingAPIV1RuleUpdateRuleIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutAlertingAPIV1RuleUpdateRuleIDParamsWithTimeout creates a new PutAlertingAPIV1RuleUpdateRuleIDParams object
// with the ability to set a timeout on a request.
func NewPutAlertingAPIV1RuleUpdateRuleIDParamsWithTimeout(timeout time.Duration) *PutAlertingAPIV1RuleUpdateRuleIDParams {
	return &PutAlertingAPIV1RuleUpdateRuleIDParams{
		timeout: timeout,
	}
}

// NewPutAlertingAPIV1RuleUpdateRuleIDParamsWithContext creates a new PutAlertingAPIV1RuleUpdateRuleIDParams object
// with the ability to set a context for a request.
func NewPutAlertingAPIV1RuleUpdateRuleIDParamsWithContext(ctx context.Context) *PutAlertingAPIV1RuleUpdateRuleIDParams {
	return &PutAlertingAPIV1RuleUpdateRuleIDParams{
		Context: ctx,
	}
}

// NewPutAlertingAPIV1RuleUpdateRuleIDParamsWithHTTPClient creates a new PutAlertingAPIV1RuleUpdateRuleIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutAlertingAPIV1RuleUpdateRuleIDParamsWithHTTPClient(client *http.Client) *PutAlertingAPIV1RuleUpdateRuleIDParams {
	return &PutAlertingAPIV1RuleUpdateRuleIDParams{
		HTTPClient: client,
	}
}

/*
PutAlertingAPIV1RuleUpdateRuleIDParams contains all the parameters to send to the API endpoint

	for the put alerting API v1 rule update rule ID operation.

	Typically these are written to a http.Request.
*/
type PutAlertingAPIV1RuleUpdateRuleIDParams struct {

	/* Request.

	   Request Body
	*/
	Request *models.GithubComKaytuIoKaytuEnginePkgAlertingAPIUpdateRuleRequest

	/* RuleID.

	   ruleId
	*/
	RuleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put alerting API v1 rule update rule ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutAlertingAPIV1RuleUpdateRuleIDParams) WithDefaults() *PutAlertingAPIV1RuleUpdateRuleIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put alerting API v1 rule update rule ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutAlertingAPIV1RuleUpdateRuleIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put alerting API v1 rule update rule ID params
func (o *PutAlertingAPIV1RuleUpdateRuleIDParams) WithTimeout(timeout time.Duration) *PutAlertingAPIV1RuleUpdateRuleIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put alerting API v1 rule update rule ID params
func (o *PutAlertingAPIV1RuleUpdateRuleIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put alerting API v1 rule update rule ID params
func (o *PutAlertingAPIV1RuleUpdateRuleIDParams) WithContext(ctx context.Context) *PutAlertingAPIV1RuleUpdateRuleIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put alerting API v1 rule update rule ID params
func (o *PutAlertingAPIV1RuleUpdateRuleIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put alerting API v1 rule update rule ID params
func (o *PutAlertingAPIV1RuleUpdateRuleIDParams) WithHTTPClient(client *http.Client) *PutAlertingAPIV1RuleUpdateRuleIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put alerting API v1 rule update rule ID params
func (o *PutAlertingAPIV1RuleUpdateRuleIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the put alerting API v1 rule update rule ID params
func (o *PutAlertingAPIV1RuleUpdateRuleIDParams) WithRequest(request *models.GithubComKaytuIoKaytuEnginePkgAlertingAPIUpdateRuleRequest) *PutAlertingAPIV1RuleUpdateRuleIDParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the put alerting API v1 rule update rule ID params
func (o *PutAlertingAPIV1RuleUpdateRuleIDParams) SetRequest(request *models.GithubComKaytuIoKaytuEnginePkgAlertingAPIUpdateRuleRequest) {
	o.Request = request
}

// WithRuleID adds the ruleID to the put alerting API v1 rule update rule ID params
func (o *PutAlertingAPIV1RuleUpdateRuleIDParams) WithRuleID(ruleID string) *PutAlertingAPIV1RuleUpdateRuleIDParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the put alerting API v1 rule update rule ID params
func (o *PutAlertingAPIV1RuleUpdateRuleIDParams) SetRuleID(ruleID string) {
	o.RuleID = ruleID
}

// WriteToRequest writes these params to a swagger request
func (o *PutAlertingAPIV1RuleUpdateRuleIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	// path param ruleId
	if err := r.SetPathParam("ruleId", o.RuleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}