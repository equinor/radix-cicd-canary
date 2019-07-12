package happypath

import (
	"fmt"
	"net/http"

	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils"
	log "github.com/sirupsen/logrus"
)

func deleteApplications() string {
	const testname = "DeleteApplication"

	log.Infof("Sending HTTP GET request...")
	deleteApplication(app1Name, testname)
	deleteApplication(app2Name, testname)

	return testname
}

func deleteApplication(appName, testname string) {
	const (
		path   = "/api/v1/applications"
		method = "DELETE"
	)

	req := utils.CreateHTTPRequest(fmt.Sprintf("%s/%s", path, appName), method, nil)
	client := http.DefaultClient

	resp, err := client.Do(req)

	if err != nil {
		addTestError(testname)
		log.Errorf("HTTP DELETE error: %v", err)
	} else {
		if resp.StatusCode == 200 {
			addTestSuccess(testname)
			log.Infof("Response: %s", resp.Status)
		} else {
			addTestError(testname)
			log.Errorf("Error response code: %v", resp.StatusCode)
		}
	}
}
