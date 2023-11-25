package main

import (
	"log/slog"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func initMenu(l *tview.List) {
	l.AddItem("Quit", "Quit", rune('Q'), nil)
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
			confirmQuit()
		}
	})

	l.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch key := event.Rune(); {
		case key == 'q':
			confirmQuit()
		}
		return event
	})
}

func confirmQuit() {
	slog.Info("Confirm Quit")
	quitModal := tview.NewModal().SetText("Quit?").AddButtons([]string{"Yes", "No"})
	quitModal.SetBackgroundColor(tcell.ColorDarkSlateGrey)
	quitModal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "Yes" {
			slog.Info("Exiting")
			app.Stop()
		} else {
			app.SetRoot(flex, false).SetFocus(menu)
		}
	})
	app.SetRoot(quitModal, false).SetFocus(quitModal)
}
