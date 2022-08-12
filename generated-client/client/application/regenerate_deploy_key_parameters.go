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

	"github.com/equinor/radix-cicd-canary/generated-client/models"
)

// NewRegenerateDeployKeyParams creates a new RegenerateDeployKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRegenerateDeployKeyParams() *RegenerateDeployKeyParams {
	return &RegenerateDeployKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRegenerateDeployKeyParamsWithTimeout creates a new RegenerateDeployKeyParams object
// with the ability to set a timeout on a request.
func NewRegenerateDeployKeyParamsWithTimeout(timeout time.Duration) *RegenerateDeployKeyParams {
	return &RegenerateDeployKeyParams{
		timeout: timeout,
	}
}

// NewRegenerateDeployKeyParamsWithContext creates a new RegenerateDeployKeyParams object
// with the ability to set a context for a request.
func NewRegenerateDeployKeyParamsWithContext(ctx context.Context) *RegenerateDeployKeyParams {
	return &RegenerateDeployKeyParams{
		Context: ctx,
	}
}

// NewRegenerateDeployKeyParamsWithHTTPClient creates a new RegenerateDeployKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewRegenerateDeployKeyParamsWithHTTPClient(client *http.Client) *RegenerateDeployKeyParams {
	return &RegenerateDeployKeyParams{
		HTTPClient: client,
	}
}

/* RegenerateDeployKeyParams contains all the parameters to send to the API endpoint
   for the regenerate deploy key operation.

   Typically these are written to a http.Request.
*/
type RegenerateDeployKeyParams struct {

	/* ImpersonateGroup.

	   Works only with custom setup of cluster. Allow impersonation of test group (Required if Impersonate-User is set)
	*/
	ImpersonateGroup *string

	/* ImpersonateUser.

	   Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)
	*/
	ImpersonateUser *string

	/* AppName.

	   name of application
	*/
	AppName string

	/* RegenerateDeployKeyAndSecretData.

	   Regenerate deploy key and secret data
	*/
	RegenerateDeployKeyAndSecretData *models.RegenerateDeployKeyAndSecretData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the regenerate deploy key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegenerateDeployKeyParams) WithDefaults() *RegenerateDeployKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the regenerate deploy key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RegenerateDeployKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the regenerate deploy key params
func (o *RegenerateDeployKeyParams) WithTimeout(timeout time.Duration) *RegenerateDeployKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the regenerate deploy key params
func (o *RegenerateDeployKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the regenerate deploy key params
func (o *RegenerateDeployKeyParams) WithContext(ctx context.Context) *RegenerateDeployKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the regenerate deploy key params
func (o *RegenerateDeployKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the regenerate deploy key params
func (o *RegenerateDeployKeyParams) WithHTTPClient(client *http.Client) *RegenerateDeployKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the regenerate deploy key params
func (o *RegenerateDeployKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the regenerate deploy key params
func (o *RegenerateDeployKeyParams) WithImpersonateGroup(impersonateGroup *string) *RegenerateDeployKeyParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the regenerate deploy key params
func (o *RegenerateDeployKeyParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the regenerate deploy key params
func (o *RegenerateDeployKeyParams) WithImpersonateUser(impersonateUser *string) *RegenerateDeployKeyParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the regenerate deploy key params
func (o *RegenerateDeployKeyParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the regenerate deploy key params
func (o *RegenerateDeployKeyParams) WithAppName(appName string) *RegenerateDeployKeyParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the regenerate deploy key params
func (o *RegenerateDeployKeyParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithRegenerateDeployKeyAndSecretData adds the regenerateDeployKeyAndSecretData to the regenerate deploy key params
func (o *RegenerateDeployKeyParams) WithRegenerateDeployKeyAndSecretData(regenerateDeployKeyAndSecretData *models.RegenerateDeployKeyAndSecretData) *RegenerateDeployKeyParams {
	o.SetRegenerateDeployKeyAndSecretData(regenerateDeployKeyAndSecretData)
	return o
}

// SetRegenerateDeployKeyAndSecretData adds the regenerateDeployKeyAndSecretData to the regenerate deploy key params
func (o *RegenerateDeployKeyParams) SetRegenerateDeployKeyAndSecretData(regenerateDeployKeyAndSecretData *models.RegenerateDeployKeyAndSecretData) {
	o.RegenerateDeployKeyAndSecretData = regenerateDeployKeyAndSecretData
}

// WriteToRequest writes these params to a swagger request
func (o *RegenerateDeployKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.RegenerateDeployKeyAndSecretData != nil {
		if err := r.SetBodyParam(o.RegenerateDeployKeyAndSecretData); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
