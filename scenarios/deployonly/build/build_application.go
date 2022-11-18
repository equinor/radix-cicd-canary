package build

import (
	"errors"
	"fmt"

	"github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/job"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
)

type expectedStep struct {
	name       string
	components []string
}

// Application Tests that we are able to successfully build an application
func Application(env envUtil.Env, suiteName string) error {

	// Trigger build via web hook
	err := httpUtils.TriggerWebhookPush(env, config.App3BranchToBuildFrom, config.App3CommitID, config.App3SSHRepository, config.App3SharedSecret)
	if err != nil {
		return fmt.Errorf("failed to push webhook push for App3, error %v", err)
	}

	// Get job
	jobSummary, err := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (*models.JobSummary, error) {
		jobSummary, err := job.IsListedWithStatus(env, config.App3Name, "Succeeded")
		if err != nil {
			return nil, err
		}
		if jobSummary == nil {
			return nil, fmt.Errorf("could not get listed job for application %s status '%s' - exiting", config.App3Name, "Succeeded")
		}
		return jobSummary, err
	})

	if err != nil {
		return err
	}

	jobName := jobSummary.Name
	steps := job.GetSteps(env, config.App3Name, jobName)
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
