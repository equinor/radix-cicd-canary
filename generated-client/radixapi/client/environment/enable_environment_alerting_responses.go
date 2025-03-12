// Code generated by go-swagger; DO NOT EDIT.

package environment

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

// EnableEnvironmentAlertingReader is a Reader for the EnableEnvironmentAlerting structure.
type EnableEnvironmentAlertingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnableEnvironmentAlertingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnableEnvironmentAlertingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEnableEnvironmentAlertingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEnableEnvironmentAlertingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEnableEnvironmentAlertingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEnableEnvironmentAlertingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEnableEnvironmentAlertingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /applications/{appName}/environments/{envName}/alerting/enable] enableEnvironmentAlerting", response, response.Code())
	}
}

// NewEnableEnvironmentAlertingOK creates a EnableEnvironmentAlertingOK with default headers values
func NewEnableEnvironmentAlertingOK() *EnableEnvironmentAlertingOK {
	return &EnableEnvironmentAlertingOK{}
}

/*
EnableEnvironmentAlertingOK describes a response with status code 200, with default header values.

Successful enable alerting
*/
type EnableEnvironmentAlertingOK struct {
	Payload *models.AlertingConfig
}

// IsSuccess returns true when this enable environment alerting o k response has a 2xx status code
func (o *EnableEnvironmentAlertingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this enable environment alerting o k response has a 3xx status code
func (o *EnableEnvironmentAlertingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enable environment alerting o k response has a 4xx status code
func (o *EnableEnvironmentAlertingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this enable environment alerting o k response has a 5xx status code
func (o *EnableEnvironmentAlertingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this enable environment alerting o k response a status code equal to that given
func (o *EnableEnvironmentAlertingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the enable environment alerting o k response
func (o *EnableEnvironmentAlertingOK) Code() int {
	return 200
}

func (o *EnableEnvironmentAlertingOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/enable][%d] enableEnvironmentAlertingOK %s", 200, payload)
}

func (o *EnableEnvironmentAlertingOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/enable][%d] enableEnvironmentAlertingOK %s", 200, payload)
}

func (o *EnableEnvironmentAlertingOK) GetPayload() *models.AlertingConfig {
	return o.Payload
}

func (o *EnableEnvironmentAlertingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertingConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnableEnvironmentAlertingBadRequest creates a EnableEnvironmentAlertingBadRequest with default headers values
func NewEnableEnvironmentAlertingBadRequest() *EnableEnvironmentAlertingBadRequest {
	return &EnableEnvironmentAlertingBadRequest{}
}

/*
EnableEnvironmentAlertingBadRequest describes a response with status code 400, with default header values.

Alerting already enabled
*/
type EnableEnvironmentAlertingBadRequest struct {
}

// IsSuccess returns true when this enable environment alerting bad request response has a 2xx status code
func (o *EnableEnvironmentAlertingBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enable environment alerting bad request response has a 3xx status code
func (o *EnableEnvironmentAlertingBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enable environment alerting bad request response has a 4xx status code
func (o *EnableEnvironmentAlertingBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this enable environment alerting bad request response has a 5xx status code
func (o *EnableEnvironmentAlertingBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this enable environment alerting bad request response a status code equal to that given
func (o *EnableEnvironmentAlertingBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the enable environment alerting bad request response
func (o *EnableEnvironmentAlertingBadRequest) Code() int {
	return 400
}

func (o *EnableEnvironmentAlertingBadRequest) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/enable][%d] enableEnvironmentAlertingBadRequest", 400)
}

func (o *EnableEnvironmentAlertingBadRequest) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/enable][%d] enableEnvironmentAlertingBadRequest", 400)
}

func (o *EnableEnvironmentAlertingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnableEnvironmentAlertingUnauthorized creates a EnableEnvironmentAlertingUnauthorized with default headers values
func NewEnableEnvironmentAlertingUnauthorized() *EnableEnvironmentAlertingUnauthorized {
	return &EnableEnvironmentAlertingUnauthorized{}
}

/*
EnableEnvironmentAlertingUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type EnableEnvironmentAlertingUnauthorized struct {
}

// IsSuccess returns true when this enable environment alerting unauthorized response has a 2xx status code
func (o *EnableEnvironmentAlertingUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enable environment alerting unauthorized response has a 3xx status code
func (o *EnableEnvironmentAlertingUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enable environment alerting unauthorized response has a 4xx status code
func (o *EnableEnvironmentAlertingUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this enable environment alerting unauthorized response has a 5xx status code
func (o *EnableEnvironmentAlertingUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this enable environment alerting unauthorized response a status code equal to that given
func (o *EnableEnvironmentAlertingUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the enable environment alerting unauthorized response
func (o *EnableEnvironmentAlertingUnauthorized) Code() int {
	return 401
}

func (o *EnableEnvironmentAlertingUnauthorized) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/enable][%d] enableEnvironmentAlertingUnauthorized", 401)
}

func (o *EnableEnvironmentAlertingUnauthorized) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/enable][%d] enableEnvironmentAlertingUnauthorized", 401)
}

func (o *EnableEnvironmentAlertingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnableEnvironmentAlertingForbidden creates a EnableEnvironmentAlertingForbidden with default headers values
func NewEnableEnvironmentAlertingForbidden() *EnableEnvironmentAlertingForbidden {
	return &EnableEnvironmentAlertingForbidden{}
}

/*
EnableEnvironmentAlertingForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type EnableEnvironmentAlertingForbidden struct {
}

// IsSuccess returns true when this enable environment alerting forbidden response has a 2xx status code
func (o *EnableEnvironmentAlertingForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enable environment alerting forbidden response has a 3xx status code
func (o *EnableEnvironmentAlertingForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enable environment alerting forbidden response has a 4xx status code
func (o *EnableEnvironmentAlertingForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this enable environment alerting forbidden response has a 5xx status code
func (o *EnableEnvironmentAlertingForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this enable environment alerting forbidden response a status code equal to that given
func (o *EnableEnvironmentAlertingForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the enable environment alerting forbidden response
func (o *EnableEnvironmentAlertingForbidden) Code() int {
	return 403
}

func (o *EnableEnvironmentAlertingForbidden) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/enable][%d] enableEnvironmentAlertingForbidden", 403)
}

func (o *EnableEnvironmentAlertingForbidden) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/enable][%d] enableEnvironmentAlertingForbidden", 403)
}

func (o *EnableEnvironmentAlertingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnableEnvironmentAlertingNotFound creates a EnableEnvironmentAlertingNotFound with default headers values
func NewEnableEnvironmentAlertingNotFound() *EnableEnvironmentAlertingNotFound {
	return &EnableEnvironmentAlertingNotFound{}
}

/*
EnableEnvironmentAlertingNotFound describes a response with status code 404, with default header values.

Not found
*/
type EnableEnvironmentAlertingNotFound struct {
}

// IsSuccess returns true when this enable environment alerting not found response has a 2xx status code
func (o *EnableEnvironmentAlertingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enable environment alerting not found response has a 3xx status code
func (o *EnableEnvironmentAlertingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enable environment alerting not found response has a 4xx status code
func (o *EnableEnvironmentAlertingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this enable environment alerting not found response has a 5xx status code
func (o *EnableEnvironmentAlertingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this enable environment alerting not found response a status code equal to that given
func (o *EnableEnvironmentAlertingNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the enable environment alerting not found response
func (o *EnableEnvironmentAlertingNotFound) Code() int {
	return 404
}

func (o *EnableEnvironmentAlertingNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/enable][%d] enableEnvironmentAlertingNotFound", 404)
}

func (o *EnableEnvironmentAlertingNotFound) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/enable][%d] enableEnvironmentAlertingNotFound", 404)
}

func (o *EnableEnvironmentAlertingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnableEnvironmentAlertingInternalServerError creates a EnableEnvironmentAlertingInternalServerError with default headers values
func NewEnableEnvironmentAlertingInternalServerError() *EnableEnvironmentAlertingInternalServerError {
	return &EnableEnvironmentAlertingInternalServerError{}
}

/*
EnableEnvironmentAlertingInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type EnableEnvironmentAlertingInternalServerError struct {
}

// IsSuccess returns true when this enable environment alerting internal server error response has a 2xx status code
func (o *EnableEnvironmentAlertingInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this enable environment alerting internal server error response has a 3xx status code
func (o *EnableEnvironmentAlertingInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enable environment alerting internal server error response has a 4xx status code
func (o *EnableEnvironmentAlertingInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this enable environment alerting internal server error response has a 5xx status code
func (o *EnableEnvironmentAlertingInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this enable environment alerting internal server error response a status code equal to that given
func (o *EnableEnvironmentAlertingInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the enable environment alerting internal server error response
func (o *EnableEnvironmentAlertingInternalServerError) Code() int {
	return 500
}

func (o *EnableEnvironmentAlertingInternalServerError) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/enable][%d] enableEnvironmentAlertingInternalServerError", 500)
}

func (o *EnableEnvironmentAlertingInternalServerError) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/enable][%d] enableEnvironmentAlertingInternalServerError", 500)
}

func (o *EnableEnvironmentAlertingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
