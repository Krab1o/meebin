FROM golang:1.24.1-alpine AS builder

COPY . /github.com/Krab1o/meebin/src
WORKDIR /github.com/Krab1o/meebin/src

RUN go mod download \
    && go build -o bin/crud_server cmd/http-server/main.go

FROM alpine:3.21

WORKDIR /app/

COPY --from=builder /github.com/Krab1o/meebin/src/bin/crud_server .

CMD ["./crud_server"]