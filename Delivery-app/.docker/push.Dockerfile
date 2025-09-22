FROM golang:1.21-alpine AS builder

COPY ../w_push_service /app/push
RUN ls
WORKDIR /app/push

RUN go mod download

RUN apk --no-cache update && \
apk --no-cache add gcc g++ libc-dev librdkafka-dev pkgconf


RUN go build -tags musl -o /bin/push_service ./cmd/main.go

FROM alpine

RUN apk --no-cache add bash curl
WORKDIR /root/

COPY --from=builder ./bin/push_service .

CMD ["./push_service"]