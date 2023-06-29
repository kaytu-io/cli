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

// GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDReader is a Reader for the GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceID structure.
type GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /workspace/api/v1/workspaces/limits/byid/{workspace_id}] GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceID", response, response.Code())
	}
}

// NewGetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDOK creates a GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDOK with default headers values
func NewGetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDOK() *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDOK {
	return &GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDOK{}
}

/*
GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDOK describes a response with status code 200, with default header values.

OK
*/
type GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDOK struct {
	Payload *models.GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceLimits
}

// IsSuccess returns true when this get workspace Api v1 workspaces limits byid workspace Id o k response has a 2xx status code
func (o *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workspace Api v1 workspaces limits byid workspace Id o k response has a 3xx status code
func (o *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workspace Api v1 workspaces limits byid workspace Id o k response has a 4xx status code
func (o *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workspace Api v1 workspaces limits byid workspace Id o k response has a 5xx status code
func (o *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workspace Api v1 workspaces limits byid workspace Id o k response a status code equal to that given
func (o *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get workspace Api v1 workspaces limits byid workspace Id o k response
func (o *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDOK) Code() int {
	return 200
}

func (o *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDOK) Error() string {
	return fmt.Sprintf("[GET /workspace/api/v1/workspaces/limits/byid/{workspace_id}][%d] getWorkspaceApiV1WorkspacesLimitsByidWorkspaceIdOK  %+v", 200, o.Payload)
}

func (o *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDOK) String() string {
	return fmt.Sprintf("[GET /workspace/api/v1/workspaces/limits/byid/{workspace_id}][%d] getWorkspaceApiV1WorkspacesLimitsByidWorkspaceIdOK  %+v", 200, o.Payload)
}

func (o *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDOK) GetPayload() *models.GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceLimits {
	return o.Payload
}

func (o *GetWorkspaceAPIV1WorkspacesLimitsByidWorkspaceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GithubComKaytuIoKaytuEnginePkgWorkspaceAPIWorkspaceLimits)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
