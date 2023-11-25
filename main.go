package main

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/rivo/tview"
)

var (
	app       = tview.NewApplication() // The main tview application. contains all other elements.
	flex      = tview.NewFlex()
	menu      = tview.NewList().ShowSecondaryText(false)
	bottomBox = tview.NewTextView() // Container for the lower 2 lines of text
)

// MAIN
func main() {

	// Open a file for logging
	logFile, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0660)
	if err != nil {
		slog.Error("boom!", err)
		os.Exit(1)
	}
	defer logFile.Close()

	logOpts := slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug, ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.SourceKey {
			source := a.Value.Any().(*slog.Source)
			source.File = filepath.Base(source.File)
		}
		return a
	}}

	// Create a logger using the  log file handle
	logger := slog.New(slog.NewTextHandler(logFile, &logOpts))
	slog.SetDefault(logger)

	// Log startup
	slog.Info("Application Startup")

	// Init the menu
	initMenu(menu)

	bottomText := "This is the bottom text"
	bottomBox.SetText(bottomText)

	// Configure flexbox
	flex.SetDirection(tview.FlexRow).
		AddItem(menu, 0, 1, true).
		AddItem(bottomBox, 2, 0, false)

	// Render
	if err := app.SetRoot(flex, true).SetFocus(menu).Run(); err != nil {
		panic(err)
	}

}
