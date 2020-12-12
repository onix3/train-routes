package src

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/fogleman/gg"
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

func thisDay0000(t time.Time) time.Time {
	y,m,d := t.Year(), t.Month(), t.Day()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
}

var (
	today0000              = thisDay0000(time.Now())
	routesCache            = map[string][]route{} // кэш запросов
	PastCity1, PastCity2   string // последние выбранные города
	СколькоДиаграммСделано int
	SaveRequest            bool
)

var (
	A                              fyne.App       // приложение
	W                              fyne.Window    // главное окно
	select1, select2               *widget.Select // селекторы
	forwardPattern, reversePattern gg.Pattern
	resultImage                    image.Image   // изображение диаграммы
	resultImageWidget              *canvas.Image // миниатюра диаграммы
	resultText                     *canvas.Text
	resultBox                      *tappableBox
	ScreenWidth, ScreenHeight      int
	mapImage                       image.Image                 // изображение карты
	T                              *userTheme                  // текущая тема
	mainButton                     *widget.Button
)