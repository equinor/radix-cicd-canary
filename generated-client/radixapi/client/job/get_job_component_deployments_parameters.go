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
	"github.com/go-openapi/swag"
)

// NewGetJobComponentDeploymentsParams creates a new GetJobComponentDeploymentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetJobComponentDeploymentsParams() *GetJobComponentDeploymentsParams {
	return &GetJobComponentDeploymentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetJobComponentDeploymentsParamsWithTimeout creates a new GetJobComponentDeploymentsParams object
// with the ability to set a timeout on a request.
func NewGetJobComponentDeploymentsParamsWithTimeout(timeout time.Duration) *GetJobComponentDeploymentsParams {
	return &GetJobComponentDeploymentsParams{
		timeout: timeout,
	}
}

// NewGetJobComponentDeploymentsParamsWithContext creates a new GetJobComponentDeploymentsParams object
// with the ability to set a context for a request.
func NewGetJobComponentDeploymentsParamsWithContext(ctx context.Context) *GetJobComponentDeploymentsParams {
	return &GetJobComponentDeploymentsParams{
		Context: ctx,
	}
}

// NewGetJobComponentDeploymentsParamsWithHTTPClient creates a new GetJobComponentDeploymentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetJobComponentDeploymentsParamsWithHTTPClient(client *http.Client) *GetJobComponentDeploymentsParams {
	return &GetJobComponentDeploymentsParams{
		HTTPClient: client,
	}
}

/* GetJobComponentDeploymentsParams contains all the parameters to send to the API endpoint
   for the get job component deployments operation.

   Typically these are written to a http.Request.
*/
type GetJobComponentDeploymentsParams struct {

	/* ImpersonateGroup.

	   Works only with custom setup of cluster. Allow impersonation of test group (Required if Impersonate-User is set)
	*/
	ImpersonateGroup []string

	/* ImpersonateUser.

	   Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)
	*/
	ImpersonateUser *string

	/* AppName.

	   Name of application
	*/
	AppName string

	/* EnvName.

	   Name of environment
	*/
	EnvName string

	/* JobComponentName.

	   Name of job-component
	*/
	JobComponentName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get job component deployments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetJobComponentDeploymentsParams) WithDefaults() *GetJobComponentDeploymentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get job component deployments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetJobComponentDeploymentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get job component deployments params
func (o *GetJobComponentDeploymentsParams) WithTimeout(timeout time.Duration) *GetJobComponentDeploymentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get job component deployments params
func (o *GetJobComponentDeploymentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get job component deployments params
func (o *GetJobComponentDeploymentsParams) WithContext(ctx context.Context) *GetJobComponentDeploymentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get job component deployments params
func (o *GetJobComponentDeploymentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get job component deployments params
func (o *GetJobComponentDeploymentsParams) WithHTTPClient(client *http.Client) *GetJobComponentDeploymentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get job component deployments params
func (o *GetJobComponentDeploymentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the get job component deployments params
func (o *GetJobComponentDeploymentsParams) WithImpersonateGroup(impersonateGroup []string) *GetJobComponentDeploymentsParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the get job component deployments params
func (o *GetJobComponentDeploymentsParams) SetImpersonateGroup(impersonateGroup []string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the get job component deployments params
func (o *GetJobComponentDeploymentsParams) WithImpersonateUser(impersonateUser *string) *GetJobComponentDeploymentsParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the get job component deployments params
func (o *GetJobComponentDeploymentsParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the get job component deployments params
func (o *GetJobComponentDeploymentsParams) WithAppName(appName string) *GetJobComponentDeploymentsParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the get job component deployments params
func (o *GetJobComponentDeploymentsParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithEnvName adds the envName to the get job component deployments params
func (o *GetJobComponentDeploymentsParams) WithEnvName(envName string) *GetJobComponentDeploymentsParams {
	o.SetEnvName(envName)
	return o
}

// SetEnvName adds the envName to the get job component deployments params
func (o *GetJobComponentDeploymentsParams) SetEnvName(envName string) {
	o.EnvName = envName
}

// WithJobComponentName adds the jobComponentName to the get job component deployments params
func (o *GetJobComponentDeploymentsParams) WithJobComponentName(jobComponentName string) *GetJobComponentDeploymentsParams {
	o.SetJobComponentName(jobComponentName)
	return o
}

// SetJobComponentName adds the jobComponentName to the get job component deployments params
func (o *GetJobComponentDeploymentsParams) SetJobComponentName(jobComponentName string) {
	o.JobComponentName = jobComponentName
}

// WriteToRequest writes these params to a swagger request
func (o *GetJobComponentDeploymentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param envName
	if err := r.SetPathParam("envName", o.EnvName); err != nil {
		return err
	}

	// path param jobComponentName
	if err := r.SetPathParam("jobComponentName", o.JobComponentName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetJobComponentDeployments binds the parameter Impersonate-Group
func (o *GetJobComponentDeploymentsParams) bindParamImpersonateGroup(formats strfmt.Registry) []string {
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
