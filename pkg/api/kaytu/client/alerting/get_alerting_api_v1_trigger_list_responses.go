// Code generated by go-swagger; DO NOT EDIT.

package alerting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

// GetAlertingAPIV1TriggerListReader is a Reader for the GetAlertingAPIV1TriggerList structure.
type GetAlertingAPIV1TriggerListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertingAPIV1TriggerListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertingAPIV1TriggerListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /alerting/api/v1/trigger/list] GetAlertingAPIV1TriggerList", response, response.Code())
	}
}

// NewGetAlertingAPIV1TriggerListOK creates a GetAlertingAPIV1TriggerListOK with default headers values
func NewGetAlertingAPIV1TriggerListOK() *GetAlertingAPIV1TriggerListOK {
	return &GetAlertingAPIV1TriggerListOK{}
}

/*
GetAlertingAPIV1TriggerListOK describes a response with status code 200, with default header values.

OK
*/
type GetAlertingAPIV1TriggerListOK struct {
	Payload []*models.GithubComKaytuIoKaytuEnginePkgAlertingAPITriggers
}

// IsSuccess returns true when this get alerting Api v1 trigger list o k response has a 2xx status code
func (o *GetAlertingAPIV1TriggerListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get alerting Api v1 trigger list o k response has a 3xx status code
func (o *GetAlertingAPIV1TriggerListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting Api v1 trigger list o k response has a 4xx status code
func (o *GetAlertingAPIV1TriggerListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alerting Api v1 trigger list o k response has a 5xx status code
func (o *GetAlertingAPIV1TriggerListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting Api v1 trigger list o k response a status code equal to that given
func (o *GetAlertingAPIV1TriggerListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get alerting Api v1 trigger list o k response
func (o *GetAlertingAPIV1TriggerListOK) Code() int {
	return 200
}

func (o *GetAlertingAPIV1TriggerListOK) Error() string {
	return fmt.Sprintf("[GET /alerting/api/v1/trigger/list][%d] getAlertingApiV1TriggerListOK  %+v", 200, o.Payload)
}

func (o *GetAlertingAPIV1TriggerListOK) String() string {
	return fmt.Sprintf("[GET /alerting/api/v1/trigger/list][%d] getAlertingApiV1TriggerListOK  %+v", 200, o.Payload)
}

func (o *GetAlertingAPIV1TriggerListOK) GetPayload() []*models.GithubComKaytuIoKaytuEnginePkgAlertingAPITriggers {
	return o.Payload
}

func (o *GetAlertingAPIV1TriggerListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
