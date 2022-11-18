package test

import (
	"errors"
	"fmt"
	"time"

	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
)

// CheckFn The prototype function for any check function without return value
type CheckFn func(env envUtil.Env) error

// CheckFnNew The prototype function for any check function with return value
type CheckFnNew[T any] func(env envUtil.Env) (T, error)

// WaitForCheckFuncOrTimeout Call this to ensure we wait until a check is reached, or time out
func WaitForCheckFuncOrTimeout(env envUtil.Env, checkFunc CheckFn) error {
	_, err := WaitForCheckFuncWithValueOrTimeout(env, func(env envUtil.Env) (any, error) { return nil, checkFunc(env) })
	return err
}

// WaitForCheckFuncWithValueOrTimeout Call this to ensure we wait until a check is reached, or time out, returning a value
func WaitForCheckFuncWithValueOrTimeout[T any](env envUtil.Env, checkFunc CheckFnNew[T]) (T, error) {
	timeout := env.GetTimeoutOfTest()
	sleepIntervalBetweenCheckFunc := env.GetSleepIntervalBetweenCheckFunc()
	firstSleepBetweenCheckFunc := time.Second
	var accumulatedWait time.Duration

	for {
		startTime := time.Now()
		obj, checkFuncErr := checkFunc(env)
		if checkFuncErr == nil {
			return obj, nil
		}

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
