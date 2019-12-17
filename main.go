package main

import (
	"net/http"
	"time"

	"github.com/equinor/radix-cicd-canary/scenarios/deployonly"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath"
	"github.com/equinor/radix-cicd-canary/scenarios/nsp"
	"github.com/equinor/radix-cicd-canary/scenarios/test"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func init() {
	// If you get GOAWAY calling API with token using:
	// az account get-access-token
	// ...enable this line
	//os.Setenv("GODEBUG", "http2server=0,http2client=0")

}

func main() {
	log.Infof("Starting...")

	environmentVariables := env.NewEnv()

	sleepInterval := environmentVariables.GetSleepIntervalBetweenTestRuns()
	happyPathSuite := happypath.TestSuite()
	deployOnlySuite := deployonly.TestSuite()

	nspSleepInterval := environmentVariables.GetNSPSleepInterval()
	nspSuite := nsp.TestSuite()

	go runSuites(environmentVariables, sleepInterval, happyPathSuite)
	go runSuites(environmentVariables, sleepInterval, deployOnlySuite)
	go runSuites(environmentVariables, nspSleepInterval, nspSuite)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":5000", nil)
}

func runSuites(environmentVariables env.Env, sleepInterval time.Duration, suites ...test.Suite) {
	runner := test.NewRunner(environmentVariables)
	for {
		runner.Run(suites...)
		time.Sleep(sleepInterval)
	}
}
