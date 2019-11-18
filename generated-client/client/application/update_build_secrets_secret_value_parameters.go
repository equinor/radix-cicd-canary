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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/equinor/radix-cicd-canary/generated-client/models"
)

// NewUpdateBuildSecretsSecretValueParams creates a new UpdateBuildSecretsSecretValueParams object
// with the default values initialized.
func NewUpdateBuildSecretsSecretValueParams() *UpdateBuildSecretsSecretValueParams {
	var ()
	return &UpdateBuildSecretsSecretValueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateBuildSecretsSecretValueParamsWithTimeout creates a new UpdateBuildSecretsSecretValueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateBuildSecretsSecretValueParamsWithTimeout(timeout time.Duration) *UpdateBuildSecretsSecretValueParams {
	var ()
	return &UpdateBuildSecretsSecretValueParams{

		timeout: timeout,
	}
}

// NewUpdateBuildSecretsSecretValueParamsWithContext creates a new UpdateBuildSecretsSecretValueParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateBuildSecretsSecretValueParamsWithContext(ctx context.Context) *UpdateBuildSecretsSecretValueParams {
	var ()
	return &UpdateBuildSecretsSecretValueParams{

		Context: ctx,
	}
}

// NewUpdateBuildSecretsSecretValueParamsWithHTTPClient creates a new UpdateBuildSecretsSecretValueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateBuildSecretsSecretValueParamsWithHTTPClient(client *http.Client) *UpdateBuildSecretsSecretValueParams {
	var ()
	return &UpdateBuildSecretsSecretValueParams{
		HTTPClient: client,
	}
}

/*UpdateBuildSecretsSecretValueParams contains all the parameters to send to the API endpoint
for the update build secrets secret value operation typically these are written to a http.Request
*/
type UpdateBuildSecretsSecretValueParams struct {

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
	/*SecretName
	  name of secret

	*/
	SecretName string
	/*SecretValue
	  New secret value

	*/
	SecretValue *models.SecretParameters

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update build secrets secret value params
func (o *UpdateBuildSecretsSecretValueParams) WithTimeout(timeout time.Duration) *UpdateBuildSecretsSecretValueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update build secrets secret value params
func (o *UpdateBuildSecretsSecretValueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update build secrets secret value params
func (o *UpdateBuildSecretsSecretValueParams) WithContext(ctx context.Context) *UpdateBuildSecretsSecretValueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update build secrets secret value params
func (o *UpdateBuildSecretsSecretValueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update build secrets secret value params
func (o *UpdateBuildSecretsSecretValueParams) WithHTTPClient(client *http.Client) *UpdateBuildSecretsSecretValueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update build secrets secret value params
func (o *UpdateBuildSecretsSecretValueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the update build secrets secret value params
func (o *UpdateBuildSecretsSecretValueParams) WithImpersonateGroup(impersonateGroup *string) *UpdateBuildSecretsSecretValueParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the update build secrets secret value params
func (o *UpdateBuildSecretsSecretValueParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the update build secrets secret value params
func (o *UpdateBuildSecretsSecretValueParams) WithImpersonateUser(impersonateUser *string) *UpdateBuildSecretsSecretValueParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the update build secrets secret value params
func (o *UpdateBuildSecretsSecretValueParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the update build secrets secret value params
func (o *UpdateBuildSecretsSecretValueParams) WithAppName(appName string) *UpdateBuildSecretsSecretValueParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the update build secrets secret value params
func (o *UpdateBuildSecretsSecretValueParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithSecretName adds the secretName to the update build secrets secret value params
func (o *UpdateBuildSecretsSecretValueParams) WithSecretName(secretName string) *UpdateBuildSecretsSecretValueParams {
	o.SetSecretName(secretName)
	return o
}

// SetSecretName adds the secretName to the update build secrets secret value params
func (o *UpdateBuildSecretsSecretValueParams) SetSecretName(secretName string) {
	o.SecretName = secretName
}

// WithSecretValue adds the secretValue to the update build secrets secret value params
func (o *UpdateBuildSecretsSecretValueParams) WithSecretValue(secretValue *models.SecretParameters) *UpdateBuildSecretsSecretValueParams {
	o.SetSecretValue(secretValue)
	return o
}

// SetSecretValue adds the secretValue to the update build secrets secret value params
func (o *UpdateBuildSecretsSecretValueParams) SetSecretValue(secretValue *models.SecretParameters) {
	o.SecretValue = secretValue
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateBuildSecretsSecretValueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param secretName
	if err := r.SetPathParam("secretName", o.SecretName); err != nil {
		return err
	}

	if o.SecretValue != nil {
		if err := r.SetBodyParam(o.SecretValue); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}