// Code generated by go-swagger; DO NOT EDIT.

package stack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

// PostScheduleAPIV1StacksBenchmarkTriggerReader is a Reader for the PostScheduleAPIV1StacksBenchmarkTrigger structure.
type PostScheduleAPIV1StacksBenchmarkTriggerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostScheduleAPIV1StacksBenchmarkTriggerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostScheduleAPIV1StacksBenchmarkTriggerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /schedule/api/v1/stacks/benchmark/trigger] PostScheduleAPIV1StacksBenchmarkTrigger", response, response.Code())
	}
}

// NewPostScheduleAPIV1StacksBenchmarkTriggerOK creates a PostScheduleAPIV1StacksBenchmarkTriggerOK with default headers values
func NewPostScheduleAPIV1StacksBenchmarkTriggerOK() *PostScheduleAPIV1StacksBenchmarkTriggerOK {
	return &PostScheduleAPIV1StacksBenchmarkTriggerOK{}
}

/*
PostScheduleAPIV1StacksBenchmarkTriggerOK describes a response with status code 200, with default header values.

OK
*/
type PostScheduleAPIV1StacksBenchmarkTriggerOK struct {
	Payload []*models.DescribeComplianceReportJob
}

// IsSuccess returns true when this post schedule Api v1 stacks benchmark trigger o k response has a 2xx status code
func (o *PostScheduleAPIV1StacksBenchmarkTriggerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post schedule Api v1 stacks benchmark trigger o k response has a 3xx status code
func (o *PostScheduleAPIV1StacksBenchmarkTriggerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post schedule Api v1 stacks benchmark trigger o k response has a 4xx status code
func (o *PostScheduleAPIV1StacksBenchmarkTriggerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post schedule Api v1 stacks benchmark trigger o k response has a 5xx status code
func (o *PostScheduleAPIV1StacksBenchmarkTriggerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post schedule Api v1 stacks benchmark trigger o k response a status code equal to that given
func (o *PostScheduleAPIV1StacksBenchmarkTriggerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post schedule Api v1 stacks benchmark trigger o k response
func (o *PostScheduleAPIV1StacksBenchmarkTriggerOK) Code() int {
	return 200
}

func (o *PostScheduleAPIV1StacksBenchmarkTriggerOK) Error() string {
	return fmt.Sprintf("[POST /schedule/api/v1/stacks/benchmark/trigger][%d] postScheduleApiV1StacksBenchmarkTriggerOK  %+v", 200, o.Payload)
}

func (o *PostScheduleAPIV1StacksBenchmarkTriggerOK) String() string {
	return fmt.Sprintf("[POST /schedule/api/v1/stacks/benchmark/trigger][%d] postScheduleApiV1StacksBenchmarkTriggerOK  %+v", 200, o.Payload)
}

func (o *PostScheduleAPIV1StacksBenchmarkTriggerOK) GetPayload() []*models.DescribeComplianceReportJob {
	return o.Payload
}

func (o *PostScheduleAPIV1StacksBenchmarkTriggerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
