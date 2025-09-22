FROM golang:1.21-alpine AS builder

COPY ../shop /app/shop
RUN ls
WORKDIR /app/shop

RUN go mod download

RUN apk --no-cache update && \
apk --no-cache add gcc g++ libc-dev librdkafka-dev pkgconf


RUN go build -tags musl -o /bin/shop_service ./cmd/main.go

FROM alpine
RUN apk --no-cache add bash curl

WORKDIR /root/

COPY --from=builder ./bin/shop_service .
COPY --from=builder ./app/shop/dev/local.env dev/local.env

CMD ["./shop_service"]
