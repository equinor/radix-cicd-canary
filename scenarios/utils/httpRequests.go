package utils

import (
	"fmt"
	"net/http"
	"os"
)

// CreateHTTPRequest setup correct header for running tests
func CreateHTTPRequest(endpoint, method string) *http.Request {
	req, _ := http.NewRequest(method, endpoint, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", getBearerToken()))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Impersonate-User", getImpersonateUser())
	req.Header.Add("Impersonate-Group", getImpersonateGroup())

	return req
}

func getBearerToken() string {
	return os.Getenv("BEARER_TOKEN")
}

func getImpersonateUser() string {
	return os.Getenv("IMPERSONATE_USER")
}

func getImpersonateGroup() string {
	return os.Getenv("IMPERSONATE_GROUP")
}
