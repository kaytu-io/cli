// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new inventory API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for inventory API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetInventoryAPIV2AnalyticsSpendComposition(params *GetInventoryAPIV2AnalyticsSpendCompositionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2AnalyticsSpendCompositionOK, error)

	GetInventoryAPIV2AnalyticsSpendMetric(params *GetInventoryAPIV2AnalyticsSpendMetricParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2AnalyticsSpendMetricOK, error)

	GetInventoryAPIV2AnalyticsSpendMetricsTrend(params *GetInventoryAPIV2AnalyticsSpendMetricsTrendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2AnalyticsSpendMetricsTrendOK, error)

	GetInventoryAPIV2AnalyticsSpendTrend(params *GetInventoryAPIV2AnalyticsSpendTrendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2AnalyticsSpendTrendOK, error)

	GetInventoryAPIV2CostComposition(params *GetInventoryAPIV2CostCompositionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2CostCompositionOK, error)

	GetInventoryAPIV2CostMetric(params *GetInventoryAPIV2CostMetricParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2CostMetricOK, error)

	GetInventoryAPIV2CostTrend(params *GetInventoryAPIV2CostTrendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2CostTrendOK, error)

	GetInventoryAPIV2ResourcesMetricResourceType(params *GetInventoryAPIV2ResourcesMetricResourceTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ResourcesMetricResourceTypeOK, error)

	GetInventoryAPIV2ResourcesTag(params *GetInventoryAPIV2ResourcesTagParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ResourcesTagOK, error)

	GetInventoryAPIV2ServicesCostTrend(params *GetInventoryAPIV2ServicesCostTrendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ServicesCostTrendOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetInventoryAPIV2AnalyticsSpendComposition lists cost composition

This API allows users to retrieve the cost composition with respect to specified filters. The API returns information such as the total cost for the given time range, and the top services by cost.
*/
func (a *Client) GetInventoryAPIV2AnalyticsSpendComposition(params *GetInventoryAPIV2AnalyticsSpendCompositionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2AnalyticsSpendCompositionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryAPIV2AnalyticsSpendCompositionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInventoryAPIV2AnalyticsSpendComposition",
		Method:             "GET",
		PathPattern:        "/inventory/api/v2/analytics/spend/composition",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryAPIV2AnalyticsSpendCompositionReader{formats: a.formats},
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
	success, ok := result.(*GetInventoryAPIV2AnalyticsSpendCompositionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInventoryAPIV2AnalyticsSpendComposition: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInventoryAPIV2AnalyticsSpendMetric lists spend metrics

This API allows users to retrieve cost metrics with respect to specified filters. The API returns information such as the total cost and costs per each service based on the specified filters.
*/
func (a *Client) GetInventoryAPIV2AnalyticsSpendMetric(params *GetInventoryAPIV2AnalyticsSpendMetricParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2AnalyticsSpendMetricOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryAPIV2AnalyticsSpendMetricParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInventoryAPIV2AnalyticsSpendMetric",
		Method:             "GET",
		PathPattern:        "/inventory/api/v2/analytics/spend/metric",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryAPIV2AnalyticsSpendMetricReader{formats: a.formats},
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
	success, ok := result.(*GetInventoryAPIV2AnalyticsSpendMetricOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInventoryAPIV2AnalyticsSpendMetric: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInventoryAPIV2AnalyticsSpendMetricsTrend gets cost trend

This API allows users to retrieve a list of costs over the course of the specified time frame based on the given input filters. If startTime and endTime are empty, the API returns the last month trend.
*/
func (a *Client) GetInventoryAPIV2AnalyticsSpendMetricsTrend(params *GetInventoryAPIV2AnalyticsSpendMetricsTrendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2AnalyticsSpendMetricsTrendOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryAPIV2AnalyticsSpendMetricsTrendParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInventoryAPIV2AnalyticsSpendMetricsTrend",
		Method:             "GET",
		PathPattern:        "/inventory/api/v2/analytics/spend/metrics/trend",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryAPIV2AnalyticsSpendMetricsTrendReader{formats: a.formats},
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
	success, ok := result.(*GetInventoryAPIV2AnalyticsSpendMetricsTrendOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInventoryAPIV2AnalyticsSpendMetricsTrend: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInventoryAPIV2AnalyticsSpendTrend gets cost trend

This API allows users to retrieve a list of costs over the course of the specified time frame based on the given input filters. If startTime and endTime are empty, the API returns the last month trend.
*/
func (a *Client) GetInventoryAPIV2AnalyticsSpendTrend(params *GetInventoryAPIV2AnalyticsSpendTrendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2AnalyticsSpendTrendOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryAPIV2AnalyticsSpendTrendParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInventoryAPIV2AnalyticsSpendTrend",
		Method:             "GET",
		PathPattern:        "/inventory/api/v2/analytics/spend/trend",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryAPIV2AnalyticsSpendTrendReader{formats: a.formats},
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
	success, ok := result.(*GetInventoryAPIV2AnalyticsSpendTrendOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInventoryAPIV2AnalyticsSpendTrend: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInventoryAPIV2CostComposition lists cost composition

This API allows users to retrieve the cost composition with respect to specified filters. The API returns information such as the total cost for the given time range, and the top services by cost.
*/
func (a *Client) GetInventoryAPIV2CostComposition(params *GetInventoryAPIV2CostCompositionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2CostCompositionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryAPIV2CostCompositionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInventoryAPIV2CostComposition",
		Method:             "GET",
		PathPattern:        "/inventory/api/v2/cost/composition",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryAPIV2CostCompositionReader{formats: a.formats},
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
	success, ok := result.(*GetInventoryAPIV2CostCompositionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInventoryAPIV2CostComposition: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInventoryAPIV2CostMetric lists cost metrics

This API allows users to retrieve cost metrics with respect to specified filters. The API returns information such as the total cost and costs per each service based on the specified filters.
*/
func (a *Client) GetInventoryAPIV2CostMetric(params *GetInventoryAPIV2CostMetricParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2CostMetricOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryAPIV2CostMetricParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInventoryAPIV2CostMetric",
		Method:             "GET",
		PathPattern:        "/inventory/api/v2/cost/metric",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryAPIV2CostMetricReader{formats: a.formats},
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
	success, ok := result.(*GetInventoryAPIV2CostMetricOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInventoryAPIV2CostMetric: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInventoryAPIV2CostTrend gets cost trend

This API allows users to retrieve a list of costs over the course of the specified time frame based on the given input filters. If startTime and endTime are empty, the API returns the last month trend.
*/
func (a *Client) GetInventoryAPIV2CostTrend(params *GetInventoryAPIV2CostTrendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2CostTrendOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryAPIV2CostTrendParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInventoryAPIV2CostTrend",
		Method:             "GET",
		PathPattern:        "/inventory/api/v2/cost/trend",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryAPIV2CostTrendReader{formats: a.formats},
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
	success, ok := result.(*GetInventoryAPIV2CostTrendOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInventoryAPIV2CostTrend: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInventoryAPIV2ResourcesMetricResourceType gets resource metrics

This API allows users to retrieve metrics for a specific resource type.
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
GetInventoryAPIV2ResourcesTag lists resourcetype tags

This API allows users to retrieve a list of tag keys with their possible values for all resource types.
*/
func (a *Client) GetInventoryAPIV2ResourcesTag(params *GetInventoryAPIV2ResourcesTagParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ResourcesTagOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryAPIV2ResourcesTagParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInventoryAPIV2ResourcesTag",
		Method:             "GET",
		PathPattern:        "/inventory/api/v2/resources/tag",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryAPIV2ResourcesTagReader{formats: a.formats},
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
	success, ok := result.(*GetInventoryAPIV2ResourcesTagOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInventoryAPIV2ResourcesTag: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInventoryAPIV2ServicesCostTrend gets services cost trend

This API allows users to retrieve a list of costs over the course of the specified time frame for the given services. If startTime and endTime are empty, the API returns the last month trend.
*/
func (a *Client) GetInventoryAPIV2ServicesCostTrend(params *GetInventoryAPIV2ServicesCostTrendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ServicesCostTrendOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryAPIV2ServicesCostTrendParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInventoryAPIV2ServicesCostTrend",
		Method:             "GET",
		PathPattern:        "/inventory/api/v2/services/cost/trend",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryAPIV2ServicesCostTrendReader{formats: a.formats},
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
	success, ok := result.(*GetInventoryAPIV2ServicesCostTrendOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInventoryAPIV2ServicesCostTrend: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
