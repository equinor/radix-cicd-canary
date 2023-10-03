package promote

import (
	"context"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/job"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/pkg/errors"
)

const environmentToPromoteWithin = "qa"

// DeploymentWithinEnvironment Checks that a deployment can be promoted within env
func DeploymentWithinEnvironment(ctx context.Context, cfg config.Config) error {
	appName := defaults.App2Name

	// Get deployments
	deploymentToPromote, err := getLastDeployment(ctx, cfg, appName, environmentToPromoteWithin)
	if err != nil {
		return err
	}

	// Assert that we no deployments within environment
	deploymentsInEnvironment, err := getDeployments(ctx, cfg, appName, environmentToPromoteWithin)
	if err != nil {
		return err
	}

	numDeploymentsBefore := len(deploymentsInEnvironment)
	promoteJobName, err := promote(ctx, cfg, deploymentToPromote, appName, environmentToPromoteWithin, environmentToPromoteWithin)
	if err != nil {
		return err
	}

	// Get job
	jobStatus, err := test.WaitForCheckFuncWithValueOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) (string, error) {
		return job.IsDone(cfg, appName, promoteJobName, ctx)
	})
	if err != nil {
		return err
	}
	if jobStatus != "Succeeded" {
		return errors.Errorf("job %s completed with status %s", promoteJobName, jobStatus)
	}
	return test.WaitForCheckFuncOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) error {
		return isNewDeploymentExist(ctx, cfg, appName, numDeploymentsBefore)
	})
}

func isNewDeploymentExist(ctx context.Context, cfg config.Config, appName string, numDeploymentsBefore int) error {
	deploymentsInEnvironment, err := getDeployments(ctx, cfg, appName, environmentToPromoteWithin)
	if err != nil {
		return err
	}

	numDeploymentsAfter := len(deploymentsInEnvironment)
	if (numDeploymentsAfter - numDeploymentsBefore) != 1 {
		return errors.Errorf("new expected deployment does not exist")
	}
	return nil
}
