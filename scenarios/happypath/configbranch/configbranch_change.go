package configbranch

import (
	"fmt"
	"strings"

	apiclient "github.com/equinor/radix-cicd-canary/generated-client/client/application"
	models "github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/array"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
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

// Change Tests that radixconfig is read from the branch defined as configBanch
func Change(env envUtil.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	// Trigger first build via web hook
	ok, err := httpUtils.TriggerWebhookPush(env, config.App4ConfigBranch, config.App4CommitID, config.App4SSHRepository, config.App4SharedSecret)
	if !ok {
		return false, err
	}

	logger.Infof("First job was triggered")
	jobSummary, err := waitForJobRunning(env)

	if err != nil {
		return false, errors.WithMessage(err, fmt.Sprintf("first job for application %s", config.App4Name))
	}

	jobName := jobSummary.Name
	logger.Infof("First job name: %s", jobName)

	if ok, err = waitForJobDone(env, jobName); !ok {
		return false, errors.WithMessage(err, fmt.Sprintf("first job for application %s", config.App4Name))
	}

	logger.Info("First job was completed")

	expectedSteps := []expectedStep{
		{name: "clone-config", components: []string{}},
		{name: "prepare-pipelines", components: []string{}},
		{name: "radix-pipeline", components: []string{}},
		{name: "clone", components: []string{}},
		{name: "build-www", components: []string{"www"}},
		{name: "run-pipelines", components: []string{}},
		{name: "scan-www", components: []string{"www"}},
	}

	if ok, err := validateJobSteps(env, jobName, expectedSteps); !ok {
		return false, err
	}

	// Change config branch, trigger second webhook and verify job
	if err := patchConfigBranch(env, config.App4NewConfigBranch); err != nil {
		return false, err
	}

	ok, err = httpUtils.TriggerWebhookPush(env, config.App4NewConfigBranch, config.App4NewCommitID, config.App4SSHRepository, config.App4SharedSecret)
	if !ok {
		return false, err
	}

	logger.Infof("Second job was triggered")
	jobSummary, err = waitForJobRunning(env)

	if err != nil {
		return false, errors.WithMessage(err, fmt.Sprintf("second job for application %s", config.App4Name))
	}

	jobName = jobSummary.Name
	logger.Infof("Second job name: %s", jobName)

	if ok, err = waitForJobDone(env, jobName); !ok {
		return false, errors.WithMessage(err, fmt.Sprintf("second job for application %s", config.App4Name))
	}

	logger.Info("Second job was completed")

	expectedSteps = []expectedStep{
		{name: "clone-config", components: []string{}},
		{name: "prepare-pipelines", components: []string{}},
		{name: "radix-pipeline", components: []string{}},
		{name: "clone", components: []string{}},
		{name: "build-www2", components: []string{"www2"}},
		{name: "run-pipelines", components: []string{}},
		{name: "scan-www2", components: []string{"www2"}},
	}

	if ok, err := validateJobSteps(env, jobName, expectedSteps); !ok {
		return false, err
	}

	return true, nil
}

func waitForJobRunning(env envUtil.Env) (*models.JobSummary, error) {
	status := "Running"

	ok, obj := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return job.IsListedWithStatus(env, config.App4Name, status)
	})

	if !ok {
		return nil, fmt.Errorf("could not get job with status %s", status)
	}

	if jobSummary, ok := obj.(*models.JobSummary); ok {
		return jobSummary, nil
	}

	return nil, fmt.Errorf("could not unmarshal jobSummary")
}

func waitForJobDone(env envUtil.Env, jobName string) (bool, error) {
	ok, status := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return job.IsDone(env, config.App4Name, jobName)
	})

	if !ok {
		return false, fmt.Errorf("job %s did not complete within the specified timeout period", jobName)
	}
	if status.(string) != "Succeeded" {
		return false, fmt.Errorf("job %s completed with status %s", jobName, status)
	}

	return true, nil
}

func patchConfigBranch(env env.Env, newConfigBranch string) error {
	logger.Debugf("Set ConfigBranch to %v", newConfigBranch)
	patchRequest := models.ApplicationPatchRequest{
		ConfigBranch: newConfigBranch,
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

func validateJobSteps(env env.Env, jobName string, expectedSteps []expectedStep) (bool, error) {
	steps := job.GetSteps(env, config.App4Name, jobName)

	if steps == nil && len(steps) != len(expectedSteps) {
		return false, fmt.Errorf("pipeline steps was not as expected")
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
