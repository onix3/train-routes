package main

import (
	"fyne.io/fyne/app"
	"github.com/onix3/train-routes/resource"
	"github.com/onix3/train-routes/src"
)

func main() {
	//_ = src.GetAllSortedRoutes("Брест","Минск",time.Now().Format("2006-01-02"),"train")

	src.A = app.NewWithID("train-routes")
	src.A.SetIcon(resource.TrainPng)

	// запросы кэшируются: при повторном запросе тех же рейсов запрос к API не осуществляется
	src.LoadCache()

	src.W = src.A.NewWindow("Расписаньице")
	src.W.SetContent(src.Content())
	src.W.CenterOnScreen()
	src.W.ShowAndRun()
}