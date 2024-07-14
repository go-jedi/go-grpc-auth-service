FROM golang:1.22.5-alpine AS builder

WORKDIR /github.com/go-jedi/auth/app
COPY . /github.com/go-jedi/auth/app

RUN go mod download
RUN go build -o .bin/app cmd/app/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/go-jedi/auth/app/.bin/app .
COPY migrations /root/migrations
COPY config.yaml /root/

CMD ["./app", "--config", "config.yaml"]