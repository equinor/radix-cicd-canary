// Code generated by go-swagger; DO NOT EDIT.

package environment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new environment API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for environment API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ChangeEnvironmentComponentSecret updates an application environment component secret
*/
func (a *Client) ChangeEnvironmentComponentSecret(params *ChangeEnvironmentComponentSecretParams, authInfo runtime.ClientAuthInfoWriter) (*ChangeEnvironmentComponentSecretOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeEnvironmentComponentSecretParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changeEnvironmentComponentSecret",
		Method:             "PUT",
		PathPattern:        "/applications/{appName}/environments/{envName}/components/{componentName}/secrets/{secretName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeEnvironmentComponentSecretReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChangeEnvironmentComponentSecretOK), nil

}

/*
DeleteEnvironment deletes application environment
*/
func (a *Client) DeleteEnvironment(params *DeleteEnvironmentParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteEnvironmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEnvironmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteEnvironment",
		Method:             "DELETE",
		PathPattern:        "/applications/{appName}/environments/{envName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteEnvironmentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteEnvironmentOK), nil

}

/*
GetApplicationEnvironmentDeployments lists the application environment deployments
*/
func (a *Client) GetApplicationEnvironmentDeployments(params *GetApplicationEnvironmentDeploymentsParams, authInfo runtime.ClientAuthInfoWriter) (*GetApplicationEnvironmentDeploymentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetApplicationEnvironmentDeploymentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getApplicationEnvironmentDeployments",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/environments/{envName}/deployments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetApplicationEnvironmentDeploymentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetApplicationEnvironmentDeploymentsOK), nil

}

/*
GetEnvironment gets details for an application environment
*/
func (a *Client) GetEnvironment(params *GetEnvironmentParams, authInfo runtime.ClientAuthInfoWriter) (*GetEnvironmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEnvironmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEnvironment",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/environments/{envName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEnvironmentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEnvironmentOK), nil

}

/*
GetEnvironmentSummary lists the environments for an application
*/
func (a *Client) GetEnvironmentSummary(params *GetEnvironmentSummaryParams, authInfo runtime.ClientAuthInfoWriter) (*GetEnvironmentSummaryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEnvironmentSummaryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEnvironmentSummary",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/environments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEnvironmentSummaryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEnvironmentSummaryOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}