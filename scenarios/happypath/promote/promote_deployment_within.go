package promote

import (
	"fmt"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/job"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

const environmentToPromoteWithin = "qa"

// DeploymentWithinEnvironment Checks that a deployment can be promoted within env
func DeploymentWithinEnvironment(env envUtil.Env, suiteName string) error {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	// Get deployments
	deploymentToPromote, err := getLastDeployment(env, environmentToPromoteWithin)
	if err != nil {
		return err
	}

	// Assert that we no deployments within environment
	deploymentsInEnvironment, err := getDeployments(env, environmentToPromoteWithin)
	if err != nil {
		return err
	}

	numDeploymentsBefore := len(deploymentsInEnvironment)
	promoteJobName, err := promote(env, deploymentToPromote, environmentToPromoteWithin, environmentToPromoteWithin)
	if err != nil {
		return err
	}

	// Get job
	status, err := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (string, error) {
		return job.IsDone(env, config.App2Name, promoteJobName)
	})
	if err != nil {
		return err
	}
	if status != "Succeeded" {
		return fmt.Errorf("expected status Success, but got %s", status)
	}
	_, err = test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, error) {
		return false, isNewDeploymentExist(env, numDeploymentsBefore)
	})
	return err
}

func isNewDeploymentExist(env envUtil.Env, numDeploymentsBefore int) error {
	deploymentsInEnvironment, err := getDeployments(env, environmentToPromoteWithin)
	if err != nil {
		return err
	}

	numDeploymentsAfter := len(deploymentsInEnvironment)
	if (numDeploymentsAfter - numDeploymentsBefore) == 1 {
		return nil
	}

	return fmt.Errorf("new expected deployment does not exist")
}
