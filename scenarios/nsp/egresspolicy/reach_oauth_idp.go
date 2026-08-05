package egresspolicy

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/equinor/radix-cicd-canary/metrics"
	nspMetrics "github.com/equinor/radix-cicd-canary/metrics/scenarios/nsp"
	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/equinor/radix-common/utils/slice"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// ReachOauthIdp tests that IDP endpoint can be reached from Oauth Aux pod
func ReachOauthIdp(ctx context.Context, cfg config.Config) error {
	baseUrl, err := url.Parse(cfg.GetNetworkPolicyCanaryUrl("oauthdenyall"))
	if err != nil {
		return fmt.Errorf("failed to parse base URL: %w", err)
	}
	client := http.Client{
		Timeout: 15 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Get(baseUrl.String())
	if err != nil {
		return fmt.Errorf("failed initial request: %w", err)
	}

	loc, err := resp.Location()
	if err != nil {
		return fmt.Errorf("failed to get location from response: %w", err)
	}

	state := loc.Query().Get("state")
	if len(state) == 0 {
		return errors.New("location does not contain state in location")
	}

	csrfCookie, hasCsrfCookie := slice.FindFirst(resp.Cookies(), func(c *http.Cookie) bool { return c.Name == "_oauth2_proxy_csrf" })
	if !hasCsrfCookie {
		return errors.New("response does not CSRF cookie")
	}

	callbackUrl := baseUrl.JoinPath("oauth2", "callback")
	callbackQuery := callbackUrl.Query()
	callbackQuery.Add("code", "fake-code")
	callbackQuery.Add("state", state)
	callbackUrl.RawQuery = callbackQuery.Encode()

	req, err := http.NewRequest("GET", callbackUrl.String(), nil)
	if err != nil {
		panic(err)
	}
	req.AddCookie(csrfCookie)

	_, err = client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to get response from %s within %v, which likely means oauth pod could not connect to IDP which should be allowed by nsp: %w", callbackUrl.Path, client.Timeout, err)
	}

	return nil
}

// ReachOauthIdpSuccess is a function after a call to ReachOauthIdp succeeds
func ReachOauthIdpSuccess(ctx context.Context, testName string) {
	nspMetrics.AddOauthIdpReachable()
	metrics.AddTestOne(testName, nspMetrics.Success)
	metrics.AddTestZero(testName, nspMetrics.Errors)
	log.Ctx(ctx).Info().Msg("Test: SUCCESS")
}

// ReachOauthIdpFail is a function after a call to ReachOauthIdp failed
func ReachOauthIdpFail(ctx context.Context, testName string) {
	nspMetrics.AddOauthIdpUnreachable()
	metrics.AddTestZero(testName, nspMetrics.Success)
	metrics.AddTestOne(testName, nspMetrics.Errors)
	log.Ctx(ctx).Info().Msg("Test: FAIL")
}
