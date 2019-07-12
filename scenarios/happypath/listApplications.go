package happypath

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils"
	log "github.com/sirupsen/logrus"
)

const (
	listAppPath     = "/api/v1/applications"
	listAppMethod   = "GET"
	listAppTestName = "ListApplications"
)

func listApplications() {
	log.Infof("Sending HTTP GET request...")

	start := time.Now()
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

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Error reading response body: %v", err)
	}

	log.Infof(string(body))

	end := time.Now()
	elapsed := end.Sub(start)

	addTestDuration(listAppTestName, elapsed.Seconds())
	log.Infof("Elapsed time: %v", elapsed)
}
