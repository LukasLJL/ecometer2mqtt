FROM golang:1.26-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ecometer2mqtt

FROM alpine:3.24
WORKDIR /app
COPY --from=builder /src/ecometer2mqtt /app/ecometer2mqtt
COPY config-example.yaml /app/config.yaml
ENTRYPOINT ["/app/ecometer2mqtt"]