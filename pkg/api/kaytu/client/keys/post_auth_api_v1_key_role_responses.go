// Code generated by go-swagger; DO NOT EDIT.

package keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

// PostAuthAPIV1KeyRoleReader is a Reader for the PostAuthAPIV1KeyRole structure.
type PostAuthAPIV1KeyRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAuthAPIV1KeyRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAuthAPIV1KeyRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /auth/api/v1/key/role] PostAuthAPIV1KeyRole", response, response.Code())
	}
}

// NewPostAuthAPIV1KeyRoleOK creates a PostAuthAPIV1KeyRoleOK with default headers values
func NewPostAuthAPIV1KeyRoleOK() *PostAuthAPIV1KeyRoleOK {
	return &PostAuthAPIV1KeyRoleOK{}
}

/*
PostAuthAPIV1KeyRoleOK describes a response with status code 200, with default header values.

OK
*/
type PostAuthAPIV1KeyRoleOK struct {
	Payload *models.GithubComKaytuIoKaytuEnginePkgAuthAPIWorkspaceAPIKey
}

// IsSuccess returns true when this post auth Api v1 key role o k response has a 2xx status code
func (o *PostAuthAPIV1KeyRoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post auth Api v1 key role o k response has a 3xx status code
func (o *PostAuthAPIV1KeyRoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post auth Api v1 key role o k response has a 4xx status code
func (o *PostAuthAPIV1KeyRoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post auth Api v1 key role o k response has a 5xx status code
func (o *PostAuthAPIV1KeyRoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post auth Api v1 key role o k response a status code equal to that given
func (o *PostAuthAPIV1KeyRoleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post auth Api v1 key role o k response
func (o *PostAuthAPIV1KeyRoleOK) Code() int {
	return 200
}

func (o *PostAuthAPIV1KeyRoleOK) Error() string {
	return fmt.Sprintf("[POST /auth/api/v1/key/role][%d] postAuthApiV1KeyRoleOK  %+v", 200, o.Payload)
}

func (o *PostAuthAPIV1KeyRoleOK) String() string {
	return fmt.Sprintf("[POST /auth/api/v1/key/role][%d] postAuthApiV1KeyRoleOK  %+v", 200, o.Payload)
}

func (o *PostAuthAPIV1KeyRoleOK) GetPayload() *models.GithubComKaytuIoKaytuEnginePkgAuthAPIWorkspaceAPIKey {
	return o.Payload
}

func (o *PostAuthAPIV1KeyRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GithubComKaytuIoKaytuEnginePkgAuthAPIWorkspaceAPIKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
