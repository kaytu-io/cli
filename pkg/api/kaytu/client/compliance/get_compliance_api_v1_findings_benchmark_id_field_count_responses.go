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

// GetComplianceAPIV1FindingsBenchmarkIDFieldCountReader is a Reader for the GetComplianceAPIV1FindingsBenchmarkIDFieldCount structure.
type GetComplianceAPIV1FindingsBenchmarkIDFieldCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetComplianceAPIV1FindingsBenchmarkIDFieldCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /compliance/api/v1/findings/{benchmarkId}/{field}/count] GetComplianceAPIV1FindingsBenchmarkIDFieldCount", response, response.Code())
	}
}

// NewGetComplianceAPIV1FindingsBenchmarkIDFieldCountOK creates a GetComplianceAPIV1FindingsBenchmarkIDFieldCountOK with default headers values
func NewGetComplianceAPIV1FindingsBenchmarkIDFieldCountOK() *GetComplianceAPIV1FindingsBenchmarkIDFieldCountOK {
	return &GetComplianceAPIV1FindingsBenchmarkIDFieldCountOK{}
}

/*
GetComplianceAPIV1FindingsBenchmarkIDFieldCountOK describes a response with status code 200, with default header values.

OK
*/
type GetComplianceAPIV1FindingsBenchmarkIDFieldCountOK struct {
	Payload *models.GithubComKaytuIoKaytuEnginePkgComplianceAPIGetTopFieldResponse
}

// IsSuccess returns true when this get compliance Api v1 findings benchmark Id field count o k response has a 2xx status code
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldCountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get compliance Api v1 findings benchmark Id field count o k response has a 3xx status code
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldCountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get compliance Api v1 findings benchmark Id field count o k response has a 4xx status code
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldCountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get compliance Api v1 findings benchmark Id field count o k response has a 5xx status code
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldCountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get compliance Api v1 findings benchmark Id field count o k response a status code equal to that given
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldCountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get compliance Api v1 findings benchmark Id field count o k response
func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldCountOK) Code() int {
	return 200
}

func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldCountOK) Error() string {
	return fmt.Sprintf("[GET /compliance/api/v1/findings/{benchmarkId}/{field}/count][%d] getComplianceApiV1FindingsBenchmarkIdFieldCountOK  %+v", 200, o.Payload)
}

func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldCountOK) String() string {
	return fmt.Sprintf("[GET /compliance/api/v1/findings/{benchmarkId}/{field}/count][%d] getComplianceApiV1FindingsBenchmarkIdFieldCountOK  %+v", 200, o.Payload)
}

func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldCountOK) GetPayload() *models.GithubComKaytuIoKaytuEnginePkgComplianceAPIGetTopFieldResponse {
	return o.Payload
}

func (o *GetComplianceAPIV1FindingsBenchmarkIDFieldCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GithubComKaytuIoKaytuEnginePkgComplianceAPIGetTopFieldResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
