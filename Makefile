
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

binaries:
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.VERSION=`git rev-parse --abbrev-ref HEAD`-`git rev-parse --short HEAD``git diff --quiet || echo '-dirty'` -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`'" -a -o ./transmitter/transmitter github.com/containers-ai/federatorai-agent/cmd
	# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags \"-X main.VERSION=`git rev-parse --abbrev-ref HEAD`-`git rev-parse --short HEAD``git diff --quiet || echo '-dirty'` -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`'\" -buildmode=plugin -a -o ./lib/inputlib/datahub.so github.com/containers-ai/federatorai-agent/pkg/inputlib/alameda_datapipe/datapipe.go
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -buildmode=plugin -a -o ./lib/inputlib/datapipe.so github.com/containers-ai/federatorai-agent/pkg/inputlib/alameda_datapipe
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -buildmode=plugin -a -o ./lib/outputlib/datapipe_recommender.so github.com/containers-ai/federatorai-agent/pkg/outputlib/alameda_recommender

install_dir:
	mkdir -pv /etc/alameda/federatorai-agent /etc/alameda/federatorai-agent/inputlib /lib/inputlib /lib/outputlib /root/

install: install_dir
	cp -fv /go/src/github.com/containers-ai/federatorai-agent/etc/transmitter.toml /etc/alameda/federatorai-agent/transmitter.toml
	cp -fv /go/src/github.com/containers-ai/federatorai-agent/etc/inputlib/datapipe.toml /etc/alameda/federatorai-agent/inputlib/datapipe.toml
	cp -fv /go/src/github.com/containers-ai/federatorai-agent/etc/inputlib/datapipe.toml /etc/alameda/federatorai-agent/outputlib/datapipe.toml
	cp -fv /go/src/github.com/containers-ai/federatorai-agent/transmitter/transmitter /root/
	cp -fv /go/src/github.com/containers-ai/federatorai-agent/lib/inputlib/datapipe.so /lib/inputlib/datapipe.so
	cp -fv /go/src/github.com/containers-ai/federatorai-agent/lib/outputlib/datapipe_recommender.so /lib/outputlib/datapipe_recommender.so

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
	go build -buildmode=plugin -a -o ./lib/outputlib/datapipe_recommender.so github.com/containers-ai/federatorai-agent/pkg/outputlib/alameda_recommender

# Run go vet against code
vet:
	go vet ./cmd/...

# Generate code
generate:
	go generate ./cmd/...

# Build the docker image
docker-build:
	docker build . -t ${IMG} -f Dockerfile
