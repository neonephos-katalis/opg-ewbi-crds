# Makefile for opg-ewbi-crds (multi-component)

SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

.PHONY: all
all: generate manifests

##@ General

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: manifests
manifests: controller-gen ## Generate CRD YAML manifests for all components.
	@for comp_dir in apis/*; do \
		if [ -d "$$comp_dir" ]; then \
			comp=$$(basename "$$comp_dir"); \
			echo "==> Generating CRDs for component: $$comp"; \
			mkdir -p "crds/$$comp/bases"; \
			$(CONTROLLER_GEN) crd paths="./apis/$$comp/..." output:crd:artifacts:config="crds/$$comp/bases"; \
			cat crds/$$comp/bases/*.yaml > "crds/$$comp/install.yaml"; \
		fi \
	done
	@echo "==> Aggregating all component manifests into install.yaml"
	cat crds/*/bases/*.yaml > install.yaml

.PHONY: generate
generate: controller-gen ## Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject.
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./apis/..."

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

##@ Deployment

.PHONY: install
install: manifests ## Install all component CRDs into the target cluster.
	kubectl apply -f install.yaml

.PHONY: uninstall
uninstall: ## Uninstall all component CRDs from the target cluster.
	kubectl delete -f install.yaml

.PHONY: install-operator
install-operator: manifests ## Install only operator CRDs.
	kubectl apply -f crds/operator/install.yaml

.PHONY: uninstall-operator
uninstall-operator: ## Uninstall only operator CRDs.
	kubectl delete -f crds/operator/install.yaml

##@ Dependencies

LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

CONTROLLER_GEN ?= $(LOCALBIN)/controller-gen
CONTROLLER_TOOLS_VERSION ?= v0.16.5

.PHONY: controller-gen
controller-gen: $(CONTROLLER_GEN) ## Download controller-gen locally if necessary.
$(CONTROLLER_GEN): $(LOCALBIN)
	$(call go-install-tool,$(CONTROLLER_GEN),sigs.k8s.io/controller-tools/cmd/controller-gen,$(CONTROLLER_TOOLS_VERSION))

define go-install-tool
@[ -f $(1) ] || { \
set -e; \
package=$(2)@$(3) ;\
echo "Downloading $${package}" ;\
GOBIN=$(LOCALBIN) go install $${package} ;\
mv "$$(echo "$(1)" | sed "s/-$(3)$$//")" $(1) ;\
}
endef
