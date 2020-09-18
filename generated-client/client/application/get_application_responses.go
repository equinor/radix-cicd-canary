// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/equinor/radix-cicd-canary/generated-client/models"
)

// GetApplicationReader is a Reader for the GetApplication structure.
type GetApplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetApplicationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetApplicationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApplicationOK creates a GetApplicationOK with default headers values
func NewGetApplicationOK() *GetApplicationOK {
	return &GetApplicationOK{}
}

/*GetApplicationOK handles this case with default header values.

Successful get application
*/
type GetApplicationOK struct {
	Payload *models.Application
}

func (o *GetApplicationOK) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}][%d] getApplicationOK  %+v", 200, o.Payload)
}

func (o *GetApplicationOK) GetPayload() *models.Application {
	return o.Payload
}

func (o *GetApplicationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Application)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationUnauthorized creates a GetApplicationUnauthorized with default headers values
func NewGetApplicationUnauthorized() *GetApplicationUnauthorized {
	return &GetApplicationUnauthorized{}
}

/*GetApplicationUnauthorized handles this case with default header values.

Unauthorized
*/
type GetApplicationUnauthorized struct {
}

func (o *GetApplicationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}][%d] getApplicationUnauthorized ", 401)
}

func (o *GetApplicationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationNotFound creates a GetApplicationNotFound with default headers values
func NewGetApplicationNotFound() *GetApplicationNotFound {
	return &GetApplicationNotFound{}
}

/*GetApplicationNotFound handles this case with default header values.

Not found
*/
type GetApplicationNotFound struct {
}

func (o *GetApplicationNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}][%d] getApplicationNotFound ", 404)
}

func (o *GetApplicationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
