// Code generated by go-swagger; DO NOT EDIT.

package onboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

// PostOnboardAPIV1ConnectionsAwsReader is a Reader for the PostOnboardAPIV1ConnectionsAws structure.
type PostOnboardAPIV1ConnectionsAwsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOnboardAPIV1ConnectionsAwsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOnboardAPIV1ConnectionsAwsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /onboard/api/v1/connections/aws] PostOnboardAPIV1ConnectionsAws", response, response.Code())
	}
}

// NewPostOnboardAPIV1ConnectionsAwsOK creates a PostOnboardAPIV1ConnectionsAwsOK with default headers values
func NewPostOnboardAPIV1ConnectionsAwsOK() *PostOnboardAPIV1ConnectionsAwsOK {
	return &PostOnboardAPIV1ConnectionsAwsOK{}
}

/*
PostOnboardAPIV1ConnectionsAwsOK describes a response with status code 200, with default header values.

OK
*/
type PostOnboardAPIV1ConnectionsAwsOK struct {
	Payload *models.GithubComKaytuIoKaytuEnginePkgOnboardAPICreateConnectionResponse
}

// IsSuccess returns true when this post onboard Api v1 connections aws o k response has a 2xx status code
func (o *PostOnboardAPIV1ConnectionsAwsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post onboard Api v1 connections aws o k response has a 3xx status code
func (o *PostOnboardAPIV1ConnectionsAwsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post onboard Api v1 connections aws o k response has a 4xx status code
func (o *PostOnboardAPIV1ConnectionsAwsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post onboard Api v1 connections aws o k response has a 5xx status code
func (o *PostOnboardAPIV1ConnectionsAwsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post onboard Api v1 connections aws o k response a status code equal to that given
func (o *PostOnboardAPIV1ConnectionsAwsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post onboard Api v1 connections aws o k response
func (o *PostOnboardAPIV1ConnectionsAwsOK) Code() int {
	return 200
}

func (o *PostOnboardAPIV1ConnectionsAwsOK) Error() string {
	return fmt.Sprintf("[POST /onboard/api/v1/connections/aws][%d] postOnboardApiV1ConnectionsAwsOK  %+v", 200, o.Payload)
}

func (o *PostOnboardAPIV1ConnectionsAwsOK) String() string {
	return fmt.Sprintf("[POST /onboard/api/v1/connections/aws][%d] postOnboardApiV1ConnectionsAwsOK  %+v", 200, o.Payload)
}

func (o *PostOnboardAPIV1ConnectionsAwsOK) GetPayload() *models.GithubComKaytuIoKaytuEnginePkgOnboardAPICreateConnectionResponse {
	return o.Payload
}

func (o *PostOnboardAPIV1ConnectionsAwsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GithubComKaytuIoKaytuEnginePkgOnboardAPICreateConnectionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
