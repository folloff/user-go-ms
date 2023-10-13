FROM golang:1.20.3-alpine AS builder

# will copy all container inside, new generated path inside of the docker container
COPY . /github.com/folloff/auth-go-ms/grpc/source/
WORKDIR /github.com/folloff/auth-go-ms/grpc/source/

RUN go mod download
RUN go build -o ./bin/auth_server cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/folloff/auth-go-ms/grpc/source/bin/auth_server .

CMD ["./auth_server"]
