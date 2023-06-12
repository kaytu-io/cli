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

// PostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDReader is a Reader for the PostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionID structure.
type PostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /compliance/api/v1/assignments/{benchmark_id}/connection/{connection_id}] PostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionID", response, response.Code())
	}
}

// NewPostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK creates a PostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK with default headers values
func NewPostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK() *PostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK {
	return &PostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK{}
}

/*
PostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK describes a response with status code 200, with default header values.

OK
*/
type PostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK struct {
	Payload *models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmarkAssignment
}

// IsSuccess returns true when this post compliance Api v1 assignments benchmark Id connection connection Id o k response has a 2xx status code
func (o *PostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post compliance Api v1 assignments benchmark Id connection connection Id o k response has a 3xx status code
func (o *PostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post compliance Api v1 assignments benchmark Id connection connection Id o k response has a 4xx status code
func (o *PostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post compliance Api v1 assignments benchmark Id connection connection Id o k response has a 5xx status code
func (o *PostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post compliance Api v1 assignments benchmark Id connection connection Id o k response a status code equal to that given
func (o *PostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post compliance Api v1 assignments benchmark Id connection connection Id o k response
func (o *PostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK) Code() int {
	return 200
}

func (o *PostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK) Error() string {
	return fmt.Sprintf("[POST /compliance/api/v1/assignments/{benchmark_id}/connection/{connection_id}][%d] postComplianceApiV1AssignmentsBenchmarkIdConnectionConnectionIdOK  %+v", 200, o.Payload)
}

func (o *PostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK) String() string {
	return fmt.Sprintf("[POST /compliance/api/v1/assignments/{benchmark_id}/connection/{connection_id}][%d] postComplianceApiV1AssignmentsBenchmarkIdConnectionConnectionIdOK  %+v", 200, o.Payload)
}

func (o *PostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK) GetPayload() *models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmarkAssignment {
	return o.Payload
}

func (o *PostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmarkAssignment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
