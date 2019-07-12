package happypath

import (
	"fmt"
	"net/http"

	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils"
	log "github.com/sirupsen/logrus"
)

const (
	deleteAppPath     = "/api/v1/applications"
	deleteAppMethod   = "DELETE"
	deleteAppTestName = "DeleteApplication"
)

func deleteApplications() string {
	log.Infof("Sending HTTP GET request...")
	deleteApplication(app1Name)
	deleteApplication(app2Name)

	return deleteAppTestName
}

func deleteApplication(appName string) {
	req := utils.CreateHTTPRequest(deleteSpecificAppPath(appName), deleteAppMethod)
	client := http.DefaultClient

	resp, err := client.Do(req)

	if err != nil {
		addTestError(deleteAppTestName)
		log.Errorf("HTTP DELETE error: %v", err)
	} else {
		if resp.StatusCode == 200 {
			addTestSuccess(deleteAppTestName)
			log.Infof("Response: %s", resp.Status)
		} else {
			addTestError(deleteAppTestName)
			log.Errorf("Error response code: %v", resp.StatusCode)
		}
	}
}

func deleteSpecificAppPath(appName string) string {
	return fmt.Sprintf("%s/%s", deleteAppPath, appName)
}
