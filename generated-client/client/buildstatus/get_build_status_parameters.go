// Code generated by go-swagger; DO NOT EDIT.

package buildstatus

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

// NewGetBuildStatusParams creates a new GetBuildStatusParams object
// with the default values initialized.
func NewGetBuildStatusParams() *GetBuildStatusParams {
	var (
		pipelineDefault = string("build-deploy")
	)
	return &GetBuildStatusParams{
		Pipeline: &pipelineDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBuildStatusParamsWithTimeout creates a new GetBuildStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBuildStatusParamsWithTimeout(timeout time.Duration) *GetBuildStatusParams {
	var (
		pipelineDefault = string("build-deploy")
	)
	return &GetBuildStatusParams{
		Pipeline: &pipelineDefault,

		timeout: timeout,
	}
}

// NewGetBuildStatusParamsWithContext creates a new GetBuildStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBuildStatusParamsWithContext(ctx context.Context) *GetBuildStatusParams {
	var (
		pipelineDefault = string("build-deploy")
	)
	return &GetBuildStatusParams{
		Pipeline: &pipelineDefault,

		Context: ctx,
	}
}

// NewGetBuildStatusParamsWithHTTPClient creates a new GetBuildStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBuildStatusParamsWithHTTPClient(client *http.Client) *GetBuildStatusParams {
	var (
		pipelineDefault = string("build-deploy")
	)
	return &GetBuildStatusParams{
		Pipeline:   &pipelineDefault,
		HTTPClient: client,
	}
}

/*GetBuildStatusParams contains all the parameters to send to the API endpoint
for the get build status operation typically these are written to a http.Request
*/
type GetBuildStatusParams struct {

	/*AppName
	  name of Radix application

	*/
	AppName string
	/*EnvName
	  name of the environment

	*/
	EnvName string
	/*Pipeline
	  Type of pipeline job to get status for.

	*/
	Pipeline *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get build status params
func (o *GetBuildStatusParams) WithTimeout(timeout time.Duration) *GetBuildStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get build status params
func (o *GetBuildStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get build status params
func (o *GetBuildStatusParams) WithContext(ctx context.Context) *GetBuildStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get build status params
func (o *GetBuildStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get build status params
func (o *GetBuildStatusParams) WithHTTPClient(client *http.Client) *GetBuildStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get build status params
func (o *GetBuildStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the get build status params
func (o *GetBuildStatusParams) WithAppName(appName string) *GetBuildStatusParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the get build status params
func (o *GetBuildStatusParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithEnvName adds the envName to the get build status params
func (o *GetBuildStatusParams) WithEnvName(envName string) *GetBuildStatusParams {
	o.SetEnvName(envName)
	return o
}

// SetEnvName adds the envName to the get build status params
func (o *GetBuildStatusParams) SetEnvName(envName string) {
	o.EnvName = envName
}

// WithPipeline adds the pipeline to the get build status params
func (o *GetBuildStatusParams) WithPipeline(pipeline *string) *GetBuildStatusParams {
	o.SetPipeline(pipeline)
	return o
}

// SetPipeline adds the pipeline to the get build status params
func (o *GetBuildStatusParams) SetPipeline(pipeline *string) {
	o.Pipeline = pipeline
}

// WriteToRequest writes these params to a swagger request
func (o *GetBuildStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appName
	if err := r.SetPathParam("appName", o.AppName); err != nil {
		return err
	}

	// path param envName
	if err := r.SetPathParam("envName", o.EnvName); err != nil {
		return err
	}

	if o.Pipeline != nil {

		// query param pipeline
		var qrPipeline string
		if o.Pipeline != nil {
			qrPipeline = *o.Pipeline
		}
		qPipeline := qrPipeline
		if qPipeline != "" {
			if err := r.SetQueryParam("pipeline", qPipeline); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
