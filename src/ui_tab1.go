package src

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"image/color"
)

// Содержимое первой вкладки "Расписание"
func content1() (C fyne.CanvasObject) {
	// два селектора и кнопка
	select1 = &widget.Select{
		Selected:  Last1,
		Options:   allCities(),
	}
	select2 = &widget.Select{
		Selected:  Last2,
		Options:   allCities(),
	}

	mainButton := &widget.Button{
		Text:     "Намутить",
		Style:    widget.PrimaryButton,
		Icon:     theme.DocumentCreateIcon(),
		OnTapped: func() {
			// незачем что-то делать, если в селекторах станции не менялись
			s1,s2 := select1.Selected, select2.Selected
			if !(s1 == Last1 && s2 == Last2) {
				mainButtonClick(s1,s2,"train")
				Last1,Last2 = s1,s2
				saveLastCities(s1,s2)
			}
		},
	}

	// когда приложение только что запущено и в селекторах последние выбранные станции
	// нужно чтобы всё же кнопка срабатывала
	Last1 = ""

	////////////////////////////////////////////////////////

	// миниатюра диаграммы и подпись
	resultImageWidget = &canvas.Image{
		FillMode:     canvas.ImageFillOriginal,
		ScaleMode:    canvas.ImageScalePixels,
	}
	resultImageWidget.Resize(fyne.NewSize(320,180))
	resultText = &canvas.Text{
		Color: color.RGBA{242, 78, 124, 255},
		Text: "",
		Alignment: fyne.TextAlignCenter,
		TextSize: 16,
		TextStyle: fyne.TextStyle{Bold: true},
	}
	resultBox = newTappableBox(widget.NewVBox(resultImageWidget, resultText))
	resultBox.Hide()

	box := widget.NewVBox(
		widget.NewHBox(select1, select2),
		mainButton,
	)
	C = widget.NewVBox(
		widget.NewHBox(layout.NewSpacer(),box,layout.NewSpacer()),
		layout.NewSpacer(),
		widget.NewVBox(layout.NewSpacer(), resultBox,layout.NewSpacer()),
		layout.NewSpacer(),
	)

	return
}