package defaults

// TODO: Add test for build-deploy for a specific commitId.
// Before user story https://dev.azure.com/Equinor/Radix/_workitems/edit/106471, the AppXCommitId didn't have any effect on the build-deploy jobs
// To keep consistent behavioru between dev and prod clusters in a transition period until the user story is release to prod
// we have to use "" as commit Id.
// When functionality in the user story is released to all clusters, we must add a test to verify that the commit id is actually used by the pipeline job(s)

const (
	App1Name       = "canarycicd-test1"
	App1Repository = "https://github.com/equinor/radix-canarycicd-test-1"
	// file deepcode ignore HardcodedPassword: This is a fake shared secret
	App1SharedSecret                 = "AnySharedSecret"
	App1Creator                      = "a-user@equinor.com"
	App1ConfigBranch                 = "master"
	App1ConfigurationItem            = "2b0781a7db131784551ea1ea4b9619c9"
	App2Name                         = "canarycicd-test2"
	App2Creator                      = "a-user@equinor.com"
	App2BranchToBuildFrom            = "master"
	App2CommitID                     = "" //"580f3ae3c0c23503a7cd02a8b60a5dec49279d0f"
	App2Repository                   = "https://github.com/equinor/radix-canarycicd-test-2"
	App2SSHRepository                = "git@github.com:equinor/radix-canarycicd-test-2.git"
	App2SharedSecret                 = "a sportless strobic spinach"
	App2EnvironmentName              = "prod"
	App2Component1Name               = "app"
	App2Component2Name               = "redis"
	App2ComponentPrivateImageHubName = "from-private-image-hub"
	App2BuildSecretName              = "SECRET_1"
	App2SecretName                   = "DB_PASS"
	App2SecretValue                  = "Test"
	App2ConfigBranch                 = "master"
	App2ConfigurationItem            = "2b0781a7db131784551ea1ea4b9619c9"
	App3Name                         = "canarycicd-test3"
	App3Creator                      = "a-user@equinor.com"
	App3BranchToBuildFrom            = "master"
	App3CommitID                     = "" // "580f3ae3c0c23503a7cd02a8b60a5dec49279d0f"
	App3Repository                   = "https://github.com/equinor/radix-canarycicd-test-3"
	App3SSHRepository                = "git@github.com:equinor/radix-canarycicd-test-3.git"
	App3SharedSecret                 = "a sportless strobic spinach"
	App3EnvironmentName              = "prod"
	App3Component1Name               = "client"
	App3ConfigBranch                 = "master"
	App3ConfigurationItem            = "2b0781a7db131784551ea1ea4b9619c9"
	App4Name                         = "canarycicd-test4"
	App4Repository                   = "https://github.com/equinor/radix-canarycicd-test-4"
	App4SSHRepository                = "git@github.com:equinor/radix-canarycicd-test-4.git"
	App4CommitID                     = "" //"02fdd66c44bad6dcbb887186e1a2a2bc9ee075f4"
	App4NewCommitID                  = "" //"97f71362819b3505087dfe56b5624f88aeebf6cb"
	App4SharedSecret                 = "AnySharedSecret"
	App4Creator                      = "a-user@equinor.com"
	App4ConfigBranch                 = "main"
	App4ConfigurationItem            = "2b0781a7db131784551ea1ea4b9619c9"
	App4NewConfigBranch              = "config"
	RestrictedApplicationName        = "radix-api"
)
