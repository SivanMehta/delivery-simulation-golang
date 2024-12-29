package factory

import (
	"main/tools"

	"github.com/rivo/tview"
)

type UI struct {
	View *tview.Flex
	App  *tview.Application
}

func newText(text string) *tview.TextView {
	return tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText(text)
}

func GenerateUI(settings tools.SimulationSettings) *UI {
	app := tview.NewApplication()

	flex := tview.NewFlex().
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Left (1/2 x width of Top)"), 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Middle (3 x height of Top)"), 0, 3, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Bottom (5 rows)"), 5, 1, false), 0, 2, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Right (20 cols)"), 20, 1, false)
	app.SetRoot(flex, true).Run()

	return &UI{
		View: flex,
		App:  app,
	}
}
