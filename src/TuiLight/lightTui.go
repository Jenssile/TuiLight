package main

import (
	"github.com/rivo/tview"
	"os/exec"
	"log"
	"fmt"
)

func main() {
	out, err := exec.Command("light").Output()
	if err != nil {
		log.Fatal(err)
	}
	app := tview.NewApplication()
	box := tview.NewBox().SetBorder(true).SetTitle(fmt.Sprintf("%s", out))
	if err := app.SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
