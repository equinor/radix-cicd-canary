// Code generated by go-swagger; DO NOT EDIT.

package component

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// LogReader is a Reader for the Log structure.
type LogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewLogNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLogOK creates a LogOK with default headers values
func NewLogOK() *LogOK {
	return &LogOK{}
}

/*LogOK handles this case with default header values.

pod log
*/
type LogOK struct {
}

func (o *LogOK) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/deployments/{deploymentName}/components/{componentName}/replicas/{podName}/logs][%d] logOK ", 200)
}

func (o *LogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLogNotFound creates a LogNotFound with default headers values
func NewLogNotFound() *LogNotFound {
	return &LogNotFound{}
}

/*LogNotFound handles this case with default header values.

Not found
*/
type LogNotFound struct {
}

func (o *LogNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/deployments/{deploymentName}/components/{componentName}/replicas/{podName}/logs][%d] logNotFound ", 404)
}

func (o *LogNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
