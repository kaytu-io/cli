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

// PostComplianceAPIV1AssignmentsBenchmarkIDConnectionReader is a Reader for the PostComplianceAPIV1AssignmentsBenchmarkIDConnection structure.
type PostComplianceAPIV1AssignmentsBenchmarkIDConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostComplianceAPIV1AssignmentsBenchmarkIDConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostComplianceAPIV1AssignmentsBenchmarkIDConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /compliance/api/v1/assignments/{benchmark_id}/connection] PostComplianceAPIV1AssignmentsBenchmarkIDConnection", response, response.Code())
	}
}

// NewPostComplianceAPIV1AssignmentsBenchmarkIDConnectionOK creates a PostComplianceAPIV1AssignmentsBenchmarkIDConnectionOK with default headers values
func NewPostComplianceAPIV1AssignmentsBenchmarkIDConnectionOK() *PostComplianceAPIV1AssignmentsBenchmarkIDConnectionOK {
	return &PostComplianceAPIV1AssignmentsBenchmarkIDConnectionOK{}
}

/*
PostComplianceAPIV1AssignmentsBenchmarkIDConnectionOK describes a response with status code 200, with default header values.

OK
*/
type PostComplianceAPIV1AssignmentsBenchmarkIDConnectionOK struct {
	Payload []*models.GithubComKaytuIoKaytuEnginePkgComplianceAPIBenchmarkAssignment
}

// IsSuccess returns true when this post compliance Api v1 assignments benchmark Id connection o k response has a 2xx status code
func (o *PostComplianceAPIV1AssignmentsBenchmarkIDConnectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post compliance Api v1 assignments benchmark Id connection o k response has a 3xx status code
func (o *PostComplianceAPIV1AssignmentsBenchmarkIDConnectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post compliance Api v1 assignments benchmark Id connection o k response has a 4xx status code
func (o *PostComplianceAPIV1AssignmentsBenchmarkIDConnectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post compliance Api v1 assignments benchmark Id connection o k response has a 5xx status code
func (o *PostComplianceAPIV1AssignmentsBenchmarkIDConnectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post compliance Api v1 assignments benchmark Id connection o k response a status code equal to that given
func (o *PostComplianceAPIV1AssignmentsBenchmarkIDConnectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post compliance Api v1 assignments benchmark Id connection o k response
func (o *PostComplianceAPIV1AssignmentsBenchmarkIDConnectionOK) Code() int {
	return 200
}

func (o *PostComplianceAPIV1AssignmentsBenchmarkIDConnectionOK) Error() string {
	return fmt.Sprintf("[POST /compliance/api/v1/assignments/{benchmark_id}/connection][%d] postComplianceApiV1AssignmentsBenchmarkIdConnectionOK  %+v", 200, o.Payload)
}

func (o *PostComplianceAPIV1AssignmentsBenchmarkIDConnectionOK) String() string {
	return fmt.Sprintf("[POST /compliance/api/v1/assignments/{benchmark_id}/connection][%d] postComplianceApiV1AssignmentsBenchmarkIdConnectionOK  %+v", 200, o.Payload)
}

func (o *PostComplianceAPIV1AssignmentsBenchmarkIDConnectionOK) GetPayload() []*models.GithubComKaytuIoKaytuEnginePkgComplianceAPIBenchmarkAssignment {
	return o.Payload
}

func (o *PostComplianceAPIV1AssignmentsBenchmarkIDConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
