// Code generated by go-swagger; DO NOT EDIT.

package keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAuthAPIV1KeyIDDeleteReader is a Reader for the DeleteAuthAPIV1KeyIDDelete structure.
type DeleteAuthAPIV1KeyIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAuthAPIV1KeyIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAuthAPIV1KeyIDDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /auth/api/v1/key/{id}/delete] DeleteAuthAPIV1KeyIDDelete", response, response.Code())
	}
}

// NewDeleteAuthAPIV1KeyIDDeleteOK creates a DeleteAuthAPIV1KeyIDDeleteOK with default headers values
func NewDeleteAuthAPIV1KeyIDDeleteOK() *DeleteAuthAPIV1KeyIDDeleteOK {
	return &DeleteAuthAPIV1KeyIDDeleteOK{}
}

/*
DeleteAuthAPIV1KeyIDDeleteOK describes a response with status code 200, with default header values.

OK
*/
type DeleteAuthAPIV1KeyIDDeleteOK struct {
}

// IsSuccess returns true when this delete auth Api v1 key Id delete o k response has a 2xx status code
func (o *DeleteAuthAPIV1KeyIDDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete auth Api v1 key Id delete o k response has a 3xx status code
func (o *DeleteAuthAPIV1KeyIDDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete auth Api v1 key Id delete o k response has a 4xx status code
func (o *DeleteAuthAPIV1KeyIDDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete auth Api v1 key Id delete o k response has a 5xx status code
func (o *DeleteAuthAPIV1KeyIDDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete auth Api v1 key Id delete o k response a status code equal to that given
func (o *DeleteAuthAPIV1KeyIDDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete auth Api v1 key Id delete o k response
func (o *DeleteAuthAPIV1KeyIDDeleteOK) Code() int {
	return 200
}

func (o *DeleteAuthAPIV1KeyIDDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /auth/api/v1/key/{id}/delete][%d] deleteAuthApiV1KeyIdDeleteOK ", 200)
}

func (o *DeleteAuthAPIV1KeyIDDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /auth/api/v1/key/{id}/delete][%d] deleteAuthApiV1KeyIdDeleteOK ", 200)
}

func (o *DeleteAuthAPIV1KeyIDDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
