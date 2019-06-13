
# Image URL to use all building/pushing image targets
IMG ?= federatorai-agent:latest

SRC_DIR = $(shell pwd)
INSTALL_ROOT = $(SRC_DIR)/install_root
PRODUCT_ROOT = /opt/alameda/federatorai-agent
DEST_PREFIX = $(INSTALL_ROOT)$(PRODUCT_ROOT)
######################################################################

.PHONY: all test transmitter
all: test federatorai-agent

# Run tests
test: generate fmt vet
	go test ./cmd/... -coverprofile cover.out

# Build transmitter binary
federatorai-agent: generate fmt vet binaries lib
	# go build -ldflags "-X main.VERSION=`git rev-parse --abbrev-ref HEAD`-`git rev-parse --short HEAD``git diff --quiet || echo '-dirty'` -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`'" -o transmitter/transmitter github.com/containers-ai/federatorai-agent/cmd
	#go build -o transmitter/transmitter github.com/containers-ai/federatorai-agent/cmd

# Run go build library
lib:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -buildmode=plugin \
	  -a -o ./lib/inputlib/datapipe.so github.com/containers-ai/federatorai-agent/pkg/inputlib/alameda_datapipe
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -buildmode=plugin \
	  -a -o ./lib/outputlib/datapipe_recommender.so github.com/containers-ai/federatorai-agent/pkg/outputlib/alameda_recommender

binaries:
	GOOS=linux GOARCH=amd64 go build \
	  -ldflags "-X main.VERSION=`git rev-parse --abbrev-ref HEAD`-`git rev-parse --short HEAD``git diff --quiet || echo '-dirty'` -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`'" \
	  -a -o ./transmitter/transmitter github.com/containers-ai/federatorai-agent/cmd

install_dir:
	mkdir -pv $(INSTALL_ROOT)/etc/alameda $(INSTALL_ROOT)/lib/inputlib $(INSTALL_ROOT)/lib/outputlib $(DEST_PREFIX)/bin $(DEST_PREFIX)/etc/input $(DEST_PREFIX)/etc/output

install: install_dir
	cp -fv etc/transmitter.toml $(DEST_PREFIX)/etc/transmitter.toml
	cp -fv etc/input/datapipe.toml $(DEST_PREFIX)/etc/input/datapipe.toml
	cp -fv etc/input/datapipe.toml $(DEST_PREFIX)/etc/output/datapipe.toml
	cp -fv lib/inputlib/datapipe.so $(INSTALL_ROOT)/lib/inputlib/datapipe.so
	cp -fv lib/outputlib/datapipe_recommender.so $(INSTALL_ROOT)/lib/outputlib/datapipe_recommender.so
	cp -fv transmitter/transmitter $(DEST_PREFIX)/bin/
	ln -sfv $(PRODUCT_ROOT)/etc $(INSTALL_ROOT)/etc/alameda/federatorai-agent
	cd $(INSTALL_ROOT); tar -czvf $(SRC_DIR)/install_root.tgz .; cd -

clean:
	rm -fv build/build-image/bin/apiserver install_root.tgz *~

clobber: clean
	rm -rf install_root

build: lib binaries install

.PHONY: run

# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet
	go run ./cmd/transmitter.go run

.PHONY: fmt vet generate docker-build docker-push

# Run go fmt against code
fmt:
	go fmt ./cmd/...

# Run go vet against code
vet:
	go vet ./cmd/...

# Generate code
generate:
	go generate ./cmd/...

## docker-build: Build the docker image.
docker-build-alpine:
	docker build . -t ${IMG} -f Dockerfile

docker-build-ubi:
	docker build . -t ${IMG} -f Dockerfile.ubi

docker-build: docker-build-ubi
