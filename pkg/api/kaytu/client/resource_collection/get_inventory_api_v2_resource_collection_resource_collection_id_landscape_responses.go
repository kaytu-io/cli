// Code generated by go-swagger; DO NOT EDIT.

package resource_collection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

// GetInventoryAPIV2ResourceCollectionResourceCollectionIDLandscapeReader is a Reader for the GetInventoryAPIV2ResourceCollectionResourceCollectionIDLandscape structure.
type GetInventoryAPIV2ResourceCollectionResourceCollectionIDLandscapeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventoryAPIV2ResourceCollectionResourceCollectionIDLandscapeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInventoryAPIV2ResourceCollectionResourceCollectionIDLandscapeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /inventory/api/v2/resource-collection/{resourceCollectionId}/landscape] GetInventoryAPIV2ResourceCollectionResourceCollectionIDLandscape", response, response.Code())
	}
}

// NewGetInventoryAPIV2ResourceCollectionResourceCollectionIDLandscapeOK creates a GetInventoryAPIV2ResourceCollectionResourceCollectionIDLandscapeOK with default headers values
func NewGetInventoryAPIV2ResourceCollectionResourceCollectionIDLandscapeOK() *GetInventoryAPIV2ResourceCollectionResourceCollectionIDLandscapeOK {
	return &GetInventoryAPIV2ResourceCollectionResourceCollectionIDLandscapeOK{}
}

/*
GetInventoryAPIV2ResourceCollectionResourceCollectionIDLandscapeOK describes a response with status code 200, with default header values.

OK
*/
type GetInventoryAPIV2ResourceCollectionResourceCollectionIDLandscapeOK struct {
	Payload *models.GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceCollectionLandscape
}

// IsSuccess returns true when this get inventory Api v2 resource collection resource collection Id landscape o k response has a 2xx status code
func (o *GetInventoryAPIV2ResourceCollectionResourceCollectionIDLandscapeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get inventory Api v2 resource collection resource collection Id landscape o k response has a 3xx status code
func (o *GetInventoryAPIV2ResourceCollectionResourceCollectionIDLandscapeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get inventory Api v2 resource collection resource collection Id landscape o k response has a 4xx status code
func (o *GetInventoryAPIV2ResourceCollectionResourceCollectionIDLandscapeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get inventory Api v2 resource collection resource collection Id landscape o k response has a 5xx status code
func (o *GetInventoryAPIV2ResourceCollectionResourceCollectionIDLandscapeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get inventory Api v2 resource collection resource collection Id landscape o k response a status code equal to that given
func (o *GetInventoryAPIV2ResourceCollectionResourceCollectionIDLandscapeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get inventory Api v2 resource collection resource collection Id landscape o k response
func (o *GetInventoryAPIV2ResourceCollectionResourceCollectionIDLandscapeOK) Code() int {
	return 200
}

func (o *GetInventoryAPIV2ResourceCollectionResourceCollectionIDLandscapeOK) Error() string {
	return fmt.Sprintf("[GET /inventory/api/v2/resource-collection/{resourceCollectionId}/landscape][%d] getInventoryApiV2ResourceCollectionResourceCollectionIdLandscapeOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2ResourceCollectionResourceCollectionIDLandscapeOK) String() string {
	return fmt.Sprintf("[GET /inventory/api/v2/resource-collection/{resourceCollectionId}/landscape][%d] getInventoryApiV2ResourceCollectionResourceCollectionIdLandscapeOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2ResourceCollectionResourceCollectionIDLandscapeOK) GetPayload() *models.GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceCollectionLandscape {
	return o.Payload
}

func (o *GetInventoryAPIV2ResourceCollectionResourceCollectionIDLandscapeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceCollectionLandscape)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}