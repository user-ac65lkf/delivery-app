FROM golang:1.21-alpine AS builder

COPY ../admin /app/admin
RUN ls
WORKDIR /app/admin

RUN go mod download

RUN go build -o /bin/admin_service ./cmd/api/main.go

FROM alpine
RUN apk --no-cache add bash curl
WORKDIR /root/

COPY --from=builder ./bin/admin_service .
COPY --from=builder ./app/admin/dev/local.env dev/local.env

CMD ["./admin_service"]