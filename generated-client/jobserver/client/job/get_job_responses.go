// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/equinor/radix-cicd-canary/generated-client/jobserver/models"
)

// GetJobReader is a Reader for the GetJob structure.
type GetJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetJobOK creates a GetJobOK with default headers values
func NewGetJobOK() *GetJobOK {
	return &GetJobOK{}
}

/* GetJobOK describes a response with status code 200, with default header values.

Successful get job
*/
type GetJobOK struct {
	Payload *models.JobStatus
}

func (o *GetJobOK) Error() string {
	return fmt.Sprintf("[GET /jobs/{jobName}][%d] getJobOK  %+v", 200, o.Payload)
}
func (o *GetJobOK) GetPayload() *models.JobStatus {
	return o.Payload
}

func (o *GetJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobNotFound creates a GetJobNotFound with default headers values
func NewGetJobNotFound() *GetJobNotFound {
	return &GetJobNotFound{}
}

/* GetJobNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetJobNotFound struct {
	Payload *models.Status
}

func (o *GetJobNotFound) Error() string {
	return fmt.Sprintf("[GET /jobs/{jobName}][%d] getJobNotFound  %+v", 404, o.Payload)
}
func (o *GetJobNotFound) GetPayload() *models.Status {
	return o.Payload
}

func (o *GetJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobInternalServerError creates a GetJobInternalServerError with default headers values
func NewGetJobInternalServerError() *GetJobInternalServerError {
	return &GetJobInternalServerError{}
}

/* GetJobInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetJobInternalServerError struct {
	Payload *models.Status
}

func (o *GetJobInternalServerError) Error() string {
	return fmt.Sprintf("[GET /jobs/{jobName}][%d] getJobInternalServerError  %+v", 500, o.Payload)
}
func (o *GetJobInternalServerError) GetPayload() *models.Status {
	return o.Payload
}

func (o *GetJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}