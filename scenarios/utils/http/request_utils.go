package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	applicationAPIClient "github.com/equinor/radix-cicd-canary-golang/generated-client/client/application"
	environmentAPIClient "github.com/equinor/radix-cicd-canary-golang/generated-client/client/environment"
	jobAPIClient "github.com/equinor/radix-cicd-canary-golang/generated-client/client/job"
	platformAPIClient "github.com/equinor/radix-cicd-canary-golang/generated-client/client/platform"
	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils/crypto"
	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils/env"
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

// CreateRequest setup correct header for running tests
func CreateRequest(url, method string, parameters interface{}) *http.Request {
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
func TriggerWebhookPush(branch, commit, repository, sharedSecret string) bool {
	parameters := WebhookPayload{
		Ref:   fmt.Sprintf("refs/heads/%s", branch),
		After: commit,
		Repository: Repository{
			SSHURL: repository,
		},
	}

	req := CreateRequest(env.GetWebhookURL(), "POST", parameters)
	client := http.DefaultClient
	payload, _ := json.Marshal(parameters)

	req.Header.Add("X-GitHub-Event", "push")
	req.Header.Add("X-Hub-Signature", crypto.SHA1HMAC([]byte(sharedSecret), payload))
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Error %v", err)
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
		return true
	}
	return false
}

// GetClientBearerToken Gets bearer token in order to make call to API server
func GetClientBearerToken() runtime.ClientAuthInfoWriter {
	return httptransport.BearerToken(env.GetBearerToken())
}

// GetPlatformClient Gets the Platform API client
func GetPlatformClient() *platformAPIClient.Client {
	return platformAPIClient.New(getTransport(), strfmt.Default)
}

// GetApplicationClient Gets the Application API client
func GetApplicationClient() *applicationAPIClient.Client {
	return applicationAPIClient.New(getTransport(), strfmt.Default)
}

// GetJobClient Gets the Job API client
func GetJobClient() *jobAPIClient.Client {
	return jobAPIClient.New(getTransport(), strfmt.Default)
}

// GetEnvironmentClient Gets the Environment API client
func GetEnvironmentClient() *environmentAPIClient.Client {
	return environmentAPIClient.New(getTransport(), strfmt.Default)
}

func getTransport() *httptransport.Runtime {
	radixAPIURL := env.GetRadixAPIURL()
	schemes := []string{"https"}

	return httptransport.New(radixAPIURL, basePath, schemes)
}
