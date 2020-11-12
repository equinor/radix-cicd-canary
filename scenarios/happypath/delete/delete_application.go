package delete

import (
	"fmt"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

// Applications Tests that we are able to delete applications
func Applications(env env.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	isAllSuccess := true
	var allErrors error
	var errorMessages string
	success, err := application.Delete(env, config.App1Name)
	if !success {
		isAllSuccess = false
		errorMessages += fmt.Sprintf("%s\n", err.Error())
	}

	success, err = application.Delete(env, config.App2Name)
	if !success {
		isAllSuccess = false
		errorMessages += fmt.Sprintf("%s\n", err.Error())
	}

	success, err = application.Delete(env, config.App4Name)
	if !success {
		isAllSuccess = false
		errorMessages += fmt.Sprintf("%s\n", err.Error())
	}

	if !isAllSuccess {
		allErrors = fmt.Errorf("Errors:\n%s", errorMessages)
	}

	return isAllSuccess, allErrors
}
