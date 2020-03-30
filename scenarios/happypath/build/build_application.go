package build

import (
	"strings"
	"time"

	models "github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/array"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/job"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

const (
	Secret1      = "SECRET_1"
	Secret2      = "SECRET_2"
	Secret1Value = "SECRET_1_VALUE"
	Secret2Value = "SECRET_2_VALUE"
)

type expectedStep struct {
	name       string
	components []string
}

// Application Tests that we are able to successfully build an application
func Application(env envUtil.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	// Trigger build via web hook
	ok := httpUtils.TriggerWebhookPush(env, config.App2BranchToBuildFrom, config.App2CommitID, config.App2SSHRepository, config.App2SharedSecret)
	if !ok {
		return false, nil
	}
	logger.Infof("First job was triggered")

	// Get job
	ok, jobSummary := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return job.IsListedWithStatus(env, config.App2Name, "Running")
	})

	if !ok {
		return false, nil
	}

	jobName := (jobSummary.(*models.JobSummary)).Name
	logger.Infof("First job name: %s", jobName)

	// Another build should cause second job to queue up
	// Trigger another build via web hook
	time.Sleep(1 * time.Second)
	ok = httpUtils.TriggerWebhookPush(env, config.App2BranchToBuildFrom, config.App2CommitID, config.App2SSHRepository, config.App2SharedSecret)
	if !ok {
		return false, nil
	}
	logger.Infof("Second job was triggered")

	ok, jobSummary = test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return job.IsListedWithStatus(env, config.App2Name, "Queued")
	})

	if !ok {
		return false, nil
	}

	logger.Info("Second job was queued")
	ok, status := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return job.IsDone(env, config.App2Name, jobName)
	})

	if !ok {
		return false, nil
	}
	if status.(string) != "Succeeded" {
		return false, nil
	}

	logger.Info("First job was completed")
	steps := job.GetSteps(env, config.App2Name, jobName)

	expectedSteps := []expectedStep{
		{name: "clone-config", components: []string{}},
		{name: "config-2-map", components: []string{}},
		{name: "radix-pipeline", components: []string{}},
		{name: "clone", components: []string{}},
		{name: "build-app", components: []string{"app"}},
		{name: "build-redis", components: []string{"redis"}},
		{name: "scan-app", components: []string{"app"}},
		{name: "scan-redis", components: []string{"redis"}}}

	if steps == nil && len(steps) != len(expectedSteps) {
		logger.Error("Pipeline steps was not as expected")
		return false, nil
	}

	for index, step := range steps {
		if !strings.EqualFold(step.Name, expectedSteps[index].name) {
			logger.Errorf("Expeced step %s, but got %s", expectedSteps[index].name, step.Name)
			return false, nil
		}

		if !array.EqualElements(step.Components, expectedSteps[index].components) {
			logger.Errorf("Expeced components %s, but got %s", expectedSteps[index].components, step.Components)
			return false, nil
		}
	}

	log := job.GetLogForStep(env, config.App2Name, jobName, "build-app")
	if !strings.Contains(log, Secret1Value) || !strings.Contains(log, Secret2Value) {
		logger.Error("Build secrets are not contained in build log")
		return false, nil
	}

	ok, jobSummary = test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return job.IsListedWithStatus(env, config.App2Name, "Running")
	})

	if !ok {
		return false, nil
	}

	// Stop job and verify that it has been stopped
	jobName = (jobSummary.(*models.JobSummary)).Name
	logger.Infof("Second job name: %s", jobName)
	ok = job.Stop(env, config.App2Name, jobName)
	if !ok {
		return false, nil
	}

	ok, jobSummary = test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return job.IsListedWithStatus(env, config.App2Name, "Stopped")
	})

	if !ok {
		return false, nil
	}

	logger.Info("Second job was stopped")
	return true, nil
}
