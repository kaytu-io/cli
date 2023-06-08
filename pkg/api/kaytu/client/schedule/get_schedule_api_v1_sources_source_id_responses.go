// Code generated by go-swagger; DO NOT EDIT.

package schedule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

// GetScheduleAPIV1SourcesSourceIDReader is a Reader for the GetScheduleAPIV1SourcesSourceID structure.
type GetScheduleAPIV1SourcesSourceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScheduleAPIV1SourcesSourceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScheduleAPIV1SourcesSourceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScheduleAPIV1SourcesSourceIDOK creates a GetScheduleAPIV1SourcesSourceIDOK with default headers values
func NewGetScheduleAPIV1SourcesSourceIDOK() *GetScheduleAPIV1SourcesSourceIDOK {
	return &GetScheduleAPIV1SourcesSourceIDOK{}
}

/*
GetScheduleAPIV1SourcesSourceIDOK describes a response with status code 200, with default header values.

OK
*/
type GetScheduleAPIV1SourcesSourceIDOK struct {
	Payload *models.GitlabComKeibiengineKeibiEnginePkgDescribeAPISource
}

// IsSuccess returns true when this get schedule Api v1 sources source Id o k response has a 2xx status code
func (o *GetScheduleAPIV1SourcesSourceIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get schedule Api v1 sources source Id o k response has a 3xx status code
func (o *GetScheduleAPIV1SourcesSourceIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get schedule Api v1 sources source Id o k response has a 4xx status code
func (o *GetScheduleAPIV1SourcesSourceIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get schedule Api v1 sources source Id o k response has a 5xx status code
func (o *GetScheduleAPIV1SourcesSourceIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get schedule Api v1 sources source Id o k response a status code equal to that given
func (o *GetScheduleAPIV1SourcesSourceIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get schedule Api v1 sources source Id o k response
func (o *GetScheduleAPIV1SourcesSourceIDOK) Code() int {
	return 200
}

func (o *GetScheduleAPIV1SourcesSourceIDOK) Error() string {
	return fmt.Sprintf("[GET /schedule/api/v1/sources/{source_id}][%d] getScheduleApiV1SourcesSourceIdOK  %+v", 200, o.Payload)
}

func (o *GetScheduleAPIV1SourcesSourceIDOK) String() string {
	return fmt.Sprintf("[GET /schedule/api/v1/sources/{source_id}][%d] getScheduleApiV1SourcesSourceIdOK  %+v", 200, o.Payload)
}

func (o *GetScheduleAPIV1SourcesSourceIDOK) GetPayload() *models.GitlabComKeibiengineKeibiEnginePkgDescribeAPISource {
	return o.Payload
}

func (o *GetScheduleAPIV1SourcesSourceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitlabComKeibiengineKeibiEnginePkgDescribeAPISource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}