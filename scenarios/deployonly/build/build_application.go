package build

import (
	"context"
	"errors"
	"fmt"

	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/job"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/rs/zerolog/log"
)

type expectedStep struct {
	name       string
	components []string
}

// Application Tests that we are able to successfully build an application
func Application(ctx context.Context, cfg config.Config) error {
	appName := defaults.App3Name
	appCtx := log.Ctx(ctx).With().Str("app", appName).Logger().WithContext(ctx)

	// Trigger build via web hook
	err := httpUtils.TriggerWebhookPush(cfg, defaults.App3BranchToBuildFrom, defaults.App3CommitID, defaults.App3SSHRepository, defaults.App3SharedSecret, appCtx)
	if err != nil {
		return fmt.Errorf("failed to push webhook push for App3, error %v", err)
	}

	// Get job
	jobSummary, err := test.WaitForCheckFuncWithValueOrTimeout(appCtx, cfg, func(cfg config.Config, ctx context.Context) (*models.JobSummary, error) {
		jobSummary, err := job.GetLastPipelineJobWithStatus(ctx, cfg, appName, "Succeeded")
		if err != nil {
			return nil, err
		}
		return jobSummary, err
	})

	if err != nil {
		return err
	}
	if jobSummary == nil {
		return fmt.Errorf("could not get listed job for application %s status '%s'", appName, "Succeeded")
	}

	jobName := jobSummary.Name
	steps := job.GetSteps(ctx, cfg, appName, jobName)
	expectedSteps := []expectedStep{
		{name: "clone-config", components: []string{}},
		{name: "prepare-pipelines", components: []string{}},
		{name: "radix-pipeline", components: []string{}},
		{name: "run-pipelines", components: []string{}},
	}

	if steps == nil && len(steps) != len(expectedSteps) {
		return errors.New("pipeline steps was not as expected")
	}

	return nil
}
