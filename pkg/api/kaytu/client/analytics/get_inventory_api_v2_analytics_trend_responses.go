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

// GetInventoryAPIV2AnalyticsTrendReader is a Reader for the GetInventoryAPIV2AnalyticsTrend structure.
type GetInventoryAPIV2AnalyticsTrendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventoryAPIV2AnalyticsTrendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInventoryAPIV2AnalyticsTrendOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /inventory/api/v2/analytics/trend] GetInventoryAPIV2AnalyticsTrend", response, response.Code())
	}
}

// NewGetInventoryAPIV2AnalyticsTrendOK creates a GetInventoryAPIV2AnalyticsTrendOK with default headers values
func NewGetInventoryAPIV2AnalyticsTrendOK() *GetInventoryAPIV2AnalyticsTrendOK {
	return &GetInventoryAPIV2AnalyticsTrendOK{}
}

/*
GetInventoryAPIV2AnalyticsTrendOK describes a response with status code 200, with default header values.

OK
*/
type GetInventoryAPIV2AnalyticsTrendOK struct {
	Payload []*models.GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceTypeTrendDatapoint
}

// IsSuccess returns true when this get inventory Api v2 analytics trend o k response has a 2xx status code
func (o *GetInventoryAPIV2AnalyticsTrendOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get inventory Api v2 analytics trend o k response has a 3xx status code
func (o *GetInventoryAPIV2AnalyticsTrendOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get inventory Api v2 analytics trend o k response has a 4xx status code
func (o *GetInventoryAPIV2AnalyticsTrendOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get inventory Api v2 analytics trend o k response has a 5xx status code
func (o *GetInventoryAPIV2AnalyticsTrendOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get inventory Api v2 analytics trend o k response a status code equal to that given
func (o *GetInventoryAPIV2AnalyticsTrendOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get inventory Api v2 analytics trend o k response
func (o *GetInventoryAPIV2AnalyticsTrendOK) Code() int {
	return 200
}

func (o *GetInventoryAPIV2AnalyticsTrendOK) Error() string {
	return fmt.Sprintf("[GET /inventory/api/v2/analytics/trend][%d] getInventoryApiV2AnalyticsTrendOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2AnalyticsTrendOK) String() string {
	return fmt.Sprintf("[GET /inventory/api/v2/analytics/trend][%d] getInventoryApiV2AnalyticsTrendOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2AnalyticsTrendOK) GetPayload() []*models.GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceTypeTrendDatapoint {
	return o.Payload
}

func (o *GetInventoryAPIV2AnalyticsTrendOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
