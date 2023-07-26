// Code generated by go-swagger; DO NOT EDIT.

package environment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RestartEnvironmentReader is a Reader for the RestartEnvironment structure.
type RestartEnvironmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestartEnvironmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestartEnvironmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRestartEnvironmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRestartEnvironmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRestartEnvironmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRestartEnvironmentOK creates a RestartEnvironmentOK with default headers values
func NewRestartEnvironmentOK() *RestartEnvironmentOK {
	return &RestartEnvironmentOK{}
}

/* RestartEnvironmentOK describes a response with status code 200, with default header values.

Environment started ok
*/
type RestartEnvironmentOK struct {
}

func (o *RestartEnvironmentOK) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/restart][%d] restartEnvironmentOK ", 200)
}

func (o *RestartEnvironmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRestartEnvironmentUnauthorized creates a RestartEnvironmentUnauthorized with default headers values
func NewRestartEnvironmentUnauthorized() *RestartEnvironmentUnauthorized {
	return &RestartEnvironmentUnauthorized{}
}

/* RestartEnvironmentUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RestartEnvironmentUnauthorized struct {
}

func (o *RestartEnvironmentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/restart][%d] restartEnvironmentUnauthorized ", 401)
}

func (o *RestartEnvironmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRestartEnvironmentForbidden creates a RestartEnvironmentForbidden with default headers values
func NewRestartEnvironmentForbidden() *RestartEnvironmentForbidden {
	return &RestartEnvironmentForbidden{}
}

/* RestartEnvironmentForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RestartEnvironmentForbidden struct {
}

func (o *RestartEnvironmentForbidden) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/restart][%d] restartEnvironmentForbidden ", 403)
}

func (o *RestartEnvironmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRestartEnvironmentNotFound creates a RestartEnvironmentNotFound with default headers values
func NewRestartEnvironmentNotFound() *RestartEnvironmentNotFound {
	return &RestartEnvironmentNotFound{}
}

/* RestartEnvironmentNotFound describes a response with status code 404, with default header values.

Not found
*/
type RestartEnvironmentNotFound struct {
}

func (o *RestartEnvironmentNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/restart][%d] restartEnvironmentNotFound ", 404)
}

func (o *RestartEnvironmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
