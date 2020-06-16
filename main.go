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
	suites = filterSuites(suites, environmentVariables)
	if len(suites) == 0 {
		return
	}

	runner := test.NewRunner(environmentVariables)
	for {
		runner.Run(suites...)
		time.Sleep(sleepInterval)
	}
}

func filterSuites(suites []test.Suite, environmentVariables env.Env) []test.Suite {
	filter := environmentVariables.GetSuiteList()
	if len(filter) == 0 {
		return suites
	}

	suitesToRun := make([]test.Suite, len(suites))
	isBlacklist := environmentVariables.GetSuiteListIsBlacklist()
	for _, suite := range suites {
		// pass the filter if mentioned and !isBlacklist OR if !mentioned and isBlacklist
		if contains(filter, suite.Name) != isBlacklist {
			suitesToRun = append(suitesToRun, suite)
		}
	}
	return suitesToRun
}

func contains(list []string, target string) bool {
	for _, item := range list {
		if target == item {
			return true
		}
	}
	return false
}
