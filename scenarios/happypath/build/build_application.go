package build

import (
	"time"

	jobclient "github.com/equinor/radix-cicd-canary/generated-client/client/job"
	models "github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

// Application Tests that we are able to successfully build an application
func Application(env env.Env) (bool, error) {
	// Trigger build via web hook
	ok := httpUtils.TriggerWebhookPush(env, config.App2BranchToBuildFrom, config.App2CommitID, config.App2SSHRepository, config.App2SharedSecret)
	if !ok {
		return false, nil
	}
	log.Infof("First job was triggered")

	// Get job
	ok, jobSummary := test.WaitForCheckFunc(env, isJobListed)
	if !ok {
		return false, nil
	}

	jobName := (jobSummary.(*models.JobSummary)).Name
	log.Infof("First job name: %s", jobName)

	// Another build should cause second job to queue up
	// Trigger another build via web hook
	time.Sleep(1 * time.Second)
	ok = httpUtils.TriggerWebhookPush(env, config.App2BranchToBuildFrom, config.App2CommitID, config.App2SSHRepository, config.App2SharedSecret)
	if !ok {
		return false, nil
	}
	log.Infof("Second job was triggered")

	ok, jobSummary = test.WaitForCheckFuncWithArguments(env, isSecondJobExpectedStatus, []string{"Queued"})
	if !ok {
		return false, nil
	}

	log.Info("Second job was queued")
	ok, status := test.WaitForCheckFuncWithArguments(env, isJobDone, []string{jobName})

	if ok && status.(string) == "Succeeded" {
		log.Info("First job was completed")
		ok, jobSummary = test.WaitForCheckFuncWithArguments(env, isSecondJobExpectedStatus, []string{"Running"})
		if !ok {
			return false, nil
		}

		// Stop job and verify that it has been stopped
		jobName = (jobSummary.(*models.JobSummary)).Name
		log.Infof("Second job name: %s", jobName)
		ok = stopJob(env, config.App2Name, jobName)
		if !ok {
			return false, nil
		}

		ok, jobSummary = test.WaitForCheckFuncWithArguments(env, isSecondJobExpectedStatus, []string{"Stopped"})
		if !ok {
			return false, nil
		}

		log.Info("Second job was stopped")
		return true, nil
	}

	return false, nil
}

func stopJob(env env.Env, appName, jobName string) bool {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetJobClient(env)

	params := jobclient.NewStopApplicationJobParams().
		WithAppName(appName).
		WithJobName(jobName).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)

	jobStopped, err := client.StopApplicationJob(params, clientBearerToken)
	if err == nil && jobStopped != nil {
		return true
	}

	log.Infof("Failed stopping job %s. Error: %v", jobName, err)
	return false
}

func isJobListed(env env.Env, args []string) (bool, interface{}) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := jobclient.NewGetApplicationJobsParams().
		WithAppName(config.App2Name).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)
	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetJobClient(env)

	applicationJobs, err := client.GetApplicationJobs(params, clientBearerToken)
	if err == nil && applicationJobs.Payload != nil && len(applicationJobs.Payload) > 0 {
		return true, applicationJobs.Payload[0]
	}

	log.Info("Job was not listed yet")
	return false, nil
}

func isSecondJobExpectedStatus(env env.Env, args []string) (bool, interface{}) {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	expectedStatus := args[0]
	params := jobclient.NewGetApplicationJobsParams().
		WithAppName(config.App2Name).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)
	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetJobClient(env)

	applicationJobs, err := client.GetApplicationJobs(params, clientBearerToken)
	if err == nil && applicationJobs.Payload != nil && len(applicationJobs.Payload) > 0 && applicationJobs.Payload[0].Status == expectedStatus {
		log.Infof("Second job was listed with status: %s", expectedStatus)
		return true, applicationJobs.Payload[0]
	}

	log.Infof("Second job was not listed yet. Expected status: %s", expectedStatus)
	return false, nil
}

func isJobDone(env env.Env, args []string) (bool, interface{}) {
	jobStatus := getJobStatus(env, args[0])
	if jobStatus == "Succeeded" || jobStatus == "Failed" {
		log.Infof("Job is done with status: %s", jobStatus)
		return true, jobStatus
	}

	log.Info("Job is not done yet")
	return false, nil
}

func getJobStatus(env env.Env, jobName string) string {
	impersonateUser := env.GetImpersonateUser()
	impersonateGroup := env.GetImpersonateGroup()

	params := jobclient.NewGetApplicationJobParams().
		WithAppName(config.App2Name).
		WithJobName(jobName).
		WithImpersonateUser(&impersonateUser).
		WithImpersonateGroup(&impersonateGroup)

	clientBearerToken := httpUtils.GetClientBearerToken(env)
	client := httpUtils.GetJobClient(env)

	applicationJob, err := client.GetApplicationJob(params, clientBearerToken)
	if err == nil && applicationJob.Payload != nil {
		return applicationJob.Payload.Status
	}

	log.Info("Job was not listed yet")
	return ""
}
