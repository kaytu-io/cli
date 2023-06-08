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

// GetInventoryAPIV2ResourcesMetricReader is a Reader for the GetInventoryAPIV2ResourcesMetric structure.
type GetInventoryAPIV2ResourcesMetricReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventoryAPIV2ResourcesMetricReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInventoryAPIV2ResourcesMetricOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInventoryAPIV2ResourcesMetricOK creates a GetInventoryAPIV2ResourcesMetricOK with default headers values
func NewGetInventoryAPIV2ResourcesMetricOK() *GetInventoryAPIV2ResourcesMetricOK {
	return &GetInventoryAPIV2ResourcesMetricOK{}
}

/*
GetInventoryAPIV2ResourcesMetricOK describes a response with status code 200, with default header values.

OK
*/
type GetInventoryAPIV2ResourcesMetricOK struct {
	Payload *models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIListResourceTypeMetricsResponse
}

// IsSuccess returns true when this get inventory Api v2 resources metric o k response has a 2xx status code
func (o *GetInventoryAPIV2ResourcesMetricOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get inventory Api v2 resources metric o k response has a 3xx status code
func (o *GetInventoryAPIV2ResourcesMetricOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get inventory Api v2 resources metric o k response has a 4xx status code
func (o *GetInventoryAPIV2ResourcesMetricOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get inventory Api v2 resources metric o k response has a 5xx status code
func (o *GetInventoryAPIV2ResourcesMetricOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get inventory Api v2 resources metric o k response a status code equal to that given
func (o *GetInventoryAPIV2ResourcesMetricOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get inventory Api v2 resources metric o k response
func (o *GetInventoryAPIV2ResourcesMetricOK) Code() int {
	return 200
}

func (o *GetInventoryAPIV2ResourcesMetricOK) Error() string {
	return fmt.Sprintf("[GET /inventory/api/v2/resources/metric][%d] getInventoryApiV2ResourcesMetricOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2ResourcesMetricOK) String() string {
	return fmt.Sprintf("[GET /inventory/api/v2/resources/metric][%d] getInventoryApiV2ResourcesMetricOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2ResourcesMetricOK) GetPayload() *models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIListResourceTypeMetricsResponse {
	return o.Payload
}

func (o *GetInventoryAPIV2ResourcesMetricOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIListResourceTypeMetricsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}