FROM golang:1.23-alpine AS builder
WORKDIR /app
ADD . .
COPY go.mod go.sum ./
RUN go build -o bin/devopsmon ./cmd/main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/bin/devopsmon ./devopsmon
COPY --from=builder /app/config/config.yaml ./config/config.yaml
EXPOSE 8000
ENTRYPOINT ["./devopsmon"]