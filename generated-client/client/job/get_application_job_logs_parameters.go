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
)

// NewGetApplicationJobLogsParams creates a new GetApplicationJobLogsParams object
// with the default values initialized.
func NewGetApplicationJobLogsParams() *GetApplicationJobLogsParams {
	var ()
	return &GetApplicationJobLogsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetApplicationJobLogsParamsWithTimeout creates a new GetApplicationJobLogsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetApplicationJobLogsParamsWithTimeout(timeout time.Duration) *GetApplicationJobLogsParams {
	var ()
	return &GetApplicationJobLogsParams{

		timeout: timeout,
	}
}

// NewGetApplicationJobLogsParamsWithContext creates a new GetApplicationJobLogsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetApplicationJobLogsParamsWithContext(ctx context.Context) *GetApplicationJobLogsParams {
	var ()
	return &GetApplicationJobLogsParams{

		Context: ctx,
	}
}

// NewGetApplicationJobLogsParamsWithHTTPClient creates a new GetApplicationJobLogsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetApplicationJobLogsParamsWithHTTPClient(client *http.Client) *GetApplicationJobLogsParams {
	var ()
	return &GetApplicationJobLogsParams{
		HTTPClient: client,
	}
}

/*GetApplicationJobLogsParams contains all the parameters to send to the API endpoint
for the get application job logs operation typically these are written to a http.Request
*/
type GetApplicationJobLogsParams struct {

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
	/*JobName
	  Name of pipeline job

	*/
	JobName string
	/*SinceTime
	  Get log only from sinceTime (example 2020-03-18T07:20:41+00:00)

	*/
	SinceTime *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get application job logs params
func (o *GetApplicationJobLogsParams) WithTimeout(timeout time.Duration) *GetApplicationJobLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get application job logs params
func (o *GetApplicationJobLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get application job logs params
func (o *GetApplicationJobLogsParams) WithContext(ctx context.Context) *GetApplicationJobLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get application job logs params
func (o *GetApplicationJobLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get application job logs params
func (o *GetApplicationJobLogsParams) WithHTTPClient(client *http.Client) *GetApplicationJobLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get application job logs params
func (o *GetApplicationJobLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the get application job logs params
func (o *GetApplicationJobLogsParams) WithImpersonateGroup(impersonateGroup *string) *GetApplicationJobLogsParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the get application job logs params
func (o *GetApplicationJobLogsParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the get application job logs params
func (o *GetApplicationJobLogsParams) WithImpersonateUser(impersonateUser *string) *GetApplicationJobLogsParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the get application job logs params
func (o *GetApplicationJobLogsParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the get application job logs params
func (o *GetApplicationJobLogsParams) WithAppName(appName string) *GetApplicationJobLogsParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the get application job logs params
func (o *GetApplicationJobLogsParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithJobName adds the jobName to the get application job logs params
func (o *GetApplicationJobLogsParams) WithJobName(jobName string) *GetApplicationJobLogsParams {
	o.SetJobName(jobName)
	return o
}

// SetJobName adds the jobName to the get application job logs params
func (o *GetApplicationJobLogsParams) SetJobName(jobName string) {
	o.JobName = jobName
}

// WithSinceTime adds the sinceTime to the get application job logs params
func (o *GetApplicationJobLogsParams) WithSinceTime(sinceTime *strfmt.DateTime) *GetApplicationJobLogsParams {
	o.SetSinceTime(sinceTime)
	return o
}

// SetSinceTime adds the sinceTime to the get application job logs params
func (o *GetApplicationJobLogsParams) SetSinceTime(sinceTime *strfmt.DateTime) {
	o.SinceTime = sinceTime
}

// WriteToRequest writes these params to a swagger request
func (o *GetApplicationJobLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
