package egresspolicy

import (
	"context"
	"testing"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

const suiteName = "NSP"

/*
Allow us to run this as a single test. Note that we need the previous tests of the suite to have passed for this to work.
Also running the test may fail, because it may time out.

Its best use is when debugging a single test
*/
func TestLookupInternalDNS(t *testing.T) {
	config.SetRequiredEnvironmentVariablesForTest()
	environmentVariables := config.NewConfig()
	ctx := log.With().Str("suite", suiteName).Logger().WithContext(context.Background())

	err := LookupInternalDNS(ctx, environmentVariables)
	assert.NoError(t, err)
}

func TestLookupPublicDNS(t *testing.T) {
	config.SetRequiredEnvironmentVariablesForTest()
	environmentVariables := config.NewConfig()
	ctx := log.With().Str("suite", suiteName).Logger().WithContext(context.Background())

	err := LookupPublicDNS(ctx, environmentVariables)
	assert.NoError(t, err)
}
