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

// NewStopJobParams creates a new StopJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStopJobParams() *StopJobParams {
	return &StopJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStopJobParamsWithTimeout creates a new StopJobParams object
// with the ability to set a timeout on a request.
func NewStopJobParamsWithTimeout(timeout time.Duration) *StopJobParams {
	return &StopJobParams{
		timeout: timeout,
	}
}

// NewStopJobParamsWithContext creates a new StopJobParams object
// with the ability to set a context for a request.
func NewStopJobParamsWithContext(ctx context.Context) *StopJobParams {
	return &StopJobParams{
		Context: ctx,
	}
}

// NewStopJobParamsWithHTTPClient creates a new StopJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewStopJobParamsWithHTTPClient(client *http.Client) *StopJobParams {
	return &StopJobParams{
		HTTPClient: client,
	}
}

/*
StopJobParams contains all the parameters to send to the API endpoint

	for the stop job operation.

	Typically these are written to a http.Request.
*/
type StopJobParams struct {

	/* JobName.

	   Name of job
	*/
	JobName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stop job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopJobParams) WithDefaults() *StopJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stop job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stop job params
func (o *StopJobParams) WithTimeout(timeout time.Duration) *StopJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop job params
func (o *StopJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop job params
func (o *StopJobParams) WithContext(ctx context.Context) *StopJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop job params
func (o *StopJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop job params
func (o *StopJobParams) WithHTTPClient(client *http.Client) *StopJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop job params
func (o *StopJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobName adds the jobName to the stop job params
func (o *StopJobParams) WithJobName(jobName string) *StopJobParams {
	o.SetJobName(jobName)
	return o
}

// SetJobName adds the jobName to the stop job params
func (o *StopJobParams) SetJobName(jobName string) {
	o.JobName = jobName
}

// WriteToRequest writes these params to a swagger request
func (o *StopJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param jobName
	if err := r.SetPathParam("jobName", o.JobName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
