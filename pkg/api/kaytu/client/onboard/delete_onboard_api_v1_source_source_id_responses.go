// Code generated by go-swagger; DO NOT EDIT.

package onboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteOnboardAPIV1SourceSourceIDReader is a Reader for the DeleteOnboardAPIV1SourceSourceID structure.
type DeleteOnboardAPIV1SourceSourceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOnboardAPIV1SourceSourceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOnboardAPIV1SourceSourceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOnboardAPIV1SourceSourceIDOK creates a DeleteOnboardAPIV1SourceSourceIDOK with default headers values
func NewDeleteOnboardAPIV1SourceSourceIDOK() *DeleteOnboardAPIV1SourceSourceIDOK {
	return &DeleteOnboardAPIV1SourceSourceIDOK{}
}

/*
DeleteOnboardAPIV1SourceSourceIDOK describes a response with status code 200, with default header values.

OK
*/
type DeleteOnboardAPIV1SourceSourceIDOK struct {
}

// IsSuccess returns true when this delete onboard Api v1 source source Id o k response has a 2xx status code
func (o *DeleteOnboardAPIV1SourceSourceIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete onboard Api v1 source source Id o k response has a 3xx status code
func (o *DeleteOnboardAPIV1SourceSourceIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete onboard Api v1 source source Id o k response has a 4xx status code
func (o *DeleteOnboardAPIV1SourceSourceIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete onboard Api v1 source source Id o k response has a 5xx status code
func (o *DeleteOnboardAPIV1SourceSourceIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete onboard Api v1 source source Id o k response a status code equal to that given
func (o *DeleteOnboardAPIV1SourceSourceIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete onboard Api v1 source source Id o k response
func (o *DeleteOnboardAPIV1SourceSourceIDOK) Code() int {
	return 200
}

func (o *DeleteOnboardAPIV1SourceSourceIDOK) Error() string {
	return fmt.Sprintf("[DELETE /onboard/api/v1/source/{sourceId}][%d] deleteOnboardApiV1SourceSourceIdOK ", 200)
}

func (o *DeleteOnboardAPIV1SourceSourceIDOK) String() string {
	return fmt.Sprintf("[DELETE /onboard/api/v1/source/{sourceId}][%d] deleteOnboardApiV1SourceSourceIdOK ", 200)
}

func (o *DeleteOnboardAPIV1SourceSourceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
