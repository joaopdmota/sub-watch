package logger

import (
	"log/slog"
	"os"
	"sub-watch-microservice/application/services"
)

type SlogLogger struct {
    logger services.Logger
}

func New() *SlogLogger {
    return &SlogLogger{
        logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
    }
}

func (s *SlogLogger) Info(msg string, kv ...any)  { s.logger.Info(msg, kv...) }
func (s *SlogLogger) Warn(msg string, kv ...any)  { s.logger.Warn(msg, kv...) }
func (s *SlogLogger) Error(msg string, kv ...any) { s.logger.Error(msg, kv...) }
func (s *SlogLogger) Debug(msg string, kv ...any) { s.logger.Debug(msg, kv...) }
