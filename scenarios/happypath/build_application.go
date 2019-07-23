package happypath

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"

	httpUtils "github.com/equinor/radix-cicd-canary-golang/scenarios/utils/http"
	log "github.com/sirupsen/logrus"
)

const (
	testName = "BuildApplications"
	basePath = "/api/v1"
)

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

func buildApplications() string {
	buildApplication(app1Name, testName)

	return testName
}

func buildApplication(appName, testName string) {
	// Trigger build via web hook
	var payload = WebhookPayload{
		Ref:   fmt.Sprintf("refs/heads/%s", app2BranchToBuildFrom),
		After: app2CommitID,
		Repository: Repository{
			SSHURL: app2SSHRepository,
		},
	}

	ok := triggerWebhook("push", payload, app2SharedSecret)
	if ok {
		addTestSuccess(testName)
	} else {
		addTestError(testName)
	}

	//

	/*
		radixAPIURL := utils.GetRadixAPIURL()
		impersonateUser := utils.GetImpersonateUser()
		impersonateGroup := utils.GetImpersonateGroup()
		bearerToken := utils.GetBearerToken()

		params := apiclient.NewTriggerPipelineParams().
			WithImpersonateUser(&impersonateUser).
			WithImpersonateGroup(&impersonateGroup).
			WithAppName(app2Name).
			WithPipelineName("build-deploy").
			WithPipelineParameters(

				&models.PipelineParameters{
					PipelineParametersBuild: models.PipelineParametersBuild{
						Branch: "master",
					},
				},
			)

		clientBearerToken := httptransport.BearerToken(bearerToken)
		schemes := []string{"https"}

		transport := httptransport.New(radixAPIURL, basePath, schemes)
		client := apiclient.New(transport, strfmt.Default)

		_, err := client.TriggerPipeline(params, clientBearerToken)
		if err != nil {
			addTestError(testName)
			log.Errorf("Error calling Build for application %s: %v", appName, err)
		}*/

}

func triggerWebhook(event string, parameters interface{}, sharedSecret string) bool {
	req := httpUtils.CreateRequest(utils.GetWebhookURL(), "POST", parameters)
	client := http.DefaultClient
	payload, _ := json.Marshal(parameters)

	req.Header.Add("X-GitHub-Event", event)
	req.Header.Add("X-Hub-Signature", SHA1HMAC([]byte(sharedSecret), payload))
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Error %v", err)
		return false
	}

	return checkResponse(resp)
}

// SHA1HMAC computes the GitHub SHA1 HMAC.
func SHA1HMAC(salt, message []byte) string {
	// GitHub creates a SHA1 HMAC, where the key is the GitHub secret and the
	// message is the JSON body.
	digest := hmac.New(sha1.New, salt)
	digest.Write(message)
	sum := digest.Sum(nil)
	return fmt.Sprintf("sha1=%x", sum)
}
