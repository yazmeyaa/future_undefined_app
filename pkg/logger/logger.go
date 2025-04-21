package logger

import (
	"log/slog"
	"os"
)

func NewLogger(svcName string) *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	l := slog.New(handler)

	svcLog := l.With(slog.Group(
		"service_info",
		slog.String("name", svcName),
	))

	return svcLog
}
