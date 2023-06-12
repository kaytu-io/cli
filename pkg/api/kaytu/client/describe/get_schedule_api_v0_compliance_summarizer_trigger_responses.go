// Code generated by go-swagger; DO NOT EDIT.

package describe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetScheduleAPIV0ComplianceSummarizerTriggerReader is a Reader for the GetScheduleAPIV0ComplianceSummarizerTrigger structure.
type GetScheduleAPIV0ComplianceSummarizerTriggerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScheduleAPIV0ComplianceSummarizerTriggerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScheduleAPIV0ComplianceSummarizerTriggerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /schedule/api/v0/compliance/summarizer/trigger] GetScheduleAPIV0ComplianceSummarizerTrigger", response, response.Code())
	}
}

// NewGetScheduleAPIV0ComplianceSummarizerTriggerOK creates a GetScheduleAPIV0ComplianceSummarizerTriggerOK with default headers values
func NewGetScheduleAPIV0ComplianceSummarizerTriggerOK() *GetScheduleAPIV0ComplianceSummarizerTriggerOK {
	return &GetScheduleAPIV0ComplianceSummarizerTriggerOK{}
}

/*
GetScheduleAPIV0ComplianceSummarizerTriggerOK describes a response with status code 200, with default header values.

OK
*/
type GetScheduleAPIV0ComplianceSummarizerTriggerOK struct {
}

// IsSuccess returns true when this get schedule Api v0 compliance summarizer trigger o k response has a 2xx status code
func (o *GetScheduleAPIV0ComplianceSummarizerTriggerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get schedule Api v0 compliance summarizer trigger o k response has a 3xx status code
func (o *GetScheduleAPIV0ComplianceSummarizerTriggerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get schedule Api v0 compliance summarizer trigger o k response has a 4xx status code
func (o *GetScheduleAPIV0ComplianceSummarizerTriggerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get schedule Api v0 compliance summarizer trigger o k response has a 5xx status code
func (o *GetScheduleAPIV0ComplianceSummarizerTriggerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get schedule Api v0 compliance summarizer trigger o k response a status code equal to that given
func (o *GetScheduleAPIV0ComplianceSummarizerTriggerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get schedule Api v0 compliance summarizer trigger o k response
func (o *GetScheduleAPIV0ComplianceSummarizerTriggerOK) Code() int {
	return 200
}

func (o *GetScheduleAPIV0ComplianceSummarizerTriggerOK) Error() string {
	return fmt.Sprintf("[GET /schedule/api/v0/compliance/summarizer/trigger][%d] getScheduleApiV0ComplianceSummarizerTriggerOK ", 200)
}

func (o *GetScheduleAPIV0ComplianceSummarizerTriggerOK) String() string {
	return fmt.Sprintf("[GET /schedule/api/v0/compliance/summarizer/trigger][%d] getScheduleApiV0ComplianceSummarizerTriggerOK ", 200)
}

func (o *GetScheduleAPIV0ComplianceSummarizerTriggerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
