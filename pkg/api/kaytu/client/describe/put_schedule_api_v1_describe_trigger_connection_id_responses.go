// Code generated by go-swagger; DO NOT EDIT.

package describe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutScheduleAPIV1DescribeTriggerConnectionIDReader is a Reader for the PutScheduleAPIV1DescribeTriggerConnectionID structure.
type PutScheduleAPIV1DescribeTriggerConnectionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutScheduleAPIV1DescribeTriggerConnectionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutScheduleAPIV1DescribeTriggerConnectionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /schedule/api/v1/describe/trigger/{connection_id}] PutScheduleAPIV1DescribeTriggerConnectionID", response, response.Code())
	}
}

// NewPutScheduleAPIV1DescribeTriggerConnectionIDOK creates a PutScheduleAPIV1DescribeTriggerConnectionIDOK with default headers values
func NewPutScheduleAPIV1DescribeTriggerConnectionIDOK() *PutScheduleAPIV1DescribeTriggerConnectionIDOK {
	return &PutScheduleAPIV1DescribeTriggerConnectionIDOK{}
}

/*
PutScheduleAPIV1DescribeTriggerConnectionIDOK describes a response with status code 200, with default header values.

OK
*/
type PutScheduleAPIV1DescribeTriggerConnectionIDOK struct {
}

// IsSuccess returns true when this put schedule Api v1 describe trigger connection Id o k response has a 2xx status code
func (o *PutScheduleAPIV1DescribeTriggerConnectionIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put schedule Api v1 describe trigger connection Id o k response has a 3xx status code
func (o *PutScheduleAPIV1DescribeTriggerConnectionIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put schedule Api v1 describe trigger connection Id o k response has a 4xx status code
func (o *PutScheduleAPIV1DescribeTriggerConnectionIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put schedule Api v1 describe trigger connection Id o k response has a 5xx status code
func (o *PutScheduleAPIV1DescribeTriggerConnectionIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put schedule Api v1 describe trigger connection Id o k response a status code equal to that given
func (o *PutScheduleAPIV1DescribeTriggerConnectionIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put schedule Api v1 describe trigger connection Id o k response
func (o *PutScheduleAPIV1DescribeTriggerConnectionIDOK) Code() int {
	return 200
}

func (o *PutScheduleAPIV1DescribeTriggerConnectionIDOK) Error() string {
	return fmt.Sprintf("[PUT /schedule/api/v1/describe/trigger/{connection_id}][%d] putScheduleApiV1DescribeTriggerConnectionIdOK ", 200)
}

func (o *PutScheduleAPIV1DescribeTriggerConnectionIDOK) String() string {
	return fmt.Sprintf("[PUT /schedule/api/v1/describe/trigger/{connection_id}][%d] putScheduleApiV1DescribeTriggerConnectionIdOK ", 200)
}

func (o *PutScheduleAPIV1DescribeTriggerConnectionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
