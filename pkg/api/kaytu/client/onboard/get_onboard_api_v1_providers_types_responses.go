// Code generated by go-swagger; DO NOT EDIT.

package onboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

// GetOnboardAPIV1ProvidersTypesReader is a Reader for the GetOnboardAPIV1ProvidersTypes structure.
type GetOnboardAPIV1ProvidersTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOnboardAPIV1ProvidersTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOnboardAPIV1ProvidersTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOnboardAPIV1ProvidersTypesOK creates a GetOnboardAPIV1ProvidersTypesOK with default headers values
func NewGetOnboardAPIV1ProvidersTypesOK() *GetOnboardAPIV1ProvidersTypesOK {
	return &GetOnboardAPIV1ProvidersTypesOK{}
}

/*
GetOnboardAPIV1ProvidersTypesOK describes a response with status code 200, with default header values.

OK
*/
type GetOnboardAPIV1ProvidersTypesOK struct {
	Payload []*models.GitlabComKeibiengineKeibiEnginePkgOnboardAPIProviderType
}

// IsSuccess returns true when this get onboard Api v1 providers types o k response has a 2xx status code
func (o *GetOnboardAPIV1ProvidersTypesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get onboard Api v1 providers types o k response has a 3xx status code
func (o *GetOnboardAPIV1ProvidersTypesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get onboard Api v1 providers types o k response has a 4xx status code
func (o *GetOnboardAPIV1ProvidersTypesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get onboard Api v1 providers types o k response has a 5xx status code
func (o *GetOnboardAPIV1ProvidersTypesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get onboard Api v1 providers types o k response a status code equal to that given
func (o *GetOnboardAPIV1ProvidersTypesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get onboard Api v1 providers types o k response
func (o *GetOnboardAPIV1ProvidersTypesOK) Code() int {
	return 200
}

func (o *GetOnboardAPIV1ProvidersTypesOK) Error() string {
	return fmt.Sprintf("[GET /onboard/api/v1/providers/types][%d] getOnboardApiV1ProvidersTypesOK  %+v", 200, o.Payload)
}

func (o *GetOnboardAPIV1ProvidersTypesOK) String() string {
	return fmt.Sprintf("[GET /onboard/api/v1/providers/types][%d] getOnboardApiV1ProvidersTypesOK  %+v", 200, o.Payload)
}

func (o *GetOnboardAPIV1ProvidersTypesOK) GetPayload() []*models.GitlabComKeibiengineKeibiEnginePkgOnboardAPIProviderType {
	return o.Payload
}

func (o *GetOnboardAPIV1ProvidersTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
