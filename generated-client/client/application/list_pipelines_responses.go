// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ListPipelinesReader is a Reader for the ListPipelines structure.
type ListPipelinesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPipelinesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPipelinesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListPipelinesOK creates a ListPipelinesOK with default headers values
func NewListPipelinesOK() *ListPipelinesOK {
	return &ListPipelinesOK{}
}

/*ListPipelinesOK handles this case with default header values.

Successful operation
*/
type ListPipelinesOK struct {
	Payload []string
}

func (o *ListPipelinesOK) Error() string {
	return fmt.Sprintf("[GET /applications/{appName}/pipelines][%d] listPipelinesOK  %+v", 200, o.Payload)
}

func (o *ListPipelinesOK) GetPayload() []string {
	return o.Payload
}

func (o *ListPipelinesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
