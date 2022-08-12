// Code generated by go-swagger; DO NOT EDIT.

package platform

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/equinor/radix-cicd-canary/generated-client/models"
)

// RegisterApplicationReader is a Reader for the RegisterApplication structure.
type RegisterApplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegisterApplicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRegisterApplicationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRegisterApplicationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRegisterApplicationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewRegisterApplicationConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRegisterApplicationOK creates a RegisterApplicationOK with default headers values
func NewRegisterApplicationOK() *RegisterApplicationOK {
	return &RegisterApplicationOK{}
}

/* RegisterApplicationOK describes a response with status code 200, with default header values.

Successful application registration
*/
type RegisterApplicationOK struct {
	Payload *models.ApplicationRegistration
}

func (o *RegisterApplicationOK) Error() string {
	return fmt.Sprintf("[POST /applications][%d] registerApplicationOK  %+v", 200, o.Payload)
}
func (o *RegisterApplicationOK) GetPayload() *models.ApplicationRegistration {
	return o.Payload
}

func (o *RegisterApplicationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationRegistration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterApplicationBadRequest creates a RegisterApplicationBadRequest with default headers values
func NewRegisterApplicationBadRequest() *RegisterApplicationBadRequest {
	return &RegisterApplicationBadRequest{}
}

/* RegisterApplicationBadRequest describes a response with status code 400, with default header values.

Invalid application registration
*/
type RegisterApplicationBadRequest struct {
}

func (o *RegisterApplicationBadRequest) Error() string {
	return fmt.Sprintf("[POST /applications][%d] registerApplicationBadRequest ", 400)
}

func (o *RegisterApplicationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegisterApplicationUnauthorized creates a RegisterApplicationUnauthorized with default headers values
func NewRegisterApplicationUnauthorized() *RegisterApplicationUnauthorized {
	return &RegisterApplicationUnauthorized{}
}

/* RegisterApplicationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RegisterApplicationUnauthorized struct {
}

func (o *RegisterApplicationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /applications][%d] registerApplicationUnauthorized ", 401)
}

func (o *RegisterApplicationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegisterApplicationConflict creates a RegisterApplicationConflict with default headers values
func NewRegisterApplicationConflict() *RegisterApplicationConflict {
	return &RegisterApplicationConflict{}
}

/* RegisterApplicationConflict describes a response with status code 409, with default header values.

Conflict
*/
type RegisterApplicationConflict struct {
}

func (o *RegisterApplicationConflict) Error() string {
	return fmt.Sprintf("[POST /applications][%d] registerApplicationConflict ", 409)
}

func (o *RegisterApplicationConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
