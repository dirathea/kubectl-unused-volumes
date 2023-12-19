export GO111MODULE=on

.PHONY: test
test: ## test
	go test ./pkg/... ./cmd/... -coverprofile cover.out

.PHONY: bin
bin: fmt vet ## build
	go build -o bin/unused-volumes github.com/ava-labs/kubectl-unused-volumes/cmd/plugin

.PHONY: fmt
fmt: ## fmt
	go fmt ./pkg/... ./cmd/...

.PHONY: vet
vet: ## vet
	go vet ./pkg/... ./cmd/...

.PHONY: setup
setup: ## setup
	make -C setup

.PHONY: snapshot
snapshot: fmt vet ## snapshot
	goreleaser build --snapshot --rm-dist

.PHONY: release
release: fmt vet ## Release
	goreleaser

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
