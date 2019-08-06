package env

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	kubeUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/kubernetes"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	namespace                        = "radix-cicd-canary"
	configMapName                    = "radix-cicd-canary"
	impersonateUserConfig            = "impersonateUser"
	impersonateGroupConfig           = "impersonateGroup"
	clusterFQDNConfig                = "clusterFqdn"
	radixAPIPrefixConfig             = "radixApiPrefix"
	radixWebhookPrefixConfig         = "radixWebhookPrefix"
	publicKeyConfig                  = "publicKey"
	privateKeyBase64Config           = "privateKeyBase64"
	timeoutOfTestConfig              = "timeoutOfTest"
	sleepIntervalBetweenChecksConfig = "sleepIntervalBetweenChecks"
	sleepIntervalTestRunsConfig      = "sleepIntervalTestRuns"
)

// GetBearerToken get bearer token either from token file or environment variable
func GetBearerToken() string {
	token, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/token")
	if err != nil {
		return os.Getenv("BEARER_TOKEN")
	}
	return string(token)
}

// GetImpersonateUser get impersonate user from config map
func GetImpersonateUser() string {
	return getConfigFromMap(impersonateUserConfig)
}

// GetImpersonateGroup get impersonate group from config map
func GetImpersonateGroup() string {
	return getConfigFromMap(impersonateGroupConfig)
}

// GetClusterFQDN get Radix cluster FQDN from config map
func GetClusterFQDN() string {
	return getConfigFromMap(clusterFQDNConfig)
}

// GetRadixAPIPrefix get Radix API prefix from config map
func GetRadixAPIPrefix() string {
	return getConfigFromMap(radixAPIPrefixConfig)
}

// GetWebhookPrefix get Radix Webhook prefix
func GetWebhookPrefix() string {
	return getConfigFromMap(radixWebhookPrefixConfig)
}

// GetPublicKey get public deploy key from config map
func GetPublicKey() string {
	return getConfigFromMap(publicKeyConfig)
}

// GetPrivateKey get private deploy key from config map
func GetPrivateKey() string {
	data, _ := base64.StdEncoding.DecodeString(getConfigFromMap(privateKeyBase64Config))
	return string(data)
}

// TimeoutOfTest Get the time it should take before a test should time out from config map
func TimeoutOfTest() time.Duration {
	timeout, err := strconv.Atoi(getConfigFromMap(timeoutOfTestConfig))
	if err != nil {
		log.Fatalf("Could not read %s. Err: %v", timeoutOfTestConfig, err)
	}

	return time.Duration(timeout) * time.Second
}

// GetSleepIntervalBetweenCheckFunc Gets the sleep inteval between two checks from config map
func GetSleepIntervalBetweenCheckFunc() time.Duration {
	sleepInterval, err := strconv.Atoi(getConfigFromMap(sleepIntervalBetweenChecksConfig))
	if err != nil {
		log.Fatalf("Could not read %s. Err: %v", sleepIntervalBetweenChecksConfig, err)
	}

	return time.Duration(sleepInterval) * time.Second
}

// GetSleepIntervalBetweenTestRuns Gets the sleep inteval between two test runs from config map
func GetSleepIntervalBetweenTestRuns() time.Duration {
	sleepInterval, err := strconv.Atoi(getConfigFromMap(sleepIntervalTestRunsConfig))
	if err != nil {
		log.Fatalf("Could not read %s. Err: %v", sleepIntervalTestRunsConfig, err)
	}

	return time.Duration(sleepInterval) * time.Second
}

// SetRequiredEnvironmentVariablesForTest Sets test environment variables, that would come from
// launch config, when running complete scenario
func SetRequiredEnvironmentVariablesForTest() {
	os.Setenv("BEARER_TOKEN", "")
}

func getConfigFromMap(config string) string {
	kubeClient := kubeUtils.GetKubernetesClient()
	configmap, err := kubeClient.CoreV1().ConfigMaps(namespace).Get(configMapName, metav1.GetOptions{})
	if err != nil {
		log.Fatalf("Error reading config map: %v", err)
	}
	configValue := configmap.Data[config]
	return configValue
}
