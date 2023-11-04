FROM golang:1.21-alpine AS builder

COPY . /github.com/sagata1999/chat/
WORKDIR /github.com/sagata1999/chat/

RUN go mod download
RUN go build -o ./bin/chat_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/sagata1999/chat/bin/chat_server .

CMD ["./chat_server"]