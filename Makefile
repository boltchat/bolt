build-server:
	go build -o ./build/server cmd/server/server.go

build-server-static:
	CGO_ENABLED=0 GOOS=linux \
		go build -ldflags "-s" -a -installsuffix cgo \
		-o build/server \
		cmd/server/server.go

docker-build:
	docker build . -t bolt.chat

clean:
	rm -r build/