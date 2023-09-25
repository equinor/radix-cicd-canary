package config

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	kubeUtils "github.com/equinor/radix-cicd-canary/scenarios/utils/kubernetes"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/tokensource"
	log "github.com/sirupsen/logrus"
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
	privateImageHubPasswordConfig             = "privateImageHubPassword"
	networkPolicyCanaryPasswordConfig         = "networkPolicyCanaryPassword"
	networkPolicyCanaryAppNameConfig          = "networkPolicyCanaryAppName"
	networkPolicyCanaryJobComponentNameConfig = "networkPolicyCanaryJobComponentName"
	appAdminGroupConfig                       = "appAdminGroup"
	appReaderGroupsConfig                     = "appReaderGroup"
	envVarSuiteList                           = "SUITE_LIST"
	envVarIsBlacklist                         = "SUITE_LIST_IS_BLACKLIST"
	envVarLogLevel                            = "LOG_LEVEL"
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
		log.Fatalf("error reading config map: %v", err)
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
		GetNSPLongSleepInterval(),
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
func (cfg *Config) GetImpersonateGroups() []string {
	return []string{cfg.impersonateGroup, cfg.GetAppAdminGroup()}
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
func (cfg *Config) GetLogLevel() log.Level {
	lvl, err := log.ParseLevel(os.Getenv(envVarLogLevel))
	if err != nil {
		return log.InfoLevel
	}
	return lvl
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
			return string(token), err
		})
	} else if envToken := os.Getenv("BEARER_TOKEN"); len(envToken) > 0 {
		ts = tokensource.FromJwtCallback(func() (string, error) {
			return envToken, nil
		})
	}

	if ts == nil {
		panic(errors.New("failed to create TokenSource from service account token or environmnt variable"))
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
	os.Setenv("BEARER_TOKEN", "eyJhbGciOiJSUzI1NiIsImtpZCI6IjBuSzdVN2xtZVZSRlpKM20xVnhYdEJvZk56RFRrVVJQTjczdjFnTWxfUFUifQ.eyJhdWQiOlsiaHR0cHM6Ly9ub3J0aGV1cm9wZS5vaWMucHJvZC1ha3MuYXp1cmUuY29tLzNhYTRhMjM1LWI2ZTItNDhkNS05MTk1LTdmY2YwNWI0NTliMC82ODA5OWQ5NC00ODE1LTQ4MDgtOWU0MS1lOTZkMGRhNjEwZWEvIiwiaHR0cHM6Ly93ZWVrbHktMzktY2x1c3RlcnMtMTZlZGU0LWo2bDg3YTdwLmhjcC5ub3J0aGV1cm9wZS5hem1rOHMuaW8iLCJcIndlZWtseS0zOS1jbHVzdGVycy0xNmVkZTQtajZsODdhN3AuaGNwLm5vcnRoZXVyb3BlLmF6bWs4cy5pb1wiIl0sImV4cCI6MTY5NTY0NTY4OCwiaWF0IjoxNjk1NjQyMDg4LCJpc3MiOiJodHRwczovL25vcnRoZXVyb3BlLm9pYy5wcm9kLWFrcy5henVyZS5jb20vM2FhNGEyMzUtYjZlMi00OGQ1LTkxOTUtN2ZjZjA1YjQ1OWIwLzY4MDk5ZDk0LTQ4MTUtNDgwOC05ZTQxLWU5NmQwZGE2MTBlYS8iLCJrdWJlcm5ldGVzLmlvIjp7Im5hbWVzcGFjZSI6InJhZGl4LWNpY2QtY2FuYXJ5Iiwic2VydmljZWFjY291bnQiOnsibmFtZSI6InJhZGl4LWNpY2QtY2FuYXJ5IiwidWlkIjoiNmRjZTY5NjEtNGVjMS00MjdlLWJiYTAtMTNiNjJkYWVjN2JlIn19LCJuYmYiOjE2OTU2NDIwODgsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpyYWRpeC1jaWNkLWNhbmFyeTpyYWRpeC1jaWNkLWNhbmFyeSJ9.R46WnGwWkpcrwQrVs0QqnCqnk895drU3pb6Ng8pMiE3nlOzXBYz4ePqLVHsV_Ltiti2jBIHJC6z1RjHc1OjZqC2ePWIVVdfUG2M7G34dNVky80hfdNJN9ry3Dslno5_fsZariifN0kHMwSALfGoLfh62zTRGl5TUMKIT_eLKEoKvThvmzQ2WkM6heYSclqGF4JntOHFS_hsiW4yibaSSdtx4nMkSQb_WIwVpTAanJVPPkWDBa5740Y4HXQVgJx9xqBT6ePVWVo5qifXq2EEEK3dEYol2BFv49qQU9xeF3Vj6G-LBsYqWAsKnwWJ4uRhHoGMYtOl5xQdSbK1BEhTYQhqt0-YG08HJD1DWyzPOBJyA9mBh-xJoCvTIRAuFyCLfPspg8y-d28V0W85MPZ3rzlpYF3AIXGtcAE9xM_RMhWXBXloewQnG2WJFwkbQ1vkaCY6QYaA0SQ-Cl3JGMftOXGQbDQhZUMrFUFx8RmNE1LEBwsidQjrmBn48XvD8JQS4x-O3QglSCOlyN53VaLfX9-qDh6JBfe3oPZwByd2ZPlBWAzjRqWhic1f_rG7UBkbu5L0CTNJNlB7ztIHRctU6nRfJLu34k4Im4ZcGTN8wqoETbDy-t-hAHl3I68xaJWbvOV5Dsx5m0ZoQqVe5hDaMyfdw74ip5yfCbliU8CSrumU")
}

func getConfigFromMap(config string) string {
	configValue, found := configmap.Data[config]
	if !found {
		log.Fatalf("%s not found in configmap", config)
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

func useLocalRadixApi() bool {
	return envVarIsTrueOrYes(os.Getenv(envUseLocalRadixApi))
}

func useLocalGitHubWebHookApi() bool {
	return envVarIsTrueOrYes(os.Getenv(envUseLocalGitHubWebHookApi))
}

func envVarIsTrueOrYes(envVar string) bool {
	return strings.EqualFold(envVar, "true") || strings.EqualFold(envVar, "yes")
}
