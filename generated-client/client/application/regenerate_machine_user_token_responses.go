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

// RegenerateMachineUserTokenReader is a Reader for the RegenerateMachineUserToken structure.
type RegenerateMachineUserTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegenerateMachineUserTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRegenerateMachineUserTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRegenerateMachineUserTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRegenerateMachineUserTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRegenerateMachineUserTokenOK creates a RegenerateMachineUserTokenOK with default headers values
func NewRegenerateMachineUserTokenOK() *RegenerateMachineUserTokenOK {
	return &RegenerateMachineUserTokenOK{}
}

/* RegenerateMachineUserTokenOK describes a response with status code 200, with default header values.

Successful regenerate machine-user token
*/
type RegenerateMachineUserTokenOK struct {
	Payload *models.MachineUser
}

func (o *RegenerateMachineUserTokenOK) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/regenerate-machine-user-token][%d] regenerateMachineUserTokenOK  %+v", 200, o.Payload)
}
func (o *RegenerateMachineUserTokenOK) GetPayload() *models.MachineUser {
	return o.Payload
}

func (o *RegenerateMachineUserTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MachineUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegenerateMachineUserTokenUnauthorized creates a RegenerateMachineUserTokenUnauthorized with default headers values
func NewRegenerateMachineUserTokenUnauthorized() *RegenerateMachineUserTokenUnauthorized {
	return &RegenerateMachineUserTokenUnauthorized{}
}

/* RegenerateMachineUserTokenUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RegenerateMachineUserTokenUnauthorized struct {
}

func (o *RegenerateMachineUserTokenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/regenerate-machine-user-token][%d] regenerateMachineUserTokenUnauthorized ", 401)
}

func (o *RegenerateMachineUserTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegenerateMachineUserTokenNotFound creates a RegenerateMachineUserTokenNotFound with default headers values
func NewRegenerateMachineUserTokenNotFound() *RegenerateMachineUserTokenNotFound {
	return &RegenerateMachineUserTokenNotFound{}
}

/* RegenerateMachineUserTokenNotFound describes a response with status code 404, with default header values.

Not found
*/
type RegenerateMachineUserTokenNotFound struct {
}

func (o *RegenerateMachineUserTokenNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/regenerate-machine-user-token][%d] regenerateMachineUserTokenNotFound ", 404)
}

func (o *RegenerateMachineUserTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
