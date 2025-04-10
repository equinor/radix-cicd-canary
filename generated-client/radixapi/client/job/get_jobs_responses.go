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

	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
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
	case 404:
		result := NewGetJobsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs] getJobs", response, response.Code())
	}
}

// NewGetJobsOK creates a GetJobsOK with default headers values
func NewGetJobsOK() *GetJobsOK {
	return &GetJobsOK{}
}

/*
GetJobsOK describes a response with status code 200, with default header values.

scheduled jobs
*/
type GetJobsOK struct {
	Payload []*models.ScheduledJobSummary
}

// IsSuccess returns true when this get jobs o k response has a 2xx status code
func (o *GetJobsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get jobs o k response has a 3xx status code
func (o *GetJobsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get jobs o k response has a 4xx status code
func (o *GetJobsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get jobs o k response has a 5xx status code
func (o *GetJobsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get jobs o k response a status code equal to that given
func (o *GetJobsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get jobs o k response
func (o *GetJobsOK) Code() int {
	return 200
}

func (o *GetJobsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs][%d] getJobsOK %s", 200, payload)
}

func (o *GetJobsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs][%d] getJobsOK %s", 200, payload)
}

func (o *GetJobsOK) GetPayload() []*models.ScheduledJobSummary {
	return o.Payload
}

func (o *GetJobsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobsNotFound creates a GetJobsNotFound with default headers values
func NewGetJobsNotFound() *GetJobsNotFound {
	return &GetJobsNotFound{}
}

/*
GetJobsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetJobsNotFound struct {
}

// IsSuccess returns true when this get jobs not found response has a 2xx status code
func (o *GetJobsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get jobs not found response has a 3xx status code
func (o *GetJobsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get jobs not found response has a 4xx status code
func (o *GetJobsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get jobs not found response has a 5xx status code
func (o *GetJobsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get jobs not found response a status code equal to that given
func (o *GetJobsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get jobs not found response
func (o *GetJobsNotFound) Code() int {
	return 404
}

func (o *GetJobsNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs][%d] getJobsNotFound", 404)
}

func (o *GetJobsNotFound) String() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs][%d] getJobsNotFound", 404)
}

func (o *GetJobsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
