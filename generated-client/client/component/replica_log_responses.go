// Code generated by go-swagger; DO NOT EDIT.

package component

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ReplicaLogReader is a Reader for the ReplicaLog structure.
type ReplicaLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplicaLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplicaLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewReplicaLogNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplicaLogOK creates a ReplicaLogOK with default headers values
func NewReplicaLogOK() *ReplicaLogOK {
	return &ReplicaLogOK{}
}

/*ReplicaLogOK handles this case with default header values.

pod log
*/
type ReplicaLogOK struct {
	Payload string
}

func (o *ReplicaLogOK) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/components/{componentName}/replicas/{podName}/logs][%d] replicaLogOK  %+v", 200, o.Payload)
}

func (o *ReplicaLogOK) GetPayload() string {
	return o.Payload
}

func (o *ReplicaLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplicaLogNotFound creates a ReplicaLogNotFound with default headers values
func NewReplicaLogNotFound() *ReplicaLogNotFound {
	return &ReplicaLogNotFound{}
}

/*ReplicaLogNotFound handles this case with default header values.

Not found
*/
type ReplicaLogNotFound struct {
}

func (o *ReplicaLogNotFound) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/environments/{envName}/components/{componentName}/replicas/{podName}/logs][%d] replicaLogNotFound ", 404)
}

func (o *ReplicaLogNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
