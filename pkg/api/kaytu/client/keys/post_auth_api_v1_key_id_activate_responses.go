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

// PostAuthAPIV1KeyIDActivateReader is a Reader for the PostAuthAPIV1KeyIDActivate structure.
type PostAuthAPIV1KeyIDActivateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAuthAPIV1KeyIDActivateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAuthAPIV1KeyIDActivateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAuthAPIV1KeyIDActivateOK creates a PostAuthAPIV1KeyIDActivateOK with default headers values
func NewPostAuthAPIV1KeyIDActivateOK() *PostAuthAPIV1KeyIDActivateOK {
	return &PostAuthAPIV1KeyIDActivateOK{}
}

/*
PostAuthAPIV1KeyIDActivateOK describes a response with status code 200, with default header values.

OK
*/
type PostAuthAPIV1KeyIDActivateOK struct {
	Payload *models.GitlabComKeibiengineKeibiEnginePkgAuthAPIWorkspaceAPIKey
}

// IsSuccess returns true when this post auth Api v1 key Id activate o k response has a 2xx status code
func (o *PostAuthAPIV1KeyIDActivateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post auth Api v1 key Id activate o k response has a 3xx status code
func (o *PostAuthAPIV1KeyIDActivateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post auth Api v1 key Id activate o k response has a 4xx status code
func (o *PostAuthAPIV1KeyIDActivateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post auth Api v1 key Id activate o k response has a 5xx status code
func (o *PostAuthAPIV1KeyIDActivateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post auth Api v1 key Id activate o k response a status code equal to that given
func (o *PostAuthAPIV1KeyIDActivateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post auth Api v1 key Id activate o k response
func (o *PostAuthAPIV1KeyIDActivateOK) Code() int {
	return 200
}

func (o *PostAuthAPIV1KeyIDActivateOK) Error() string {
	return fmt.Sprintf("[POST /auth/api/v1/key/{id}/activate][%d] postAuthApiV1KeyIdActivateOK  %+v", 200, o.Payload)
}

func (o *PostAuthAPIV1KeyIDActivateOK) String() string {
	return fmt.Sprintf("[POST /auth/api/v1/key/{id}/activate][%d] postAuthApiV1KeyIdActivateOK  %+v", 200, o.Payload)
}

func (o *PostAuthAPIV1KeyIDActivateOK) GetPayload() *models.GitlabComKeibiengineKeibiEnginePkgAuthAPIWorkspaceAPIKey {
	return o.Payload
}

func (o *PostAuthAPIV1KeyIDActivateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitlabComKeibiengineKeibiEnginePkgAuthAPIWorkspaceAPIKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}