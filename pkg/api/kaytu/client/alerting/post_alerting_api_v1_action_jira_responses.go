// Code generated by go-swagger; DO NOT EDIT.

package alerting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

// PostAlertingAPIV1ActionJiraReader is a Reader for the PostAlertingAPIV1ActionJira structure.
type PostAlertingAPIV1ActionJiraReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAlertingAPIV1ActionJiraReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAlertingAPIV1ActionJiraOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /alerting/api/v1/action/jira] PostAlertingAPIV1ActionJira", response, response.Code())
	}
}

// NewPostAlertingAPIV1ActionJiraOK creates a PostAlertingAPIV1ActionJiraOK with default headers values
func NewPostAlertingAPIV1ActionJiraOK() *PostAlertingAPIV1ActionJiraOK {
	return &PostAlertingAPIV1ActionJiraOK{}
}

/*
PostAlertingAPIV1ActionJiraOK describes a response with status code 200, with default header values.

OK
*/
type PostAlertingAPIV1ActionJiraOK struct {
	Payload *models.GithubComKaytuIoKaytuEnginePkgAlertingAPIJiraAndStackResponse
}

// IsSuccess returns true when this post alerting Api v1 action jira o k response has a 2xx status code
func (o *PostAlertingAPIV1ActionJiraOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post alerting Api v1 action jira o k response has a 3xx status code
func (o *PostAlertingAPIV1ActionJiraOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post alerting Api v1 action jira o k response has a 4xx status code
func (o *PostAlertingAPIV1ActionJiraOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post alerting Api v1 action jira o k response has a 5xx status code
func (o *PostAlertingAPIV1ActionJiraOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post alerting Api v1 action jira o k response a status code equal to that given
func (o *PostAlertingAPIV1ActionJiraOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post alerting Api v1 action jira o k response
func (o *PostAlertingAPIV1ActionJiraOK) Code() int {
	return 200
}

func (o *PostAlertingAPIV1ActionJiraOK) Error() string {
	return fmt.Sprintf("[POST /alerting/api/v1/action/jira][%d] postAlertingApiV1ActionJiraOK  %+v", 200, o.Payload)
}

func (o *PostAlertingAPIV1ActionJiraOK) String() string {
	return fmt.Sprintf("[POST /alerting/api/v1/action/jira][%d] postAlertingApiV1ActionJiraOK  %+v", 200, o.Payload)
}

func (o *PostAlertingAPIV1ActionJiraOK) GetPayload() *models.GithubComKaytuIoKaytuEnginePkgAlertingAPIJiraAndStackResponse {
	return o.Payload
}

func (o *PostAlertingAPIV1ActionJiraOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GithubComKaytuIoKaytuEnginePkgAlertingAPIJiraAndStackResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}