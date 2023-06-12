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

// GetInventoryAPIV2InsightsInsightIDReader is a Reader for the GetInventoryAPIV2InsightsInsightID structure.
type GetInventoryAPIV2InsightsInsightIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventoryAPIV2InsightsInsightIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInventoryAPIV2InsightsInsightIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /inventory/api/v2/insights/{insightId}] GetInventoryAPIV2InsightsInsightID", response, response.Code())
	}
}

// NewGetInventoryAPIV2InsightsInsightIDOK creates a GetInventoryAPIV2InsightsInsightIDOK with default headers values
func NewGetInventoryAPIV2InsightsInsightIDOK() *GetInventoryAPIV2InsightsInsightIDOK {
	return &GetInventoryAPIV2InsightsInsightIDOK{}
}

/*
GetInventoryAPIV2InsightsInsightIDOK describes a response with status code 200, with default header values.

OK
*/
type GetInventoryAPIV2InsightsInsightIDOK struct {
	Payload *models.GitlabComKeibiengineKeibiEnginePkgInsightEsInsightResource
}

// IsSuccess returns true when this get inventory Api v2 insights insight Id o k response has a 2xx status code
func (o *GetInventoryAPIV2InsightsInsightIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get inventory Api v2 insights insight Id o k response has a 3xx status code
func (o *GetInventoryAPIV2InsightsInsightIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get inventory Api v2 insights insight Id o k response has a 4xx status code
func (o *GetInventoryAPIV2InsightsInsightIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get inventory Api v2 insights insight Id o k response has a 5xx status code
func (o *GetInventoryAPIV2InsightsInsightIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get inventory Api v2 insights insight Id o k response a status code equal to that given
func (o *GetInventoryAPIV2InsightsInsightIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get inventory Api v2 insights insight Id o k response
func (o *GetInventoryAPIV2InsightsInsightIDOK) Code() int {
	return 200
}

func (o *GetInventoryAPIV2InsightsInsightIDOK) Error() string {
	return fmt.Sprintf("[GET /inventory/api/v2/insights/{insightId}][%d] getInventoryApiV2InsightsInsightIdOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2InsightsInsightIDOK) String() string {
	return fmt.Sprintf("[GET /inventory/api/v2/insights/{insightId}][%d] getInventoryApiV2InsightsInsightIdOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2InsightsInsightIDOK) GetPayload() *models.GitlabComKeibiengineKeibiEnginePkgInsightEsInsightResource {
	return o.Payload
}

func (o *GetInventoryAPIV2InsightsInsightIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitlabComKeibiengineKeibiEnginePkgInsightEsInsightResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
