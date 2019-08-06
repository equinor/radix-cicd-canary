package happypath

import (
	"testing"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/stretchr/testify/assert"
)

/*
Allow us to run this as a single test. Note that we need the previous tests of the suite to have passed for this to work.
Also running the test may fail, because it may time out.

Its best use is when debugging a single test
*/
func TestRegisterApplicationWithNoDeployKey(t *testing.T) {
	env.SetRequiredEnvironmentVariablesForTest()
	ok, err := registerApplicationWithNoDeployKey()
	assert.NoError(t, err)
	assert.True(t, ok)
}
