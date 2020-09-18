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

// NewGetPrivateImageHubsParams creates a new GetPrivateImageHubsParams object
// with the default values initialized.
func NewGetPrivateImageHubsParams() *GetPrivateImageHubsParams {
	var ()
	return &GetPrivateImageHubsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateImageHubsParamsWithTimeout creates a new GetPrivateImageHubsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateImageHubsParamsWithTimeout(timeout time.Duration) *GetPrivateImageHubsParams {
	var ()
	return &GetPrivateImageHubsParams{

		timeout: timeout,
	}
}

// NewGetPrivateImageHubsParamsWithContext creates a new GetPrivateImageHubsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateImageHubsParamsWithContext(ctx context.Context) *GetPrivateImageHubsParams {
	var ()
	return &GetPrivateImageHubsParams{

		Context: ctx,
	}
}

// NewGetPrivateImageHubsParamsWithHTTPClient creates a new GetPrivateImageHubsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateImageHubsParamsWithHTTPClient(client *http.Client) *GetPrivateImageHubsParams {
	var ()
	return &GetPrivateImageHubsParams{
		HTTPClient: client,
	}
}

/*GetPrivateImageHubsParams contains all the parameters to send to the API endpoint
for the get private image hubs operation typically these are written to a http.Request
*/
type GetPrivateImageHubsParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private image hubs params
func (o *GetPrivateImageHubsParams) WithTimeout(timeout time.Duration) *GetPrivateImageHubsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private image hubs params
func (o *GetPrivateImageHubsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private image hubs params
func (o *GetPrivateImageHubsParams) WithContext(ctx context.Context) *GetPrivateImageHubsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private image hubs params
func (o *GetPrivateImageHubsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private image hubs params
func (o *GetPrivateImageHubsParams) WithHTTPClient(client *http.Client) *GetPrivateImageHubsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private image hubs params
func (o *GetPrivateImageHubsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the get private image hubs params
func (o *GetPrivateImageHubsParams) WithImpersonateGroup(impersonateGroup *string) *GetPrivateImageHubsParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the get private image hubs params
func (o *GetPrivateImageHubsParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the get private image hubs params
func (o *GetPrivateImageHubsParams) WithImpersonateUser(impersonateUser *string) *GetPrivateImageHubsParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the get private image hubs params
func (o *GetPrivateImageHubsParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithAppName adds the appName to the get private image hubs params
func (o *GetPrivateImageHubsParams) WithAppName(appName string) *GetPrivateImageHubsParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the get private image hubs params
func (o *GetPrivateImageHubsParams) SetAppName(appName string) {
	o.AppName = appName
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateImageHubsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
