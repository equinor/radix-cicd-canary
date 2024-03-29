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

// NewGetBatchesParams creates a new GetBatchesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBatchesParams() *GetBatchesParams {
	return &GetBatchesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBatchesParamsWithTimeout creates a new GetBatchesParams object
// with the ability to set a timeout on a request.
func NewGetBatchesParamsWithTimeout(timeout time.Duration) *GetBatchesParams {
	return &GetBatchesParams{
		timeout: timeout,
	}
}

// NewGetBatchesParamsWithContext creates a new GetBatchesParams object
// with the ability to set a context for a request.
func NewGetBatchesParamsWithContext(ctx context.Context) *GetBatchesParams {
	return &GetBatchesParams{
		Context: ctx,
	}
}

// NewGetBatchesParamsWithHTTPClient creates a new GetBatchesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBatchesParamsWithHTTPClient(client *http.Client) *GetBatchesParams {
	return &GetBatchesParams{
		HTTPClient: client,
	}
}

/*
GetBatchesParams contains all the parameters to send to the API endpoint

	for the get batches operation.

	Typically these are written to a http.Request.
*/
type GetBatchesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get batches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBatchesParams) WithDefaults() *GetBatchesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get batches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBatchesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get batches params
func (o *GetBatchesParams) WithTimeout(timeout time.Duration) *GetBatchesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get batches params
func (o *GetBatchesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get batches params
func (o *GetBatchesParams) WithContext(ctx context.Context) *GetBatchesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get batches params
func (o *GetBatchesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get batches params
func (o *GetBatchesParams) WithHTTPClient(client *http.Client) *GetBatchesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get batches params
func (o *GetBatchesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetBatchesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
