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

// GetComplianceAPIV1FindingsMetricsReader is a Reader for the GetComplianceAPIV1FindingsMetrics structure.
type GetComplianceAPIV1FindingsMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComplianceAPIV1FindingsMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetComplianceAPIV1FindingsMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /compliance/api/v1/findings/metrics] GetComplianceAPIV1FindingsMetrics", response, response.Code())
	}
}

// NewGetComplianceAPIV1FindingsMetricsOK creates a GetComplianceAPIV1FindingsMetricsOK with default headers values
func NewGetComplianceAPIV1FindingsMetricsOK() *GetComplianceAPIV1FindingsMetricsOK {
	return &GetComplianceAPIV1FindingsMetricsOK{}
}

/*
GetComplianceAPIV1FindingsMetricsOK describes a response with status code 200, with default header values.

OK
*/
type GetComplianceAPIV1FindingsMetricsOK struct {
	Payload *models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIGetFindingsMetricsResponse
}

// IsSuccess returns true when this get compliance Api v1 findings metrics o k response has a 2xx status code
func (o *GetComplianceAPIV1FindingsMetricsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get compliance Api v1 findings metrics o k response has a 3xx status code
func (o *GetComplianceAPIV1FindingsMetricsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get compliance Api v1 findings metrics o k response has a 4xx status code
func (o *GetComplianceAPIV1FindingsMetricsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get compliance Api v1 findings metrics o k response has a 5xx status code
func (o *GetComplianceAPIV1FindingsMetricsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get compliance Api v1 findings metrics o k response a status code equal to that given
func (o *GetComplianceAPIV1FindingsMetricsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get compliance Api v1 findings metrics o k response
func (o *GetComplianceAPIV1FindingsMetricsOK) Code() int {
	return 200
}

func (o *GetComplianceAPIV1FindingsMetricsOK) Error() string {
	return fmt.Sprintf("[GET /compliance/api/v1/findings/metrics][%d] getComplianceApiV1FindingsMetricsOK  %+v", 200, o.Payload)
}

func (o *GetComplianceAPIV1FindingsMetricsOK) String() string {
	return fmt.Sprintf("[GET /compliance/api/v1/findings/metrics][%d] getComplianceApiV1FindingsMetricsOK  %+v", 200, o.Payload)
}

func (o *GetComplianceAPIV1FindingsMetricsOK) GetPayload() *models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIGetFindingsMetricsResponse {
	return o.Payload
}

func (o *GetComplianceAPIV1FindingsMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIGetFindingsMetricsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
