package src

import (
	"bytes"
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/anthonynsimon/bild/transform"
	"github.com/onix3/train-timetable/resource"
	"image"
	"image/color"
)

// Содержимое первой вкладки "Расписание"
func content1() (C fyne.CanvasObject) {
	// миниатюра карты
	mapImage, _, _ = image.Decode(bytes.NewReader(resource.BelarusJpg.StaticContent))
	bigSize := int(float64(ScreenHeight)*0.875)
	mapImage = transform.Resize(mapImage,2362*bigSize/2100,bigSize, transform.Linear)
	size := 200
	mapImageWidget := &canvas.Image{
		Image:        transform.Resize(mapImage,size,size, transform.Linear),
		FillMode:     canvas.ImageFillOriginal,
		ScaleMode:    canvas.ImageScalePixels,
	}
	mapImageWidget.Resize(fyne.NewSize(size, size))
	mapBox := newTappableBox(
		widget.NewVBox(mapImageWidget),
		deployMap,
	)

	////////////////////////////////////////////////////////

	// два селектора и кнопка
	select1 = &widget.Select{
		Selected: PastCity1,
		Options:  allCities(),
	}
	select2 = &widget.Select{
		Selected: PastCity2,
		Options:  allCities(),
	}

	mainButton = &widget.Button{
		Text:     "Намутить",
		Style:    widget.PrimaryButton,
		Icon:     theme.DocumentCreateIcon(),
		OnTapped: func() {
			mainButtonClick(select1.Selected, select2.Selected,"train")
		},
	}

	// когда приложение только что запущено и в селекторах последние выбранные станции
	// нужно чтобы всё же кнопка срабатывала
	PastCity1 = ""

	box := widget.NewVBox(
		widget.NewHBox(select1, select2),
		mainButton,
	)

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
	resultBox = newTappableBox(widget.NewVBox(resultImageWidget, resultText), deployResult)
	resultBox.Hide()

	C = widget.NewVBox(
		widget.NewLabel(""),
		widget.NewHBox(layout.NewSpacer(), mapBox, layout.NewSpacer()),
		widget.NewLabel(""),
		widget.NewHBox(layout.NewSpacer(), box, layout.NewSpacer()),
		layout.NewSpacer(),
		widget.NewVBox(layout.NewSpacer(), resultBox, layout.NewSpacer()),
		layout.NewSpacer(),
	)
	return
}