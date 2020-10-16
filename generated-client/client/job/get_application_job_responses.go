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

// GetApplicationJobReader is a Reader for the GetApplicationJob structure.
type GetApplicationJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetApplicationJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetApplicationJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetApplicationJobOK creates a GetApplicationJobOK with default headers values
func NewGetApplicationJobOK() *GetApplicationJobOK {
	return &GetApplicationJobOK{}
}

/*GetApplicationJobOK handles this case with default header values.

Successful get job
*/
type GetApplicationJobOK struct {
	Payload *models.Job
}

func (o *GetApplicationJobOK) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}][%d] getApplicationJobOK  %+v", 200, o.Payload)
}

func (o *GetApplicationJobOK) GetPayload() *models.Job {
	return o.Payload
}

func (o *GetApplicationJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Job)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationJobUnauthorized creates a GetApplicationJobUnauthorized with default headers values
func NewGetApplicationJobUnauthorized() *GetApplicationJobUnauthorized {
	return &GetApplicationJobUnauthorized{}
}

/*GetApplicationJobUnauthorized handles this case with default header values.

Unauthorized
*/
type GetApplicationJobUnauthorized struct {
}

func (o *GetApplicationJobUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}][%d] getApplicationJobUnauthorized ", 401)
}

func (o *GetApplicationJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationJobNotFound creates a GetApplicationJobNotFound with default headers values
func NewGetApplicationJobNotFound() *GetApplicationJobNotFound {
	return &GetApplicationJobNotFound{}
}

/*GetApplicationJobNotFound handles this case with default header values.

Not found
*/
type GetApplicationJobNotFound struct {
}

func (o *GetApplicationJobNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/jobs/{jobName}][%d] getApplicationJobNotFound ", 404)
}

func (o *GetApplicationJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
