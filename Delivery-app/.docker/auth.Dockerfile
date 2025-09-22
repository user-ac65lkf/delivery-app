FROM golang:1.21-alpine AS builder


COPY ../auth /app/auth
RUN ls
WORKDIR /app/auth
#
#COPY ../auth/go.mod .
#COPY ../auth/go.sum .

RUN go mod download

RUN go build -o /bin/auth_service ./cmd/main.go

FROM alpine
RUN apk --no-cache add bash curl
WORKDIR /root/

COPY --from=builder ./bin/auth_service .
COPY --from=builder ./app/auth/dev/local.env dev/local.env
COPY --from=builder ./app/auth/migrations ./migrations

CMD ["./auth_service"]