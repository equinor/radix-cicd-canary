ENVIRONMENT ?= dev
VERSION 	?= latest

BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
VAULT_NAME ?= radix-vault-$(ENVIRONMENT)

CLUSTER_FQDN ?= weekly-60.dev.radix.equinor.com

RADIX_API_PREFIX ?= server-radix-api-qa
RADIX_WEBHOOK_PREFIX ?= webhook-radix-github-webhook-qa

CONTAINER_REPO ?= radix$(ENVIRONMENT)
DOCKER_REGISTRY	?= $(CONTAINER_REPO).azurecr.io

.PHONY: lint
lint: bootstrap
	golangci-lint run --max-same-issues 0 --timeout 10m

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: generate-client
generate-client: SHELL:=/bin/bash
generate-client: bootstrap
	rm -Rf ./generated-client
	mkdir -p ./generated-client/radixapi
	mkdir -p ./generated-client/jobserver
	swagger generate client -t ./generated-client/radixapi -f https://api.dev.radix.equinor.com/swaggerui/swagger.json -A radixapi
	swagger generate client -t ./generated-client/jobserver -f https://raw.githubusercontent.com/equinor/radix-job-scheduler/main/swaggerui/html/swagger.json -A jobserver

build:
	docker build -t radix-cicd-canary:$(BRANCH)-$(VERSION) .

run:
	docker run -it --rm -p 5000:5000 radix-cicd-canary

.PHONY: deploy
deploy:
	az acr login --name $(CONTAINER_REPO)
	docker build -t $(DOCKER_REGISTRY)/radix-cicd-canary:$(BRANCH)-$(VERSION) .
	docker push $(DOCKER_REGISTRY)/radix-cicd-canary:$(BRANCH)-$(VERSION)

delete-dev-image:
	az acr repository delete --n radixdev  --image  radix-cicd-canary:$(BRANCH)-$(VERSION) --yes

delete-image-and-deploy:
	make delete-dev-image
	make deploy-via-helm

.PHONY: test
test:
	go test -cover `go list ./...`


.PHONY: generate
generate: tidy generate-client

.PHONY: verify-generate
verify-generate: generate
	git diff --exit-code

HAS_SWAGGER       := $(shell command -v swagger;)
HAS_GOLANGCI_LINT := $(shell command -v golangci-lint;)

bootstrap:
ifndef HAS_SWAGGER
	go install github.com/go-swagger/go-swagger/cmd/swagger@v0.31.0
endif
ifndef HAS_GOLANGCI_LINT
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.3
endif
