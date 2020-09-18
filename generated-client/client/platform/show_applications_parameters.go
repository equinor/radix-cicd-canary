// Code generated by go-swagger; DO NOT EDIT.

package platform

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

// NewShowApplicationsParams creates a new ShowApplicationsParams object
// with the default values initialized.
func NewShowApplicationsParams() *ShowApplicationsParams {
	var ()
	return &ShowApplicationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewShowApplicationsParamsWithTimeout creates a new ShowApplicationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewShowApplicationsParamsWithTimeout(timeout time.Duration) *ShowApplicationsParams {
	var ()
	return &ShowApplicationsParams{

		timeout: timeout,
	}
}

// NewShowApplicationsParamsWithContext creates a new ShowApplicationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewShowApplicationsParamsWithContext(ctx context.Context) *ShowApplicationsParams {
	var ()
	return &ShowApplicationsParams{

		Context: ctx,
	}
}

// NewShowApplicationsParamsWithHTTPClient creates a new ShowApplicationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewShowApplicationsParamsWithHTTPClient(client *http.Client) *ShowApplicationsParams {
	var ()
	return &ShowApplicationsParams{
		HTTPClient: client,
	}
}

/*ShowApplicationsParams contains all the parameters to send to the API endpoint
for the show applications operation typically these are written to a http.Request
*/
type ShowApplicationsParams struct {

	/*ImpersonateGroup
	  Works only with custom setup of cluster. Allow impersonation of test group (Required if Impersonate-User is set)

	*/
	ImpersonateGroup *string
	/*ImpersonateUser
	  Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)

	*/
	ImpersonateUser *string
	/*SSHRepo
	  ssh repo to identify Radix application if exists

	*/
	SSHRepo *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the show applications params
func (o *ShowApplicationsParams) WithTimeout(timeout time.Duration) *ShowApplicationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the show applications params
func (o *ShowApplicationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the show applications params
func (o *ShowApplicationsParams) WithContext(ctx context.Context) *ShowApplicationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the show applications params
func (o *ShowApplicationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the show applications params
func (o *ShowApplicationsParams) WithHTTPClient(client *http.Client) *ShowApplicationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the show applications params
func (o *ShowApplicationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the show applications params
func (o *ShowApplicationsParams) WithImpersonateGroup(impersonateGroup *string) *ShowApplicationsParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the show applications params
func (o *ShowApplicationsParams) SetImpersonateGroup(impersonateGroup *string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the show applications params
func (o *ShowApplicationsParams) WithImpersonateUser(impersonateUser *string) *ShowApplicationsParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the show applications params
func (o *ShowApplicationsParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithSSHRepo adds the sSHRepo to the show applications params
func (o *ShowApplicationsParams) WithSSHRepo(sSHRepo *string) *ShowApplicationsParams {
	o.SetSSHRepo(sSHRepo)
	return o
}

// SetSSHRepo adds the sshRepo to the show applications params
func (o *ShowApplicationsParams) SetSSHRepo(sSHRepo *string) {
	o.SSHRepo = sSHRepo
}

// WriteToRequest writes these params to a swagger request
func (o *ShowApplicationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.SSHRepo != nil {

		// query param sshRepo
		var qrSSHRepo string
		if o.SSHRepo != nil {
			qrSSHRepo = *o.SSHRepo
		}
		qSSHRepo := qrSSHRepo
		if qSSHRepo != "" {
			if err := r.SetQueryParam("sshRepo", qSSHRepo); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
