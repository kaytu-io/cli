// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

// GetInventoryAPIV2AnalyticsSpendTrendReader is a Reader for the GetInventoryAPIV2AnalyticsSpendTrend structure.
type GetInventoryAPIV2AnalyticsSpendTrendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventoryAPIV2AnalyticsSpendTrendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInventoryAPIV2AnalyticsSpendTrendOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /inventory/api/v2/analytics/spend/trend] GetInventoryAPIV2AnalyticsSpendTrend", response, response.Code())
	}
}

// NewGetInventoryAPIV2AnalyticsSpendTrendOK creates a GetInventoryAPIV2AnalyticsSpendTrendOK with default headers values
func NewGetInventoryAPIV2AnalyticsSpendTrendOK() *GetInventoryAPIV2AnalyticsSpendTrendOK {
	return &GetInventoryAPIV2AnalyticsSpendTrendOK{}
}

/*
GetInventoryAPIV2AnalyticsSpendTrendOK describes a response with status code 200, with default header values.

OK
*/
type GetInventoryAPIV2AnalyticsSpendTrendOK struct {
	Payload []*models.GithubComKaytuIoKaytuEnginePkgInventoryAPICostTrendDatapoint
}

// IsSuccess returns true when this get inventory Api v2 analytics spend trend o k response has a 2xx status code
func (o *GetInventoryAPIV2AnalyticsSpendTrendOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get inventory Api v2 analytics spend trend o k response has a 3xx status code
func (o *GetInventoryAPIV2AnalyticsSpendTrendOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get inventory Api v2 analytics spend trend o k response has a 4xx status code
func (o *GetInventoryAPIV2AnalyticsSpendTrendOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get inventory Api v2 analytics spend trend o k response has a 5xx status code
func (o *GetInventoryAPIV2AnalyticsSpendTrendOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get inventory Api v2 analytics spend trend o k response a status code equal to that given
func (o *GetInventoryAPIV2AnalyticsSpendTrendOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get inventory Api v2 analytics spend trend o k response
func (o *GetInventoryAPIV2AnalyticsSpendTrendOK) Code() int {
	return 200
}

func (o *GetInventoryAPIV2AnalyticsSpendTrendOK) Error() string {
	return fmt.Sprintf("[GET /inventory/api/v2/analytics/spend/trend][%d] getInventoryApiV2AnalyticsSpendTrendOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2AnalyticsSpendTrendOK) String() string {
	return fmt.Sprintf("[GET /inventory/api/v2/analytics/spend/trend][%d] getInventoryApiV2AnalyticsSpendTrendOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2AnalyticsSpendTrendOK) GetPayload() []*models.GithubComKaytuIoKaytuEnginePkgInventoryAPICostTrendDatapoint {
	return o.Payload
}

func (o *GetInventoryAPIV2AnalyticsSpendTrendOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
