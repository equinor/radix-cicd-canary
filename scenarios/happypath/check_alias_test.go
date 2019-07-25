package happypath

import (
	"testing"

	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils/env"
	"github.com/stretchr/testify/assert"
)

/*
Allow us to run this as a single test. Note that we need the previous tests of the suite to have passed for this to work.
Also running the test may fail, because it may time out.

Its best use is when debugging a single test
*/
func TestDefaultAliasResponding(t *testing.T) {
	env.SetRequiredEnvironmentVariablesForTest()
	ok, err := defaultAliasResponding()
	assert.NoError(t, err)
	assert.True(t, ok)
}
