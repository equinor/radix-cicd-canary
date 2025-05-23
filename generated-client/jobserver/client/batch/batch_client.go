// Code generated by go-swagger; DO NOT EDIT.

package batch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new batch API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new batch API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new batch API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for batch API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateBatch(params *CreateBatchParams, opts ...ClientOption) (*CreateBatchOK, error)

	DeleteBatch(params *DeleteBatchParams, opts ...ClientOption) (*DeleteBatchOK, error)

	GetBatch(params *GetBatchParams, opts ...ClientOption) (*GetBatchOK, error)

	GetBatchJob(params *GetBatchJobParams, opts ...ClientOption) (*GetBatchJobOK, error)

	GetBatches(params *GetBatchesParams, opts ...ClientOption) (*GetBatchesOK, error)

	StopAllBatches(params *StopAllBatchesParams, opts ...ClientOption) (*StopAllBatchesOK, error)

	StopBatch(params *StopBatchParams, opts ...ClientOption) (*StopBatchOK, error)

	StopBatchJob(params *StopBatchJobParams, opts ...ClientOption) (*StopBatchJobOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateBatch creates batch
*/
func (a *Client) CreateBatch(params *CreateBatchParams, opts ...ClientOption) (*CreateBatchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBatchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createBatch",
		Method:             "POST",
		PathPattern:        "/batches",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateBatchReader{formats: a.formats},
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
	success, ok := result.(*CreateBatchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createBatch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteBatch deletes batch
*/
func (a *Client) DeleteBatch(params *DeleteBatchParams, opts ...ClientOption) (*DeleteBatchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBatchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteBatch",
		Method:             "DELETE",
		PathPattern:        "/batches/{batchName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteBatchReader{formats: a.formats},
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
	success, ok := result.(*DeleteBatchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteBatch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBatch gets batch
*/
func (a *Client) GetBatch(params *GetBatchParams, opts ...ClientOption) (*GetBatchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBatchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBatch",
		Method:             "GET",
		PathPattern:        "/batches/{batchName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetBatchReader{formats: a.formats},
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
	success, ok := result.(*GetBatchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBatch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBatchJob gets batch job
*/
func (a *Client) GetBatchJob(params *GetBatchJobParams, opts ...ClientOption) (*GetBatchJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBatchJobParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBatchJob",
		Method:             "GET",
		PathPattern:        "/batches/{batchName}/jobs/{jobName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetBatchJobReader{formats: a.formats},
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
	success, ok := result.(*GetBatchJobOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBatchJob: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBatches gets batches
*/
func (a *Client) GetBatches(params *GetBatchesParams, opts ...ClientOption) (*GetBatchesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBatchesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBatches",
		Method:             "GET",
		PathPattern:        "/batches/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetBatchesReader{formats: a.formats},
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
	success, ok := result.(*GetBatchesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBatches: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StopAllBatches stops all batches
*/
func (a *Client) StopAllBatches(params *StopAllBatchesParams, opts ...ClientOption) (*StopAllBatchesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopAllBatchesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "stopAllBatches",
		Method:             "POST",
		PathPattern:        "/batches/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopAllBatchesReader{formats: a.formats},
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
	success, ok := result.(*StopAllBatchesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for stopAllBatches: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StopBatch stops batch
*/
func (a *Client) StopBatch(params *StopBatchParams, opts ...ClientOption) (*StopBatchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopBatchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "stopBatch",
		Method:             "POST",
		PathPattern:        "/batches/{batchName}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopBatchReader{formats: a.formats},
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
	success, ok := result.(*StopBatchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for stopBatch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StopBatchJob stops batch job
*/
func (a *Client) StopBatchJob(params *StopBatchJobParams, opts ...ClientOption) (*StopBatchJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopBatchJobParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "stopBatchJob",
		Method:             "POST",
		PathPattern:        "/batches/{batchName}/jobs/{jobName}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopBatchJobReader{formats: a.formats},
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
	success, ok := result.(*StopBatchJobOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for stopBatchJob: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
