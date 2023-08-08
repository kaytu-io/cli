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

// GetComplianceAPIV1BenchmarksBenchmarkIDTreeReader is a Reader for the GetComplianceAPIV1BenchmarksBenchmarkIDTree structure.
type GetComplianceAPIV1BenchmarksBenchmarkIDTreeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDTreeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetComplianceAPIV1BenchmarksBenchmarkIDTreeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /compliance/api/v1/benchmarks/{benchmark_id}/tree] GetComplianceAPIV1BenchmarksBenchmarkIDTree", response, response.Code())
	}
}

// NewGetComplianceAPIV1BenchmarksBenchmarkIDTreeOK creates a GetComplianceAPIV1BenchmarksBenchmarkIDTreeOK with default headers values
func NewGetComplianceAPIV1BenchmarksBenchmarkIDTreeOK() *GetComplianceAPIV1BenchmarksBenchmarkIDTreeOK {
	return &GetComplianceAPIV1BenchmarksBenchmarkIDTreeOK{}
}

/*
GetComplianceAPIV1BenchmarksBenchmarkIDTreeOK describes a response with status code 200, with default header values.

OK
*/
type GetComplianceAPIV1BenchmarksBenchmarkIDTreeOK struct {
	Payload *models.GithubComKaytuIoKaytuEnginePkgComplianceAPIBenchmarkTree
}

// IsSuccess returns true when this get compliance Api v1 benchmarks benchmark Id tree o k response has a 2xx status code
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDTreeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get compliance Api v1 benchmarks benchmark Id tree o k response has a 3xx status code
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDTreeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get compliance Api v1 benchmarks benchmark Id tree o k response has a 4xx status code
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDTreeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get compliance Api v1 benchmarks benchmark Id tree o k response has a 5xx status code
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDTreeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get compliance Api v1 benchmarks benchmark Id tree o k response a status code equal to that given
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDTreeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get compliance Api v1 benchmarks benchmark Id tree o k response
func (o *GetComplianceAPIV1BenchmarksBenchmarkIDTreeOK) Code() int {
	return 200
}

func (o *GetComplianceAPIV1BenchmarksBenchmarkIDTreeOK) Error() string {
	return fmt.Sprintf("[GET /compliance/api/v1/benchmarks/{benchmark_id}/tree][%d] getComplianceApiV1BenchmarksBenchmarkIdTreeOK  %+v", 200, o.Payload)
}

func (o *GetComplianceAPIV1BenchmarksBenchmarkIDTreeOK) String() string {
	return fmt.Sprintf("[GET /compliance/api/v1/benchmarks/{benchmark_id}/tree][%d] getComplianceApiV1BenchmarksBenchmarkIdTreeOK  %+v", 200, o.Payload)
}

func (o *GetComplianceAPIV1BenchmarksBenchmarkIDTreeOK) GetPayload() *models.GithubComKaytuIoKaytuEnginePkgComplianceAPIBenchmarkTree {
	return o.Payload
}

func (o *GetComplianceAPIV1BenchmarksBenchmarkIDTreeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GithubComKaytuIoKaytuEnginePkgComplianceAPIBenchmarkTree)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
