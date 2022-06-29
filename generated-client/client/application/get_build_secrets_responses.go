// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/equinor/radix-cicd-canary/generated-client/models"
)

// GetBuildSecretsReader is a Reader for the GetBuildSecrets structure.
type GetBuildSecretsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBuildSecretsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBuildSecretsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetBuildSecretsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBuildSecretsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBuildSecretsOK creates a GetBuildSecretsOK with default headers values
func NewGetBuildSecretsOK() *GetBuildSecretsOK {
	return &GetBuildSecretsOK{}
}

/*GetBuildSecretsOK handles this case with default header values.

Successful operation
*/
type GetBuildSecretsOK struct {
	Payload []*models.BuildSecret
}

func (o *GetBuildSecretsOK) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/buildsecrets][%d] getBuildSecretsOK  %+v", 200, o.Payload)
}

func (o *GetBuildSecretsOK) GetPayload() []*models.BuildSecret {
	return o.Payload
}

func (o *GetBuildSecretsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBuildSecretsUnauthorized creates a GetBuildSecretsUnauthorized with default headers values
func NewGetBuildSecretsUnauthorized() *GetBuildSecretsUnauthorized {
	return &GetBuildSecretsUnauthorized{}
}

/*GetBuildSecretsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetBuildSecretsUnauthorized struct {
}

func (o *GetBuildSecretsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/buildsecrets][%d] getBuildSecretsUnauthorized ", 401)
}

func (o *GetBuildSecretsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBuildSecretsNotFound creates a GetBuildSecretsNotFound with default headers values
func NewGetBuildSecretsNotFound() *GetBuildSecretsNotFound {
	return &GetBuildSecretsNotFound{}
}

/*GetBuildSecretsNotFound handles this case with default header values.

Not found
*/
type GetBuildSecretsNotFound struct {
}

func (o *GetBuildSecretsNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/buildsecrets][%d] getBuildSecretsNotFound ", 404)
}

func (o *GetBuildSecretsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
