// Code generated by go-swagger; DO NOT EDIT.

package deployment

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

// NewGetDeploymentParams creates a new GetDeploymentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeploymentParams() *GetDeploymentParams {
	return &GetDeploymentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentParamsWithTimeout creates a new GetDeploymentParams object
// with the ability to set a timeout on a request.
func NewGetDeploymentParamsWithTimeout(timeout time.Duration) *GetDeploymentParams {
	return &GetDeploymentParams{
		timeout: timeout,
	}
}

// NewGetDeploymentParamsWithContext creates a new GetDeploymentParams object
// with the ability to set a context for a request.
func NewGetDeploymentParamsWithContext(ctx context.Context) *GetDeploymentParams {
	return &GetDeploymentParams{
		Context: ctx,
	}
}

// NewGetDeploymentParamsWithHTTPClient creates a new GetDeploymentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeploymentParamsWithHTTPClient(client *http.Client) *GetDeploymentParams {
	return &GetDeploymentParams{
		HTTPClient: client,
	}
}

/* GetDeploymentParams contains all the parameters to send to the API endpoint
   for the get deployment operation.

   Typically these are written to a http.Request.
*/
type GetDeploymentParams struct {

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

	/* DeploymentName.

	   name of deployment
	*/
	DeploymentName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get deployment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploymentParams) WithDefaults() *GetDeploymentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get deployment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploymentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get deployment params
func (o *GetDeploymentParams) WithTimeout(timeout time.Duration) *GetDeploymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deployment params
func (o *GetDeploymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deployment params
func (o *GetDeploymentParams) WithContext(ctx context.Context) *GetDeploymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deployment params
func (o *GetDeploymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deployment params
func (o *GetDeploymentParams) WithHTTPClient(client *http.Client) *GetDeploymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deployment params
func (o *GetDeploymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the get deployment params
func (o *GetDeploymentParams) WithImpersonateGroup(impersonateGroup []string) *GetDeploymentParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the get deployment params
func (o *GetDeploymentParams) SetImpersonateGroup(impersonateGroup []string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the get deployment params
func (o *GetDeploymentParams) WithImpersonateUser(impersonateUser *string) *GetDeploymentParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the get deployment params
func (o *GetDeploymentParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the get deployment params
func (o *GetDeploymentParams) WithAppName(appName string) *GetDeploymentParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the get deployment params
func (o *GetDeploymentParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithDeploymentName adds the deploymentName to the get deployment params
func (o *GetDeploymentParams) WithDeploymentName(deploymentName string) *GetDeploymentParams {
	o.SetDeploymentName(deploymentName)
	return o
}

// SetDeploymentName adds the deploymentName to the get deployment params
func (o *GetDeploymentParams) SetDeploymentName(deploymentName string) {
	o.DeploymentName = deploymentName
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param deploymentName
	if err := r.SetPathParam("deploymentName", o.DeploymentName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetDeployment binds the parameter Impersonate-Group
func (o *GetDeploymentParams) bindParamImpersonateGroup(formats strfmt.Registry) []string {
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
