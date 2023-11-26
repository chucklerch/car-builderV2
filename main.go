package main

import (
	"log/slog"

	"github.com/rivo/tview"
)

var (
	app = tview.NewApplication() // The main tview application. contains all other elements.
)

// MAIN
func main() {

	// Create a logger
	logger := setupLogging()
	slog.SetDefault(logger)

	// Log startup
	slog.Info("Application Startup")

	// Init the menu
	initMenu(menu)

	topText := "[::d]This is the top text"
	topBox.SetText(topText).SetDynamicColors(true).SetBorder(false)
	topBox.SetTextAlign(tview.AlignCenter)

	bottomText := "This is the bottom text"
	bottomBox.SetText(bottomText)

	// Configure flexbox
	flex.SetDirection(tview.FlexRow).
		AddItem(topBox, 1, 1, false).
		AddItem(menu, 0, 1, true).
		AddItem(bottomBox, 2, 0, false)

	// Render
	if err := app.SetRoot(flex, true).SetFocus(menu).Run(); err != nil {
		panic(err)
	}

}
