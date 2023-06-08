// Code generated by go-swagger; DO NOT EDIT.

package insights

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

// GetComplianceAPIV1InsightInsightIDTrendReader is a Reader for the GetComplianceAPIV1InsightInsightIDTrend structure.
type GetComplianceAPIV1InsightInsightIDTrendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComplianceAPIV1InsightInsightIDTrendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetComplianceAPIV1InsightInsightIDTrendOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetComplianceAPIV1InsightInsightIDTrendOK creates a GetComplianceAPIV1InsightInsightIDTrendOK with default headers values
func NewGetComplianceAPIV1InsightInsightIDTrendOK() *GetComplianceAPIV1InsightInsightIDTrendOK {
	return &GetComplianceAPIV1InsightInsightIDTrendOK{}
}

/*
GetComplianceAPIV1InsightInsightIDTrendOK describes a response with status code 200, with default header values.

OK
*/
type GetComplianceAPIV1InsightInsightIDTrendOK struct {
	Payload []*models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIInsightTrendDatapoint
}

// IsSuccess returns true when this get compliance Api v1 insight insight Id trend o k response has a 2xx status code
func (o *GetComplianceAPIV1InsightInsightIDTrendOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get compliance Api v1 insight insight Id trend o k response has a 3xx status code
func (o *GetComplianceAPIV1InsightInsightIDTrendOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get compliance Api v1 insight insight Id trend o k response has a 4xx status code
func (o *GetComplianceAPIV1InsightInsightIDTrendOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get compliance Api v1 insight insight Id trend o k response has a 5xx status code
func (o *GetComplianceAPIV1InsightInsightIDTrendOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get compliance Api v1 insight insight Id trend o k response a status code equal to that given
func (o *GetComplianceAPIV1InsightInsightIDTrendOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get compliance Api v1 insight insight Id trend o k response
func (o *GetComplianceAPIV1InsightInsightIDTrendOK) Code() int {
	return 200
}

func (o *GetComplianceAPIV1InsightInsightIDTrendOK) Error() string {
	return fmt.Sprintf("[GET /compliance/api/v1/insight/{insightId}/trend][%d] getComplianceApiV1InsightInsightIdTrendOK  %+v", 200, o.Payload)
}

func (o *GetComplianceAPIV1InsightInsightIDTrendOK) String() string {
	return fmt.Sprintf("[GET /compliance/api/v1/insight/{insightId}/trend][%d] getComplianceApiV1InsightInsightIdTrendOK  %+v", 200, o.Payload)
}

func (o *GetComplianceAPIV1InsightInsightIDTrendOK) GetPayload() []*models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIInsightTrendDatapoint {
	return o.Payload
}

func (o *GetComplianceAPIV1InsightInsightIDTrendOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}