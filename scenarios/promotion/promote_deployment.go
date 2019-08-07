package promotion

import (
	applicationclient "github.com/equinor/radix-cicd-canary/generated-client/client/application"
	environmentclient "github.com/equinor/radix-cicd-canary/generated-client/client/environment"
	jobclient "github.com/equinor/radix-cicd-canary/generated-client/client/job"
	models "github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

const (
	pipelineName    = "promote"
	app2Name        = "canarycicd-test2"
	envToDeployFrom = "qa"
	envToDeployTo   = "dev"
)

func promoteDeployment(env env.Env) (bool, error) {
	test.WaitForCheckFunc(env, isApplicationDefined)

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
	ok, status := test.WaitForCheckFuncWithArguments(env, isJobDone, []string{promoteJobName})
	if ok && status.(string) == "Succeeded" {
		deploymentsInEnvironment, err := getDeployments(env, envToDeployTo)
		if err != nil {
			return false, err
		}

		numDeploymentsAfter := len(deploymentsInEnvironment)
		return (numDeploymentsAfter - numDeploymentsBefore) == 1, nil
	}

	return false, nil
}

func isApplicationDefined(env env.Env, args []string) (bool, interface{}) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := applicationclient.NewGetApplicationParams().
		WithAppName(app2Name).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)
	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetApplicationClient(env)

	_, err := client.GetApplication(params, clientBearerToken)
	if err == nil {
		return true, nil
	}

	log.Info("Application is not defined")
	return false, nil
}

func getLastDeployment(env env.Env, environment string) (*models.DeploymentSummary, error) {
	deployments, err := getDeployments(env, environment)
	if err != nil || len(deployments) == 0 {
		return nil, err
	}

	// Which deployment is irrelevant
	return deployments[0], nil
}

func getDeployments(env env.Env, environment string) ([]*models.DeploymentSummary, error) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := environmentclient.NewGetApplicationEnvironmentDeploymentsParams().
		WithAppName(app2Name).
		WithEnvName(environment).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)
	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetEnvironmentClient(env)

	deployments, err := client.GetApplicationEnvironmentDeployments(params, clientBearerToken)
	if err == nil {
		return nil, err
	}

	return deployments.Payload, nil
}

func promote(env env.Env, deployment *models.DeploymentSummary, from, to string) (string, error) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	bodyParameters := models.PipelineParameters{
		PipelineParametersPromote: models.PipelineParametersPromote{
			DeploymentName:  deployment.Name,
			FromEnvironment: envToDeployFrom,
			ToEnvironment:   envToDeployTo,
		},
	}

	params := applicationclient.NewTriggerPipelineParams().
		WithAppName(app2Name).
		WithPipelineName(pipelineName).
		WithPipelineParameters(&bodyParameters).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)
	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetApplicationClient(env)

	returnValue, err := client.TriggerPipeline(params, clientBearerToken)
	if err != nil {
		return "", err
	}

	return returnValue.Payload.Name, nil
}

func isJobDone(env env.Env, args []string) (bool, interface{}) {
	jobStatus := getJobStatus(env, args[0])
	if jobStatus == "Succeeded" || jobStatus == "Failed" {
		log.Info("Job is done")
		return true, jobStatus
	}

	log.Info("Job is not done yet")
	return false, nil
}

func getJobStatus(env env.Env, jobName string) string {
	applicationJob, err := getJob(env, jobName)
	if err == nil && applicationJob != nil {
		return applicationJob.Status
	}

	log.Info("Job was not listed yet")
	return ""
}

func getJob(env env.Env, jobName string) (*models.Job, error) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := jobclient.NewGetApplicationJobParams().
		WithAppName(app2Name).
		WithJobName(jobName).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetJobClient(env)

	applicationJob, err := client.GetApplicationJob(params, clientBearerToken)
	if err != nil {
		return nil, err
	}

	return applicationJob.Payload, nil
}
