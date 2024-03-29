// Code generated by go-swagger; DO NOT EDIT.

package workspace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetWorkspaceAPIV1WorkspacesBynameWorkspaceNameReader is a Reader for the GetWorkspaceAPIV1WorkspacesBynameWorkspaceName structure.
type GetWorkspaceAPIV1WorkspacesBynameWorkspaceNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkspaceAPIV1WorkspacesBynameWorkspaceNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkspaceAPIV1WorkspacesBynameWorkspaceNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /workspace/api/v1/workspaces/byname/{workspace_name}] GetWorkspaceAPIV1WorkspacesBynameWorkspaceName", response, response.Code())
	}
}

// NewGetWorkspaceAPIV1WorkspacesBynameWorkspaceNameOK creates a GetWorkspaceAPIV1WorkspacesBynameWorkspaceNameOK with default headers values
func NewGetWorkspaceAPIV1WorkspacesBynameWorkspaceNameOK() *GetWorkspaceAPIV1WorkspacesBynameWorkspaceNameOK {
	return &GetWorkspaceAPIV1WorkspacesBynameWorkspaceNameOK{}
}

/*
GetWorkspaceAPIV1WorkspacesBynameWorkspaceNameOK describes a response with status code 200, with default header values.

OK
*/
type GetWorkspaceAPIV1WorkspacesBynameWorkspaceNameOK struct {
}

// IsSuccess returns true when this get workspace Api v1 workspaces byname workspace name o k response has a 2xx status code
func (o *GetWorkspaceAPIV1WorkspacesBynameWorkspaceNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workspace Api v1 workspaces byname workspace name o k response has a 3xx status code
func (o *GetWorkspaceAPIV1WorkspacesBynameWorkspaceNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workspace Api v1 workspaces byname workspace name o k response has a 4xx status code
func (o *GetWorkspaceAPIV1WorkspacesBynameWorkspaceNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workspace Api v1 workspaces byname workspace name o k response has a 5xx status code
func (o *GetWorkspaceAPIV1WorkspacesBynameWorkspaceNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workspace Api v1 workspaces byname workspace name o k response a status code equal to that given
func (o *GetWorkspaceAPIV1WorkspacesBynameWorkspaceNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get workspace Api v1 workspaces byname workspace name o k response
func (o *GetWorkspaceAPIV1WorkspacesBynameWorkspaceNameOK) Code() int {
	return 200
}

func (o *GetWorkspaceAPIV1WorkspacesBynameWorkspaceNameOK) Error() string {
	return fmt.Sprintf("[GET /workspace/api/v1/workspaces/byname/{workspace_name}][%d] getWorkspaceApiV1WorkspacesBynameWorkspaceNameOK ", 200)
}

func (o *GetWorkspaceAPIV1WorkspacesBynameWorkspaceNameOK) String() string {
	return fmt.Sprintf("[GET /workspace/api/v1/workspaces/byname/{workspace_name}][%d] getWorkspaceApiV1WorkspacesBynameWorkspaceNameOK ", 200)
}

func (o *GetWorkspaceAPIV1WorkspacesBynameWorkspaceNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
