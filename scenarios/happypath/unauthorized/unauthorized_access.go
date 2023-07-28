package unauthorized

import (
	"errors"
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
	readerGroup := cfg.GetAppReaderGroup()
	clientBearerToken := httpUtils.GetClientBearerToken(cfg)

	type impersonateParam interface {
		SetImpersonateUser(*string)
		SetImpersonateGroup([]string)
		SetAppName(string)
	}

	type scenarioSpec struct {
		name          string
		logMsg        string
		expectedError error
		testFunc      func(impersonationSetter func(impersonateParam)) error
	}

	scenarios := []scenarioSpec{
		{
			name:          "reader-user-can-read-RR",
			logMsg:        fmt.Sprintf("checking that user with reader role can read RR for application %s", defaults.App2Name),
			expectedError: nil,
			testFunc: func(impersonationSetter func(impersonateParam)) error {
				param := application.NewGetApplicationParams()
				impersonationSetter(param)
				_, err := httpUtils.GetApplicationClient(cfg).GetApplication(param, clientBearerToken)
				return err
			},
		},
		{
			name:          "reader-user-cannot-restart-env",
			logMsg:        fmt.Sprintf("checking that user with reader role cannot restart env %s for application %s", defaults.App2EnvironmentName, defaults.App2Name),
			expectedError: environment.NewRestartEnvironmentForbidden(),
			testFunc: func(impersonationSetter func(impersonateParam)) error {
				param := environment.NewRestartEnvironmentParams().WithEnvName(defaults.App2EnvironmentName)
				impersonationSetter(param)
				_, err := httpUtils.GetEnvironmentClient(cfg).RestartEnvironment(param, clientBearerToken)
				return err
			},
		},
		{
			name:          "reader-user-cannot-trigger-pipeline",
			logMsg:        fmt.Sprintf("checking that user with read role cannot trigger build-deploy pipeline for application %s", defaults.App2Name),
			expectedError: application.NewTriggerPipelineBuildDeployForbidden(),
			testFunc: func(impersonationSetter func(impersonateParam)) error {
				param := application.NewTriggerPipelineBuildDeployParams().
					WithPipelineParametersBuild(
						&models.PipelineParametersBuild{
							Branch:   defaults.App2BranchToBuildFrom,
							CommitID: "this-commit-is-invalid-and-this-job-will-never-be-created",
						},
					)
				impersonationSetter(param)
				_, err := httpUtils.GetApplicationClient(cfg).TriggerPipelineBuildDeploy(param, clientBearerToken)
				return err
			},
		},
	}

	setImpersonation := func(p impersonateParam) {
		p.SetImpersonateUser(impersonateUser)
		p.SetImpersonateGroup([]string{readerGroup})
		p.SetAppName(defaults.App2Name)
	}

	for _, scenario := range scenarios {
		logger.Debugf(scenario.logMsg)
		err := scenario.testFunc(setImpersonation)
		if !errors.Is(err, scenario.expectedError) {
			return fmt.Errorf("incorrect response on scenario %s: Got %v, expected %v", scenario.name, err, scenario.expectedError)
		}
	}
	return nil
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
