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

// PostOnboardAPIV1CredentialCredentialIDAutoonboardReader is a Reader for the PostOnboardAPIV1CredentialCredentialIDAutoonboard structure.
type PostOnboardAPIV1CredentialCredentialIDAutoonboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOnboardAPIV1CredentialCredentialIDAutoonboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOnboardAPIV1CredentialCredentialIDAutoonboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /onboard/api/v1/credential/{credentialId}/autoonboard] PostOnboardAPIV1CredentialCredentialIDAutoonboard", response, response.Code())
	}
}

// NewPostOnboardAPIV1CredentialCredentialIDAutoonboardOK creates a PostOnboardAPIV1CredentialCredentialIDAutoonboardOK with default headers values
func NewPostOnboardAPIV1CredentialCredentialIDAutoonboardOK() *PostOnboardAPIV1CredentialCredentialIDAutoonboardOK {
	return &PostOnboardAPIV1CredentialCredentialIDAutoonboardOK{}
}

/*
PostOnboardAPIV1CredentialCredentialIDAutoonboardOK describes a response with status code 200, with default header values.

OK
*/
type PostOnboardAPIV1CredentialCredentialIDAutoonboardOK struct {
	Payload []*models.GithubComKaytuIoKaytuEnginePkgOnboardAPIConnection
}

// IsSuccess returns true when this post onboard Api v1 credential credential Id autoonboard o k response has a 2xx status code
func (o *PostOnboardAPIV1CredentialCredentialIDAutoonboardOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post onboard Api v1 credential credential Id autoonboard o k response has a 3xx status code
func (o *PostOnboardAPIV1CredentialCredentialIDAutoonboardOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post onboard Api v1 credential credential Id autoonboard o k response has a 4xx status code
func (o *PostOnboardAPIV1CredentialCredentialIDAutoonboardOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post onboard Api v1 credential credential Id autoonboard o k response has a 5xx status code
func (o *PostOnboardAPIV1CredentialCredentialIDAutoonboardOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post onboard Api v1 credential credential Id autoonboard o k response a status code equal to that given
func (o *PostOnboardAPIV1CredentialCredentialIDAutoonboardOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post onboard Api v1 credential credential Id autoonboard o k response
func (o *PostOnboardAPIV1CredentialCredentialIDAutoonboardOK) Code() int {
	return 200
}

func (o *PostOnboardAPIV1CredentialCredentialIDAutoonboardOK) Error() string {
	return fmt.Sprintf("[POST /onboard/api/v1/credential/{credentialId}/autoonboard][%d] postOnboardApiV1CredentialCredentialIdAutoonboardOK  %+v", 200, o.Payload)
}

func (o *PostOnboardAPIV1CredentialCredentialIDAutoonboardOK) String() string {
	return fmt.Sprintf("[POST /onboard/api/v1/credential/{credentialId}/autoonboard][%d] postOnboardApiV1CredentialCredentialIdAutoonboardOK  %+v", 200, o.Payload)
}

func (o *PostOnboardAPIV1CredentialCredentialIDAutoonboardOK) GetPayload() []*models.GithubComKaytuIoKaytuEnginePkgOnboardAPIConnection {
	return o.Payload
}

func (o *PostOnboardAPIV1CredentialCredentialIDAutoonboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
