package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	grpcTaskList "github.com/yazmeyaa/future_undefined_app/api/proto/task_list"
	"github.com/yazmeyaa/future_undefined_app/pkg/logger"
	"github.com/yazmeyaa/future_undefined_app/task_list/internal/handler/grpchandler"
	"github.com/yazmeyaa/future_undefined_app/task_list/internal/repository/postgres"
	sqlctasklist "github.com/yazmeyaa/future_undefined_app/task_list/internal/repository/postgres/sqlc/generated"
	"github.com/yazmeyaa/future_undefined_app/task_list/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}
		return intValue
	}
	return fallback
}

func resolveDBConfig() *pgx.ConnConfig {
	config, _ := pgx.ParseConfig("")
	config.Config.Host = getEnv("DB_HOST", "localhost")
	config.Config.User = getEnv("DB_USER", "postgres")
	config.Config.Password = getEnv("DB_PASSWORD", "postgres")
	config.Config.Port = uint16(getEnvInt("DB_PORT", 5432))
	config.Config.Database = getEnv("DB_NAME", "postgres")
	return config
}

func main() {
	svcLog := logger.NewLogger("task_list")

	dbcfg := resolveDBConfig()
	conn, err := pgx.ConnectConfig(context.Background(), dbcfg)
	if err != nil {
		log.Fatal(err)
	}

	q := sqlctasklist.New(conn)
	repo := postgres.NewPostgresqlRepository(q, svcLog)

	svc := service.New(repo, svcLog)
	h := grpchandler.NewGRPCHandler(svc, svcLog)
	server := grpc.NewServer()
	grpcTaskList.RegisterTaskListServiceServer(server, h)

	exit := make(chan os.Signal, 1)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	svcLog.Info("server started")
	<-exit
	svcLog.Info("shutting down server...")
	server.Stop()
}
