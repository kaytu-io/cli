// Code generated by go-swagger; DO NOT EDIT.

package metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

// GetInventoryAPIV2MetadataServicesServiceNameReader is a Reader for the GetInventoryAPIV2MetadataServicesServiceName structure.
type GetInventoryAPIV2MetadataServicesServiceNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventoryAPIV2MetadataServicesServiceNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInventoryAPIV2MetadataServicesServiceNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /inventory/api/v2/metadata/services/{serviceName}] GetInventoryAPIV2MetadataServicesServiceName", response, response.Code())
	}
}

// NewGetInventoryAPIV2MetadataServicesServiceNameOK creates a GetInventoryAPIV2MetadataServicesServiceNameOK with default headers values
func NewGetInventoryAPIV2MetadataServicesServiceNameOK() *GetInventoryAPIV2MetadataServicesServiceNameOK {
	return &GetInventoryAPIV2MetadataServicesServiceNameOK{}
}

/*
GetInventoryAPIV2MetadataServicesServiceNameOK describes a response with status code 200, with default header values.

OK
*/
type GetInventoryAPIV2MetadataServicesServiceNameOK struct {
	Payload *models.GithubComKaytuIoKaytuEnginePkgInventoryAPIService
}

// IsSuccess returns true when this get inventory Api v2 metadata services service name o k response has a 2xx status code
func (o *GetInventoryAPIV2MetadataServicesServiceNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get inventory Api v2 metadata services service name o k response has a 3xx status code
func (o *GetInventoryAPIV2MetadataServicesServiceNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get inventory Api v2 metadata services service name o k response has a 4xx status code
func (o *GetInventoryAPIV2MetadataServicesServiceNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get inventory Api v2 metadata services service name o k response has a 5xx status code
func (o *GetInventoryAPIV2MetadataServicesServiceNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get inventory Api v2 metadata services service name o k response a status code equal to that given
func (o *GetInventoryAPIV2MetadataServicesServiceNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get inventory Api v2 metadata services service name o k response
func (o *GetInventoryAPIV2MetadataServicesServiceNameOK) Code() int {
	return 200
}

func (o *GetInventoryAPIV2MetadataServicesServiceNameOK) Error() string {
	return fmt.Sprintf("[GET /inventory/api/v2/metadata/services/{serviceName}][%d] getInventoryApiV2MetadataServicesServiceNameOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2MetadataServicesServiceNameOK) String() string {
	return fmt.Sprintf("[GET /inventory/api/v2/metadata/services/{serviceName}][%d] getInventoryApiV2MetadataServicesServiceNameOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2MetadataServicesServiceNameOK) GetPayload() *models.GithubComKaytuIoKaytuEnginePkgInventoryAPIService {
	return o.Payload
}

func (o *GetInventoryAPIV2MetadataServicesServiceNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GithubComKaytuIoKaytuEnginePkgInventoryAPIService)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
