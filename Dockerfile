# Build the manager binary
FROM golang:1.11.5-stretch as builder

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# Copy in the go src
WORKDIR /go/src/github.com/containers-ai/federatorai-agent
ADD . .

# Build
RUN make build

# Prepare the package into a thin image
FROM ubuntu:18.04
# FROM alpine:latest
# FROM busybox:latest

ENV AIHOME=/opt/alameda/federatorai-agent \
    USER_UID=1001 \
    USER_NAME=alameda

WORKDIR ${AIHOME}
COPY --from=builder /go/src/github.com/containers-ai/federatorai-agent/tini /sbin/tini
COPY --from=builder /go/src/github.com/containers-ai/federatorai-agent/install_root.tgz /tmp/

RUN set -x \
    && echo "${USER_NAME}:x:${USER_UID}:0:Federator.ai:${AIHOME}:/bin/sh" >> /etc/passwd \
    && chmod ug+w /var/log \
    # install packages
    && cd / && tar xzvf /tmp/install_root.tgz && rm -fv /tmp/install_root.tgz \
    && chown -R ${USER_UID}:root ${AIHOME} && chmod -R ug+w ${AIHOME} \
    && mkdir -pv -m 775 /var/log/alameda && chown -R ${USER_UID}:root /var/log/alameda

USER ${USER_UID}
ENTRYPOINT ["/sbin/tini","-v", "--"]
CMD ["/init.sh"]
