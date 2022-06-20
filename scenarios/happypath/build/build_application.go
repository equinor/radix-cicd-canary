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
	Secret1            = "SECRET_1"
	Secret2            = "SECRET_2"
	Secret1Value       = "SECRET_1_VALUE"
	Secret2Value       = "SECRET_2_VALUE"
	Secret1ValueSha256 = "7cb08032ffb66e835ceeb10b849a8728440b0631ccb11f652b807534df26275e"
	Secret2ValueSha256 = "087f38fb04a52265ad5394fc20a6bfaa78c44bd58097dbcb690031a85b6e8313"
)

type expectedStep struct {
	name       string
	components []string
}

// Application Tests that we are able to successfully build an application
func Application(env envUtil.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	// Trigger build via web hook
	ok, err := httpUtils.TriggerWebhookPush(env, config.App2BranchToBuildFrom, config.App2CommitID, config.App2SSHRepository, config.App2SharedSecret)
	if !ok {
		return false, err
	}
	logger.Infof("First job was triggered")

	// Get job
	ok, jobSummary := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return job.IsListedWithStatus(env, config.App2Name, "Running")
	})

	if !ok {
		log.Errorf("Could not get listed job for application %s status \"%s\" - exiting.", config.App2Name, "Running")
		return false, nil
	}

	jobName := (jobSummary.(*models.JobSummary)).Name
	logger.Infof("First job name: %s", jobName)

	// Another build should cause second job to queue up
	// Trigger another build via web hook
	time.Sleep(1 * time.Second)
	ok, err = httpUtils.TriggerWebhookPush(env, config.App2BranchToBuildFrom, config.App2CommitID, config.App2SSHRepository, config.App2SharedSecret)
	if !ok {
		return false, err
	}
	logger.Infof("Second job was triggered")

	ok, jobSummary = test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return job.IsListedWithStatus(env, config.App2Name, "Queued")
	})

	if !ok {
		log.Errorf("Could not get listed job for application %s status \"%s\" - exiting.", config.App2Name, "Queued")
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
		{name: "prepare-pipelines", components: []string{}},
		{name: "radix-pipeline", components: []string{}},
		{name: "clone", components: []string{}},
		{name: "build-app", components: []string{"app"}},
		{name: "build-redis", components: []string{"redis"}},
		{name: "run-pipelines", components: []string{}},
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

	stepLog := job.GetLogForStep(env, config.App2Name, jobName, "build-app")
	//Validate if Dockerfile build output contains SHA256 hash of build secrets:
	//https://github.com/equinor/radix-canarycicd-test-2/blob/master/Dockerfile#L9
	if !strings.Contains(stepLog, Secret1ValueSha256) || !strings.Contains(stepLog, Secret2ValueSha256) {
		logger.Error("Build secrets are not contained in build log")
		return false, nil
	}

	ok, jobSummary = test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return job.IsListedWithStatus(env, config.App2Name, "Running")
	})

	if !ok {
		log.Errorf("Could not get listed job for application %s status \"%s\" - exiting.", config.App2Name, "Running")
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
		log.Errorf("Could not get listed job for application %s status \"%s\" - exiting.", config.App2Name, "Stopped")
		return false, nil
	}

	logger.Info("Second job was stopped")
	return true, nil
}
