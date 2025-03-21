// Code generated by go-swagger; DO NOT EDIT.

package component

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
)

// ComponentsReader is a Reader for the Components structure.
type ComponentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ComponentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewComponentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewComponentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /applications/{appName}/deployments/{deploymentName}/components] components", response, response.Code())
	}
}

// NewComponentsOK creates a ComponentsOK with default headers values
func NewComponentsOK() *ComponentsOK {
	return &ComponentsOK{}
}

/*
ComponentsOK describes a response with status code 200, with default header values.

pod log
*/
type ComponentsOK struct {
	Payload []*models.Component
}

// IsSuccess returns true when this components o k response has a 2xx status code
func (o *ComponentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this components o k response has a 3xx status code
func (o *ComponentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this components o k response has a 4xx status code
func (o *ComponentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this components o k response has a 5xx status code
func (o *ComponentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this components o k response a status code equal to that given
func (o *ComponentsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the components o k response
func (o *ComponentsOK) Code() int {
	return 200
}

func (o *ComponentsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /applications/{appName}/deployments/{deploymentName}/components][%d] componentsOK %s", 200, payload)
}

func (o *ComponentsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /applications/{appName}/deployments/{deploymentName}/components][%d] componentsOK %s", 200, payload)
}

func (o *ComponentsOK) GetPayload() []*models.Component {
	return o.Payload
}

func (o *ComponentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewComponentsNotFound creates a ComponentsNotFound with default headers values
func NewComponentsNotFound() *ComponentsNotFound {
	return &ComponentsNotFound{}
}

/*
ComponentsNotFound describes a response with status code 404, with default header values.

Not found
*/
type ComponentsNotFound struct {
}

// IsSuccess returns true when this components not found response has a 2xx status code
func (o *ComponentsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this components not found response has a 3xx status code
func (o *ComponentsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this components not found response has a 4xx status code
func (o *ComponentsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this components not found response has a 5xx status code
func (o *ComponentsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this components not found response a status code equal to that given
func (o *ComponentsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the components not found response
func (o *ComponentsNotFound) Code() int {
	return 404
}

func (o *ComponentsNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/deployments/{deploymentName}/components][%d] componentsNotFound", 404)
}

func (o *ComponentsNotFound) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/deployments/{deploymentName}/components][%d] componentsNotFound", 404)
}

func (o *ComponentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
