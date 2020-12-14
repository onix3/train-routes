package src

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"image"
	"time"
)

// "Анимация" иконки кнопки во время запроса
func button1() {
	mainButton.Icon = theme.MoveUpIcon()
	mainButton.Refresh()
}
func button2() {
	mainButton.Icon = theme.MoveDownIcon()
	mainButton.Refresh()
}
func buttonBack(d time.Duration) {
	time.Sleep(d)
	mainButton.Icon = theme.DocumentCreateIcon()
	mainButton.Refresh()
}

// Для даты/времени время свести к 00:00:00
//func thisDay0000(t time.Time) time.Time {
//	y, m, d := t.Year(), t.Month(), t.Day()
//	return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
//}

func Now0000() time.Time {
	t := time.Now()
	y, m, d := t.Year(), t.Month(), t.Day()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
}

var (
	// кэш запросов
	routesCache = map[string][]route{}
	// последние выбранные города
	PastCity1, PastCity2   string
	// надпись Янд-кс отображается только раз, поэтому считаю, сколько диаграмм было показано
	// (Go поддерживает unicode в названиях идентификаторов — почему бы этим не воспользоваться)
	СколькоДиаграммПоказано int
	// Раз в 2 секунды проверяется, есть ли запрос на сохранение изменений темы
	SaveRequest            bool
)

var (
	// приложение
	A fyne.App
	// главное окно
	W fyne.Window
	// селекторы
	//select1, select2 *widget.Select
	// диаграмма
	diagram image.Image
	// миниатюра диаграммы, название и объединяющий контейнер
	diagramWidget *canvas.Image
	diagramName   *canvas.Text
	diagramBox    *tappableBox
	// разрешение экрана
	ScreenWidth, ScreenHeight int
	// карта
	mapImage image.Image
	// текущая тема
	T *userTheme
	// наиглавнейшая кнопка в приложении
	mainButton *widget.Button
	// область диаграммы, которая будет показана для пояснения узоров
	cropRect image.Rectangle
)
