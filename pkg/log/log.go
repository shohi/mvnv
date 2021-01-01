package log

import (
	slog "github.com/sirupsen/logrus"
)

func InitLog(level string) {
	l, err := slog.ParseLevel(level)
	if err != nil {
		l = slog.InfoLevel
	}

	slog.SetLevel(l)
	slog.SetFormatter(&slog.TextFormatter{
		DisableColors:    true,
		DisableTimestamp: true,
	})
}
