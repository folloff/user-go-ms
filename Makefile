LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-auth-api

generate-auth-api:
	mkdir -p pkg/auth_v1
	protoc --proto_path internal/controller/grpc/auth_v1 \
	--go_out=pkg/auth_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/auth_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	internal/controller/grpc/auth_v1/user.proto

build-local:
	go build -o service cmd/grpc_server/main.go

build:
	GOOS=linux GOARCH=amd64 go build -o service_linux cmd/main.go

copy-to-server:
	scp service_linux root@134.209.224.105:

# doctl registry login OR docker login registry.digitalocean.com
docker-build-and-push:
	docker buildx build --no-cache --platform linux/amd64 -t registry.digitalocean.com/folloff/auth-ms:v0.0.2 .
	# Notice: Login valid for 30 days. Use the --expiry-seconds flag to set a shorter expiration or --never-expire for no expiration.
	doctl registry login -c config/docker-config.json -t TOKEN
	# docker login -u  TOKEN -p TOKEN registry.digitalocean.com
	docker push registry.digitalocean.com/folloff/auth-ms:v0.0.2

