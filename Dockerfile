# Build the manager binary
FROM golang:1.11.5-stretch as builder

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# Copy in the go src
WORKDIR /go/src/github.com/containers-ai/federatorai-agent
ADD . .

# Build
RUN make build

# Copy the controller-manager into a thin image
FROM ubuntu:18.04
# FROM alpine:latest
# FROM busybox:latest

WORKDIR /root/
COPY --from=builder /go/src/github.com/containers-ai/federatorai-agent/install_root.tgz /tmp/

RUN set -x \
    && cd / && tar xzvf /tmp/install_root.tgz && rm -fv /tmp/install_root.tgz \
    && chown -R 1001:0 /opt/alameda \
    && mkdir -p /var/log/alameda && chown -R 1001:0 /var/log/alameda && chmod ug+w /var/log/alameda

USER 1001
ENTRYPOINT ["/opt/alameda/federatorai-agent/bin/transmitter", "run"]
