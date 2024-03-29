// Code generated by go-swagger; DO NOT EDIT.

package describe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutScheduleAPIV1ComplianceTriggerBenchmarkIDReader is a Reader for the PutScheduleAPIV1ComplianceTriggerBenchmarkID structure.
type PutScheduleAPIV1ComplianceTriggerBenchmarkIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutScheduleAPIV1ComplianceTriggerBenchmarkIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutScheduleAPIV1ComplianceTriggerBenchmarkIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /schedule/api/v1/compliance/trigger/{benchmark_id}] PutScheduleAPIV1ComplianceTriggerBenchmarkID", response, response.Code())
	}
}

// NewPutScheduleAPIV1ComplianceTriggerBenchmarkIDOK creates a PutScheduleAPIV1ComplianceTriggerBenchmarkIDOK with default headers values
func NewPutScheduleAPIV1ComplianceTriggerBenchmarkIDOK() *PutScheduleAPIV1ComplianceTriggerBenchmarkIDOK {
	return &PutScheduleAPIV1ComplianceTriggerBenchmarkIDOK{}
}

/*
PutScheduleAPIV1ComplianceTriggerBenchmarkIDOK describes a response with status code 200, with default header values.

OK
*/
type PutScheduleAPIV1ComplianceTriggerBenchmarkIDOK struct {
}

// IsSuccess returns true when this put schedule Api v1 compliance trigger benchmark Id o k response has a 2xx status code
func (o *PutScheduleAPIV1ComplianceTriggerBenchmarkIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put schedule Api v1 compliance trigger benchmark Id o k response has a 3xx status code
func (o *PutScheduleAPIV1ComplianceTriggerBenchmarkIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put schedule Api v1 compliance trigger benchmark Id o k response has a 4xx status code
func (o *PutScheduleAPIV1ComplianceTriggerBenchmarkIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put schedule Api v1 compliance trigger benchmark Id o k response has a 5xx status code
func (o *PutScheduleAPIV1ComplianceTriggerBenchmarkIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put schedule Api v1 compliance trigger benchmark Id o k response a status code equal to that given
func (o *PutScheduleAPIV1ComplianceTriggerBenchmarkIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put schedule Api v1 compliance trigger benchmark Id o k response
func (o *PutScheduleAPIV1ComplianceTriggerBenchmarkIDOK) Code() int {
	return 200
}

func (o *PutScheduleAPIV1ComplianceTriggerBenchmarkIDOK) Error() string {
	return fmt.Sprintf("[PUT /schedule/api/v1/compliance/trigger/{benchmark_id}][%d] putScheduleApiV1ComplianceTriggerBenchmarkIdOK ", 200)
}

func (o *PutScheduleAPIV1ComplianceTriggerBenchmarkIDOK) String() string {
	return fmt.Sprintf("[PUT /schedule/api/v1/compliance/trigger/{benchmark_id}][%d] putScheduleApiV1ComplianceTriggerBenchmarkIdOK ", 200)
}

func (o *PutScheduleAPIV1ComplianceTriggerBenchmarkIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
