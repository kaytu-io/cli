// Code generated by go-swagger; DO NOT EDIT.

package cost_estimator

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new cost estimator API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cost estimator API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetCostEstimatorAPIV1CostAws(params *GetCostEstimatorAPIV1CostAwsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCostEstimatorAPIV1CostAwsOK, error)

	GetCostEstimatorAPIV1CostAzure(params *GetCostEstimatorAPIV1CostAzureParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCostEstimatorAPIV1CostAzureOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetCostEstimatorAPIV1CostAws gets a w s cost

Get AWS cost for each resource
*/
func (a *Client) GetCostEstimatorAPIV1CostAws(params *GetCostEstimatorAPIV1CostAwsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCostEstimatorAPIV1CostAwsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCostEstimatorAPIV1CostAwsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCostEstimatorAPIV1CostAws",
		Method:             "GET",
		PathPattern:        "/cost_estimator/api/v1/cost/aws",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCostEstimatorAPIV1CostAwsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCostEstimatorAPIV1CostAwsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetCostEstimatorAPIV1CostAws: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCostEstimatorAPIV1CostAzure gets azure cost

Get Azure cost for each resource
*/
func (a *Client) GetCostEstimatorAPIV1CostAzure(params *GetCostEstimatorAPIV1CostAzureParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCostEstimatorAPIV1CostAzureOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCostEstimatorAPIV1CostAzureParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCostEstimatorAPIV1CostAzure",
		Method:             "GET",
		PathPattern:        "/cost_estimator/api/v1/cost/azure",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCostEstimatorAPIV1CostAzureReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCostEstimatorAPIV1CostAzureOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetCostEstimatorAPIV1CostAzure: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
