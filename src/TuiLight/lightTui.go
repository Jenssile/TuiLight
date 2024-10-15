package main

import (
	"fmt"
	"github.com/rivo/tview"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("light").Output()
	if err != nil {
		log.Fatal(err)
	}
	app := tview.NewApplication()
	box := tview.NewFlex().
		AddItem(tview.NewBox().SetBorder(false), 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(false), 0, 1, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle(fmt.Sprintf("%s", out)), 0, 3, false).
			AddItem(tview.NewBox().SetBorder(false), 5, 1, false), 0, 2, false).
		AddItem(tview.NewBox().SetBorder(false), 0, 1, false)
	if err := app.SetRoot(box, true).SetFocus(box).Run(); err != nil {
		panic(err)
	}
}
