// Code generated by go-swagger; DO NOT EDIT.

package onboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutOnboardAPIV1CredentialCredentialIDReader is a Reader for the PutOnboardAPIV1CredentialCredentialID structure.
type PutOnboardAPIV1CredentialCredentialIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOnboardAPIV1CredentialCredentialIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutOnboardAPIV1CredentialCredentialIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /onboard/api/v1/credential/{credentialId}] PutOnboardAPIV1CredentialCredentialID", response, response.Code())
	}
}

// NewPutOnboardAPIV1CredentialCredentialIDOK creates a PutOnboardAPIV1CredentialCredentialIDOK with default headers values
func NewPutOnboardAPIV1CredentialCredentialIDOK() *PutOnboardAPIV1CredentialCredentialIDOK {
	return &PutOnboardAPIV1CredentialCredentialIDOK{}
}

/*
PutOnboardAPIV1CredentialCredentialIDOK describes a response with status code 200, with default header values.

OK
*/
type PutOnboardAPIV1CredentialCredentialIDOK struct {
}

// IsSuccess returns true when this put onboard Api v1 credential credential Id o k response has a 2xx status code
func (o *PutOnboardAPIV1CredentialCredentialIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put onboard Api v1 credential credential Id o k response has a 3xx status code
func (o *PutOnboardAPIV1CredentialCredentialIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put onboard Api v1 credential credential Id o k response has a 4xx status code
func (o *PutOnboardAPIV1CredentialCredentialIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put onboard Api v1 credential credential Id o k response has a 5xx status code
func (o *PutOnboardAPIV1CredentialCredentialIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put onboard Api v1 credential credential Id o k response a status code equal to that given
func (o *PutOnboardAPIV1CredentialCredentialIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put onboard Api v1 credential credential Id o k response
func (o *PutOnboardAPIV1CredentialCredentialIDOK) Code() int {
	return 200
}

func (o *PutOnboardAPIV1CredentialCredentialIDOK) Error() string {
	return fmt.Sprintf("[PUT /onboard/api/v1/credential/{credentialId}][%d] putOnboardApiV1CredentialCredentialIdOK ", 200)
}

func (o *PutOnboardAPIV1CredentialCredentialIDOK) String() string {
	return fmt.Sprintf("[PUT /onboard/api/v1/credential/{credentialId}][%d] putOnboardApiV1CredentialCredentialIdOK ", 200)
}

func (o *PutOnboardAPIV1CredentialCredentialIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
