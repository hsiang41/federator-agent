# Build the manager binary
FROM golang:1.11.5-stretch as builder

# Copy in the go src
WORKDIR /go/src/github.com/containers-ai/federator-agent
ADD . .
# Build
RUN ["/bin/bash", "-c", "GOOS=linux GOARCH=amd64 go build -ldflags \"-X main.VERSION=`git rev-parse --abbrev-ref HEAD`-`git rev-parse --short HEAD``git diff --quiet || echo '-dirty'` -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`'\" -a -o ./transmitter/transmitter github.com/containers-ai/federator-agent/cmd"]
# RUN ["/bin/bash", "-c", "CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags \"-X main.VERSION=`git rev-parse --abbrev-ref HEAD`-`git rev-parse --short HEAD``git diff --quiet || echo '-dirty'` -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`'\" -buildmode=plugin -a -o ./lib/inputlib/datahub.so github.com/containers-ai/federator-agent/pkg/inputlib/alameda_datahub/datahub.go"]
RUN ["/bin/bash", "-c", "CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -buildmode=plugin -a -o ./lib/inputlib/datahub.so github.com/containers-ai/federator-agent/pkg/inputlib/alameda_datahub"]

# Copy the controller-manager into a thin image
FROM ubuntu:18.04
# FROM alpine:latest
# FROM busybox:latest

WORKDIR /root/
COPY --from=builder /go/src/github.com/containers-ai/federator-agent/etc/transmitter.yml /etc/alameda/federator-agent/transmitter.yml
COPY --from=builder /go/src/github.com/containers-ai/federator-agent/etc/inputlib/datahub.yml /etc/alameda/federator-agent/inputlib/datahub.yml
#COPY --from=builder /go/src/github.com/containers-ai/federator-agent/etc/transmitter.yml /etc/alameda/federator-agent/transmitter.yml
COPY --from=builder /go/src/github.com/containers-ai/federator-agent/transmitter/transmitter .
COPY --from=builder /go/src/github.com/containers-ai/federator-agent/lib/inputlib/datahub.so /lib/inputlib/datahub.so
# EXPOSE 50050/tcp
ENTRYPOINT ["./transmitter"]
CMD [ "run" ]
