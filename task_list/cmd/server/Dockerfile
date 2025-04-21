FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN GOOS=linux GOARCH=arm64 go build -o /app/server ./task_list/cmd/server;

FROM golang:1.24-alpine AS runner
WORKDIR /app
COPY --from=builder /app/server /app/task_list_bin

ENTRYPOINT ["/app/task_list_bin"]