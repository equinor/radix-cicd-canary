package config

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"errors"

	kubeUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/kubernetes"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/tokensource"
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
	envVarRadixAPIServer                      = "RADIX_API_SERVER"
	serviceAccountTokenFile                   = "/var/run/secrets/kubernetes.io/serviceaccount/token"
)

// Config Holds all the environment variables
type Config struct {
	impersonateUser                     string
	impersonateGroup                    string
	clusterFQDN                         string
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
	privateImageHubPassword             string
	tokenSource                         oauth2.TokenSource
}

// NewConfig Constructor
func NewConfig() Config {
	kubeClient := kubeUtils.GetKubernetesClient()
	cm, err := kubeClient.CoreV1().ConfigMaps(namespace).Get(context.Background(), configMapName, metav1.GetOptions{})
	if err != nil {
		log.Fatal().Err(err).Msg("error reading config map")
	}

	return Config{
		impersonateUser:                     getImpersonateUser(cm),
		impersonateGroup:                    getImpersonateGroup(cm),
		clusterFQDN:                         getClusterFQDN(cm),
		publicKey:                           getPublicKey(cm),
		privateKey:                          getPrivateKey(cm),
		publicKeyCanary3:                    getPublicKeyCanary3(cm),
		privateKeyCanary3:                   getPrivateKeyCanary3(cm),
		publicKeyCanary4:                    getPublicKeyCanary4(cm),
		privateKeyCanary4:                   getPrivateKeyCanary4(cm),
		networkPolicyCanaryPassword:         getNetworkPolicyCanaryPassword(cm),
		timeoutOfTest:                       timeoutOfTest(cm),
		sleepIntervalBetweenCheckFunc:       getSleepIntervalBetweenCheckFunc(cm),
		sleepIntervalBetweenTestRuns:        getSleepIntervalBetweenTestRuns(cm),
		nspSleepInterval:                    getNSPSleepInterval(cm),
		nspLongSleepInterval:                GetNSPLongSleepInterval(cm),
		nspDNSLookupTimeout:                 getNSPDNSLookupTimeout(cm),
		nspNSPReachIngressTimeout:           getNSPReachIngressTimeout(cm),
		nspNSPReachServiceTimeout:           getNSPReachServiceTimeout(cm),
		suiteList:                           getSuiteList(),
		suiteListIsBlacklist:                getIsBlacklist(),
		networkPolicyCanaryAppName:          getNetworkPolicyCanaryAppName(cm),
		networkPolicyCanaryJobComponentName: getNetworkPolicyCanaryJobComponentName(cm),
		appAdminGroup:                       getAppAdminGroup(cm),
		appReaderGroup:                      getAppReaderGroup(cm),
		privateImageHubPassword:             getPrivateImageHubPassword(cm),
		tokenSource:                         getTokenSource(),
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
	return cfg.privateImageHubPassword
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
	if os.Getenv(envVarRadixAPIServer) != "" {
		return os.Getenv(envVarRadixAPIServer)
	}

	return fmt.Sprintf("https://api.%s", cfg.GetClusterFQDN())
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

func getTokenSource() oauth2.TokenSource {
	var ts oauth2.TokenSource
	if _, err := os.Stat(serviceAccountTokenFile); err == nil {
		ts = tokensource.FromJwtCallback(func() (string, error) {
			token, err := os.ReadFile(serviceAccountTokenFile)
			if err != nil {
				return "", fmt.Errorf("failed to read file %s: %w", serviceAccountTokenFile, err)
			}
			return string(token), nil
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

func getImpersonateUser(configMap *v1.ConfigMap) string {
	return getConfigFromMap(configMap, impersonateUserConfig)
}

func getImpersonateGroup(configMap *v1.ConfigMap) string {
	return getConfigFromMap(configMap, impersonateGroupConfig)
}

func getAppAdminGroup(configMap *v1.ConfigMap) string {
	return getConfigFromMap(configMap, appAdminGroupConfig)
}

func getAppReaderGroup(configMap *v1.ConfigMap) string {
	return getConfigFromMap(configMap, appReaderGroupsConfig)
}

func getClusterFQDN(configMap *v1.ConfigMap) string {
	return getConfigFromMap(configMap, clusterFQDNConfig)
}

func getPublicKey(configMap *v1.ConfigMap) string {
	return getConfigFromMap(configMap, publicKeyConfig)
}

func getPrivateKey(configMap *v1.ConfigMap) string {
	data, _ := base64.StdEncoding.DecodeString(getConfigFromMap(configMap, privateKeyBase64Config))
	return string(data)
}

func getPublicKeyCanary3(configMap *v1.ConfigMap) string {
	return getConfigFromMap(configMap, publicKeyCanary3Config)
}

func getPrivateKeyCanary3(configMap *v1.ConfigMap) string {
	data, _ := base64.StdEncoding.DecodeString(getConfigFromMap(configMap, privateKeyCanary3Base64Config))
	return string(data)
}

func getPublicKeyCanary4(configMap *v1.ConfigMap) string {
	return getConfigFromMap(configMap, publicKeyCanary4Config)
}

func getPrivateKeyCanary4(configMap *v1.ConfigMap) string {
	data, _ := base64.StdEncoding.DecodeString(getConfigFromMap(configMap, privateKeyCanary4Base64Config))
	return string(data)
}

func getNetworkPolicyCanaryPassword(configMap *v1.ConfigMap) string {
	return getConfigFromMap(configMap, networkPolicyCanaryPasswordConfig)
}

func timeoutOfTest(configMap *v1.ConfigMap) time.Duration {
	timeout, err := strconv.Atoi(getConfigFromMap(configMap, timeoutOfTestConfig))
	if err != nil {
		log.Fatal().Err(err).Str("key", timeoutOfTestConfig).Msg("Could not read config")
	}

	return time.Duration(timeout) * time.Second
}

func getSleepIntervalBetweenCheckFunc(configMap *v1.ConfigMap) time.Duration {
	sleepInterval, err := strconv.Atoi(getConfigFromMap(configMap, sleepIntervalBetweenChecksConfig))
	if err != nil {
		log.Fatal().Err(err).Str("key", sleepIntervalBetweenChecksConfig).Msg("Could not read config")
	}

	return time.Duration(sleepInterval) * time.Second
}

func getSleepIntervalBetweenTestRuns(configMap *v1.ConfigMap) time.Duration {
	sleepInterval, err := strconv.Atoi(getConfigFromMap(configMap, sleepIntervalTestRunsConfig))
	if err != nil {
		log.Fatal().Err(err).Str("key", sleepIntervalTestRunsConfig).Msg("Could not read config")
	}

	return time.Duration(sleepInterval) * time.Second
}

func getNetworkPolicyCanaryAppName(configMap *v1.ConfigMap) string {
	appName := getConfigFromMap(configMap, networkPolicyCanaryAppNameConfig)
	if appName == "" {
		log.Fatal().Str("key", networkPolicyCanaryAppNameConfig).Msg("Could not read config")
	}
	return appName
}

func getPrivateImageHubPassword(configMap *v1.ConfigMap) string {
	password := getConfigFromMap(configMap, privateImageHubPasswordConfig)
	if password == "" {
		log.Fatal().Str("key", privateImageHubPasswordConfig).Msg("Could not read config")
	}
	return password
}

func getNetworkPolicyCanaryJobComponentName(configMap *v1.ConfigMap) string {
	jobComponentName := getConfigFromMap(configMap, networkPolicyCanaryJobComponentNameConfig)
	if jobComponentName == "" {
		log.Fatal().Str("key", networkPolicyCanaryJobComponentNameConfig).Msg("Could not read from configmap")
	}
	return jobComponentName
}

func getNSPSleepInterval(configMap *v1.ConfigMap) time.Duration {
	return time.Duration(getConfigMapInt(configMap, nspSleepIntervalConfig)) * time.Second
}

func getNSPDNSLookupTimeout(configMap *v1.ConfigMap) time.Duration {
	return time.Duration(getConfigMapInt(configMap, nspDNSLookupTimeoutConfig)) * time.Second
}

func getNSPReachIngressTimeout(configMap *v1.ConfigMap) time.Duration {
	return time.Duration(getConfigMapInt(configMap, nspNSPReachIngressTimeoutConfig)) * time.Second
}

func getNSPReachServiceTimeout(configMap *v1.ConfigMap) time.Duration {
	return time.Duration(getConfigMapInt(configMap, nspNSPReachServiceTimeoutConfig)) * time.Second
}

func GetNSPLongSleepInterval(configMap *v1.ConfigMap) time.Duration {
	return time.Duration(getConfigMapInt(configMap, nspLongSleepIntervalConfig)) * time.Second
}

func getConfigMapInt(configMap *v1.ConfigMap, key string) int {
	value, err := strconv.Atoi(getConfigFromMap(configMap, key))
	if err != nil {
		log.Fatal().Err(err).Str("key", key).Msg("Could not read config")
	}
	return value
}

// SetRequiredEnvironmentVariablesForTest Sets test environment variables, that would come from
// launch config, when running complete scenario
func SetRequiredEnvironmentVariablesForTest() {
	_ = os.Setenv("GODEBUG", "http2server=0,http2client=0")
	_ = os.Setenv("BEARER_TOKEN", "")
}

func getConfigFromMap(configMap *v1.ConfigMap, key string) string {
	configValue, found := configMap.Data[key]
	if !found {
		log.Fatal().Str("key", key).Msg("key not found in configmap")
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

func envVarIsTrueOrYes(envVar string) bool {
	return strings.EqualFold(envVar, "true") || strings.EqualFold(envVar, "yes")
}
