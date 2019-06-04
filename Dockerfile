# Build the manager binary
FROM golang:1.11.5-stretch as builder

# Copy in the go src
WORKDIR /go/src/github.com/containers-ai/federatorai-agent
ADD . .
# Build
RUN ["/bin/bash", "-c", "GOOS=linux GOARCH=amd64 go build -ldflags \"-X main.VERSION=`git rev-parse --abbrev-ref HEAD`-`git rev-parse --short HEAD``git diff --quiet || echo '-dirty'` -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`'\" -a -o ./transmitter/transmitter github.com/containers-ai/federatorai-agent/cmd"]
# RUN ["/bin/bash", "-c", "CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags \"-X main.VERSION=`git rev-parse --abbrev-ref HEAD`-`git rev-parse --short HEAD``git diff --quiet || echo '-dirty'` -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`'\" -buildmode=plugin -a -o ./lib/inputlib/datahub.so github.com/containers-ai/federatorai-agent/pkg/inputlib/alameda_datapipe/datapipe.go"]
RUN ["/bin/bash", "-c", "CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -buildmode=plugin -a -o ./lib/inputlib/datapipe.so github.com/containers-ai/federatorai-agent/pkg/inputlib/alameda_datapipe"]
RUN ["/bin/bash", "-c", "CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -buildmode=plugin -a -o ./lib/outputlib/datapipe_writer.so github.com/containers-ai/federatorai-agent/pkg/outputlib/alameda_datapipe"]


# Copy the controller-manager into a thin image
FROM ubuntu:18.04
# FROM alpine:latest
# FROM busybox:latest

WORKDIR /root/
COPY --from=builder /go/src/github.com/containers-ai/federatorai-agent/etc/transmitter.toml /etc/alameda/federatorai-agent/transmitter.toml
COPY --from=builder /go/src/github.com/containers-ai/federatorai-agent/etc/inputlib/datapipe.toml /etc/alameda/federatorai-agent/inputlib/datapipe.toml
#COPY --from=builder /go/src/github.com/containers-ai/federatorai-agent/etc/transmitter.toml /etc/alameda/federatorai-agent/transmitter.toml
COPY --from=builder /go/src/github.com/containers-ai/federatorai-agent/transmitter/transmitter .
COPY --from=builder /go/src/github.com/containers-ai/federatorai-agent/lib/inputlib/datapipe.so /lib/inputlib/datapipe.so
COPY --from=builder /go/src/github.com/containers-ai/federatorai-agent/lib/outputlib/datapipe_writer.so /lib/outputlib/datapipe_writer.so
# EXPOSE 50050/tcp
ENTRYPOINT ["./transmitter"]
CMD [ "run" ]
