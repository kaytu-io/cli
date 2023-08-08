// Code generated by go-swagger; DO NOT EDIT.

package insights

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new insights API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for insights API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetComplianceAPIV1Insight(params *GetComplianceAPIV1InsightParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetComplianceAPIV1InsightOK, error)

	GetComplianceAPIV1InsightGroup(params *GetComplianceAPIV1InsightGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetComplianceAPIV1InsightGroupOK, error)

	GetComplianceAPIV1InsightInsightID(params *GetComplianceAPIV1InsightInsightIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetComplianceAPIV1InsightInsightIDOK, error)

	GetComplianceAPIV1InsightInsightIDTrend(params *GetComplianceAPIV1InsightInsightIDTrendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetComplianceAPIV1InsightInsightIDTrendOK, error)

	GetComplianceAPIV1MetadataInsightInsightID(params *GetComplianceAPIV1MetadataInsightInsightIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetComplianceAPIV1MetadataInsightInsightIDOK, error)

	GetComplianceAPIV1MetadataTagInsight(params *GetComplianceAPIV1MetadataTagInsightParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetComplianceAPIV1MetadataTagInsightOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	GetComplianceAPIV1Insight lists insights

	This API returns a list of insights based on specified filters. The API provides details of insights, including results during the specified time period for the specified connection.

Returns "all:provider" job results if connectionId is not defined.
*/
func (a *Client) GetComplianceAPIV1Insight(params *GetComplianceAPIV1InsightParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetComplianceAPIV1InsightOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetComplianceAPIV1InsightParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetComplianceAPIV1Insight",
		Method:             "GET",
		PathPattern:        "/compliance/api/v1/insight",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetComplianceAPIV1InsightReader{formats: a.formats},
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
	success, ok := result.(*GetComplianceAPIV1InsightOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetComplianceAPIV1Insight: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetComplianceAPIV1InsightGroup lists insight groups

	This API returns a list of insight groups based on specified filters. The API provides details of insights, including results during the specified time period for the specified connection.

Returns "all:provider" job results if connectionId is not defined.
*/
func (a *Client) GetComplianceAPIV1InsightGroup(params *GetComplianceAPIV1InsightGroupParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetComplianceAPIV1InsightGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetComplianceAPIV1InsightGroupParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetComplianceAPIV1InsightGroup",
		Method:             "GET",
		PathPattern:        "/compliance/api/v1/insight/group",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetComplianceAPIV1InsightGroupReader{formats: a.formats},
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
	success, ok := result.(*GetComplianceAPIV1InsightGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetComplianceAPIV1InsightGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetComplianceAPIV1InsightInsightID gets insight

	This API returns the specified insight with ID. The API provides details of the insight, including results during the specified time period for the specified connection.

Returns "all:provider" job results if connectionId is not defined.
*/
func (a *Client) GetComplianceAPIV1InsightInsightID(params *GetComplianceAPIV1InsightInsightIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetComplianceAPIV1InsightInsightIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetComplianceAPIV1InsightInsightIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetComplianceAPIV1InsightInsightID",
		Method:             "GET",
		PathPattern:        "/compliance/api/v1/insight/{insightId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetComplianceAPIV1InsightInsightIDReader{formats: a.formats},
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
	success, ok := result.(*GetComplianceAPIV1InsightInsightIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetComplianceAPIV1InsightInsightID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetComplianceAPIV1InsightInsightIDTrend gets insight trend

	This API allows users to retrieve insight results datapoints for a specified connection during a specified time period.

Returns "all:provider" job results if connectionId is not defined.
*/
func (a *Client) GetComplianceAPIV1InsightInsightIDTrend(params *GetComplianceAPIV1InsightInsightIDTrendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetComplianceAPIV1InsightInsightIDTrendOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetComplianceAPIV1InsightInsightIDTrendParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetComplianceAPIV1InsightInsightIDTrend",
		Method:             "GET",
		PathPattern:        "/compliance/api/v1/insight/{insightId}/trend",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetComplianceAPIV1InsightInsightIDTrendReader{formats: a.formats},
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
	success, ok := result.(*GetComplianceAPIV1InsightInsightIDTrendOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetComplianceAPIV1InsightInsightIDTrend: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetComplianceAPIV1MetadataInsightInsightID gets insight metadata

Get insight metadata by id
*/
func (a *Client) GetComplianceAPIV1MetadataInsightInsightID(params *GetComplianceAPIV1MetadataInsightInsightIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetComplianceAPIV1MetadataInsightInsightIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetComplianceAPIV1MetadataInsightInsightIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetComplianceAPIV1MetadataInsightInsightID",
		Method:             "GET",
		PathPattern:        "/compliance/api/v1/metadata/insight/{insightId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetComplianceAPIV1MetadataInsightInsightIDReader{formats: a.formats},
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
	success, ok := result.(*GetComplianceAPIV1MetadataInsightInsightIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetComplianceAPIV1MetadataInsightInsightID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetComplianceAPIV1MetadataTagInsight lists insights tag keys

This API allows users to retrieve a list of insights tag keys with their possible values.
*/
func (a *Client) GetComplianceAPIV1MetadataTagInsight(params *GetComplianceAPIV1MetadataTagInsightParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetComplianceAPIV1MetadataTagInsightOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetComplianceAPIV1MetadataTagInsightParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetComplianceAPIV1MetadataTagInsight",
		Method:             "GET",
		PathPattern:        "/compliance/api/v1/metadata/tag/insight",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetComplianceAPIV1MetadataTagInsightReader{formats: a.formats},
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
	success, ok := result.(*GetComplianceAPIV1MetadataTagInsightOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetComplianceAPIV1MetadataTagInsight: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
