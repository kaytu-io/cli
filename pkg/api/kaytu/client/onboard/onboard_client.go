// Code generated by go-swagger; DO NOT EDIT.

package onboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new onboard API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for onboard API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteOnboardAPIV1CredentialCredentialID(params *DeleteOnboardAPIV1CredentialCredentialIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOnboardAPIV1CredentialCredentialIDOK, error)

	DeleteOnboardAPIV1SourceSourceID(params *DeleteOnboardAPIV1SourceSourceIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOnboardAPIV1SourceSourceIDOK, error)

	GetOnboardAPIV1CatalogMetrics(params *GetOnboardAPIV1CatalogMetricsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOnboardAPIV1CatalogMetricsOK, error)

	GetOnboardAPIV1Connector(params *GetOnboardAPIV1ConnectorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOnboardAPIV1ConnectorOK, error)

	GetOnboardAPIV1Credential(params *GetOnboardAPIV1CredentialParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOnboardAPIV1CredentialOK, error)

	GetOnboardAPIV1CredentialCredentialID(params *GetOnboardAPIV1CredentialCredentialIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOnboardAPIV1CredentialCredentialIDOK, error)

	GetOnboardAPIV1SourceSourceIDHealthcheck(params *GetOnboardAPIV1SourceSourceIDHealthcheckParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOnboardAPIV1SourceSourceIDHealthcheckOK, error)

	PostOnboardAPIV1Credential(params *PostOnboardAPIV1CredentialParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostOnboardAPIV1CredentialOK, error)

	PostOnboardAPIV1CredentialCredentialIDAutoonboard(params *PostOnboardAPIV1CredentialCredentialIDAutoonboardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostOnboardAPIV1CredentialCredentialIDAutoonboardOK, error)

	PostOnboardAPIV1SourceAws(params *PostOnboardAPIV1SourceAwsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostOnboardAPIV1SourceAwsOK, error)

	PostOnboardAPIV1SourceAzure(params *PostOnboardAPIV1SourceAzureParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostOnboardAPIV1SourceAzureOK, error)

	PutOnboardAPIV1CredentialCredentialID(params *PutOnboardAPIV1CredentialCredentialIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutOnboardAPIV1CredentialCredentialIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteOnboardAPIV1CredentialCredentialID deletes credential

Remove a credential by ID
*/
func (a *Client) DeleteOnboardAPIV1CredentialCredentialID(params *DeleteOnboardAPIV1CredentialCredentialIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOnboardAPIV1CredentialCredentialIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOnboardAPIV1CredentialCredentialIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteOnboardAPIV1CredentialCredentialID",
		Method:             "DELETE",
		PathPattern:        "/onboard/api/v1/credential/{credentialId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteOnboardAPIV1CredentialCredentialIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteOnboardAPIV1CredentialCredentialIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteOnboardAPIV1CredentialCredentialID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteOnboardAPIV1SourceSourceID deletes source

Deleting a single source either AWS / Azure for the given source id.
*/
func (a *Client) DeleteOnboardAPIV1SourceSourceID(params *DeleteOnboardAPIV1SourceSourceIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOnboardAPIV1SourceSourceIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOnboardAPIV1SourceSourceIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteOnboardAPIV1SourceSourceID",
		Method:             "DELETE",
		PathPattern:        "/onboard/api/v1/source/{sourceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteOnboardAPIV1SourceSourceIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteOnboardAPIV1SourceSourceIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteOnboardAPIV1SourceSourceID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetOnboardAPIV1CatalogMetrics lists catalog metrics

Retrieving the list of metrics for catalog page.
*/
func (a *Client) GetOnboardAPIV1CatalogMetrics(params *GetOnboardAPIV1CatalogMetricsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOnboardAPIV1CatalogMetricsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOnboardAPIV1CatalogMetricsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetOnboardAPIV1CatalogMetrics",
		Method:             "GET",
		PathPattern:        "/onboard/api/v1/catalog/metrics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOnboardAPIV1CatalogMetricsReader{formats: a.formats},
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
	success, ok := result.(*GetOnboardAPIV1CatalogMetricsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOnboardAPIV1CatalogMetrics: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetOnboardAPIV1Connector lists connectors

Returns list of all connectors
*/
func (a *Client) GetOnboardAPIV1Connector(params *GetOnboardAPIV1ConnectorParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOnboardAPIV1ConnectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOnboardAPIV1ConnectorParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetOnboardAPIV1Connector",
		Method:             "GET",
		PathPattern:        "/onboard/api/v1/connector",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOnboardAPIV1ConnectorReader{formats: a.formats},
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
	success, ok := result.(*GetOnboardAPIV1ConnectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOnboardAPIV1Connector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetOnboardAPIV1Credential lists credentials

Retrieving list of credentials with their details
*/
func (a *Client) GetOnboardAPIV1Credential(params *GetOnboardAPIV1CredentialParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOnboardAPIV1CredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOnboardAPIV1CredentialParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetOnboardAPIV1Credential",
		Method:             "GET",
		PathPattern:        "/onboard/api/v1/credential",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOnboardAPIV1CredentialReader{formats: a.formats},
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
	success, ok := result.(*GetOnboardAPIV1CredentialOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOnboardAPIV1Credential: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetOnboardAPIV1CredentialCredentialID gets credential

Retrieving credential details by credential ID
*/
func (a *Client) GetOnboardAPIV1CredentialCredentialID(params *GetOnboardAPIV1CredentialCredentialIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOnboardAPIV1CredentialCredentialIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOnboardAPIV1CredentialCredentialIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetOnboardAPIV1CredentialCredentialID",
		Method:             "GET",
		PathPattern:        "/onboard/api/v1/credential/{credentialId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOnboardAPIV1CredentialCredentialIDReader{formats: a.formats},
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
	success, ok := result.(*GetOnboardAPIV1CredentialCredentialIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOnboardAPIV1CredentialCredentialID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetOnboardAPIV1SourceSourceIDHealthcheck gets source health

Get live source health status with given source ID.
*/
func (a *Client) GetOnboardAPIV1SourceSourceIDHealthcheck(params *GetOnboardAPIV1SourceSourceIDHealthcheckParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOnboardAPIV1SourceSourceIDHealthcheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOnboardAPIV1SourceSourceIDHealthcheckParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetOnboardAPIV1SourceSourceIDHealthcheck",
		Method:             "GET",
		PathPattern:        "/onboard/api/v1/source/{sourceId}/healthcheck",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOnboardAPIV1SourceSourceIDHealthcheckReader{formats: a.formats},
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
	success, ok := result.(*GetOnboardAPIV1SourceSourceIDHealthcheckOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetOnboardAPIV1SourceSourceIDHealthcheck: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostOnboardAPIV1Credential creates connection credentials

Creating connection credentials
*/
func (a *Client) PostOnboardAPIV1Credential(params *PostOnboardAPIV1CredentialParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostOnboardAPIV1CredentialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOnboardAPIV1CredentialParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostOnboardAPIV1Credential",
		Method:             "POST",
		PathPattern:        "/onboard/api/v1/credential",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostOnboardAPIV1CredentialReader{formats: a.formats},
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
	success, ok := result.(*PostOnboardAPIV1CredentialOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostOnboardAPIV1Credential: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostOnboardAPIV1CredentialCredentialIDAutoonboard onboards credential connections

Onboard all available connections for a credential
*/
func (a *Client) PostOnboardAPIV1CredentialCredentialIDAutoonboard(params *PostOnboardAPIV1CredentialCredentialIDAutoonboardParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostOnboardAPIV1CredentialCredentialIDAutoonboardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOnboardAPIV1CredentialCredentialIDAutoonboardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostOnboardAPIV1CredentialCredentialIDAutoonboard",
		Method:             "POST",
		PathPattern:        "/onboard/api/v1/credential/{credentialId}/autoonboard",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostOnboardAPIV1CredentialCredentialIDAutoonboardReader{formats: a.formats},
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
	success, ok := result.(*PostOnboardAPIV1CredentialCredentialIDAutoonboardOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostOnboardAPIV1CredentialCredentialIDAutoonboard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostOnboardAPIV1SourceAws creates a w s source

Creating AWS source
*/
func (a *Client) PostOnboardAPIV1SourceAws(params *PostOnboardAPIV1SourceAwsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostOnboardAPIV1SourceAwsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOnboardAPIV1SourceAwsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostOnboardAPIV1SourceAws",
		Method:             "POST",
		PathPattern:        "/onboard/api/v1/source/aws",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostOnboardAPIV1SourceAwsReader{formats: a.formats},
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
	success, ok := result.(*PostOnboardAPIV1SourceAwsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostOnboardAPIV1SourceAws: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostOnboardAPIV1SourceAzure creates azure source

Creating Azure source
*/
func (a *Client) PostOnboardAPIV1SourceAzure(params *PostOnboardAPIV1SourceAzureParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostOnboardAPIV1SourceAzureOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOnboardAPIV1SourceAzureParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostOnboardAPIV1SourceAzure",
		Method:             "POST",
		PathPattern:        "/onboard/api/v1/source/azure",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostOnboardAPIV1SourceAzureReader{formats: a.formats},
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
	success, ok := result.(*PostOnboardAPIV1SourceAzureOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostOnboardAPIV1SourceAzure: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PutOnboardAPIV1CredentialCredentialID edits credential

Edit a credential by ID
*/
func (a *Client) PutOnboardAPIV1CredentialCredentialID(params *PutOnboardAPIV1CredentialCredentialIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutOnboardAPIV1CredentialCredentialIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutOnboardAPIV1CredentialCredentialIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutOnboardAPIV1CredentialCredentialID",
		Method:             "PUT",
		PathPattern:        "/onboard/api/v1/credential/{credentialId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutOnboardAPIV1CredentialCredentialIDReader{formats: a.formats},
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
	success, ok := result.(*PutOnboardAPIV1CredentialCredentialIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutOnboardAPIV1CredentialCredentialID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
