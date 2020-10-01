package env

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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
	publicKeyCanary3Config           = "publicKeyCanary3"
	privateKeyCanary3Base64Config    = "privateKeyCanary3Base64"
	timeoutOfTestConfig              = "timeoutOfTest"
	sleepIntervalBetweenChecksConfig = "sleepIntervalBetweenChecks"
	sleepIntervalTestRunsConfig      = "sleepIntervalTestRuns"
	nspSleepIntervalConfig           = "nspSleepInterval"
	privateImageHubPasswordConfig    = "privateImageHubPassword"
	envVarSuiteList                  = "SUITE_LIST"
	envVarIsBlacklist                = "SUITE_LIST_IS_BLACKLIST"
	envVarLogLevel                   = "LOG_LEVEL"
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
	publicKeyCanary3              string
	privateKeyCanary3             string
	timeoutOfTest                 time.Duration
	sleepIntervalBetweenCheckFunc time.Duration
	sleepIntervalBetweenTestRuns  time.Duration
	nspSleepInterval              time.Duration
	suiteList                     []string
	suiteListIsBlacklist          bool // suiteList is a whitelist by default
	isDebugLogLevel               bool
	isWarningLogLevel             bool
	isErrorLogLevel               bool
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
		getPublicKeyCanary3(),
		getPrivateKeyCanary3(),
		timeoutOfTest(),
		getSleepIntervalBetweenCheckFunc(),
		getSleepIntervalBetweenTestRuns(),
		getNSPSleepInterval(),
		getSuiteList(),
		getIsBlacklist(),
		isDebugLogLevel(),
		isWarningLogLevel(),
		isErrorLogLevel(),
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

// GetPublicKeyCanary3 get public deploy key from config map
func (env Env) GetPublicKeyCanary3() string {
	return env.publicKeyCanary3
}

// GetPrivateKeyCanary3 get private deploy key from config map
func (env Env) GetPrivateKeyCanary3() string {
	return env.privateKeyCanary3
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

// GetSuiteList Gets a filter list for which suites to run
func (env Env) GetSuiteList() []string {
	return env.suiteList
}

// GetSuiteListIsBlacklist Gets whether suiteList is considered a blacklist
func (env Env) GetSuiteListIsBlacklist() bool {
	return env.suiteListIsBlacklist
}

// GetLogLevel Gets log level
func (env Env) GetLogLevel() log.Level {
	switch {
	case isDebugLogLevel():
		return log.DebugLevel
	case isWarningLogLevel():
		return log.WarnLevel
	case isErrorLogLevel():
		return log.ErrorLevel
	default:
		return log.InfoLevel
	}
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

func getPublicKeyCanary3() string {
	return getConfigFromMap(publicKeyCanary3Config)
}

func getPrivateKeyCanary3() string {
	data, _ := base64.StdEncoding.DecodeString(getConfigFromMap(privateKeyCanary3Base64Config))
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
	os.Setenv("GODEBUG", "http2server=0,http2client=0")
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

func getSuiteList() []string {
	suiteList := os.Getenv(envVarSuiteList)
	// return empty list if no values (Split would return [""])
	if len(suiteList) == 0 {
		return make([]string, 0)
	} else {
		return strings.Split(suiteList, ":")
	}
}

func getIsBlacklist() bool {
	suiteListIsBlacklist := strings.ToLower(os.Getenv(envVarIsBlacklist))
	return suiteListIsBlacklist == "true" || suiteListIsBlacklist == "yes"
}

func isDebugLogLevel() bool {
	return strings.EqualFold(os.Getenv(envVarLogLevel), "DEBUG")
}

func isWarningLogLevel() bool {
	return strings.EqualFold(os.Getenv(envVarLogLevel), "WARNING")
}

func isErrorLogLevel() bool {
	return strings.EqualFold(os.Getenv(envVarLogLevel), "ERROR")
}
