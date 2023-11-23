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

// PostOnboardAPIV2CredentialReader is a Reader for the PostOnboardAPIV2Credential structure.
type PostOnboardAPIV2CredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOnboardAPIV2CredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOnboardAPIV2CredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /onboard/api/v2/credential] PostOnboardAPIV2Credential", response, response.Code())
	}
}

// NewPostOnboardAPIV2CredentialOK creates a PostOnboardAPIV2CredentialOK with default headers values
func NewPostOnboardAPIV2CredentialOK() *PostOnboardAPIV2CredentialOK {
	return &PostOnboardAPIV2CredentialOK{}
}

/*
PostOnboardAPIV2CredentialOK describes a response with status code 200, with default header values.

OK
*/
type PostOnboardAPIV2CredentialOK struct {
	Payload *models.GithubComKaytuIoKaytuEnginePkgOnboardAPIV2CreateCredentialV2Response
}

// IsSuccess returns true when this post onboard Api v2 credential o k response has a 2xx status code
func (o *PostOnboardAPIV2CredentialOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post onboard Api v2 credential o k response has a 3xx status code
func (o *PostOnboardAPIV2CredentialOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post onboard Api v2 credential o k response has a 4xx status code
func (o *PostOnboardAPIV2CredentialOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post onboard Api v2 credential o k response has a 5xx status code
func (o *PostOnboardAPIV2CredentialOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post onboard Api v2 credential o k response a status code equal to that given
func (o *PostOnboardAPIV2CredentialOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post onboard Api v2 credential o k response
func (o *PostOnboardAPIV2CredentialOK) Code() int {
	return 200
}

func (o *PostOnboardAPIV2CredentialOK) Error() string {
	return fmt.Sprintf("[POST /onboard/api/v2/credential][%d] postOnboardApiV2CredentialOK  %+v", 200, o.Payload)
}

func (o *PostOnboardAPIV2CredentialOK) String() string {
	return fmt.Sprintf("[POST /onboard/api/v2/credential][%d] postOnboardApiV2CredentialOK  %+v", 200, o.Payload)
}

func (o *PostOnboardAPIV2CredentialOK) GetPayload() *models.GithubComKaytuIoKaytuEnginePkgOnboardAPIV2CreateCredentialV2Response {
	return o.Payload
}

func (o *PostOnboardAPIV2CredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GithubComKaytuIoKaytuEnginePkgOnboardAPIV2CreateCredentialV2Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
