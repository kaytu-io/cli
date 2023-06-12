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

// PostAuthAPIV1KeyIDSuspendReader is a Reader for the PostAuthAPIV1KeyIDSuspend structure.
type PostAuthAPIV1KeyIDSuspendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAuthAPIV1KeyIDSuspendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAuthAPIV1KeyIDSuspendOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /auth/api/v1/key/{id}/suspend] PostAuthAPIV1KeyIDSuspend", response, response.Code())
	}
}

// NewPostAuthAPIV1KeyIDSuspendOK creates a PostAuthAPIV1KeyIDSuspendOK with default headers values
func NewPostAuthAPIV1KeyIDSuspendOK() *PostAuthAPIV1KeyIDSuspendOK {
	return &PostAuthAPIV1KeyIDSuspendOK{}
}

/*
PostAuthAPIV1KeyIDSuspendOK describes a response with status code 200, with default header values.

OK
*/
type PostAuthAPIV1KeyIDSuspendOK struct {
	Payload *models.GitlabComKeibiengineKeibiEnginePkgAuthAPIWorkspaceAPIKey
}

// IsSuccess returns true when this post auth Api v1 key Id suspend o k response has a 2xx status code
func (o *PostAuthAPIV1KeyIDSuspendOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post auth Api v1 key Id suspend o k response has a 3xx status code
func (o *PostAuthAPIV1KeyIDSuspendOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post auth Api v1 key Id suspend o k response has a 4xx status code
func (o *PostAuthAPIV1KeyIDSuspendOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post auth Api v1 key Id suspend o k response has a 5xx status code
func (o *PostAuthAPIV1KeyIDSuspendOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post auth Api v1 key Id suspend o k response a status code equal to that given
func (o *PostAuthAPIV1KeyIDSuspendOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post auth Api v1 key Id suspend o k response
func (o *PostAuthAPIV1KeyIDSuspendOK) Code() int {
	return 200
}

func (o *PostAuthAPIV1KeyIDSuspendOK) Error() string {
	return fmt.Sprintf("[POST /auth/api/v1/key/{id}/suspend][%d] postAuthApiV1KeyIdSuspendOK  %+v", 200, o.Payload)
}

func (o *PostAuthAPIV1KeyIDSuspendOK) String() string {
	return fmt.Sprintf("[POST /auth/api/v1/key/{id}/suspend][%d] postAuthApiV1KeyIdSuspendOK  %+v", 200, o.Payload)
}

func (o *PostAuthAPIV1KeyIDSuspendOK) GetPayload() *models.GitlabComKeibiengineKeibiEnginePkgAuthAPIWorkspaceAPIKey {
	return o.Payload
}

func (o *PostAuthAPIV1KeyIDSuspendOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitlabComKeibiengineKeibiEnginePkgAuthAPIWorkspaceAPIKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
