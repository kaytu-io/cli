// Code generated by go-swagger; DO NOT EDIT.

package schedule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new schedule API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for schedule API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetScheduleAPIV1ComplianceReportLastCompleted(params *GetScheduleAPIV1ComplianceReportLastCompletedParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScheduleAPIV1ComplianceReportLastCompletedOK, error)

	GetScheduleAPIV1DescribeResourceJobsPending(params *GetScheduleAPIV1DescribeResourceJobsPendingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScheduleAPIV1DescribeResourceJobsPendingOK, error)

	GetScheduleAPIV1DescribeSourceJobsPending(params *GetScheduleAPIV1DescribeSourceJobsPendingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScheduleAPIV1DescribeSourceJobsPendingOK, error)

	GetScheduleAPIV1InsightJobsPending(params *GetScheduleAPIV1InsightJobsPendingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScheduleAPIV1InsightJobsPendingOK, error)

	GetScheduleAPIV1ResourceTypeProvider(params *GetScheduleAPIV1ResourceTypeProviderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScheduleAPIV1ResourceTypeProviderOK, error)

	GetScheduleAPIV1Sources(params *GetScheduleAPIV1SourcesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScheduleAPIV1SourcesOK, error)

	GetScheduleAPIV1SourcesSourceID(params *GetScheduleAPIV1SourcesSourceIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScheduleAPIV1SourcesSourceIDOK, error)

	GetScheduleAPIV1SourcesSourceIDJobsCompliance(params *GetScheduleAPIV1SourcesSourceIDJobsComplianceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScheduleAPIV1SourcesSourceIDJobsComplianceOK, error)

	GetScheduleAPIV1SourcesSourceIDJobsDescribe(params *GetScheduleAPIV1SourcesSourceIDJobsDescribeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScheduleAPIV1SourcesSourceIDJobsDescribeOK, error)

	GetScheduleAPIV1SummarizeJobsPending(params *GetScheduleAPIV1SummarizeJobsPendingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScheduleAPIV1SummarizeJobsPendingOK, error)

	PostScheduleAPIV1SourcesSourceIDJobsComplianceRefresh(params *PostScheduleAPIV1SourcesSourceIDJobsComplianceRefreshParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostScheduleAPIV1SourcesSourceIDJobsComplianceRefreshOK, error)

	PostScheduleAPIV1SourcesSourceIDJobsDescribeRefresh(params *PostScheduleAPIV1SourcesSourceIDJobsDescribeRefreshParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostScheduleAPIV1SourcesSourceIDJobsDescribeRefreshOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetScheduleAPIV1ComplianceReportLastCompleted gets last completed compliance report
*/
func (a *Client) GetScheduleAPIV1ComplianceReportLastCompleted(params *GetScheduleAPIV1ComplianceReportLastCompletedParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScheduleAPIV1ComplianceReportLastCompletedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetScheduleAPIV1ComplianceReportLastCompletedParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetScheduleAPIV1ComplianceReportLastCompleted",
		Method:             "GET",
		PathPattern:        "/schedule/api/v1/compliance/report/last/completed",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScheduleAPIV1ComplianceReportLastCompletedReader{formats: a.formats},
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
	success, ok := result.(*GetScheduleAPIV1ComplianceReportLastCompletedOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetScheduleAPIV1ComplianceReportLastCompleted: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetScheduleAPIV1DescribeResourceJobsPending listings describe resource jobs
*/
func (a *Client) GetScheduleAPIV1DescribeResourceJobsPending(params *GetScheduleAPIV1DescribeResourceJobsPendingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScheduleAPIV1DescribeResourceJobsPendingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetScheduleAPIV1DescribeResourceJobsPendingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetScheduleAPIV1DescribeResourceJobsPending",
		Method:             "GET",
		PathPattern:        "/schedule/api/v1/describe/resource/jobs/pending",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScheduleAPIV1DescribeResourceJobsPendingReader{formats: a.formats},
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
	success, ok := result.(*GetScheduleAPIV1DescribeResourceJobsPendingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetScheduleAPIV1DescribeResourceJobsPending: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetScheduleAPIV1DescribeSourceJobsPending listings describe source jobs
*/
func (a *Client) GetScheduleAPIV1DescribeSourceJobsPending(params *GetScheduleAPIV1DescribeSourceJobsPendingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScheduleAPIV1DescribeSourceJobsPendingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetScheduleAPIV1DescribeSourceJobsPendingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetScheduleAPIV1DescribeSourceJobsPending",
		Method:             "GET",
		PathPattern:        "/schedule/api/v1/describe/source/jobs/pending",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScheduleAPIV1DescribeSourceJobsPendingReader{formats: a.formats},
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
	success, ok := result.(*GetScheduleAPIV1DescribeSourceJobsPendingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetScheduleAPIV1DescribeSourceJobsPending: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetScheduleAPIV1InsightJobsPending listings insight jobs
*/
func (a *Client) GetScheduleAPIV1InsightJobsPending(params *GetScheduleAPIV1InsightJobsPendingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScheduleAPIV1InsightJobsPendingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetScheduleAPIV1InsightJobsPendingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetScheduleAPIV1InsightJobsPending",
		Method:             "GET",
		PathPattern:        "/schedule/api/v1/insight/jobs/pending",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScheduleAPIV1InsightJobsPendingReader{formats: a.formats},
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
	success, ok := result.(*GetScheduleAPIV1InsightJobsPendingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetScheduleAPIV1InsightJobsPending: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetScheduleAPIV1ResourceTypeProvider gets resource type by provider

get resource type by provider
*/
func (a *Client) GetScheduleAPIV1ResourceTypeProvider(params *GetScheduleAPIV1ResourceTypeProviderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScheduleAPIV1ResourceTypeProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetScheduleAPIV1ResourceTypeProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetScheduleAPIV1ResourceTypeProvider",
		Method:             "GET",
		PathPattern:        "/schedule/api/v1/resource_type/{provider}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScheduleAPIV1ResourceTypeProviderReader{formats: a.formats},
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
	success, ok := result.(*GetScheduleAPIV1ResourceTypeProviderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetScheduleAPIV1ResourceTypeProvider: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetScheduleAPIV1Sources lists sources

Getting all of Keibi sources
*/
func (a *Client) GetScheduleAPIV1Sources(params *GetScheduleAPIV1SourcesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScheduleAPIV1SourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetScheduleAPIV1SourcesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetScheduleAPIV1Sources",
		Method:             "GET",
		PathPattern:        "/schedule/api/v1/sources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScheduleAPIV1SourcesReader{formats: a.formats},
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
	success, ok := result.(*GetScheduleAPIV1SourcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetScheduleAPIV1Sources: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetScheduleAPIV1SourcesSourceID gets source by id

Getting Keibi source by id
*/
func (a *Client) GetScheduleAPIV1SourcesSourceID(params *GetScheduleAPIV1SourcesSourceIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScheduleAPIV1SourcesSourceIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetScheduleAPIV1SourcesSourceIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetScheduleAPIV1SourcesSourceID",
		Method:             "GET",
		PathPattern:        "/schedule/api/v1/sources/{source_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScheduleAPIV1SourcesSourceIDReader{formats: a.formats},
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
	success, ok := result.(*GetScheduleAPIV1SourcesSourceIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetScheduleAPIV1SourcesSourceID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetScheduleAPIV1SourcesSourceIDJobsCompliance lists source compliance reports

List source compliance reports
*/
func (a *Client) GetScheduleAPIV1SourcesSourceIDJobsCompliance(params *GetScheduleAPIV1SourcesSourceIDJobsComplianceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScheduleAPIV1SourcesSourceIDJobsComplianceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetScheduleAPIV1SourcesSourceIDJobsComplianceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetScheduleAPIV1SourcesSourceIDJobsCompliance",
		Method:             "GET",
		PathPattern:        "/schedule/api/v1/sources/{source_id}/jobs/compliance",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScheduleAPIV1SourcesSourceIDJobsComplianceReader{formats: a.formats},
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
	success, ok := result.(*GetScheduleAPIV1SourcesSourceIDJobsComplianceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetScheduleAPIV1SourcesSourceIDJobsCompliance: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetScheduleAPIV1SourcesSourceIDJobsDescribe lists source describe jobs

List source describe jobs
*/
func (a *Client) GetScheduleAPIV1SourcesSourceIDJobsDescribe(params *GetScheduleAPIV1SourcesSourceIDJobsDescribeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScheduleAPIV1SourcesSourceIDJobsDescribeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetScheduleAPIV1SourcesSourceIDJobsDescribeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetScheduleAPIV1SourcesSourceIDJobsDescribe",
		Method:             "GET",
		PathPattern:        "/schedule/api/v1/sources/{source_id}/jobs/describe",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScheduleAPIV1SourcesSourceIDJobsDescribeReader{formats: a.formats},
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
	success, ok := result.(*GetScheduleAPIV1SourcesSourceIDJobsDescribeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetScheduleAPIV1SourcesSourceIDJobsDescribe: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetScheduleAPIV1SummarizeJobsPending listings summarize jobs
*/
func (a *Client) GetScheduleAPIV1SummarizeJobsPending(params *GetScheduleAPIV1SummarizeJobsPendingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetScheduleAPIV1SummarizeJobsPendingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetScheduleAPIV1SummarizeJobsPendingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetScheduleAPIV1SummarizeJobsPending",
		Method:             "GET",
		PathPattern:        "/schedule/api/v1/summarize/jobs/pending",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScheduleAPIV1SummarizeJobsPendingReader{formats: a.formats},
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
	success, ok := result.(*GetScheduleAPIV1SummarizeJobsPendingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetScheduleAPIV1SummarizeJobsPending: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostScheduleAPIV1SourcesSourceIDJobsComplianceRefresh runs compliance report jobs

Run compliance report jobs
*/
func (a *Client) PostScheduleAPIV1SourcesSourceIDJobsComplianceRefresh(params *PostScheduleAPIV1SourcesSourceIDJobsComplianceRefreshParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostScheduleAPIV1SourcesSourceIDJobsComplianceRefreshOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostScheduleAPIV1SourcesSourceIDJobsComplianceRefreshParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostScheduleAPIV1SourcesSourceIDJobsComplianceRefresh",
		Method:             "POST",
		PathPattern:        "/schedule/api/v1/sources/{source_id}/jobs/compliance/refresh",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostScheduleAPIV1SourcesSourceIDJobsComplianceRefreshReader{formats: a.formats},
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
	success, ok := result.(*PostScheduleAPIV1SourcesSourceIDJobsComplianceRefreshOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostScheduleAPIV1SourcesSourceIDJobsComplianceRefresh: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostScheduleAPIV1SourcesSourceIDJobsDescribeRefresh runs describe jobs

Run describe jobs
*/
func (a *Client) PostScheduleAPIV1SourcesSourceIDJobsDescribeRefresh(params *PostScheduleAPIV1SourcesSourceIDJobsDescribeRefreshParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostScheduleAPIV1SourcesSourceIDJobsDescribeRefreshOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostScheduleAPIV1SourcesSourceIDJobsDescribeRefreshParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostScheduleAPIV1SourcesSourceIDJobsDescribeRefresh",
		Method:             "POST",
		PathPattern:        "/schedule/api/v1/sources/{source_id}/jobs/describe/refresh",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostScheduleAPIV1SourcesSourceIDJobsDescribeRefreshReader{formats: a.formats},
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
	success, ok := result.(*PostScheduleAPIV1SourcesSourceIDJobsDescribeRefreshOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostScheduleAPIV1SourcesSourceIDJobsDescribeRefresh: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
