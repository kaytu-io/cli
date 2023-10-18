// Code generated by go-swagger; DO NOT EDIT.

package alerting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetAlertingAPIV1RuleRuleIDTriggerReader is a Reader for the GetAlertingAPIV1RuleRuleIDTrigger structure.
type GetAlertingAPIV1RuleRuleIDTriggerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertingAPIV1RuleRuleIDTriggerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertingAPIV1RuleRuleIDTriggerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /alerting/api/v1/rule/{ruleId}/trigger] GetAlertingAPIV1RuleRuleIDTrigger", response, response.Code())
	}
}

// NewGetAlertingAPIV1RuleRuleIDTriggerOK creates a GetAlertingAPIV1RuleRuleIDTriggerOK with default headers values
func NewGetAlertingAPIV1RuleRuleIDTriggerOK() *GetAlertingAPIV1RuleRuleIDTriggerOK {
	return &GetAlertingAPIV1RuleRuleIDTriggerOK{}
}

/*
GetAlertingAPIV1RuleRuleIDTriggerOK describes a response with status code 200, with default header values.

OK
*/
type GetAlertingAPIV1RuleRuleIDTriggerOK struct {
	Payload string
}

// IsSuccess returns true when this get alerting Api v1 rule rule Id trigger o k response has a 2xx status code
func (o *GetAlertingAPIV1RuleRuleIDTriggerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get alerting Api v1 rule rule Id trigger o k response has a 3xx status code
func (o *GetAlertingAPIV1RuleRuleIDTriggerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting Api v1 rule rule Id trigger o k response has a 4xx status code
func (o *GetAlertingAPIV1RuleRuleIDTriggerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alerting Api v1 rule rule Id trigger o k response has a 5xx status code
func (o *GetAlertingAPIV1RuleRuleIDTriggerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting Api v1 rule rule Id trigger o k response a status code equal to that given
func (o *GetAlertingAPIV1RuleRuleIDTriggerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get alerting Api v1 rule rule Id trigger o k response
func (o *GetAlertingAPIV1RuleRuleIDTriggerOK) Code() int {
	return 200
}

func (o *GetAlertingAPIV1RuleRuleIDTriggerOK) Error() string {
	return fmt.Sprintf("[GET /alerting/api/v1/rule/{ruleId}/trigger][%d] getAlertingApiV1RuleRuleIdTriggerOK  %+v", 200, o.Payload)
}

func (o *GetAlertingAPIV1RuleRuleIDTriggerOK) String() string {
	return fmt.Sprintf("[GET /alerting/api/v1/rule/{ruleId}/trigger][%d] getAlertingApiV1RuleRuleIdTriggerOK  %+v", 200, o.Payload)
}

func (o *GetAlertingAPIV1RuleRuleIDTriggerOK) GetPayload() string {
	return o.Payload
}

func (o *GetAlertingAPIV1RuleRuleIDTriggerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}