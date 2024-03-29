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

// PostOnboardAPIV1CredentialReader is a Reader for the PostOnboardAPIV1Credential structure.
type PostOnboardAPIV1CredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOnboardAPIV1CredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOnboardAPIV1CredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /onboard/api/v1/credential] PostOnboardAPIV1Credential", response, response.Code())
	}
}

// NewPostOnboardAPIV1CredentialOK creates a PostOnboardAPIV1CredentialOK with default headers values
func NewPostOnboardAPIV1CredentialOK() *PostOnboardAPIV1CredentialOK {
	return &PostOnboardAPIV1CredentialOK{}
}

/*
PostOnboardAPIV1CredentialOK describes a response with status code 200, with default header values.

OK
*/
type PostOnboardAPIV1CredentialOK struct {
	Payload *models.GithubComKaytuIoKaytuEnginePkgOnboardAPICreateCredentialResponse
}

// IsSuccess returns true when this post onboard Api v1 credential o k response has a 2xx status code
func (o *PostOnboardAPIV1CredentialOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post onboard Api v1 credential o k response has a 3xx status code
func (o *PostOnboardAPIV1CredentialOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post onboard Api v1 credential o k response has a 4xx status code
func (o *PostOnboardAPIV1CredentialOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post onboard Api v1 credential o k response has a 5xx status code
func (o *PostOnboardAPIV1CredentialOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post onboard Api v1 credential o k response a status code equal to that given
func (o *PostOnboardAPIV1CredentialOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post onboard Api v1 credential o k response
func (o *PostOnboardAPIV1CredentialOK) Code() int {
	return 200
}

func (o *PostOnboardAPIV1CredentialOK) Error() string {
	return fmt.Sprintf("[POST /onboard/api/v1/credential][%d] postOnboardApiV1CredentialOK  %+v", 200, o.Payload)
}

func (o *PostOnboardAPIV1CredentialOK) String() string {
	return fmt.Sprintf("[POST /onboard/api/v1/credential][%d] postOnboardApiV1CredentialOK  %+v", 200, o.Payload)
}

func (o *PostOnboardAPIV1CredentialOK) GetPayload() *models.GithubComKaytuIoKaytuEnginePkgOnboardAPICreateCredentialResponse {
	return o.Payload
}

func (o *PostOnboardAPIV1CredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GithubComKaytuIoKaytuEnginePkgOnboardAPICreateCredentialResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
