package src

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/driver/desktop"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/fogleman/gg"
	"net/url"
	"os"
)

func makeF1Content() fyne.CanvasObject {
	/////// Клавиши ///////
	labelEnter := widget.NewLabel("Enter")
	labelEnter.TextStyle.Bold = true
	labelF1 := widget.NewLabel("\nF1")
	labelF1.TextStyle.Bold = true

	vbox1 := widget.NewVBox(labelEnter, labelF1)
	vbox2 := widget.NewVBox(widget.NewLabel("Сохранить диаграмму\nна Рабочий стол"), widget.NewLabel("Вызвать это окно :)"))

	group1 := widget.NewGroup("Клавиши",
		widget.NewHBox(vbox1, vbox2),
	)

	/////// Инфо ///////

	label1 := widget.NewLabel("Подробная информация")
	parsed1,_ := url.Parse("https://github.com/onix3/train-timetable#train-timetable")
	hyperLink1 := widget.NewHyperlink("здесь", parsed1)

	group2 := widget.NewGroup("Инфо",
		widget.NewHBox(label1, layout.NewSpacer(), hyperLink1),
	)


	return widget.NewVBox(
		group1,
		widget.NewLabel(""),
		group2,
		widget.NewLabel(""),
	)
}

// Захват нажатий клавиш
func KeyBindings(win fyne.Window) {
	f1Content := makeF1Content()

	f1DialogIsDisplayed := false

	// функция обработки
	f := func(ev *fyne.KeyEvent) {
		s := ev.Name
		fmt.Println(s)
		// если буква, сохранить последнюю диаграмму на Рабочий Стол
		if s == fyne.KeyEnter || s == fyne.KeyReturn {
			if diagram != nil && len(diagramName.Text) > 0 {
				// сохранение диаграммы на Рабочий Стол
				dir,err := os.UserHomeDir()
				IsErr(err)
				err = gg.SavePNG(dir + "\\Desktop\\" + diagramName.Text + ".png", diagram)
				IsErr(err)
			}
		}

		// если F1, отобразить диалог со справкой
		// притом нужно отобразить только если ещё не отображается
		if s == fyne.KeyF1 && !f1DialogIsDisplayed {
			helpD := dialog.NewCustom("Справка", "OK", f1Content, win)
			helpD.SetOnClosed(func() {
				f1DialogIsDisplayed = false
			})
			helpD.Show()
			f1DialogIsDisplayed = true
		}
	}

	if desk,ok := win.Canvas().(desktop.Canvas); ok {
		desk.SetOnKeyDown(f)
	}
}