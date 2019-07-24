package happypath

const (
	app1Name                  = "canarycicd-test1"
	app1Repository            = "https://github.com/equinor/radix-canarycicd-test-1"
	app1SharedSecret          = "AnySharedSecret"
	app2Name                  = "canarycicd-test2"
	app2BranchToBuildFrom     = "master"
	app2CommitID              = "580f3ae3c0c23503a7cd02a8b60a5dec49279d0f"
	app2Repository            = "https://github.com/equinor/radix-canarycicd-test-2"
	app2SSHRepository         = "git@github.com:equinor/radix-canarycicd-test-2.git"
	app2SharedSecret          = "a sportless strobic spinach"
	app2EnvironmentName       = "prod"
	app2Component1Name        = "app"
	app2Component2Name        = "redis"
	app2SecretName            = "DB_PASS"
	app2SecretValue           = "Test"
	restrictedApplicationName = "radix-api"
)
