// Code generated by go-swagger; DO NOT EDIT.

package workspace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

// GetWorkspaceAPIV1BootstrapWorkspaceNameReader is a Reader for the GetWorkspaceAPIV1BootstrapWorkspaceName structure.
type GetWorkspaceAPIV1BootstrapWorkspaceNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkspaceAPIV1BootstrapWorkspaceNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkspaceAPIV1BootstrapWorkspaceNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /workspace/api/v1/bootstrap/{workspace_name}] GetWorkspaceAPIV1BootstrapWorkspaceName", response, response.Code())
	}
}

// NewGetWorkspaceAPIV1BootstrapWorkspaceNameOK creates a GetWorkspaceAPIV1BootstrapWorkspaceNameOK with default headers values
func NewGetWorkspaceAPIV1BootstrapWorkspaceNameOK() *GetWorkspaceAPIV1BootstrapWorkspaceNameOK {
	return &GetWorkspaceAPIV1BootstrapWorkspaceNameOK{}
}

/*
GetWorkspaceAPIV1BootstrapWorkspaceNameOK describes a response with status code 200, with default header values.

OK
*/
type GetWorkspaceAPIV1BootstrapWorkspaceNameOK struct {
	Payload *models.GithubComKaytuIoKaytuEnginePkgWorkspaceAPIBootstrapStatusResponse
}

// IsSuccess returns true when this get workspace Api v1 bootstrap workspace name o k response has a 2xx status code
func (o *GetWorkspaceAPIV1BootstrapWorkspaceNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workspace Api v1 bootstrap workspace name o k response has a 3xx status code
func (o *GetWorkspaceAPIV1BootstrapWorkspaceNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workspace Api v1 bootstrap workspace name o k response has a 4xx status code
func (o *GetWorkspaceAPIV1BootstrapWorkspaceNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workspace Api v1 bootstrap workspace name o k response has a 5xx status code
func (o *GetWorkspaceAPIV1BootstrapWorkspaceNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workspace Api v1 bootstrap workspace name o k response a status code equal to that given
func (o *GetWorkspaceAPIV1BootstrapWorkspaceNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get workspace Api v1 bootstrap workspace name o k response
func (o *GetWorkspaceAPIV1BootstrapWorkspaceNameOK) Code() int {
	return 200
}

func (o *GetWorkspaceAPIV1BootstrapWorkspaceNameOK) Error() string {
	return fmt.Sprintf("[GET /workspace/api/v1/bootstrap/{workspace_name}][%d] getWorkspaceApiV1BootstrapWorkspaceNameOK  %+v", 200, o.Payload)
}

func (o *GetWorkspaceAPIV1BootstrapWorkspaceNameOK) String() string {
	return fmt.Sprintf("[GET /workspace/api/v1/bootstrap/{workspace_name}][%d] getWorkspaceApiV1BootstrapWorkspaceNameOK  %+v", 200, o.Payload)
}

func (o *GetWorkspaceAPIV1BootstrapWorkspaceNameOK) GetPayload() *models.GithubComKaytuIoKaytuEnginePkgWorkspaceAPIBootstrapStatusResponse {
	return o.Payload
}

func (o *GetWorkspaceAPIV1BootstrapWorkspaceNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GithubComKaytuIoKaytuEnginePkgWorkspaceAPIBootstrapStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
