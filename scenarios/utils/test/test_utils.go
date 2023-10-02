package test

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/config"
	"github.com/rs/zerolog/log"
)

// CheckFn The prototype function for any check function without return value
type CheckFn func(cfg config.Config, ctx context.Context) error

// CheckFnNew The prototype function for any check function with return value
type CheckFnNew[T any] func(cfg config.Config, ctx context.Context) (T, error)

// WaitForCheckFuncOrTimeout Call this to ensure we wait until a check is reached, or time out
func WaitForCheckFuncOrTimeout(cfg config.Config, checkFunc CheckFn, ctx context.Context) error {
	_, err := WaitForCheckFuncWithValueOrTimeout(cfg, func(cfg config.Config, ctx context.Context) (any, error) { return nil, checkFunc(cfg, ctx) }, ctx)
	return err
}

// WaitForCheckFuncWithValueOrTimeout Call this to ensure we wait until a check is reached, or time out, returning a value
func WaitForCheckFuncWithValueOrTimeout[T any](cfg config.Config, checkFunc CheckFnNew[T], ctx context.Context) (T, error) {
	timeout := cfg.GetTimeoutOfTest()
	sleepIntervalBetweenCheckFunc := cfg.GetSleepIntervalBetweenCheckFunc()
	firstSleepBetweenCheckFunc := time.Second
	var accumulatedWait time.Duration

	for {
		startTime := time.Now()
		obj, checkFuncErr := checkFunc(cfg, ctx)
		if checkFuncErr == nil {
			return obj, nil
		}
		log.Ctx(ctx).Debug().Err(checkFuncErr).Msg("check function fails in WaitForCheck")

		// should accumulatedWait include sleep?
		if sleepIntervalBetweenCheckFunc > 0 {
			sleepTime := sleepIntervalBetweenCheckFunc

			if accumulatedWait == 0 && firstSleepBetweenCheckFunc < sleepIntervalBetweenCheckFunc {
				sleepTime = firstSleepBetweenCheckFunc
			}

			time.Sleep(sleepTime)
		}

		waitPeriod := time.Since(startTime)

		if timeout > 0 {
			accumulatedWait = accumulatedWait + waitPeriod
			if accumulatedWait > timeout {
				message := "timeout exceeded"
				if checkFuncErr != nil {
					message = fmt.Sprintf("%s. Last error: %s", message, checkFuncErr.Error())
				}
				return obj, errors.New(message)
			}
		}

	}
}
