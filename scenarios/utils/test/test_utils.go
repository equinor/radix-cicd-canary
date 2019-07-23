package test

import "time"

// CheckFn The prototype function for any check function
type CheckFn func(args []string) (bool, interface{})

const (
	defaultTimeout                       = 1200000
	defaultSleepIntervalBetweenCheckFunc = 5
)

// WaitForCheckFunc Call this to ensure we wait until a check is reached, or time out
func WaitForCheckFunc(checkFunc CheckFn) (bool, interface{}) {
	return waitForCheckFuncOrTimeout(defaultTimeout, defaultSleepIntervalBetweenCheckFunc, checkFunc, []string{})
}

// WaitForCheckFuncWithArguments Function to pass arguments to check function
func WaitForCheckFuncWithArguments(checkFunc CheckFn, args []string) (bool, interface{}) {
	return waitForCheckFuncOrTimeout(defaultTimeout, defaultSleepIntervalBetweenCheckFunc, checkFunc, args)
}

func waitForCheckFuncOrTimeout(timeout, sleepIntervalBetweenCheckFunc int, checkFunc CheckFn, args []string) (bool, interface{}) {
	accumulatedWait := 0

	for {
		startTime := time.Now()
		success, obj := checkFunc(args)
		if success {
			return true, obj
		}

		// should accumulatedWait include sleep?
		if sleepIntervalBetweenCheckFunc > 0 {
			time.Sleep(defaultSleepIntervalBetweenCheckFunc * time.Second)
		}

		waitPeriod := time.Now().Sub(startTime)

		if timeout > 0 {
			accumulatedWait = accumulatedWait + (int(waitPeriod.Seconds()) * 1000)
			if accumulatedWait > timeout {
				return false, nil
			}
		}

	}
}
