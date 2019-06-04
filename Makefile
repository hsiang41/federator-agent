
# Image URL to use all building/pushing image targets
IMG ?= federatorai-agent:latest

.PHONY: all test transmitter
all: test federatorai-agent

# Run tests
test: generate fmt vet
	go test ./cmd/... -coverprofile cover.out

# Build transmitter binary
federatorai-agent: generate fmt vet
	# go build -ldflags "-X main.VERSION=`git rev-parse --abbrev-ref HEAD`-`git rev-parse --short HEAD``git diff --quiet || echo '-dirty'` -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`'" -o transmitter/transmitter github.com/containers-ai/federatorai-agent/cmd
	go build -o transmitter/transmitter github.com/containers-ai/federatorai-agent/cmd

.PHONY: run

# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet
	go run ./cmd/transmitter.go run

.PHONY: fmt vet generate docker-build docker-push

# Run go fmt against code
fmt:
	go fmt ./cmd/...

# Run go build library
lib:
    go build -buildmode=plugin -a -o ./lib/inputlib/datapipe.so github.com/containers-ai/federatorai-agent/pkg/inputlib/alameda_datapipe
    go build -buildmode=plugin -a -o ./lib/outputlib/datapipe_writer.so github.com/containers-ai/federatorai-agent/pkg/outputlib/alameda_datapipe

# Run go vet against code
vet:
	go vet ./cmd/...

# Generate code
generate:
	go generate ./cmd/...

# Build the docker image
docker-build: test
	docker build . -t ${IMG} -f Dockerfile