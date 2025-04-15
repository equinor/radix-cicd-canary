package component

import (
	"context"
	"fmt"

	environmentclient "github.com/equinor/radix-cicd-canary/generated-client/radixapi/client/environment"
	"github.com/equinor/radix-cicd-canary/generated-client/radixapi/models"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	httpUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/http"
	"github.com/equinor/radix-common/utils/slice"
)

func AllReplicasHaveExpectedStatus(ctx context.Context, cfg config.Config, appName, envName, componentName, expectedStatus string) error {
	env, err := getEnvironment(ctx, cfg, appName, envName)
	if err != nil {
		return fmt.Errorf("failed to get environment %s for application %s: %w", envName, appName, err)
	}

	deployment := env.ActiveDeployment
	if deployment == nil {
		return fmt.Errorf("no active deployment found in environment %s for application %s ", envName, appName)
	}

	component, found := slice.FindFirst(deployment.Components, func(c *models.Component) bool { return *c.Name == componentName })
	if !found {
		return fmt.Errorf("component %s not found in environment %s for application %s", componentName, envName, appName)
	}

	allHasExpectedStatus := slice.All(component.ReplicaList, func(r *models.ReplicaSummary) bool {
		if r == nil || r.ReplicaStatus == nil {
			return false
		}
		return *r.ReplicaStatus.Status == expectedStatus
	})
	if !allHasExpectedStatus {
		return fmt.Errorf("one or more replicas did not have expected status %s for component %s in environment %s for application %s", expectedStatus, componentName, envName, appName)
	}

	return nil
}

func getEnvironment(ctx context.Context, cfg config.Config, appName, envName string) (*models.Environment, error) {
	params := environmentclient.NewGetEnvironmentParams().
		WithContext(ctx).
		WithAppName(appName).
		WithEnvName(envName).
		WithImpersonateUser(cfg.GetImpersonateUser()).
		WithImpersonateGroup(cfg.GetImpersonateGroups())
	client := httpUtils.GetEnvironmentClient(cfg)
	result, err := client.GetEnvironment(params, nil)
	if err != nil {
		return nil, err
	}
	return result.Payload, nil
}
