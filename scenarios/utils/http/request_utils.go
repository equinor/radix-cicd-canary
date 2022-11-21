package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	applicationAPIClient "github.com/equinor/radix-cicd-canary/generated-client/client/application"
	componentAPIClient "github.com/equinor/radix-cicd-canary/generated-client/client/component"
	deploymentAPIClient "github.com/equinor/radix-cicd-canary/generated-client/client/deployment"
	environmentAPIClient "github.com/equinor/radix-cicd-canary/generated-client/client/environment"
	jobAPIClient "github.com/equinor/radix-cicd-canary/generated-client/client/job"
	pipelineJobAPIClient "github.com/equinor/radix-cicd-canary/generated-client/client/pipeline_job"
	platformAPIClient "github.com/equinor/radix-cicd-canary/generated-client/client/platform"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/crypto"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
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
func CreateRequest(env env.Env, url, method string, parameters interface{}) *http.Request {
	var reader io.Reader
	if parameters != nil {
		payload, _ := json.Marshal(parameters)
		reader = bytes.NewReader(payload)
	}

	req, _ := http.NewRequest(method, url, reader)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", env.GetBearerToken()))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Impersonate-User", env.GetImpersonateUser())
	req.Header.Add("Impersonate-Group", env.GetImpersonateGroup())

	return req
}

// TriggerWebhookPush Makes call to webhook
func TriggerWebhookPush(env env.Env, branch, commit, repository, sharedSecret string, logger *log.Entry) error {
	parameters := WebhookPayload{
		Ref:   fmt.Sprintf("refs/heads/%s", branch),
		After: commit,
		Repository: Repository{
			SSHURL: repository,
		},
	}

	req := CreateRequest(env, fmt.Sprintf("%s/events/github", env.GetGitHubWebHookAPIURL()), "POST", parameters)
	client := http.DefaultClient
	payload, _ := json.Marshal(parameters)

	req.Header.Add("X-GitHub-Event", "push")
	req.Header.Add("X-Hub-Signature-256", crypto.SHA256HMAC([]byte(sharedSecret), payload))

	logger.Debugf("Trigger webhook push for '%s' branch of repository %s, for commit %s", branch, repository, commit)

	resp, err := client.Do(req)
	if err != nil {
		return errors.WithMessage(err,
			fmt.Sprintf("error trigger webhook push for '%s' branch of repository %s, for commit %s", branch, repository, commit))
	}

	if err := CheckResponse(resp, logger); err != nil {
		return errors.WithMessage(err,
			fmt.Sprintf("error checking webhook response for '%s' branch of repository %s, for commit %s", branch, repository, commit))
	}

	return nil
}

// CheckResponse Checks that the response was successful
func CheckResponse(resp *http.Response, logger *log.Entry) error {
	defer resp.Body.Close()
	_, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.WithMessage(err, "error reading response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		logger.Debugf("Response code: %d", resp.StatusCode)
		return nil
	}

	return fmt.Errorf("response status code is %d", resp.StatusCode)
}

// CheckUrl Checks that a GET request to specified URL returns 200 without errors
func CheckUrl(url string, logger *log.Entry) error {
	logger.Debugf("Sending request to %s", url)
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	return CheckResponse(response, logger)
}

// GetClientBearerToken Gets bearer token in order to make call to API server
func GetClientBearerToken(env env.Env) runtime.ClientAuthInfoWriter {
	return httptransport.BearerToken(env.GetBearerToken())
}

// GetPlatformClient Gets the Platform API client
func GetPlatformClient(env env.Env) platformAPIClient.ClientService {
	return platformAPIClient.New(getTransport(env), strfmt.Default)
}

// GetApplicationClient Gets the Application API client
func GetApplicationClient(env env.Env) applicationAPIClient.ClientService {
	return applicationAPIClient.New(getTransport(env), strfmt.Default)
}

// GetJobClient Gets the Job API client
func GetJobClient(env env.Env) pipelineJobAPIClient.ClientService {
	return pipelineJobAPIClient.New(getTransport(env), strfmt.Default)
}

// GetEnvironmentClient Gets the Environment API client
func GetEnvironmentClient(env env.Env) environmentAPIClient.ClientService {
	return environmentAPIClient.New(getTransport(env), strfmt.Default)
}

// GetDeploymentClient Gets the Deployment API client
func GetDeploymentClient(env env.Env) deploymentAPIClient.ClientService {
	return deploymentAPIClient.New(getTransport(env), strfmt.Default)
}

// GetComponentClient Gets the Component API client
func GetComponentClient(env env.Env) componentAPIClient.ClientService {
	return componentAPIClient.New(getTransport(env), strfmt.Default)
}

// GetK8sJobClient Gets the K8s job API client
func GetK8sJobClient(env env.Env) jobAPIClient.ClientService {
	return jobAPIClient.New(getTransport(env), strfmt.Default)
}

func getTransport(env env.Env) *httptransport.Runtime {
	radixAPIURL := env.GetRadixAPIURL()
	schemes := env.GetRadixAPISchemes()

	return httptransport.New(radixAPIURL, basePath, schemes)
}

func GetUrl(schema string, domainName string) string {
	if strings.HasPrefix("http://", domainName) || strings.HasPrefix("https://", domainName) {
		return domainName
	}
	return fmt.Sprintf("%s://%s", schema, domainName)
}
