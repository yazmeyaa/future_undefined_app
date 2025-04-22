package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	grpcProjects "github.com/yazmeyaa/future_undefined_app/api/proto/projects"
	"github.com/yazmeyaa/future_undefined_app/pkg/logger"
	"github.com/yazmeyaa/future_undefined_app/projects/internal/handler/projectsGRPCHandler"
	"github.com/yazmeyaa/future_undefined_app/projects/internal/repository/postgres"
	sqlcproj "github.com/yazmeyaa/future_undefined_app/projects/internal/repository/postgres/sqlc/generated"
	"github.com/yazmeyaa/future_undefined_app/projects/internal/service"
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
	svcLogger := logger.NewLogger("projects")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	cfg := resolveDBConfig()
	db, err := pgx.ConnectConfig(context.Background(), cfg)
	if err != nil {
		log.Fatal(err)
	}

	sqlcdb := sqlcproj.New(db)
	repo := postgres.NewRepository(sqlcdb)
	svc := service.NewProjectsService(repo, svcLogger)
	srv := grpc.NewServer()
	handler := projectsGRPCHandler.NewGRPCHandler(svc)
	grpcProjects.RegisterProjectsServiceServer(srv, handler)

	exit := make(chan os.Signal, 1)

	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	svcLogger.Info("server started")
	<-exit
	svcLogger.Info("shutting down server...")
	srv.Stop()
}
