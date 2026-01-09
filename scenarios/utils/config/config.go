package config

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	kubeUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/kubernetes"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/tokensource"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
	v1 "k8s.io/api/core/v1"
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
	nspDNSLookupTimeoutConfig                 = "nspDNSLookupTimeout"
	nspNSPReachIngressTimeoutConfig           = "nspNSPReachIngressTimeout"
	nspNSPReachServiceTimeoutConfig           = "nspNSPReachServiceTimeout"
	privateImageHubPasswordConfig             = "privateImageHubPassword"
	networkPolicyCanaryPasswordConfig         = "networkPolicyCanaryPassword"
	networkPolicyCanaryAppNameConfig          = "networkPolicyCanaryAppName"
	networkPolicyCanaryJobComponentNameConfig = "networkPolicyCanaryJobComponentName"
	appAdminGroupConfig                       = "appAdminGroup"
	appReaderGroupsConfig                     = "appReaderGroup"
	envVarSuiteList                           = "SUITE_LIST"
	envVarIsBlacklist                         = "SUITE_LIST_IS_BLACKLIST"
	envVarLogLevel                            = "LOG_LEVEL"
	envVarPrettyPrint                         = "PRETTY_PRINT"
	envUseLocalRadixApi                       = "USE_LOCAL_RADIX_API"
	envUseLocalGitHubWebHookApi               = "USE_LOCAL_GITHUB_WEBHOOK_API"
	serviceAccountTokenFile                   = "/var/run/secrets/kubernetes.io/serviceaccount/token"
)

// Config Holds all the environment variables
type Config struct {
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
	networkPolicyCanaryPassword         string
	timeoutOfTest                       time.Duration
	sleepIntervalBetweenCheckFunc       time.Duration
	sleepIntervalBetweenTestRuns        time.Duration
	nspSleepInterval                    time.Duration
	nspLongSleepInterval                time.Duration
	nspDNSLookupTimeout                 time.Duration
	nspNSPReachIngressTimeout           time.Duration
	nspNSPReachServiceTimeout           time.Duration
	suiteList                           []string
	suiteListIsBlacklist                bool // suiteList is a whitelist by default
	networkPolicyCanaryAppName          string
	networkPolicyCanaryJobComponentName string
	appAdminGroup                       string
	appReaderGroup                      string
	tokenSource                         oauth2.TokenSource
}

var configmap *v1.ConfigMap

func init() {
	kubeClient := kubeUtils.GetKubernetesClient()
	cm, err := kubeClient.CoreV1().ConfigMaps(namespace).Get(context.Background(), configMapName, metav1.GetOptions{})
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("error reading config map")
	}
	configmap = cm
}

// NewConfig Constructor
func NewConfig() Config {
	return Config{
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
		getNetworkPolicyCanaryPassword(),
		timeoutOfTest(),
		getSleepIntervalBetweenCheckFunc(),
		getSleepIntervalBetweenTestRuns(),
		getNSPSleepInterval(),
		getNSPSleepInterval(),
		getNSPDNSLookupTimeout(),
		getNSPReachIngressTimeout(),
		getNSPReachServiceTimeout(),
		getSuiteList(),
		getIsBlacklist(),
		getNetworkPolicyCanaryAppName(),
		getNetworkPolicyCanaryJobComponentName(),
		getAppAdminGroup(),
		getAppReaderGroup(),
		getTokenSource(),
	}
}

// GetImpersonateUser get impersonate user from config map
func (cfg *Config) GetImpersonateUser() *string {
	if len(cfg.impersonateUser) > 0 {
		return &cfg.impersonateUser
	}
	return nil
}

// GetImpersonateGroups get list of groups for impersonation
func (cfg *Config) GetImpersonateGroups() *string {
	s := cfg.impersonateGroup + "," + cfg.GetAppAdminGroup()
	return &s
}

func (cfg *Config) GetAppAdminGroup() string {
	return cfg.appAdminGroup
}

// GetClusterFQDN get Radix cluster FQDN from config map
func (cfg *Config) GetClusterFQDN() string {
	return cfg.clusterFQDN
}

// GetPublicKey get public deploy key from config map
func (cfg *Config) GetPublicKey() string {
	return cfg.publicKey
}

// GetPrivateKey get private deploy key from config map
func (cfg *Config) GetPrivateKey() string {
	return cfg.privateKey
}

// GetPublicKeyCanary3 get public deploy key from config map
func (cfg *Config) GetPublicKeyCanary3() string {
	return cfg.publicKeyCanary3
}

// GetPrivateKeyCanary3 get private deploy key from config map
func (cfg *Config) GetPrivateKeyCanary3() string {
	return cfg.privateKeyCanary3
}

// GetPublicKeyCanary4 get public deploy key from config map
func (cfg *Config) GetPublicKeyCanary4() string {
	return cfg.publicKeyCanary4
}

// GetPrivateKeyCanary4 get private deploy key from config map
func (cfg *Config) GetPrivateKeyCanary4() string {
	return cfg.privateKeyCanary4
}

// GetPrivateImageHubPassword get private image hub password
func (cfg *Config) GetPrivateImageHubPassword() string {
	return getConfigFromMap(privateImageHubPasswordConfig)
}

// GetNetworkPolicyCanaryPassword get networkpolicy-canary HTTP password from environment
func (cfg *Config) GetNetworkPolicyCanaryPassword() string {
	return cfg.networkPolicyCanaryPassword
}

// GetTimeoutOfTest Get the time it should take before a test should time out from config map
func (cfg *Config) GetTimeoutOfTest() time.Duration {
	return cfg.timeoutOfTest
}

// GetSleepIntervalBetweenCheckFunc Gets the sleep inteval between two checks from config map
func (cfg *Config) GetSleepIntervalBetweenCheckFunc() time.Duration {
	return cfg.sleepIntervalBetweenCheckFunc
}

// GetSleepIntervalBetweenTestRuns Gets the sleep inteval between two test runs from config map
func (cfg *Config) GetSleepIntervalBetweenTestRuns() time.Duration {
	return cfg.sleepIntervalBetweenTestRuns
}

// GetNSPSleepInterval Gets the sleep inteval between NSP test runs from config map
func (cfg *Config) GetNSPSleepInterval() time.Duration {
	return cfg.nspSleepInterval
}

// GetNSPLongSleepInterval Gets the sleep inteval between NSPLong test runs from config map
func (cfg *Config) GetNSPLongSleepInterval() time.Duration {
	return cfg.nspLongSleepInterval
}

// GetSuiteList Gets a filter list for which suites to run
func (cfg *Config) GetSuiteList() []string {
	return cfg.suiteList
}

// GetSuiteListIsBlacklist Gets whether suiteList is considered a blacklist
func (cfg *Config) GetSuiteListIsBlacklist() bool {
	return cfg.suiteListIsBlacklist
}

func (cfg *Config) GetTokenSource() oauth2.TokenSource {
	return cfg.tokenSource
}

// GetLogLevel Gets log level
func (cfg *Config) GetLogLevel() zerolog.Level {
	lvl, err := zerolog.ParseLevel(os.Getenv(envVarLogLevel))
	if err != nil || lvl == zerolog.NoLevel {
		return zerolog.InfoLevel
	}
	return lvl
}
func (cfg *Config) GetPrettyPrint() bool {
	pretty := os.Getenv(envVarPrettyPrint)
	return envVarIsTrueOrYes(pretty)
}

func (cfg *Config) GetRadixAPIURL() string {
	if useLocalRadixApi() {
		return "localhost:3002"
	} else {
		return fmt.Sprintf("%s.%s", cfg.getRadixAPIPrefix(), cfg.GetClusterFQDN())
	}
}

func (cfg *Config) GetGitHubWebHookAPIURL() string {
	if useLocalGitHubWebHookApi() {
		return "http://localhost:3001"
	} else {
		return fmt.Sprintf("https://%s.%s", cfg.getWebHookPrefix(), cfg.GetClusterFQDN())
	}
}

func (cfg *Config) GetNetworkPolicyCanaryUrl(appEnv string) string {
	canaryURLPrefix := fmt.Sprintf("https://web-%s-%s", cfg.GetNetworkPolicyCanaryAppName(), appEnv)
	return fmt.Sprintf("%s.%s", canaryURLPrefix, cfg.GetClusterFQDN())
}

func (cfg *Config) GetNetworkPolicyCanaryAppName() string {
	return cfg.networkPolicyCanaryAppName
}

func (cfg *Config) GetNetworkPolicyCanaryJobComponentName() string {
	return cfg.networkPolicyCanaryJobComponentName
}

func (cfg *Config) GetRadixAPISchemes() []string {
	if useLocalRadixApi() {
		return []string{"http"}
	} else {
		return []string{"https"}
	}
}

func getTokenSource() oauth2.TokenSource {
	var ts oauth2.TokenSource
	if _, err := os.Stat(serviceAccountTokenFile); err == nil {
		ts = tokensource.FromJwtCallback(func() (string, error) {
			token, err := os.ReadFile(serviceAccountTokenFile)
			return string(token), errors.WithStack(err)
		})
	} else if envToken := os.Getenv("BEARER_TOKEN"); len(envToken) > 0 {
		ts = tokensource.FromJwtCallback(func() (string, error) {
			return envToken, nil
		})
	}

	if ts == nil {
		panic(errors.New("failed to create TokenSource from service account token or environment variable"))
	}

	return oauth2.ReuseTokenSource(nil, ts)
}

func getImpersonateUser() string {
	return getConfigFromMap(impersonateUserConfig)
}

func getImpersonateGroup() string {
	return getConfigFromMap(impersonateGroupConfig)
}

func getAppAdminGroup() string {
	return getConfigFromMap(appAdminGroupConfig)
}

func getAppReaderGroup() string {
	return getConfigFromMap(appReaderGroupsConfig)
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

func getNetworkPolicyCanaryPassword() string {
	return getConfigFromMap(networkPolicyCanaryPasswordConfig)
}

func timeoutOfTest() time.Duration {
	timeout, err := strconv.Atoi(getConfigFromMap(timeoutOfTestConfig))
	if err != nil {
		log.Fatal().Stack().Err(err).Str("key", timeoutOfTestConfig).Msg("Could not read config")
	}

	return time.Duration(timeout) * time.Second
}

func getSleepIntervalBetweenCheckFunc() time.Duration {
	sleepInterval, err := strconv.Atoi(getConfigFromMap(sleepIntervalBetweenChecksConfig))
	if err != nil {
		log.Fatal().Stack().Err(err).Str("key", sleepIntervalBetweenChecksConfig).Msg("Could not read config")
	}

	return time.Duration(sleepInterval) * time.Second
}

func getSleepIntervalBetweenTestRuns() time.Duration {
	sleepInterval, err := strconv.Atoi(getConfigFromMap(sleepIntervalTestRunsConfig))
	if err != nil {
		log.Fatal().Stack().Err(err).Str("key", sleepIntervalTestRunsConfig).Msg("Could not read config")
	}

	return time.Duration(sleepInterval) * time.Second
}

func getNetworkPolicyCanaryAppName() string {
	appName := getConfigFromMap(networkPolicyCanaryAppNameConfig)
	if appName == "" {
		log.Fatal().Str("key", networkPolicyCanaryAppNameConfig).Msg("Could not read config")
	}
	return appName
}

func getNetworkPolicyCanaryJobComponentName() string {
	jobComponentName := getConfigFromMap(networkPolicyCanaryJobComponentNameConfig)
	if jobComponentName == "" {
		log.Fatal().Str("key", networkPolicyCanaryJobComponentNameConfig).Msg("Could not read from configmap")
	}
	return jobComponentName
}

func getNSPSleepInterval() time.Duration {
	return time.Duration(getConfigMapInt(nspSleepIntervalConfig)) * time.Second
}

func getNSPDNSLookupTimeout() time.Duration {
	return time.Duration(getConfigMapInt(nspDNSLookupTimeoutConfig)) * time.Second
}

func getNSPReachIngressTimeout() time.Duration {
	return time.Duration(getConfigMapInt(nspNSPReachIngressTimeoutConfig)) * time.Second
}

func getNSPReachServiceTimeout() time.Duration {
	return time.Duration(getConfigMapInt(nspNSPReachServiceTimeoutConfig)) * time.Second
}

func GetNSPLongSleepInterval() time.Duration {
	return time.Duration(getConfigMapInt(nspLongSleepIntervalConfig)) * time.Second
}

func getConfigMapInt(key string) int {
	value, err := strconv.Atoi(getConfigFromMap(key))
	if err != nil {
		log.Fatal().Stack().Err(err).Str("key", key).Msg("Could not read config")
	}
	return value
}

// SetRequiredEnvironmentVariablesForTest Sets test environment variables, that would come from
// launch config, when running complete scenario
func SetRequiredEnvironmentVariablesForTest() {
	_ = os.Setenv("GODEBUG", "http2server=0,http2client=0")
	_ = os.Setenv("BEARER_TOKEN", "")
}

func getConfigFromMap(config string) string {
	configValue, found := configmap.Data[config]
	if !found {
		log.Fatal().Str("key", config).Msg("key not found in configmap")
	}
	return configValue
}

func getSuiteList() []string {
	suiteList := os.Getenv(envVarSuiteList)
	// return empty list if no values (Split would return [""])
	if len(suiteList) == 0 {
		return make([]string, 0)
	}
	split := strings.Split(suiteList, ":")
	return split
}

func getIsBlacklist() bool {
	return envVarIsTrueOrYes(strings.ToLower(os.Getenv(envVarIsBlacklist)))
}

func (cfg *Config) getRadixAPIPrefix() string {
	return cfg.radixAPIPrefix
}

func (cfg *Config) getWebHookPrefix() string {
	return cfg.webhookPrefix
}

func (cfg *Config) GetAppReaderGroup() string {
	return cfg.appReaderGroup
}

// GetNSPDNSLookupTimeout Get DNS lookup timeout
func (cfg *Config) GetNSPDNSLookupTimeout() time.Duration {
	return cfg.nspDNSLookupTimeout
}

// GetNSPReachIngressTimeout Get reach ingress timeout
func (cfg *Config) GetNSPReachIngressTimeout() time.Duration {
	return cfg.nspNSPReachIngressTimeout
}

// GetNSPReachServiceTimeout Get reach service timeout
func (cfg *Config) GetNSPReachServiceTimeout() time.Duration {
	return cfg.nspNSPReachServiceTimeout
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
