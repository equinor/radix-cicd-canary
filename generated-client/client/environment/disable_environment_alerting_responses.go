// Code generated by go-swagger; DO NOT EDIT.

package environment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/equinor/radix-cicd-canary/generated-client/models"
)

// DisableEnvironmentAlertingReader is a Reader for the DisableEnvironmentAlerting structure.
type DisableEnvironmentAlertingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DisableEnvironmentAlertingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDisableEnvironmentAlertingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDisableEnvironmentAlertingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDisableEnvironmentAlertingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDisableEnvironmentAlertingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDisableEnvironmentAlertingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDisableEnvironmentAlertingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDisableEnvironmentAlertingOK creates a DisableEnvironmentAlertingOK with default headers values
func NewDisableEnvironmentAlertingOK() *DisableEnvironmentAlertingOK {
	return &DisableEnvironmentAlertingOK{}
}

/* DisableEnvironmentAlertingOK describes a response with status code 200, with default header values.

Successful disable alerting
*/
type DisableEnvironmentAlertingOK struct {
	Payload *models.AlertingConfig
}

func (o *DisableEnvironmentAlertingOK) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/disable][%d] disableEnvironmentAlertingOK  %+v", 200, o.Payload)
}
func (o *DisableEnvironmentAlertingOK) GetPayload() *models.AlertingConfig {
	return o.Payload
}

func (o *DisableEnvironmentAlertingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertingConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDisableEnvironmentAlertingBadRequest creates a DisableEnvironmentAlertingBadRequest with default headers values
func NewDisableEnvironmentAlertingBadRequest() *DisableEnvironmentAlertingBadRequest {
	return &DisableEnvironmentAlertingBadRequest{}
}

/* DisableEnvironmentAlertingBadRequest describes a response with status code 400, with default header values.

Alerting already enabled
*/
type DisableEnvironmentAlertingBadRequest struct {
}

func (o *DisableEnvironmentAlertingBadRequest) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/disable][%d] disableEnvironmentAlertingBadRequest ", 400)
}

func (o *DisableEnvironmentAlertingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisableEnvironmentAlertingUnauthorized creates a DisableEnvironmentAlertingUnauthorized with default headers values
func NewDisableEnvironmentAlertingUnauthorized() *DisableEnvironmentAlertingUnauthorized {
	return &DisableEnvironmentAlertingUnauthorized{}
}

/* DisableEnvironmentAlertingUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DisableEnvironmentAlertingUnauthorized struct {
}

func (o *DisableEnvironmentAlertingUnauthorized) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/disable][%d] disableEnvironmentAlertingUnauthorized ", 401)
}

func (o *DisableEnvironmentAlertingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisableEnvironmentAlertingForbidden creates a DisableEnvironmentAlertingForbidden with default headers values
func NewDisableEnvironmentAlertingForbidden() *DisableEnvironmentAlertingForbidden {
	return &DisableEnvironmentAlertingForbidden{}
}

/* DisableEnvironmentAlertingForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DisableEnvironmentAlertingForbidden struct {
}

func (o *DisableEnvironmentAlertingForbidden) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/disable][%d] disableEnvironmentAlertingForbidden ", 403)
}

func (o *DisableEnvironmentAlertingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisableEnvironmentAlertingNotFound creates a DisableEnvironmentAlertingNotFound with default headers values
func NewDisableEnvironmentAlertingNotFound() *DisableEnvironmentAlertingNotFound {
	return &DisableEnvironmentAlertingNotFound{}
}

/* DisableEnvironmentAlertingNotFound describes a response with status code 404, with default header values.

Not found
*/
type DisableEnvironmentAlertingNotFound struct {
}

func (o *DisableEnvironmentAlertingNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/disable][%d] disableEnvironmentAlertingNotFound ", 404)
}

func (o *DisableEnvironmentAlertingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDisableEnvironmentAlertingInternalServerError creates a DisableEnvironmentAlertingInternalServerError with default headers values
func NewDisableEnvironmentAlertingInternalServerError() *DisableEnvironmentAlertingInternalServerError {
	return &DisableEnvironmentAlertingInternalServerError{}
}

/* DisableEnvironmentAlertingInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DisableEnvironmentAlertingInternalServerError struct {
}

func (o *DisableEnvironmentAlertingInternalServerError) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/alerting/disable][%d] disableEnvironmentAlertingInternalServerError ", 500)
}

func (o *DisableEnvironmentAlertingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
