// Code generated by go-swagger; DO NOT EDIT.

package cost_estimator

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetCostEstimatorAPIV1CostAwsReader is a Reader for the GetCostEstimatorAPIV1CostAws structure.
type GetCostEstimatorAPIV1CostAwsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCostEstimatorAPIV1CostAwsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCostEstimatorAPIV1CostAwsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /cost_estimator/api/v1/cost/aws] GetCostEstimatorAPIV1CostAws", response, response.Code())
	}
}

// NewGetCostEstimatorAPIV1CostAwsOK creates a GetCostEstimatorAPIV1CostAwsOK with default headers values
func NewGetCostEstimatorAPIV1CostAwsOK() *GetCostEstimatorAPIV1CostAwsOK {
	return &GetCostEstimatorAPIV1CostAwsOK{}
}

/*
GetCostEstimatorAPIV1CostAwsOK describes a response with status code 200, with default header values.

OK
*/
type GetCostEstimatorAPIV1CostAwsOK struct {
	Payload int64
}

// IsSuccess returns true when this get cost estimator Api v1 cost aws o k response has a 2xx status code
func (o *GetCostEstimatorAPIV1CostAwsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cost estimator Api v1 cost aws o k response has a 3xx status code
func (o *GetCostEstimatorAPIV1CostAwsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cost estimator Api v1 cost aws o k response has a 4xx status code
func (o *GetCostEstimatorAPIV1CostAwsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cost estimator Api v1 cost aws o k response has a 5xx status code
func (o *GetCostEstimatorAPIV1CostAwsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cost estimator Api v1 cost aws o k response a status code equal to that given
func (o *GetCostEstimatorAPIV1CostAwsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get cost estimator Api v1 cost aws o k response
func (o *GetCostEstimatorAPIV1CostAwsOK) Code() int {
	return 200
}

func (o *GetCostEstimatorAPIV1CostAwsOK) Error() string {
	return fmt.Sprintf("[GET /cost_estimator/api/v1/cost/aws][%d] getCostEstimatorApiV1CostAwsOK  %+v", 200, o.Payload)
}

func (o *GetCostEstimatorAPIV1CostAwsOK) String() string {
	return fmt.Sprintf("[GET /cost_estimator/api/v1/cost/aws][%d] getCostEstimatorApiV1CostAwsOK  %+v", 200, o.Payload)
}

func (o *GetCostEstimatorAPIV1CostAwsOK) GetPayload() int64 {
	return o.Payload
}

func (o *GetCostEstimatorAPIV1CostAwsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}