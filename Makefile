ENVIRONMENT ?= dev
VERSION 	?= latest

BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
VAULT_NAME ?= radix-vault-$(ENVIRONMENT)
CLUSTER_FQDN ?= weekly-47.dev.radix.equinor.com
RADIX_API_PREFIX ?= server-radix-api-qa
RADIX_WEBHOOK_PREFIX ?= webhook-radix-github-webhook-qa

CONTAINER_REPO ?= radix$(ENVIRONMENT)
DOCKER_REGISTRY	?= $(CONTAINER_REPO).azurecr.io

generate-client:
	swagger generate client -t ./generated-client -f https://api.dev.radix.equinor.com/swaggerui/swagger.json -A radix-api

build:
	docker build -t radix-cicd-canary .

run:
	docker run -it --rm -p 5000:5000 radix-cicd-canary

build-push:
	az acr login --name $(CONTAINER_REPO)
	docker build -t $(DOCKER_REGISTRY)/radix-cicd-canary:$(BRANCH)-$(VERSION) .
	docker push $(DOCKER_REGISTRY)/radix-cicd-canary:$(BRANCH)-$(VERSION)

deploy-via-helm:
	make build-push

	az keyvault secret download \
		--vault-name $(VAULT_NAME) \
		--name radix-cicd-canary-values \
		--file radix-cicd-canary-values.yaml

	helm upgrade --install radix-cicd-canary \
	    ./charts/radix-cicd-canary/ \
		--namespace radix-cicd-canary \
		--set image.tag=$(BRANCH)-$(VERSION) \
		--set radixApiPrefix=$(RADIX_API_PREFIX) \
		--set radixWebhookPrefix=$(RADIX_WEBHOOK_PREFIX) \
		--set clusterFqdn=$(CLUSTER_FQDN) \
		-f radix-cicd-canary-values.yaml

	rm -f radix-cicd-canary-values.yaml