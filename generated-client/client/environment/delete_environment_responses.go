// Code generated by go-swagger; DO NOT EDIT.

package environment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteEnvironmentReader is a Reader for the DeleteEnvironment structure.
type DeleteEnvironmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEnvironmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteEnvironmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteEnvironmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteEnvironmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteEnvironmentOK creates a DeleteEnvironmentOK with default headers values
func NewDeleteEnvironmentOK() *DeleteEnvironmentOK {
	return &DeleteEnvironmentOK{}
}

/* DeleteEnvironmentOK describes a response with status code 200, with default header values.

Environment deleted ok
*/
type DeleteEnvironmentOK struct {
}

func (o *DeleteEnvironmentOK) Error() string {
	return fmt.Sprintf("[DELETE /applications/{appName}/environments/{envName}][%d] deleteEnvironmentOK ", 200)
}

func (o *DeleteEnvironmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEnvironmentUnauthorized creates a DeleteEnvironmentUnauthorized with default headers values
func NewDeleteEnvironmentUnauthorized() *DeleteEnvironmentUnauthorized {
	return &DeleteEnvironmentUnauthorized{}
}

/* DeleteEnvironmentUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteEnvironmentUnauthorized struct {
}

func (o *DeleteEnvironmentUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /applications/{appName}/environments/{envName}][%d] deleteEnvironmentUnauthorized ", 401)
}

func (o *DeleteEnvironmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEnvironmentNotFound creates a DeleteEnvironmentNotFound with default headers values
func NewDeleteEnvironmentNotFound() *DeleteEnvironmentNotFound {
	return &DeleteEnvironmentNotFound{}
}

/* DeleteEnvironmentNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteEnvironmentNotFound struct {
}

func (o *DeleteEnvironmentNotFound) Error() string {
	return fmt.Sprintf("[DELETE /applications/{appName}/environments/{envName}][%d] deleteEnvironmentNotFound ", 404)
}

func (o *DeleteEnvironmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
