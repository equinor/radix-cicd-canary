package secret

import (
	"testing"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/stretchr/testify/assert"
)

const suiteName = "Happy path"

/*
Allow us to run this as a single test. Note that we need the previous tests of the suite to have passed for this to work.
Also running the test may fail, because it may time out.

Its best use is when debugging a single test
*/
func TestSetSecret(t *testing.T) {
	env.SetRequiredEnvironmentVariablesForTest()
	environmentVariables := env.NewEnv()

	err := Set(environmentVariables, suiteName)
	assert.NoError(t, err)
}
