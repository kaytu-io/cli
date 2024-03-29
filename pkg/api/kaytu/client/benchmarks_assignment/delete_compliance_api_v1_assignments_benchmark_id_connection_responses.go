// Code generated by go-swagger; DO NOT EDIT.

package benchmarks_assignment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionReader is a Reader for the DeleteComplianceAPIV1AssignmentsBenchmarkIDConnection structure.
type DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /compliance/api/v1/assignments/{benchmark_id}/connection] DeleteComplianceAPIV1AssignmentsBenchmarkIDConnection", response, response.Code())
	}
}

// NewDeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionOK creates a DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionOK with default headers values
func NewDeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionOK() *DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionOK {
	return &DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionOK{}
}

/*
DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionOK describes a response with status code 200, with default header values.

OK
*/
type DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionOK struct {
}

// IsSuccess returns true when this delete compliance Api v1 assignments benchmark Id connection o k response has a 2xx status code
func (o *DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete compliance Api v1 assignments benchmark Id connection o k response has a 3xx status code
func (o *DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete compliance Api v1 assignments benchmark Id connection o k response has a 4xx status code
func (o *DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete compliance Api v1 assignments benchmark Id connection o k response has a 5xx status code
func (o *DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete compliance Api v1 assignments benchmark Id connection o k response a status code equal to that given
func (o *DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete compliance Api v1 assignments benchmark Id connection o k response
func (o *DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionOK) Code() int {
	return 200
}

func (o *DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionOK) Error() string {
	return fmt.Sprintf("[DELETE /compliance/api/v1/assignments/{benchmark_id}/connection][%d] deleteComplianceApiV1AssignmentsBenchmarkIdConnectionOK ", 200)
}

func (o *DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionOK) String() string {
	return fmt.Sprintf("[DELETE /compliance/api/v1/assignments/{benchmark_id}/connection][%d] deleteComplianceApiV1AssignmentsBenchmarkIdConnectionOK ", 200)
}

func (o *DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
