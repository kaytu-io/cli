// Code generated by go-swagger; DO NOT EDIT.

package benchmarks_assignment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDReader is a Reader for the DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionID structure.
type DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /compliance/api/v1/assignments/{benchmark_id}/connection/{connection_id}] DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionID", response, response.Code())
	}
}

// NewDeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK creates a DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK with default headers values
func NewDeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK() *DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK {
	return &DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK{}
}

/*
DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK describes a response with status code 200, with default header values.

OK
*/
type DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK struct {
}

// IsSuccess returns true when this delete compliance Api v1 assignments benchmark Id connection connection Id o k response has a 2xx status code
func (o *DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete compliance Api v1 assignments benchmark Id connection connection Id o k response has a 3xx status code
func (o *DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete compliance Api v1 assignments benchmark Id connection connection Id o k response has a 4xx status code
func (o *DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete compliance Api v1 assignments benchmark Id connection connection Id o k response has a 5xx status code
func (o *DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete compliance Api v1 assignments benchmark Id connection connection Id o k response a status code equal to that given
func (o *DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete compliance Api v1 assignments benchmark Id connection connection Id o k response
func (o *DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK) Code() int {
	return 200
}

func (o *DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK) Error() string {
	return fmt.Sprintf("[DELETE /compliance/api/v1/assignments/{benchmark_id}/connection/{connection_id}][%d] deleteComplianceApiV1AssignmentsBenchmarkIdConnectionConnectionIdOK ", 200)
}

func (o *DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK) String() string {
	return fmt.Sprintf("[DELETE /compliance/api/v1/assignments/{benchmark_id}/connection/{connection_id}][%d] deleteComplianceApiV1AssignmentsBenchmarkIdConnectionConnectionIdOK ", 200)
}

func (o *DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
