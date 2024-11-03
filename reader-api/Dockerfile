FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN go mod tidy
RUN go mod vendor

RUN go build -o main .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
COPY --from=builder /app/sql/schema /app/sql/schema
COPY --from=builder /go/bin/goose /usr/local/bin/goose

ENTRYPOINT sh -c "goose -dir /app/sql/schema postgres ${DATABASE_URL} up && ./main"
