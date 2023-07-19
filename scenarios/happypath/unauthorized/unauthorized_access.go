package unauthorized

import (
	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/application"
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

	logger.Debugf("check that inpersonated user has no access to the application %s", defaults.RestrictedApplicationName)
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
	client := httpUtils.GetApplicationClient(cfg)

	logger.Debugf("check that impersonated user has read access to the application %s", defaults.App2Name)
	_, err := client.GetApplication(getApplicationParams, clientBearerToken)
	if err != nil {
		return err
	}

	restartApplicationParams := application.NewRestartApplicationParams().
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroups).
		WithAppName(defaults.App2Name)

	logger.Debugf("check that impersonated user cannot restart application %s", defaults.App2Name)
	_, err = client.RestartApplication(restartApplicationParams, clientBearerToken)
	wrongAccessError := givesAccessError(err)
	if wrongAccessError != nil {
		return wrongAccessError
	}

	triggerPipelineForApplicationParams := application.NewTriggerPipelineBuildDeployParams().
		WithImpersonateUser(impersonateUser).
		WithImpersonateGroup(impersonateGroups).
		WithAppName(defaults.App2Name)

	logger.Debugf("check that impersonated user cannot trigger pipeline for application %s", defaults.App2Name)
	_, err = client.TriggerPipelineBuildDeploy(triggerPipelineForApplicationParams, clientBearerToken)
	if err != nil {
		return err
	}
	return givesAccessError(err)
}

func givesAccessError(err error) error {
	switch err.(type) {
	case *application.GetApplicationForbidden:
		return nil
	}
	return err
}
