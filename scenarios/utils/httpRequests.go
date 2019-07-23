package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// CreateRequestForPath setup correct header for running tests
func CreateRequestForPath(apiPath, method string, parameters interface{}) *http.Request {
	var reader io.Reader
	endpoint := fmt.Sprintf("%s%s", GetRadixAPIURL(), apiPath)
	if parameters != nil {
		payload, _ := json.Marshal(parameters)
		reader = bytes.NewReader(payload)
	}
	req, _ := http.NewRequest(method, endpoint, reader)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", GetBearerToken()))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Impersonate-User", GetImpersonateUser())
	req.Header.Add("Impersonate-Group", GetImpersonateGroup())

	return req
}

// CreateRequest setup correct header for running tests
func CreateRequest(url, method string, parameters interface{}) *http.Request {
	var reader io.Reader
	if parameters != nil {
		payload, _ := json.Marshal(parameters)
		reader = bytes.NewReader(payload)
	}
	req, _ := http.NewRequest(method, url, reader)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", GetBearerToken()))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Impersonate-User", GetImpersonateUser())
	req.Header.Add("Impersonate-Group", GetImpersonateGroup())

	return req
}

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

// GetPrivateKeyBase64 get private deploy key from environment variable
func GetPrivateKeyBase64() string {
	data, _ := base64.StdEncoding.DecodeString(os.Getenv("PRIVATE_KEY_BASE64"))
	return string(data)
}
