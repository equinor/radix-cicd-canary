package egresspolicy

import (
	"fmt"
	"github.com/equinor/radix-cicd-canary/metrics"
	nspMetrics "github.com/equinor/radix-cicd-canary/metrics/scenarios/nsp"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-common/utils/errors"
	log "github.com/sirupsen/logrus"
)

// GetJobList tests that we are able to retrieve job list from job scheduler
func GetJobList(env env.Env, suiteName string) error {
	logger = log.WithFields(log.Fields{"Suite": suiteName})
	appEnvs := []string{"egressrulestopublicdns", "allowradix"}
	var errs []error
	for _, appEnv := range appEnvs {
		jobListUrl := fmt.Sprintf("%s/testjobscheduler", env.GetNetworkPolicyCanaryUrl(appEnv))
		err := httpUtils.CheckUrl(jobListUrl)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errors.Concat(errs)
	}
	return nil
}

// GetJobListSuccess is a function after a call to GetJobList succeeds
func GetJobListSuccess(testName string) {
	nspMetrics.AddGetJobListSuccess()
	metrics.AddTestOne(testName, nspMetrics.Success)
	metrics.AddTestZero(testName, nspMetrics.Errors)
	logger.Infof("Test %s: SUCCESS", testName)
}

// GetJobListFail is a function after a call to GetJobList failed
func GetJobListFail(testName string) {
	nspMetrics.AddGetJobListFail()
	metrics.AddTestZero(testName, nspMetrics.Success)
	metrics.AddTestOne(testName, nspMetrics.Errors)
	logger.Infof("Test %s: FAIL", testName)
}
