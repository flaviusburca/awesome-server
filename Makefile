GOPROXY?="https://proxy.golang.org,direct"

all: server-amd64 client-amd64 server-arm64 client-arm64

server-amd64:
	docker run --rm -v .:/usr/src/myapp -w /usr/src/myapp -e GOOS=linux -e GOARCH=amd64 golang:1.19 go build -o ./build/server-amd64 ./cmd/server/server.go

server-arm64:
	docker run --rm -v .:/usr/src/myapp -w /usr/src/myapp -e GOOS=linux -e GOARCH=arm64 golang:1.19 go build -o ./build/server-arm64 ./cmd/server/server.go

client-amd64:
	docker run --rm -v .:/usr/src/myapp -w /usr/src/myapp -e GOOS=linux -e GOARCH=amd64 golang:1.19 go build -o ./build/client-amd64 ./cmd/client/client.go

client-arm64:
	docker run --rm -v .:/usr/src/myapp -w /usr/src/myapp -e GOOS=linux -e GOARCH=arm64 golang:1.19 go build -o ./build/client-arm64 ./cmd/client/client.go

