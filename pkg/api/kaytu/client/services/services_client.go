// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new services API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for services API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetInventoryAPIV2ServicesSummary(params *GetInventoryAPIV2ServicesSummaryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ServicesSummaryOK, error)

	GetInventoryAPIV2ServicesSummaryServiceName(params *GetInventoryAPIV2ServicesSummaryServiceNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ServicesSummaryServiceNameOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetInventoryAPIV2ServicesSummary lists cloud services summary

Retrieves list of summaries of the services including the number of them and the API filters and a list of services with more details. Including connector and the resource counts.
*/
func (a *Client) GetInventoryAPIV2ServicesSummary(params *GetInventoryAPIV2ServicesSummaryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ServicesSummaryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryAPIV2ServicesSummaryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInventoryAPIV2ServicesSummary",
		Method:             "GET",
		PathPattern:        "/inventory/api/v2/services/summary",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryAPIV2ServicesSummaryReader{formats: a.formats},
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
	success, ok := result.(*GetInventoryAPIV2ServicesSummaryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInventoryAPIV2ServicesSummary: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInventoryAPIV2ServicesSummaryServiceName gets cloud service summary

Retrieves Cloud Service Summary for the specified service name. Including connector, the resource counts.
*/
func (a *Client) GetInventoryAPIV2ServicesSummaryServiceName(params *GetInventoryAPIV2ServicesSummaryServiceNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ServicesSummaryServiceNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryAPIV2ServicesSummaryServiceNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInventoryAPIV2ServicesSummaryServiceName",
		Method:             "GET",
		PathPattern:        "/inventory/api/v2/services/summary/{serviceName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryAPIV2ServicesSummaryServiceNameReader{formats: a.formats},
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
	success, ok := result.(*GetInventoryAPIV2ServicesSummaryServiceNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInventoryAPIV2ServicesSummaryServiceName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}