// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteApplicationReader is a Reader for the DeleteApplication structure.
type DeleteApplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteApplicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteApplicationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteApplicationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteApplicationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteApplicationOK creates a DeleteApplicationOK with default headers values
func NewDeleteApplicationOK() *DeleteApplicationOK {
	return &DeleteApplicationOK{}
}

/*DeleteApplicationOK handles this case with default header values.

Application deleted ok
*/
type DeleteApplicationOK struct {
}

func (o *DeleteApplicationOK) Error() string {
	return fmt.Sprintf("[DELETE /applications/{appName}][%d] deleteApplicationOK ", 200)
}

func (o *DeleteApplicationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApplicationUnauthorized creates a DeleteApplicationUnauthorized with default headers values
func NewDeleteApplicationUnauthorized() *DeleteApplicationUnauthorized {
	return &DeleteApplicationUnauthorized{}
}

/*DeleteApplicationUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteApplicationUnauthorized struct {
}

func (o *DeleteApplicationUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /applications/{appName}][%d] deleteApplicationUnauthorized ", 401)
}

func (o *DeleteApplicationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApplicationNotFound creates a DeleteApplicationNotFound with default headers values
func NewDeleteApplicationNotFound() *DeleteApplicationNotFound {
	return &DeleteApplicationNotFound{}
}

/*DeleteApplicationNotFound handles this case with default header values.

Not found
*/
type DeleteApplicationNotFound struct {
}

func (o *DeleteApplicationNotFound) Error() string {
	return fmt.Sprintf("[DELETE /applications/{appName}][%d] deleteApplicationNotFound ", 404)
}

func (o *DeleteApplicationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
