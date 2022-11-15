package promote

import (
	"errors"
	"fmt"

	applicationclient "github.com/equinor/radix-cicd-canary/generated-client/client/application"
	environmentclient "github.com/equinor/radix-cicd-canary/generated-client/client/environment"
	models "github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/job"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

const (
	pipelineName    = "promote"
	envToDeployFrom = "qa"
	envToDeployTo   = "dev"
)

var logger *log.Entry

// DeploymentToAnotherEnvironment Checks that deployment can be promoted to other environment
func DeploymentToAnotherEnvironment(env envUtil.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	// Get deployments
	deploymentToPromote, err := getLastDeployment(env, envToDeployFrom)
	if err != nil {
		return false, err
	}

	// Assert that we no deployments within environment
	deploymentsInEnvironment, err := getDeployments(env, envToDeployTo)
	if err != nil {
		return false, err
	}

	numDeploymentsBefore := len(deploymentsInEnvironment)
	promoteJobName, err := promote(env, deploymentToPromote, envToDeployFrom, envToDeployTo)
	if err != nil {
		return false, err
	}

	// Get job
	ok, status := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return job.IsDone(env, config.App2Name, promoteJobName)
	})
	if ok && status.(string) == "Succeeded" {
		deploymentsInEnvironment, err := getDeployments(env, envToDeployTo)
		if err != nil {
			return false, err
		}

		numDeploymentsAfter := len(deploymentsInEnvironment)
		return (numDeploymentsAfter - numDeploymentsBefore) == 1, nil
	}

	return false, errors.New(fmt.Sprintf("expected status Success, but got %s", status))
}

func getLastDeployment(env envUtil.Env, environment string) (*models.DeploymentSummary, error) {
	deployments, err := getDeployments(env, environment)
	if err != nil || len(deployments) == 0 {
		return nil, err
	}

	// Which deployment is irrelevant
	return deployments[0], nil
}

func getDeployments(env envUtil.Env, environment string) ([]*models.DeploymentSummary, error) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := environmentclient.NewGetApplicationEnvironmentDeploymentsParams().
		WithAppName(config.App2Name).
		WithEnvName(environment).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)
	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetEnvironmentClient(env)

	deployments, err := client.GetApplicationEnvironmentDeployments(params, clientBearerToken)
	if err != nil {
		return nil, err
	}

	return deployments.Payload, nil
}

func promote(env envUtil.Env, deployment *models.DeploymentSummary, from, to string) (string, error) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	bodyParameters := models.PipelineParametersPromote{
		DeploymentName:  *deployment.Name,
		FromEnvironment: from,
		ToEnvironment:   to,
	}

	params := applicationclient.NewTriggerPipelinePromoteParams().
		WithAppName(config.App2Name).
		WithPipelineParametersPromote(&bodyParameters).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)
	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetApplicationClient(env)

	returnValue, err := client.TriggerPipelinePromote(params, clientBearerToken)
	if err != nil {
		return "", err
	}

	return returnValue.Payload.Name, nil
}
