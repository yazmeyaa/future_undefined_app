FROM golang:1.24-alpine3.21 AS builder
WORKDIR /app

COPY . .
RUN go mod download
RUN GOOS=linux GOARCH=arm64 go build -o /app/server ./projects/cmd;

FROM golang:1.24-alpine3.21 as runner
WORKDIR /app

COPY --from=builder /app/server /app/server
ENTRYPOINT ["/app/server"]
