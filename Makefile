ENVIRONMENT ?= dev
VERSION 	?= latest

# BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
VAULT_NAME ?= radix-vault-$(ENVIRONMENT)

CONTAINER_REPO ?= radix$(ENVIRONMENT)
DOCKER_REGISTRY	?= $(CONTAINER_REPO).azurecr.io

build:
	docker build -t radix-cicd-canary-golang .

run:
	docker run -it --rm -p 5000:5000 radix-cicd-canary-golang

build-push:
	az acr login --name $(CONTAINER_REPO)
	docker build -t $(DOCKER_REGISTRY)/radix-cicd-canary-golang:$(BRANCH)-$(VERSION) .
	docker push $(DOCKER_REGISTRY)/radix-cicd-canary-golang:$(BRANCH)-$(VERSION)

deploy-via-helm:
	make build-push

	az keyvault secret download \
		--vault-name $(VAULT_NAME) \
		--name radix-cicd-canary-values \
		--file radix-cicd-canary-values.yaml

	helm upgrade --install radix-cicd-canary-golang \
	    ./charts/radix-cicd-canary-golang/ \
		--namespace radix-cicd-canary-golang \
		--set image.tag=$(BRANCH)-$(VERSION) \
		-f radix-cicd-canary-values.yaml

	rm -f radix-cicd-canary-values.yaml