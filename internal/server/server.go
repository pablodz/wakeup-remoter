package server

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func ServeWakeUpRemoter() {

	a := app.New()
	w := a.NewWindow("WakeUp Remoter")
	// change size of window
	w.Resize(fyne.NewSize(400, 400))
	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
