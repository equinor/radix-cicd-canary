package deploy

import (
	"errors"
	"fmt"
	"strings"

	"github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/array"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/job"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

type expectedStep struct {
	name       string
	components []string
}

// Application Tests that we are able to successfully deploy an application by calling Radix API server
func Application(env envUtil.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	appName := config.App3Name
	toEnvironment := config.App3EnvironmentName

	// Trigger deploy via Radix API
	_, err := application.Deploy(env, appName, toEnvironment)
	if err != nil {
		logger.Errorf("Error calling Deploy for application %s:  %v", appName, err)
		return false, err
	}

	// Get job
	ok, jobSummary := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return job.IsListedWithStatus(env, config.App3Name, "Succeeded")
	})
	if !ok {
		return false, errors.New(fmt.Sprintf("Could not get listed job for application %s status \"%s\" - exiting.", config.App3Name, "Succeeded"))
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
		return false, errors.New("pipeline steps was not as expected")
	}

	for index, step := range steps {
		if !strings.EqualFold(step.Name, expectedSteps[index].name) {
			return false, errors.New(fmt.Sprintf("Expected step %s, but got %s", expectedSteps[index].name, step.Name))
		}

		if !array.EqualElements(step.Components, expectedSteps[index].components) {
			return false, errors.New(fmt.Sprintf("Expected components %s, but got %s", expectedSteps[index].components, step.Components))
		}
	}

	return true, nil
}
