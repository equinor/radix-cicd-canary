package egresspolicy

import (
	"encoding/json"
	"fmt"
	"github.com/equinor/radix-cicd-canary/generated-client/models"
	"github.com/equinor/radix-cicd-canary/metrics"
	nspMetrics "github.com/equinor/radix-cicd-canary/metrics/scenarios/nsp"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/k8sjob"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

var logger *log.Entry

// StartAndCheckJobBatch starts a job batch and confirms that jobs were created
func StartAndCheckJobBatch(env envUtil.Env, suiteName string) (bool, error) {
	appEnvs := []string{"egressrulestopublicdns", "allowradix"}
	jobComponentName := env.GetNetworkPolicyCanaryJobComponentName()
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	var batchNames []string
	for _, appEnv := range appEnvs {
		err, batchName := startJobBatch(env, appEnv)
		if err != nil {
			return false, err
		} else {
			batchNames = append(batchNames, batchName)
		}
	}

	for i := 0; i < len(appEnvs); i++ {
		appEnv := appEnvs[i]
		batchName := batchNames[i]
		batchWasRun := checkJobBatch(env, env.GetNetworkPolicyCanaryAppName(), appEnv, jobComponentName, batchName)
		if !batchWasRun {
			return false, fmt.Errorf("could not find batch job %s after running it", batchName)
		}
	}

	return true, nil
}

func checkJobBatch(env envUtil.Env, appName, appEnv string, jobComponentName string, batchName string) bool {

	ok, _ := test.WaitForCheckFuncOrTimeout(
		env,
		func(env envUtil.Env) (bool, interface{}) {
			return k8sjob.IsListedWithStatus(env, appName, appEnv, jobComponentName, batchName, "Succeeded")
		},
	)

	return ok
}

func startJobBatch(env envUtil.Env, appEnv string) (error, string) {
	jobBatchUrl := fmt.Sprintf("%s/startjobbatch", env.GetNetworkPolicyCanaryUrl(appEnv))
	response, err := http.Get(jobBatchUrl)
	if err != nil {
		return err, ""
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	var batchStatus models.BatchStatus
	unMarshalErr := json.Unmarshal(body, &batchStatus)
	if unMarshalErr != nil {
		return unMarshalErr, ""
	}
	if batchStatus.BatchName == "" {
		err = fmt.Errorf("no batchName attribute in job batch creation response. appEnv %s", appEnv)
		return err, ""
	}
	return nil, batchStatus.BatchName
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
