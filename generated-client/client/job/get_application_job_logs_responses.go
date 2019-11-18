// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/equinor/radix-cicd-canary/generated-client/models"
)

// GetApplicationJobLogsReader is a Reader for the GetApplicationJobLogs structure.
type GetApplicationJobLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationJobLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationJobLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetApplicationJobLogsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetApplicationJobLogsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApplicationJobLogsOK creates a GetApplicationJobLogsOK with default headers values
func NewGetApplicationJobLogsOK() *GetApplicationJobLogsOK {
	return &GetApplicationJobLogsOK{}
}

/*GetApplicationJobLogsOK handles this case with default header values.

Successful operation
*/
type GetApplicationJobLogsOK struct {
	Payload []*models.StepLog
}

func (o *GetApplicationJobLogsOK) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}/logs][%d] getApplicationJobLogsOK  %+v", 200, o.Payload)
}

func (o *GetApplicationJobLogsOK) GetPayload() []*models.StepLog {
	return o.Payload
}

func (o *GetApplicationJobLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationJobLogsUnauthorized creates a GetApplicationJobLogsUnauthorized with default headers values
func NewGetApplicationJobLogsUnauthorized() *GetApplicationJobLogsUnauthorized {
	return &GetApplicationJobLogsUnauthorized{}
}

/*GetApplicationJobLogsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetApplicationJobLogsUnauthorized struct {
}

func (o *GetApplicationJobLogsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}/logs][%d] getApplicationJobLogsUnauthorized ", 401)
}

func (o *GetApplicationJobLogsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationJobLogsNotFound creates a GetApplicationJobLogsNotFound with default headers values
func NewGetApplicationJobLogsNotFound() *GetApplicationJobLogsNotFound {
	return &GetApplicationJobLogsNotFound{}
}

/*GetApplicationJobLogsNotFound handles this case with default header values.

Not found
*/
type GetApplicationJobLogsNotFound struct {
}

func (o *GetApplicationJobLogsNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}/logs][%d] getApplicationJobLogsNotFound ", 404)
}

func (o *GetApplicationJobLogsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
