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

// GetComplianceAPIV1InsightGroupInsightGroupIDTrendReader is a Reader for the GetComplianceAPIV1InsightGroupInsightGroupIDTrend structure.
type GetComplianceAPIV1InsightGroupInsightGroupIDTrendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetComplianceAPIV1InsightGroupInsightGroupIDTrendOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /compliance/api/v1/insight/group/{insightGroupId}/trend] GetComplianceAPIV1InsightGroupInsightGroupIDTrend", response, response.Code())
	}
}

// NewGetComplianceAPIV1InsightGroupInsightGroupIDTrendOK creates a GetComplianceAPIV1InsightGroupInsightGroupIDTrendOK with default headers values
func NewGetComplianceAPIV1InsightGroupInsightGroupIDTrendOK() *GetComplianceAPIV1InsightGroupInsightGroupIDTrendOK {
	return &GetComplianceAPIV1InsightGroupInsightGroupIDTrendOK{}
}

/*
GetComplianceAPIV1InsightGroupInsightGroupIDTrendOK describes a response with status code 200, with default header values.

OK
*/
type GetComplianceAPIV1InsightGroupInsightGroupIDTrendOK struct {
	Payload []*models.GithubComKaytuIoKaytuEnginePkgComplianceAPIInsightTrendDatapoint
}

// IsSuccess returns true when this get compliance Api v1 insight group insight group Id trend o k response has a 2xx status code
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get compliance Api v1 insight group insight group Id trend o k response has a 3xx status code
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get compliance Api v1 insight group insight group Id trend o k response has a 4xx status code
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get compliance Api v1 insight group insight group Id trend o k response has a 5xx status code
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get compliance Api v1 insight group insight group Id trend o k response a status code equal to that given
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get compliance Api v1 insight group insight group Id trend o k response
func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendOK) Code() int {
	return 200
}

func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendOK) Error() string {
	return fmt.Sprintf("[GET /compliance/api/v1/insight/group/{insightGroupId}/trend][%d] getComplianceApiV1InsightGroupInsightGroupIdTrendOK  %+v", 200, o.Payload)
}

func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendOK) String() string {
	return fmt.Sprintf("[GET /compliance/api/v1/insight/group/{insightGroupId}/trend][%d] getComplianceApiV1InsightGroupInsightGroupIdTrendOK  %+v", 200, o.Payload)
}

func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendOK) GetPayload() []*models.GithubComKaytuIoKaytuEnginePkgComplianceAPIInsightTrendDatapoint {
	return o.Payload
}

func (o *GetComplianceAPIV1InsightGroupInsightGroupIDTrendOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
