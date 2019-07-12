package happypath

import (
	"net/http"

	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils"
	log "github.com/sirupsen/logrus"
)

func listApplications() string {
	const (
		path     = "/api/v1/applications"
		method   = "GET"
		testName = "ListApplications"
	)

	log.Infof("Sending HTTP GET request...")

	req := utils.CreateHTTPRequest(path, method)
	client := http.DefaultClient

	resp, err := client.Do(req)

	if err != nil {
		addTestError(testName)
		log.Errorf("HTTP GET error: %v", err)
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
