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
	"github.com/go-openapi/swag"
)

// NewGetTektonPipelineRunTaskStepLogsParams creates a new GetTektonPipelineRunTaskStepLogsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTektonPipelineRunTaskStepLogsParams() *GetTektonPipelineRunTaskStepLogsParams {
	return &GetTektonPipelineRunTaskStepLogsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTektonPipelineRunTaskStepLogsParamsWithTimeout creates a new GetTektonPipelineRunTaskStepLogsParams object
// with the ability to set a timeout on a request.
func NewGetTektonPipelineRunTaskStepLogsParamsWithTimeout(timeout time.Duration) *GetTektonPipelineRunTaskStepLogsParams {
	return &GetTektonPipelineRunTaskStepLogsParams{
		timeout: timeout,
	}
}

// NewGetTektonPipelineRunTaskStepLogsParamsWithContext creates a new GetTektonPipelineRunTaskStepLogsParams object
// with the ability to set a context for a request.
func NewGetTektonPipelineRunTaskStepLogsParamsWithContext(ctx context.Context) *GetTektonPipelineRunTaskStepLogsParams {
	return &GetTektonPipelineRunTaskStepLogsParams{
		Context: ctx,
	}
}

// NewGetTektonPipelineRunTaskStepLogsParamsWithHTTPClient creates a new GetTektonPipelineRunTaskStepLogsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTektonPipelineRunTaskStepLogsParamsWithHTTPClient(client *http.Client) *GetTektonPipelineRunTaskStepLogsParams {
	return &GetTektonPipelineRunTaskStepLogsParams{
		HTTPClient: client,
	}
}

/* GetTektonPipelineRunTaskStepLogsParams contains all the parameters to send to the API endpoint
   for the get tekton pipeline run task step logs operation.

   Typically these are written to a http.Request.
*/
type GetTektonPipelineRunTaskStepLogsParams struct {

	/* ImpersonateGroup.

	   Works only with custom setup of cluster. Allow impersonation of test group (Required if Impersonate-User is set)
	*/
	ImpersonateGroup []string

	/* ImpersonateUser.

	   Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)
	*/
	ImpersonateUser *string

	/* AppName.

	   name of Radix application
	*/
	AppName string

	/* File.

	   Get log as a file if true

	   Format: boolean
	*/
	File *string

	/* JobName.

	   Name of pipeline job
	*/
	JobName string

	/* Lines.

	   Get log lines (example 1000)

	   Format: number
	*/
	Lines *string

	/* PipelineRunName.

	   Name of pipeline run
	*/
	PipelineRunName string

	/* SinceTime.

	   Get log only from sinceTime (example 2020-03-18T07:20:41+00:00)

	   Format: date-time
	*/
	SinceTime *strfmt.DateTime

	/* StepName.

	   Name of pipeline run task step
	*/
	StepName string

	/* TaskName.

	   Name of pipeline run task
	*/
	TaskName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get tekton pipeline run task step logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTektonPipelineRunTaskStepLogsParams) WithDefaults() *GetTektonPipelineRunTaskStepLogsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get tekton pipeline run task step logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTektonPipelineRunTaskStepLogsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get tekton pipeline run task step logs params
func (o *GetTektonPipelineRunTaskStepLogsParams) WithTimeout(timeout time.Duration) *GetTektonPipelineRunTaskStepLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tekton pipeline run task step logs params
func (o *GetTektonPipelineRunTaskStepLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tekton pipeline run task step logs params
func (o *GetTektonPipelineRunTaskStepLogsParams) WithContext(ctx context.Context) *GetTektonPipelineRunTaskStepLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tekton pipeline run task step logs params
func (o *GetTektonPipelineRunTaskStepLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tekton pipeline run task step logs params
func (o *GetTektonPipelineRunTaskStepLogsParams) WithHTTPClient(client *http.Client) *GetTektonPipelineRunTaskStepLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tekton pipeline run task step logs params
func (o *GetTektonPipelineRunTaskStepLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the get tekton pipeline run task step logs params
func (o *GetTektonPipelineRunTaskStepLogsParams) WithImpersonateGroup(impersonateGroup []string) *GetTektonPipelineRunTaskStepLogsParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the get tekton pipeline run task step logs params
func (o *GetTektonPipelineRunTaskStepLogsParams) SetImpersonateGroup(impersonateGroup []string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the get tekton pipeline run task step logs params
func (o *GetTektonPipelineRunTaskStepLogsParams) WithImpersonateUser(impersonateUser *string) *GetTektonPipelineRunTaskStepLogsParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the get tekton pipeline run task step logs params
func (o *GetTektonPipelineRunTaskStepLogsParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the get tekton pipeline run task step logs params
func (o *GetTektonPipelineRunTaskStepLogsParams) WithAppName(appName string) *GetTektonPipelineRunTaskStepLogsParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the get tekton pipeline run task step logs params
func (o *GetTektonPipelineRunTaskStepLogsParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithFile adds the file to the get tekton pipeline run task step logs params
func (o *GetTektonPipelineRunTaskStepLogsParams) WithFile(file *string) *GetTektonPipelineRunTaskStepLogsParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the get tekton pipeline run task step logs params
func (o *GetTektonPipelineRunTaskStepLogsParams) SetFile(file *string) {
	o.File = file
}

// WithJobName adds the jobName to the get tekton pipeline run task step logs params
func (o *GetTektonPipelineRunTaskStepLogsParams) WithJobName(jobName string) *GetTektonPipelineRunTaskStepLogsParams {
	o.SetJobName(jobName)
	return o
}

// SetJobName adds the jobName to the get tekton pipeline run task step logs params
func (o *GetTektonPipelineRunTaskStepLogsParams) SetJobName(jobName string) {
	o.JobName = jobName
}

// WithLines adds the lines to the get tekton pipeline run task step logs params
func (o *GetTektonPipelineRunTaskStepLogsParams) WithLines(lines *string) *GetTektonPipelineRunTaskStepLogsParams {
	o.SetLines(lines)
	return o
}

// SetLines adds the lines to the get tekton pipeline run task step logs params
func (o *GetTektonPipelineRunTaskStepLogsParams) SetLines(lines *string) {
	o.Lines = lines
}

// WithPipelineRunName adds the pipelineRunName to the get tekton pipeline run task step logs params
func (o *GetTektonPipelineRunTaskStepLogsParams) WithPipelineRunName(pipelineRunName string) *GetTektonPipelineRunTaskStepLogsParams {
	o.SetPipelineRunName(pipelineRunName)
	return o
}

// SetPipelineRunName adds the pipelineRunName to the get tekton pipeline run task step logs params
func (o *GetTektonPipelineRunTaskStepLogsParams) SetPipelineRunName(pipelineRunName string) {
	o.PipelineRunName = pipelineRunName
}

// WithSinceTime adds the sinceTime to the get tekton pipeline run task step logs params
func (o *GetTektonPipelineRunTaskStepLogsParams) WithSinceTime(sinceTime *strfmt.DateTime) *GetTektonPipelineRunTaskStepLogsParams {
	o.SetSinceTime(sinceTime)
	return o
}

// SetSinceTime adds the sinceTime to the get tekton pipeline run task step logs params
func (o *GetTektonPipelineRunTaskStepLogsParams) SetSinceTime(sinceTime *strfmt.DateTime) {
	o.SinceTime = sinceTime
}

// WithStepName adds the stepName to the get tekton pipeline run task step logs params
func (o *GetTektonPipelineRunTaskStepLogsParams) WithStepName(stepName string) *GetTektonPipelineRunTaskStepLogsParams {
	o.SetStepName(stepName)
	return o
}

// SetStepName adds the stepName to the get tekton pipeline run task step logs params
func (o *GetTektonPipelineRunTaskStepLogsParams) SetStepName(stepName string) {
	o.StepName = stepName
}

// WithTaskName adds the taskName to the get tekton pipeline run task step logs params
func (o *GetTektonPipelineRunTaskStepLogsParams) WithTaskName(taskName string) *GetTektonPipelineRunTaskStepLogsParams {
	o.SetTaskName(taskName)
	return o
}

// SetTaskName adds the taskName to the get tekton pipeline run task step logs params
func (o *GetTektonPipelineRunTaskStepLogsParams) SetTaskName(taskName string) {
	o.TaskName = taskName
}

// WriteToRequest writes these params to a swagger request
func (o *GetTektonPipelineRunTaskStepLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ImpersonateGroup != nil {

		// binding items for Impersonate-Group
		joinedImpersonateGroup := o.bindParamImpersonateGroup(reg)

		// header array param Impersonate-Group
		if len(joinedImpersonateGroup) > 0 {
			if err := r.SetHeaderParam("Impersonate-Group", joinedImpersonateGroup[0]); err != nil {
				return err
			}
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

	if o.File != nil {

		// query param file
		var qrFile string

		if o.File != nil {
			qrFile = *o.File
		}
		qFile := qrFile
		if qFile != "" {

			if err := r.SetQueryParam("file", qFile); err != nil {
				return err
			}
		}
	}

	// path param jobName
	if err := r.SetPathParam("jobName", o.JobName); err != nil {
		return err
	}

	if o.Lines != nil {

		// query param lines
		var qrLines string

		if o.Lines != nil {
			qrLines = *o.Lines
		}
		qLines := qrLines
		if qLines != "" {

			if err := r.SetQueryParam("lines", qLines); err != nil {
				return err
			}
		}
	}

	// path param pipelineRunName
	if err := r.SetPathParam("pipelineRunName", o.PipelineRunName); err != nil {
		return err
	}

	if o.SinceTime != nil {

		// query param sinceTime
		var qrSinceTime strfmt.DateTime

		if o.SinceTime != nil {
			qrSinceTime = *o.SinceTime
		}
		qSinceTime := qrSinceTime.String()
		if qSinceTime != "" {

			if err := r.SetQueryParam("sinceTime", qSinceTime); err != nil {
				return err
			}
		}
	}

	// path param stepName
	if err := r.SetPathParam("stepName", o.StepName); err != nil {
		return err
	}

	// path param taskName
	if err := r.SetPathParam("taskName", o.TaskName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetTektonPipelineRunTaskStepLogs binds the parameter Impersonate-Group
func (o *GetTektonPipelineRunTaskStepLogsParams) bindParamImpersonateGroup(formats strfmt.Registry) []string {
	impersonateGroupIR := o.ImpersonateGroup

	var impersonateGroupIC []string
	for _, impersonateGroupIIR := range impersonateGroupIR { // explode []string

		impersonateGroupIIV := impersonateGroupIIR // string as string
		impersonateGroupIC = append(impersonateGroupIC, impersonateGroupIIV)
	}

	// items.CollectionFormat: ""
	impersonateGroupIS := swag.JoinByFormat(impersonateGroupIC, "")

	return impersonateGroupIS
}
