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

// GetInventoryAPIV2MetadataServicesReader is a Reader for the GetInventoryAPIV2MetadataServices structure.
type GetInventoryAPIV2MetadataServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventoryAPIV2MetadataServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInventoryAPIV2MetadataServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInventoryAPIV2MetadataServicesOK creates a GetInventoryAPIV2MetadataServicesOK with default headers values
func NewGetInventoryAPIV2MetadataServicesOK() *GetInventoryAPIV2MetadataServicesOK {
	return &GetInventoryAPIV2MetadataServicesOK{}
}

/*
GetInventoryAPIV2MetadataServicesOK describes a response with status code 200, with default header values.

OK
*/
type GetInventoryAPIV2MetadataServicesOK struct {
	Payload *models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIListServiceMetadataResponse
}

// IsSuccess returns true when this get inventory Api v2 metadata services o k response has a 2xx status code
func (o *GetInventoryAPIV2MetadataServicesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get inventory Api v2 metadata services o k response has a 3xx status code
func (o *GetInventoryAPIV2MetadataServicesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get inventory Api v2 metadata services o k response has a 4xx status code
func (o *GetInventoryAPIV2MetadataServicesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get inventory Api v2 metadata services o k response has a 5xx status code
func (o *GetInventoryAPIV2MetadataServicesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get inventory Api v2 metadata services o k response a status code equal to that given
func (o *GetInventoryAPIV2MetadataServicesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get inventory Api v2 metadata services o k response
func (o *GetInventoryAPIV2MetadataServicesOK) Code() int {
	return 200
}

func (o *GetInventoryAPIV2MetadataServicesOK) Error() string {
	return fmt.Sprintf("[GET /inventory/api/v2/metadata/services][%d] getInventoryApiV2MetadataServicesOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2MetadataServicesOK) String() string {
	return fmt.Sprintf("[GET /inventory/api/v2/metadata/services][%d] getInventoryApiV2MetadataServicesOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2MetadataServicesOK) GetPayload() *models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIListServiceMetadataResponse {
	return o.Payload
}

func (o *GetInventoryAPIV2MetadataServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIListServiceMetadataResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}