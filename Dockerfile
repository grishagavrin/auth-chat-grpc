FROM golang:1.21.4-alpine AS builder

COPY . /github.com/grishagavrin/auth-chat-grpc/source/
WORKDIR /github.com/grishagavrin/auth-chat-grpc/source/

RUN go mod download
RUN go build -o ./bin/crud_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /usr/local/src/
COPY --from=builder /github.com/grishagavrin/auth-chat-grpc/source/bin/crud_server .

CMD ["./crud_server"]