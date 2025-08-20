![Build Status](https://github.com/equinor/radix-cicd-canary/workflows/radix-cicd-canary-build/badge.svg)  [![SCM Compliance](https://scm-compliance-api.radix.equinor.com/repos/equinor/radix-cicd-canary/badge)](https://developer.equinor.com/governance/scm-policy/)  
# Radix CI/CD Canary

This application is an automated end-to-end test tool to be run continuously in a [Radix](https://www.radix.equinor.com) cluster to verify that the most important functionalities are behaving as expected. This document is for Radix developers, or anyone interested in poking around.

Radix CI/CD Canary is deployed to cluster through a Helm release using [Flux Operator](https://github.com/weaveworks/flux), whenever a new image is pushed to the container registry for the corresponding branch, or a change has been made to the Helm chart. Build and Push to container registry is done using Github actions. 

The application is implemented in [Go](https://golang.org/). It provides metrics to the Radix [external monitoring solution](https://github.com/equinor/radix-monitoring/tree/master/cluster-external-monitoring) via [Prometheus](https://prometheus.io/). It relies on being able to impersonate users (test users and groups are defined in the Helm chart), and it interacts with the [Radix API](https://github.com/equinor/radix-api/) and the [Radix GitHub Webhook](https://github.com/equinor/radix-github-webhook) in the cluster it runs.

![pic](diagrams/radix-cicd-canary.png)

Currently, there are following scenarios (or suites) implemented, named `Happy path`, `NSP`, `NSP-Long` and `Deploy only`.

The `Happy path` suite contains the following tests.

1. Register application
2. Register application with no deploy key
3. Register application with main as config branch
4. List applications
5. Set build secrets
6. Build application
7. Set secret
8. Check private image hub
9. Check alias responding
10. Check access to application user should not be able to access
11. Promote deployment to other environment
12. Promote deployment to same environment
13. Checks that access can be locked down by updating AD group
14. Checks that machine user can be created and get proper access
15. Checks that radixconfig.yaml is read from correct config branch
16. Delete applications

The `NSP` (Network Security Policy) suite contains the following tests.

1. Reach ingress
2. Reach service in different namespace
3. Do DNS lookup toward public nameservers from [networkpolicy-canary](https://console.dev.radix.equinor.com/applications/radix-networkpolicy-canary).
4. Do DNS lookup toward internal K8s nameserver from networkpolicy-canary.
5. Get list of Radix jobs in networkpolicy-canary's namespace from networkpolicy-canary.
6. Test that http://login.microsoft.com/ can be reached from networkpolicy-canary.

The `NSP-Long` suite contains the following tests. The `NSP-Long` suite has longer test interval than the `NSP` suite.

1. Test scheduling of job batch from networkpolicy-canary.
2. Test whether [radix-canary-golang](https://github.com/equinor/radix-canary-golang) can be reached from networkpolicy-canary.
3. Test whether well known external websites can be reached from networkpolicy-canary.

The `Deploy only` suite contains the following tests.

1. Register application
1. Deploy application
1. Check private image hub func
1. Check alias responding
1. Delete applications


## Development Process

The `radix-cicd-canary` project follows a **trunk-based development** approach.

### 🔁 Workflow

- **External contributors** should:
  - Fork the repository
  - Create a feature branch in their fork

- **Maintainers** may create feature branches directly in the main repository.

### ✅ Merging Changes

All changes must be merged into the `master` branch using **pull requests** with **squash commits**.

The squash commit message must follow the [Conventional Commits](https://www.conventionalcommits.org/en/about/) specification.


## Release Process

Merging a pull request into `master` triggers the **Prepare release pull request** workflow.  
This workflow analyzes the commit messages to determine whether the version number should be bumped — and if so, whether it's a major, minor, or patch change.  

It then creates two pull requests:

- one for the new stable version (e.g. `1.2.3`), and  
- one for a pre-release version where `-rc.[number]` is appended (e.g. `1.2.3-rc.1`).

---

Merging either of these pull requests triggers the **Create releases and tags** workflow.  
This workflow reads the version stored in `version.txt`, creates a GitHub release, and tags it accordingly.

The new tag triggers the **Build and deploy Docker and Helm** workflow, which:

- builds and pushes a new container image and Helm chart to `ghcr.io`, and  
- uploads the Helm chart as an artifact to the corresponding GitHub release.



## Debugging

The application can be run locally for debugging purposes, but it will still interact with `radix-api` and `radix-github-webhook` in a cluster. A config map named `radix-cicd-canary` should exist in in the cluster, under the `radix-cicd-canary` namespace (i.e. `kubectl get configmap -n radix-cicd-canary -oyaml`); its format can be found at `charts/templates/config.yaml`. Normally, though, you don't need to do anything with this configmap. When debugging in a cluster it is wise to turn of the canary in the cluster. Do that by setting replica to zero for the deployment (i.e. `kubectl edit deploy -n radix-cicd-canary`). Also make sure you start a test from scratch by deleting the registration for the apps used in the tests `kubectl delete rr $(kubectl get rr -o custom-columns=':metadata.name' --no-headers | grep canarycicd-)`

### Entire application

The tests can be debugged in their entirety by setting the `BEARER_TOKEN` value in the `.env` file, and then running debug from VSCode (F5). You will most likely need to comment in the `os.Setenv("GODEBUG", "http2server=0,http2client=0")` line in the `main.go` to allow for a large token in the API.

### Unit tests

Unit tests can be debugged individually by setting the `BEARER_TOKEN` value in the `env_utils.go` file, and then running debug on each unit test. Note the unit tests, are not really unit tests, but an ability to test a single functionality. Make sure that all scenarios before the test has executed before you start debugging a single test.

### Custom configuration

By default `Info` and `Error` messages have being logged. This can be configured via environment variable `LOG_LEVEL` (pods need to be restarted after changes)
* `LOG_LEVEL=ERROR` - log only `Error` messages
* `LOG_LEVEL=INFO` or not set - log `Info` and `Error` messages
* `LOG_LEVEL=WARN` or not set - log `Info`, `Warning` and `Error` messages
* `LOG_LEVEL=DEBUG` - log `Debug`, `Warning`, `Info` and `Error` messages
* `PRETTY_PRINT=yes` - Print human readable text instead of json messages

By default, all suites are running. This can be configured with environment variables
* `SUITE_LIST` - list of suite names, separated by `:`
* `SUITE_LIST_IS_BLACKLIST`
  * `false`, `no` or not set - `SUITE_LIST` contains suites to be only running
  * `true` or `yes` - `SUITE_LIST` contains suites, which should be skipped

To debug locally with connecting to the local services - set following environment variables:
* `USE_LOCAL_GITHUB_WEBHOOK_API`
  * `false`, `no` or not set - connecting to in-cluster `radix-api`
  * `true` or `yes` - connecting to `radix-api`, running on `http://localhost:3001`
* `USE_LOCAL_RADIX_API`
  * `false`, `no` or not set - connecting to in-cluster `radix-api`
  * `true` or `yes` - connecting to `radix-api`, running on `http://localhost:3002`

## Contribution

Want to contribute? Read our [contributing guidelines](./CONTRIBUTING.md)

## Security

This is how we handle [security issues](./SECURITY.md)
