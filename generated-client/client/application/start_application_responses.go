// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// StartApplicationReader is a Reader for the StartApplication structure.
type StartApplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartApplicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartApplicationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewStartApplicationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStartApplicationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStartApplicationOK creates a StartApplicationOK with default headers values
func NewStartApplicationOK() *StartApplicationOK {
	return &StartApplicationOK{}
}

/* StartApplicationOK describes a response with status code 200, with default header values.

Application started ok
*/
type StartApplicationOK struct {
}

func (o *StartApplicationOK) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/start][%d] startApplicationOK ", 200)
}

func (o *StartApplicationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStartApplicationUnauthorized creates a StartApplicationUnauthorized with default headers values
func NewStartApplicationUnauthorized() *StartApplicationUnauthorized {
	return &StartApplicationUnauthorized{}
}

/* StartApplicationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StartApplicationUnauthorized struct {
}

func (o *StartApplicationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/start][%d] startApplicationUnauthorized ", 401)
}

func (o *StartApplicationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStartApplicationNotFound creates a StartApplicationNotFound with default headers values
func NewStartApplicationNotFound() *StartApplicationNotFound {
	return &StartApplicationNotFound{}
}

/* StartApplicationNotFound describes a response with status code 404, with default header values.

Not found
*/
type StartApplicationNotFound struct {
}

func (o *StartApplicationNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/start][%d] startApplicationNotFound ", 404)
}

func (o *StartApplicationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
