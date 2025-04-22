package main

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v5"
	"github.com/yazmeyaa/future_undefined_app/pkg/logger"
	"github.com/yazmeyaa/future_undefined_app/task_list/internal/model"
	"github.com/yazmeyaa/future_undefined_app/task_list/internal/repository"
	"github.com/yazmeyaa/future_undefined_app/task_list/internal/repository/postgres"
	sqlctasklist "github.com/yazmeyaa/future_undefined_app/task_list/internal/repository/postgres/sqlc/generated"
	"log"
	"os"
	"strconv"
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

	tasklistID, err := repo.CreateTaskList(context.Background(), repository.CreateTaskListDTO{
		UserID:      1,
		Title:       "TEST LIST",
		Description: sql.NullString{},
	})

	task, err := repo.CreateTask(context.Background(), repository.CreateTaskDTO{
		TaskListID: tasklistID,
		UserID:     1,
		Title:      "hello world",
		Description: sql.NullString{
			String: "hello world",
			Valid:  true,
		},
		Status: model.TaskStatusCreated,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", task)
}
