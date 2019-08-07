package main

import (
	"net/http"
	"time"

	"github.com/equinor/radix-cicd-canary/scenarios/adgroup"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath"
	"github.com/equinor/radix-cicd-canary/scenarios/promotion"
	"github.com/equinor/radix-cicd-canary/scenarios/test"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Infof("Starting...")

	go runSuites()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":5000", nil)
}

func runSuites() {
	environmentVariables := env.NewEnv()
	runner := test.NewRunner(environmentVariables)

	sleepInterval := environmentVariables.GetSleepIntervalBetweenTestRuns()
	happyPathSuite := happypath.TestSuite()
	promotionSuite := promotion.TestSuite()
	adgroupSuite := adgroup.TestSuite()

	for {
		runner.Run(happyPathSuite, promotionSuite, adgroupSuite)
		time.Sleep(sleepInterval)
	}
}
