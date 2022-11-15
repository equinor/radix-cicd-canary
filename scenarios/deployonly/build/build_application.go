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
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

type expectedStep struct {
	name       string
	components []string
}

// Application Tests that we are able to successfully build an application
func Application(env envUtil.Env, suiteName string) error {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	// Trigger build via web hook
	err := httpUtils.TriggerWebhookPush(env, config.App3BranchToBuildFrom, config.App3CommitID, config.App3SSHRepository, config.App3SharedSecret)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to push webhook push for App3, error %v", err))
	}

	// Get job
	ok, jobSummary := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return job.IsListedWithStatus(env, config.App3Name, "Succeeded")
	})

	if !ok {
		return errors.New(fmt.Sprintf("Could not get listed job for application %s status \"%s\" - exiting.", config.App3Name, "Succeeded"))
	}

	jobName := (jobSummary.(*models.JobSummary)).Name
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
