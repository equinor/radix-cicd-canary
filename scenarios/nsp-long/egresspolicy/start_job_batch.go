package egresspolicy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/metrics"
	nspMetrics "github.com/equinor/radix-cicd-canary/metrics/scenarios/nsp"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/k8sjob"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

// StartAndCheckJobBatch starts a job batch and confirms that jobs were created
func StartAndCheckJobBatch(cfg config.Config, suiteName string) error {
	appEnvs := []string{"egressrulestopublicdns", "allowradix"}
	jobComponentName := cfg.GetNetworkPolicyCanaryJobComponentName()
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	var batchNames []string
	for _, appEnv := range appEnvs {
		baseUrl := cfg.GetNetworkPolicyCanaryUrl(appEnv)
		password := cfg.GetNetworkPolicyCanaryPassword()
		batchName, err := startJobBatch(baseUrl, password, appEnv)
		if err != nil {
			return err
		} else {
			batchNames = append(batchNames, batchName)
		}
	}

	for i := 0; i < len(appEnvs); i++ {
		appEnv := appEnvs[i]
		batchName := batchNames[i]
		err := checkJobBatch(cfg, cfg.GetNetworkPolicyCanaryAppName(), appEnv, jobComponentName, batchName)
		if err != nil {
			return err
		}
	}

	return nil
}

func checkJobBatch(cfg config.Config, appName, appEnv string, jobComponentName string, batchName string) error {
	return test.WaitForCheckFuncOrTimeout(cfg, func(cfg config.Config) error {
		return k8sjob.IsListedWithStatus(cfg, appName, appEnv, jobComponentName, batchName, "Succeeded")
	}, logger)
}

func startJobBatch(baseUrl string, password string, appEnv string) (string, error) {
	jobBatchUrl := fmt.Sprintf("%s/startjobbatch", baseUrl)
	httpClient := &http.Client{}
	req, _ := http.NewRequest("GET", jobBatchUrl, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", password))
	response, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	if response.StatusCode != 200 {
		return "", fmt.Errorf("got non-200 OK from %s", jobBatchUrl)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	var batchStatus models.BatchStatus
	unMarshalErr := json.Unmarshal(body, &batchStatus)
	if unMarshalErr != nil {
		return "", unMarshalErr
	}
	if batchStatus.BatchName == "" {
		err = fmt.Errorf("no batchName attribute in job batch creation response. appEnv %s", appEnv)
		return "", err
	}
	return batchStatus.BatchName, nil
}

// StartAndCheckJobBatchSuccess is a function after a call to Lookup succeeds
func StartAndCheckJobBatchSuccess(testName string) {
	nspMetrics.AddStartAndCheckJobBatchSuccess()
	metrics.AddTestOne(testName, nspMetrics.Success)
	metrics.AddTestZero(testName, nspMetrics.Errors)
	logger.Infof("Test %s: SUCCESS", testName)
}

// StartAndCheckJobBatchFail is a function after a call to Lookup failed
func StartAndCheckJobBatchFail(testName string) {
	nspMetrics.AddStartAndCheckJobBatchFail()
	metrics.AddTestZero(testName, nspMetrics.Success)
	metrics.AddTestOne(testName, nspMetrics.Errors)
	logger.Infof("Test %s: FAIL", testName)
}
