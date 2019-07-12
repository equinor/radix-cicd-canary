package happypath

import (
	"net/http"

	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils"
	log "github.com/sirupsen/logrus"
)

type applicationRegistration struct {
	Name         string   `json:"name"`
	Repository   string   `json:"repository"`
	SharedSecret string   `json:"sharedSecret"`
	AdGroups     []string `json:"adGroups"`
	PublicKey    string   `json:"publicKey,omitempty"`
	PrivateKey   string   `json:"privateKey,omitempty"`
}

func registerApplication() string {
	const (
		path     = "/api/v1/applications"
		method   = "POST"
		testName = "RegisterApplication"
	)

	log.Infof("Sending HTTP POST request...")

	parameters := applicationRegistration{
		Name:         app2Name,
		Repository:   app2Repository,
		SharedSecret: app2SharedSecret,
		AdGroups:     nil,
		PublicKey:    utils.GetPublicKey(),
		PrivateKey:   utils.GetPrivateKeyBase64(),
	}

	req := utils.CreateHTTPRequest(path, method, parameters)
	client := http.DefaultClient

	resp, err := client.Do(req)

	if err != nil {
		addTestError(testName)
		log.Errorf("HTTP POST error: %v", err)
	} else {
		if resp.StatusCode == 200 {
			addTestSuccess(testName)
			log.Infof("Response: %s", resp.Status)
		} else {
			addTestError(testName)
			log.Errorf("Error response code: %v", resp.StatusCode)
		}
	}

	return testName
}
