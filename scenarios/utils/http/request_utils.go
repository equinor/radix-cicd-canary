package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	applicationAPIClient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/application"
	componentAPIClient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/component"
	deploymentAPIClient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/deployment"
	environmentAPIClient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/environment"
	jobAPIClient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/job"
	pipelineJobAPIClient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/pipeline_job"
	platformAPIClient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/platform"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/crypto"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

const basePath = "/api/v1"

// WebhookPayload For serializing payload
type WebhookPayload struct {
	Ref        string     `json:"ref"`
	After      string     `json:"after"`
	Repository Repository `json:"repository"`
}

// Repository For serializing payload -> Repository
type Repository struct {
	SSHURL string `json:"ssh_url"`
}

// GetHTTPDefaultClient returns a new simple HTTP client
func GetHTTPDefaultClient() *http.Client {
	return &http.Client{Timeout: time.Second * 5}
}

// CreateRequest setup correct header for running tests
func CreateRequest(url, method string, parameters interface{}) *http.Request {
	var reader io.Reader
	if parameters != nil {
		payload, _ := json.Marshal(parameters)
		reader = bytes.NewReader(payload)
	}

	req, _ := http.NewRequest(method, url, reader)
	req.Header.Add("Content-Type", "application/json")
	return req
}

// TriggerWebhookPush Makes call to webhook
func TriggerWebhookPush(ctx context.Context, cfg config.Config, branch, commit, repository, sharedSecret string) error {
	parameters := WebhookPayload{
		Ref:   fmt.Sprintf("refs/heads/%s", branch),
		After: commit,
		Repository: Repository{
			SSHURL: repository,
		},
	}

	req := CreateRequest(fmt.Sprintf("%s/events/github", cfg.GetGitHubWebHookAPIURL()), "POST", parameters)
	client := http.DefaultClient
	payload, _ := json.Marshal(parameters)

	req.Header.Add("X-GitHub-Event", "push")
	req.Header.Add("X-Hub-Signature-256", crypto.SHA256HMAC([]byte(sharedSecret), payload))

	log.Ctx(ctx).Debug().Str("branch", branch).Str("repository", repository).Str("comit", commit).Msg("Trigger webhook push")

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrapf(err,
			"error trigger webhook push for '%s' branch of repository %s, for commit %s", branch, repository, commit)
	}

	if err := CheckResponse(resp, ctx); err != nil {
		return errors.Wrapf(err, "error checking webhook response for '%s' branch of repository %s, for commit %s", branch, repository, commit)
	}

	return nil
}

// CheckResponse Checks that the response was successful
func CheckResponse(resp *http.Response, ctx context.Context) error {
	defer resp.Body.Close()
	_, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "error reading response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		log.Ctx(ctx).Debug().Int("statusCode", resp.StatusCode).Msg("Response code")
		return nil
	}

	return errors.Errorf("response status code is %d", resp.StatusCode)
}

// CheckUrl Checks that a GET request to specified URL returns 200 without errors
func CheckUrl(url string, ctx context.Context) error {
	log.Ctx(ctx).Debug().Str("url", url).Msg("Sending request")

	request, _ := http.NewRequest("GET", url, nil)
	request = request.WithContext(ctx)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return errors.WithStack(err)
	}
	return CheckResponse(response, ctx)
}

// GetPlatformClient Gets the Platform API client
func GetPlatformClient(cfg config.Config) platformAPIClient.ClientService {
	return platformAPIClient.New(getTransport(cfg), strfmt.Default)
}

// GetApplicationClient Gets the Application API client
func GetApplicationClient(cfg config.Config) applicationAPIClient.ClientService {
	return applicationAPIClient.New(getTransport(cfg), strfmt.Default)
}

// GetJobClient Gets the Job API client
func GetJobClient(cfg config.Config) pipelineJobAPIClient.ClientService {
	return pipelineJobAPIClient.New(getTransport(cfg), strfmt.Default)
}

// GetEnvironmentClient Gets the Environment API client
func GetEnvironmentClient(cfg config.Config) environmentAPIClient.ClientService {
	return environmentAPIClient.New(getTransport(cfg), strfmt.Default)
}

// GetDeploymentClient Gets the Deployment API client
func GetDeploymentClient(cfg config.Config) deploymentAPIClient.ClientService {
	return deploymentAPIClient.New(getTransport(cfg), strfmt.Default)
}

// GetComponentClient Gets the Component API client
func GetComponentClient(cfg config.Config) componentAPIClient.ClientService {
	return componentAPIClient.New(getTransport(cfg), strfmt.Default)
}

// GetK8sJobClient Gets the K8s job API client
func GetK8sJobClient(cfg config.Config) jobAPIClient.ClientService {
	return jobAPIClient.New(getTransport(cfg), strfmt.Default)
}

func GetUrl(schema string, domainName string) string {
	if strings.HasPrefix("http://", domainName) || strings.HasPrefix("https://", domainName) {
		return domainName
	}
	return fmt.Sprintf("%s://%s", schema, domainName)
}

func getTransport(cfg config.Config) *httptransport.Runtime {
	radixAPIURL := cfg.GetRadixAPIURL()
	schemes := cfg.GetRadixAPISchemes()
	transport := httptransport.New(radixAPIURL, basePath, schemes)
	transport.DefaultAuthentication = getClientAuthInfoWriter(cfg)
	return transport
}

func getClientAuthInfoWriter(cfg config.Config) runtime.ClientAuthInfoWriter {
	if ts := cfg.GetTokenSource(); ts != nil {
		return runtime.ClientAuthInfoWriterFunc(func(cr runtime.ClientRequest, r strfmt.Registry) error {
			token, err := ts.Token()
			if err != nil {
				return err
			}
			cr.SetHeaderParam(runtime.HeaderAuthorization, "Bearer "+token.AccessToken)
			return nil
		})
	}

	return nil
}
