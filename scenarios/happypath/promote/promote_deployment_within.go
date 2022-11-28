package promote

import (
	"fmt"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/job"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

const environmentToPromoteWithin = "qa"

// DeploymentWithinEnvironment Checks that a deployment can be promoted within env
func DeploymentWithinEnvironment(cfg config.Config, suiteName string) error {
	logger := log.WithFields(log.Fields{"Suite": suiteName})

	// Get deployments
	deploymentToPromote, err := getLastDeployment(cfg, environmentToPromoteWithin)
	if err != nil {
		return err
	}

	// Assert that we no deployments within environment
	deploymentsInEnvironment, err := getDeployments(cfg, environmentToPromoteWithin)
	if err != nil {
		return err
	}

	numDeploymentsBefore := len(deploymentsInEnvironment)
	promoteJobName, err := promote(cfg, deploymentToPromote, environmentToPromoteWithin, environmentToPromoteWithin)
	if err != nil {
		return err
	}

	// Get job
	jobStatus, err := test.WaitForCheckFuncWithValueOrTimeout(cfg, func(cfg config.Config) (string, error) {
		return job.IsDone(cfg, defaults.App2Name, promoteJobName, logger)
	}, logger)
	if err != nil {
		return err
	}
	if jobStatus != "Succeeded" {
		return fmt.Errorf("job %s completed with status %s", promoteJobName, jobStatus)
	}
	return test.WaitForCheckFuncOrTimeout(cfg, func(cfg config.Config) error {
		return isNewDeploymentExist(cfg, numDeploymentsBefore)
	}, logger)
}

func isNewDeploymentExist(cfg config.Config, numDeploymentsBefore int) error {
	deploymentsInEnvironment, err := getDeployments(cfg, environmentToPromoteWithin)
	if err != nil {
		return err
	}

	numDeploymentsAfter := len(deploymentsInEnvironment)
	if (numDeploymentsAfter - numDeploymentsBefore) != 1 {
		return fmt.Errorf("new expected deployment does not exist")
	}
	return nil
}
