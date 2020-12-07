package main

import (
	"fyne.io/fyne/app"
	"github.com/onix3/train-routes/src"
)

func main() {
	src.A = app.New()

	src.W = src.A.NewWindow("Расписаньице")
	src.W.SetContent(src.Content())
	src.W.CenterOnScreen()
	src.W.ShowAndRun()
}