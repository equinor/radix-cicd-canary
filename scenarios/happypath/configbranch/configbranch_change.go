package configbranch

import (
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

	// Trigger build via web hook
	ok := httpUtils.TriggerWebhookPush(env, config.App4ConfigBranch, config.App4CommitID, config.App4SSHRepository, config.App4SharedSecret)
	if !ok {
		return false, nil
	}
	logger.Infof("First job was triggered")

	// Get first job
	ok, jobSummary := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return job.IsListedWithStatus(env, config.App4Name, "Running")
	})

	if !ok {
		log.Errorf("Could not get listed job for application %s status \"%s\" - exiting.", config.App4Name, "Running")
		return false, nil
	}

	jobName := (jobSummary.(*models.JobSummary)).Name
	logger.Infof("First job name: %s", jobName)

	ok, status := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return job.IsDone(env, config.App4Name, jobName)
	})

	if !ok {
		return false, nil
	}
	if status.(string) != "Succeeded" {
		return false, nil
	}

	logger.Info("First job was completed")

	expectedSteps := []expectedStep{
		{name: "clone-config", components: []string{}},
		{name: "config-2-map", components: []string{}},
		{name: "radix-pipeline", components: []string{}},
		{name: "clone", components: []string{}},
		{name: "build-www", components: []string{"www"}},
		{name: "scan-www", components: []string{"www"}},
	}

	if ok := validateJobSteps(env, jobName, expectedSteps); !ok {
		return false, nil
	}

	// Change config branch, trigger webhook and verify job
	if err := patchConfigBranch(env, config.App4NewConfigBranch); err != nil {
		return false, err
	}

	ok = httpUtils.TriggerWebhookPush(env, config.App4NewConfigBranch, config.App4NewCommitID, config.App4SSHRepository, config.App4SharedSecret)
	if !ok {
		return false, nil
	}
	logger.Infof("Second job was triggered")

	ok, jobSummary = test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return job.IsListedWithStatus(env, config.App4Name, "Running")
	})

	jobName = (jobSummary.(*models.JobSummary)).Name
	logger.Infof("Second job name: %s", jobName)

	ok, status = test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return job.IsDone(env, config.App4Name, jobName)
	})

	if !ok {
		return false, nil
	}
	if status.(string) != "Succeeded" {
		return false, nil
	}

	logger.Info("Second job was completed")

	expectedSteps = []expectedStep{
		{name: "clone-config", components: []string{}},
		{name: "config-2-map", components: []string{}},
		{name: "radix-pipeline", components: []string{}},
		{name: "clone", components: []string{}},
		{name: "build-www2", components: []string{"www2"}},
		{name: "scan-www2", components: []string{"www2"}},
	}

	if ok := validateJobSteps(env, jobName, expectedSteps); !ok {
		return false, nil
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

func validateJobSteps(env env.Env, jobName string, expectedSteps []expectedStep) bool {
	steps := job.GetSteps(env, config.App4Name, jobName)

	if steps == nil && len(steps) != len(expectedSteps) {
		logger.Error("Pipeline steps was not as expected")
		return false
	}

	for index, step := range steps {
		if !strings.EqualFold(step.Name, expectedSteps[index].name) {
			logger.Errorf("Expeced step %s, but got %s", expectedSteps[index].name, step.Name)
			return false
		}

		if !array.EqualElements(step.Components, expectedSteps[index].components) {
			logger.Errorf("Expeced components %s, but got %s", expectedSteps[index].components, step.Components)
			return false
		}
	}

	return true
}
