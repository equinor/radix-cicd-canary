// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RestartApplicationReader is a Reader for the RestartApplication structure.
type RestartApplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestartApplicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestartApplicationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRestartApplicationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRestartApplicationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRestartApplicationOK creates a RestartApplicationOK with default headers values
func NewRestartApplicationOK() *RestartApplicationOK {
	return &RestartApplicationOK{}
}

/* RestartApplicationOK describes a response with status code 200, with default header values.

Application started ok
*/
type RestartApplicationOK struct {
}

func (o *RestartApplicationOK) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/restart][%d] restartApplicationOK ", 200)
}

func (o *RestartApplicationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRestartApplicationUnauthorized creates a RestartApplicationUnauthorized with default headers values
func NewRestartApplicationUnauthorized() *RestartApplicationUnauthorized {
	return &RestartApplicationUnauthorized{}
}

/* RestartApplicationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RestartApplicationUnauthorized struct {
}

func (o *RestartApplicationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/restart][%d] restartApplicationUnauthorized ", 401)
}

func (o *RestartApplicationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRestartApplicationNotFound creates a RestartApplicationNotFound with default headers values
func NewRestartApplicationNotFound() *RestartApplicationNotFound {
	return &RestartApplicationNotFound{}
}

/* RestartApplicationNotFound describes a response with status code 404, with default header values.

Not found
*/
type RestartApplicationNotFound struct {
}

func (o *RestartApplicationNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/restart][%d] restartApplicationNotFound ", 404)
}

func (o *RestartApplicationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
