package test

import (
	"errors"
	"time"

	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
)

// CheckFn The prototype function for any check function
type CheckFn[T any] func(env envUtil.Env, args []string) (T, error)

// CheckFnNew The prototype function for any check function
type CheckFnNew[T any] func(env envUtil.Env) (T, error)

// WaitForCheckFuncOrTimeout Call this to ensure we wait until a check is reached, or time out
func WaitForCheckFuncOrTimeout[T any](env envUtil.Env, checkFunc CheckFnNew[T]) (T, error) {
	timeout := env.GetTimeoutOfTest()
	sleepIntervalBetweenCheckFunc := env.GetSleepIntervalBetweenCheckFunc()
	firstSleepBetweenCheckFunc := time.Second
	var accumulatedWait time.Duration

	for {
		startTime := time.Now()
		obj, err := checkFunc(env)
		if err == nil {
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
				return nil, errors.New("timeout exceeded")
			}
		}

	}
}
