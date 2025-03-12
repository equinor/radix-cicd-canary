// Code generated by go-swagger; DO NOT EDIT.

package job

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

// DeleteJobReader is a Reader for the DeleteJob structure.
type DeleteJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /jobs/{jobName}] deleteJob", response, response.Code())
	}
}

// NewDeleteJobOK creates a DeleteJobOK with default headers values
func NewDeleteJobOK() *DeleteJobOK {
	return &DeleteJobOK{}
}

/*
DeleteJobOK describes a response with status code 200, with default header values.

Successful delete job
*/
type DeleteJobOK struct {
	Payload *models.Status
}

// IsSuccess returns true when this delete job o k response has a 2xx status code
func (o *DeleteJobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete job o k response has a 3xx status code
func (o *DeleteJobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete job o k response has a 4xx status code
func (o *DeleteJobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete job o k response has a 5xx status code
func (o *DeleteJobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete job o k response a status code equal to that given
func (o *DeleteJobOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete job o k response
func (o *DeleteJobOK) Code() int {
	return 200
}

func (o *DeleteJobOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /jobs/{jobName}][%d] deleteJobOK %s", 200, payload)
}

func (o *DeleteJobOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /jobs/{jobName}][%d] deleteJobOK %s", 200, payload)
}

func (o *DeleteJobOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *DeleteJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJobNotFound creates a DeleteJobNotFound with default headers values
func NewDeleteJobNotFound() *DeleteJobNotFound {
	return &DeleteJobNotFound{}
}

/*
DeleteJobNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteJobNotFound struct {
	Payload *models.Status
}

// IsSuccess returns true when this delete job not found response has a 2xx status code
func (o *DeleteJobNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete job not found response has a 3xx status code
func (o *DeleteJobNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete job not found response has a 4xx status code
func (o *DeleteJobNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete job not found response has a 5xx status code
func (o *DeleteJobNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete job not found response a status code equal to that given
func (o *DeleteJobNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete job not found response
func (o *DeleteJobNotFound) Code() int {
	return 404
}

func (o *DeleteJobNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /jobs/{jobName}][%d] deleteJobNotFound %s", 404, payload)
}

func (o *DeleteJobNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /jobs/{jobName}][%d] deleteJobNotFound %s", 404, payload)
}

func (o *DeleteJobNotFound) GetPayload() *models.Status {
	return o.Payload
}

func (o *DeleteJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJobInternalServerError creates a DeleteJobInternalServerError with default headers values
func NewDeleteJobInternalServerError() *DeleteJobInternalServerError {
	return &DeleteJobInternalServerError{}
}

/*
DeleteJobInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteJobInternalServerError struct {
	Payload *models.Status
}

// IsSuccess returns true when this delete job internal server error response has a 2xx status code
func (o *DeleteJobInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete job internal server error response has a 3xx status code
func (o *DeleteJobInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete job internal server error response has a 4xx status code
func (o *DeleteJobInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete job internal server error response has a 5xx status code
func (o *DeleteJobInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete job internal server error response a status code equal to that given
func (o *DeleteJobInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete job internal server error response
func (o *DeleteJobInternalServerError) Code() int {
	return 500
}

func (o *DeleteJobInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /jobs/{jobName}][%d] deleteJobInternalServerError %s", 500, payload)
}

func (o *DeleteJobInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /jobs/{jobName}][%d] deleteJobInternalServerError %s", 500, payload)
}

func (o *DeleteJobInternalServerError) GetPayload() *models.Status {
	return o.Payload
}

func (o *DeleteJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
