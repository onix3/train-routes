package main

import (
	"fyne.io/fyne/app"
	"github.com/onix3/train-routes/src"
	"time"
)

func main() {
	_ = src.GetAllSortedRoutes("Брест","Минск",time.Now().Format("2006-01-02"),"train")
	_ = src.GetAllSortedRoutes("Брест","Минск",time.Now().Format("2006-01-02"),"train")

	src.A = app.New()

	src.W = src.A.NewWindow("Расписаньице")
	src.W.SetContent(src.Content())
	src.W.CenterOnScreen()
	src.W.ShowAndRun()
}