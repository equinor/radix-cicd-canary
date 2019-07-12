package happypath

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

const (
	listAppEndpoint = "https://server-radix-api-prod.weekly-27.dev.radix.equinor.com/api/v1/applications"
	listAppMethod   = "GET"
)

func listApplications() {
	log.Infof("Sending HTTP GET request...")

	start := time.Now()
	req := utils.CreateHTTPRequest(listAppEndpoint, listAppMethod)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		errors.With(prometheus.Labels{"testName": "ListApplications"}).Add(1)
		log.Errorf("HTTP GET error: %v", err)
	} else {
		if resp.StatusCode == 200 {
			errors.With(prometheus.Labels{"testName": "ListApplications"}).Add(0)
			log.Infof("Response: %s", resp.Status)
		} else {
			errors.With(prometheus.Labels{"testName": "ListApplications"}).Add(1)
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

	testDurations.With(prometheus.Labels{"testName": "ListApplications"}).Add(elapsed.Seconds())
	log.Infof("Elapsed time: %v", elapsed)

	time.Sleep(10 * time.Second)
}
