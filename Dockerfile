FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o api-server ./cmd/main.go
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/api-server ./api-server
COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY internal/database/migrations ./internal/database/migrations
COPY entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
