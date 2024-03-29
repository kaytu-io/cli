// Code generated by go-swagger; DO NOT EDIT.

package analytics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

// GetInventoryAPIV2AnalyticsMetricsListReader is a Reader for the GetInventoryAPIV2AnalyticsMetricsList structure.
type GetInventoryAPIV2AnalyticsMetricsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventoryAPIV2AnalyticsMetricsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInventoryAPIV2AnalyticsMetricsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /inventory/api/v2/analytics/metrics/list] GetInventoryAPIV2AnalyticsMetricsList", response, response.Code())
	}
}

// NewGetInventoryAPIV2AnalyticsMetricsListOK creates a GetInventoryAPIV2AnalyticsMetricsListOK with default headers values
func NewGetInventoryAPIV2AnalyticsMetricsListOK() *GetInventoryAPIV2AnalyticsMetricsListOK {
	return &GetInventoryAPIV2AnalyticsMetricsListOK{}
}

/*
GetInventoryAPIV2AnalyticsMetricsListOK describes a response with status code 200, with default header values.

OK
*/
type GetInventoryAPIV2AnalyticsMetricsListOK struct {
	Payload []*models.GithubComKaytuIoKaytuEnginePkgInventoryAPIAnalyticsMetric
}

// IsSuccess returns true when this get inventory Api v2 analytics metrics list o k response has a 2xx status code
func (o *GetInventoryAPIV2AnalyticsMetricsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get inventory Api v2 analytics metrics list o k response has a 3xx status code
func (o *GetInventoryAPIV2AnalyticsMetricsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get inventory Api v2 analytics metrics list o k response has a 4xx status code
func (o *GetInventoryAPIV2AnalyticsMetricsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get inventory Api v2 analytics metrics list o k response has a 5xx status code
func (o *GetInventoryAPIV2AnalyticsMetricsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get inventory Api v2 analytics metrics list o k response a status code equal to that given
func (o *GetInventoryAPIV2AnalyticsMetricsListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get inventory Api v2 analytics metrics list o k response
func (o *GetInventoryAPIV2AnalyticsMetricsListOK) Code() int {
	return 200
}

func (o *GetInventoryAPIV2AnalyticsMetricsListOK) Error() string {
	return fmt.Sprintf("[GET /inventory/api/v2/analytics/metrics/list][%d] getInventoryApiV2AnalyticsMetricsListOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2AnalyticsMetricsListOK) String() string {
	return fmt.Sprintf("[GET /inventory/api/v2/analytics/metrics/list][%d] getInventoryApiV2AnalyticsMetricsListOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2AnalyticsMetricsListOK) GetPayload() []*models.GithubComKaytuIoKaytuEnginePkgInventoryAPIAnalyticsMetric {
	return o.Payload
}

func (o *GetInventoryAPIV2AnalyticsMetricsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
