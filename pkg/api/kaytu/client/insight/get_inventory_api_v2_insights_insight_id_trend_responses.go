// Code generated by go-swagger; DO NOT EDIT.

package insight

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

// GetInventoryAPIV2InsightsInsightIDTrendReader is a Reader for the GetInventoryAPIV2InsightsInsightIDTrend structure.
type GetInventoryAPIV2InsightsInsightIDTrendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventoryAPIV2InsightsInsightIDTrendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInventoryAPIV2InsightsInsightIDTrendOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /inventory/api/v2/insights/{insightId}/trend] GetInventoryAPIV2InsightsInsightIDTrend", response, response.Code())
	}
}

// NewGetInventoryAPIV2InsightsInsightIDTrendOK creates a GetInventoryAPIV2InsightsInsightIDTrendOK with default headers values
func NewGetInventoryAPIV2InsightsInsightIDTrendOK() *GetInventoryAPIV2InsightsInsightIDTrendOK {
	return &GetInventoryAPIV2InsightsInsightIDTrendOK{}
}

/*
GetInventoryAPIV2InsightsInsightIDTrendOK describes a response with status code 200, with default header values.

OK
*/
type GetInventoryAPIV2InsightsInsightIDTrendOK struct {
	Payload map[string][]models.GitlabComKeibiengineKeibiEnginePkgInsightEsInsightResource
}

// IsSuccess returns true when this get inventory Api v2 insights insight Id trend o k response has a 2xx status code
func (o *GetInventoryAPIV2InsightsInsightIDTrendOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get inventory Api v2 insights insight Id trend o k response has a 3xx status code
func (o *GetInventoryAPIV2InsightsInsightIDTrendOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get inventory Api v2 insights insight Id trend o k response has a 4xx status code
func (o *GetInventoryAPIV2InsightsInsightIDTrendOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get inventory Api v2 insights insight Id trend o k response has a 5xx status code
func (o *GetInventoryAPIV2InsightsInsightIDTrendOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get inventory Api v2 insights insight Id trend o k response a status code equal to that given
func (o *GetInventoryAPIV2InsightsInsightIDTrendOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get inventory Api v2 insights insight Id trend o k response
func (o *GetInventoryAPIV2InsightsInsightIDTrendOK) Code() int {
	return 200
}

func (o *GetInventoryAPIV2InsightsInsightIDTrendOK) Error() string {
	return fmt.Sprintf("[GET /inventory/api/v2/insights/{insightId}/trend][%d] getInventoryApiV2InsightsInsightIdTrendOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2InsightsInsightIDTrendOK) String() string {
	return fmt.Sprintf("[GET /inventory/api/v2/insights/{insightId}/trend][%d] getInventoryApiV2InsightsInsightIdTrendOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2InsightsInsightIDTrendOK) GetPayload() map[string][]models.GitlabComKeibiengineKeibiEnginePkgInsightEsInsightResource {
	return o.Payload
}

func (o *GetInventoryAPIV2InsightsInsightIDTrendOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
