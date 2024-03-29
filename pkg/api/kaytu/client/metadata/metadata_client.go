// Code generated by go-swagger; DO NOT EDIT.

package metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new metadata API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for metadata API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetMetadataAPIV1Filter(params *GetMetadataAPIV1FilterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetMetadataAPIV1FilterOK, error)

	GetMetadataAPIV1MetadataKey(params *GetMetadataAPIV1MetadataKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetMetadataAPIV1MetadataKeyOK, error)

	PostMetadataAPIV1Filter(params *PostMetadataAPIV1FilterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostMetadataAPIV1FilterOK, error)

	PostMetadataAPIV1Metadata(params *PostMetadataAPIV1MetadataParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostMetadataAPIV1MetadataOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetMetadataAPIV1Filter lists filters
*/
func (a *Client) GetMetadataAPIV1Filter(params *GetMetadataAPIV1FilterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetMetadataAPIV1FilterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMetadataAPIV1FilterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetMetadataAPIV1Filter",
		Method:             "GET",
		PathPattern:        "/metadata/api/v1/filter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMetadataAPIV1FilterReader{formats: a.formats},
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
	success, ok := result.(*GetMetadataAPIV1FilterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetMetadataAPIV1Filter: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetMetadataAPIV1MetadataKey gets key metadata

Returns the config metadata for the given key
*/
func (a *Client) GetMetadataAPIV1MetadataKey(params *GetMetadataAPIV1MetadataKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetMetadataAPIV1MetadataKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMetadataAPIV1MetadataKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetMetadataAPIV1MetadataKey",
		Method:             "GET",
		PathPattern:        "/metadata/api/v1/metadata/{key}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMetadataAPIV1MetadataKeyReader{formats: a.formats},
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
	success, ok := result.(*GetMetadataAPIV1MetadataKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetMetadataAPIV1MetadataKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostMetadataAPIV1Filter adds filter
*/
func (a *Client) PostMetadataAPIV1Filter(params *PostMetadataAPIV1FilterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostMetadataAPIV1FilterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostMetadataAPIV1FilterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostMetadataAPIV1Filter",
		Method:             "POST",
		PathPattern:        "/metadata/api/v1/filter",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostMetadataAPIV1FilterReader{formats: a.formats},
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
	success, ok := result.(*PostMetadataAPIV1FilterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostMetadataAPIV1Filter: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostMetadataAPIV1Metadata sets key metadata

Sets the config metadata for the given key
*/
func (a *Client) PostMetadataAPIV1Metadata(params *PostMetadataAPIV1MetadataParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostMetadataAPIV1MetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostMetadataAPIV1MetadataParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostMetadataAPIV1Metadata",
		Method:             "POST",
		PathPattern:        "/metadata/api/v1/metadata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostMetadataAPIV1MetadataReader{formats: a.formats},
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
	success, ok := result.(*PostMetadataAPIV1MetadataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostMetadataAPIV1Metadata: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
