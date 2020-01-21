// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/equinor/radix-cicd-canary/generated-client/models"
)

// NewTriggerPipelineBuildDeployParams creates a new TriggerPipelineBuildDeployParams object
// with the default values initialized.
func NewTriggerPipelineBuildDeployParams() *TriggerPipelineBuildDeployParams {
	var ()
	return &TriggerPipelineBuildDeployParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTriggerPipelineBuildDeployParamsWithTimeout creates a new TriggerPipelineBuildDeployParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTriggerPipelineBuildDeployParamsWithTimeout(timeout time.Duration) *TriggerPipelineBuildDeployParams {
	var ()
	return &TriggerPipelineBuildDeployParams{

		timeout: timeout,
	}
}

// NewTriggerPipelineBuildDeployParamsWithContext creates a new TriggerPipelineBuildDeployParams object
// with the default values initialized, and the ability to set a context for a request
func NewTriggerPipelineBuildDeployParamsWithContext(ctx context.Context) *TriggerPipelineBuildDeployParams {
	var ()
	return &TriggerPipelineBuildDeployParams{

		Context: ctx,
	}
}

// NewTriggerPipelineBuildDeployParamsWithHTTPClient creates a new TriggerPipelineBuildDeployParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTriggerPipelineBuildDeployParamsWithHTTPClient(client *http.Client) *TriggerPipelineBuildDeployParams {
	var ()
	return &TriggerPipelineBuildDeployParams{
		HTTPClient: client,
	}
}

/*TriggerPipelineBuildDeployParams contains all the parameters to send to the API endpoint
for the trigger pipeline build deploy operation typically these are written to a http.Request
*/
type TriggerPipelineBuildDeployParams struct {

	/*ImpersonateGroup
	  Works only with custom setup of cluster. Allow impersonation of test group (Required if Impersonate-User is set)

	*/
	ImpersonateGroup *string
	/*ImpersonateUser
	  Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)

	*/
	ImpersonateUser *string
	/*PipelineParametersBuild
	  Pipeline parameters

	*/
	PipelineParametersBuild *models.PipelineParametersBuild
	/*AppName
	  Name of application

	*/
	AppName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the trigger pipeline build deploy params
func (o *TriggerPipelineBuildDeployParams) WithTimeout(timeout time.Duration) *TriggerPipelineBuildDeployParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the trigger pipeline build deploy params
func (o *TriggerPipelineBuildDeployParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the trigger pipeline build deploy params
func (o *TriggerPipelineBuildDeployParams) WithContext(ctx context.Context) *TriggerPipelineBuildDeployParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the trigger pipeline build deploy params
func (o *TriggerPipelineBuildDeployParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the trigger pipeline build deploy params
func (o *TriggerPipelineBuildDeployParams) WithHTTPClient(client *http.Client) *TriggerPipelineBuildDeployParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the trigger pipeline build deploy params
func (o *TriggerPipelineBuildDeployParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the trigger pipeline build deploy params
func (o *TriggerPipelineBuildDeployParams) WithImpersonateGroup(impersonateGroup *string) *TriggerPipelineBuildDeployParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the trigger pipeline build deploy params
func (o *TriggerPipelineBuildDeployParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the trigger pipeline build deploy params
func (o *TriggerPipelineBuildDeployParams) WithImpersonateUser(impersonateUser *string) *TriggerPipelineBuildDeployParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the trigger pipeline build deploy params
func (o *TriggerPipelineBuildDeployParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithPipelineParametersBuild adds the pipelineParametersBuild to the trigger pipeline build deploy params
func (o *TriggerPipelineBuildDeployParams) WithPipelineParametersBuild(pipelineParametersBuild *models.PipelineParametersBuild) *TriggerPipelineBuildDeployParams {
	o.SetPipelineParametersBuild(pipelineParametersBuild)
	return o
}

// SetPipelineParametersBuild adds the pipelineParametersBuild to the trigger pipeline build deploy params
func (o *TriggerPipelineBuildDeployParams) SetPipelineParametersBuild(pipelineParametersBuild *models.PipelineParametersBuild) {
	o.PipelineParametersBuild = pipelineParametersBuild
}

// WithAppName adds the appName to the trigger pipeline build deploy params
func (o *TriggerPipelineBuildDeployParams) WithAppName(appName string) *TriggerPipelineBuildDeployParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the trigger pipeline build deploy params
func (o *TriggerPipelineBuildDeployParams) SetAppName(appName string) {
	o.AppName = appName
}

// WriteToRequest writes these params to a swagger request
func (o *TriggerPipelineBuildDeployParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.PipelineParametersBuild != nil {
		if err := r.SetBodyParam(o.PipelineParametersBuild); err != nil {
			return err
		}
	}

	// path param appName
	if err := r.SetPathParam("appName", o.AppName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
