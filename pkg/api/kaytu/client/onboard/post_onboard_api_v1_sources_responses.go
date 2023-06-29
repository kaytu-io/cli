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

// PostOnboardAPIV1SourcesReader is a Reader for the PostOnboardAPIV1Sources structure.
type PostOnboardAPIV1SourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOnboardAPIV1SourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOnboardAPIV1SourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /onboard/api/v1/sources] PostOnboardAPIV1Sources", response, response.Code())
	}
}

// NewPostOnboardAPIV1SourcesOK creates a PostOnboardAPIV1SourcesOK with default headers values
func NewPostOnboardAPIV1SourcesOK() *PostOnboardAPIV1SourcesOK {
	return &PostOnboardAPIV1SourcesOK{}
}

/*
PostOnboardAPIV1SourcesOK describes a response with status code 200, with default header values.

OK
*/
type PostOnboardAPIV1SourcesOK struct {
	Payload []*models.GithubComKaytuIoKaytuEnginePkgOnboardAPIConnection
}

// IsSuccess returns true when this post onboard Api v1 sources o k response has a 2xx status code
func (o *PostOnboardAPIV1SourcesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post onboard Api v1 sources o k response has a 3xx status code
func (o *PostOnboardAPIV1SourcesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post onboard Api v1 sources o k response has a 4xx status code
func (o *PostOnboardAPIV1SourcesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post onboard Api v1 sources o k response has a 5xx status code
func (o *PostOnboardAPIV1SourcesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post onboard Api v1 sources o k response a status code equal to that given
func (o *PostOnboardAPIV1SourcesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post onboard Api v1 sources o k response
func (o *PostOnboardAPIV1SourcesOK) Code() int {
	return 200
}

func (o *PostOnboardAPIV1SourcesOK) Error() string {
	return fmt.Sprintf("[POST /onboard/api/v1/sources][%d] postOnboardApiV1SourcesOK  %+v", 200, o.Payload)
}

func (o *PostOnboardAPIV1SourcesOK) String() string {
	return fmt.Sprintf("[POST /onboard/api/v1/sources][%d] postOnboardApiV1SourcesOK  %+v", 200, o.Payload)
}

func (o *PostOnboardAPIV1SourcesOK) GetPayload() []*models.GithubComKaytuIoKaytuEnginePkgOnboardAPIConnection {
	return o.Payload
}

func (o *PostOnboardAPIV1SourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
