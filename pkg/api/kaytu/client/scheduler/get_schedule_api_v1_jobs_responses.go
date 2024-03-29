// Code generated by go-swagger; DO NOT EDIT.

package scheduler

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

// GetScheduleAPIV1JobsReader is a Reader for the GetScheduleAPIV1Jobs structure.
type GetScheduleAPIV1JobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScheduleAPIV1JobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScheduleAPIV1JobsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /schedule/api/v1/jobs] GetScheduleAPIV1Jobs", response, response.Code())
	}
}

// NewGetScheduleAPIV1JobsOK creates a GetScheduleAPIV1JobsOK with default headers values
func NewGetScheduleAPIV1JobsOK() *GetScheduleAPIV1JobsOK {
	return &GetScheduleAPIV1JobsOK{}
}

/*
GetScheduleAPIV1JobsOK describes a response with status code 200, with default header values.

OK
*/
type GetScheduleAPIV1JobsOK struct {
	Payload *models.GithubComKaytuIoKaytuEnginePkgDescribeAPIListJobsResponse
}

// IsSuccess returns true when this get schedule Api v1 jobs o k response has a 2xx status code
func (o *GetScheduleAPIV1JobsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get schedule Api v1 jobs o k response has a 3xx status code
func (o *GetScheduleAPIV1JobsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get schedule Api v1 jobs o k response has a 4xx status code
func (o *GetScheduleAPIV1JobsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get schedule Api v1 jobs o k response has a 5xx status code
func (o *GetScheduleAPIV1JobsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get schedule Api v1 jobs o k response a status code equal to that given
func (o *GetScheduleAPIV1JobsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get schedule Api v1 jobs o k response
func (o *GetScheduleAPIV1JobsOK) Code() int {
	return 200
}

func (o *GetScheduleAPIV1JobsOK) Error() string {
	return fmt.Sprintf("[GET /schedule/api/v1/jobs][%d] getScheduleApiV1JobsOK  %+v", 200, o.Payload)
}

func (o *GetScheduleAPIV1JobsOK) String() string {
	return fmt.Sprintf("[GET /schedule/api/v1/jobs][%d] getScheduleApiV1JobsOK  %+v", 200, o.Payload)
}

func (o *GetScheduleAPIV1JobsOK) GetPayload() *models.GithubComKaytuIoKaytuEnginePkgDescribeAPIListJobsResponse {
	return o.Payload
}

func (o *GetScheduleAPIV1JobsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GithubComKaytuIoKaytuEnginePkgDescribeAPIListJobsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
