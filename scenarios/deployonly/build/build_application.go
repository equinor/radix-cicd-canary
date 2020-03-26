package build

import (
	models "github.com/equinor/radix-cicd-canary/generated-client/models"
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
func Application(env envUtil.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	// Trigger build via web hook
	ok := httpUtils.TriggerWebhookPush(env, config.App3BranchToBuildFrom, config.App3CommitID, config.App3SSHRepository, config.App3SharedSecret)
	if !ok {
		return false, nil
	}

	// Get job
	ok, jobSummary := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return job.IsListedWithStatus(env, config.App3Name, "Succeeded")
	})

	if !ok {
		return false, nil
	}

	jobName := (jobSummary.(*models.JobSummary)).Name
	steps := job.GetSteps(env, config.App3Name, jobName)
	expectedSteps := []expectedStep{
		{name: "clone-config", components: []string{}},
		{name: "config-2-map", components: []string{}},
		{name: "radix-pipeline", components: []string{}}}

	if steps == nil && len(steps) != len(expectedSteps) {
		logger.Error("Pipeline steps was not as expected")
		return false, nil
	}

	return true, nil
}
