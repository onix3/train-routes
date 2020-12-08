package src

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"time"
)

func thisDay0000(t time.Time) time.Time {
	y, m, d := t.Year(), t.Month(), t.Day()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
}

var (
	today0000   = thisDay0000(time.Now())
	routesCache = map[string][]route{} // кэш запросов
	Last2       string
	Last1       string // последние выбранные города
)

var (
	A                fyne.App       // приложение
	W                fyne.Window    // главное окно
	select1, select2 *widget.Select // селекторы
)
