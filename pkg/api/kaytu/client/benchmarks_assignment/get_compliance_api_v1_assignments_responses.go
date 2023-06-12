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

// GetComplianceAPIV1AssignmentsReader is a Reader for the GetComplianceAPIV1Assignments structure.
type GetComplianceAPIV1AssignmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComplianceAPIV1AssignmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetComplianceAPIV1AssignmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /compliance/api/v1/assignments] GetComplianceAPIV1Assignments", response, response.Code())
	}
}

// NewGetComplianceAPIV1AssignmentsOK creates a GetComplianceAPIV1AssignmentsOK with default headers values
func NewGetComplianceAPIV1AssignmentsOK() *GetComplianceAPIV1AssignmentsOK {
	return &GetComplianceAPIV1AssignmentsOK{}
}

/*
GetComplianceAPIV1AssignmentsOK describes a response with status code 200, with default header values.

OK
*/
type GetComplianceAPIV1AssignmentsOK struct {
	Payload []*models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmarkAssignment
}

// IsSuccess returns true when this get compliance Api v1 assignments o k response has a 2xx status code
func (o *GetComplianceAPIV1AssignmentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get compliance Api v1 assignments o k response has a 3xx status code
func (o *GetComplianceAPIV1AssignmentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get compliance Api v1 assignments o k response has a 4xx status code
func (o *GetComplianceAPIV1AssignmentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get compliance Api v1 assignments o k response has a 5xx status code
func (o *GetComplianceAPIV1AssignmentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get compliance Api v1 assignments o k response a status code equal to that given
func (o *GetComplianceAPIV1AssignmentsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get compliance Api v1 assignments o k response
func (o *GetComplianceAPIV1AssignmentsOK) Code() int {
	return 200
}

func (o *GetComplianceAPIV1AssignmentsOK) Error() string {
	return fmt.Sprintf("[GET /compliance/api/v1/assignments][%d] getComplianceApiV1AssignmentsOK  %+v", 200, o.Payload)
}

func (o *GetComplianceAPIV1AssignmentsOK) String() string {
	return fmt.Sprintf("[GET /compliance/api/v1/assignments][%d] getComplianceApiV1AssignmentsOK  %+v", 200, o.Payload)
}

func (o *GetComplianceAPIV1AssignmentsOK) GetPayload() []*models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmarkAssignment {
	return o.Payload
}

func (o *GetComplianceAPIV1AssignmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
