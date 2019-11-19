// Code generated by go-swagger; DO NOT EDIT.

package deployment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/equinor/radix-cicd-canary/generated-client/models"
)

// GetDeploymentReader is a Reader for the GetDeployment structure.
type GetDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDeploymentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDeploymentOK creates a GetDeploymentOK with default headers values
func NewGetDeploymentOK() *GetDeploymentOK {
	return &GetDeploymentOK{}
}

/*GetDeploymentOK handles this case with default header values.

Successful get deployment
*/
type GetDeploymentOK struct {
	Payload *models.Deployment
}

func (o *GetDeploymentOK) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/deployments/{deploymentName}][%d] getDeploymentOK  %+v", 200, o.Payload)
}

func (o *GetDeploymentOK) GetPayload() *models.Deployment {
	return o.Payload
}

func (o *GetDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Deployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentUnauthorized creates a GetDeploymentUnauthorized with default headers values
func NewGetDeploymentUnauthorized() *GetDeploymentUnauthorized {
	return &GetDeploymentUnauthorized{}
}

/*GetDeploymentUnauthorized handles this case with default header values.

Unauthorized
*/
type GetDeploymentUnauthorized struct {
}

func (o *GetDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/deployments/{deploymentName}][%d] getDeploymentUnauthorized ", 401)
}

func (o *GetDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeploymentNotFound creates a GetDeploymentNotFound with default headers values
func NewGetDeploymentNotFound() *GetDeploymentNotFound {
	return &GetDeploymentNotFound{}
}

/*GetDeploymentNotFound handles this case with default header values.

Not found
*/
type GetDeploymentNotFound struct {
}

func (o *GetDeploymentNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/deployments/{deploymentName}][%d] getDeploymentNotFound ", 404)
}

func (o *GetDeploymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
