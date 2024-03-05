package deploy

import (
	"context"

	"github.com/pkg/errors"

	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	jobUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/job"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
)

// Application Tests that we are able to successfully deploy an application by calling Radix API server
func Application(ctx context.Context, cfg config.Config) error {

	// Trigger deploy via Radix API
	_, err := application.Deploy(ctx, cfg, defaults.App3Name, defaults.App3EnvironmentName)
	if err != nil {
		return errors.Errorf("failed to deploy the application %s:  %v", defaults.App3Name, err)
	}

	// Get job
	jobSummary, err := test.WaitForCheckFuncWithValueOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) (*models.JobSummary, error) {
		jobSummary, err := jobUtils.GetLastPipelineJobWithStatus(ctx, cfg, defaults.App3Name, "Succeeded")
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
	steps := jobUtils.GetSteps(ctx, cfg, defaults.App3Name, jobName)

	expectedSteps := jobUtils.NewExpectedSteps().
		Add("clone-config").
		Add("prepare-pipelines").
		Add("radix-pipeline")

	if len(steps) != expectedSteps.Count() {
		return errors.New("number of pipeline steps was not as expected")
	}

	for _, step := range steps {
		if !expectedSteps.HasStepWithComponent(step.Name, step.Components) {
			return errors.Errorf("missing expected step %s", step.Name)
		}
	}

	return nil
}
