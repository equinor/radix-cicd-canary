package build

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/equinor/radix-cicd-canary/generated-client/models"
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
func Application(env envUtil.Env, suiteName string) error {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	// Trigger build via web hook
	err := httpUtils.TriggerWebhookPush(env, config.App2BranchToBuildFrom, config.App2CommitID, config.App2SSHRepository, config.App2SharedSecret)
	if err != nil {
		return err
	}
	logger.Infof("First job was triggered")

	// Get job
	jobSummary, err := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (*models.JobSummary, error) {
		return job.IsListedWithStatus(env, config.App2Name, "Running")
	})

	if err != nil {
		return err
	}

	jobName := jobSummary.Name
	logger.Infof("First job name: %s", jobName)

	// Another build should cause second job to queue up
	// Trigger another build via web hook
	time.Sleep(1 * time.Second)
	err = httpUtils.TriggerWebhookPush(env, config.App2BranchToBuildFrom, config.App2CommitID, config.App2SSHRepository, config.App2SharedSecret)
	if err != nil {
		return err
	}
	logger.Infof("Second job was triggered")

	_, err = test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (*models.JobSummary, error) {
		return job.IsListedWithStatus(env, config.App2Name, "Queued")
	})

	if err != nil {
		return err
	}

	logger.Info("Second job was queued")
	_, err = test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (string, error) {
		return job.IsDone(env, config.App2Name, jobName)
	})

	if err != nil {
		return err
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
	}

	if len(steps) != len(expectedSteps) {
		return errors.New("number of pipeline steps was not as expected")
	}

	for index, step := range steps {
		if !strings.EqualFold(step.Name, expectedSteps[index].name) {
			return errors.New(fmt.Sprintf("Expeced step %s, but got %s", expectedSteps[index].name, step.Name))
		}

		if !array.EqualElements(step.Components, expectedSteps[index].components) {
			return errors.New(fmt.Sprintf("Expeced components %s, but got %s", expectedSteps[index].components, step.Components))
		}
	}

	stepLog := job.GetLogForStep(env, config.App2Name, jobName, "build-app")
	//Validate if Dockerfile build output contains SHA256 hash of build secrets:
	//https://github.com/equinor/radix-canarycicd-test-2/blob/master/Dockerfile#L9
	if !strings.Contains(stepLog, Secret1ValueSha256) || !strings.Contains(stepLog, Secret2ValueSha256) {
		return errors.New("build secrets are not contained in build log")
	}

	jobSummary, err = test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (*models.JobSummary, error) {
		return job.IsListedWithStatus(env, config.App2Name, "Running")
	})

	if err != nil {
		return err
	}

	// Stop job and verify that it has been stopped
	jobName = jobSummary.Name
	logger.Infof("Second job name: %s", jobName)
	err = job.Stop(env, config.App2Name, jobName)
	if err != nil {
		return err
	}

	_, err = test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (*models.JobSummary, error) {
		return job.IsListedWithStatus(env, config.App2Name, "Stopped")
	})

	if err != nil {
		return err
	}

	logger.Info("Second job was stopped")
	return nil
}
