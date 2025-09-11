# Image URL to use all building/pushing image targets
VERSION=$(shell cat VERSION)
IMG ?= wego:${VERSION}


# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

.PHONY: all
all: build

##@ Development

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: swag
swag: ## Run swag init against code.
	swag init --parseDependency --parseInternal -g  ./main.go  -o ./docs -d ./cmd/api,./internal/httpservice

##@ Build

.PHONY: build
build: fmt vet swag ## Build manager binary.
	go build -o ./main cmd/api/main.go
build-linux-amd64: fmt swag
	GOOS=linux GOARCH=amd64 go build -o ./main cmd/api/main.go

.PHONY: run
run:  fmt vet ## Run a controller from your host.
	go run ./cmd/api/main.go

.PHONY: docker-build
docker-build: build-linux-amd64  ## Build docker image with the manager.
	docker build -f build/Dockerfile --build-arg GO_FLAGS=${GO_LD_FLAGS} -t ${IMG} . 
	docker save ${IMG} > build/images/wego.tar
