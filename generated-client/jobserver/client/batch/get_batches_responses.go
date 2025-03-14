// Code generated by go-swagger; DO NOT EDIT.

package batch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/equinor/radix-cicd-canary/generated-client/jobserver/models"
)

// GetBatchesReader is a Reader for the GetBatches structure.
type GetBatchesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBatchesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBatchesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetBatchesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /batches/] getBatches", response, response.Code())
	}
}

// NewGetBatchesOK creates a GetBatchesOK with default headers values
func NewGetBatchesOK() *GetBatchesOK {
	return &GetBatchesOK{}
}

/*
GetBatchesOK describes a response with status code 200, with default header values.

Successful get batches
*/
type GetBatchesOK struct {
	Payload []*models.BatchStatus
}

// IsSuccess returns true when this get batches o k response has a 2xx status code
func (o *GetBatchesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get batches o k response has a 3xx status code
func (o *GetBatchesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get batches o k response has a 4xx status code
func (o *GetBatchesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get batches o k response has a 5xx status code
func (o *GetBatchesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get batches o k response a status code equal to that given
func (o *GetBatchesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get batches o k response
func (o *GetBatchesOK) Code() int {
	return 200
}

func (o *GetBatchesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /batches/][%d] getBatchesOK %s", 200, payload)
}

func (o *GetBatchesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /batches/][%d] getBatchesOK %s", 200, payload)
}

func (o *GetBatchesOK) GetPayload() []*models.BatchStatus {
	return o.Payload
}

func (o *GetBatchesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBatchesInternalServerError creates a GetBatchesInternalServerError with default headers values
func NewGetBatchesInternalServerError() *GetBatchesInternalServerError {
	return &GetBatchesInternalServerError{}
}

/*
GetBatchesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetBatchesInternalServerError struct {
	Payload *models.Status
}

// IsSuccess returns true when this get batches internal server error response has a 2xx status code
func (o *GetBatchesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get batches internal server error response has a 3xx status code
func (o *GetBatchesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get batches internal server error response has a 4xx status code
func (o *GetBatchesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get batches internal server error response has a 5xx status code
func (o *GetBatchesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get batches internal server error response a status code equal to that given
func (o *GetBatchesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get batches internal server error response
func (o *GetBatchesInternalServerError) Code() int {
	return 500
}

func (o *GetBatchesInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /batches/][%d] getBatchesInternalServerError %s", 500, payload)
}

func (o *GetBatchesInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /batches/][%d] getBatchesInternalServerError %s", 500, payload)
}

func (o *GetBatchesInternalServerError) GetPayload() *models.Status {
	return o.Payload
}

func (o *GetBatchesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
