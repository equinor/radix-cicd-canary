// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RestartBatchReader is a Reader for the RestartBatch structure.
type RestartBatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestartBatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRestartBatchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRestartBatchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRestartBatchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRestartBatchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRestartBatchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRestartBatchNoContent creates a RestartBatchNoContent with default headers values
func NewRestartBatchNoContent() *RestartBatchNoContent {
	return &RestartBatchNoContent{}
}

/* RestartBatchNoContent describes a response with status code 204, with default header values.

Success
*/
type RestartBatchNoContent struct {
}

func (o *RestartBatchNoContent) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/batches/{batchName}/restart][%d] restartBatchNoContent ", 204)
}

func (o *RestartBatchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRestartBatchBadRequest creates a RestartBatchBadRequest with default headers values
func NewRestartBatchBadRequest() *RestartBatchBadRequest {
	return &RestartBatchBadRequest{}
}

/* RestartBatchBadRequest describes a response with status code 400, with default header values.

Invalid batch
*/
type RestartBatchBadRequest struct {
}

func (o *RestartBatchBadRequest) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/batches/{batchName}/restart][%d] restartBatchBadRequest ", 400)
}

func (o *RestartBatchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRestartBatchUnauthorized creates a RestartBatchUnauthorized with default headers values
func NewRestartBatchUnauthorized() *RestartBatchUnauthorized {
	return &RestartBatchUnauthorized{}
}

/* RestartBatchUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RestartBatchUnauthorized struct {
}

func (o *RestartBatchUnauthorized) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/batches/{batchName}/restart][%d] restartBatchUnauthorized ", 401)
}

func (o *RestartBatchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRestartBatchForbidden creates a RestartBatchForbidden with default headers values
func NewRestartBatchForbidden() *RestartBatchForbidden {
	return &RestartBatchForbidden{}
}

/* RestartBatchForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RestartBatchForbidden struct {
}

func (o *RestartBatchForbidden) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/batches/{batchName}/restart][%d] restartBatchForbidden ", 403)
}

func (o *RestartBatchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRestartBatchNotFound creates a RestartBatchNotFound with default headers values
func NewRestartBatchNotFound() *RestartBatchNotFound {
	return &RestartBatchNotFound{}
}

/* RestartBatchNotFound describes a response with status code 404, with default header values.

Not found
*/
type RestartBatchNotFound struct {
}

func (o *RestartBatchNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/batches/{batchName}/restart][%d] restartBatchNotFound ", 404)
}

func (o *RestartBatchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
