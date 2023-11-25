package main

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	app  = tview.NewApplication() // The main tview application. contains all other elements.
	flex = tview.NewFlex()
	menu = tview.NewList().ShowSecondaryText(false)
)

func initMenu(l *tview.List) {
	l.AddItem("Quit", "Weapon", rune('Q'), nil)
	l.SetBorderPadding(1, 1, 1, 1)

	// Configure the list's style
	l.SetBorder(true)
	l.SetMainTextColor(tcell.ColorDarkSlateGrey)
	l.SetSelectedBackgroundColor(tcell.ColorBlack)
	l.SetSelectedTextColor(tcell.Color(tcell.AttrBold))
	l.SetShortcutColor(tcell.ColorDarkCyan)
	l.SetSelectedFunc(func(index int, name string, secondaryText string, shortcut rune) {
		switch name {
		case "Quit":
			app.Stop()
		}
	})

	l.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch key := event.Rune(); {
		case key == 'q':
			app.Stop()
		}
		return event
	})
}

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

	initMenu(menu)

	// Render
	if err := app.SetRoot(flex, true).SetFocus(menu).Run(); err != nil {
		panic(err)
	}

}
