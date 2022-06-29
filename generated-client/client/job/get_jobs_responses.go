// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/equinor/radix-cicd-canary/generated-client/models"
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
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetJobsOK creates a GetJobsOK with default headers values
func NewGetJobsOK() *GetJobsOK {
	return &GetJobsOK{}
}

/*GetJobsOK handles this case with default header values.

scheduled jobs
*/
type GetJobsOK struct {
	Payload []*models.ScheduledJobSummary
}

func (o *GetJobsOK) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs][%d] getJobsOK  %+v", 200, o.Payload)
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

/*GetJobsNotFound handles this case with default header values.

Not found
*/
type GetJobsNotFound struct {
}

func (o *GetJobsNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs][%d] getJobsNotFound ", 404)
}

func (o *GetJobsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
