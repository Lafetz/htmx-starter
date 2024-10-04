package customlogger

import (
	"log/slog"
	"os"
)
const JsonHandler="json"
func NewLogger(level slog.Level, handler string) *slog.Logger {

	var logHandler slog.Handler
	switch handler {
	case JsonHandler:
		logHandler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     level,
		})
	default:
		logHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: true,
			Level:     level,
		})

	}
	logger := slog.New(logHandler)
	return logger
}