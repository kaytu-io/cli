// Code generated by go-swagger; DO NOT EDIT.

package benchmarks_assignment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

// GetComplianceAPIV1AssignmentsResourceCollectionResourceCollectionIDReader is a Reader for the GetComplianceAPIV1AssignmentsResourceCollectionResourceCollectionID structure.
type GetComplianceAPIV1AssignmentsResourceCollectionResourceCollectionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComplianceAPIV1AssignmentsResourceCollectionResourceCollectionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetComplianceAPIV1AssignmentsResourceCollectionResourceCollectionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /compliance/api/v1/assignments/resource_collection/{resource_collection_id}] GetComplianceAPIV1AssignmentsResourceCollectionResourceCollectionID", response, response.Code())
	}
}

// NewGetComplianceAPIV1AssignmentsResourceCollectionResourceCollectionIDOK creates a GetComplianceAPIV1AssignmentsResourceCollectionResourceCollectionIDOK with default headers values
func NewGetComplianceAPIV1AssignmentsResourceCollectionResourceCollectionIDOK() *GetComplianceAPIV1AssignmentsResourceCollectionResourceCollectionIDOK {
	return &GetComplianceAPIV1AssignmentsResourceCollectionResourceCollectionIDOK{}
}

/*
GetComplianceAPIV1AssignmentsResourceCollectionResourceCollectionIDOK describes a response with status code 200, with default header values.

OK
*/
type GetComplianceAPIV1AssignmentsResourceCollectionResourceCollectionIDOK struct {
	Payload []*models.GithubComKaytuIoKaytuEnginePkgComplianceAPIAssignedBenchmark
}

// IsSuccess returns true when this get compliance Api v1 assignments resource collection resource collection Id o k response has a 2xx status code
func (o *GetComplianceAPIV1AssignmentsResourceCollectionResourceCollectionIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get compliance Api v1 assignments resource collection resource collection Id o k response has a 3xx status code
func (o *GetComplianceAPIV1AssignmentsResourceCollectionResourceCollectionIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get compliance Api v1 assignments resource collection resource collection Id o k response has a 4xx status code
func (o *GetComplianceAPIV1AssignmentsResourceCollectionResourceCollectionIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get compliance Api v1 assignments resource collection resource collection Id o k response has a 5xx status code
func (o *GetComplianceAPIV1AssignmentsResourceCollectionResourceCollectionIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get compliance Api v1 assignments resource collection resource collection Id o k response a status code equal to that given
func (o *GetComplianceAPIV1AssignmentsResourceCollectionResourceCollectionIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get compliance Api v1 assignments resource collection resource collection Id o k response
func (o *GetComplianceAPIV1AssignmentsResourceCollectionResourceCollectionIDOK) Code() int {
	return 200
}

func (o *GetComplianceAPIV1AssignmentsResourceCollectionResourceCollectionIDOK) Error() string {
	return fmt.Sprintf("[GET /compliance/api/v1/assignments/resource_collection/{resource_collection_id}][%d] getComplianceApiV1AssignmentsResourceCollectionResourceCollectionIdOK  %+v", 200, o.Payload)
}

func (o *GetComplianceAPIV1AssignmentsResourceCollectionResourceCollectionIDOK) String() string {
	return fmt.Sprintf("[GET /compliance/api/v1/assignments/resource_collection/{resource_collection_id}][%d] getComplianceApiV1AssignmentsResourceCollectionResourceCollectionIdOK  %+v", 200, o.Payload)
}

func (o *GetComplianceAPIV1AssignmentsResourceCollectionResourceCollectionIDOK) GetPayload() []*models.GithubComKaytuIoKaytuEnginePkgComplianceAPIAssignedBenchmark {
	return o.Payload
}

func (o *GetComplianceAPIV1AssignmentsResourceCollectionResourceCollectionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
