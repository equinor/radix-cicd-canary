// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// StopAllJobsReader is a Reader for the StopAllJobs structure.
type StopAllJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopAllJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewStopAllJobsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStopAllJobsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStopAllJobsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStopAllJobsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStopAllJobsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/stop] stopAllJobs", response, response.Code())
	}
}

// NewStopAllJobsNoContent creates a StopAllJobsNoContent with default headers values
func NewStopAllJobsNoContent() *StopAllJobsNoContent {
	return &StopAllJobsNoContent{}
}

/*
StopAllJobsNoContent describes a response with status code 204, with default header values.

Success
*/
type StopAllJobsNoContent struct {
}

// IsSuccess returns true when this stop all jobs no content response has a 2xx status code
func (o *StopAllJobsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stop all jobs no content response has a 3xx status code
func (o *StopAllJobsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop all jobs no content response has a 4xx status code
func (o *StopAllJobsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop all jobs no content response has a 5xx status code
func (o *StopAllJobsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this stop all jobs no content response a status code equal to that given
func (o *StopAllJobsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the stop all jobs no content response
func (o *StopAllJobsNoContent) Code() int {
	return 204
}

func (o *StopAllJobsNoContent) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/stop][%d] stopAllJobsNoContent", 204)
}

func (o *StopAllJobsNoContent) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/stop][%d] stopAllJobsNoContent", 204)
}

func (o *StopAllJobsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopAllJobsBadRequest creates a StopAllJobsBadRequest with default headers values
func NewStopAllJobsBadRequest() *StopAllJobsBadRequest {
	return &StopAllJobsBadRequest{}
}

/*
StopAllJobsBadRequest describes a response with status code 400, with default header values.

Invalid job
*/
type StopAllJobsBadRequest struct {
}

// IsSuccess returns true when this stop all jobs bad request response has a 2xx status code
func (o *StopAllJobsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop all jobs bad request response has a 3xx status code
func (o *StopAllJobsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop all jobs bad request response has a 4xx status code
func (o *StopAllJobsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop all jobs bad request response has a 5xx status code
func (o *StopAllJobsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stop all jobs bad request response a status code equal to that given
func (o *StopAllJobsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the stop all jobs bad request response
func (o *StopAllJobsBadRequest) Code() int {
	return 400
}

func (o *StopAllJobsBadRequest) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/stop][%d] stopAllJobsBadRequest", 400)
}

func (o *StopAllJobsBadRequest) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/stop][%d] stopAllJobsBadRequest", 400)
}

func (o *StopAllJobsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopAllJobsUnauthorized creates a StopAllJobsUnauthorized with default headers values
func NewStopAllJobsUnauthorized() *StopAllJobsUnauthorized {
	return &StopAllJobsUnauthorized{}
}

/*
StopAllJobsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StopAllJobsUnauthorized struct {
}

// IsSuccess returns true when this stop all jobs unauthorized response has a 2xx status code
func (o *StopAllJobsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop all jobs unauthorized response has a 3xx status code
func (o *StopAllJobsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop all jobs unauthorized response has a 4xx status code
func (o *StopAllJobsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop all jobs unauthorized response has a 5xx status code
func (o *StopAllJobsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stop all jobs unauthorized response a status code equal to that given
func (o *StopAllJobsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the stop all jobs unauthorized response
func (o *StopAllJobsUnauthorized) Code() int {
	return 401
}

func (o *StopAllJobsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/stop][%d] stopAllJobsUnauthorized", 401)
}

func (o *StopAllJobsUnauthorized) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/stop][%d] stopAllJobsUnauthorized", 401)
}

func (o *StopAllJobsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopAllJobsForbidden creates a StopAllJobsForbidden with default headers values
func NewStopAllJobsForbidden() *StopAllJobsForbidden {
	return &StopAllJobsForbidden{}
}

/*
StopAllJobsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StopAllJobsForbidden struct {
}

// IsSuccess returns true when this stop all jobs forbidden response has a 2xx status code
func (o *StopAllJobsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop all jobs forbidden response has a 3xx status code
func (o *StopAllJobsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop all jobs forbidden response has a 4xx status code
func (o *StopAllJobsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop all jobs forbidden response has a 5xx status code
func (o *StopAllJobsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stop all jobs forbidden response a status code equal to that given
func (o *StopAllJobsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the stop all jobs forbidden response
func (o *StopAllJobsForbidden) Code() int {
	return 403
}

func (o *StopAllJobsForbidden) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/stop][%d] stopAllJobsForbidden", 403)
}

func (o *StopAllJobsForbidden) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/stop][%d] stopAllJobsForbidden", 403)
}

func (o *StopAllJobsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopAllJobsNotFound creates a StopAllJobsNotFound with default headers values
func NewStopAllJobsNotFound() *StopAllJobsNotFound {
	return &StopAllJobsNotFound{}
}

/*
StopAllJobsNotFound describes a response with status code 404, with default header values.

Not found
*/
type StopAllJobsNotFound struct {
}

// IsSuccess returns true when this stop all jobs not found response has a 2xx status code
func (o *StopAllJobsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop all jobs not found response has a 3xx status code
func (o *StopAllJobsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop all jobs not found response has a 4xx status code
func (o *StopAllJobsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop all jobs not found response has a 5xx status code
func (o *StopAllJobsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stop all jobs not found response a status code equal to that given
func (o *StopAllJobsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the stop all jobs not found response
func (o *StopAllJobsNotFound) Code() int {
	return 404
}

func (o *StopAllJobsNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/stop][%d] stopAllJobsNotFound", 404)
}

func (o *StopAllJobsNotFound) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/stop][%d] stopAllJobsNotFound", 404)
}

func (o *StopAllJobsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
