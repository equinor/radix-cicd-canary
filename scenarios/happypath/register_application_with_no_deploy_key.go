package happypath

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	models "github.com/equinor/radix-cicd-canary-golang/scenarios/happypath/models"
	"github.com/equinor/radix-cicd-canary-golang/scenarios/utils"
	log "github.com/sirupsen/logrus"
)

func registerApplicationWithNoDeployKey() string {
	const (
		path     = "/api/v1/applications"
		method   = "POST"
		testName = "RegisterApplicationWithNoDeployKey"
	)

	log.Infof("Sending HTTP POST request...")

	parameters := models.ApplicationRegistration{
		Name:         app1Name,
		Repository:   app1Repository,
		SharedSecret: app1SharedSecret,
	}

	req := utils.CreateHTTPRequest(path, method, parameters)
	client := http.DefaultClient

	resp, err := client.Do(req)

	if err != nil {
		addTestError(testName)
		log.Errorf("HTTP POST error: %v", err)
	} else {
		ok := checkResponse(resp)
		if ok {
			addTestSuccess(testName)
			log.Infof("Response: %s", resp.Status)
		} else {
			addTestError(testName)
			log.Errorf("Error response code: %v", resp.StatusCode)
		}
	}

	return testName
}

func checkResponse(resp *http.Response) bool {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Error reading response body: %v", err)
		return false
	}
	appRegistration := models.ApplicationRegistration{}
	err = json.Unmarshal(body, &appRegistration)
	if err != nil {
		log.Errorf("Error unmarshalling response body: %v", err)
		return false
	}
	if resp.StatusCode == 200 && appRegistration.PublicKey != "" {
		return true
	}
	return false
}
