// Code generated by go-swagger; DO NOT EDIT.

package batch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewStopBatchParams creates a new StopBatchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStopBatchParams() *StopBatchParams {
	return &StopBatchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStopBatchParamsWithTimeout creates a new StopBatchParams object
// with the ability to set a timeout on a request.
func NewStopBatchParamsWithTimeout(timeout time.Duration) *StopBatchParams {
	return &StopBatchParams{
		timeout: timeout,
	}
}

// NewStopBatchParamsWithContext creates a new StopBatchParams object
// with the ability to set a context for a request.
func NewStopBatchParamsWithContext(ctx context.Context) *StopBatchParams {
	return &StopBatchParams{
		Context: ctx,
	}
}

// NewStopBatchParamsWithHTTPClient creates a new StopBatchParams object
// with the ability to set a custom HTTPClient for a request.
func NewStopBatchParamsWithHTTPClient(client *http.Client) *StopBatchParams {
	return &StopBatchParams{
		HTTPClient: client,
	}
}

/*
StopBatchParams contains all the parameters to send to the API endpoint

	for the stop batch operation.

	Typically these are written to a http.Request.
*/
type StopBatchParams struct {

	/* BatchName.

	   Name of batch
	*/
	BatchName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stop batch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopBatchParams) WithDefaults() *StopBatchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stop batch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopBatchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stop batch params
func (o *StopBatchParams) WithTimeout(timeout time.Duration) *StopBatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop batch params
func (o *StopBatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop batch params
func (o *StopBatchParams) WithContext(ctx context.Context) *StopBatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop batch params
func (o *StopBatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop batch params
func (o *StopBatchParams) WithHTTPClient(client *http.Client) *StopBatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop batch params
func (o *StopBatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBatchName adds the batchName to the stop batch params
func (o *StopBatchParams) WithBatchName(batchName string) *StopBatchParams {
	o.SetBatchName(batchName)
	return o
}

// SetBatchName adds the batchName to the stop batch params
func (o *StopBatchParams) SetBatchName(batchName string) {
	o.BatchName = batchName
}

// WriteToRequest writes these params to a swagger request
func (o *StopBatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param batchName
	if err := r.SetPathParam("batchName", o.BatchName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
