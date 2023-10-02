package promote

import (
	"context"
	"fmt"

	applicationclient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/application"
	environmentclient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/environment"
	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/job"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	"github.com/rs/zerolog/log"
)

const (
	envToDeployFrom = "qa"
	envToDeployTo   = "dev"
)

// DeploymentToAnotherEnvironment Checks that deployment can be promoted to other environment
func DeploymentToAnotherEnvironment(ctx context.Context, cfg config.Config, suiteName string) error {
	appName := defaults.App2Name
	appCtx := log.Ctx(ctx).With().Str("app", appName).Logger().WithContext(ctx)

	// Get deployments
	deploymentToPromote, err := getLastDeployment(ctx, cfg, appName, envToDeployFrom)
	if err != nil {
		return err
	}

	// Assert that we have no deployments within environment
	deploymentsInEnvironment, err := getDeployments(ctx, cfg, appName, envToDeployTo)
	if err != nil {
		return err
	}
	log.Ctx(appCtx).Debug().Msg("no deployments within environment")

	numDeploymentsBefore := len(deploymentsInEnvironment)
	promoteJobName, err := promote(ctx, cfg, deploymentToPromote, appName, envToDeployFrom, envToDeployTo)
	if err != nil {
		return err
	}

	// Get job
	jobStatus, err := test.WaitForCheckFuncWithValueOrTimeout(cfg, func(cfg config.Config, ctx context.Context) (string, error) {
		return job.IsDone(cfg, appName, promoteJobName, ctx)
	}, appCtx)
	if err != nil {
		return err
	}
	if jobStatus != "Succeeded" {
		return fmt.Errorf("job %s completed with status %s", promoteJobName, jobStatus)
	}
	deploymentsInEnvironment, err = getDeployments(ctx, cfg, appName, envToDeployTo)
	if err != nil {
		return err
	}

	numDeploymentsAfter := len(deploymentsInEnvironment)
	newDeploymentCount := numDeploymentsAfter - numDeploymentsBefore
	if newDeploymentCount != 1 {
		return fmt.Errorf("new expected deployment does not exist")
	}
	return nil
}

func getLastDeployment(ctx context.Context, cfg config.Config, appName string, environment string) (*models.DeploymentSummary, error) {
	deployments, err := getDeployments(ctx, cfg, appName, environment)
	if err != nil || len(deployments) == 0 {
		return nil, err
	}

	// Which deployment is irrelevant
	return deployments[0], nil
}

func getDeployments(ctx context.Context, cfg config.Config, appName string, environment string) ([]*models.DeploymentSummary, error) {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := environmentclient.NewGetApplicationEnvironmentDeploymentsParams().
		WithAppName(appName).
		WithEnvName(environment).
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup)
	client := httpUtils.GetEnvironmentClient(cfg)
	deployments, err := client.GetApplicationEnvironmentDeployments(params, nil)
	if err != nil {
		return nil, err
	}

	return deployments.Payload, nil
}

func promote(ctx context.Context, cfg config.Config, deployment *models.DeploymentSummary, appName string, from, to string) (string, error) {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	bodyParameters := models.PipelineParametersPromote{
		DeploymentName:  *deployment.Name,
		FromEnvironment: from,
		ToEnvironment:   to,
	}

	params := applicationclient.NewTriggerPipelinePromoteParams().
		WithAppName(appName).
		WithContext(ctx).
		WithPipelineParametersPromote(&bodyParameters).
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup)
	client := httpUtils.GetApplicationClient(cfg)
	returnValue, err := client.TriggerPipelinePromote(params, nil)
	if err != nil {
		return "", err
	}

	return returnValue.Payload.Name, nil
}
