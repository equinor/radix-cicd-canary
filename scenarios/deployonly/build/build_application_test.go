package build

import (
	"context"
	"testing"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/stretchr/testify/assert"
)

const suiteName = "Deploy only"

/*
Allow us to run this as a single test. Note that we need the previous tests of the suite to have passed for this to work.
Also running the test may fail, because it may time out.

Its best use is when debugging a single test
*/
func TestBuildApplicationCreated(t *testing.T) {
	config.SetRequiredEnvironmentVariablesForTest()
	environmentVariables := config.NewConfig()

	err := Application(context.Background(), environmentVariables, suiteName)
	assert.NoError(t, err)
}
