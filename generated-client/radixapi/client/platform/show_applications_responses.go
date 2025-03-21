// Code generated by go-swagger; DO NOT EDIT.

package platform

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

// ShowApplicationsReader is a Reader for the ShowApplications structure.
type ShowApplicationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowApplicationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewShowApplicationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewShowApplicationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewShowApplicationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewShowApplicationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewShowApplicationsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewShowApplicationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /applications] showApplications", response, response.Code())
	}
}

// NewShowApplicationsOK creates a ShowApplicationsOK with default headers values
func NewShowApplicationsOK() *ShowApplicationsOK {
	return &ShowApplicationsOK{}
}

/*
ShowApplicationsOK describes a response with status code 200, with default header values.

Successful operation
*/
type ShowApplicationsOK struct {
	Payload []*models.ApplicationSummary
}

// IsSuccess returns true when this show applications o k response has a 2xx status code
func (o *ShowApplicationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this show applications o k response has a 3xx status code
func (o *ShowApplicationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this show applications o k response has a 4xx status code
func (o *ShowApplicationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this show applications o k response has a 5xx status code
func (o *ShowApplicationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this show applications o k response a status code equal to that given
func (o *ShowApplicationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the show applications o k response
func (o *ShowApplicationsOK) Code() int {
	return 200
}

func (o *ShowApplicationsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /applications][%d] showApplicationsOK %s", 200, payload)
}

func (o *ShowApplicationsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /applications][%d] showApplicationsOK %s", 200, payload)
}

func (o *ShowApplicationsOK) GetPayload() []*models.ApplicationSummary {
	return o.Payload
}

func (o *ShowApplicationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowApplicationsUnauthorized creates a ShowApplicationsUnauthorized with default headers values
func NewShowApplicationsUnauthorized() *ShowApplicationsUnauthorized {
	return &ShowApplicationsUnauthorized{}
}

/*
ShowApplicationsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ShowApplicationsUnauthorized struct {
}

// IsSuccess returns true when this show applications unauthorized response has a 2xx status code
func (o *ShowApplicationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this show applications unauthorized response has a 3xx status code
func (o *ShowApplicationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this show applications unauthorized response has a 4xx status code
func (o *ShowApplicationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this show applications unauthorized response has a 5xx status code
func (o *ShowApplicationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this show applications unauthorized response a status code equal to that given
func (o *ShowApplicationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the show applications unauthorized response
func (o *ShowApplicationsUnauthorized) Code() int {
	return 401
}

func (o *ShowApplicationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications][%d] showApplicationsUnauthorized", 401)
}

func (o *ShowApplicationsUnauthorized) String() string {
	return fmt.Sprintf("[GET /applications][%d] showApplicationsUnauthorized", 401)
}

func (o *ShowApplicationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewShowApplicationsForbidden creates a ShowApplicationsForbidden with default headers values
func NewShowApplicationsForbidden() *ShowApplicationsForbidden {
	return &ShowApplicationsForbidden{}
}

/*
ShowApplicationsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ShowApplicationsForbidden struct {
}

// IsSuccess returns true when this show applications forbidden response has a 2xx status code
func (o *ShowApplicationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this show applications forbidden response has a 3xx status code
func (o *ShowApplicationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this show applications forbidden response has a 4xx status code
func (o *ShowApplicationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this show applications forbidden response has a 5xx status code
func (o *ShowApplicationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this show applications forbidden response a status code equal to that given
func (o *ShowApplicationsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the show applications forbidden response
func (o *ShowApplicationsForbidden) Code() int {
	return 403
}

func (o *ShowApplicationsForbidden) Error() string {
	return fmt.Sprintf("[GET /applications][%d] showApplicationsForbidden", 403)
}

func (o *ShowApplicationsForbidden) String() string {
	return fmt.Sprintf("[GET /applications][%d] showApplicationsForbidden", 403)
}

func (o *ShowApplicationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewShowApplicationsNotFound creates a ShowApplicationsNotFound with default headers values
func NewShowApplicationsNotFound() *ShowApplicationsNotFound {
	return &ShowApplicationsNotFound{}
}

/*
ShowApplicationsNotFound describes a response with status code 404, with default header values.

Not found
*/
type ShowApplicationsNotFound struct {
}

// IsSuccess returns true when this show applications not found response has a 2xx status code
func (o *ShowApplicationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this show applications not found response has a 3xx status code
func (o *ShowApplicationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this show applications not found response has a 4xx status code
func (o *ShowApplicationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this show applications not found response has a 5xx status code
func (o *ShowApplicationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this show applications not found response a status code equal to that given
func (o *ShowApplicationsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the show applications not found response
func (o *ShowApplicationsNotFound) Code() int {
	return 404
}

func (o *ShowApplicationsNotFound) Error() string {
	return fmt.Sprintf("[GET /applications][%d] showApplicationsNotFound", 404)
}

func (o *ShowApplicationsNotFound) String() string {
	return fmt.Sprintf("[GET /applications][%d] showApplicationsNotFound", 404)
}

func (o *ShowApplicationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewShowApplicationsConflict creates a ShowApplicationsConflict with default headers values
func NewShowApplicationsConflict() *ShowApplicationsConflict {
	return &ShowApplicationsConflict{}
}

/*
ShowApplicationsConflict describes a response with status code 409, with default header values.

Conflict
*/
type ShowApplicationsConflict struct {
}

// IsSuccess returns true when this show applications conflict response has a 2xx status code
func (o *ShowApplicationsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this show applications conflict response has a 3xx status code
func (o *ShowApplicationsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this show applications conflict response has a 4xx status code
func (o *ShowApplicationsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this show applications conflict response has a 5xx status code
func (o *ShowApplicationsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this show applications conflict response a status code equal to that given
func (o *ShowApplicationsConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the show applications conflict response
func (o *ShowApplicationsConflict) Code() int {
	return 409
}

func (o *ShowApplicationsConflict) Error() string {
	return fmt.Sprintf("[GET /applications][%d] showApplicationsConflict", 409)
}

func (o *ShowApplicationsConflict) String() string {
	return fmt.Sprintf("[GET /applications][%d] showApplicationsConflict", 409)
}

func (o *ShowApplicationsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewShowApplicationsInternalServerError creates a ShowApplicationsInternalServerError with default headers values
func NewShowApplicationsInternalServerError() *ShowApplicationsInternalServerError {
	return &ShowApplicationsInternalServerError{}
}

/*
ShowApplicationsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ShowApplicationsInternalServerError struct {
}

// IsSuccess returns true when this show applications internal server error response has a 2xx status code
func (o *ShowApplicationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this show applications internal server error response has a 3xx status code
func (o *ShowApplicationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this show applications internal server error response has a 4xx status code
func (o *ShowApplicationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this show applications internal server error response has a 5xx status code
func (o *ShowApplicationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this show applications internal server error response a status code equal to that given
func (o *ShowApplicationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the show applications internal server error response
func (o *ShowApplicationsInternalServerError) Code() int {
	return 500
}

func (o *ShowApplicationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /applications][%d] showApplicationsInternalServerError", 500)
}

func (o *ShowApplicationsInternalServerError) String() string {
	return fmt.Sprintf("[GET /applications][%d] showApplicationsInternalServerError", 500)
}

func (o *ShowApplicationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
