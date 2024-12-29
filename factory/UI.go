package factory

import (
	"fmt"
	"main/tools"
	"os"
	"time"

	"github.com/rivo/tview"
)

type UI struct {
	View *tview.Modal
	App  *tview.Application
}

func currentTimeString() string {
	t := time.Now()
	return fmt.Sprintf(t.Format("Current time is 15:04:05"))
}

func GenerateUI(settings tools.SimulationSettings) *UI {
	settingsString := tools.PrintArgs(settings)

	app := tview.NewApplication()
	view := tview.NewModal().
		SetText(settingsString).
		AddButtons([]string{"Quit"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Quit" {
				app.Stop()
				os.Exit(0)
			}
		})

	app.SetRoot(view, false).Run()

	return &UI{
		View: view,
		App:  app,
	}
}
