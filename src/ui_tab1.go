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
	select1 := &widget.Select{
		Selected: PastCity1,
		Options:  allCities(),
	}
	select2 := &widget.Select{
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

	box := widget.NewVBox(
		widget.NewHBox(select1, select2),
		mainButton,
	)

	////////////////////////////////////////////////////////

	// миниатюра диаграммы и подпись
	diagramWidget = &canvas.Image{
		FillMode:     canvas.ImageFillOriginal,
		ScaleMode:    canvas.ImageScalePixels,
	}
	diagramWidget.Resize(fyne.NewSize(320,180))
	diagramName = &canvas.Text{
		Color: color.RGBA{242, 78, 124, 255},
		Text: "",
		Alignment: fyne.TextAlignCenter,
		TextSize: 16,
		TextStyle: fyne.TextStyle{Bold: true},
	}
	diagramBox = newTappableBox(widget.NewVBox(diagramWidget, diagramName), deployDiagram)
	diagramBox.Hide()

	C = widget.NewVBox(
		widget.NewLabel(""),
		widget.NewHBox(layout.NewSpacer(), mapBox, layout.NewSpacer()),
		widget.NewLabel(""),
		widget.NewHBox(layout.NewSpacer(), box, layout.NewSpacer()),
		widget.NewLabel(""),
		layout.NewSpacer(),
		widget.NewVBox(layout.NewSpacer(), diagramBox, layout.NewSpacer()),
		widget.NewLabel(""),
		layout.NewSpacer(),
	)
	return
}