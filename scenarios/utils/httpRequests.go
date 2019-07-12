package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// CreateHTTPRequest setup correct header for running tests
func CreateHTTPRequest(apiPath, method string, parameters interface{}) *http.Request {
	var reader io.Reader
	endpoint := fmt.Sprintf("%s%s", getRadixAPIURL(), apiPath)
	if parameters != nil {
		payload, _ := json.Marshal(parameters)
		reader = bytes.NewReader(payload)
	}
	req, _ := http.NewRequest(method, endpoint, reader)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", getBearerToken()))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Impersonate-User", getImpersonateUser())
	req.Header.Add("Impersonate-Group", getImpersonateGroup())

	return req
}

func getBearerToken() string {
	token, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/token")
	if err != nil {
		return os.Getenv("BEARER_TOKEN")
	}
	return string(token)
}

func getImpersonateUser() string {
	return os.Getenv("IMPERSONATE_USER")
}

func getImpersonateGroup() string {
	return os.Getenv("IMPERSONATE_GROUP")
}

func getRadixAPIURL() string {
	return os.Getenv("RADIX_API_URL")
}

// GetPublicKey get public deploy key from environment variable
func GetPublicKey() string {
	return os.Getenv("PUBLIC_KEY")
}

// GetPrivateKeyBase64 get private deploy key from environment variable
func GetPrivateKeyBase64() string {
	return os.Getenv("PRIVATE_KEY_BASE64")
}
