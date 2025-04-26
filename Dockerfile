FROM golang:1.24-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ecometer2mqtt

FROM alpine:3.21
WORKDIR /app
COPY --from=builder /src/ecometer2mqtt /app/ecometer2mqtt
COPY config-example.yaml /app/config.yaml
ENTRYPOINT ["/app/ecometer2mqtt"]