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

// GetOnboardAPIV1SourcesCountReader is a Reader for the GetOnboardAPIV1SourcesCount structure.
type GetOnboardAPIV1SourcesCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOnboardAPIV1SourcesCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOnboardAPIV1SourcesCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOnboardAPIV1SourcesCountOK creates a GetOnboardAPIV1SourcesCountOK with default headers values
func NewGetOnboardAPIV1SourcesCountOK() *GetOnboardAPIV1SourcesCountOK {
	return &GetOnboardAPIV1SourcesCountOK{}
}

/*
GetOnboardAPIV1SourcesCountOK describes a response with status code 200, with default header values.

OK
*/
type GetOnboardAPIV1SourcesCountOK struct {
	Payload int64
}

// IsSuccess returns true when this get onboard Api v1 sources count o k response has a 2xx status code
func (o *GetOnboardAPIV1SourcesCountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get onboard Api v1 sources count o k response has a 3xx status code
func (o *GetOnboardAPIV1SourcesCountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get onboard Api v1 sources count o k response has a 4xx status code
func (o *GetOnboardAPIV1SourcesCountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get onboard Api v1 sources count o k response has a 5xx status code
func (o *GetOnboardAPIV1SourcesCountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get onboard Api v1 sources count o k response a status code equal to that given
func (o *GetOnboardAPIV1SourcesCountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get onboard Api v1 sources count o k response
func (o *GetOnboardAPIV1SourcesCountOK) Code() int {
	return 200
}

func (o *GetOnboardAPIV1SourcesCountOK) Error() string {
	return fmt.Sprintf("[GET /onboard/api/v1/sources/count][%d] getOnboardApiV1SourcesCountOK  %+v", 200, o.Payload)
}

func (o *GetOnboardAPIV1SourcesCountOK) String() string {
	return fmt.Sprintf("[GET /onboard/api/v1/sources/count][%d] getOnboardApiV1SourcesCountOK  %+v", 200, o.Payload)
}

func (o *GetOnboardAPIV1SourcesCountOK) GetPayload() int64 {
	return o.Payload
}

func (o *GetOnboardAPIV1SourcesCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}