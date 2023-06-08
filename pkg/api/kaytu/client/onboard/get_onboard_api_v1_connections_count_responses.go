// Code generated by go-swagger; DO NOT EDIT.

package onboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetOnboardAPIV1ConnectionsCountReader is a Reader for the GetOnboardAPIV1ConnectionsCount structure.
type GetOnboardAPIV1ConnectionsCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOnboardAPIV1ConnectionsCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOnboardAPIV1ConnectionsCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOnboardAPIV1ConnectionsCountOK creates a GetOnboardAPIV1ConnectionsCountOK with default headers values
func NewGetOnboardAPIV1ConnectionsCountOK() *GetOnboardAPIV1ConnectionsCountOK {
	return &GetOnboardAPIV1ConnectionsCountOK{}
}

/*
GetOnboardAPIV1ConnectionsCountOK describes a response with status code 200, with default header values.

OK
*/
type GetOnboardAPIV1ConnectionsCountOK struct {
	Payload int64
}

// IsSuccess returns true when this get onboard Api v1 connections count o k response has a 2xx status code
func (o *GetOnboardAPIV1ConnectionsCountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get onboard Api v1 connections count o k response has a 3xx status code
func (o *GetOnboardAPIV1ConnectionsCountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get onboard Api v1 connections count o k response has a 4xx status code
func (o *GetOnboardAPIV1ConnectionsCountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get onboard Api v1 connections count o k response has a 5xx status code
func (o *GetOnboardAPIV1ConnectionsCountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get onboard Api v1 connections count o k response a status code equal to that given
func (o *GetOnboardAPIV1ConnectionsCountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get onboard Api v1 connections count o k response
func (o *GetOnboardAPIV1ConnectionsCountOK) Code() int {
	return 200
}

func (o *GetOnboardAPIV1ConnectionsCountOK) Error() string {
	return fmt.Sprintf("[GET /onboard/api/v1/connections/count][%d] getOnboardApiV1ConnectionsCountOK  %+v", 200, o.Payload)
}

func (o *GetOnboardAPIV1ConnectionsCountOK) String() string {
	return fmt.Sprintf("[GET /onboard/api/v1/connections/count][%d] getOnboardApiV1ConnectionsCountOK  %+v", 200, o.Payload)
}

func (o *GetOnboardAPIV1ConnectionsCountOK) GetPayload() int64 {
	return o.Payload
}

func (o *GetOnboardAPIV1ConnectionsCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}