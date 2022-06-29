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
)

// NewGetEnvironmentEventsParams creates a new GetEnvironmentEventsParams object
// with the default values initialized.
func NewGetEnvironmentEventsParams() *GetEnvironmentEventsParams {
	var ()
	return &GetEnvironmentEventsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEnvironmentEventsParamsWithTimeout creates a new GetEnvironmentEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEnvironmentEventsParamsWithTimeout(timeout time.Duration) *GetEnvironmentEventsParams {
	var ()
	return &GetEnvironmentEventsParams{

		timeout: timeout,
	}
}

// NewGetEnvironmentEventsParamsWithContext creates a new GetEnvironmentEventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEnvironmentEventsParamsWithContext(ctx context.Context) *GetEnvironmentEventsParams {
	var ()
	return &GetEnvironmentEventsParams{

		Context: ctx,
	}
}

// NewGetEnvironmentEventsParamsWithHTTPClient creates a new GetEnvironmentEventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEnvironmentEventsParamsWithHTTPClient(client *http.Client) *GetEnvironmentEventsParams {
	var ()
	return &GetEnvironmentEventsParams{
		HTTPClient: client,
	}
}

/*GetEnvironmentEventsParams contains all the parameters to send to the API endpoint
for the get environment events operation typically these are written to a http.Request
*/
type GetEnvironmentEventsParams struct {

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

// WithTimeout adds the timeout to the get environment events params
func (o *GetEnvironmentEventsParams) WithTimeout(timeout time.Duration) *GetEnvironmentEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get environment events params
func (o *GetEnvironmentEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get environment events params
func (o *GetEnvironmentEventsParams) WithContext(ctx context.Context) *GetEnvironmentEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get environment events params
func (o *GetEnvironmentEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get environment events params
func (o *GetEnvironmentEventsParams) WithHTTPClient(client *http.Client) *GetEnvironmentEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get environment events params
func (o *GetEnvironmentEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the get environment events params
func (o *GetEnvironmentEventsParams) WithImpersonateGroup(impersonateGroup *string) *GetEnvironmentEventsParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the get environment events params
func (o *GetEnvironmentEventsParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the get environment events params
func (o *GetEnvironmentEventsParams) WithImpersonateUser(impersonateUser *string) *GetEnvironmentEventsParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the get environment events params
func (o *GetEnvironmentEventsParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the get environment events params
func (o *GetEnvironmentEventsParams) WithAppName(appName string) *GetEnvironmentEventsParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the get environment events params
func (o *GetEnvironmentEventsParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithEnvName adds the envName to the get environment events params
func (o *GetEnvironmentEventsParams) WithEnvName(envName string) *GetEnvironmentEventsParams {
	o.SetEnvName(envName)
	return o
}

// SetEnvName adds the envName to the get environment events params
func (o *GetEnvironmentEventsParams) SetEnvName(envName string) {
	o.EnvName = envName
}

// WriteToRequest writes these params to a swagger request
func (o *GetEnvironmentEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
