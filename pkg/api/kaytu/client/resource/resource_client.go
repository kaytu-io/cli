// Code generated by go-swagger; DO NOT EDIT.

package resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new resource API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for resource API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetInventoryAPIV2ResourcesMetricResourceType(params *GetInventoryAPIV2ResourcesMetricResourceTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ResourcesMetricResourceTypeOK, error)

	PostAiAPIV1GptRun(params *PostAiAPIV1GptRunParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostAiAPIV1GptRunOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetInventoryAPIV2ResourcesMetricResourceType lists resource type metrics

Retrieving metrics for a specific resource type.
*/
func (a *Client) GetInventoryAPIV2ResourcesMetricResourceType(params *GetInventoryAPIV2ResourcesMetricResourceTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ResourcesMetricResourceTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryAPIV2ResourcesMetricResourceTypeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInventoryAPIV2ResourcesMetricResourceType",
		Method:             "GET",
		PathPattern:        "/inventory/api/v2/resources/metric/{resourceType}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryAPIV2ResourcesMetricResourceTypeReader{formats: a.formats},
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
	success, ok := result.(*GetInventoryAPIV2ResourcesMetricResourceTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInventoryAPIV2ResourcesMetricResourceType: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostAiAPIV1GptRun runs the query on kaytu g p t and returns the generated query
*/
func (a *Client) PostAiAPIV1GptRun(params *PostAiAPIV1GptRunParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostAiAPIV1GptRunOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAiAPIV1GptRunParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostAiAPIV1GptRun",
		Method:             "POST",
		PathPattern:        "/ai/api/v1/gpt/run",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAiAPIV1GptRunReader{formats: a.formats},
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
	success, ok := result.(*PostAiAPIV1GptRunOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAiAPIV1GptRun: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
