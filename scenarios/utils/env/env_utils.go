package env

import (
	"encoding/base64"
	"io/ioutil"
	"os"
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

// GetPrivateKey get private deploy key from environment variable
func GetPrivateKey() string {
	data, _ := base64.StdEncoding.DecodeString(os.Getenv("PRIVATE_KEY_BASE64"))
	return string(data)
}
