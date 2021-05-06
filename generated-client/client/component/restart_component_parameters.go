// Code generated by go-swagger; DO NOT EDIT.

package component

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

// NewRestartComponentParams creates a new RestartComponentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRestartComponentParams() *RestartComponentParams {
	return &RestartComponentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRestartComponentParamsWithTimeout creates a new RestartComponentParams object
// with the ability to set a timeout on a request.
func NewRestartComponentParamsWithTimeout(timeout time.Duration) *RestartComponentParams {
	return &RestartComponentParams{
		timeout: timeout,
	}
}

// NewRestartComponentParamsWithContext creates a new RestartComponentParams object
// with the ability to set a context for a request.
func NewRestartComponentParamsWithContext(ctx context.Context) *RestartComponentParams {
	return &RestartComponentParams{
		Context: ctx,
	}
}

// NewRestartComponentParamsWithHTTPClient creates a new RestartComponentParams object
// with the ability to set a custom HTTPClient for a request.
func NewRestartComponentParamsWithHTTPClient(client *http.Client) *RestartComponentParams {
	return &RestartComponentParams{
		HTTPClient: client,
	}
}

/* RestartComponentParams contains all the parameters to send to the API endpoint
   for the restart component operation.

   Typically these are written to a http.Request.
*/
type RestartComponentParams struct {

	/* ImpersonateGroup.

	   Works only with custom setup of cluster. Allow impersonation of test group (Required if Impersonate-User is set)
	*/
	ImpersonateGroup *string

	/* ImpersonateUser.

	   Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)
	*/
	ImpersonateUser *string

	/* AppName.

	   Name of application
	*/
	AppName string

	/* ComponentName.

	   Name of component
	*/
	ComponentName string

	/* EnvName.

	   Name of environment
	*/
	EnvName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the restart component params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestartComponentParams) WithDefaults() *RestartComponentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the restart component params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestartComponentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the restart component params
func (o *RestartComponentParams) WithTimeout(timeout time.Duration) *RestartComponentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restart component params
func (o *RestartComponentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restart component params
func (o *RestartComponentParams) WithContext(ctx context.Context) *RestartComponentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restart component params
func (o *RestartComponentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restart component params
func (o *RestartComponentParams) WithHTTPClient(client *http.Client) *RestartComponentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restart component params
func (o *RestartComponentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the restart component params
func (o *RestartComponentParams) WithImpersonateGroup(impersonateGroup *string) *RestartComponentParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the restart component params
func (o *RestartComponentParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the restart component params
func (o *RestartComponentParams) WithImpersonateUser(impersonateUser *string) *RestartComponentParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the restart component params
func (o *RestartComponentParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the restart component params
func (o *RestartComponentParams) WithAppName(appName string) *RestartComponentParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the restart component params
func (o *RestartComponentParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithComponentName adds the componentName to the restart component params
func (o *RestartComponentParams) WithComponentName(componentName string) *RestartComponentParams {
	o.SetComponentName(componentName)
	return o
}

// SetComponentName adds the componentName to the restart component params
func (o *RestartComponentParams) SetComponentName(componentName string) {
	o.ComponentName = componentName
}

// WithEnvName adds the envName to the restart component params
func (o *RestartComponentParams) WithEnvName(envName string) *RestartComponentParams {
	o.SetEnvName(envName)
	return o
}

// SetEnvName adds the envName to the restart component params
func (o *RestartComponentParams) SetEnvName(envName string) {
	o.EnvName = envName
}

// WriteToRequest writes these params to a swagger request
func (o *RestartComponentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param componentName
	if err := r.SetPathParam("componentName", o.ComponentName); err != nil {
		return err
	}

	// path param envName
	if err := r.SetPathParam("envName", o.EnvName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
