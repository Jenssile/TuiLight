package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	out, err := exec.Command("light").Output()
	if err != nil {
		log.Fatal(err)
	}
	app := tview.NewApplication()
	flex := tview.NewFlex().
		AddItem(tview.NewBox().SetBorder(false), 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(false), 0, 1, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle(fmt.Sprintf("%s", out)), 0, 3, false).
			AddItem(tview.NewBox().SetBorder(false), 5, 1, false), 0, 2, false).
		AddItem(tview.NewBox().SetBorder(false), 0, 1, false)
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}
