package test

import (
	"context"
	"time"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/rs/zerolog/log"
	"k8s.io/apimachinery/pkg/util/wait"
)

// CheckFn The prototype function for any check function without return value
type CheckFn func(cfg config.Config, ctx context.Context) error

// CheckFnNew The prototype function for any check function with return value
type CheckFnNew[T any] func(cfg config.Config, ctx context.Context) (T, error)

// WaitForCheckFuncOrTimeout Call this to ensure we wait until a check is reached, or time out
func WaitForCheckFuncOrTimeout(ctx context.Context, cfg config.Config, checkFunc CheckFn) error {
	_, err := WaitForCheckFuncWithValueOrTimeout(ctx, cfg, func(cfg config.Config, ctx context.Context) (any, error) { return nil, checkFunc(cfg, ctx) })
	return err
}

// WaitForCheckFuncWithValueOrTimeout Call this to ensure we wait until a check is reached, or time out, returning a value
func WaitForCheckFuncWithValueOrTimeout[T any](ctx context.Context, cfg config.Config, checkFunc CheckFnNew[T]) (T, error) {
	timeout := cfg.GetTimeoutOfTest()
	sleepInterval := cfg.GetSleepIntervalBetweenCheckFunc()
	timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var result T
	time.Sleep(1 * time.Second) // Wait for operator to pick up any changes HACK
	err := wait.PollUntilContextCancel(timeoutCtx, sleepInterval, true, func(ctx context.Context) (done bool, err error) {
		tmp, err := checkFunc(cfg, ctx)
		if err != nil {
			log.Ctx(ctx).Debug().Msgf("check function fails in WaitForCheck: %v", err)
			return false, nil
		}

		result = tmp
		return true, nil
	})

	return result, err
}
