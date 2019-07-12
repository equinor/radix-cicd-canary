package happypath

import (
	"net/http"

	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils"
	log "github.com/sirupsen/logrus"
)

const (
	listAppPath     = "/api/v1/applications"
	listAppMethod   = "GET"
	listAppTestName = "ListApplications"
)

func listApplications() string {
	log.Infof("Sending HTTP GET request...")

	req := utils.CreateHTTPRequest(listAppPath, listAppMethod)
	client := http.DefaultClient

	resp, err := client.Do(req)

	if err != nil {
		addTestError(listAppTestName)
		log.Errorf("HTTP GET error: %v", err)
	} else {
		if resp.StatusCode == 200 {
			addTestSuccess(listAppTestName)
			log.Infof("Response: %s", resp.Status)
		} else {
			addTestError(listAppTestName)
			log.Errorf("Error response code: %v", resp.StatusCode)
		}
	}
	return listAppTestName
}
