package configbranch

import (
	"fmt"
	"strings"

	apiclient "github.com/equinor/radix-cicd-canary/generated-client/client/application"
	"github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/array"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/job"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

type expectedStep struct {
	name       string
	components []string
}

// Change Tests that radixconfig is read from the branch defined as configBranch
func Change(env envUtil.Env, suiteName string) error {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	// Trigger first build via web hook
	err := httpUtils.TriggerWebhookPush(env, config.App4ConfigBranch, config.App4CommitID, config.App4SSHRepository, config.App4SharedSecret, logger)
	if err != nil {
		return err
	}

	logger.Infof("First job was triggered")
	jobSummary, err := waitForJobRunning(env)

	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("first job for application %s", config.App4Name))
	}

	jobName := jobSummary.Name
	logger.Infof("First job name: %s", jobName)

	if err = waitForJobDone(env, jobName); err != nil {
		return errors.WithMessage(err, fmt.Sprintf("first job for application %s", config.App4Name))
	}

	logger.Info("First job was completed")

	expectedSteps := []expectedStep{
		{name: "clone-config", components: []string{}},
		{name: "prepare-pipelines", components: []string{}},
		{name: "radix-pipeline", components: []string{}},
		{name: "clone", components: []string{}},
		{name: "build-www", components: []string{"www"}},
		//{name: "run-pipelines", components: []string{}},//skip due to there is no sub-pipeline
	}

	if ok, err := validateJobSteps(env, jobName, expectedSteps); !ok {
		return err
	}

	// Change config branch, trigger second webhook and verify job
	if err := patchConfigBranch(env, config.App4NewConfigBranch); err != nil {
		return err
	}

	err = httpUtils.TriggerWebhookPush(env, config.App4NewConfigBranch, config.App4NewCommitID, config.App4SSHRepository, config.App4SharedSecret, logger)
	if err != nil {
		return err
	}

	logger.Infof("Second job was triggered")
	jobSummary, err = waitForJobRunning(env)

	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("second job for application %s", config.App4Name))
	}

	jobName = jobSummary.Name
	logger.Infof("Second job name: %s", jobName)

	if err = waitForJobDone(env, jobName); err != nil {
		return errors.WithMessage(err, fmt.Sprintf("second job for application %s", config.App4Name))
	}

	logger.Info("Second job was completed")

	expectedSteps = []expectedStep{
		{name: "clone-config", components: []string{}},
		{name: "prepare-pipelines", components: []string{}},
		{name: "radix-pipeline", components: []string{}},
		{name: "clone", components: []string{}},
		{name: "build-www2", components: []string{"www2"}},
		//{name: "run-pipelines", components: []string{}},//skip due to there is no sub-pipeline
	}

	if ok, err := validateJobSteps(env, jobName, expectedSteps); !ok {
		return err
	}

	return nil
}

func waitForJobRunning(env envUtil.Env) (*models.JobSummary, error) {
	status := "Running"

	return test.WaitForCheckFuncWithValueOrTimeout(env, func(env envUtil.Env) (*models.JobSummary, error) {
		return job.IsListedWithStatus(env, config.App4Name, status, logger)
	}, logger)
}

func waitForJobDone(env envUtil.Env, jobName string) error {
	jobStatus, err := test.WaitForCheckFuncWithValueOrTimeout(env, func(env envUtil.Env) (string, error) {
		return job.IsDone(config.App4Name, jobName, env, logger)
	}, logger)
	if err != nil {
		return err
	}
	if jobStatus != "Succeeded" {
		return fmt.Errorf("job %s completed with status %s", jobName, jobStatus)
	}
	return nil
}

func patchConfigBranch(env envUtil.Env, newConfigBranch string) error {
	logger.Debugf("Set ConfigBranch to %v", newConfigBranch)
	patchRequest := models.ApplicationRegistrationPatchRequest{
		ApplicationRegistrationPatch: &models.ApplicationRegistrationPatch{
			ConfigBranch: newConfigBranch,
		},
	}

	params := apiclient.NewModifyRegistrationDetailsParams().
		WithAppName(config.App4Name).
		WithPatchRequest(&patchRequest)

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetApplicationClient(env)

	_, err := client.ModifyRegistrationDetails(params, clientBearerToken)
	if err != nil {
		return err
	}
	logger.Debugf("ConfigBranch has been set to %v", newConfigBranch)
	return nil
}

func validateJobSteps(env envUtil.Env, jobName string, expectedSteps []expectedStep) (bool, error) {
	steps := job.GetSteps(env, config.App4Name, jobName)

	if len(steps) != len(expectedSteps) {
		return false, fmt.Errorf("number of pipeline steps was not as expected")
	}

	for index, step := range steps {
		if !strings.EqualFold(step.Name, expectedSteps[index].name) {
			return false, fmt.Errorf("expeced step %s, but got %s", expectedSteps[index].name, step.Name)
		}

		if !array.EqualElements(step.Components, expectedSteps[index].components) {
			return false, fmt.Errorf("expeced components %s, but got %s", expectedSteps[index].components, step.Components)
		}
	}

	return true, nil
}
