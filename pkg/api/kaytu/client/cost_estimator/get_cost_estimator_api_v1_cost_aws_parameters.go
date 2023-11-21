// Code generated by go-swagger; DO NOT EDIT.

package cost_estimator

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

// NewGetCostEstimatorAPIV1CostAwsParams creates a new GetCostEstimatorAPIV1CostAwsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCostEstimatorAPIV1CostAwsParams() *GetCostEstimatorAPIV1CostAwsParams {
	return &GetCostEstimatorAPIV1CostAwsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCostEstimatorAPIV1CostAwsParamsWithTimeout creates a new GetCostEstimatorAPIV1CostAwsParams object
// with the ability to set a timeout on a request.
func NewGetCostEstimatorAPIV1CostAwsParamsWithTimeout(timeout time.Duration) *GetCostEstimatorAPIV1CostAwsParams {
	return &GetCostEstimatorAPIV1CostAwsParams{
		timeout: timeout,
	}
}

// NewGetCostEstimatorAPIV1CostAwsParamsWithContext creates a new GetCostEstimatorAPIV1CostAwsParams object
// with the ability to set a context for a request.
func NewGetCostEstimatorAPIV1CostAwsParamsWithContext(ctx context.Context) *GetCostEstimatorAPIV1CostAwsParams {
	return &GetCostEstimatorAPIV1CostAwsParams{
		Context: ctx,
	}
}

// NewGetCostEstimatorAPIV1CostAwsParamsWithHTTPClient creates a new GetCostEstimatorAPIV1CostAwsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCostEstimatorAPIV1CostAwsParamsWithHTTPClient(client *http.Client) *GetCostEstimatorAPIV1CostAwsParams {
	return &GetCostEstimatorAPIV1CostAwsParams{
		HTTPClient: client,
	}
}

/*
GetCostEstimatorAPIV1CostAwsParams contains all the parameters to send to the API endpoint

	for the get cost estimator API v1 cost aws operation.

	Typically these are written to a http.Request.
*/
type GetCostEstimatorAPIV1CostAwsParams struct {

	/* ResourceID.

	   Connection ID
	*/
	ResourceID string

	/* ResourceType.

	   ResourceType
	*/
	ResourceType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cost estimator API v1 cost aws params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCostEstimatorAPIV1CostAwsParams) WithDefaults() *GetCostEstimatorAPIV1CostAwsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cost estimator API v1 cost aws params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCostEstimatorAPIV1CostAwsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cost estimator API v1 cost aws params
func (o *GetCostEstimatorAPIV1CostAwsParams) WithTimeout(timeout time.Duration) *GetCostEstimatorAPIV1CostAwsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cost estimator API v1 cost aws params
func (o *GetCostEstimatorAPIV1CostAwsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cost estimator API v1 cost aws params
func (o *GetCostEstimatorAPIV1CostAwsParams) WithContext(ctx context.Context) *GetCostEstimatorAPIV1CostAwsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cost estimator API v1 cost aws params
func (o *GetCostEstimatorAPIV1CostAwsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cost estimator API v1 cost aws params
func (o *GetCostEstimatorAPIV1CostAwsParams) WithHTTPClient(client *http.Client) *GetCostEstimatorAPIV1CostAwsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cost estimator API v1 cost aws params
func (o *GetCostEstimatorAPIV1CostAwsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceID adds the resourceID to the get cost estimator API v1 cost aws params
func (o *GetCostEstimatorAPIV1CostAwsParams) WithResourceID(resourceID string) *GetCostEstimatorAPIV1CostAwsParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the get cost estimator API v1 cost aws params
func (o *GetCostEstimatorAPIV1CostAwsParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WithResourceType adds the resourceType to the get cost estimator API v1 cost aws params
func (o *GetCostEstimatorAPIV1CostAwsParams) WithResourceType(resourceType string) *GetCostEstimatorAPIV1CostAwsParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the get cost estimator API v1 cost aws params
func (o *GetCostEstimatorAPIV1CostAwsParams) SetResourceType(resourceType string) {
	o.ResourceType = resourceType
}

// WriteToRequest writes these params to a swagger request
func (o *GetCostEstimatorAPIV1CostAwsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param resourceId
	qrResourceID := o.ResourceID
	qResourceID := qrResourceID
	if qResourceID != "" {

		if err := r.SetQueryParam("resourceId", qResourceID); err != nil {
			return err
		}
	}

	// query param resourceType
	qrResourceType := o.ResourceType
	qResourceType := qrResourceType
	if qResourceType != "" {

		if err := r.SetQueryParam("resourceType", qResourceType); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
