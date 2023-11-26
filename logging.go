package main

import (
	"log/slog"
	"os"
	"path/filepath"
)

func setupLogging() *slog.Logger {
	// Open a file for logging
	logFile, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0660)
	if err != nil {
		slog.Error("Fatal Error!", err)
		os.Exit(1)
	}

	logOpts := slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug, ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.SourceKey {
			source := a.Value.Any().(*slog.Source)
			source.File = filepath.Base(source.File)
		}
		return a
	}}

	// Create a logger using the  log file handle
	return slog.New(slog.NewTextHandler(logFile, &logOpts))
}
