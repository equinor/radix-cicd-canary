// Code generated by go-swagger; DO NOT EDIT.

package environment

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
		return nil, runtime.NewAPIError("[GET /applications/{appName}/environments/{envName}] getEnvironment", response, response.Code())
	}
}

// NewGetEnvironmentOK creates a GetEnvironmentOK with default headers values
func NewGetEnvironmentOK() *GetEnvironmentOK {
	return &GetEnvironmentOK{}
}

/*
GetEnvironmentOK describes a response with status code 200, with default header values.

Successful get environment
*/
type GetEnvironmentOK struct {
	Payload *models.Environment
}

// IsSuccess returns true when this get environment o k response has a 2xx status code
func (o *GetEnvironmentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get environment o k response has a 3xx status code
func (o *GetEnvironmentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get environment o k response has a 4xx status code
func (o *GetEnvironmentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get environment o k response has a 5xx status code
func (o *GetEnvironmentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get environment o k response a status code equal to that given
func (o *GetEnvironmentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get environment o k response
func (o *GetEnvironmentOK) Code() int {
	return 200
}

func (o *GetEnvironmentOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}][%d] getEnvironmentOK %s", 200, payload)
}

func (o *GetEnvironmentOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}][%d] getEnvironmentOK %s", 200, payload)
}

func (o *GetEnvironmentOK) GetPayload() *models.Environment {
	return o.Payload
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

/*
GetEnvironmentUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetEnvironmentUnauthorized struct {
}

// IsSuccess returns true when this get environment unauthorized response has a 2xx status code
func (o *GetEnvironmentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get environment unauthorized response has a 3xx status code
func (o *GetEnvironmentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get environment unauthorized response has a 4xx status code
func (o *GetEnvironmentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get environment unauthorized response has a 5xx status code
func (o *GetEnvironmentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get environment unauthorized response a status code equal to that given
func (o *GetEnvironmentUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get environment unauthorized response
func (o *GetEnvironmentUnauthorized) Code() int {
	return 401
}

func (o *GetEnvironmentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}][%d] getEnvironmentUnauthorized", 401)
}

func (o *GetEnvironmentUnauthorized) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}][%d] getEnvironmentUnauthorized", 401)
}

func (o *GetEnvironmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEnvironmentNotFound creates a GetEnvironmentNotFound with default headers values
func NewGetEnvironmentNotFound() *GetEnvironmentNotFound {
	return &GetEnvironmentNotFound{}
}

/*
GetEnvironmentNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetEnvironmentNotFound struct {
}

// IsSuccess returns true when this get environment not found response has a 2xx status code
func (o *GetEnvironmentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get environment not found response has a 3xx status code
func (o *GetEnvironmentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get environment not found response has a 4xx status code
func (o *GetEnvironmentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get environment not found response has a 5xx status code
func (o *GetEnvironmentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get environment not found response a status code equal to that given
func (o *GetEnvironmentNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get environment not found response
func (o *GetEnvironmentNotFound) Code() int {
	return 404
}

func (o *GetEnvironmentNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}][%d] getEnvironmentNotFound", 404)
}

func (o *GetEnvironmentNotFound) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}][%d] getEnvironmentNotFound", 404)
}

func (o *GetEnvironmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
