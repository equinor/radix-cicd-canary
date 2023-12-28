package egresspolicy

import (
	"context"
	"errors"
	"fmt"

	"github.com/equinor/radix-cicd-canary/metrics"
	nspMetrics "github.com/equinor/radix-cicd-canary/metrics/scenarios/nsp"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/rs/zerolog/log"
)

// GetJobList tests that we are able to retrieve job list from job scheduler
func GetJobList(ctx context.Context, cfg config.Config) error {
	appEnvs := []string{"egressrulestopublicdns", "allowradix"}
	var errs []error
	for _, appEnv := range appEnvs {
		jobListUrl := fmt.Sprintf("%s/testjobscheduler", cfg.GetNetworkPolicyCanaryUrl(appEnv))
		err := httpUtils.CheckUrl(jobListUrl, ctx)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

// GetJobListSuccess is a function after a call to GetJobList succeeds
func GetJobListSuccess(ctx context.Context, testName string) {
	nspMetrics.AddGetJobListSuccess()
	metrics.AddTestOne(testName, nspMetrics.Success)
	metrics.AddTestZero(testName, nspMetrics.Errors)
	log.Ctx(ctx).Info().Msg("Test: SUCCESS")
}

// GetJobListFail is a function after a call to GetJobList failed
func GetJobListFail(ctx context.Context, testName string) {
	nspMetrics.AddGetJobListFail()
	metrics.AddTestZero(testName, nspMetrics.Success)
	metrics.AddTestOne(testName, nspMetrics.Errors)
	log.Ctx(ctx).Info().Msg("Test: FAIL")
}
