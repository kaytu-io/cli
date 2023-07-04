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
	GetInventoryAPIV1ResourcesRegions(params *GetInventoryAPIV1ResourcesRegionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV1ResourcesRegionsOK, error)

	GetInventoryAPIV1ResourcesTopRegions(params *GetInventoryAPIV1ResourcesTopRegionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV1ResourcesTopRegionsOK, error)

	GetInventoryAPIV2ResourcesCount(params *GetInventoryAPIV2ResourcesCountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ResourcesCountOK, error)

	GetInventoryAPIV2ResourcesRegionsComposition(params *GetInventoryAPIV2ResourcesRegionsCompositionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ResourcesRegionsCompositionOK, error)

	GetInventoryAPIV2ResourcesRegionsSummary(params *GetInventoryAPIV2ResourcesRegionsSummaryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ResourcesRegionsSummaryOK, error)

	GetInventoryAPIV2ResourcesRegionsTrend(params *GetInventoryAPIV2ResourcesRegionsTrendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ResourcesRegionsTrendOK, error)

	PostInventoryAPIV1Resource(params *PostInventoryAPIV1ResourceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostInventoryAPIV1ResourceOK, error)

	PostInventoryAPIV1Resources(params *PostInventoryAPIV1ResourcesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostInventoryAPIV1ResourcesOK, error)

	PostInventoryAPIV1ResourcesAws(params *PostInventoryAPIV1ResourcesAwsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostInventoryAPIV1ResourcesAwsOK, error)

	PostInventoryAPIV1ResourcesAzure(params *PostInventoryAPIV1ResourcesAzureParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostInventoryAPIV1ResourcesAzureOK, error)

	PostInventoryAPIV1ResourcesFilters(params *PostInventoryAPIV1ResourcesFiltersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostInventoryAPIV1ResourcesFiltersOK, error)

	SetTransport(transport runtime.ClientTransport)
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
GetInventoryAPIV2ResourcesCount counts resources

Number of all resources
*/
func (a *Client) GetInventoryAPIV2ResourcesCount(params *GetInventoryAPIV2ResourcesCountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ResourcesCountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryAPIV2ResourcesCountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInventoryAPIV2ResourcesCount",
		Method:             "GET",
		PathPattern:        "/inventory/api/v2/resources/count",
		ProducesMediaTypes: []string{"application/json", "text/csv"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryAPIV2ResourcesCountReader{formats: a.formats},
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
	success, ok := result.(*GetInventoryAPIV2ResourcesCountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInventoryAPIV2ResourcesCount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInventoryAPIV2ResourcesRegionsComposition lists resources regions composition

Returns list of top regions per given connector type and connection IDs
*/
func (a *Client) GetInventoryAPIV2ResourcesRegionsComposition(params *GetInventoryAPIV2ResourcesRegionsCompositionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ResourcesRegionsCompositionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryAPIV2ResourcesRegionsCompositionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInventoryAPIV2ResourcesRegionsComposition",
		Method:             "GET",
		PathPattern:        "/inventory/api/v2/resources/regions/composition",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryAPIV2ResourcesRegionsCompositionReader{formats: a.formats},
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
	success, ok := result.(*GetInventoryAPIV2ResourcesRegionsCompositionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInventoryAPIV2ResourcesRegionsComposition: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInventoryAPIV2ResourcesRegionsSummary lists regions summary

Returns list of regions resources summary
*/
func (a *Client) GetInventoryAPIV2ResourcesRegionsSummary(params *GetInventoryAPIV2ResourcesRegionsSummaryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ResourcesRegionsSummaryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryAPIV2ResourcesRegionsSummaryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInventoryAPIV2ResourcesRegionsSummary",
		Method:             "GET",
		PathPattern:        "/inventory/api/v2/resources/regions/summary",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryAPIV2ResourcesRegionsSummaryReader{formats: a.formats},
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
	success, ok := result.(*GetInventoryAPIV2ResourcesRegionsSummaryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInventoryAPIV2ResourcesRegionsSummary: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInventoryAPIV2ResourcesRegionsTrend returns trend of resources count in given regions

Returns list of regions resources summary
*/
func (a *Client) GetInventoryAPIV2ResourcesRegionsTrend(params *GetInventoryAPIV2ResourcesRegionsTrendParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInventoryAPIV2ResourcesRegionsTrendOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryAPIV2ResourcesRegionsTrendParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInventoryAPIV2ResourcesRegionsTrend",
		Method:             "GET",
		PathPattern:        "/inventory/api/v2/resources/regions/trend",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryAPIV2ResourcesRegionsTrendReader{formats: a.formats},
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
	success, ok := result.(*GetInventoryAPIV2ResourcesRegionsTrendOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInventoryAPIV2ResourcesRegionsTrend: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostInventoryAPIV1Resource gets details of a resource

Getting resource details by id and resource type
*/
func (a *Client) PostInventoryAPIV1Resource(params *PostInventoryAPIV1ResourceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostInventoryAPIV1ResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostInventoryAPIV1ResourceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostInventoryAPIV1Resource",
		Method:             "POST",
		PathPattern:        "/inventory/api/v1/resource",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostInventoryAPIV1ResourceReader{formats: a.formats},
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
	success, ok := result.(*PostInventoryAPIV1ResourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostInventoryAPIV1Resource: API contract not enforced by server. Client expected to get an error, but got: %T", result)
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
