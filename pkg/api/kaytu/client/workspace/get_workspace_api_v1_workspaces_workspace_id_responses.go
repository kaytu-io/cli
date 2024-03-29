// Code generated by go-swagger; DO NOT EDIT.

package workspace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetWorkspaceAPIV1WorkspacesWorkspaceIDReader is a Reader for the GetWorkspaceAPIV1WorkspacesWorkspaceID structure.
type GetWorkspaceAPIV1WorkspacesWorkspaceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkspaceAPIV1WorkspacesWorkspaceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkspaceAPIV1WorkspacesWorkspaceIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /workspace/api/v1/workspaces/{workspace_id}] GetWorkspaceAPIV1WorkspacesWorkspaceID", response, response.Code())
	}
}

// NewGetWorkspaceAPIV1WorkspacesWorkspaceIDOK creates a GetWorkspaceAPIV1WorkspacesWorkspaceIDOK with default headers values
func NewGetWorkspaceAPIV1WorkspacesWorkspaceIDOK() *GetWorkspaceAPIV1WorkspacesWorkspaceIDOK {
	return &GetWorkspaceAPIV1WorkspacesWorkspaceIDOK{}
}

/*
GetWorkspaceAPIV1WorkspacesWorkspaceIDOK describes a response with status code 200, with default header values.

OK
*/
type GetWorkspaceAPIV1WorkspacesWorkspaceIDOK struct {
}

// IsSuccess returns true when this get workspace Api v1 workspaces workspace Id o k response has a 2xx status code
func (o *GetWorkspaceAPIV1WorkspacesWorkspaceIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workspace Api v1 workspaces workspace Id o k response has a 3xx status code
func (o *GetWorkspaceAPIV1WorkspacesWorkspaceIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workspace Api v1 workspaces workspace Id o k response has a 4xx status code
func (o *GetWorkspaceAPIV1WorkspacesWorkspaceIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workspace Api v1 workspaces workspace Id o k response has a 5xx status code
func (o *GetWorkspaceAPIV1WorkspacesWorkspaceIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workspace Api v1 workspaces workspace Id o k response a status code equal to that given
func (o *GetWorkspaceAPIV1WorkspacesWorkspaceIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get workspace Api v1 workspaces workspace Id o k response
func (o *GetWorkspaceAPIV1WorkspacesWorkspaceIDOK) Code() int {
	return 200
}

func (o *GetWorkspaceAPIV1WorkspacesWorkspaceIDOK) Error() string {
	return fmt.Sprintf("[GET /workspace/api/v1/workspaces/{workspace_id}][%d] getWorkspaceApiV1WorkspacesWorkspaceIdOK ", 200)
}

func (o *GetWorkspaceAPIV1WorkspacesWorkspaceIDOK) String() string {
	return fmt.Sprintf("[GET /workspace/api/v1/workspaces/{workspace_id}][%d] getWorkspaceApiV1WorkspacesWorkspaceIdOK ", 200)
}

func (o *GetWorkspaceAPIV1WorkspacesWorkspaceIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
