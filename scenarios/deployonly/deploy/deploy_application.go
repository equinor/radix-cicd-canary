package deploy

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/array"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/job"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/rs/zerolog/log"
)

type expectedStep struct {
	name       string
	components []string
}

// Application Tests that we are able to successfully deploy an application by calling Radix API server
func Application(ctx context.Context, cfg config.Config, suiteName string) error {
	appName := defaults.App3Name
	toEnvironment := defaults.App3EnvironmentName
	appCtx := log.Ctx(ctx).With().Str("app", appName).Logger().WithContext(ctx)

	// Trigger deploy via Radix API
	_, err := application.Deploy(cfg, appName, toEnvironment)
	if err != nil {
		return fmt.Errorf("failed to deploy the application %s:  %v", appName, err)
	}

	// Get job
	jobSummary, err := test.WaitForCheckFuncWithValueOrTimeout(cfg, func(cfg config.Config, ctx context.Context) (*models.JobSummary, error) {
		jobSummary, err := job.GetLastPipelineJobWithStatus(cfg, appName, "Succeeded", ctx)
		if err != nil {
			return nil, err
		}
		return jobSummary, err
	}, appCtx)
	if err != nil {
		return err
	}
	if jobSummary == nil {
		return fmt.Errorf("could not get listed job for application %s status '%s'", defaults.App3Name, "Succeeded")
	}

	jobName := jobSummary.Name
	steps := job.GetSteps(cfg, defaults.App3Name, jobName)

	expectedSteps := []expectedStep{
		{name: "clone-config", components: []string{}},
		{name: "prepare-pipelines", components: []string{}},
		{name: "radix-pipeline", components: []string{}},
		{name: "run-pipelines", components: []string{}},
	}

	if steps == nil && len(steps) != len(expectedSteps) {
		return errors.New("pipeline steps was not as expected")
	}

	for index, step := range steps {
		if !strings.EqualFold(step.Name, expectedSteps[index].name) {
			return fmt.Errorf("expected step %s, but got %s", expectedSteps[index].name, step.Name)
		}

		if !array.EqualElements(step.Components, expectedSteps[index].components) {
			return fmt.Errorf("expected components %s, but got %s", expectedSteps[index].components, step.Components)
		}
	}

	return nil
}
