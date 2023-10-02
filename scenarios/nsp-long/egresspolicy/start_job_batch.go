package egresspolicy

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	jobServerModels "github.com/equinor/radix-cicd-canary/generated-client/jobserver/models"
	"github.com/equinor/radix-cicd-canary/metrics"
	nspMetrics "github.com/equinor/radix-cicd-canary/metrics/scenarios/nsp"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/rs/zerolog/log"
)

// StartAndCheckJobBatch starts a job batch and confirms that jobs were created
func StartAndCheckJobBatch(ctx context.Context, cfg config.Config, suiteName string) error {
	appEnvs := []string{"egressrulestopublicdns", "allowradix"}

	for _, appEnv := range appEnvs {
		baseUrl := cfg.GetNetworkPolicyCanaryUrl(appEnv)
		password := cfg.GetNetworkPolicyCanaryPassword()
		if err := startJobBatch(baseUrl, password, appEnv); err != nil {
			return err
		}
	}

	return nil
}

func startJobBatch(baseUrl string, password string, appEnv string) error {
	jobBatchUrl := fmt.Sprintf("%s/startjobbatch", baseUrl)
	httpClient := &http.Client{}
	req, _ := http.NewRequest("GET", jobBatchUrl, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", password))
	response, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return fmt.Errorf("got non-200 OK from %s", jobBatchUrl)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	var batchStatus jobServerModels.BatchStatus
	unMarshalErr := json.Unmarshal(body, &batchStatus)
	if unMarshalErr != nil {
		return unMarshalErr
	}
	if batchStatus.Name == nil || *batchStatus.Name == "" {
		err = fmt.Errorf("no name attribute in job batch creation response. appEnv %s", appEnv)
		return err
	}
	return nil
}

// StartAndCheckJobBatchSuccess is a function after a call to Lookup succeeds
func StartAndCheckJobBatchSuccess(ctx context.Context, testName string) {
	nspMetrics.AddStartAndCheckJobBatchSuccess()
	metrics.AddTestOne(testName, nspMetrics.Success)
	metrics.AddTestZero(testName, nspMetrics.Errors)
	log.Ctx(ctx).Info().Msg("Test: SUCCESS")
}

// StartAndCheckJobBatchFail is a function after a call to Lookup failed
func StartAndCheckJobBatchFail(ctx context.Context, testName string) {
	nspMetrics.AddStartAndCheckJobBatchFail()
	metrics.AddTestZero(testName, nspMetrics.Success)
	metrics.AddTestOne(testName, nspMetrics.Errors)
	log.Ctx(ctx).Info().Msg("Test: FAIL")
}
