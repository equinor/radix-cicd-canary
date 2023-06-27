// Code generated by go-swagger; DO NOT EDIT.

package environment

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

// NewStopEnvironmentParams creates a new StopEnvironmentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStopEnvironmentParams() *StopEnvironmentParams {
	return &StopEnvironmentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStopEnvironmentParamsWithTimeout creates a new StopEnvironmentParams object
// with the ability to set a timeout on a request.
func NewStopEnvironmentParamsWithTimeout(timeout time.Duration) *StopEnvironmentParams {
	return &StopEnvironmentParams{
		timeout: timeout,
	}
}

// NewStopEnvironmentParamsWithContext creates a new StopEnvironmentParams object
// with the ability to set a context for a request.
func NewStopEnvironmentParamsWithContext(ctx context.Context) *StopEnvironmentParams {
	return &StopEnvironmentParams{
		Context: ctx,
	}
}

// NewStopEnvironmentParamsWithHTTPClient creates a new StopEnvironmentParams object
// with the ability to set a custom HTTPClient for a request.
func NewStopEnvironmentParamsWithHTTPClient(client *http.Client) *StopEnvironmentParams {
	return &StopEnvironmentParams{
		HTTPClient: client,
	}
}

/* StopEnvironmentParams contains all the parameters to send to the API endpoint
   for the stop environment operation.

   Typically these are written to a http.Request.
*/
type StopEnvironmentParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stop environment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopEnvironmentParams) WithDefaults() *StopEnvironmentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stop environment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopEnvironmentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stop environment params
func (o *StopEnvironmentParams) WithTimeout(timeout time.Duration) *StopEnvironmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop environment params
func (o *StopEnvironmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop environment params
func (o *StopEnvironmentParams) WithContext(ctx context.Context) *StopEnvironmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop environment params
func (o *StopEnvironmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop environment params
func (o *StopEnvironmentParams) WithHTTPClient(client *http.Client) *StopEnvironmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop environment params
func (o *StopEnvironmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the stop environment params
func (o *StopEnvironmentParams) WithImpersonateGroup(impersonateGroup []string) *StopEnvironmentParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the stop environment params
func (o *StopEnvironmentParams) SetImpersonateGroup(impersonateGroup []string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the stop environment params
func (o *StopEnvironmentParams) WithImpersonateUser(impersonateUser *string) *StopEnvironmentParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the stop environment params
func (o *StopEnvironmentParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the stop environment params
func (o *StopEnvironmentParams) WithAppName(appName string) *StopEnvironmentParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the stop environment params
func (o *StopEnvironmentParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithEnvName adds the envName to the stop environment params
func (o *StopEnvironmentParams) WithEnvName(envName string) *StopEnvironmentParams {
	o.SetEnvName(envName)
	return o
}

// SetEnvName adds the envName to the stop environment params
func (o *StopEnvironmentParams) SetEnvName(envName string) {
	o.EnvName = envName
}

// WriteToRequest writes these params to a swagger request
func (o *StopEnvironmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamStopEnvironment binds the parameter Impersonate-Group
func (o *StopEnvironmentParams) bindParamImpersonateGroup(formats strfmt.Registry) []string {
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
