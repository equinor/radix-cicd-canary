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

	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
)

// NewModifyRegistrationDetailsParams creates a new ModifyRegistrationDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewModifyRegistrationDetailsParams() *ModifyRegistrationDetailsParams {
	return &ModifyRegistrationDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewModifyRegistrationDetailsParamsWithTimeout creates a new ModifyRegistrationDetailsParams object
// with the ability to set a timeout on a request.
func NewModifyRegistrationDetailsParamsWithTimeout(timeout time.Duration) *ModifyRegistrationDetailsParams {
	return &ModifyRegistrationDetailsParams{
		timeout: timeout,
	}
}

// NewModifyRegistrationDetailsParamsWithContext creates a new ModifyRegistrationDetailsParams object
// with the ability to set a context for a request.
func NewModifyRegistrationDetailsParamsWithContext(ctx context.Context) *ModifyRegistrationDetailsParams {
	return &ModifyRegistrationDetailsParams{
		Context: ctx,
	}
}

// NewModifyRegistrationDetailsParamsWithHTTPClient creates a new ModifyRegistrationDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewModifyRegistrationDetailsParamsWithHTTPClient(client *http.Client) *ModifyRegistrationDetailsParams {
	return &ModifyRegistrationDetailsParams{
		HTTPClient: client,
	}
}

/*
ModifyRegistrationDetailsParams contains all the parameters to send to the API endpoint

	for the modify registration details operation.

	Typically these are written to a http.Request.
*/
type ModifyRegistrationDetailsParams struct {

	/* ImpersonateGroup.

	   Works only with custom setup of cluster. Allow impersonation of a comma-seperated list of test groups (Required if Impersonate-User is set)
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

	/* PatchRequest.

	   Request for Application to patch
	*/
	PatchRequest *models.ApplicationRegistrationPatchRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the modify registration details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModifyRegistrationDetailsParams) WithDefaults() *ModifyRegistrationDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the modify registration details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModifyRegistrationDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the modify registration details params
func (o *ModifyRegistrationDetailsParams) WithTimeout(timeout time.Duration) *ModifyRegistrationDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify registration details params
func (o *ModifyRegistrationDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify registration details params
func (o *ModifyRegistrationDetailsParams) WithContext(ctx context.Context) *ModifyRegistrationDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify registration details params
func (o *ModifyRegistrationDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify registration details params
func (o *ModifyRegistrationDetailsParams) WithHTTPClient(client *http.Client) *ModifyRegistrationDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify registration details params
func (o *ModifyRegistrationDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the modify registration details params
func (o *ModifyRegistrationDetailsParams) WithImpersonateGroup(impersonateGroup *string) *ModifyRegistrationDetailsParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the modify registration details params
func (o *ModifyRegistrationDetailsParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the modify registration details params
func (o *ModifyRegistrationDetailsParams) WithImpersonateUser(impersonateUser *string) *ModifyRegistrationDetailsParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the modify registration details params
func (o *ModifyRegistrationDetailsParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the modify registration details params
func (o *ModifyRegistrationDetailsParams) WithAppName(appName string) *ModifyRegistrationDetailsParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the modify registration details params
func (o *ModifyRegistrationDetailsParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithPatchRequest adds the patchRequest to the modify registration details params
func (o *ModifyRegistrationDetailsParams) WithPatchRequest(patchRequest *models.ApplicationRegistrationPatchRequest) *ModifyRegistrationDetailsParams {
	o.SetPatchRequest(patchRequest)
	return o
}

// SetPatchRequest adds the patchRequest to the modify registration details params
func (o *ModifyRegistrationDetailsParams) SetPatchRequest(patchRequest *models.ApplicationRegistrationPatchRequest) {
	o.PatchRequest = patchRequest
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyRegistrationDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.PatchRequest != nil {
		if err := r.SetBodyParam(o.PatchRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
