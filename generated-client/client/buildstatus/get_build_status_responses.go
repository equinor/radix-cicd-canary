// Code generated by go-swagger; DO NOT EDIT.

package buildstatus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetBuildStatusReader is a Reader for the GetBuildStatus structure.
type GetBuildStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBuildStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBuildStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetBuildStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBuildStatusOK creates a GetBuildStatusOK with default headers values
func NewGetBuildStatusOK() *GetBuildStatusOK {
	return &GetBuildStatusOK{}
}

/* GetBuildStatusOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetBuildStatusOK struct {
}

func (o *GetBuildStatusOK) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/buildstatus][%d] getBuildStatusOK ", 200)
}

func (o *GetBuildStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBuildStatusNotFound creates a GetBuildStatusNotFound with default headers values
func NewGetBuildStatusNotFound() *GetBuildStatusNotFound {
	return &GetBuildStatusNotFound{}
}

/* GetBuildStatusNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetBuildStatusNotFound struct {
}

func (o *GetBuildStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/buildstatus][%d] getBuildStatusNotFound ", 404)
}

func (o *GetBuildStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
