package unauthorized

import (
	"fmt"
	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/application"
	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/environment"
	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	log "github.com/sirupsen/logrus"
)

// Access Checks that we are not able to enter any application we should not
// have access to
func Access(cfg config.Config, suiteName string) error {
	logger := log.WithFields(log.Fields{"Suite": suiteName})
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := application.NewGetApplicationParams().
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup).
		WithAppName(defaults.RestrictedApplicationName)

	clientBearerToken := httpUtils.GetClientBearerToken(cfg)
	client := httpUtils.GetApplicationClient(cfg)

	logger.Debugf("check that impersonated user has no access to the application %s", defaults.RestrictedApplicationName)
	_, err := client.GetApplication(params, clientBearerToken)
	return givesAccessError(err)
}

// ReaderAccess Checks that we have appropriate access to the application as readers
func ReaderAccess(cfg config.Config, suiteName string) error {
	logger := log.WithFields(log.Fields{"Suite": suiteName})
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroups := cfg.GetAppReaderGroups()

	getApplicationParams := application.NewGetApplicationParams().
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroups).
		WithAppName(defaults.App2Name)

	clientBearerToken := httpUtils.GetClientBearerToken(cfg)
	applicationClient := httpUtils.GetApplicationClient(cfg)

	logger.Debugf("check that impersonated user has read access to the application %s", defaults.App2Name)
	_, err := applicationClient.GetApplication(getApplicationParams, clientBearerToken)
	if err != nil {
		return err
	}

	restartEnvironmentParameters := environment.NewRestartEnvironmentParams().
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroups).
		WithEnvName("qa").
		WithAppName(defaults.App2Name)

	logger.Debugf("check that impersonated user cannot restart env qa in application %s", defaults.App2Name)
	environmentClient := httpUtils.GetEnvironmentClient(cfg)
	_, err = environmentClient.RestartEnvironment(restartEnvironmentParameters, clientBearerToken)
	wrongAccessError := givesAccessError(err)
	if wrongAccessError != nil {
		return wrongAccessError
	}

	nonExistingUser := "non-existing-user"
	triggerPipelineForApplicationParams := application.NewTriggerPipelineBuildDeployParams().
		WithImpersonateUser(&nonExistingUser).
		WithImpersonateGroup(impersonateGroups).
		WithAppName(defaults.App2Name).
		WithPipelineParametersBuild(
			&models.PipelineParametersBuild{
				Branch:   defaults.App2BranchToBuildFrom,
				CommitID: "this-commit-is-invalid-and-this-job-will-never-be-created",
			},
		)

	logger.Debugf("check that impersonated user cannot trigger pipeline for application %s", defaults.App2Name)
	_, err = applicationClient.TriggerPipelineBuildDeploy(triggerPipelineForApplicationParams, clientBearerToken)
	wrongAccessError = givesAccessError(err)
	return wrongAccessError
}

func givesAccessError(err error) error {
	switch err.(type) {
	case *application.GetApplicationForbidden:
		return nil
	case *environment.RestartEnvironmentForbidden:
		return nil
	case *application.TriggerPipelineBuildDeployForbidden:
		return nil
	case nil:
		return fmt.Errorf("expected 403 from radix-api, but got nil")
	default:
		return fmt.Errorf("expected 403 from radix-api, but got %v", err)
	}
}
