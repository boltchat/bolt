# Build stage
FROM golang:1.15.6-alpine AS build

## Git is apparently needed for Mage
RUN apk add git

RUN go get github.com/magefile/mage && \
  go install github.com/magefile/mage

WORKDIR /src
COPY . .

## Compile the static server binary
RUN mage build:serverStatic

# Deploy stage
FROM busybox:1.32.1
WORKDIR /app
COPY --from=build /src/build/boltchat-server-linux-amd64 ./server
COPY scripts/entrypoint.sh /entrypoint.sh

## Executable permissions
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
