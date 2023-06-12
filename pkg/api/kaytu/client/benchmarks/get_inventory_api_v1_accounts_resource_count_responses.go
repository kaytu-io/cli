// Code generated by go-swagger; DO NOT EDIT.

package benchmarks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

// GetInventoryAPIV1AccountsResourceCountReader is a Reader for the GetInventoryAPIV1AccountsResourceCount structure.
type GetInventoryAPIV1AccountsResourceCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventoryAPIV1AccountsResourceCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInventoryAPIV1AccountsResourceCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /inventory/api/v1/accounts/resource/count] GetInventoryAPIV1AccountsResourceCount", response, response.Code())
	}
}

// NewGetInventoryAPIV1AccountsResourceCountOK creates a GetInventoryAPIV1AccountsResourceCountOK with default headers values
func NewGetInventoryAPIV1AccountsResourceCountOK() *GetInventoryAPIV1AccountsResourceCountOK {
	return &GetInventoryAPIV1AccountsResourceCountOK{}
}

/*
GetInventoryAPIV1AccountsResourceCountOK describes a response with status code 200, with default header values.

OK
*/
type GetInventoryAPIV1AccountsResourceCountOK struct {
	Payload []*models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIConnectionResourceCountResponse
}

// IsSuccess returns true when this get inventory Api v1 accounts resource count o k response has a 2xx status code
func (o *GetInventoryAPIV1AccountsResourceCountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get inventory Api v1 accounts resource count o k response has a 3xx status code
func (o *GetInventoryAPIV1AccountsResourceCountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get inventory Api v1 accounts resource count o k response has a 4xx status code
func (o *GetInventoryAPIV1AccountsResourceCountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get inventory Api v1 accounts resource count o k response has a 5xx status code
func (o *GetInventoryAPIV1AccountsResourceCountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get inventory Api v1 accounts resource count o k response a status code equal to that given
func (o *GetInventoryAPIV1AccountsResourceCountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get inventory Api v1 accounts resource count o k response
func (o *GetInventoryAPIV1AccountsResourceCountOK) Code() int {
	return 200
}

func (o *GetInventoryAPIV1AccountsResourceCountOK) Error() string {
	return fmt.Sprintf("[GET /inventory/api/v1/accounts/resource/count][%d] getInventoryApiV1AccountsResourceCountOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV1AccountsResourceCountOK) String() string {
	return fmt.Sprintf("[GET /inventory/api/v1/accounts/resource/count][%d] getInventoryApiV1AccountsResourceCountOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV1AccountsResourceCountOK) GetPayload() []*models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIConnectionResourceCountResponse {
	return o.Payload
}

func (o *GetInventoryAPIV1AccountsResourceCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
