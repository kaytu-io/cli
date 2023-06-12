// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostAuthAPIV1UserInviteReader is a Reader for the PostAuthAPIV1UserInvite structure.
type PostAuthAPIV1UserInviteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAuthAPIV1UserInviteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAuthAPIV1UserInviteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /auth/api/v1/user/invite] PostAuthAPIV1UserInvite", response, response.Code())
	}
}

// NewPostAuthAPIV1UserInviteOK creates a PostAuthAPIV1UserInviteOK with default headers values
func NewPostAuthAPIV1UserInviteOK() *PostAuthAPIV1UserInviteOK {
	return &PostAuthAPIV1UserInviteOK{}
}

/*
PostAuthAPIV1UserInviteOK describes a response with status code 200, with default header values.

OK
*/
type PostAuthAPIV1UserInviteOK struct {
}

// IsSuccess returns true when this post auth Api v1 user invite o k response has a 2xx status code
func (o *PostAuthAPIV1UserInviteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post auth Api v1 user invite o k response has a 3xx status code
func (o *PostAuthAPIV1UserInviteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post auth Api v1 user invite o k response has a 4xx status code
func (o *PostAuthAPIV1UserInviteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post auth Api v1 user invite o k response has a 5xx status code
func (o *PostAuthAPIV1UserInviteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post auth Api v1 user invite o k response a status code equal to that given
func (o *PostAuthAPIV1UserInviteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post auth Api v1 user invite o k response
func (o *PostAuthAPIV1UserInviteOK) Code() int {
	return 200
}

func (o *PostAuthAPIV1UserInviteOK) Error() string {
	return fmt.Sprintf("[POST /auth/api/v1/user/invite][%d] postAuthApiV1UserInviteOK ", 200)
}

func (o *PostAuthAPIV1UserInviteOK) String() string {
	return fmt.Sprintf("[POST /auth/api/v1/user/invite][%d] postAuthApiV1UserInviteOK ", 200)
}

func (o *PostAuthAPIV1UserInviteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
