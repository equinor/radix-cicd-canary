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
	nspSleepIntervalConfig           = "nspSleepInterval"
	privateImageHubPasswordConfig    = "privateImageHubPassword"
)

// Env Holds all the environment variables
type Env struct {
	bearerToken                   string
	impersonateUser               string
	impersonateGroup              string
	clusterFQDN                   string
	radixAPIPrefix                string
	webhookPrefix                 string
	publicKey                     string
	privateKey                    string
	timeoutOfTest                 time.Duration
	sleepIntervalBetweenCheckFunc time.Duration
	sleepIntervalBetweenTestRuns  time.Duration
	nspSleepInterval              time.Duration
}

// NewEnv Constructor
func NewEnv() Env {
	return Env{
		getBearerToken(),
		getImpersonateUser(),
		getImpersonateGroup(),
		getClusterFQDN(),
		getRadixAPIPrefix(),
		getWebhookPrefix(),
		getPublicKey(),
		getPrivateKey(),
		timeoutOfTest(),
		getSleepIntervalBetweenCheckFunc(),
		getSleepIntervalBetweenTestRuns(),
		getNSPSleepInterval(),
	}
}

// GetBearerToken get bearer token either from token file or environment variable
func (env Env) GetBearerToken() string {
	return env.bearerToken
}

// GetImpersonateUser get impersonate user from config map
func (env Env) GetImpersonateUser() string {
	return env.impersonateUser
}

// GetImpersonateUserPointer get impersonate user from config map
func (env Env) GetImpersonateUserPointer() *string {
	return &env.impersonateUser
}

// GetImpersonateGroup get impersonate group from config map
func (env Env) GetImpersonateGroup() string {
	return env.impersonateGroup
}

// GetImpersonateGroupPointer get impersonate group from config map
func (env Env) GetImpersonateGroupPointer() *string {
	return &env.impersonateGroup
}

// GetClusterFQDN get Radix cluster FQDN from config map
func (env Env) GetClusterFQDN() string {
	return env.clusterFQDN
}

// GetRadixAPIPrefix get Radix API prefix from config map
func (env Env) GetRadixAPIPrefix() string {
	return env.radixAPIPrefix
}

// GetWebhookPrefix get Radix Webhook prefix
func (env Env) GetWebhookPrefix() string {
	return env.webhookPrefix
}

// GetPublicKey get public deploy key from config map
func (env Env) GetPublicKey() string {
	return env.publicKey
}

// GetPrivateKey get private deploy key from config map
func (env Env) GetPrivateKey() string {
	return env.privateKey
}

// GetPrivateImageHubPassword get private image hub password
func (env Env) GetPrivateImageHubPassword() string {
	return getConfigFromMap(privateImageHubPasswordConfig)
}

// GetTimeoutOfTest Get the time it should take before a test should time out from config map
func (env Env) GetTimeoutOfTest() time.Duration {
	return env.timeoutOfTest
}

// GetSleepIntervalBetweenCheckFunc Gets the sleep inteval between two checks from config map
func (env Env) GetSleepIntervalBetweenCheckFunc() time.Duration {
	return env.sleepIntervalBetweenCheckFunc
}

// GetSleepIntervalBetweenTestRuns Gets the sleep inteval between two test runs from config map
func (env Env) GetSleepIntervalBetweenTestRuns() time.Duration {
	return env.sleepIntervalBetweenTestRuns
}

// GetNSPSleepInterval Gets the sleep inteval between NSP test runs from config map
func (env Env) GetNSPSleepInterval() time.Duration {
	return env.nspSleepInterval
}

func getBearerToken() string {
	token, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/token")
	if err != nil {
		return os.Getenv("BEARER_TOKEN")
	}
	return string(token)
}

func getImpersonateUser() string {
	return getConfigFromMap(impersonateUserConfig)
}

func getImpersonateGroup() string {
	return getConfigFromMap(impersonateGroupConfig)
}

func getClusterFQDN() string {
	return getConfigFromMap(clusterFQDNConfig)
}

func getRadixAPIPrefix() string {
	return getConfigFromMap(radixAPIPrefixConfig)
}

func getWebhookPrefix() string {
	return getConfigFromMap(radixWebhookPrefixConfig)
}

func getPublicKey() string {
	return getConfigFromMap(publicKeyConfig)
}

func getPrivateKey() string {
	data, _ := base64.StdEncoding.DecodeString(getConfigFromMap(privateKeyBase64Config))
	return string(data)
}

func timeoutOfTest() time.Duration {
	timeout, err := strconv.Atoi(getConfigFromMap(timeoutOfTestConfig))
	if err != nil {
		log.Fatalf("Could not read %s. Err: %v", timeoutOfTestConfig, err)
	}

	return time.Duration(timeout) * time.Second
}

func getSleepIntervalBetweenCheckFunc() time.Duration {
	sleepInterval, err := strconv.Atoi(getConfigFromMap(sleepIntervalBetweenChecksConfig))
	if err != nil {
		log.Fatalf("Could not read %s. Err: %v", sleepIntervalBetweenChecksConfig, err)
	}

	return time.Duration(sleepInterval) * time.Second
}

func getSleepIntervalBetweenTestRuns() time.Duration {
	sleepInterval, err := strconv.Atoi(getConfigFromMap(sleepIntervalTestRunsConfig))
	if err != nil {
		log.Fatalf("Could not read %s. Err: %v", sleepIntervalTestRunsConfig, err)
	}

	return time.Duration(sleepInterval) * time.Second
}

func getNSPSleepInterval() time.Duration {
	sleepInterval, err := strconv.Atoi(getConfigFromMap(nspSleepIntervalConfig))
	if err != nil {
		log.Fatalf("Could not read %s. Err: %v", nspSleepIntervalConfig, err)
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
