FROM golang:1.22.0-alpine AS builder

WORKDIR /github.com/go-jedi/auth-service/app/
COPY . /github.com/go-jedi/auth-service/app/

RUN go mod download
RUN go build -o .bin/grpc_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/go-jedi/auth-service/app/.bin/grpc_server .
COPY .env /root/

CMD ["./grpc_server"]