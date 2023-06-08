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

// GetComplianceAPIV1BenchmarksReader is a Reader for the GetComplianceAPIV1Benchmarks structure.
type GetComplianceAPIV1BenchmarksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComplianceAPIV1BenchmarksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetComplianceAPIV1BenchmarksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetComplianceAPIV1BenchmarksOK creates a GetComplianceAPIV1BenchmarksOK with default headers values
func NewGetComplianceAPIV1BenchmarksOK() *GetComplianceAPIV1BenchmarksOK {
	return &GetComplianceAPIV1BenchmarksOK{}
}

/*
GetComplianceAPIV1BenchmarksOK describes a response with status code 200, with default header values.

OK
*/
type GetComplianceAPIV1BenchmarksOK struct {
	Payload []*models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmark
}

// IsSuccess returns true when this get compliance Api v1 benchmarks o k response has a 2xx status code
func (o *GetComplianceAPIV1BenchmarksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get compliance Api v1 benchmarks o k response has a 3xx status code
func (o *GetComplianceAPIV1BenchmarksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get compliance Api v1 benchmarks o k response has a 4xx status code
func (o *GetComplianceAPIV1BenchmarksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get compliance Api v1 benchmarks o k response has a 5xx status code
func (o *GetComplianceAPIV1BenchmarksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get compliance Api v1 benchmarks o k response a status code equal to that given
func (o *GetComplianceAPIV1BenchmarksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get compliance Api v1 benchmarks o k response
func (o *GetComplianceAPIV1BenchmarksOK) Code() int {
	return 200
}

func (o *GetComplianceAPIV1BenchmarksOK) Error() string {
	return fmt.Sprintf("[GET /compliance/api/v1/benchmarks][%d] getComplianceApiV1BenchmarksOK  %+v", 200, o.Payload)
}

func (o *GetComplianceAPIV1BenchmarksOK) String() string {
	return fmt.Sprintf("[GET /compliance/api/v1/benchmarks][%d] getComplianceApiV1BenchmarksOK  %+v", 200, o.Payload)
}

func (o *GetComplianceAPIV1BenchmarksOK) GetPayload() []*models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmark {
	return o.Payload
}

func (o *GetComplianceAPIV1BenchmarksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}