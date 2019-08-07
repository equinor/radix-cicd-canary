package promotion

import (
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
)

const environmentToPromoteWithin = "qa"

func promoteDeploymentWithinEnvironment(env env.Env) (bool, error) {
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
	ok, status := test.WaitForCheckFuncWithArguments(env, isJobDone, []string{promoteJobName})
	if ok && status.(string) == "Succeeded" {
		deploymentsInEnvironment, err := getDeployments(env, environmentToPromoteWithin)
		if err != nil {
			return false, err
		}

		numDeploymentsAfter := len(deploymentsInEnvironment)
		return (numDeploymentsAfter - numDeploymentsBefore) == 1, nil
	}

	return false, nil
}
