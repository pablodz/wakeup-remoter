package server

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	xwidget "fyne.io/x/fyne/widget"
)

var data = []string{"a", "string", "list"}

func ServeWakeUpRemoter() {
	a := app.New()
	w := a.NewWindow("WakeUp Remoter")
	//
	w.Resize(fyne.NewSize(600, 800))
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})
	list.OnSelected = func(id widget.ListItemID) {
		editWaker(a)
	}

	w.SetContent(list)
	w.ShowAndRun()
}

func editWaker(a fyne.App) {
	w2 := a.NewWindow("Calendar")

	i := widget.NewLabel("Please Choose a Date")
	i.Alignment = fyne.TextAlignCenter

	l := widget.NewLabel("No Date Selected")
	l.Alignment = fyne.TextAlignCenter

	d := &date{instruction: i, dateChosen: l}
	startingDate := time.Now()
	calendar := xwidget.NewCalendar(startingDate, d.onSelected)

	titleHour := widget.NewLabel("Hour")
	titleHour.Alignment = fyne.TextAlignCenter

	statusHour := widget.NewLabel("-")
	statusHour.Alignment = fyne.TextAlignCenter

	hour := widget.NewEntry()
	hour.OnChanged = func(s string) {
		// validate hour format
		validDate := ValidateStringHourFormat(s)
		if validDate {
			statusHour.SetText("Valid")
		} else {
			statusHour.SetText("Invalid")
		}
	}

	c := container.NewVBox(i, l, calendar, titleHour, hour, statusHour)

	w2.SetContent(c)
	w2.Show()
}

type date struct {
	instruction *widget.Label
	dateChosen  *widget.Label
}

func (d *date) onSelected(t time.Time) {
	// use time object to set text on label with given format
	d.instruction.SetText("Date Selected:")
	d.dateChosen.SetText(t.Format("Mon 02 Jan 2006"))
}

func ValidateStringHourFormat(s string) bool {

	if len(s) != 5 {
		return false
	}

	if s[0] < '0' || s[0] > '2' {
		return false
	}

	if s[1] < '0' || s[1] > '9' {
		return false
	}

	if s[2] != ':' {
		return false
	}

	if s[3] < '0' || s[3] > '5' {
		return false
	}

	if s[4] < '0' || s[4] > '9' {
		return false
	}

	if s[0] == '2' && s[1] > '3' {
		return false
	}

	return true
}
