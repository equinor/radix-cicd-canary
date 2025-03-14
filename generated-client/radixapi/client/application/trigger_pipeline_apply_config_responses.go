// Code generated by go-swagger; DO NOT EDIT.

package application

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

// TriggerPipelineApplyConfigReader is a Reader for the TriggerPipelineApplyConfig structure.
type TriggerPipelineApplyConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TriggerPipelineApplyConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTriggerPipelineApplyConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewTriggerPipelineApplyConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTriggerPipelineApplyConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /applications/{appName}/pipelines/apply-config] triggerPipelineApplyConfig", response, response.Code())
	}
}

// NewTriggerPipelineApplyConfigOK creates a TriggerPipelineApplyConfigOK with default headers values
func NewTriggerPipelineApplyConfigOK() *TriggerPipelineApplyConfigOK {
	return &TriggerPipelineApplyConfigOK{}
}

/*
TriggerPipelineApplyConfigOK describes a response with status code 200, with default header values.

Successful trigger pipeline
*/
type TriggerPipelineApplyConfigOK struct {
	Payload *models.JobSummary
}

// IsSuccess returns true when this trigger pipeline apply config o k response has a 2xx status code
func (o *TriggerPipelineApplyConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this trigger pipeline apply config o k response has a 3xx status code
func (o *TriggerPipelineApplyConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trigger pipeline apply config o k response has a 4xx status code
func (o *TriggerPipelineApplyConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this trigger pipeline apply config o k response has a 5xx status code
func (o *TriggerPipelineApplyConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this trigger pipeline apply config o k response a status code equal to that given
func (o *TriggerPipelineApplyConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the trigger pipeline apply config o k response
func (o *TriggerPipelineApplyConfigOK) Code() int {
	return 200
}

func (o *TriggerPipelineApplyConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/apply-config][%d] triggerPipelineApplyConfigOK %s", 200, payload)
}

func (o *TriggerPipelineApplyConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/apply-config][%d] triggerPipelineApplyConfigOK %s", 200, payload)
}

func (o *TriggerPipelineApplyConfigOK) GetPayload() *models.JobSummary {
	return o.Payload
}

func (o *TriggerPipelineApplyConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTriggerPipelineApplyConfigForbidden creates a TriggerPipelineApplyConfigForbidden with default headers values
func NewTriggerPipelineApplyConfigForbidden() *TriggerPipelineApplyConfigForbidden {
	return &TriggerPipelineApplyConfigForbidden{}
}

/*
TriggerPipelineApplyConfigForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type TriggerPipelineApplyConfigForbidden struct {
}

// IsSuccess returns true when this trigger pipeline apply config forbidden response has a 2xx status code
func (o *TriggerPipelineApplyConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this trigger pipeline apply config forbidden response has a 3xx status code
func (o *TriggerPipelineApplyConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trigger pipeline apply config forbidden response has a 4xx status code
func (o *TriggerPipelineApplyConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this trigger pipeline apply config forbidden response has a 5xx status code
func (o *TriggerPipelineApplyConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this trigger pipeline apply config forbidden response a status code equal to that given
func (o *TriggerPipelineApplyConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the trigger pipeline apply config forbidden response
func (o *TriggerPipelineApplyConfigForbidden) Code() int {
	return 403
}

func (o *TriggerPipelineApplyConfigForbidden) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/apply-config][%d] triggerPipelineApplyConfigForbidden", 403)
}

func (o *TriggerPipelineApplyConfigForbidden) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/apply-config][%d] triggerPipelineApplyConfigForbidden", 403)
}

func (o *TriggerPipelineApplyConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTriggerPipelineApplyConfigNotFound creates a TriggerPipelineApplyConfigNotFound with default headers values
func NewTriggerPipelineApplyConfigNotFound() *TriggerPipelineApplyConfigNotFound {
	return &TriggerPipelineApplyConfigNotFound{}
}

/*
TriggerPipelineApplyConfigNotFound describes a response with status code 404, with default header values.

Not found
*/
type TriggerPipelineApplyConfigNotFound struct {
}

// IsSuccess returns true when this trigger pipeline apply config not found response has a 2xx status code
func (o *TriggerPipelineApplyConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this trigger pipeline apply config not found response has a 3xx status code
func (o *TriggerPipelineApplyConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trigger pipeline apply config not found response has a 4xx status code
func (o *TriggerPipelineApplyConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this trigger pipeline apply config not found response has a 5xx status code
func (o *TriggerPipelineApplyConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this trigger pipeline apply config not found response a status code equal to that given
func (o *TriggerPipelineApplyConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the trigger pipeline apply config not found response
func (o *TriggerPipelineApplyConfigNotFound) Code() int {
	return 404
}

func (o *TriggerPipelineApplyConfigNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/apply-config][%d] triggerPipelineApplyConfigNotFound", 404)
}

func (o *TriggerPipelineApplyConfigNotFound) String() string {
	return fmt.Sprintf("[POST /applications/{appName}/pipelines/apply-config][%d] triggerPipelineApplyConfigNotFound", 404)
}

func (o *TriggerPipelineApplyConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
