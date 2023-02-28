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

// GetJobsReader is a Reader for the GetJobs structure.
type GetJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetJobsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetJobsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetJobsOK creates a GetJobsOK with default headers values
func NewGetJobsOK() *GetJobsOK {
	return &GetJobsOK{}
}

/* GetJobsOK describes a response with status code 200, with default header values.

Successful get jobs
*/
type GetJobsOK struct {
	Payload []*models.JobStatus
}

func (o *GetJobsOK) Error() string {
	return fmt.Sprintf("[GET /jobs/][%d] getJobsOK  %+v", 200, o.Payload)
}
func (o *GetJobsOK) GetPayload() []*models.JobStatus {
	return o.Payload
}

func (o *GetJobsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobsInternalServerError creates a GetJobsInternalServerError with default headers values
func NewGetJobsInternalServerError() *GetJobsInternalServerError {
	return &GetJobsInternalServerError{}
}

/* GetJobsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetJobsInternalServerError struct {
	Payload *models.Status
}

func (o *GetJobsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /jobs/][%d] getJobsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetJobsInternalServerError) GetPayload() *models.Status {
	return o.Payload
}

func (o *GetJobsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
