package main

import (
	"net/http"
	"time"

	"github.com/equinor/radix-cicd-canary/scenarios/deployonly"
	"github.com/equinor/radix-cicd-canary/scenarios/happypath"
	"github.com/equinor/radix-cicd-canary/scenarios/nsp"
	nsplong "github.com/equinor/radix-cicd-canary/scenarios/nsp-long"
	"github.com/equinor/radix-cicd-canary/scenarios/test"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func init() {
	// If you get GOAWAY calling API with token using:
	// az account get-access-token
	// ...enable this line
	// os.Setenv("GODEBUG", "http2server=0,http2client=0")
}

func main() {
	log.Info("Starting...")

	cfg := config.NewConfig()
	logLevel := cfg.GetLogLevel()
	log.Infof("Log level: %v", logLevel)
	log.SetLevel(logLevel)

	sleepInterval := cfg.GetSleepIntervalBetweenTestRuns()
	happyPathSuite := happypath.TestSuite()
	deployOnlySuite := deployonly.TestSuite()

	nspSleepInterval := cfg.GetNSPSleepInterval()
	nspLongSleepInterval := cfg.GetNSPLongSleepInterval()
	nspSuite := nsp.TestSuite()
	nspLongSuite := nsplong.TestSuite()

	go runSuites(cfg, sleepInterval, happyPathSuite)
	go runSuites(cfg, sleepInterval, deployOnlySuite)
	go runSuites(cfg, nspSleepInterval, nspSuite)
	go runSuites(cfg, nspLongSleepInterval, nspLongSuite)

	log.Info("Started suites. Start metrics service.")
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Error(err)
		return
	}
	log.Info("Complete.")
}

func runSuites(environmentVariables config.Config, sleepInterval time.Duration, suites ...test.Suite) {
	log.Debugf("Prepare to run %d suite(s)", len(suites))
	suites = filterSuites(suites, environmentVariables)
	if len(suites) == 0 {
		log.Debug("No suites to run")
		return
	}

	log.Debugf("Run %d suite(s)", len(suites))
	runner := test.NewRunner(environmentVariables)
	for {
		runner.Run(suites...)
		time.Sleep(sleepInterval)
	}
}

func filterSuites(suites []test.Suite, environmentVariables config.Config) []test.Suite {
	filter := environmentVariables.GetSuiteList()
	if len(filter) == 0 {
		return suites
	}

	log.Debug("Filtering suites...")
	suitesToRun := make([]test.Suite, 0)
	isBlacklist := environmentVariables.GetSuiteListIsBlacklist()
	for _, suite := range suites {
		// pass the filter if mentioned and !isBlacklist OR if !mentioned and isBlacklist
		if contains(filter, suite.Name) != isBlacklist {
			log.Debugf("- run suite '%s'", suite.Name)
			suitesToRun = append(suitesToRun, suite)
		} else {
			log.Debugf("- skip suite '%s'", suite.Name)
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
