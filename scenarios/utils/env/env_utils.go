package env

import (
	"context"
	"encoding/base64"
	"fmt"
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
	namespace                                 = "radix-cicd-canary"
	configMapName                             = "radix-cicd-canary"
	impersonateUserConfig                     = "impersonateUser"
	impersonateGroupConfig                    = "impersonateGroup"
	clusterFQDNConfig                         = "clusterFqdn"
	radixAPIPrefixConfig                      = "radixApiPrefix"
	radixWebhookPrefixConfig                  = "radixWebhookPrefix"
	publicKeyConfig                           = "publicKey"
	privateKeyBase64Config                    = "privateKeyBase64"
	publicKeyCanary3Config                    = "publicKeyCanary3"
	privateKeyCanary3Base64Config             = "privateKeyCanary3Base64"
	publicKeyCanary4Config                    = "publicKeyCanary4"
	privateKeyCanary4Base64Config             = "privateKeyCanary4Base64"
	timeoutOfTestConfig                       = "timeoutOfTest"
	sleepIntervalBetweenChecksConfig          = "sleepIntervalBetweenChecks"
	sleepIntervalTestRunsConfig               = "sleepIntervalTestRuns"
	nspSleepIntervalConfig                    = "nspSleepInterval"
	nspLongSleepIntervalConfig                = "nspLongSleepInterval"
	privateImageHubPasswordConfig             = "privateImageHubPassword"
	networkPolicyCanaryAppNameConfig          = "networkPolicyCanaryAppName"
	golangCanaryUrlConfig                     = "golangCanaryUrl"
	networkPolicyCanaryJobComponentNameConfig = "networkPolicyCanaryJobComponentName"
	envVarSuiteList                           = "SUITE_LIST"
	envVarIsBlacklist                         = "SUITE_LIST_IS_BLACKLIST"
	envVarLogLevel                            = "LOG_LEVEL"
	envUseLocalRadixApi                       = "USE_LOCAL_RADIX_API"
	envUseLocalGitHubWebHookApi               = "USE_LOCAL_GITHUB_WEBHOOK_API"
)

// Env Holds all the environment variables
type Env struct {
	bearerToken                         string
	impersonateUser                     string
	impersonateGroup                    string
	clusterFQDN                         string
	radixAPIPrefix                      string
	webhookPrefix                       string
	publicKey                           string
	privateKey                          string
	publicKeyCanary3                    string
	privateKeyCanary3                   string
	publicKeyCanary4                    string
	privateKeyCanary4                   string
	timeoutOfTest                       time.Duration
	sleepIntervalBetweenCheckFunc       time.Duration
	sleepIntervalBetweenTestRuns        time.Duration
	nspSleepInterval                    time.Duration
	nspLongSleepInterval                time.Duration
	suiteList                           []string
	suiteListIsBlacklist                bool // suiteList is a whitelist by default
	isDebugLogLevel                     bool
	isWarningLogLevel                   bool
	isErrorLogLevel                     bool
	networkPolicyCanaryAppName          string
	networkPolicyCanaryJobComponentName string
	golangCanaryUrl                     string
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
		getPublicKeyCanary4(),
		getPrivateKeyCanary4(),
		timeoutOfTest(),
		getSleepIntervalBetweenCheckFunc(),
		getSleepIntervalBetweenTestRuns(),
		getNSPSleepInterval(),
		GetNSPLongSleepInterval(),
		getSuiteList(),
		getIsBlacklist(),
		isDebugLogLevel(),
		isWarningLogLevel(),
		isErrorLogLevel(),
		getNetworkPolicyCanaryAppName(),
		getNetworkPolicyCanaryJobComponentName(),
		getGolangCanaryUrl(),
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

// GetPublicKeyCanary4 get public deploy key from config map
func (env Env) GetPublicKeyCanary4() string {
	return env.publicKeyCanary4
}

// GetPrivateKeyCanary4 get private deploy key from config map
func (env Env) GetPrivateKeyCanary4() string {
	return env.privateKeyCanary4
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

// GetNSPLongSleepInterval Gets the sleep inteval between NSPLong test runs from config map
func (env Env) GetNSPLongSleepInterval() time.Duration {
	return env.nspLongSleepInterval
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

func (env Env) GetRadixAPIURL() string {
	if useLocalRadixApi() {
		return "localhost:3002"
	} else {
		return fmt.Sprintf("%s.%s", env.getRadixAPIPrefix(), env.GetClusterFQDN())
	}
}

func (env Env) GetGitHubWebHookAPIURL() string {
	if useLocalGitHubWebHookApi() {
		return "http://localhost:3001"
	} else {
		return fmt.Sprintf("https://%s.%s", env.getWebHookPrefix(), env.GetClusterFQDN())
	}
}

func (env Env) GetNetworkPolicyCanaryUrl(appEnv string) string {
	canaryURLPrefix := fmt.Sprintf("https://web-%s-%s", env.GetNetworkPolicyCanaryAppName(), appEnv)
	return fmt.Sprintf("%s.%s", canaryURLPrefix, env.GetClusterFQDN())
}

func (env Env) GetGolangCanaryUrl() string {
	return env.golangCanaryUrl
}

func (env Env) GetNetworkPolicyCanaryAppName() string {
	return env.networkPolicyCanaryAppName
}

func (env Env) GetNetworkPolicyCanaryJobComponentName() string {
	return env.networkPolicyCanaryJobComponentName
}

func (env Env) GetRadixAPISchemes() []string {
	if useLocalRadixApi() {
		return []string{"http"}
	} else {
		return []string{"https"}
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

func getPublicKeyCanary4() string {
	return getConfigFromMap(publicKeyCanary4Config)
}

func getPrivateKeyCanary4() string {
	data, _ := base64.StdEncoding.DecodeString(getConfigFromMap(privateKeyCanary4Base64Config))
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

func getNetworkPolicyCanaryAppName() string {
	appName := getConfigFromMap(networkPolicyCanaryAppNameConfig)
	if appName == "" {
		log.Fatalf("Could not read %s from configmap", networkPolicyCanaryAppNameConfig)
	}
	return appName
}

func getGolangCanaryUrl() string {
	url := getConfigFromMap(golangCanaryUrlConfig)
	if url == "" {
		log.Fatalf("Could not read %s from configmap", golangCanaryUrlConfig)
	}
	return url
}

func getNetworkPolicyCanaryJobComponentName() string {
	jobComponentName := getConfigFromMap(networkPolicyCanaryJobComponentNameConfig)
	if jobComponentName == "" {
		log.Fatalf("Could not read %s from configmap", networkPolicyCanaryJobComponentNameConfig)
	}
	return jobComponentName
}

func getNSPSleepInterval() time.Duration {
	sleepInterval, err := strconv.Atoi(getConfigFromMap(nspSleepIntervalConfig))
	if err != nil {
		log.Fatalf("Could not read %s. Err: %v", nspSleepIntervalConfig, err)
	}

	return time.Duration(sleepInterval) * time.Second
}

func GetNSPLongSleepInterval() time.Duration {
	sleepInterval, err := strconv.Atoi(getConfigFromMap(nspLongSleepIntervalConfig))
	if err != nil {
		log.Fatalf("Could not read %s. Err: %v", nspLongSleepIntervalConfig, err)
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
	configmap, err := kubeClient.CoreV1().ConfigMaps(namespace).Get(context.Background(), configMapName, metav1.GetOptions{})
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
	return envVarIsTrueOrYes(strings.ToLower(os.Getenv(envVarIsBlacklist)))
}

func (env Env) getRadixAPIPrefix() string {
	return env.radixAPIPrefix
}

func (env Env) getWebHookPrefix() string {
	return env.webhookPrefix
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

func useLocalRadixApi() bool {
	return envVarIsTrueOrYes(os.Getenv(envUseLocalRadixApi))
}

func useLocalGitHubWebHookApi() bool {
	return envVarIsTrueOrYes(os.Getenv(envUseLocalGitHubWebHookApi))
}

func envVarIsTrueOrYes(envVar string) bool {
	return strings.EqualFold(envVar, "true") || strings.EqualFold(envVar, "yes")
}
