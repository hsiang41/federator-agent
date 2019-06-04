# Build the manager binary
FROM golang:1.11.5-stretch as builder

# Copy in the go src
WORKDIR /go/src/github.com/containers-ai/federatorai-agent
ADD . .

# Build
RUN make test && make binaries && make install

# Copy the controller-manager into a thin image
FROM ubuntu:18.04
# FROM alpine:latest
# FROM busybox:latest

WORKDIR /root/
COPY --from=builder /etc/alameda /etc/alameda
COPY --from=builder /root/transmitter /root/transmitter
COPY --from=builder /lib/inputlib /lib/inputlib
COPY --from=builder /lib/outputlib /lib/outputlib

# EXPOSE 50050/tcp
ENTRYPOINT ["./transmitter"]
CMD [ "run" ]
