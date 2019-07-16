package happypath

import (
	"net/http"

	models "github.com/equinor/radix-cicd-canary-golang/scenarios/happypath/models"
	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils"
	log "github.com/sirupsen/logrus"
)

func registerApplication() string {
	const (
		path     = "/api/v1/applications"
		method   = "POST"
		testName = "RegisterApplication"
	)

	log.Infof("Sending HTTP POST request...")

	parameters := models.ApplicationRegistration{
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
