package server

import (
	"fmt"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func ServeWakeUpRemoter() {

	t := time.Now().Local().Format("2006-01-02 15:04:05")
	fmt.Println("today date is = ", t)

	a := app.New()
	w := a.NewWindow("WakeUp Remoter")
	// change size of window
	w.Resize(fyne.NewSize(400, 400))

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
