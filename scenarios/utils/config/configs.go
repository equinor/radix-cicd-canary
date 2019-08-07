package config

const (
	App1Name                  = "canarycicd-test1"
	App1Repository            = "https://github.com/equinor/radix-canarycicd-test-1"
	App1SharedSecret          = "AnySharedSecret"
	App2Name                  = "canarycicd-test2"
	App2BranchToBuildFrom     = "master"
	App2CommitID              = "580f3ae3c0c23503a7cd02a8b60a5dec49279d0f"
	App2Repository            = "https://github.com/equinor/radix-canarycicd-test-2"
	App2SSHRepository         = "git@github.com:equinor/radix-canarycicd-test-2.git"
	App2SharedSecret          = "a sportless strobic spinach"
	App2EnvironmentName       = "prod"
	App2Component1Name        = "app"
	App2Component2Name        = "redis"
	App2SecretName            = "DB_PASS"
	App2SecretValue           = "Test"
	RestrictedApplicationName = "radix-api"
)
