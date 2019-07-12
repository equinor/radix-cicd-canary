package happypath

import (
	"fmt"
	"net/http"

	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils"
	log "github.com/sirupsen/logrus"
)

const (
	unauthorizedAccessPath     = "/api/v1/applications"
	unauthorizedAccessMethod   = "GET"
	unauthorizedAccessTestName = "UnauthorizedAccess"
)

func unauthorizedAccess() string {
	log.Infof("Sending HTTP GET request...")

	req := utils.CreateHTTPRequest(fmt.Sprintf("%s/%s", listAppPath, restrictedApplicationName), unauthorizedAccessMethod)
	client := http.DefaultClient

	resp, err := client.Do(req)

	if err != nil {
		addTestError(unauthorizedAccessTestName)
		log.Errorf("HTTP GET error: %v", err)
	} else {
		if resp.StatusCode == 403 {
			addTestSuccess(unauthorizedAccessTestName)
			log.Infof("Response: %s", resp.Status)
		} else {
			addTestError(unauthorizedAccessTestName)
			log.Errorf("Error response code: %v", resp.StatusCode)
		}
	}
	return unauthorizedAccessTestName
}
