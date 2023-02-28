// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new application API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for application API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ChangeRegistrationDetails(params *ChangeRegistrationDetailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChangeRegistrationDetailsOK, error)

	DeleteApplication(params *DeleteApplicationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteApplicationOK, error)

	DisableApplicationAlerting(params *DisableApplicationAlertingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DisableApplicationAlertingOK, error)

	EnableApplicationAlerting(params *EnableApplicationAlertingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EnableApplicationAlertingOK, error)

	GetApplication(params *GetApplicationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetApplicationOK, error)

	GetApplicationAlertingConfig(params *GetApplicationAlertingConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetApplicationAlertingConfigOK, error)

	GetBuildSecrets(params *GetBuildSecretsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBuildSecretsOK, error)

	GetDeployments(params *GetDeploymentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeploymentsOK, error)

	GetPrivateImageHubs(params *GetPrivateImageHubsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPrivateImageHubsOK, error)

	IsDeployKeyValid(params *IsDeployKeyValidParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IsDeployKeyValidOK, error)

	ListPipelines(params *ListPipelinesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListPipelinesOK, error)

	ModifyRegistrationDetails(params *ModifyRegistrationDetailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ModifyRegistrationDetailsOK, error)

	RegenerateDeployKey(params *RegenerateDeployKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegenerateDeployKeyOK, error)

	RegenerateMachineUserToken(params *RegenerateMachineUserTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegenerateMachineUserTokenOK, error)

	RestartApplication(params *RestartApplicationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestartApplicationOK, error)

	StartApplication(params *StartApplicationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StartApplicationOK, error)

	StopApplication(params *StopApplicationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StopApplicationOK, error)

	TriggerPipelineBuild(params *TriggerPipelineBuildParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TriggerPipelineBuildOK, error)

	TriggerPipelineBuildDeploy(params *TriggerPipelineBuildDeployParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TriggerPipelineBuildDeployOK, error)

	TriggerPipelineDeploy(params *TriggerPipelineDeployParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TriggerPipelineDeployOK, error)

	TriggerPipelinePromote(params *TriggerPipelinePromoteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TriggerPipelinePromoteOK, error)

	UpdateApplicationAlertingConfig(params *UpdateApplicationAlertingConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateApplicationAlertingConfigOK, error)

	UpdateBuildSecretsSecretValue(params *UpdateBuildSecretsSecretValueParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateBuildSecretsSecretValueOK, error)

	UpdatePrivateImageHubsSecretValue(params *UpdatePrivateImageHubsSecretValueParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdatePrivateImageHubsSecretValueOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ChangeRegistrationDetails updates application registration
*/
func (a *Client) ChangeRegistrationDetails(params *ChangeRegistrationDetailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChangeRegistrationDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeRegistrationDetailsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "changeRegistrationDetails",
		Method:             "PUT",
		PathPattern:        "/applications/{appName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeRegistrationDetailsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChangeRegistrationDetailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for changeRegistrationDetails: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteApplication deletes application
*/
func (a *Client) DeleteApplication(params *DeleteApplicationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteApplicationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteApplicationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteApplication",
		Method:             "DELETE",
		PathPattern:        "/applications/{appName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteApplicationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteApplicationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteApplication: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DisableApplicationAlerting disables alerting for application namespace
*/
func (a *Client) DisableApplicationAlerting(params *DisableApplicationAlertingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DisableApplicationAlertingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisableApplicationAlertingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "disableApplicationAlerting",
		Method:             "POST",
		PathPattern:        "/applications/{appName}/alerting/disable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DisableApplicationAlertingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DisableApplicationAlertingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for disableApplicationAlerting: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EnableApplicationAlerting enables alerting for application namespace
*/
func (a *Client) EnableApplicationAlerting(params *EnableApplicationAlertingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EnableApplicationAlertingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEnableApplicationAlertingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "enableApplicationAlerting",
		Method:             "POST",
		PathPattern:        "/applications/{appName}/alerting/enable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EnableApplicationAlertingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EnableApplicationAlertingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for enableApplicationAlerting: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetApplication gets the application by name
*/
func (a *Client) GetApplication(params *GetApplicationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetApplicationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetApplicationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getApplication",
		Method:             "GET",
		PathPattern:        "/applications/{appName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetApplicationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetApplicationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getApplication: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetApplicationAlertingConfig gets alerts configuration for application namespace
*/
func (a *Client) GetApplicationAlertingConfig(params *GetApplicationAlertingConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetApplicationAlertingConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetApplicationAlertingConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getApplicationAlertingConfig",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/alerting",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetApplicationAlertingConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetApplicationAlertingConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getApplicationAlertingConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetBuildSecrets lists the application build secrets
*/
func (a *Client) GetBuildSecrets(params *GetBuildSecretsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBuildSecretsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBuildSecretsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBuildSecrets",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/buildsecrets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetBuildSecretsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBuildSecretsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBuildSecrets: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDeployments lists the application deployments
*/
func (a *Client) GetDeployments(params *GetDeploymentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeploymentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploymentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDeployments",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/deployments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDeploymentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeploymentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDeployments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPrivateImageHubs lists the application private image hubs
*/
func (a *Client) GetPrivateImageHubs(params *GetPrivateImageHubsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetPrivateImageHubsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateImageHubsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPrivateImageHubs",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/privateimagehubs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPrivateImageHubsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPrivateImageHubsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPrivateImageHubs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  IsDeployKeyValid checks if the deploy key is correctly setup for application by cloning the repository
*/
func (a *Client) IsDeployKeyValid(params *IsDeployKeyValidParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IsDeployKeyValidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIsDeployKeyValidParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "isDeployKeyValid",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/deploykey-valid",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &IsDeployKeyValidReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IsDeployKeyValidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for isDeployKeyValid: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListPipelines lists the supported pipelines
*/
func (a *Client) ListPipelines(params *ListPipelinesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListPipelinesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPipelinesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listPipelines",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/pipelines",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListPipelinesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListPipelinesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listPipelines: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ModifyRegistrationDetails updates specific field s of an application registration
*/
func (a *Client) ModifyRegistrationDetails(params *ModifyRegistrationDetailsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ModifyRegistrationDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyRegistrationDetailsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "modifyRegistrationDetails",
		Method:             "PATCH",
		PathPattern:        "/applications/{appName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ModifyRegistrationDetailsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ModifyRegistrationDetailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for modifyRegistrationDetails: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RegenerateDeployKey regenerates deploy key
*/
func (a *Client) RegenerateDeployKey(params *RegenerateDeployKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegenerateDeployKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegenerateDeployKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "regenerateDeployKey",
		Method:             "POST",
		PathPattern:        "/applications/{appName}/regenerate-deploy-key",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RegenerateDeployKeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RegenerateDeployKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for regenerateDeployKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RegenerateMachineUserToken regenerates machine user token
*/
func (a *Client) RegenerateMachineUserToken(params *RegenerateMachineUserTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RegenerateMachineUserTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegenerateMachineUserTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "regenerateMachineUserToken",
		Method:             "POST",
		PathPattern:        "/applications/{appName}/regenerate-machine-user-token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RegenerateMachineUserTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RegenerateMachineUserTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for regenerateMachineUserToken: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RestartApplication restarts all components in all environments of the application stops all running components in all environments of the application pulls new images from image hub in radix configuration starts all components in all environments of the application again using up to date image
*/
func (a *Client) RestartApplication(params *RestartApplicationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestartApplicationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestartApplicationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "restartApplication",
		Method:             "POST",
		PathPattern:        "/applications/{appName}/restart",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RestartApplicationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RestartApplicationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for restartApplication: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StartApplication starts all components in all environments of the application
*/
func (a *Client) StartApplication(params *StartApplicationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StartApplicationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartApplicationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "startApplication",
		Method:             "POST",
		PathPattern:        "/applications/{appName}/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartApplicationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StartApplicationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for startApplication: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StopApplication stops all components in the environment
*/
func (a *Client) StopApplication(params *StopApplicationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StopApplicationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopApplicationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "stopApplication",
		Method:             "POST",
		PathPattern:        "/applications/{appName}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopApplicationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StopApplicationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for stopApplication: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TriggerPipelineBuild runs a build pipeline for a given application and branch
*/
func (a *Client) TriggerPipelineBuild(params *TriggerPipelineBuildParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TriggerPipelineBuildOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTriggerPipelineBuildParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "triggerPipelineBuild",
		Method:             "POST",
		PathPattern:        "/applications/{appName}/pipelines/build",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TriggerPipelineBuildReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TriggerPipelineBuildOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for triggerPipelineBuild: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TriggerPipelineBuildDeploy runs a build deploy pipeline for a given application and branch
*/
func (a *Client) TriggerPipelineBuildDeploy(params *TriggerPipelineBuildDeployParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TriggerPipelineBuildDeployOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTriggerPipelineBuildDeployParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "triggerPipelineBuildDeploy",
		Method:             "POST",
		PathPattern:        "/applications/{appName}/pipelines/build-deploy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TriggerPipelineBuildDeployReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TriggerPipelineBuildDeployOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for triggerPipelineBuildDeploy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TriggerPipelineDeploy runs a deploy pipeline for a given application and environment
*/
func (a *Client) TriggerPipelineDeploy(params *TriggerPipelineDeployParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TriggerPipelineDeployOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTriggerPipelineDeployParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "triggerPipelineDeploy",
		Method:             "POST",
		PathPattern:        "/applications/{appName}/pipelines/deploy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TriggerPipelineDeployReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TriggerPipelineDeployOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for triggerPipelineDeploy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TriggerPipelinePromote runs a promote pipeline for a given application and branch
*/
func (a *Client) TriggerPipelinePromote(params *TriggerPipelinePromoteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TriggerPipelinePromoteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTriggerPipelinePromoteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "triggerPipelinePromote",
		Method:             "POST",
		PathPattern:        "/applications/{appName}/pipelines/promote",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TriggerPipelinePromoteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TriggerPipelinePromoteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for triggerPipelinePromote: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateApplicationAlertingConfig updates alerts configuration for application namespace
*/
func (a *Client) UpdateApplicationAlertingConfig(params *UpdateApplicationAlertingConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateApplicationAlertingConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateApplicationAlertingConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateApplicationAlertingConfig",
		Method:             "PUT",
		PathPattern:        "/applications/{appName}/alerting",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateApplicationAlertingConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateApplicationAlertingConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateApplicationAlertingConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateBuildSecretsSecretValue updates an application build secret
*/
func (a *Client) UpdateBuildSecretsSecretValue(params *UpdateBuildSecretsSecretValueParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateBuildSecretsSecretValueOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateBuildSecretsSecretValueParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateBuildSecretsSecretValue",
		Method:             "PUT",
		PathPattern:        "/applications/{appName}/buildsecrets/{secretName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateBuildSecretsSecretValueReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateBuildSecretsSecretValueOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateBuildSecretsSecretValue: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdatePrivateImageHubsSecretValue updates an application private image hub secret
*/
func (a *Client) UpdatePrivateImageHubsSecretValue(params *UpdatePrivateImageHubsSecretValueParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdatePrivateImageHubsSecretValueOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePrivateImageHubsSecretValueParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updatePrivateImageHubsSecretValue",
		Method:             "PUT",
		PathPattern:        "/applications/{appName}/privateimagehubs/{serverName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdatePrivateImageHubsSecretValueReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdatePrivateImageHubsSecretValueOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updatePrivateImageHubsSecretValue: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}