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

// GetScheduleAPIV1StacksStackIDInsightReader is a Reader for the GetScheduleAPIV1StacksStackIDInsight structure.
type GetScheduleAPIV1StacksStackIDInsightReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScheduleAPIV1StacksStackIDInsightReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScheduleAPIV1StacksStackIDInsightOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /schedule/api/v1/stacks/{stackId}/insight] GetScheduleAPIV1StacksStackIDInsight", response, response.Code())
	}
}

// NewGetScheduleAPIV1StacksStackIDInsightOK creates a GetScheduleAPIV1StacksStackIDInsightOK with default headers values
func NewGetScheduleAPIV1StacksStackIDInsightOK() *GetScheduleAPIV1StacksStackIDInsightOK {
	return &GetScheduleAPIV1StacksStackIDInsightOK{}
}

/*
GetScheduleAPIV1StacksStackIDInsightOK describes a response with status code 200, with default header values.

OK
*/
type GetScheduleAPIV1StacksStackIDInsightOK struct {
	Payload *models.GithubComKaytuIoKaytuEnginePkgComplianceAPIInsight
}

// IsSuccess returns true when this get schedule Api v1 stacks stack Id insight o k response has a 2xx status code
func (o *GetScheduleAPIV1StacksStackIDInsightOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get schedule Api v1 stacks stack Id insight o k response has a 3xx status code
func (o *GetScheduleAPIV1StacksStackIDInsightOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get schedule Api v1 stacks stack Id insight o k response has a 4xx status code
func (o *GetScheduleAPIV1StacksStackIDInsightOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get schedule Api v1 stacks stack Id insight o k response has a 5xx status code
func (o *GetScheduleAPIV1StacksStackIDInsightOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get schedule Api v1 stacks stack Id insight o k response a status code equal to that given
func (o *GetScheduleAPIV1StacksStackIDInsightOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get schedule Api v1 stacks stack Id insight o k response
func (o *GetScheduleAPIV1StacksStackIDInsightOK) Code() int {
	return 200
}

func (o *GetScheduleAPIV1StacksStackIDInsightOK) Error() string {
	return fmt.Sprintf("[GET /schedule/api/v1/stacks/{stackId}/insight][%d] getScheduleApiV1StacksStackIdInsightOK  %+v", 200, o.Payload)
}

func (o *GetScheduleAPIV1StacksStackIDInsightOK) String() string {
	return fmt.Sprintf("[GET /schedule/api/v1/stacks/{stackId}/insight][%d] getScheduleApiV1StacksStackIdInsightOK  %+v", 200, o.Payload)
}

func (o *GetScheduleAPIV1StacksStackIDInsightOK) GetPayload() *models.GithubComKaytuIoKaytuEnginePkgComplianceAPIInsight {
	return o.Payload
}

func (o *GetScheduleAPIV1StacksStackIDInsightOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GithubComKaytuIoKaytuEnginePkgComplianceAPIInsight)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
