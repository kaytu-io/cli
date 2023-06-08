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

// GetComplianceAPIV1MetadataInsightReader is a Reader for the GetComplianceAPIV1MetadataInsight structure.
type GetComplianceAPIV1MetadataInsightReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComplianceAPIV1MetadataInsightReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetComplianceAPIV1MetadataInsightOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetComplianceAPIV1MetadataInsightOK creates a GetComplianceAPIV1MetadataInsightOK with default headers values
func NewGetComplianceAPIV1MetadataInsightOK() *GetComplianceAPIV1MetadataInsightOK {
	return &GetComplianceAPIV1MetadataInsightOK{}
}

/*
GetComplianceAPIV1MetadataInsightOK describes a response with status code 200, with default header values.

OK
*/
type GetComplianceAPIV1MetadataInsightOK struct {
	Payload []*models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIInsight
}

// IsSuccess returns true when this get compliance Api v1 metadata insight o k response has a 2xx status code
func (o *GetComplianceAPIV1MetadataInsightOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get compliance Api v1 metadata insight o k response has a 3xx status code
func (o *GetComplianceAPIV1MetadataInsightOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get compliance Api v1 metadata insight o k response has a 4xx status code
func (o *GetComplianceAPIV1MetadataInsightOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get compliance Api v1 metadata insight o k response has a 5xx status code
func (o *GetComplianceAPIV1MetadataInsightOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get compliance Api v1 metadata insight o k response a status code equal to that given
func (o *GetComplianceAPIV1MetadataInsightOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get compliance Api v1 metadata insight o k response
func (o *GetComplianceAPIV1MetadataInsightOK) Code() int {
	return 200
}

func (o *GetComplianceAPIV1MetadataInsightOK) Error() string {
	return fmt.Sprintf("[GET /compliance/api/v1/metadata/insight][%d] getComplianceApiV1MetadataInsightOK  %+v", 200, o.Payload)
}

func (o *GetComplianceAPIV1MetadataInsightOK) String() string {
	return fmt.Sprintf("[GET /compliance/api/v1/metadata/insight][%d] getComplianceApiV1MetadataInsightOK  %+v", 200, o.Payload)
}

func (o *GetComplianceAPIV1MetadataInsightOK) GetPayload() []*models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIInsight {
	return o.Payload
}

func (o *GetComplianceAPIV1MetadataInsightOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
