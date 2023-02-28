package egresspolicy

import (
	"testing"

	envUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/stretchr/testify/assert"
)

const jobSchedulerBaseUrl = "http://127.0.0.1:6000"

/*
Allow us to run this as a single test. Note that we need the previous tests of the suite to have passed for this to work.
Also running the test may fail, because it may time out.

Its best use is when debugging a single test
*/
func TestStartJobBatch(t *testing.T) {
	envUtils.SetRequiredEnvironmentVariablesForTest()
	cfg := envUtils.NewConfig()

	err := startJobBatch(jobSchedulerBaseUrl, "passwordSentInHeaderFromCicdCanary", "someappenv")
	assert.NoError(t, err)

	err = startJobBatch(jobSchedulerBaseUrl, "wrongPassword", "someappenv")
	assert.Error(t, err)

	err = startJobBatch(jobSchedulerBaseUrl, cfg.GetNetworkPolicyCanaryPassword(), "someappenv")
	assert.Error(t, err)
}
