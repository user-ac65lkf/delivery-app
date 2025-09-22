FROM golang:1.21-alpine AS builder


COPY ../delivery /app/delivery
RUN ls
WORKDIR /app/delivery

RUN go mod download

RUN go build -o /bin/delivery_service ./cmd/api/main.go

FROM alpine
RUN apk --no-cache add bash curl
WORKDIR /root/

COPY --from=builder ./bin/delivery_service .
COPY --from=builder ./app/delivery/dev/local.env dev/local.env

CMD ["./delivery_service"]