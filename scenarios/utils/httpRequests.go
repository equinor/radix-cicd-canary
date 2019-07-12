package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// CreateHTTPRequest setup correct header for running tests
func CreateHTTPRequest(apiPath, method string) *http.Request {
	endpoint := fmt.Sprintf("%s/%s", getRadixAPIURL(), apiPath)
	req, _ := http.NewRequest(method, endpoint, nil)
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
