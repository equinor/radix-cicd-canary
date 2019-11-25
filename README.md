# Radix CI/CD Canary

This application is an automated end-to-end test tool to be run continuously in a [Radix](https://www.radix.equinor.com) cluster to verify that the most important functionalities are behaving as expected. This document is for Radix developers, or anyone interested in poking around.

Radix CI/CD Canary is not deployed as a standard Radix application (it requires custom setup not provided by the platform), but rather as a custom Kubernetes deployment through a Helm chart.

The application is implemented in [Go](https://golang.org/). It provides metrics to the Radix [external monitoring solution](https://github.com/equinor/radix-monitoring/tree/master/cluster-external-monitoring) via [Prometheus](https://prometheus.io/). It relies on being able to impersonate users (test users and groups are defined in the Helm chart), and it interacts with the [Radix API](https://github.com/equinor/radix-api/) and the [Radix GitHub Webhook](https://github.com/equinor/radix-github-webhook) in the cluster it runs.

![pic](diagrams/radix-cicd-canary.png)

Currently, there are two scenarios (or suites) implemented, named `Happy path` and `NSP`.

The `Happy path` suite contains the following tests.
1. Register application
1. Register application with no deploy key
1. List applications
1. Build application
1. Set secret
1. Check alias responding
1. Check private image hub
1. Check access to application user should not be able to access
1. Promote deployment to other environment
1. Promote deployment to same environment
1. Checks that access can be locked down by updating AD group
1. Delete applications

The `NSP` suite contains the following tests.
1. Reach ingress
1. Reach service in different namespace

## Deploying

### From a local machine

NOTE: The following only applies to development. Proper installation is done through installing base components and Flux Helmrelease. The tests can be deployed to the cluster through a Helm chart, but first build the docker file (default it will push to radixdev. With ENVIRONMENT=prod it will push to radixprod):

```bash
make deploy-via-helm ENVIRONMENT=<dev|prod> CLUSTER_FQDN=<clustername>.<clustertype>.radix.equinor.com

# Example:
make deploy-via-helm ENVIRONMENT=dev CLUSTER_FQDN=weekly-27.dev.radix.equinor.com
```

### In a cluster

The application is installed by the `install_base_components.sh` script (in the [radix-platform repository](https://github.com/equinor/radix-platform/tree/master/scripts)) that is typically run when a new cluster is created. Before running the script, make sure that the docker file has been built and pushed to ACR:

```bash
make build-push ENVIRONMENT=<dev|prod>
```

And make sure that the Helm chart is pushed to ACR:

```bash
cd charts/radix-cicd-canary
helm package .
az acr helm push --name radixdev <tgz file>
az acr helm push --name radixprod <tgz file>
```

## Debugging

The application can be run locally for debugging purposes, but it will still interact with `radix-api` and `radix-github-webhook` in a cluster. A config map named `radix-cicd-canary` should be created in the cluster; its format can be found at `charts/templates/config.yaml`.

### Entire application

The tests can be debugged in their entirety by setting the `BEARER_TOKEN` value in the `launch.json` file, and then running debug from VSCode (F5).

### Unit tests

Unit tests can be debugged individually by setting the `BEARER_TOKEN` value in the `env_utils.go` file, and then running debug on each unit test.
