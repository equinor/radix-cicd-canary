package happypath

func registerApplicationWithNoDeployKey() string {
	const (
		path     = "/api/v1/applications"
		method   = "POST"
		testName = "RegisterApplicationWithNoDeployKey"
	)

	return testName
}
