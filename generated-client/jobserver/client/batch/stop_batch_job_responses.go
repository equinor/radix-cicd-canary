// Code generated by go-swagger; DO NOT EDIT.

package batch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/equinor/radix-cicd-canary/generated-client/jobserver/models"
)

// StopBatchJobReader is a Reader for the StopBatchJob structure.
type StopBatchJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopBatchJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopBatchJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStopBatchJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStopBatchJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStopBatchJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /batches/{batchName}/jobs/{jobName}/stop] stopBatchJob", response, response.Code())
	}
}

// NewStopBatchJobOK creates a StopBatchJobOK with default headers values
func NewStopBatchJobOK() *StopBatchJobOK {
	return &StopBatchJobOK{}
}

/*
StopBatchJobOK describes a response with status code 200, with default header values.

Successful stop batch job
*/
type StopBatchJobOK struct {
	Payload *models.Status
}

// IsSuccess returns true when this stop batch job o k response has a 2xx status code
func (o *StopBatchJobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stop batch job o k response has a 3xx status code
func (o *StopBatchJobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop batch job o k response has a 4xx status code
func (o *StopBatchJobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop batch job o k response has a 5xx status code
func (o *StopBatchJobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stop batch job o k response a status code equal to that given
func (o *StopBatchJobOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stop batch job o k response
func (o *StopBatchJobOK) Code() int {
	return 200
}

func (o *StopBatchJobOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /batches/{batchName}/jobs/{jobName}/stop][%d] stopBatchJobOK %s", 200, payload)
}

func (o *StopBatchJobOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /batches/{batchName}/jobs/{jobName}/stop][%d] stopBatchJobOK %s", 200, payload)
}

func (o *StopBatchJobOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *StopBatchJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopBatchJobBadRequest creates a StopBatchJobBadRequest with default headers values
func NewStopBatchJobBadRequest() *StopBatchJobBadRequest {
	return &StopBatchJobBadRequest{}
}

/*
StopBatchJobBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type StopBatchJobBadRequest struct {
	Payload *models.Status
}

// IsSuccess returns true when this stop batch job bad request response has a 2xx status code
func (o *StopBatchJobBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop batch job bad request response has a 3xx status code
func (o *StopBatchJobBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop batch job bad request response has a 4xx status code
func (o *StopBatchJobBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop batch job bad request response has a 5xx status code
func (o *StopBatchJobBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stop batch job bad request response a status code equal to that given
func (o *StopBatchJobBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the stop batch job bad request response
func (o *StopBatchJobBadRequest) Code() int {
	return 400
}

func (o *StopBatchJobBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /batches/{batchName}/jobs/{jobName}/stop][%d] stopBatchJobBadRequest %s", 400, payload)
}

func (o *StopBatchJobBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /batches/{batchName}/jobs/{jobName}/stop][%d] stopBatchJobBadRequest %s", 400, payload)
}

func (o *StopBatchJobBadRequest) GetPayload() *models.Status {
	return o.Payload
}

func (o *StopBatchJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopBatchJobNotFound creates a StopBatchJobNotFound with default headers values
func NewStopBatchJobNotFound() *StopBatchJobNotFound {
	return &StopBatchJobNotFound{}
}

/*
StopBatchJobNotFound describes a response with status code 404, with default header values.

Not found
*/
type StopBatchJobNotFound struct {
	Payload *models.Status
}

// IsSuccess returns true when this stop batch job not found response has a 2xx status code
func (o *StopBatchJobNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop batch job not found response has a 3xx status code
func (o *StopBatchJobNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop batch job not found response has a 4xx status code
func (o *StopBatchJobNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop batch job not found response has a 5xx status code
func (o *StopBatchJobNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stop batch job not found response a status code equal to that given
func (o *StopBatchJobNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the stop batch job not found response
func (o *StopBatchJobNotFound) Code() int {
	return 404
}

func (o *StopBatchJobNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /batches/{batchName}/jobs/{jobName}/stop][%d] stopBatchJobNotFound %s", 404, payload)
}

func (o *StopBatchJobNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /batches/{batchName}/jobs/{jobName}/stop][%d] stopBatchJobNotFound %s", 404, payload)
}

func (o *StopBatchJobNotFound) GetPayload() *models.Status {
	return o.Payload
}

func (o *StopBatchJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopBatchJobInternalServerError creates a StopBatchJobInternalServerError with default headers values
func NewStopBatchJobInternalServerError() *StopBatchJobInternalServerError {
	return &StopBatchJobInternalServerError{}
}

/*
StopBatchJobInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type StopBatchJobInternalServerError struct {
	Payload *models.Status
}

// IsSuccess returns true when this stop batch job internal server error response has a 2xx status code
func (o *StopBatchJobInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop batch job internal server error response has a 3xx status code
func (o *StopBatchJobInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop batch job internal server error response has a 4xx status code
func (o *StopBatchJobInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop batch job internal server error response has a 5xx status code
func (o *StopBatchJobInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stop batch job internal server error response a status code equal to that given
func (o *StopBatchJobInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the stop batch job internal server error response
func (o *StopBatchJobInternalServerError) Code() int {
	return 500
}

func (o *StopBatchJobInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /batches/{batchName}/jobs/{jobName}/stop][%d] stopBatchJobInternalServerError %s", 500, payload)
}

func (o *StopBatchJobInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /batches/{batchName}/jobs/{jobName}/stop][%d] stopBatchJobInternalServerError %s", 500, payload)
}

func (o *StopBatchJobInternalServerError) GetPayload() *models.Status {
	return o.Payload
}

func (o *StopBatchJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
