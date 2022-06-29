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
	"github.com/go-openapi/strfmt"
)

// NewGetApplicationAlertingConfigParams creates a new GetApplicationAlertingConfigParams object
// with the default values initialized.
func NewGetApplicationAlertingConfigParams() *GetApplicationAlertingConfigParams {
	var ()
	return &GetApplicationAlertingConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetApplicationAlertingConfigParamsWithTimeout creates a new GetApplicationAlertingConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetApplicationAlertingConfigParamsWithTimeout(timeout time.Duration) *GetApplicationAlertingConfigParams {
	var ()
	return &GetApplicationAlertingConfigParams{

		timeout: timeout,
	}
}

// NewGetApplicationAlertingConfigParamsWithContext creates a new GetApplicationAlertingConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetApplicationAlertingConfigParamsWithContext(ctx context.Context) *GetApplicationAlertingConfigParams {
	var ()
	return &GetApplicationAlertingConfigParams{

		Context: ctx,
	}
}

// NewGetApplicationAlertingConfigParamsWithHTTPClient creates a new GetApplicationAlertingConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetApplicationAlertingConfigParamsWithHTTPClient(client *http.Client) *GetApplicationAlertingConfigParams {
	var ()
	return &GetApplicationAlertingConfigParams{
		HTTPClient: client,
	}
}

/*GetApplicationAlertingConfigParams contains all the parameters to send to the API endpoint
for the get application alerting config operation typically these are written to a http.Request
*/
type GetApplicationAlertingConfigParams struct {

	/*ImpersonateGroup
	  Works only with custom setup of cluster. Allow impersonation of test group (Required if Impersonate-User is set)

	*/
	ImpersonateGroup *string
	/*ImpersonateUser
	  Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)

	*/
	ImpersonateUser *string
	/*AppName
	  Name of application

	*/
	AppName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get application alerting config params
func (o *GetApplicationAlertingConfigParams) WithTimeout(timeout time.Duration) *GetApplicationAlertingConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get application alerting config params
func (o *GetApplicationAlertingConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get application alerting config params
func (o *GetApplicationAlertingConfigParams) WithContext(ctx context.Context) *GetApplicationAlertingConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get application alerting config params
func (o *GetApplicationAlertingConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get application alerting config params
func (o *GetApplicationAlertingConfigParams) WithHTTPClient(client *http.Client) *GetApplicationAlertingConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get application alerting config params
func (o *GetApplicationAlertingConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the get application alerting config params
func (o *GetApplicationAlertingConfigParams) WithImpersonateGroup(impersonateGroup *string) *GetApplicationAlertingConfigParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the get application alerting config params
func (o *GetApplicationAlertingConfigParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the get application alerting config params
func (o *GetApplicationAlertingConfigParams) WithImpersonateUser(impersonateUser *string) *GetApplicationAlertingConfigParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the get application alerting config params
func (o *GetApplicationAlertingConfigParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the get application alerting config params
func (o *GetApplicationAlertingConfigParams) WithAppName(appName string) *GetApplicationAlertingConfigParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the get application alerting config params
func (o *GetApplicationAlertingConfigParams) SetAppName(appName string) {
	o.AppName = appName
}

// WriteToRequest writes these params to a swagger request
func (o *GetApplicationAlertingConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
