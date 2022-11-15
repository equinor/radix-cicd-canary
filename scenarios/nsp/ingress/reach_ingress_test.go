package ingress

import (
	"testing"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/stretchr/testify/assert"
)

const suiteName = "NSP"

/*
Allow us to run this as a single test. Note that we need the previous tests of the suite to have passed for this to work.
Also running the test may fail, because it may time out.

Its best use is when debugging a single test
*/
func TestReachIngress(t *testing.T) {
	env.SetRequiredEnvironmentVariablesForTest()
	environmentVariables := env.NewEnv()

	err := Reach(environmentVariables, suiteName)
	assert.NoError(t, err)
}
