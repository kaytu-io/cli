// Code generated by go-swagger; DO NOT EDIT.

package resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

// GetInventoryAPIV2ResourcesRegionsCompositionReader is a Reader for the GetInventoryAPIV2ResourcesRegionsComposition structure.
type GetInventoryAPIV2ResourcesRegionsCompositionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventoryAPIV2ResourcesRegionsCompositionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInventoryAPIV2ResourcesRegionsCompositionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /inventory/api/v2/resources/regions/composition] GetInventoryAPIV2ResourcesRegionsComposition", response, response.Code())
	}
}

// NewGetInventoryAPIV2ResourcesRegionsCompositionOK creates a GetInventoryAPIV2ResourcesRegionsCompositionOK with default headers values
func NewGetInventoryAPIV2ResourcesRegionsCompositionOK() *GetInventoryAPIV2ResourcesRegionsCompositionOK {
	return &GetInventoryAPIV2ResourcesRegionsCompositionOK{}
}

/*
GetInventoryAPIV2ResourcesRegionsCompositionOK describes a response with status code 200, with default header values.

OK
*/
type GetInventoryAPIV2ResourcesRegionsCompositionOK struct {
	Payload *models.GithubComKaytuIoKaytuEnginePkgInventoryAPIListRegionsResourceCountCompositionResponse
}

// IsSuccess returns true when this get inventory Api v2 resources regions composition o k response has a 2xx status code
func (o *GetInventoryAPIV2ResourcesRegionsCompositionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get inventory Api v2 resources regions composition o k response has a 3xx status code
func (o *GetInventoryAPIV2ResourcesRegionsCompositionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get inventory Api v2 resources regions composition o k response has a 4xx status code
func (o *GetInventoryAPIV2ResourcesRegionsCompositionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get inventory Api v2 resources regions composition o k response has a 5xx status code
func (o *GetInventoryAPIV2ResourcesRegionsCompositionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get inventory Api v2 resources regions composition o k response a status code equal to that given
func (o *GetInventoryAPIV2ResourcesRegionsCompositionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get inventory Api v2 resources regions composition o k response
func (o *GetInventoryAPIV2ResourcesRegionsCompositionOK) Code() int {
	return 200
}

func (o *GetInventoryAPIV2ResourcesRegionsCompositionOK) Error() string {
	return fmt.Sprintf("[GET /inventory/api/v2/resources/regions/composition][%d] getInventoryApiV2ResourcesRegionsCompositionOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2ResourcesRegionsCompositionOK) String() string {
	return fmt.Sprintf("[GET /inventory/api/v2/resources/regions/composition][%d] getInventoryApiV2ResourcesRegionsCompositionOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2ResourcesRegionsCompositionOK) GetPayload() *models.GithubComKaytuIoKaytuEnginePkgInventoryAPIListRegionsResourceCountCompositionResponse {
	return o.Payload
}

func (o *GetInventoryAPIV2ResourcesRegionsCompositionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GithubComKaytuIoKaytuEnginePkgInventoryAPIListRegionsResourceCountCompositionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}