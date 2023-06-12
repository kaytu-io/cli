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

// GetOnboardAPIV1ProvidersReader is a Reader for the GetOnboardAPIV1Providers structure.
type GetOnboardAPIV1ProvidersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOnboardAPIV1ProvidersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOnboardAPIV1ProvidersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /onboard/api/v1/providers] GetOnboardAPIV1Providers", response, response.Code())
	}
}

// NewGetOnboardAPIV1ProvidersOK creates a GetOnboardAPIV1ProvidersOK with default headers values
func NewGetOnboardAPIV1ProvidersOK() *GetOnboardAPIV1ProvidersOK {
	return &GetOnboardAPIV1ProvidersOK{}
}

/*
GetOnboardAPIV1ProvidersOK describes a response with status code 200, with default header values.

OK
*/
type GetOnboardAPIV1ProvidersOK struct {
	Payload []*models.GitlabComKeibiengineKeibiEnginePkgOnboardAPIProvider
}

// IsSuccess returns true when this get onboard Api v1 providers o k response has a 2xx status code
func (o *GetOnboardAPIV1ProvidersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get onboard Api v1 providers o k response has a 3xx status code
func (o *GetOnboardAPIV1ProvidersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get onboard Api v1 providers o k response has a 4xx status code
func (o *GetOnboardAPIV1ProvidersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get onboard Api v1 providers o k response has a 5xx status code
func (o *GetOnboardAPIV1ProvidersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get onboard Api v1 providers o k response a status code equal to that given
func (o *GetOnboardAPIV1ProvidersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get onboard Api v1 providers o k response
func (o *GetOnboardAPIV1ProvidersOK) Code() int {
	return 200
}

func (o *GetOnboardAPIV1ProvidersOK) Error() string {
	return fmt.Sprintf("[GET /onboard/api/v1/providers][%d] getOnboardApiV1ProvidersOK  %+v", 200, o.Payload)
}

func (o *GetOnboardAPIV1ProvidersOK) String() string {
	return fmt.Sprintf("[GET /onboard/api/v1/providers][%d] getOnboardApiV1ProvidersOK  %+v", 200, o.Payload)
}

func (o *GetOnboardAPIV1ProvidersOK) GetPayload() []*models.GitlabComKeibiengineKeibiEnginePkgOnboardAPIProvider {
	return o.Payload
}

func (o *GetOnboardAPIV1ProvidersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
