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

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteEnvironmentParams creates a new DeleteEnvironmentParams object
// with the default values initialized.
func NewDeleteEnvironmentParams() *DeleteEnvironmentParams {
	var ()
	return &DeleteEnvironmentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEnvironmentParamsWithTimeout creates a new DeleteEnvironmentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteEnvironmentParamsWithTimeout(timeout time.Duration) *DeleteEnvironmentParams {
	var ()
	return &DeleteEnvironmentParams{

		timeout: timeout,
	}
}

// NewDeleteEnvironmentParamsWithContext creates a new DeleteEnvironmentParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteEnvironmentParamsWithContext(ctx context.Context) *DeleteEnvironmentParams {
	var ()
	return &DeleteEnvironmentParams{

		Context: ctx,
	}
}

// NewDeleteEnvironmentParamsWithHTTPClient creates a new DeleteEnvironmentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteEnvironmentParamsWithHTTPClient(client *http.Client) *DeleteEnvironmentParams {
	var ()
	return &DeleteEnvironmentParams{
		HTTPClient: client,
	}
}

/*DeleteEnvironmentParams contains all the parameters to send to the API endpoint
for the delete environment operation typically these are written to a http.Request
*/
type DeleteEnvironmentParams struct {

	/*ImpersonateGroup
	  Works only with custom setup of cluster. Allow impersonation of test group (Required if Impersonate-User is set)

	*/
	ImpersonateGroup *string
	/*ImpersonateUser
	  Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)

	*/
	ImpersonateUser *string
	/*AppName
	  name of Radix application

	*/
	AppName string
	/*EnvName
	  name of environment

	*/
	EnvName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete environment params
func (o *DeleteEnvironmentParams) WithTimeout(timeout time.Duration) *DeleteEnvironmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete environment params
func (o *DeleteEnvironmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete environment params
func (o *DeleteEnvironmentParams) WithContext(ctx context.Context) *DeleteEnvironmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete environment params
func (o *DeleteEnvironmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete environment params
func (o *DeleteEnvironmentParams) WithHTTPClient(client *http.Client) *DeleteEnvironmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete environment params
func (o *DeleteEnvironmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the delete environment params
func (o *DeleteEnvironmentParams) WithImpersonateGroup(impersonateGroup *string) *DeleteEnvironmentParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the delete environment params
func (o *DeleteEnvironmentParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the delete environment params
func (o *DeleteEnvironmentParams) WithImpersonateUser(impersonateUser *string) *DeleteEnvironmentParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the delete environment params
func (o *DeleteEnvironmentParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the delete environment params
func (o *DeleteEnvironmentParams) WithAppName(appName string) *DeleteEnvironmentParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the delete environment params
func (o *DeleteEnvironmentParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithEnvName adds the envName to the delete environment params
func (o *DeleteEnvironmentParams) WithEnvName(envName string) *DeleteEnvironmentParams {
	o.SetEnvName(envName)
	return o
}

// SetEnvName adds the envName to the delete environment params
func (o *DeleteEnvironmentParams) SetEnvName(envName string) {
	o.EnvName = envName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEnvironmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param envName
	if err := r.SetPathParam("envName", o.EnvName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
