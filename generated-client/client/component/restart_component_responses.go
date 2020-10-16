// Code generated by go-swagger; DO NOT EDIT.

package component

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RestartComponentReader is a Reader for the RestartComponent structure.
type RestartComponentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestartComponentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestartComponentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRestartComponentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRestartComponentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRestartComponentOK creates a RestartComponentOK with default headers values
func NewRestartComponentOK() *RestartComponentOK {
	return &RestartComponentOK{}
}

/*RestartComponentOK handles this case with default header values.

Component started ok
*/
type RestartComponentOK struct {
}

func (o *RestartComponentOK) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/components/{componentName}/restart][%d] restartComponentOK ", 200)
}

func (o *RestartComponentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRestartComponentUnauthorized creates a RestartComponentUnauthorized with default headers values
func NewRestartComponentUnauthorized() *RestartComponentUnauthorized {
	return &RestartComponentUnauthorized{}
}

/*RestartComponentUnauthorized handles this case with default header values.

Unauthorized
*/
type RestartComponentUnauthorized struct {
}

func (o *RestartComponentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/components/{componentName}/restart][%d] restartComponentUnauthorized ", 401)
}

func (o *RestartComponentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRestartComponentNotFound creates a RestartComponentNotFound with default headers values
func NewRestartComponentNotFound() *RestartComponentNotFound {
	return &RestartComponentNotFound{}
}

/*RestartComponentNotFound handles this case with default header values.

Not found
*/
type RestartComponentNotFound struct {
}

func (o *RestartComponentNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/components/{componentName}/restart][%d] restartComponentNotFound ", 404)
}

func (o *RestartComponentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
