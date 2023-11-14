package deploy

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/array"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/job"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
)

type expectedStep struct {
	name       string
	components []string
}

// Application Tests that we are able to successfully deploy an application by calling Radix API server
func Application(ctx context.Context, cfg config.Config) error {

	// Trigger deploy via Radix API
	_, err := application.Deploy(ctx, cfg, defaults.App3Name, defaults.App3EnvironmentName)
	if err != nil {
		return errors.Errorf("failed to deploy the application %s:  %v", defaults.App3Name, err)
	}

	// Get job
	jobSummary, err := test.WaitForCheckFuncWithValueOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) (*models.JobSummary, error) {
		jobSummary, err := job.GetLastPipelineJobWithStatus(ctx, cfg, defaults.App3Name, "Succeeded")
		if err != nil {
			return nil, err
		}
		return jobSummary, err
	})
	if err != nil {
		return errors.WithStack(err)
	}
	if jobSummary == nil {
		return errors.Errorf("could not get listed job for application %s status '%s'", defaults.App3Name, "Succeeded")
	}

	jobName := jobSummary.Name
	steps := job.GetSteps(ctx, cfg, defaults.App3Name, jobName)

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
			return errors.Errorf("expected step %s, but got %s", expectedSteps[index].name, step.Name)
		}

		if !array.EqualElements(step.Components, expectedSteps[index].components) {
			return errors.Errorf("expected components %s, but got %s", expectedSteps[index].components, step.Components)
		}
	}

	return nil
}
