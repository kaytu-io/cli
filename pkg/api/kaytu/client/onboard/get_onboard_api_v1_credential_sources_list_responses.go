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

// GetOnboardAPIV1CredentialSourcesListReader is a Reader for the GetOnboardAPIV1CredentialSourcesList structure.
type GetOnboardAPIV1CredentialSourcesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOnboardAPIV1CredentialSourcesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOnboardAPIV1CredentialSourcesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /onboard/api/v1/credential/sources/list] GetOnboardAPIV1CredentialSourcesList", response, response.Code())
	}
}

// NewGetOnboardAPIV1CredentialSourcesListOK creates a GetOnboardAPIV1CredentialSourcesListOK with default headers values
func NewGetOnboardAPIV1CredentialSourcesListOK() *GetOnboardAPIV1CredentialSourcesListOK {
	return &GetOnboardAPIV1CredentialSourcesListOK{}
}

/*
GetOnboardAPIV1CredentialSourcesListOK describes a response with status code 200, with default header values.

OK
*/
type GetOnboardAPIV1CredentialSourcesListOK struct {
	Payload []*models.GitlabComKeibiengineKeibiEnginePkgOnboardAPICredential
}

// IsSuccess returns true when this get onboard Api v1 credential sources list o k response has a 2xx status code
func (o *GetOnboardAPIV1CredentialSourcesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get onboard Api v1 credential sources list o k response has a 3xx status code
func (o *GetOnboardAPIV1CredentialSourcesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get onboard Api v1 credential sources list o k response has a 4xx status code
func (o *GetOnboardAPIV1CredentialSourcesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get onboard Api v1 credential sources list o k response has a 5xx status code
func (o *GetOnboardAPIV1CredentialSourcesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get onboard Api v1 credential sources list o k response a status code equal to that given
func (o *GetOnboardAPIV1CredentialSourcesListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get onboard Api v1 credential sources list o k response
func (o *GetOnboardAPIV1CredentialSourcesListOK) Code() int {
	return 200
}

func (o *GetOnboardAPIV1CredentialSourcesListOK) Error() string {
	return fmt.Sprintf("[GET /onboard/api/v1/credential/sources/list][%d] getOnboardApiV1CredentialSourcesListOK  %+v", 200, o.Payload)
}

func (o *GetOnboardAPIV1CredentialSourcesListOK) String() string {
	return fmt.Sprintf("[GET /onboard/api/v1/credential/sources/list][%d] getOnboardApiV1CredentialSourcesListOK  %+v", 200, o.Payload)
}

func (o *GetOnboardAPIV1CredentialSourcesListOK) GetPayload() []*models.GitlabComKeibiengineKeibiEnginePkgOnboardAPICredential {
	return o.Payload
}

func (o *GetOnboardAPIV1CredentialSourcesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
