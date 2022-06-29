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

// NewUpdatePrivateImageHubsSecretValueParams creates a new UpdatePrivateImageHubsSecretValueParams object
// with the default values initialized.
func NewUpdatePrivateImageHubsSecretValueParams() *UpdatePrivateImageHubsSecretValueParams {
	var ()
	return &UpdatePrivateImageHubsSecretValueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePrivateImageHubsSecretValueParamsWithTimeout creates a new UpdatePrivateImageHubsSecretValueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePrivateImageHubsSecretValueParamsWithTimeout(timeout time.Duration) *UpdatePrivateImageHubsSecretValueParams {
	var ()
	return &UpdatePrivateImageHubsSecretValueParams{

		timeout: timeout,
	}
}

// NewUpdatePrivateImageHubsSecretValueParamsWithContext creates a new UpdatePrivateImageHubsSecretValueParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePrivateImageHubsSecretValueParamsWithContext(ctx context.Context) *UpdatePrivateImageHubsSecretValueParams {
	var ()
	return &UpdatePrivateImageHubsSecretValueParams{

		Context: ctx,
	}
}

// NewUpdatePrivateImageHubsSecretValueParamsWithHTTPClient creates a new UpdatePrivateImageHubsSecretValueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePrivateImageHubsSecretValueParamsWithHTTPClient(client *http.Client) *UpdatePrivateImageHubsSecretValueParams {
	var ()
	return &UpdatePrivateImageHubsSecretValueParams{
		HTTPClient: client,
	}
}

/*UpdatePrivateImageHubsSecretValueParams contains all the parameters to send to the API endpoint
for the update private image hubs secret value operation typically these are written to a http.Request
*/
type UpdatePrivateImageHubsSecretValueParams struct {

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
	/*ImageHubSecret
	  New secret value

	*/
	ImageHubSecret *models.SecretParameters
	/*ServerName
	  server name to update

	*/
	ServerName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update private image hubs secret value params
func (o *UpdatePrivateImageHubsSecretValueParams) WithTimeout(timeout time.Duration) *UpdatePrivateImageHubsSecretValueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update private image hubs secret value params
func (o *UpdatePrivateImageHubsSecretValueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update private image hubs secret value params
func (o *UpdatePrivateImageHubsSecretValueParams) WithContext(ctx context.Context) *UpdatePrivateImageHubsSecretValueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update private image hubs secret value params
func (o *UpdatePrivateImageHubsSecretValueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update private image hubs secret value params
func (o *UpdatePrivateImageHubsSecretValueParams) WithHTTPClient(client *http.Client) *UpdatePrivateImageHubsSecretValueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update private image hubs secret value params
func (o *UpdatePrivateImageHubsSecretValueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the update private image hubs secret value params
func (o *UpdatePrivateImageHubsSecretValueParams) WithImpersonateGroup(impersonateGroup *string) *UpdatePrivateImageHubsSecretValueParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the update private image hubs secret value params
func (o *UpdatePrivateImageHubsSecretValueParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the update private image hubs secret value params
func (o *UpdatePrivateImageHubsSecretValueParams) WithImpersonateUser(impersonateUser *string) *UpdatePrivateImageHubsSecretValueParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the update private image hubs secret value params
func (o *UpdatePrivateImageHubsSecretValueParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the update private image hubs secret value params
func (o *UpdatePrivateImageHubsSecretValueParams) WithAppName(appName string) *UpdatePrivateImageHubsSecretValueParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the update private image hubs secret value params
func (o *UpdatePrivateImageHubsSecretValueParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithImageHubSecret adds the imageHubSecret to the update private image hubs secret value params
func (o *UpdatePrivateImageHubsSecretValueParams) WithImageHubSecret(imageHubSecret *models.SecretParameters) *UpdatePrivateImageHubsSecretValueParams {
	o.SetImageHubSecret(imageHubSecret)
	return o
}

// SetImageHubSecret adds the imageHubSecret to the update private image hubs secret value params
func (o *UpdatePrivateImageHubsSecretValueParams) SetImageHubSecret(imageHubSecret *models.SecretParameters) {
	o.ImageHubSecret = imageHubSecret
}

// WithServerName adds the serverName to the update private image hubs secret value params
func (o *UpdatePrivateImageHubsSecretValueParams) WithServerName(serverName string) *UpdatePrivateImageHubsSecretValueParams {
	o.SetServerName(serverName)
	return o
}

// SetServerName adds the serverName to the update private image hubs secret value params
func (o *UpdatePrivateImageHubsSecretValueParams) SetServerName(serverName string) {
	o.ServerName = serverName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePrivateImageHubsSecretValueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ImageHubSecret != nil {
		if err := r.SetBodyParam(o.ImageHubSecret); err != nil {
			return err
		}
	}

	// path param serverName
	if err := r.SetPathParam("serverName", o.ServerName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
