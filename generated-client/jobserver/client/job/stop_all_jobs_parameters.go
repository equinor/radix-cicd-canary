// Code generated by go-swagger; DO NOT EDIT.

package job

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

// NewStopAllJobsParams creates a new StopAllJobsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStopAllJobsParams() *StopAllJobsParams {
	return &StopAllJobsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStopAllJobsParamsWithTimeout creates a new StopAllJobsParams object
// with the ability to set a timeout on a request.
func NewStopAllJobsParamsWithTimeout(timeout time.Duration) *StopAllJobsParams {
	return &StopAllJobsParams{
		timeout: timeout,
	}
}

// NewStopAllJobsParamsWithContext creates a new StopAllJobsParams object
// with the ability to set a context for a request.
func NewStopAllJobsParamsWithContext(ctx context.Context) *StopAllJobsParams {
	return &StopAllJobsParams{
		Context: ctx,
	}
}

// NewStopAllJobsParamsWithHTTPClient creates a new StopAllJobsParams object
// with the ability to set a custom HTTPClient for a request.
func NewStopAllJobsParamsWithHTTPClient(client *http.Client) *StopAllJobsParams {
	return &StopAllJobsParams{
		HTTPClient: client,
	}
}

/*
StopAllJobsParams contains all the parameters to send to the API endpoint

	for the stop all jobs operation.

	Typically these are written to a http.Request.
*/
type StopAllJobsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stop all jobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopAllJobsParams) WithDefaults() *StopAllJobsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stop all jobs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopAllJobsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stop all jobs params
func (o *StopAllJobsParams) WithTimeout(timeout time.Duration) *StopAllJobsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop all jobs params
func (o *StopAllJobsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop all jobs params
func (o *StopAllJobsParams) WithContext(ctx context.Context) *StopAllJobsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop all jobs params
func (o *StopAllJobsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop all jobs params
func (o *StopAllJobsParams) WithHTTPClient(client *http.Client) *StopAllJobsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop all jobs params
func (o *StopAllJobsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StopAllJobsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
