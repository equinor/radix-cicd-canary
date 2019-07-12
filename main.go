package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

var (
	errors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "radix_test_errors",
			Help: "Test errors",
		},
		[]string{"testName"},
	)
	testDurations = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "radix_test_duration",
			Help: "Duration of test",
		},
		[]string{"testName"},
	)
)

func main() {
	log.Infof("Starting...")

	go runTest()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":5000", nil)
}

func runTest() {
	method := "GET"
	endpoint := "https://server-radix-api-prod.weekly-27.dev.radix.equinor.com/api/v1/applications"

	req, _ := http.NewRequest(method, endpoint, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", getBearerToken()))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Impersonate-User", getImpersonateUser())
	req.Header.Add("Impersonate-Group", getImpersonateGroup())

	log.Infof("Bearer token: %s", getBearerToken())
	log.Infof("Impersonate-user: %s", getImpersonateUser())
	log.Infof("Impersonate-group: %s", getImpersonateGroup())

	client := http.DefaultClient

	for {
		log.Infof("Sending HTTP GET request...")

		start := time.Now()

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
}

func getBearerToken() string {
	return "xx"
}

func getImpersonateUser() string {
	return os.Getenv("IMPERSONATE_USER")
}

func getImpersonateGroup() string {
	return os.Getenv("IMPERSONATE_GROUP")
}
