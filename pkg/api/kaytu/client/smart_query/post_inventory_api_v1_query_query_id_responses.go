// Code generated by go-swagger; DO NOT EDIT.

package smart_query

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

// PostInventoryAPIV1QueryQueryIDReader is a Reader for the PostInventoryAPIV1QueryQueryID structure.
type PostInventoryAPIV1QueryQueryIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostInventoryAPIV1QueryQueryIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostInventoryAPIV1QueryQueryIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostInventoryAPIV1QueryQueryIDOK creates a PostInventoryAPIV1QueryQueryIDOK with default headers values
func NewPostInventoryAPIV1QueryQueryIDOK() *PostInventoryAPIV1QueryQueryIDOK {
	return &PostInventoryAPIV1QueryQueryIDOK{}
}

/*
PostInventoryAPIV1QueryQueryIDOK describes a response with status code 200, with default header values.

OK
*/
type PostInventoryAPIV1QueryQueryIDOK struct {
	Payload *models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIRunQueryResponse
}

// IsSuccess returns true when this post inventory Api v1 query query Id o k response has a 2xx status code
func (o *PostInventoryAPIV1QueryQueryIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post inventory Api v1 query query Id o k response has a 3xx status code
func (o *PostInventoryAPIV1QueryQueryIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post inventory Api v1 query query Id o k response has a 4xx status code
func (o *PostInventoryAPIV1QueryQueryIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post inventory Api v1 query query Id o k response has a 5xx status code
func (o *PostInventoryAPIV1QueryQueryIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post inventory Api v1 query query Id o k response a status code equal to that given
func (o *PostInventoryAPIV1QueryQueryIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post inventory Api v1 query query Id o k response
func (o *PostInventoryAPIV1QueryQueryIDOK) Code() int {
	return 200
}

func (o *PostInventoryAPIV1QueryQueryIDOK) Error() string {
	return fmt.Sprintf("[POST /inventory/api/v1/query/{queryId}][%d] postInventoryApiV1QueryQueryIdOK  %+v", 200, o.Payload)
}

func (o *PostInventoryAPIV1QueryQueryIDOK) String() string {
	return fmt.Sprintf("[POST /inventory/api/v1/query/{queryId}][%d] postInventoryApiV1QueryQueryIdOK  %+v", 200, o.Payload)
}

func (o *PostInventoryAPIV1QueryQueryIDOK) GetPayload() *models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIRunQueryResponse {
	return o.Payload
}

func (o *PostInventoryAPIV1QueryQueryIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIRunQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
