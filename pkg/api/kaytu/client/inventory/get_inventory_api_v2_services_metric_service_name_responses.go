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

// GetInventoryAPIV2ServicesMetricServiceNameReader is a Reader for the GetInventoryAPIV2ServicesMetricServiceName structure.
type GetInventoryAPIV2ServicesMetricServiceNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventoryAPIV2ServicesMetricServiceNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInventoryAPIV2ServicesMetricServiceNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /inventory/api/v2/services/metric/{serviceName}] GetInventoryAPIV2ServicesMetricServiceName", response, response.Code())
	}
}

// NewGetInventoryAPIV2ServicesMetricServiceNameOK creates a GetInventoryAPIV2ServicesMetricServiceNameOK with default headers values
func NewGetInventoryAPIV2ServicesMetricServiceNameOK() *GetInventoryAPIV2ServicesMetricServiceNameOK {
	return &GetInventoryAPIV2ServicesMetricServiceNameOK{}
}

/*
GetInventoryAPIV2ServicesMetricServiceNameOK describes a response with status code 200, with default header values.

OK
*/
type GetInventoryAPIV2ServicesMetricServiceNameOK struct {
	Payload *models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIService
}

// IsSuccess returns true when this get inventory Api v2 services metric service name o k response has a 2xx status code
func (o *GetInventoryAPIV2ServicesMetricServiceNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get inventory Api v2 services metric service name o k response has a 3xx status code
func (o *GetInventoryAPIV2ServicesMetricServiceNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get inventory Api v2 services metric service name o k response has a 4xx status code
func (o *GetInventoryAPIV2ServicesMetricServiceNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get inventory Api v2 services metric service name o k response has a 5xx status code
func (o *GetInventoryAPIV2ServicesMetricServiceNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get inventory Api v2 services metric service name o k response a status code equal to that given
func (o *GetInventoryAPIV2ServicesMetricServiceNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get inventory Api v2 services metric service name o k response
func (o *GetInventoryAPIV2ServicesMetricServiceNameOK) Code() int {
	return 200
}

func (o *GetInventoryAPIV2ServicesMetricServiceNameOK) Error() string {
	return fmt.Sprintf("[GET /inventory/api/v2/services/metric/{serviceName}][%d] getInventoryApiV2ServicesMetricServiceNameOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2ServicesMetricServiceNameOK) String() string {
	return fmt.Sprintf("[GET /inventory/api/v2/services/metric/{serviceName}][%d] getInventoryApiV2ServicesMetricServiceNameOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2ServicesMetricServiceNameOK) GetPayload() *models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIService {
	return o.Payload
}

func (o *GetInventoryAPIV2ServicesMetricServiceNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIService)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
