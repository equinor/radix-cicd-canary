package happypath

const (
	newAppNoDeployKeyAppPath     = "/api/v1/applications"
	newAppNoDeployKeyAppMethod   = "POST"
	newAppNoDeployKeyAppTestName = "RegisterApplicationWithNoDeployKey"
)

func registerApplicationWithNoDeployKey() string {
	return newAppNoDeployKeyAppTestName
}
