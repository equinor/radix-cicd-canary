// Code generated by go-swagger; DO NOT EDIT.

package pipeline_job

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

// GetApplicationJobReader is a Reader for the GetApplicationJob structure.
type GetApplicationJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetApplicationJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetApplicationJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /applications/{appName}/jobs/{jobName}] getApplicationJob", response, response.Code())
	}
}

// NewGetApplicationJobOK creates a GetApplicationJobOK with default headers values
func NewGetApplicationJobOK() *GetApplicationJobOK {
	return &GetApplicationJobOK{}
}

/*
GetApplicationJobOK describes a response with status code 200, with default header values.

Successful get job
*/
type GetApplicationJobOK struct {
	Payload *models.Job
}

// IsSuccess returns true when this get application job o k response has a 2xx status code
func (o *GetApplicationJobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get application job o k response has a 3xx status code
func (o *GetApplicationJobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application job o k response has a 4xx status code
func (o *GetApplicationJobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get application job o k response has a 5xx status code
func (o *GetApplicationJobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get application job o k response a status code equal to that given
func (o *GetApplicationJobOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get application job o k response
func (o *GetApplicationJobOK) Code() int {
	return 200
}

func (o *GetApplicationJobOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}][%d] getApplicationJobOK %s", 200, payload)
}

func (o *GetApplicationJobOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}][%d] getApplicationJobOK %s", 200, payload)
}

func (o *GetApplicationJobOK) GetPayload() *models.Job {
	return o.Payload
}

func (o *GetApplicationJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Job)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationJobUnauthorized creates a GetApplicationJobUnauthorized with default headers values
func NewGetApplicationJobUnauthorized() *GetApplicationJobUnauthorized {
	return &GetApplicationJobUnauthorized{}
}

/*
GetApplicationJobUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetApplicationJobUnauthorized struct {
}

// IsSuccess returns true when this get application job unauthorized response has a 2xx status code
func (o *GetApplicationJobUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get application job unauthorized response has a 3xx status code
func (o *GetApplicationJobUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application job unauthorized response has a 4xx status code
func (o *GetApplicationJobUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get application job unauthorized response has a 5xx status code
func (o *GetApplicationJobUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get application job unauthorized response a status code equal to that given
func (o *GetApplicationJobUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get application job unauthorized response
func (o *GetApplicationJobUnauthorized) Code() int {
	return 401
}

func (o *GetApplicationJobUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}][%d] getApplicationJobUnauthorized", 401)
}

func (o *GetApplicationJobUnauthorized) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}][%d] getApplicationJobUnauthorized", 401)
}

func (o *GetApplicationJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationJobNotFound creates a GetApplicationJobNotFound with default headers values
func NewGetApplicationJobNotFound() *GetApplicationJobNotFound {
	return &GetApplicationJobNotFound{}
}

/*
GetApplicationJobNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetApplicationJobNotFound struct {
}

// IsSuccess returns true when this get application job not found response has a 2xx status code
func (o *GetApplicationJobNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get application job not found response has a 3xx status code
func (o *GetApplicationJobNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get application job not found response has a 4xx status code
func (o *GetApplicationJobNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get application job not found response has a 5xx status code
func (o *GetApplicationJobNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get application job not found response a status code equal to that given
func (o *GetApplicationJobNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get application job not found response
func (o *GetApplicationJobNotFound) Code() int {
	return 404
}

func (o *GetApplicationJobNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}][%d] getApplicationJobNotFound", 404)
}

func (o *GetApplicationJobNotFound) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}][%d] getApplicationJobNotFound", 404)
}

func (o *GetApplicationJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
