package env

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

// GetBearerToken get bearer token either from token file or environment variable
func GetBearerToken() string {
	token, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/token")
	if err != nil {
		return os.Getenv("BEARER_TOKEN")
	}
	return string(token)
}

// GetImpersonateUser get impersonate user
func GetImpersonateUser() string {
	return os.Getenv("IMPERSONATE_USER")
}

// GetImpersonateGroup get impersonate group
func GetImpersonateGroup() string {
	return os.Getenv("IMPERSONATE_GROUP")
}

// GetRadixAPIURL get Radix API URL
func GetRadixAPIURL() string {
	return os.Getenv("RADIX_API_URL")
}

// GetWebhookURL get Radix API URL
func GetWebhookURL() string {
	return os.Getenv("RADIX_GITHUB_WEBHOOK_URL")
}

// GetPublicKey get public deploy key from environment variable
func GetPublicKey() string {
	return os.Getenv("PUBLIC_KEY")
}

// TimeoutOfTest Get the time it should take before a test should time out
func TimeoutOfTest() time.Duration {
	timeout, err := strconv.Atoi(os.Getenv("TIMEOUT_OF_TEST_SEC"))
	if err != nil {
		log.Fatalf("Could not read %s. Err: %v", "TIMEOUT_OF_TEST_SEC", err)
	}

	return time.Duration(timeout) * time.Second
}

// GetSleepIntervalBetweenCheckFunc Gets the sleep inteval between two checks
func GetSleepIntervalBetweenCheckFunc() time.Duration {
	sleepInterval, err := strconv.Atoi(os.Getenv("SLEEP_INTERVAL_BETWEEN_CHECK_SEC"))
	if err != nil {
		log.Fatalf("Could not read %s. Err: %v", "SLEEP_INTERVAL_BETWEEN_CHECK_SEC", err)
	}

	return time.Duration(sleepInterval) * time.Second
}

// GetSleepIntervalBetweenTestRuns Gets the sleep inteval between two test runs
func GetSleepIntervalBetweenTestRuns() time.Duration {
	sleepInterval, err := strconv.Atoi(os.Getenv("SLEEP_INTERVAL_BETWEEN_TEST_RUNS_SEC"))
	if err != nil {
		log.Fatalf("Could not read %s. Err: %v", "SLEEP_INTERVAL_BETWEEN_TEST_RUNS_SEC", err)
	}

	return time.Duration(sleepInterval) * time.Second
}

// GetPrivateKey get private deploy key from environment variable
func GetPrivateKey() string {
	data, _ := base64.StdEncoding.DecodeString(os.Getenv("PRIVATE_KEY_BASE64"))
	return string(data)
}
