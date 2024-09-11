# Use Alpine Linux as the base image
FROM alpine:3.20.0

# Copy binaries
COPY ./api.go ./renew.sh /usr/local/bin/

ENV CHANCE=10
ENV RENEWAL_TIME=24:00:00
ENV TZ=UTC

# Make configurations
RUN \
    # Enable testing repo
    echo "@testing http://dl-cdn.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories && \
    apk update && \
    # Install dependencies
    apk add --no-cache cowsay@testing fortune openrc go tzdata && \
    # Create restricted user
    addgroup api && \
    adduser -s /bin/nologin -D -G api api && \
    # Compile binary
    CGO_ENABLED=0 go build -o /usr/local/bin/api /usr/local/bin/api.go && \
    rm /usr/local/bin/api.go && \
    apk del go openrc && \
    # Set permissions
    chown -R root:api /srv && \
    chmod -R 770 /srv && \
    chown root:api /usr/local/bin/api /usr/local/bin/renew.sh && \
    chmod 550 /usr/local/bin/api /usr/local/bin/renew.sh

# Expose HTTP port
EXPOSE 80

USER api

# Entrypoint configuration
ENTRYPOINT ["/usr/local/bin/api"]