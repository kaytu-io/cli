// Code generated by go-swagger; DO NOT EDIT.

package analytics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

// GetInventoryAPIV2AnalyticsTableReader is a Reader for the GetInventoryAPIV2AnalyticsTable structure.
type GetInventoryAPIV2AnalyticsTableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInventoryAPIV2AnalyticsTableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInventoryAPIV2AnalyticsTableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /inventory/api/v2/analytics/table] GetInventoryAPIV2AnalyticsTable", response, response.Code())
	}
}

// NewGetInventoryAPIV2AnalyticsTableOK creates a GetInventoryAPIV2AnalyticsTableOK with default headers values
func NewGetInventoryAPIV2AnalyticsTableOK() *GetInventoryAPIV2AnalyticsTableOK {
	return &GetInventoryAPIV2AnalyticsTableOK{}
}

/*
GetInventoryAPIV2AnalyticsTableOK describes a response with status code 200, with default header values.

OK
*/
type GetInventoryAPIV2AnalyticsTableOK struct {
	Payload []*models.GithubComKaytuIoKaytuEnginePkgInventoryAPIAssetTableRow
}

// IsSuccess returns true when this get inventory Api v2 analytics table o k response has a 2xx status code
func (o *GetInventoryAPIV2AnalyticsTableOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get inventory Api v2 analytics table o k response has a 3xx status code
func (o *GetInventoryAPIV2AnalyticsTableOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get inventory Api v2 analytics table o k response has a 4xx status code
func (o *GetInventoryAPIV2AnalyticsTableOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get inventory Api v2 analytics table o k response has a 5xx status code
func (o *GetInventoryAPIV2AnalyticsTableOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get inventory Api v2 analytics table o k response a status code equal to that given
func (o *GetInventoryAPIV2AnalyticsTableOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get inventory Api v2 analytics table o k response
func (o *GetInventoryAPIV2AnalyticsTableOK) Code() int {
	return 200
}

func (o *GetInventoryAPIV2AnalyticsTableOK) Error() string {
	return fmt.Sprintf("[GET /inventory/api/v2/analytics/table][%d] getInventoryApiV2AnalyticsTableOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2AnalyticsTableOK) String() string {
	return fmt.Sprintf("[GET /inventory/api/v2/analytics/table][%d] getInventoryApiV2AnalyticsTableOK  %+v", 200, o.Payload)
}

func (o *GetInventoryAPIV2AnalyticsTableOK) GetPayload() []*models.GithubComKaytuIoKaytuEnginePkgInventoryAPIAssetTableRow {
	return o.Payload
}

func (o *GetInventoryAPIV2AnalyticsTableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
