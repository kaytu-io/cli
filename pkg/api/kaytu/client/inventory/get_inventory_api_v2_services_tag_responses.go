// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetInventoryAPIV2ServicesTagReader is a Reader for the GetInventoryAPIV2ServicesTag structure.
type GetInventoryAPIV2ServicesTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventoryAPIV2ServicesTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInventoryAPIV2ServicesTagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInventoryAPIV2ServicesTagOK creates a GetInventoryAPIV2ServicesTagOK with default headers values
func NewGetInventoryAPIV2ServicesTagOK() *GetInventoryAPIV2ServicesTagOK {
	return &GetInventoryAPIV2ServicesTagOK{}
}

/*
GetInventoryAPIV2ServicesTagOK describes a response with status code 200, with default header values.

OK
*/
type GetInventoryAPIV2ServicesTagOK struct {
	Payload map[string][]string
}

// IsSuccess returns true when this get inventory Api v2 services tag o k response has a 2xx status code
func (o *GetInventoryAPIV2ServicesTagOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get inventory Api v2 services tag o k response has a 3xx status code
func (o *GetInventoryAPIV2ServicesTagOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get inventory Api v2 services tag o k response has a 4xx status code
func (o *GetInventoryAPIV2ServicesTagOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get inventory Api v2 services tag o k response has a 5xx status code
func (o *GetInventoryAPIV2ServicesTagOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get inventory Api v2 services tag o k response a status code equal to that given
func (o *GetInventoryAPIV2ServicesTagOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get inventory Api v2 services tag o k response
func (o *GetInventoryAPIV2ServicesTagOK) Code() int {
	return 200
}

func (o *GetInventoryAPIV2ServicesTagOK) Error() string {
	return fmt.Sprintf("[GET /inventory/api/v2/services/tag][%d] getInventoryApiV2ServicesTagOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2ServicesTagOK) String() string {
	return fmt.Sprintf("[GET /inventory/api/v2/services/tag][%d] getInventoryApiV2ServicesTagOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2ServicesTagOK) GetPayload() map[string][]string {
	return o.Payload
}

func (o *GetInventoryAPIV2ServicesTagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
