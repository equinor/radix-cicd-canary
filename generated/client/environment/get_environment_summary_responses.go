// Code generated by go-swagger; DO NOT EDIT.

package environment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/equinor/radix-cicd-canary-golang/generated/models"
)

// GetEnvironmentSummaryReader is a Reader for the GetEnvironmentSummary structure.
type GetEnvironmentSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEnvironmentSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEnvironmentSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetEnvironmentSummaryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetEnvironmentSummaryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEnvironmentSummaryOK creates a GetEnvironmentSummaryOK with default headers values
func NewGetEnvironmentSummaryOK() *GetEnvironmentSummaryOK {
	return &GetEnvironmentSummaryOK{}
}

/*GetEnvironmentSummaryOK handles this case with default header values.

Successful operation
*/
type GetEnvironmentSummaryOK struct {
	Payload []*models.EnvironmentSummary
}

func (o *GetEnvironmentSummaryOK) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments][%d] getEnvironmentSummaryOK  %+v", 200, o.Payload)
}

func (o *GetEnvironmentSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEnvironmentSummaryUnauthorized creates a GetEnvironmentSummaryUnauthorized with default headers values
func NewGetEnvironmentSummaryUnauthorized() *GetEnvironmentSummaryUnauthorized {
	return &GetEnvironmentSummaryUnauthorized{}
}

/*GetEnvironmentSummaryUnauthorized handles this case with default header values.

Unauthorized
*/
type GetEnvironmentSummaryUnauthorized struct {
}

func (o *GetEnvironmentSummaryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments][%d] getEnvironmentSummaryUnauthorized ", 401)
}

func (o *GetEnvironmentSummaryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEnvironmentSummaryNotFound creates a GetEnvironmentSummaryNotFound with default headers values
func NewGetEnvironmentSummaryNotFound() *GetEnvironmentSummaryNotFound {
	return &GetEnvironmentSummaryNotFound{}
}

/*GetEnvironmentSummaryNotFound handles this case with default header values.

Not found
*/
type GetEnvironmentSummaryNotFound struct {
}

func (o *GetEnvironmentSummaryNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments][%d] getEnvironmentSummaryNotFound ", 404)
}

func (o *GetEnvironmentSummaryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
