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
	"github.com/go-openapi/swag"

	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
)

// NewSearchApplicationsParams creates a new SearchApplicationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchApplicationsParams() *SearchApplicationsParams {
	return &SearchApplicationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchApplicationsParamsWithTimeout creates a new SearchApplicationsParams object
// with the ability to set a timeout on a request.
func NewSearchApplicationsParamsWithTimeout(timeout time.Duration) *SearchApplicationsParams {
	return &SearchApplicationsParams{
		timeout: timeout,
	}
}

// NewSearchApplicationsParamsWithContext creates a new SearchApplicationsParams object
// with the ability to set a context for a request.
func NewSearchApplicationsParamsWithContext(ctx context.Context) *SearchApplicationsParams {
	return &SearchApplicationsParams{
		Context: ctx,
	}
}

// NewSearchApplicationsParamsWithHTTPClient creates a new SearchApplicationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchApplicationsParamsWithHTTPClient(client *http.Client) *SearchApplicationsParams {
	return &SearchApplicationsParams{
		HTTPClient: client,
	}
}

/* SearchApplicationsParams contains all the parameters to send to the API endpoint
   for the search applications operation.

   Typically these are written to a http.Request.
*/
type SearchApplicationsParams struct {

	/* ImpersonateGroup.

	   Works only with custom setup of cluster. Allow impersonation of test group (Required if Impersonate-User is set)
	*/
	ImpersonateGroup []string

	/* ImpersonateUser.

	   Works only with custom setup of cluster. Allow impersonation of test users (Required if Impersonate-Group is set)
	*/
	ImpersonateUser *string

	/* ApplicationSearch.

	   List of application names to search for
	*/
	ApplicationSearch *models.ApplicationsSearchRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search applications params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchApplicationsParams) WithDefaults() *SearchApplicationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search applications params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchApplicationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search applications params
func (o *SearchApplicationsParams) WithTimeout(timeout time.Duration) *SearchApplicationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search applications params
func (o *SearchApplicationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search applications params
func (o *SearchApplicationsParams) WithContext(ctx context.Context) *SearchApplicationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search applications params
func (o *SearchApplicationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search applications params
func (o *SearchApplicationsParams) WithHTTPClient(client *http.Client) *SearchApplicationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search applications params
func (o *SearchApplicationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImpersonateGroup adds the impersonateGroup to the search applications params
func (o *SearchApplicationsParams) WithImpersonateGroup(impersonateGroup []string) *SearchApplicationsParams {
	o.SetImpersonateGroup(impersonateGroup)
	return o
}

// SetImpersonateGroup adds the impersonateGroup to the search applications params
func (o *SearchApplicationsParams) SetImpersonateGroup(impersonateGroup []string) {
	o.ImpersonateGroup = impersonateGroup
}

// WithImpersonateUser adds the impersonateUser to the search applications params
func (o *SearchApplicationsParams) WithImpersonateUser(impersonateUser *string) *SearchApplicationsParams {
	o.SetImpersonateUser(impersonateUser)
	return o
}

// SetImpersonateUser adds the impersonateUser to the search applications params
func (o *SearchApplicationsParams) SetImpersonateUser(impersonateUser *string) {
	o.ImpersonateUser = impersonateUser
}

// WithApplicationSearch adds the applicationSearch to the search applications params
func (o *SearchApplicationsParams) WithApplicationSearch(applicationSearch *models.ApplicationsSearchRequest) *SearchApplicationsParams {
	o.SetApplicationSearch(applicationSearch)
	return o
}

// SetApplicationSearch adds the applicationSearch to the search applications params
func (o *SearchApplicationsParams) SetApplicationSearch(applicationSearch *models.ApplicationsSearchRequest) {
	o.ApplicationSearch = applicationSearch
}

// WriteToRequest writes these params to a swagger request
func (o *SearchApplicationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.ApplicationSearch != nil {
		if err := r.SetBodyParam(o.ApplicationSearch); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSearchApplications binds the parameter Impersonate-Group
func (o *SearchApplicationsParams) bindParamImpersonateGroup(formats strfmt.Registry) []string {
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
