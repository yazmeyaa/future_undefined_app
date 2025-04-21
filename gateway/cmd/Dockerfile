FROM golang:1.24-alpine3.21

RUN mkdir -p /app
WORKDIR /app
ENV GOPATH /app

COPY go.work go.work.sum ./
COPY api/go.mod api/go.sum ./api/
COPY pkg/go.mod pkg/go.sum ./pkg/
COPY gateway/go.mod gateway/go.sum ./gateway/
COPY projects/go.mod projects/go.sum ./projects/
COPY task_list/go.mod task_list/go.sum ./task_list/

RUN go mod download

COPY gateway/ ./gateway
COPY api/ ./api
COPY pkg/ ./pkg
COPY projects/ ./projects
COPY task_list/ ./task_list


RUN GOOS=linux GOARCH=arm64 go build -o /app/gateway_server ./gateway/cmd;
ENTRYPOINT ["/app/gateway_server"]
