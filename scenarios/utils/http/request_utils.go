package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	applicationAPIClient "github.com/equinor/radix-cicd-canary/generated-client/client/application"
	componentAPIClient "github.com/equinor/radix-cicd-canary/generated-client/client/component"
	deploymentAPIClient "github.com/equinor/radix-cicd-canary/generated-client/client/deployment"
	environmentAPIClient "github.com/equinor/radix-cicd-canary/generated-client/client/environment"
	jobAPIClient "github.com/equinor/radix-cicd-canary/generated-client/client/job"
	platformAPIClient "github.com/equinor/radix-cicd-canary/generated-client/client/platform"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/crypto"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
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

	// Append protocol
	url = fmt.Sprintf("https://%s", url)

	req, _ := http.NewRequest(method, url, reader)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", env.GetBearerToken()))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Impersonate-User", env.GetImpersonateUser())
	req.Header.Add("Impersonate-Group", env.GetImpersonateGroup())

	return req
}

// TriggerWebhookPush Makes call to webhook
func TriggerWebhookPush(env env.Env, branch, commit, repository, sharedSecret string) bool {
	parameters := WebhookPayload{
		Ref:   fmt.Sprintf("refs/heads/%s", branch),
		After: commit,
		Repository: Repository{
			SSHURL: repository,
		},
	}

	req := CreateRequest(env, fmt.Sprintf("%s.%s/events/github", env.GetWebhookPrefix(), env.GetClusterFQDN()), "POST", parameters)
	client := http.DefaultClient
	payload, _ := json.Marshal(parameters)

	req.Header.Add("X-GitHub-Event", "push")
	req.Header.Add("X-Hub-Signature", crypto.SHA1HMAC([]byte(sharedSecret), payload))

	log.Debugf("Trigger webhook push for \"%s\" branch of repository %s, for commit %s", branch, repository, commit)

	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Error TriggerWebhookPush %v", err)
		return false
	}

	return CheckResponse(resp)
}

// CheckResponse Checks that the response was successful
func CheckResponse(resp *http.Response) bool {
	defer resp.Body.Close()
	_, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Error reading response body: %v", err)
		return false
	}

	if resp.StatusCode == 200 {
		log.Debug("Response code: 200")
		return true
	}

	log.Debugf("Response code: %d", resp.StatusCode)
	return false
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
func GetJobClient(env env.Env) jobAPIClient.ClientService {
	return jobAPIClient.New(getTransport(env), strfmt.Default)
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

func getTransport(env env.Env) *httptransport.Runtime {
	radixAPIURL := fmt.Sprintf("%s.%s", env.GetRadixAPIPrefix(), env.GetClusterFQDN())
	schemes := []string{"https"}

	return httptransport.New(radixAPIURL, basePath, schemes)
}
