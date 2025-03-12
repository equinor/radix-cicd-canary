// Code generated by go-swagger; DO NOT EDIT.

package pipeline_job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// StopApplicationJobReader is a Reader for the StopApplicationJob structure.
type StopApplicationJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopApplicationJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewStopApplicationJobNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewStopApplicationJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStopApplicationJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /applications/{appName}/jobs/{jobName}/stop] stopApplicationJob", response, response.Code())
	}
}

// NewStopApplicationJobNoContent creates a StopApplicationJobNoContent with default headers values
func NewStopApplicationJobNoContent() *StopApplicationJobNoContent {
	return &StopApplicationJobNoContent{}
}

/*
StopApplicationJobNoContent describes a response with status code 204, with default header values.

Job stopped ok
*/
type StopApplicationJobNoContent struct {
}

// IsSuccess returns true when this stop application job no content response has a 2xx status code
func (o *StopApplicationJobNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stop application job no content response has a 3xx status code
func (o *StopApplicationJobNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop application job no content response has a 4xx status code
func (o *StopApplicationJobNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop application job no content response has a 5xx status code
func (o *StopApplicationJobNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this stop application job no content response a status code equal to that given
func (o *StopApplicationJobNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the stop application job no content response
func (o *StopApplicationJobNoContent) Code() int {
	return 204
}

func (o *StopApplicationJobNoContent) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/jobs/{jobName}/stop][%d] stopApplicationJobNoContent", 204)
}

func (o *StopApplicationJobNoContent) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/jobs/{jobName}/stop][%d] stopApplicationJobNoContent", 204)
}

func (o *StopApplicationJobNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopApplicationJobUnauthorized creates a StopApplicationJobUnauthorized with default headers values
func NewStopApplicationJobUnauthorized() *StopApplicationJobUnauthorized {
	return &StopApplicationJobUnauthorized{}
}

/*
StopApplicationJobUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StopApplicationJobUnauthorized struct {
}

// IsSuccess returns true when this stop application job unauthorized response has a 2xx status code
func (o *StopApplicationJobUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop application job unauthorized response has a 3xx status code
func (o *StopApplicationJobUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop application job unauthorized response has a 4xx status code
func (o *StopApplicationJobUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop application job unauthorized response has a 5xx status code
func (o *StopApplicationJobUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stop application job unauthorized response a status code equal to that given
func (o *StopApplicationJobUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the stop application job unauthorized response
func (o *StopApplicationJobUnauthorized) Code() int {
	return 401
}

func (o *StopApplicationJobUnauthorized) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/jobs/{jobName}/stop][%d] stopApplicationJobUnauthorized", 401)
}

func (o *StopApplicationJobUnauthorized) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/jobs/{jobName}/stop][%d] stopApplicationJobUnauthorized", 401)
}

func (o *StopApplicationJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopApplicationJobNotFound creates a StopApplicationJobNotFound with default headers values
func NewStopApplicationJobNotFound() *StopApplicationJobNotFound {
	return &StopApplicationJobNotFound{}
}

/*
StopApplicationJobNotFound describes a response with status code 404, with default header values.

Not found
*/
type StopApplicationJobNotFound struct {
}

// IsSuccess returns true when this stop application job not found response has a 2xx status code
func (o *StopApplicationJobNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop application job not found response has a 3xx status code
func (o *StopApplicationJobNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop application job not found response has a 4xx status code
func (o *StopApplicationJobNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop application job not found response has a 5xx status code
func (o *StopApplicationJobNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stop application job not found response a status code equal to that given
func (o *StopApplicationJobNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the stop application job not found response
func (o *StopApplicationJobNotFound) Code() int {
	return 404
}

func (o *StopApplicationJobNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/jobs/{jobName}/stop][%d] stopApplicationJobNotFound", 404)
}

func (o *StopApplicationJobNotFound) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/jobs/{jobName}/stop][%d] stopApplicationJobNotFound", 404)
}

func (o *StopApplicationJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
