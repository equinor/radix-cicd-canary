package unauthorized

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/application"
	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/environment"
	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/pipeline_job"
	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/defaults"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/job"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/privateimagehub"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	commonUtils "github.com/equinor/radix-common/utils"
	"github.com/rs/zerolog/log"
)

// Access Checks that we are not able to enter any application we should not
// have access to
func Access(ctx context.Context, cfg config.Config) error {
	impersonateUser := cfg.GetImpersonateUser()
	impersonateGroup := cfg.GetImpersonateGroups()

	params := application.NewGetApplicationParams().
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroup).
		WithContext(ctx).
		WithAppName(defaults.RestrictedApplicationName)

	client := httpUtils.GetApplicationClient(cfg)
	log.Ctx(ctx).Debug().Str("app", defaults.RestrictedApplicationName).Msg("check that impersonated user has no access to the application")
	_, err := client.GetApplication(params, nil)
	return givesAccessError(err)
}

// ReaderAccess Checks that we have appropriate access to the application as readers
func ReaderAccess(ctx context.Context, cfg config.Config) error {
	impersonateUser := cfg.GetImpersonateUser()
	readerGroup := cfg.GetAppReaderGroup()
	appName := defaults.App2Name

	type impersonateParam interface {
		SetImpersonateUser(*string)
		SetImpersonateGroup([]string)
		SetAppName(string)
	}

	type scenarioSpec struct {
		name          string
		logMsg        string
		expectedError error
		testFunc      func(ctx context.Context, impersonationSetter func(impersonateParam)) error
	}

	scenarios := []scenarioSpec{
		{
			name:          "reader-user-can-read-RR",
			logMsg:        fmt.Sprintf("checking that user with reader role can read RR for application %s", appName),
			expectedError: nil,
			testFunc: func(ctx context.Context, impersonationSetter func(impersonateParam)) error {
				param := application.NewGetApplicationParams().WithContext(ctx)
				impersonationSetter(param)
				_, err := httpUtils.GetApplicationClient(cfg).GetApplication(param, nil)
				return err
			},
		},
		{
			name:          "reader-user-cannot-restart-env",
			logMsg:        fmt.Sprintf("checking that user with reader role cannot restart env %s for application %s", defaults.App2EnvironmentName, appName),
			expectedError: environment.NewRestartEnvironmentForbidden(),
			testFunc: func(ctx context.Context, impersonationSetter func(impersonateParam)) error {
				param := environment.NewRestartEnvironmentParams().
					WithEnvName(defaults.App2EnvironmentName).
					WithContext(ctx)
				impersonationSetter(param)
				_, err := httpUtils.GetEnvironmentClient(cfg).RestartEnvironment(param, nil)
				return err
			},
		},
		{
			name:          "reader-user-cannot-trigger-pipeline",
			logMsg:        fmt.Sprintf("checking that user with read role cannot trigger build-deploy pipeline for application %s", appName),
			expectedError: application.NewTriggerPipelineBuildDeployForbidden(),
			testFunc: func(ctx context.Context, impersonationSetter func(impersonateParam)) error {
				param := application.NewTriggerPipelineBuildDeployParams().
					WithContext(ctx).
					WithPipelineParametersBuild(
						&models.PipelineParametersBuild{
							Branch:   defaults.App2BranchToBuildFrom,
							CommitID: "this-commit-is-invalid-and-this-job-will-never-be-created",
						},
					)
				impersonationSetter(param)
				_, err := httpUtils.GetApplicationClient(cfg).TriggerPipelineBuildDeploy(param, nil)
				return err
			},
		},
		{
			name:          "reader-user-cannot-delete-app",
			logMsg:        fmt.Sprintf("checking that user with read role cannot delete app %s", appName),
			expectedError: application.NewDeleteApplicationForbidden(),
			testFunc: func(ctx context.Context, impersonationSetter func(impersonateParam)) error {
				param := application.NewDeleteApplicationParams().WithContext(ctx)
				impersonationSetter(param)
				_, err := httpUtils.GetApplicationClient(cfg).DeleteApplication(param, nil)
				return err
			},
		},
		{
			name:          "reader-user-cannot-set-build-secret",
			logMsg:        fmt.Sprintf("checking that user with read role cannot set build secret for app %s", appName),
			expectedError: application.NewUpdateBuildSecretsSecretValueForbidden(),
			testFunc: func(ctx context.Context, impersonationSetter func(impersonateParam)) error {
				param := application.NewUpdateBuildSecretsSecretValueParams().
					WithContext(ctx).
					WithSecretName(defaults.App2SecretName).
					WithSecretValue(&models.SecretParameters{SecretValue: commonUtils.StringPtr(defaults.App2SecretValue)})
				impersonationSetter(param)
				_, err := httpUtils.GetApplicationClient(cfg).UpdateBuildSecretsSecretValue(param, nil)
				return err
			},
		},
		{
			name:          "reader-user-cannot-set-private-image-hub-secret",
			logMsg:        fmt.Sprintf("checking that user with read role cannot set private image hub secret for app %s", appName),
			expectedError: application.NewUpdatePrivateImageHubsSecretValueBadRequest(), // TODO: should be forbidden 403. requires some rewrite of radix-api or radix-operator lib function
			testFunc: func(ctx context.Context, impersonationSetter func(impersonateParam)) error {
				imageHubs, err := privateimagehub.List(cfg, appName)
				if err != nil {
					return err
				}
				imageHub := imageHubs[0]
				param := application.NewUpdatePrivateImageHubsSecretValueParams().
					WithContext(ctx).
					WithServerName(*imageHub.Server).
					WithImageHubSecret(&models.SecretParameters{SecretValue: commonUtils.StringPtr("some-value")})
				impersonationSetter(param)
				_, err = httpUtils.GetApplicationClient(cfg).UpdatePrivateImageHubsSecretValue(param, nil)
				return err
			},
		},
		{
			name:          "reader-user-cannot-set-secret",
			logMsg:        fmt.Sprintf("checking that user with read role cannot set secret for app %s", appName),
			expectedError: environment.NewChangeComponentSecretForbidden(),
			testFunc: func(ctx context.Context, impersonationSetter func(impersonateParam)) error {
				param := environment.NewChangeComponentSecretParams().
					WithEnvName(defaults.App2EnvironmentName).
					WithContext(ctx).
					WithComponentName(defaults.App2Component2Name).
					WithSecretName(defaults.App2SecretName).
					WithComponentSecret(
						&models.SecretParameters{
							SecretValue: commonUtils.StringPtr(defaults.App2SecretValue),
						})
				impersonationSetter(param)
				_, err := httpUtils.GetEnvironmentClient(cfg).ChangeComponentSecret(param, nil)
				return err
			},
		},
		// TODO: check that reading pipeline log is allowed https://console.dev.radix.equinor.com/api/v1/applications/radix-networkpolicy-canary/jobs/radix-pipeline-20230728084604-ew3sz/logs/radix-pipeline?lines=1000
		{
			name:          "reader-user-can-read-pipeline-log",
			logMsg:        fmt.Sprintf("checking that user with read role can read pipeline log for app %s", appName),
			expectedError: nil,
			testFunc: func(ctx context.Context, impersonationSetter func(impersonateParam)) error {
				// Get job
				jobSummary, err := test.WaitForCheckFuncWithValueOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) (*models.JobSummary, error) {
					return job.GetAnyPipelineJobWithStatus(ctx, cfg, appName, "Succeeded")
				})
				if err != nil {
					return err
				}
				jobName := jobSummary.Name
				param := pipeline_job.NewGetPipelineJobStepLogsParams().
					WithContext(ctx).
					WithJobName(jobName).
					WithStepName("radix-pipeline")
				impersonationSetter(param)
				_, err = httpUtils.GetJobClient(cfg).GetPipelineJobStepLogs(param, nil)
				return err
			},
		},
		// TODO: check that reading runtime log is allowed https://console.dev.radix.equinor.com/api/v1/applications/radix-networkpolicy-canary/environments/egressrulestopublicdns/components/web/replicas/web-978d76dc4-ctgzz/logs?lines=1000
	}

	setImpersonation := func(p impersonateParam) {
		p.SetImpersonateUser(impersonateUser)
		p.SetImpersonateGroup([]string{readerGroup})
		p.SetAppName(appName)
	}

	for _, scenario := range scenarios {
		scenarioCtx := log.Ctx(ctx).With().Str("scenario", scenario.name).Logger().WithContext(ctx)
		log.Ctx(scenarioCtx).Debug().Msg(scenario.logMsg)
		err := scenario.testFunc(scenarioCtx, setImpersonation)
		if !errors.Is(err, scenario.expectedError) {
			return errors.Errorf("incorrect response on scenario %s: Got %v, expected %v", scenario.name, err, scenario.expectedError)
		}
	}
	return nil
}

func givesAccessError(err error) error {
	switch err.(type) {
	case *application.GetApplicationForbidden:
		return nil
	}
	return err
}
