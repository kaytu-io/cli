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
)

// NewDeleteAuthAPIV1UserRoleBindingParams creates a new DeleteAuthAPIV1UserRoleBindingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAuthAPIV1UserRoleBindingParams() *DeleteAuthAPIV1UserRoleBindingParams {
	return &DeleteAuthAPIV1UserRoleBindingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAuthAPIV1UserRoleBindingParamsWithTimeout creates a new DeleteAuthAPIV1UserRoleBindingParams object
// with the ability to set a timeout on a request.
func NewDeleteAuthAPIV1UserRoleBindingParamsWithTimeout(timeout time.Duration) *DeleteAuthAPIV1UserRoleBindingParams {
	return &DeleteAuthAPIV1UserRoleBindingParams{
		timeout: timeout,
	}
}

// NewDeleteAuthAPIV1UserRoleBindingParamsWithContext creates a new DeleteAuthAPIV1UserRoleBindingParams object
// with the ability to set a context for a request.
func NewDeleteAuthAPIV1UserRoleBindingParamsWithContext(ctx context.Context) *DeleteAuthAPIV1UserRoleBindingParams {
	return &DeleteAuthAPIV1UserRoleBindingParams{
		Context: ctx,
	}
}

// NewDeleteAuthAPIV1UserRoleBindingParamsWithHTTPClient creates a new DeleteAuthAPIV1UserRoleBindingParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAuthAPIV1UserRoleBindingParamsWithHTTPClient(client *http.Client) *DeleteAuthAPIV1UserRoleBindingParams {
	return &DeleteAuthAPIV1UserRoleBindingParams{
		HTTPClient: client,
	}
}

/*
DeleteAuthAPIV1UserRoleBindingParams contains all the parameters to send to the API endpoint

	for the delete auth API v1 user role binding operation.

	Typically these are written to a http.Request.
*/
type DeleteAuthAPIV1UserRoleBindingParams struct {

	/* UserID.

	   userId
	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete auth API v1 user role binding params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAuthAPIV1UserRoleBindingParams) WithDefaults() *DeleteAuthAPIV1UserRoleBindingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete auth API v1 user role binding params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAuthAPIV1UserRoleBindingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete auth API v1 user role binding params
func (o *DeleteAuthAPIV1UserRoleBindingParams) WithTimeout(timeout time.Duration) *DeleteAuthAPIV1UserRoleBindingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete auth API v1 user role binding params
func (o *DeleteAuthAPIV1UserRoleBindingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete auth API v1 user role binding params
func (o *DeleteAuthAPIV1UserRoleBindingParams) WithContext(ctx context.Context) *DeleteAuthAPIV1UserRoleBindingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete auth API v1 user role binding params
func (o *DeleteAuthAPIV1UserRoleBindingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete auth API v1 user role binding params
func (o *DeleteAuthAPIV1UserRoleBindingParams) WithHTTPClient(client *http.Client) *DeleteAuthAPIV1UserRoleBindingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete auth API v1 user role binding params
func (o *DeleteAuthAPIV1UserRoleBindingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the delete auth API v1 user role binding params
func (o *DeleteAuthAPIV1UserRoleBindingParams) WithUserID(userID string) *DeleteAuthAPIV1UserRoleBindingParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete auth API v1 user role binding params
func (o *DeleteAuthAPIV1UserRoleBindingParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAuthAPIV1UserRoleBindingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param userId
	qrUserID := o.UserID
	qUserID := qrUserID
	if qUserID != "" {

		if err := r.SetQueryParam("userId", qUserID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
