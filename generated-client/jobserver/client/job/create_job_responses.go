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

// CreateJobReader is a Reader for the CreateJob structure.
type CreateJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateJobUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /jobs] createJob", response, response.Code())
	}
}

// NewCreateJobOK creates a CreateJobOK with default headers values
func NewCreateJobOK() *CreateJobOK {
	return &CreateJobOK{}
}

/*
CreateJobOK describes a response with status code 200, with default header values.

Successful create job
*/
type CreateJobOK struct {
	Payload *models.JobStatus
}

// IsSuccess returns true when this create job o k response has a 2xx status code
func (o *CreateJobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create job o k response has a 3xx status code
func (o *CreateJobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create job o k response has a 4xx status code
func (o *CreateJobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create job o k response has a 5xx status code
func (o *CreateJobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create job o k response a status code equal to that given
func (o *CreateJobOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create job o k response
func (o *CreateJobOK) Code() int {
	return 200
}

func (o *CreateJobOK) Error() string {
	return fmt.Sprintf("[POST /jobs][%d] createJobOK  %+v", 200, o.Payload)
}

func (o *CreateJobOK) String() string {
	return fmt.Sprintf("[POST /jobs][%d] createJobOK  %+v", 200, o.Payload)
}

func (o *CreateJobOK) GetPayload() *models.JobStatus {
	return o.Payload
}

func (o *CreateJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateJobBadRequest creates a CreateJobBadRequest with default headers values
func NewCreateJobBadRequest() *CreateJobBadRequest {
	return &CreateJobBadRequest{}
}

/*
CreateJobBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateJobBadRequest struct {
	Payload *models.Status
}

// IsSuccess returns true when this create job bad request response has a 2xx status code
func (o *CreateJobBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create job bad request response has a 3xx status code
func (o *CreateJobBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create job bad request response has a 4xx status code
func (o *CreateJobBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create job bad request response has a 5xx status code
func (o *CreateJobBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create job bad request response a status code equal to that given
func (o *CreateJobBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create job bad request response
func (o *CreateJobBadRequest) Code() int {
	return 400
}

func (o *CreateJobBadRequest) Error() string {
	return fmt.Sprintf("[POST /jobs][%d] createJobBadRequest  %+v", 400, o.Payload)
}

func (o *CreateJobBadRequest) String() string {
	return fmt.Sprintf("[POST /jobs][%d] createJobBadRequest  %+v", 400, o.Payload)
}

func (o *CreateJobBadRequest) GetPayload() *models.Status {
	return o.Payload
}

func (o *CreateJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateJobNotFound creates a CreateJobNotFound with default headers values
func NewCreateJobNotFound() *CreateJobNotFound {
	return &CreateJobNotFound{}
}

/*
CreateJobNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateJobNotFound struct {
	Payload *models.Status
}

// IsSuccess returns true when this create job not found response has a 2xx status code
func (o *CreateJobNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create job not found response has a 3xx status code
func (o *CreateJobNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create job not found response has a 4xx status code
func (o *CreateJobNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create job not found response has a 5xx status code
func (o *CreateJobNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create job not found response a status code equal to that given
func (o *CreateJobNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create job not found response
func (o *CreateJobNotFound) Code() int {
	return 404
}

func (o *CreateJobNotFound) Error() string {
	return fmt.Sprintf("[POST /jobs][%d] createJobNotFound  %+v", 404, o.Payload)
}

func (o *CreateJobNotFound) String() string {
	return fmt.Sprintf("[POST /jobs][%d] createJobNotFound  %+v", 404, o.Payload)
}

func (o *CreateJobNotFound) GetPayload() *models.Status {
	return o.Payload
}

func (o *CreateJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateJobUnprocessableEntity creates a CreateJobUnprocessableEntity with default headers values
func NewCreateJobUnprocessableEntity() *CreateJobUnprocessableEntity {
	return &CreateJobUnprocessableEntity{}
}

/*
CreateJobUnprocessableEntity describes a response with status code 422, with default header values.

Invalid data in request
*/
type CreateJobUnprocessableEntity struct {
	Payload *models.Status
}

// IsSuccess returns true when this create job unprocessable entity response has a 2xx status code
func (o *CreateJobUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create job unprocessable entity response has a 3xx status code
func (o *CreateJobUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create job unprocessable entity response has a 4xx status code
func (o *CreateJobUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this create job unprocessable entity response has a 5xx status code
func (o *CreateJobUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this create job unprocessable entity response a status code equal to that given
func (o *CreateJobUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the create job unprocessable entity response
func (o *CreateJobUnprocessableEntity) Code() int {
	return 422
}

func (o *CreateJobUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /jobs][%d] createJobUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateJobUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /jobs][%d] createJobUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateJobUnprocessableEntity) GetPayload() *models.Status {
	return o.Payload
}

func (o *CreateJobUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateJobInternalServerError creates a CreateJobInternalServerError with default headers values
func NewCreateJobInternalServerError() *CreateJobInternalServerError {
	return &CreateJobInternalServerError{}
}

/*
CreateJobInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type CreateJobInternalServerError struct {
	Payload *models.Status
}

// IsSuccess returns true when this create job internal server error response has a 2xx status code
func (o *CreateJobInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create job internal server error response has a 3xx status code
func (o *CreateJobInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create job internal server error response has a 4xx status code
func (o *CreateJobInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create job internal server error response has a 5xx status code
func (o *CreateJobInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create job internal server error response a status code equal to that given
func (o *CreateJobInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create job internal server error response
func (o *CreateJobInternalServerError) Code() int {
	return 500
}

func (o *CreateJobInternalServerError) Error() string {
	return fmt.Sprintf("[POST /jobs][%d] createJobInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateJobInternalServerError) String() string {
	return fmt.Sprintf("[POST /jobs][%d] createJobInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateJobInternalServerError) GetPayload() *models.Status {
	return o.Payload
}

func (o *CreateJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
