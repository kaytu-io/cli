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

// GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendReader is a Reader for the GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrend structure.
type GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /compliance/api/v1/benchmark/{benchmark_id}/summary/result/trend] GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrend", response, response.Code())
	}
}

// NewGetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendOK creates a GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendOK with default headers values
func NewGetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendOK() *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendOK {
	return &GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendOK{}
}

/*
GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendOK describes a response with status code 200, with default header values.

OK
*/
type GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendOK struct {
	Payload *models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmarkResultTrend
}

// IsSuccess returns true when this get compliance Api v1 benchmark benchmark Id summary result trend o k response has a 2xx status code
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get compliance Api v1 benchmark benchmark Id summary result trend o k response has a 3xx status code
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get compliance Api v1 benchmark benchmark Id summary result trend o k response has a 4xx status code
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get compliance Api v1 benchmark benchmark Id summary result trend o k response has a 5xx status code
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get compliance Api v1 benchmark benchmark Id summary result trend o k response a status code equal to that given
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get compliance Api v1 benchmark benchmark Id summary result trend o k response
func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendOK) Code() int {
	return 200
}

func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendOK) Error() string {
	return fmt.Sprintf("[GET /compliance/api/v1/benchmark/{benchmark_id}/summary/result/trend][%d] getComplianceApiV1BenchmarkBenchmarkIdSummaryResultTrendOK  %+v", 200, o.Payload)
}

func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendOK) String() string {
	return fmt.Sprintf("[GET /compliance/api/v1/benchmark/{benchmark_id}/summary/result/trend][%d] getComplianceApiV1BenchmarkBenchmarkIdSummaryResultTrendOK  %+v", 200, o.Payload)
}

func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendOK) GetPayload() *models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmarkResultTrend {
	return o.Payload
}

func (o *GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIBenchmarkResultTrend)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
