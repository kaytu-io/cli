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

// GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesReader is a Reader for the GetComplianceAPIV1BenchmarksBenchmarkIDPolicies structure.
type GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetComplianceAPIV1BenchmarksBenchmarkIDPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /compliance/api/v1/benchmarks/{benchmark_id}/policies] GetComplianceAPIV1BenchmarksBenchmarkIDPolicies", response, response.Code())
	}
}

// NewGetComplianceAPIV1BenchmarksBenchmarkIDPoliciesOK creates a GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesOK with default headers values
func NewGetComplianceAPIV1BenchmarksBenchmarkIDPoliciesOK() *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesOK {
	return &GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesOK{}
}

/*
GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesOK describes a response with status code 200, with default header values.

OK
*/
type GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesOK struct {
	Payload []*models.GithubComKaytuIoKaytuEnginePkgComplianceAPIPolicy
}

// IsSuccess returns true when this get compliance Api v1 benchmarks benchmark Id policies o k response has a 2xx status code
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get compliance Api v1 benchmarks benchmark Id policies o k response has a 3xx status code
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get compliance Api v1 benchmarks benchmark Id policies o k response has a 4xx status code
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get compliance Api v1 benchmarks benchmark Id policies o k response has a 5xx status code
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get compliance Api v1 benchmarks benchmark Id policies o k response a status code equal to that given
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get compliance Api v1 benchmarks benchmark Id policies o k response
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesOK) Code() int {
	return 200
}

func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /compliance/api/v1/benchmarks/{benchmark_id}/policies][%d] getComplianceApiV1BenchmarksBenchmarkIdPoliciesOK  %+v", 200, o.Payload)
}

func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesOK) String() string {
	return fmt.Sprintf("[GET /compliance/api/v1/benchmarks/{benchmark_id}/policies][%d] getComplianceApiV1BenchmarksBenchmarkIdPoliciesOK  %+v", 200, o.Payload)
}

func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesOK) GetPayload() []*models.GithubComKaytuIoKaytuEnginePkgComplianceAPIPolicy {
	return o.Payload
}

func (o *GetComplianceAPIV1BenchmarksBenchmarkIDPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
