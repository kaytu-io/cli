// Code generated by go-swagger; DO NOT EDIT.

package schedule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

// GetScheduleAPIV1ResourceTypeProviderReader is a Reader for the GetScheduleAPIV1ResourceTypeProvider structure.
type GetScheduleAPIV1ResourceTypeProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScheduleAPIV1ResourceTypeProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScheduleAPIV1ResourceTypeProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /schedule/api/v1/resource_type/{provider}] GetScheduleAPIV1ResourceTypeProvider", response, response.Code())
	}
}

// NewGetScheduleAPIV1ResourceTypeProviderOK creates a GetScheduleAPIV1ResourceTypeProviderOK with default headers values
func NewGetScheduleAPIV1ResourceTypeProviderOK() *GetScheduleAPIV1ResourceTypeProviderOK {
	return &GetScheduleAPIV1ResourceTypeProviderOK{}
}

/*
GetScheduleAPIV1ResourceTypeProviderOK describes a response with status code 200, with default header values.

OK
*/
type GetScheduleAPIV1ResourceTypeProviderOK struct {
	Payload []*models.GitlabComKeibiengineKeibiEnginePkgDescribeAPIResourceTypeDetail
}

// IsSuccess returns true when this get schedule Api v1 resource type provider o k response has a 2xx status code
func (o *GetScheduleAPIV1ResourceTypeProviderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get schedule Api v1 resource type provider o k response has a 3xx status code
func (o *GetScheduleAPIV1ResourceTypeProviderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get schedule Api v1 resource type provider o k response has a 4xx status code
func (o *GetScheduleAPIV1ResourceTypeProviderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get schedule Api v1 resource type provider o k response has a 5xx status code
func (o *GetScheduleAPIV1ResourceTypeProviderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get schedule Api v1 resource type provider o k response a status code equal to that given
func (o *GetScheduleAPIV1ResourceTypeProviderOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get schedule Api v1 resource type provider o k response
func (o *GetScheduleAPIV1ResourceTypeProviderOK) Code() int {
	return 200
}

func (o *GetScheduleAPIV1ResourceTypeProviderOK) Error() string {
	return fmt.Sprintf("[GET /schedule/api/v1/resource_type/{provider}][%d] getScheduleApiV1ResourceTypeProviderOK  %+v", 200, o.Payload)
}

func (o *GetScheduleAPIV1ResourceTypeProviderOK) String() string {
	return fmt.Sprintf("[GET /schedule/api/v1/resource_type/{provider}][%d] getScheduleApiV1ResourceTypeProviderOK  %+v", 200, o.Payload)
}

func (o *GetScheduleAPIV1ResourceTypeProviderOK) GetPayload() []*models.GitlabComKeibiengineKeibiEnginePkgDescribeAPIResourceTypeDetail {
	return o.Payload
}

func (o *GetScheduleAPIV1ResourceTypeProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
