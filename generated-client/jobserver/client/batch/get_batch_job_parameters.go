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

// NewGetBatchJobParams creates a new GetBatchJobParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBatchJobParams() *GetBatchJobParams {
	return &GetBatchJobParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBatchJobParamsWithTimeout creates a new GetBatchJobParams object
// with the ability to set a timeout on a request.
func NewGetBatchJobParamsWithTimeout(timeout time.Duration) *GetBatchJobParams {
	return &GetBatchJobParams{
		timeout: timeout,
	}
}

// NewGetBatchJobParamsWithContext creates a new GetBatchJobParams object
// with the ability to set a context for a request.
func NewGetBatchJobParamsWithContext(ctx context.Context) *GetBatchJobParams {
	return &GetBatchJobParams{
		Context: ctx,
	}
}

// NewGetBatchJobParamsWithHTTPClient creates a new GetBatchJobParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBatchJobParamsWithHTTPClient(client *http.Client) *GetBatchJobParams {
	return &GetBatchJobParams{
		HTTPClient: client,
	}
}

/*
GetBatchJobParams contains all the parameters to send to the API endpoint

	for the get batch job operation.

	Typically these are written to a http.Request.
*/
type GetBatchJobParams struct {

	/* BatchName.

	   Name of batch
	*/
	BatchName string

	/* JobName.

	   Name of job
	*/
	JobName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get batch job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBatchJobParams) WithDefaults() *GetBatchJobParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get batch job params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBatchJobParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get batch job params
func (o *GetBatchJobParams) WithTimeout(timeout time.Duration) *GetBatchJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get batch job params
func (o *GetBatchJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get batch job params
func (o *GetBatchJobParams) WithContext(ctx context.Context) *GetBatchJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get batch job params
func (o *GetBatchJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get batch job params
func (o *GetBatchJobParams) WithHTTPClient(client *http.Client) *GetBatchJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get batch job params
func (o *GetBatchJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBatchName adds the batchName to the get batch job params
func (o *GetBatchJobParams) WithBatchName(batchName string) *GetBatchJobParams {
	o.SetBatchName(batchName)
	return o
}

// SetBatchName adds the batchName to the get batch job params
func (o *GetBatchJobParams) SetBatchName(batchName string) {
	o.BatchName = batchName
}

// WithJobName adds the jobName to the get batch job params
func (o *GetBatchJobParams) WithJobName(jobName string) *GetBatchJobParams {
	o.SetJobName(jobName)
	return o
}

// SetJobName adds the jobName to the get batch job params
func (o *GetBatchJobParams) SetJobName(jobName string) {
	o.JobName = jobName
}

// WriteToRequest writes these params to a swagger request
func (o *GetBatchJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param batchName
	if err := r.SetPathParam("batchName", o.BatchName); err != nil {
		return err
	}

	// path param jobName
	if err := r.SetPathParam("jobName", o.JobName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
