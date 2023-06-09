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
	GetInventoryAPIV1ResourcesCount(params *GetInventoryAPIV1ResourcesCountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV1ResourcesCountOK, error)

	GetInventoryAPIV1ResourcesRegions(params *GetInventoryAPIV1ResourcesRegionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV1ResourcesRegionsOK, error)

	GetInventoryAPIV1ResourcesTopRegions(params *GetInventoryAPIV1ResourcesTopRegionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV1ResourcesTopRegionsOK, error)

	GetInventoryAPIV2ResourcesCompositionKey(params *GetInventoryAPIV2ResourcesCompositionKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ResourcesCompositionKeyOK, error)

	GetInventoryAPIV2ResourcesMetric(params *GetInventoryAPIV2ResourcesMetricParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ResourcesMetricOK, error)

	GetInventoryAPIV2ResourcesMetricResourceType(params *GetInventoryAPIV2ResourcesMetricResourceTypeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ResourcesMetricResourceTypeOK, error)

	GetInventoryAPIV2ResourcesTag(params *GetInventoryAPIV2ResourcesTagParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ResourcesTagOK, error)

	GetInventoryAPIV2ResourcesTagKey(params *GetInventoryAPIV2ResourcesTagKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ResourcesTagKeyOK, error)

	GetInventoryAPIV2ResourcesTrend(params *GetInventoryAPIV2ResourcesTrendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ResourcesTrendOK, error)

	GetInventoryAPIV2ServicesCompositionKey(params *GetInventoryAPIV2ServicesCompositionKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ServicesCompositionKeyOK, error)

	GetInventoryAPIV2ServicesCostTrend(params *GetInventoryAPIV2ServicesCostTrendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ServicesCostTrendOK, error)

	GetInventoryAPIV2ServicesMetric(params *GetInventoryAPIV2ServicesMetricParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ServicesMetricOK, error)

	GetInventoryAPIV2ServicesMetricServiceName(params *GetInventoryAPIV2ServicesMetricServiceNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ServicesMetricServiceNameOK, error)

	GetInventoryAPIV2ServicesTag(params *GetInventoryAPIV2ServicesTagParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ServicesTagOK, error)

	GetInventoryAPIV2ServicesTagKey(params *GetInventoryAPIV2ServicesTagKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ServicesTagKeyOK, error)

	PostInventoryAPIV1Resources(params *PostInventoryAPIV1ResourcesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostInventoryAPIV1ResourcesOK, error)

	PostInventoryAPIV1ResourcesAws(params *PostInventoryAPIV1ResourcesAwsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostInventoryAPIV1ResourcesAwsOK, error)

	PostInventoryAPIV1ResourcesAzure(params *PostInventoryAPIV1ResourcesAzureParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostInventoryAPIV1ResourcesAzureOK, error)

	PostInventoryAPIV1ResourcesFilters(params *PostInventoryAPIV1ResourcesFiltersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostInventoryAPIV1ResourcesFiltersOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetInventoryAPIV1ResourcesCount counts resources

Number of all resources
*/
func (a *Client) GetInventoryAPIV1ResourcesCount(params *GetInventoryAPIV1ResourcesCountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV1ResourcesCountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryAPIV1ResourcesCountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInventoryAPIV1ResourcesCount",
		Method:             "GET",
		PathPattern:        "/inventory/api/v1/resources/count",
		ProducesMediaTypes: []string{"application/json", "text/csv"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryAPIV1ResourcesCountReader{formats: a.formats},
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
	success, ok := result.(*GetInventoryAPIV1ResourcesCountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInventoryAPIV1ResourcesCount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInventoryAPIV1ResourcesRegions returns top n regions of specified provider by resource count
*/
func (a *Client) GetInventoryAPIV1ResourcesRegions(params *GetInventoryAPIV1ResourcesRegionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV1ResourcesRegionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryAPIV1ResourcesRegionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInventoryAPIV1ResourcesRegions",
		Method:             "GET",
		PathPattern:        "/inventory/api/v1/resources/regions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryAPIV1ResourcesRegionsReader{formats: a.formats},
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
	success, ok := result.(*GetInventoryAPIV1ResourcesRegionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInventoryAPIV1ResourcesRegions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInventoryAPIV1ResourcesTopRegions returns top n regions of specified provider by resource count
*/
func (a *Client) GetInventoryAPIV1ResourcesTopRegions(params *GetInventoryAPIV1ResourcesTopRegionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV1ResourcesTopRegionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryAPIV1ResourcesTopRegionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInventoryAPIV1ResourcesTopRegions",
		Method:             "GET",
		PathPattern:        "/inventory/api/v1/resources/top/regions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryAPIV1ResourcesTopRegionsReader{formats: a.formats},
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
	success, ok := result.(*GetInventoryAPIV1ResourcesTopRegionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInventoryAPIV1ResourcesTopRegions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInventoryAPIV2ResourcesCompositionKey returns tag values with most resources for the given key
*/
func (a *Client) GetInventoryAPIV2ResourcesCompositionKey(params *GetInventoryAPIV2ResourcesCompositionKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ResourcesCompositionKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryAPIV2ResourcesCompositionKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInventoryAPIV2ResourcesCompositionKey",
		Method:             "GET",
		PathPattern:        "/inventory/api/v2/resources/composition/{key}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryAPIV2ResourcesCompositionKeyReader{formats: a.formats},
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
	success, ok := result.(*GetInventoryAPIV2ResourcesCompositionKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInventoryAPIV2ResourcesCompositionKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInventoryAPIV2ResourcesMetric returns list of resource types with metrics of each type based on the given input filters
*/
func (a *Client) GetInventoryAPIV2ResourcesMetric(params *GetInventoryAPIV2ResourcesMetricParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ResourcesMetricOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryAPIV2ResourcesMetricParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInventoryAPIV2ResourcesMetric",
		Method:             "GET",
		PathPattern:        "/inventory/api/v2/resources/metric",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryAPIV2ResourcesMetricReader{formats: a.formats},
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
	success, ok := result.(*GetInventoryAPIV2ResourcesMetricOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInventoryAPIV2ResourcesMetric: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInventoryAPIV2ResourcesMetricResourceType returns resource type with metrics
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
GetInventoryAPIV2ResourcesTag returns list of the keys with possible values for filtering resources types
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
GetInventoryAPIV2ResourcesTagKey returns list of the possible values for filtering resources types with specified key
*/
func (a *Client) GetInventoryAPIV2ResourcesTagKey(params *GetInventoryAPIV2ResourcesTagKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ResourcesTagKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryAPIV2ResourcesTagKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInventoryAPIV2ResourcesTagKey",
		Method:             "GET",
		PathPattern:        "/inventory/api/v2/resources/tag/{key}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryAPIV2ResourcesTagKeyReader{formats: a.formats},
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
	success, ok := result.(*GetInventoryAPIV2ResourcesTagKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInventoryAPIV2ResourcesTagKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInventoryAPIV2ResourcesTrend returns list of resource counts over the course of the specified time frame based on the given input filters
*/
func (a *Client) GetInventoryAPIV2ResourcesTrend(params *GetInventoryAPIV2ResourcesTrendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ResourcesTrendOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryAPIV2ResourcesTrendParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInventoryAPIV2ResourcesTrend",
		Method:             "GET",
		PathPattern:        "/inventory/api/v2/resources/trend",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryAPIV2ResourcesTrendReader{formats: a.formats},
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
	success, ok := result.(*GetInventoryAPIV2ResourcesTrendOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInventoryAPIV2ResourcesTrend: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInventoryAPIV2ServicesCompositionKey returns tag values with most cost for the given key
*/
func (a *Client) GetInventoryAPIV2ServicesCompositionKey(params *GetInventoryAPIV2ServicesCompositionKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ServicesCompositionKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryAPIV2ServicesCompositionKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInventoryAPIV2ServicesCompositionKey",
		Method:             "GET",
		PathPattern:        "/inventory/api/v2/services/composition/{key}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryAPIV2ServicesCompositionKeyReader{formats: a.formats},
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
	success, ok := result.(*GetInventoryAPIV2ServicesCompositionKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInventoryAPIV2ServicesCompositionKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInventoryAPIV2ServicesCostTrend returns list of costs over the course of the specified time frame based on the given input filters
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

/*
GetInventoryAPIV2ServicesMetric returns list of services with their metrics based on the given input filters
*/
func (a *Client) GetInventoryAPIV2ServicesMetric(params *GetInventoryAPIV2ServicesMetricParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ServicesMetricOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryAPIV2ServicesMetricParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInventoryAPIV2ServicesMetric",
		Method:             "GET",
		PathPattern:        "/inventory/api/v2/services/metric",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryAPIV2ServicesMetricReader{formats: a.formats},
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
	success, ok := result.(*GetInventoryAPIV2ServicesMetricOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInventoryAPIV2ServicesMetric: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInventoryAPIV2ServicesMetricServiceName returns the service with metrics for the given service name
*/
func (a *Client) GetInventoryAPIV2ServicesMetricServiceName(params *GetInventoryAPIV2ServicesMetricServiceNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ServicesMetricServiceNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryAPIV2ServicesMetricServiceNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInventoryAPIV2ServicesMetricServiceName",
		Method:             "GET",
		PathPattern:        "/inventory/api/v2/services/metric/{serviceName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryAPIV2ServicesMetricServiceNameReader{formats: a.formats},
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
	success, ok := result.(*GetInventoryAPIV2ServicesMetricServiceNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInventoryAPIV2ServicesMetricServiceName: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInventoryAPIV2ServicesTag returns list of the keys with possible values for filtering services
*/
func (a *Client) GetInventoryAPIV2ServicesTag(params *GetInventoryAPIV2ServicesTagParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ServicesTagOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryAPIV2ServicesTagParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInventoryAPIV2ServicesTag",
		Method:             "GET",
		PathPattern:        "/inventory/api/v2/services/tag",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryAPIV2ServicesTagReader{formats: a.formats},
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
	success, ok := result.(*GetInventoryAPIV2ServicesTagOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInventoryAPIV2ServicesTag: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInventoryAPIV2ServicesTagKey returns list of the possible values for filtering services with specified key
*/
func (a *Client) GetInventoryAPIV2ServicesTagKey(params *GetInventoryAPIV2ServicesTagKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ServicesTagKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryAPIV2ServicesTagKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInventoryAPIV2ServicesTagKey",
		Method:             "GET",
		PathPattern:        "/inventory/api/v2/services/tag/{key}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryAPIV2ServicesTagKeyReader{formats: a.formats},
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
	success, ok := result.(*GetInventoryAPIV2ServicesTagKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInventoryAPIV2ServicesTagKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	PostInventoryAPIV1Resources gets resources

	Getting all cloud providers resources by filters.

In order to get the results in CSV format, Accepts header must be filled with `text/csv` value.
Note that csv output doesn't process pagination and returns first 5000 records.
If sort by is empty, result will be sorted by the first column in ascending order.
*/
func (a *Client) PostInventoryAPIV1Resources(params *PostInventoryAPIV1ResourcesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostInventoryAPIV1ResourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostInventoryAPIV1ResourcesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostInventoryAPIV1Resources",
		Method:             "POST",
		PathPattern:        "/inventory/api/v1/resources",
		ProducesMediaTypes: []string{"application/json", "text/csv"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostInventoryAPIV1ResourcesReader{formats: a.formats},
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
	success, ok := result.(*PostInventoryAPIV1ResourcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostInventoryAPIV1Resources: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	PostInventoryAPIV1ResourcesAws gets a w s resources

	Getting AWS resources by filters.

In order to get the results in CSV format, Accepts header must be filled with `text/csv` value.
Note that csv output doesn't process pagination and returns first 5000 records.
If sort by is empty, result will be sorted by the first column in ascending order.
*/
func (a *Client) PostInventoryAPIV1ResourcesAws(params *PostInventoryAPIV1ResourcesAwsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostInventoryAPIV1ResourcesAwsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostInventoryAPIV1ResourcesAwsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostInventoryAPIV1ResourcesAws",
		Method:             "POST",
		PathPattern:        "/inventory/api/v1/resources/aws",
		ProducesMediaTypes: []string{"application/json", "text/csv"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostInventoryAPIV1ResourcesAwsReader{formats: a.formats},
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
	success, ok := result.(*PostInventoryAPIV1ResourcesAwsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostInventoryAPIV1ResourcesAws: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	PostInventoryAPIV1ResourcesAzure gets azure resources

	Getting Azure resources by filters.

In order to get the results in CSV format, Accepts header must be filled with `text/csv` value.
Note that csv output doesn't process pagination and returns first 5000 records.
If sort by is empty, result will be sorted by the first column in ascending order.
*/
func (a *Client) PostInventoryAPIV1ResourcesAzure(params *PostInventoryAPIV1ResourcesAzureParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostInventoryAPIV1ResourcesAzureOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostInventoryAPIV1ResourcesAzureParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostInventoryAPIV1ResourcesAzure",
		Method:             "POST",
		PathPattern:        "/inventory/api/v1/resources/azure",
		ProducesMediaTypes: []string{"application/json", "text/csv"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostInventoryAPIV1ResourcesAzureReader{formats: a.formats},
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
	success, ok := result.(*PostInventoryAPIV1ResourcesAzureOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostInventoryAPIV1ResourcesAzure: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostInventoryAPIV1ResourcesFilters gets resource filters

Getting resource filters by filters.
*/
func (a *Client) PostInventoryAPIV1ResourcesFilters(params *PostInventoryAPIV1ResourcesFiltersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostInventoryAPIV1ResourcesFiltersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostInventoryAPIV1ResourcesFiltersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostInventoryAPIV1ResourcesFilters",
		Method:             "POST",
		PathPattern:        "/inventory/api/v1/resources/filters",
		ProducesMediaTypes: []string{"application/json", "text/csv"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostInventoryAPIV1ResourcesFiltersReader{formats: a.formats},
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
	success, ok := result.(*PostInventoryAPIV1ResourcesFiltersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostInventoryAPIV1ResourcesFilters: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
