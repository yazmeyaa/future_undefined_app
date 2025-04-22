FROM golang:1.24-alpine3.21 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN GOOS=linux GOARCH=arm64 go build -o /app/gateway_server ./gateway/cmd;

FROM golang:1.24-alpine3.21 as runner
WORKDIR /app
COPY --from=builder /app/gateway_server .
ENTRYPOINT ["/app/gateway_server"]
