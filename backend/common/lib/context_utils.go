package lib

import (
	"context"
	"path"
	"runtime"
	"strconv"

	"github.com/sirupsen/logrus"
)

type FormatterType string

const (
	JSONFormatter FormatterType = "json"
	TextFormatter FormatterType = "text"
)

type contextKey string

const loggerKey contextKey = "logger"

func generateCallerPrettyfier(frame *runtime.Frame) (function string, file string) {
	fileName := path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)
	return "", fileName
}

func setFormatter(formatterType FormatterType) logrus.Formatter {
	var formatter logrus.Formatter

	switch formatterType {
	case JSONFormatter:
		formatter = &logrus.JSONFormatter{
			CallerPrettyfier: generateCallerPrettyfier,
		}
	case TextFormatter:
		formatter = &logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		}
	default:
		formatter = &logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		}
	}
	return formatter
}

func NewLogger(formatterType FormatterType) *logrus.Logger {
	if formatterType == "" {
		formatterType = JSONFormatter
	}

	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetFormatter(setFormatter(formatterType))
	logger.SetLevel(logrus.InfoLevel)
	return logger
}

func WithLogger(ctx context.Context, logger *logrus.Entry) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func LoggerFromContext(ctx context.Context) *logrus.Entry {
	logger, ok := ctx.Value(loggerKey).(*logrus.Entry)
	if !ok {
		return logrus.NewEntry(NewLogger(JSONFormatter))
	}
	return logger
}
