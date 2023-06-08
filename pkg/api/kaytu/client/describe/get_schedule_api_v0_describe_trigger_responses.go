// Code generated by go-swagger; DO NOT EDIT.

package describe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetScheduleAPIV0DescribeTriggerReader is a Reader for the GetScheduleAPIV0DescribeTrigger structure.
type GetScheduleAPIV0DescribeTriggerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScheduleAPIV0DescribeTriggerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScheduleAPIV0DescribeTriggerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScheduleAPIV0DescribeTriggerOK creates a GetScheduleAPIV0DescribeTriggerOK with default headers values
func NewGetScheduleAPIV0DescribeTriggerOK() *GetScheduleAPIV0DescribeTriggerOK {
	return &GetScheduleAPIV0DescribeTriggerOK{}
}

/*
GetScheduleAPIV0DescribeTriggerOK describes a response with status code 200, with default header values.

OK
*/
type GetScheduleAPIV0DescribeTriggerOK struct {
}

// IsSuccess returns true when this get schedule Api v0 describe trigger o k response has a 2xx status code
func (o *GetScheduleAPIV0DescribeTriggerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get schedule Api v0 describe trigger o k response has a 3xx status code
func (o *GetScheduleAPIV0DescribeTriggerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get schedule Api v0 describe trigger o k response has a 4xx status code
func (o *GetScheduleAPIV0DescribeTriggerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get schedule Api v0 describe trigger o k response has a 5xx status code
func (o *GetScheduleAPIV0DescribeTriggerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get schedule Api v0 describe trigger o k response a status code equal to that given
func (o *GetScheduleAPIV0DescribeTriggerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get schedule Api v0 describe trigger o k response
func (o *GetScheduleAPIV0DescribeTriggerOK) Code() int {
	return 200
}

func (o *GetScheduleAPIV0DescribeTriggerOK) Error() string {
	return fmt.Sprintf("[GET /schedule/api/v0/describe/trigger][%d] getScheduleApiV0DescribeTriggerOK ", 200)
}

func (o *GetScheduleAPIV0DescribeTriggerOK) String() string {
	return fmt.Sprintf("[GET /schedule/api/v0/describe/trigger][%d] getScheduleApiV0DescribeTriggerOK ", 200)
}

func (o *GetScheduleAPIV0DescribeTriggerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
