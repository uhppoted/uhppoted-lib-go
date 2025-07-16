package log

import (
	"fmt"
	"log/slog"
)

func init() {
	h := handler{}
	logger := slog.New(&h)

	slog.SetDefault(logger)
	slog.SetLogLoggerLevel(slog.LevelDebug)
}

func Debugf(tag string, format string, args ...any) {
	f := fmt.Sprintf("%-8v %v", tag, format)
	msg := fmt.Sprintf(f, args...)

	slog.Debug(msg)
}

func Infof(tag string, format string, args ...any) {
	f := fmt.Sprintf("%-8v %v", tag, format)
	msg := fmt.Sprintf(f, args...)

	slog.Info(msg)
}

func Warnf(tag string, format string, args ...any) {
	f := fmt.Sprintf("%-8v %v", tag, format)
	msg := fmt.Sprintf(f, args...)

	slog.Warn(msg)
}

func Errorf(tag string, format string, args ...any) {
	f := fmt.Sprintf("%-8v %v", tag, format)
	msg := fmt.Sprintf(f, args...)

	slog.Error(msg)
}
