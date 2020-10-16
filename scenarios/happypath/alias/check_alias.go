package alias

import (
	fmt "fmt"
	"strings"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/application"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/test"
	log "github.com/sirupsen/logrus"
)

var logger *log.Entry

// DefaultResponding Checks if default alias of application is responding
func DefaultResponding(env envUtil.Env, suiteName string) (bool, error) {
	logger = log.WithFields(log.Fields{"Suite": suiteName})

	ok, publicDomainName := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return application.TryGetPublicDomainName(env, config.App2Name, config.App2EnvironmentName, config.App2Component1Name)
	})

	if !ok {
		return false, fmt.Errorf("Public domain name of alias is empty")
	}

	ok, canonicalDomainName := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		return application.TryGetCanonicalDomainName(env, config.App2Name, config.App2EnvironmentName, config.App2Component1Name)
	})

	if !ok {
		return false, fmt.Errorf("Canonical domain name of alias is empty")
	}

	if application.IsRunningInActiveCluster(publicDomainName.(string), canonicalDomainName.(string)) {
		ok, _ := test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
			return application.IsAliasDefined(env, config.App2Name)
		})

		if !ok {
			return false, fmt.Errorf("Public alias is not defined")
		}
	}

	ok, _ = test.WaitForCheckFuncOrTimeout(env, func(env envUtil.Env) (bool, interface{}) {
		schema := "https"
		return application.AreResponding(env, getUrl(schema, canonicalDomainName.(string)), getUrl(schema, publicDomainName.(string)))
	})
	return ok, nil
}

func getUrl(schema string, domainName string) string {
	if strings.HasPrefix("http://", domainName) || strings.HasPrefix("https://", domainName) {
		return domainName
	}
	return fmt.Sprintf("%s://%s", schema, domainName)
}
