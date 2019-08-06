package test

import (
	"time"

	"github.com/equinor/radix-cicd-canary/scenarios/utils/env"
)

// CheckFn The prototype function for any check function
type CheckFn func(args []string) (bool, interface{})

// WaitForCheckFunc Call this to ensure we wait until a check is reached, or time out
func WaitForCheckFunc(checkFunc CheckFn) (bool, interface{}) {
	return waitForCheckFuncOrTimeout(checkFunc, []string{})
}

// WaitForCheckFuncWithArguments Function to pass arguments to check function
func WaitForCheckFuncWithArguments(checkFunc CheckFn, args []string) (bool, interface{}) {
	return waitForCheckFuncOrTimeout(checkFunc, args)
}

func waitForCheckFuncOrTimeout(checkFunc CheckFn, args []string) (bool, interface{}) {
	timeout := env.TimeoutOfTest()
	sleepIntervalBetweenCheckFunc := env.GetSleepIntervalBetweenCheckFunc()

	var accumulatedWait time.Duration

	for {
		startTime := time.Now()
		success, obj := checkFunc(args)
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
