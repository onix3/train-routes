package main

import (
	"fyne.io/fyne/app"
	"github.com/onix3/train-timetable/resource"
	"github.com/onix3/train-timetable/src"
	"time"
)

func main() {
	// приложение может создавать файл конфигурации
	// в Windows находится по пути c:\Users\{имя}\AppData\Roaming\fyne\train-routes\
	src.A = app.NewWithID("train-timetable")
	src.A.SetIcon(resource.TrainPng)

	// считывание сохранённой темы из preferences и её применение
	src.LoadTheme()
	src.A.Settings().SetTheme(src.T)
	// запросы кэшируются: при повторном запросе тех же рейсов запрос к API не осуществляется
	src.LoadCache()
	src.PastCity1,src.PastCity2 = src.LoadLastCities()
	src.СколькоДиаграммСделано = src.LoadCountOfCompletedDiagrams()

	// если есть запрос на сохранение темы, то сохранить
	// нельзя сохранять при каждом изменении цвета, только одни раз
	go func() {
		for {
			if src.SaveRequest {
				src.SaveTheme()
				src.SaveRequest = false
				time.Sleep(time.Second)
			}
			time.Sleep(time.Second)
		}
	}()

	src.W = src.A.NewWindow("Расписаньице")
	src.ScreenHeight = src.GetScreenHeight()
	src.W.SetContent(src.Content())
	src.W.CenterOnScreen()
	src.W.ShowAndRun()
}