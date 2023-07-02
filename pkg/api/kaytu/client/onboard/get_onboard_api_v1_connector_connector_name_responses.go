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

// GetOnboardAPIV1ConnectorConnectorNameReader is a Reader for the GetOnboardAPIV1ConnectorConnectorName structure.
type GetOnboardAPIV1ConnectorConnectorNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOnboardAPIV1ConnectorConnectorNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOnboardAPIV1ConnectorConnectorNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /onboard/api/v1/connector/{connectorName}] GetOnboardAPIV1ConnectorConnectorName", response, response.Code())
	}
}

// NewGetOnboardAPIV1ConnectorConnectorNameOK creates a GetOnboardAPIV1ConnectorConnectorNameOK with default headers values
func NewGetOnboardAPIV1ConnectorConnectorNameOK() *GetOnboardAPIV1ConnectorConnectorNameOK {
	return &GetOnboardAPIV1ConnectorConnectorNameOK{}
}

/*
GetOnboardAPIV1ConnectorConnectorNameOK describes a response with status code 200, with default header values.

OK
*/
type GetOnboardAPIV1ConnectorConnectorNameOK struct {
	Payload *models.GithubComKaytuIoKaytuEnginePkgOnboardAPIConnector
}

// IsSuccess returns true when this get onboard Api v1 connector connector name o k response has a 2xx status code
func (o *GetOnboardAPIV1ConnectorConnectorNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get onboard Api v1 connector connector name o k response has a 3xx status code
func (o *GetOnboardAPIV1ConnectorConnectorNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get onboard Api v1 connector connector name o k response has a 4xx status code
func (o *GetOnboardAPIV1ConnectorConnectorNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get onboard Api v1 connector connector name o k response has a 5xx status code
func (o *GetOnboardAPIV1ConnectorConnectorNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get onboard Api v1 connector connector name o k response a status code equal to that given
func (o *GetOnboardAPIV1ConnectorConnectorNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get onboard Api v1 connector connector name o k response
func (o *GetOnboardAPIV1ConnectorConnectorNameOK) Code() int {
	return 200
}

func (o *GetOnboardAPIV1ConnectorConnectorNameOK) Error() string {
	return fmt.Sprintf("[GET /onboard/api/v1/connector/{connectorName}][%d] getOnboardApiV1ConnectorConnectorNameOK  %+v", 200, o.Payload)
}

func (o *GetOnboardAPIV1ConnectorConnectorNameOK) String() string {
	return fmt.Sprintf("[GET /onboard/api/v1/connector/{connectorName}][%d] getOnboardApiV1ConnectorConnectorNameOK  %+v", 200, o.Payload)
}

func (o *GetOnboardAPIV1ConnectorConnectorNameOK) GetPayload() *models.GithubComKaytuIoKaytuEnginePkgOnboardAPIConnector {
	return o.Payload
}

func (o *GetOnboardAPIV1ConnectorConnectorNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GithubComKaytuIoKaytuEnginePkgOnboardAPIConnector)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}