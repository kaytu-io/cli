// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

// GetInventoryAPIV2AnalyticsSpendTableReader is a Reader for the GetInventoryAPIV2AnalyticsSpendTable structure.
type GetInventoryAPIV2AnalyticsSpendTableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventoryAPIV2AnalyticsSpendTableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInventoryAPIV2AnalyticsSpendTableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /inventory/api/v2/analytics/spend/table] GetInventoryAPIV2AnalyticsSpendTable", response, response.Code())
	}
}

// NewGetInventoryAPIV2AnalyticsSpendTableOK creates a GetInventoryAPIV2AnalyticsSpendTableOK with default headers values
func NewGetInventoryAPIV2AnalyticsSpendTableOK() *GetInventoryAPIV2AnalyticsSpendTableOK {
	return &GetInventoryAPIV2AnalyticsSpendTableOK{}
}

/*
GetInventoryAPIV2AnalyticsSpendTableOK describes a response with status code 200, with default header values.

OK
*/
type GetInventoryAPIV2AnalyticsSpendTableOK struct {
	Payload []*models.GithubComKaytuIoKaytuEnginePkgInventoryAPISpendTableRow
}

// IsSuccess returns true when this get inventory Api v2 analytics spend table o k response has a 2xx status code
func (o *GetInventoryAPIV2AnalyticsSpendTableOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get inventory Api v2 analytics spend table o k response has a 3xx status code
func (o *GetInventoryAPIV2AnalyticsSpendTableOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get inventory Api v2 analytics spend table o k response has a 4xx status code
func (o *GetInventoryAPIV2AnalyticsSpendTableOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get inventory Api v2 analytics spend table o k response has a 5xx status code
func (o *GetInventoryAPIV2AnalyticsSpendTableOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get inventory Api v2 analytics spend table o k response a status code equal to that given
func (o *GetInventoryAPIV2AnalyticsSpendTableOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get inventory Api v2 analytics spend table o k response
func (o *GetInventoryAPIV2AnalyticsSpendTableOK) Code() int {
	return 200
}

func (o *GetInventoryAPIV2AnalyticsSpendTableOK) Error() string {
	return fmt.Sprintf("[GET /inventory/api/v2/analytics/spend/table][%d] getInventoryApiV2AnalyticsSpendTableOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2AnalyticsSpendTableOK) String() string {
	return fmt.Sprintf("[GET /inventory/api/v2/analytics/spend/table][%d] getInventoryApiV2AnalyticsSpendTableOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2AnalyticsSpendTableOK) GetPayload() []*models.GithubComKaytuIoKaytuEnginePkgInventoryAPISpendTableRow {
	return o.Payload
}

func (o *GetInventoryAPIV2AnalyticsSpendTableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
