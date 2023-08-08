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

// GetInventoryAPIV2AnalyticsCompositionKeyReader is a Reader for the GetInventoryAPIV2AnalyticsCompositionKey structure.
type GetInventoryAPIV2AnalyticsCompositionKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventoryAPIV2AnalyticsCompositionKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInventoryAPIV2AnalyticsCompositionKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /inventory/api/v2/analytics/composition/{key}] GetInventoryAPIV2AnalyticsCompositionKey", response, response.Code())
	}
}

// NewGetInventoryAPIV2AnalyticsCompositionKeyOK creates a GetInventoryAPIV2AnalyticsCompositionKeyOK with default headers values
func NewGetInventoryAPIV2AnalyticsCompositionKeyOK() *GetInventoryAPIV2AnalyticsCompositionKeyOK {
	return &GetInventoryAPIV2AnalyticsCompositionKeyOK{}
}

/*
GetInventoryAPIV2AnalyticsCompositionKeyOK describes a response with status code 200, with default header values.

OK
*/
type GetInventoryAPIV2AnalyticsCompositionKeyOK struct {
	Payload *models.GithubComKaytuIoKaytuEnginePkgInventoryAPIListResourceTypeCompositionResponse
}

// IsSuccess returns true when this get inventory Api v2 analytics composition key o k response has a 2xx status code
func (o *GetInventoryAPIV2AnalyticsCompositionKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get inventory Api v2 analytics composition key o k response has a 3xx status code
func (o *GetInventoryAPIV2AnalyticsCompositionKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get inventory Api v2 analytics composition key o k response has a 4xx status code
func (o *GetInventoryAPIV2AnalyticsCompositionKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get inventory Api v2 analytics composition key o k response has a 5xx status code
func (o *GetInventoryAPIV2AnalyticsCompositionKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get inventory Api v2 analytics composition key o k response a status code equal to that given
func (o *GetInventoryAPIV2AnalyticsCompositionKeyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get inventory Api v2 analytics composition key o k response
func (o *GetInventoryAPIV2AnalyticsCompositionKeyOK) Code() int {
	return 200
}

func (o *GetInventoryAPIV2AnalyticsCompositionKeyOK) Error() string {
	return fmt.Sprintf("[GET /inventory/api/v2/analytics/composition/{key}][%d] getInventoryApiV2AnalyticsCompositionKeyOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2AnalyticsCompositionKeyOK) String() string {
	return fmt.Sprintf("[GET /inventory/api/v2/analytics/composition/{key}][%d] getInventoryApiV2AnalyticsCompositionKeyOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2AnalyticsCompositionKeyOK) GetPayload() *models.GithubComKaytuIoKaytuEnginePkgInventoryAPIListResourceTypeCompositionResponse {
	return o.Payload
}

func (o *GetInventoryAPIV2AnalyticsCompositionKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GithubComKaytuIoKaytuEnginePkgInventoryAPIListResourceTypeCompositionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
