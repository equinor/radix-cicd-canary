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
	"github.com/go-openapi/swag"
)

// NewGetApplicationParams creates a new GetApplicationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetApplicationParams() *GetApplicationParams {
	return &GetApplicationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetApplicationParamsWithTimeout creates a new GetApplicationParams object
// with the ability to set a timeout on a request.
func NewGetApplicationParamsWithTimeout(timeout time.Duration) *GetApplicationParams {
	return &GetApplicationParams{
		timeout: timeout,
	}
}

// NewGetApplicationParamsWithContext creates a new GetApplicationParams object
// with the ability to set a context for a request.
func NewGetApplicationParamsWithContext(ctx context.Context) *GetApplicationParams {
	return &GetApplicationParams{
		Context: ctx,
	}
}

// NewGetApplicationParamsWithHTTPClient creates a new GetApplicationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetApplicationParamsWithHTTPClient(client *http.Client) *GetApplicationParams {
	return &GetApplicationParams{
		HTTPClient: client,
	}
}

/* GetApplicationParams contains all the parameters to send to the API endpoint
   for the get application operation.

   Typically these are written to a http.Request.
*/
type GetApplicationParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get application params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetApplicationParams) WithDefaults() *GetApplicationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get application params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetApplicationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get application params
func (o *GetApplicationParams) WithTimeout(timeout time.Duration) *GetApplicationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get application params
func (o *GetApplicationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get application params
func (o *GetApplicationParams) WithContext(ctx context.Context) *GetApplicationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get application params
func (o *GetApplicationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get application params
func (o *GetApplicationParams) WithHTTPClient(client *http.Client) *GetApplicationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get application params
func (o *GetApplicationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the get application params
func (o *GetApplicationParams) WithImpersonateGroup(impersonateGroup []string) *GetApplicationParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the get application params
func (o *GetApplicationParams) SetImpersonateGroup(impersonateGroup []string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the get application params
func (o *GetApplicationParams) WithImpersonateUser(impersonateUser *string) *GetApplicationParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the get application params
func (o *GetApplicationParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the get application params
func (o *GetApplicationParams) WithAppName(appName string) *GetApplicationParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the get application params
func (o *GetApplicationParams) SetAppName(appName string) {
	o.AppName = appName
}

// WriteToRequest writes these params to a swagger request
func (o *GetApplicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetApplication binds the parameter Impersonate-Group
func (o *GetApplicationParams) bindParamImpersonateGroup(formats strfmt.Registry) []string {
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
