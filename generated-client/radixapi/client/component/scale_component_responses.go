// Code generated by go-swagger; DO NOT EDIT.

package component

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ScaleComponentReader is a Reader for the ScaleComponent structure.
type ScaleComponentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScaleComponentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewScaleComponentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewScaleComponentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewScaleComponentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewScaleComponentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewScaleComponentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewScaleComponentNoContent creates a ScaleComponentNoContent with default headers values
func NewScaleComponentNoContent() *ScaleComponentNoContent {
	return &ScaleComponentNoContent{}
}

/* ScaleComponentNoContent describes a response with status code 204, with default header values.

Success
*/
type ScaleComponentNoContent struct {
}

func (o *ScaleComponentNoContent) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/components/{componentName}/scale/{replicas}][%d] scaleComponentNoContent ", 204)
}

func (o *ScaleComponentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewScaleComponentBadRequest creates a ScaleComponentBadRequest with default headers values
func NewScaleComponentBadRequest() *ScaleComponentBadRequest {
	return &ScaleComponentBadRequest{}
}

/* ScaleComponentBadRequest describes a response with status code 400, with default header values.

Invalid component
*/
type ScaleComponentBadRequest struct {
}

func (o *ScaleComponentBadRequest) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/components/{componentName}/scale/{replicas}][%d] scaleComponentBadRequest ", 400)
}

func (o *ScaleComponentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewScaleComponentUnauthorized creates a ScaleComponentUnauthorized with default headers values
func NewScaleComponentUnauthorized() *ScaleComponentUnauthorized {
	return &ScaleComponentUnauthorized{}
}

/* ScaleComponentUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ScaleComponentUnauthorized struct {
}

func (o *ScaleComponentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/components/{componentName}/scale/{replicas}][%d] scaleComponentUnauthorized ", 401)
}

func (o *ScaleComponentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewScaleComponentForbidden creates a ScaleComponentForbidden with default headers values
func NewScaleComponentForbidden() *ScaleComponentForbidden {
	return &ScaleComponentForbidden{}
}

/* ScaleComponentForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ScaleComponentForbidden struct {
}

func (o *ScaleComponentForbidden) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/components/{componentName}/scale/{replicas}][%d] scaleComponentForbidden ", 403)
}

func (o *ScaleComponentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewScaleComponentNotFound creates a ScaleComponentNotFound with default headers values
func NewScaleComponentNotFound() *ScaleComponentNotFound {
	return &ScaleComponentNotFound{}
}

/* ScaleComponentNotFound describes a response with status code 404, with default header values.

Not found
*/
type ScaleComponentNotFound struct {
}

func (o *ScaleComponentNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/{appName}/environments/{envName}/components/{componentName}/scale/{replicas}][%d] scaleComponentNotFound ", 404)
}

func (o *ScaleComponentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
