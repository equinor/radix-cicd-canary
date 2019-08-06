// Code generated by go-swagger; DO NOT EDIT.

package environment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/equinor/radix-cicd-canary/generated-client/models"
)

// GetEnvironmentReader is a Reader for the GetEnvironment structure.
type GetEnvironmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEnvironmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEnvironmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetEnvironmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetEnvironmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEnvironmentOK creates a GetEnvironmentOK with default headers values
func NewGetEnvironmentOK() *GetEnvironmentOK {
	return &GetEnvironmentOK{}
}

/*GetEnvironmentOK handles this case with default header values.

Successful get environment
*/
type GetEnvironmentOK struct {
	Payload *models.Environment
}

func (o *GetEnvironmentOK) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}][%d] getEnvironmentOK  %+v", 200, o.Payload)
}

func (o *GetEnvironmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Environment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEnvironmentUnauthorized creates a GetEnvironmentUnauthorized with default headers values
func NewGetEnvironmentUnauthorized() *GetEnvironmentUnauthorized {
	return &GetEnvironmentUnauthorized{}
}

/*GetEnvironmentUnauthorized handles this case with default header values.

Unauthorized
*/
type GetEnvironmentUnauthorized struct {
}

func (o *GetEnvironmentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}][%d] getEnvironmentUnauthorized ", 401)
}

func (o *GetEnvironmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEnvironmentNotFound creates a GetEnvironmentNotFound with default headers values
func NewGetEnvironmentNotFound() *GetEnvironmentNotFound {
	return &GetEnvironmentNotFound{}
}

/*GetEnvironmentNotFound handles this case with default header values.

Not found
*/
type GetEnvironmentNotFound struct {
}

func (o *GetEnvironmentNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}][%d] getEnvironmentNotFound ", 404)
}

func (o *GetEnvironmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
