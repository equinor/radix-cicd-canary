// Code generated by go-swagger; DO NOT EDIT.

package pipeline_job

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

// NewGetPipelineJobStepScanOutputParams creates a new GetPipelineJobStepScanOutputParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPipelineJobStepScanOutputParams() *GetPipelineJobStepScanOutputParams {
	return &GetPipelineJobStepScanOutputParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPipelineJobStepScanOutputParamsWithTimeout creates a new GetPipelineJobStepScanOutputParams object
// with the ability to set a timeout on a request.
func NewGetPipelineJobStepScanOutputParamsWithTimeout(timeout time.Duration) *GetPipelineJobStepScanOutputParams {
	return &GetPipelineJobStepScanOutputParams{
		timeout: timeout,
	}
}

// NewGetPipelineJobStepScanOutputParamsWithContext creates a new GetPipelineJobStepScanOutputParams object
// with the ability to set a context for a request.
func NewGetPipelineJobStepScanOutputParamsWithContext(ctx context.Context) *GetPipelineJobStepScanOutputParams {
	return &GetPipelineJobStepScanOutputParams{
		Context: ctx,
	}
}

// NewGetPipelineJobStepScanOutputParamsWithHTTPClient creates a new GetPipelineJobStepScanOutputParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPipelineJobStepScanOutputParamsWithHTTPClient(client *http.Client) *GetPipelineJobStepScanOutputParams {
	return &GetPipelineJobStepScanOutputParams{
		HTTPClient: client,
	}
}

/* GetPipelineJobStepScanOutputParams contains all the parameters to send to the API endpoint
   for the get pipeline job step scan output operation.

   Typically these are written to a http.Request.
*/
type GetPipelineJobStepScanOutputParams struct {

	/* ImpersonateGroup.

	   Works only with custom setup of cluster. Allow impersonation of test group (Required if Impersonate-User is set)
	*/
	ImpersonateGroup *string

	/* ImpersonateUser.

	   Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)
	*/
	ImpersonateUser *string

	/* AppName.

	   name of Radix application
	*/
	AppName string

	/* JobName.

	   Name of pipeline job
	*/
	JobName string

	/* StepName.

	   Name of the step
	*/
	StepName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get pipeline job step scan output params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPipelineJobStepScanOutputParams) WithDefaults() *GetPipelineJobStepScanOutputParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get pipeline job step scan output params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPipelineJobStepScanOutputParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get pipeline job step scan output params
func (o *GetPipelineJobStepScanOutputParams) WithTimeout(timeout time.Duration) *GetPipelineJobStepScanOutputParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pipeline job step scan output params
func (o *GetPipelineJobStepScanOutputParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pipeline job step scan output params
func (o *GetPipelineJobStepScanOutputParams) WithContext(ctx context.Context) *GetPipelineJobStepScanOutputParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pipeline job step scan output params
func (o *GetPipelineJobStepScanOutputParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pipeline job step scan output params
func (o *GetPipelineJobStepScanOutputParams) WithHTTPClient(client *http.Client) *GetPipelineJobStepScanOutputParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pipeline job step scan output params
func (o *GetPipelineJobStepScanOutputParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the get pipeline job step scan output params
func (o *GetPipelineJobStepScanOutputParams) WithImpersonateGroup(impersonateGroup *string) *GetPipelineJobStepScanOutputParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the get pipeline job step scan output params
func (o *GetPipelineJobStepScanOutputParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the get pipeline job step scan output params
func (o *GetPipelineJobStepScanOutputParams) WithImpersonateUser(impersonateUser *string) *GetPipelineJobStepScanOutputParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the get pipeline job step scan output params
func (o *GetPipelineJobStepScanOutputParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the get pipeline job step scan output params
func (o *GetPipelineJobStepScanOutputParams) WithAppName(appName string) *GetPipelineJobStepScanOutputParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the get pipeline job step scan output params
func (o *GetPipelineJobStepScanOutputParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithJobName adds the jobName to the get pipeline job step scan output params
func (o *GetPipelineJobStepScanOutputParams) WithJobName(jobName string) *GetPipelineJobStepScanOutputParams {
	o.SetJobName(jobName)
	return o
}

// SetJobName adds the jobName to the get pipeline job step scan output params
func (o *GetPipelineJobStepScanOutputParams) SetJobName(jobName string) {
	o.JobName = jobName
}

// WithStepName adds the stepName to the get pipeline job step scan output params
func (o *GetPipelineJobStepScanOutputParams) WithStepName(stepName string) *GetPipelineJobStepScanOutputParams {
	o.SetStepName(stepName)
	return o
}

// SetStepName adds the stepName to the get pipeline job step scan output params
func (o *GetPipelineJobStepScanOutputParams) SetStepName(stepName string) {
	o.StepName = stepName
}

// WriteToRequest writes these params to a swagger request
func (o *GetPipelineJobStepScanOutputParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ImpersonateGroup != nil {

		// header param Impersonate-Group
		if err := r.SetHeaderParam("Impersonate-Group", *o.ImpersonateGroup); err != nil {
			return err
		}
	}

	if o.ImpersonateUser != nil {

		// header param Impersonate-User
		if err := r.SetHeaderParam("Impersonate-User", *o.ImpersonateUser); err != nil {
			return err
		}
	}

	// path param appName
	if err := r.SetPathParam("appName", o.AppName); err != nil {
		return err
	}

	// path param jobName
	if err := r.SetPathParam("jobName", o.JobName); err != nil {
		return err
	}

	// path param stepName
	if err := r.SetPathParam("stepName", o.StepName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}