package promote

import (
	"errors"
	"fmt"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/job"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

const environmentToPromoteWithin = "qa"

// DeploymentWithinEnvironment Checks that a deployment can be promoted within env
func DeploymentWithinEnvironment(env envUtil.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	// Get deployments
	deploymentToPromote, err := getLastDeployment(env, environmentToPromoteWithin)
	if err != nil {
		return false, err
	}

	// Assert that we no deployments within environment
	deploymentsInEnvironment, err := getDeployments(env, environmentToPromoteWithin)
	if err != nil {
		return false, err
	}

	numDeploymentsBefore := len(deploymentsInEnvironment)
	promoteJobName, err := promote(env, deploymentToPromote, environmentToPromoteWithin, environmentToPromoteWithin)
	if err != nil {
		return false, err
	}

	// Get job
	ok, status := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return job.IsDone(env, config.App2Name, promoteJobName)
	})
	if ok && status.(string) == "Succeeded" {
		doneCheck, ok := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
			return isNewDeploymentExist(env, numDeploymentsBefore)
		})

		if doneCheck && ok.(bool) {
			return true, nil
		}
	}

	return false, errors.New(fmt.Sprintf("expected status Success, but got %s", status))
}

func isNewDeploymentExist(env envUtil.Env, numDeploymentsBefore int) (bool, interface{}) {
	deploymentsInEnvironment, err := getDeployments(env, environmentToPromoteWithin)
	if err != nil {
		logger.Errorf("Error: %v", err)
		return true, false
	}

	if err != nil {
		logger.Errorf("Error: %v", err)
		return true, false
	}
	numDeploymentsAfter := len(deploymentsInEnvironment)
	if (numDeploymentsAfter - numDeploymentsBefore) == 1 {
		return true, true
	}

	return false, nil
}
