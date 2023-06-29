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

// GetInventoryAPIV2ResourcesRegionsSummaryReader is a Reader for the GetInventoryAPIV2ResourcesRegionsSummary structure.
type GetInventoryAPIV2ResourcesRegionsSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventoryAPIV2ResourcesRegionsSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInventoryAPIV2ResourcesRegionsSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /inventory/api/v2/resources/regions/summary] GetInventoryAPIV2ResourcesRegionsSummary", response, response.Code())
	}
}

// NewGetInventoryAPIV2ResourcesRegionsSummaryOK creates a GetInventoryAPIV2ResourcesRegionsSummaryOK with default headers values
func NewGetInventoryAPIV2ResourcesRegionsSummaryOK() *GetInventoryAPIV2ResourcesRegionsSummaryOK {
	return &GetInventoryAPIV2ResourcesRegionsSummaryOK{}
}

/*
GetInventoryAPIV2ResourcesRegionsSummaryOK describes a response with status code 200, with default header values.

OK
*/
type GetInventoryAPIV2ResourcesRegionsSummaryOK struct {
	Payload *models.GithubComKaytuIoKaytuEnginePkgInventoryAPIRegionsResourceCountResponse
}

// IsSuccess returns true when this get inventory Api v2 resources regions summary o k response has a 2xx status code
func (o *GetInventoryAPIV2ResourcesRegionsSummaryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get inventory Api v2 resources regions summary o k response has a 3xx status code
func (o *GetInventoryAPIV2ResourcesRegionsSummaryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get inventory Api v2 resources regions summary o k response has a 4xx status code
func (o *GetInventoryAPIV2ResourcesRegionsSummaryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get inventory Api v2 resources regions summary o k response has a 5xx status code
func (o *GetInventoryAPIV2ResourcesRegionsSummaryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get inventory Api v2 resources regions summary o k response a status code equal to that given
func (o *GetInventoryAPIV2ResourcesRegionsSummaryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get inventory Api v2 resources regions summary o k response
func (o *GetInventoryAPIV2ResourcesRegionsSummaryOK) Code() int {
	return 200
}

func (o *GetInventoryAPIV2ResourcesRegionsSummaryOK) Error() string {
	return fmt.Sprintf("[GET /inventory/api/v2/resources/regions/summary][%d] getInventoryApiV2ResourcesRegionsSummaryOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2ResourcesRegionsSummaryOK) String() string {
	return fmt.Sprintf("[GET /inventory/api/v2/resources/regions/summary][%d] getInventoryApiV2ResourcesRegionsSummaryOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2ResourcesRegionsSummaryOK) GetPayload() *models.GithubComKaytuIoKaytuEnginePkgInventoryAPIRegionsResourceCountResponse {
	return o.Payload
}

func (o *GetInventoryAPIV2ResourcesRegionsSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GithubComKaytuIoKaytuEnginePkgInventoryAPIRegionsResourceCountResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
