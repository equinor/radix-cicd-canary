// Code generated by go-swagger; DO NOT EDIT.

package application

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

// GetBuildSecretsReader is a Reader for the GetBuildSecrets structure.
type GetBuildSecretsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBuildSecretsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBuildSecretsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetBuildSecretsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBuildSecretsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /applications/{appName}/buildsecrets] getBuildSecrets", response, response.Code())
	}
}

// NewGetBuildSecretsOK creates a GetBuildSecretsOK with default headers values
func NewGetBuildSecretsOK() *GetBuildSecretsOK {
	return &GetBuildSecretsOK{}
}

/*
GetBuildSecretsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetBuildSecretsOK struct {
	Payload []*models.BuildSecret
}

// IsSuccess returns true when this get build secrets o k response has a 2xx status code
func (o *GetBuildSecretsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get build secrets o k response has a 3xx status code
func (o *GetBuildSecretsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get build secrets o k response has a 4xx status code
func (o *GetBuildSecretsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get build secrets o k response has a 5xx status code
func (o *GetBuildSecretsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get build secrets o k response a status code equal to that given
func (o *GetBuildSecretsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get build secrets o k response
func (o *GetBuildSecretsOK) Code() int {
	return 200
}

func (o *GetBuildSecretsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /applications/{appName}/buildsecrets][%d] getBuildSecretsOK %s", 200, payload)
}

func (o *GetBuildSecretsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /applications/{appName}/buildsecrets][%d] getBuildSecretsOK %s", 200, payload)
}

func (o *GetBuildSecretsOK) GetPayload() []*models.BuildSecret {
	return o.Payload
}

func (o *GetBuildSecretsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBuildSecretsUnauthorized creates a GetBuildSecretsUnauthorized with default headers values
func NewGetBuildSecretsUnauthorized() *GetBuildSecretsUnauthorized {
	return &GetBuildSecretsUnauthorized{}
}

/*
GetBuildSecretsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetBuildSecretsUnauthorized struct {
}

// IsSuccess returns true when this get build secrets unauthorized response has a 2xx status code
func (o *GetBuildSecretsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get build secrets unauthorized response has a 3xx status code
func (o *GetBuildSecretsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get build secrets unauthorized response has a 4xx status code
func (o *GetBuildSecretsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get build secrets unauthorized response has a 5xx status code
func (o *GetBuildSecretsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get build secrets unauthorized response a status code equal to that given
func (o *GetBuildSecretsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get build secrets unauthorized response
func (o *GetBuildSecretsUnauthorized) Code() int {
	return 401
}

func (o *GetBuildSecretsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/buildsecrets][%d] getBuildSecretsUnauthorized", 401)
}

func (o *GetBuildSecretsUnauthorized) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/buildsecrets][%d] getBuildSecretsUnauthorized", 401)
}

func (o *GetBuildSecretsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBuildSecretsNotFound creates a GetBuildSecretsNotFound with default headers values
func NewGetBuildSecretsNotFound() *GetBuildSecretsNotFound {
	return &GetBuildSecretsNotFound{}
}

/*
GetBuildSecretsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetBuildSecretsNotFound struct {
}

// IsSuccess returns true when this get build secrets not found response has a 2xx status code
func (o *GetBuildSecretsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get build secrets not found response has a 3xx status code
func (o *GetBuildSecretsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get build secrets not found response has a 4xx status code
func (o *GetBuildSecretsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get build secrets not found response has a 5xx status code
func (o *GetBuildSecretsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get build secrets not found response a status code equal to that given
func (o *GetBuildSecretsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get build secrets not found response
func (o *GetBuildSecretsNotFound) Code() int {
	return 404
}

func (o *GetBuildSecretsNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/buildsecrets][%d] getBuildSecretsNotFound", 404)
}

func (o *GetBuildSecretsNotFound) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/buildsecrets][%d] getBuildSecretsNotFound", 404)
}

func (o *GetBuildSecretsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
