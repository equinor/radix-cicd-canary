// Code generated by go-swagger; DO NOT EDIT.

package environment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ChangeComponentSecretReader is a Reader for the ChangeComponentSecret structure.
type ChangeComponentSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeComponentSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeComponentSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewChangeComponentSecretBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewChangeComponentSecretUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewChangeComponentSecretForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewChangeComponentSecretNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewChangeComponentSecretConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewChangeComponentSecretInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewChangeComponentSecretOK creates a ChangeComponentSecretOK with default headers values
func NewChangeComponentSecretOK() *ChangeComponentSecretOK {
	return &ChangeComponentSecretOK{}
}

/*ChangeComponentSecretOK handles this case with default header values.

success
*/
type ChangeComponentSecretOK struct {
}

func (o *ChangeComponentSecretOK) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}/environments/{envName}/components/{componentName}/secrets/{secretName}][%d] changeComponentSecretOK ", 200)
}

func (o *ChangeComponentSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeComponentSecretBadRequest creates a ChangeComponentSecretBadRequest with default headers values
func NewChangeComponentSecretBadRequest() *ChangeComponentSecretBadRequest {
	return &ChangeComponentSecretBadRequest{}
}

/*ChangeComponentSecretBadRequest handles this case with default header values.

Invalid application
*/
type ChangeComponentSecretBadRequest struct {
}

func (o *ChangeComponentSecretBadRequest) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}/environments/{envName}/components/{componentName}/secrets/{secretName}][%d] changeComponentSecretBadRequest ", 400)
}

func (o *ChangeComponentSecretBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeComponentSecretUnauthorized creates a ChangeComponentSecretUnauthorized with default headers values
func NewChangeComponentSecretUnauthorized() *ChangeComponentSecretUnauthorized {
	return &ChangeComponentSecretUnauthorized{}
}

/*ChangeComponentSecretUnauthorized handles this case with default header values.

Unauthorized
*/
type ChangeComponentSecretUnauthorized struct {
}

func (o *ChangeComponentSecretUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}/environments/{envName}/components/{componentName}/secrets/{secretName}][%d] changeComponentSecretUnauthorized ", 401)
}

func (o *ChangeComponentSecretUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeComponentSecretForbidden creates a ChangeComponentSecretForbidden with default headers values
func NewChangeComponentSecretForbidden() *ChangeComponentSecretForbidden {
	return &ChangeComponentSecretForbidden{}
}

/*ChangeComponentSecretForbidden handles this case with default header values.

Forbidden
*/
type ChangeComponentSecretForbidden struct {
}

func (o *ChangeComponentSecretForbidden) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}/environments/{envName}/components/{componentName}/secrets/{secretName}][%d] changeComponentSecretForbidden ", 403)
}

func (o *ChangeComponentSecretForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeComponentSecretNotFound creates a ChangeComponentSecretNotFound with default headers values
func NewChangeComponentSecretNotFound() *ChangeComponentSecretNotFound {
	return &ChangeComponentSecretNotFound{}
}

/*ChangeComponentSecretNotFound handles this case with default header values.

Not found
*/
type ChangeComponentSecretNotFound struct {
}

func (o *ChangeComponentSecretNotFound) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}/environments/{envName}/components/{componentName}/secrets/{secretName}][%d] changeComponentSecretNotFound ", 404)
}

func (o *ChangeComponentSecretNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeComponentSecretConflict creates a ChangeComponentSecretConflict with default headers values
func NewChangeComponentSecretConflict() *ChangeComponentSecretConflict {
	return &ChangeComponentSecretConflict{}
}

/*ChangeComponentSecretConflict handles this case with default header values.

Conflict
*/
type ChangeComponentSecretConflict struct {
}

func (o *ChangeComponentSecretConflict) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}/environments/{envName}/components/{componentName}/secrets/{secretName}][%d] changeComponentSecretConflict ", 409)
}

func (o *ChangeComponentSecretConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeComponentSecretInternalServerError creates a ChangeComponentSecretInternalServerError with default headers values
func NewChangeComponentSecretInternalServerError() *ChangeComponentSecretInternalServerError {
	return &ChangeComponentSecretInternalServerError{}
}

/*ChangeComponentSecretInternalServerError handles this case with default header values.

Internal server error
*/
type ChangeComponentSecretInternalServerError struct {
}

func (o *ChangeComponentSecretInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /applications/{appName}/environments/{envName}/components/{componentName}/secrets/{secretName}][%d] changeComponentSecretInternalServerError ", 500)
}

func (o *ChangeComponentSecretInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
