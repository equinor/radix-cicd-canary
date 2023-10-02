package promote

import (
	"context"
	"testing"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

const suiteName = "Happy path"

/*
Allow us to run this as a single test. Note that we need the previous tests of the suite to have passed for this to work.
Also running the test may fail, because it may time out.

Its best use is when debugging a single test
*/
func TestPromoteDeploymentToAnotherEnvironment(t *testing.T) {
	config.SetRequiredEnvironmentVariablesForTest()
	environmentVariables := config.NewConfig()
	ctx := log.With().Str("suite", suiteName).Logger().WithContext(context.Background())

	err := DeploymentToAnotherEnvironment(ctx, environmentVariables)
	assert.NoError(t, err)
}

func TestPromoteDeploymentWithinEnvironment(t *testing.T) {
	config.SetRequiredEnvironmentVariablesForTest()
	environmentVariables := config.NewConfig()
	ctx := log.With().Str("suite", suiteName).Logger().WithContext(context.Background())

	err := DeploymentWithinEnvironment(ctx, environmentVariables)
	assert.NoError(t, err)
}
