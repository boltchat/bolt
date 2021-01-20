build-server:
	go build cmd/server/server.go

build-server-static:
	CGO_ENABLED=0 GOOS=linux \
		go build -ldflags "-s" -a -installsuffix cgo \
		cmd/server/server.go

docker-build:
	docker build . -t go-tcp-chat