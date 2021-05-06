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

// NewReplicaLogParams creates a new ReplicaLogParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplicaLogParams() *ReplicaLogParams {
	return &ReplicaLogParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplicaLogParamsWithTimeout creates a new ReplicaLogParams object
// with the ability to set a timeout on a request.
func NewReplicaLogParamsWithTimeout(timeout time.Duration) *ReplicaLogParams {
	return &ReplicaLogParams{
		timeout: timeout,
	}
}

// NewReplicaLogParamsWithContext creates a new ReplicaLogParams object
// with the ability to set a context for a request.
func NewReplicaLogParamsWithContext(ctx context.Context) *ReplicaLogParams {
	return &ReplicaLogParams{
		Context: ctx,
	}
}

// NewReplicaLogParamsWithHTTPClient creates a new ReplicaLogParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplicaLogParamsWithHTTPClient(client *http.Client) *ReplicaLogParams {
	return &ReplicaLogParams{
		HTTPClient: client,
	}
}

/* ReplicaLogParams contains all the parameters to send to the API endpoint
   for the replica log operation.

   Typically these are written to a http.Request.
*/
type ReplicaLogParams struct {

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

	/* PodName.

	   Name of pod
	*/
	PodName string

	/* SinceTime.

	   Get log only from sinceTime (example 2020-03-18T07:20:41+00:00)

	   Format: date-time
	*/
	SinceTime *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the replica log params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplicaLogParams) WithDefaults() *ReplicaLogParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replica log params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplicaLogParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the replica log params
func (o *ReplicaLogParams) WithTimeout(timeout time.Duration) *ReplicaLogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replica log params
func (o *ReplicaLogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replica log params
func (o *ReplicaLogParams) WithContext(ctx context.Context) *ReplicaLogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replica log params
func (o *ReplicaLogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replica log params
func (o *ReplicaLogParams) WithHTTPClient(client *http.Client) *ReplicaLogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replica log params
func (o *ReplicaLogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the replica log params
func (o *ReplicaLogParams) WithImpersonateGroup(impersonateGroup *string) *ReplicaLogParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the replica log params
func (o *ReplicaLogParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the replica log params
func (o *ReplicaLogParams) WithImpersonateUser(impersonateUser *string) *ReplicaLogParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the replica log params
func (o *ReplicaLogParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the replica log params
func (o *ReplicaLogParams) WithAppName(appName string) *ReplicaLogParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the replica log params
func (o *ReplicaLogParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithComponentName adds the componentName to the replica log params
func (o *ReplicaLogParams) WithComponentName(componentName string) *ReplicaLogParams {
	o.SetComponentName(componentName)
	return o
}

// SetComponentName adds the componentName to the replica log params
func (o *ReplicaLogParams) SetComponentName(componentName string) {
	o.ComponentName = componentName
}

// WithEnvName adds the envName to the replica log params
func (o *ReplicaLogParams) WithEnvName(envName string) *ReplicaLogParams {
	o.SetEnvName(envName)
	return o
}

// SetEnvName adds the envName to the replica log params
func (o *ReplicaLogParams) SetEnvName(envName string) {
	o.EnvName = envName
}

// WithPodName adds the podName to the replica log params
func (o *ReplicaLogParams) WithPodName(podName string) *ReplicaLogParams {
	o.SetPodName(podName)
	return o
}

// SetPodName adds the podName to the replica log params
func (o *ReplicaLogParams) SetPodName(podName string) {
	o.PodName = podName
}

// WithSinceTime adds the sinceTime to the replica log params
func (o *ReplicaLogParams) WithSinceTime(sinceTime *strfmt.DateTime) *ReplicaLogParams {
	o.SetSinceTime(sinceTime)
	return o
}

// SetSinceTime adds the sinceTime to the replica log params
func (o *ReplicaLogParams) SetSinceTime(sinceTime *strfmt.DateTime) {
	o.SinceTime = sinceTime
}

// WriteToRequest writes these params to a swagger request
func (o *ReplicaLogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param podName
	if err := r.SetPathParam("podName", o.PodName); err != nil {
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
