# Build stage
FROM golang:1.15.6-alpine AS build
RUN apk add make

WORKDIR /src
COPY . .

## Compile the static server binary
RUN make build-server-static

# Deploy stage
FROM busybox:1.32.1
WORKDIR /app
COPY --from=build /src/build/server .
COPY scripts/entrypoint.sh /entrypoint.sh

## Executable permissions
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
