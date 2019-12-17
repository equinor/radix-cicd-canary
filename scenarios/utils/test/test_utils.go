package test

import (
	"time"

	envUtil "github.com/equinor/radix-cicd-canary/scenarios/utils/env"
)

// CheckFn The prototype function for any check function
type CheckFn func(env envUtil.Env, args []string) (bool, interface{})

// CheckFnNew The prototype function for any check function
type CheckFnNew func(env envUtil.Env) (bool, interface{})

// WaitForCheckFunc Call this to ensure we wait until a check is reached, or time out
func WaitForCheckFunc(env envUtil.Env, checkFunc CheckFn) (bool, interface{}) {
	fn := func(env envUtil.Env) (bool, interface{}) { return checkFunc(env, []string{}) }
	return WaitForCheckFuncOrTimeout(env, fn)
}

// WaitForCheckFuncOrTimeout Call this to ensure we wait until a check is reached, or time out
func WaitForCheckFuncOrTimeout(env envUtil.Env, checkFunc CheckFnNew) (bool, interface{}) {
	timeout := env.GetTimeoutOfTest()
	sleepIntervalBetweenCheckFunc := env.GetSleepIntervalBetweenCheckFunc()

	var accumulatedWait time.Duration

	for {
		startTime := time.Now()
		success, obj := checkFunc(env)
		if success {
			return true, obj
		}

		// should accumulatedWait include sleep?
		if sleepIntervalBetweenCheckFunc > 0 {
			time.Sleep(sleepIntervalBetweenCheckFunc)
		}

		waitPeriod := time.Now().Sub(startTime)

		if timeout > 0 {
			accumulatedWait = accumulatedWait + waitPeriod
			if accumulatedWait > timeout {
				return false, nil
			}
		}

	}
}
