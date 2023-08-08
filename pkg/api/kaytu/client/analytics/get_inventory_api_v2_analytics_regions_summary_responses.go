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

// GetInventoryAPIV2AnalyticsRegionsSummaryReader is a Reader for the GetInventoryAPIV2AnalyticsRegionsSummary structure.
type GetInventoryAPIV2AnalyticsRegionsSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventoryAPIV2AnalyticsRegionsSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInventoryAPIV2AnalyticsRegionsSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /inventory/api/v2/analytics/regions/summary] GetInventoryAPIV2AnalyticsRegionsSummary", response, response.Code())
	}
}

// NewGetInventoryAPIV2AnalyticsRegionsSummaryOK creates a GetInventoryAPIV2AnalyticsRegionsSummaryOK with default headers values
func NewGetInventoryAPIV2AnalyticsRegionsSummaryOK() *GetInventoryAPIV2AnalyticsRegionsSummaryOK {
	return &GetInventoryAPIV2AnalyticsRegionsSummaryOK{}
}

/*
GetInventoryAPIV2AnalyticsRegionsSummaryOK describes a response with status code 200, with default header values.

OK
*/
type GetInventoryAPIV2AnalyticsRegionsSummaryOK struct {
	Payload *models.GithubComKaytuIoKaytuEnginePkgInventoryAPIRegionsResourceCountResponse
}

// IsSuccess returns true when this get inventory Api v2 analytics regions summary o k response has a 2xx status code
func (o *GetInventoryAPIV2AnalyticsRegionsSummaryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get inventory Api v2 analytics regions summary o k response has a 3xx status code
func (o *GetInventoryAPIV2AnalyticsRegionsSummaryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get inventory Api v2 analytics regions summary o k response has a 4xx status code
func (o *GetInventoryAPIV2AnalyticsRegionsSummaryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get inventory Api v2 analytics regions summary o k response has a 5xx status code
func (o *GetInventoryAPIV2AnalyticsRegionsSummaryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get inventory Api v2 analytics regions summary o k response a status code equal to that given
func (o *GetInventoryAPIV2AnalyticsRegionsSummaryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get inventory Api v2 analytics regions summary o k response
func (o *GetInventoryAPIV2AnalyticsRegionsSummaryOK) Code() int {
	return 200
}

func (o *GetInventoryAPIV2AnalyticsRegionsSummaryOK) Error() string {
	return fmt.Sprintf("[GET /inventory/api/v2/analytics/regions/summary][%d] getInventoryApiV2AnalyticsRegionsSummaryOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2AnalyticsRegionsSummaryOK) String() string {
	return fmt.Sprintf("[GET /inventory/api/v2/analytics/regions/summary][%d] getInventoryApiV2AnalyticsRegionsSummaryOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2AnalyticsRegionsSummaryOK) GetPayload() *models.GithubComKaytuIoKaytuEnginePkgInventoryAPIRegionsResourceCountResponse {
	return o.Payload
}

func (o *GetInventoryAPIV2AnalyticsRegionsSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GithubComKaytuIoKaytuEnginePkgInventoryAPIRegionsResourceCountResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}