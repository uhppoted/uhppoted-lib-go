package log

import (
	"fmt"
	"log/slog"
)

// Debugf formats a message and logs it as a DEBUG level slog message.
func Debugf(tag string, format string, args ...any) {
	f := fmt.Sprintf("%-8v %v", tag, format)
	msg := fmt.Sprintf(f, args...)

	slog.Debug(msg)
}

// Debugf formats a message and logs it as an INFO level slog message.
func Infof(tag string, format string, args ...any) {
	f := fmt.Sprintf("%-8v %v", tag, format)
	msg := fmt.Sprintf(f, args...)

	slog.Info(msg)
}

// Debugf formats a message and logs it as a WARN level slog message.
func Warnf(tag string, format string, args ...any) {
	f := fmt.Sprintf("%-8v %v", tag, format)
	msg := fmt.Sprintf(f, args...)

	slog.Warn(msg)
}

// Debugf formats a message and logs it as an ERROR level slog message.
func Errorf(tag string, format string, args ...any) {
	f := fmt.Sprintf("%-8v %v", tag, format)
	msg := fmt.Sprintf(f, args...)

	slog.Error(msg)
}
