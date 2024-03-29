// Code generated by go-swagger; DO NOT EDIT.

package compliance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

// PostComplianceAPIV1FindingsResourceReader is a Reader for the PostComplianceAPIV1FindingsResource structure.
type PostComplianceAPIV1FindingsResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostComplianceAPIV1FindingsResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostComplianceAPIV1FindingsResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /compliance/api/v1/findings/resource] PostComplianceAPIV1FindingsResource", response, response.Code())
	}
}

// NewPostComplianceAPIV1FindingsResourceOK creates a PostComplianceAPIV1FindingsResourceOK with default headers values
func NewPostComplianceAPIV1FindingsResourceOK() *PostComplianceAPIV1FindingsResourceOK {
	return &PostComplianceAPIV1FindingsResourceOK{}
}

/*
PostComplianceAPIV1FindingsResourceOK describes a response with status code 200, with default header values.

OK
*/
type PostComplianceAPIV1FindingsResourceOK struct {
	Payload *models.GithubComKaytuIoKaytuEnginePkgComplianceAPIGetSingleResourceFindingResponse
}

// IsSuccess returns true when this post compliance Api v1 findings resource o k response has a 2xx status code
func (o *PostComplianceAPIV1FindingsResourceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post compliance Api v1 findings resource o k response has a 3xx status code
func (o *PostComplianceAPIV1FindingsResourceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post compliance Api v1 findings resource o k response has a 4xx status code
func (o *PostComplianceAPIV1FindingsResourceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post compliance Api v1 findings resource o k response has a 5xx status code
func (o *PostComplianceAPIV1FindingsResourceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post compliance Api v1 findings resource o k response a status code equal to that given
func (o *PostComplianceAPIV1FindingsResourceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post compliance Api v1 findings resource o k response
func (o *PostComplianceAPIV1FindingsResourceOK) Code() int {
	return 200
}

func (o *PostComplianceAPIV1FindingsResourceOK) Error() string {
	return fmt.Sprintf("[POST /compliance/api/v1/findings/resource][%d] postComplianceApiV1FindingsResourceOK  %+v", 200, o.Payload)
}

func (o *PostComplianceAPIV1FindingsResourceOK) String() string {
	return fmt.Sprintf("[POST /compliance/api/v1/findings/resource][%d] postComplianceApiV1FindingsResourceOK  %+v", 200, o.Payload)
}

func (o *PostComplianceAPIV1FindingsResourceOK) GetPayload() *models.GithubComKaytuIoKaytuEnginePkgComplianceAPIGetSingleResourceFindingResponse {
	return o.Payload
}

func (o *PostComplianceAPIV1FindingsResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GithubComKaytuIoKaytuEnginePkgComplianceAPIGetSingleResourceFindingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
